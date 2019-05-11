package blockchain

import (
	"bytes"
	"encoding/binary"
	"github.com/EnsicoinDevs/ensicoincoin/consensus"
	"github.com/EnsicoinDevs/ensicoincoin/network"
	"github.com/EnsicoinDevs/ensicoincoin/scripts"
	"github.com/EnsicoinDevs/ensicoincoin/utils"
	bolt "github.com/etcd-io/bbolt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/big"
	"path/filepath"
	"strconv"
	"sync"
)

type ValidationError struct {
	Reason string
}

func (err *ValidationError) Error() string {
	return "block invalid: " + err.Reason
}

type Config struct {
	DataDir       string
	OnPushedBlock func(*Block) error
	OnPoppedBlock func(*Block) error
}

type BlockchainStats struct {
	BestBlockHash   utils.Hash
	BestBlockHeight uint32
}

type Blockchain struct {
	config Config

	db *bolt.DB

	GenesisBlock *Block

	lock *sync.RWMutex

	stats *BlockchainStats

	orphans           map[utils.Hash]*Block
	prevHashToOrphans map[utils.Hash][]utils.Hash
	cache             map[utils.Hash]*network.BlockHeader
}

func NewBlockchain(config *Config) *Blockchain {
	return &Blockchain{
		config: *config,

		GenesisBlock: &genesisBlock,

		lock: &sync.RWMutex{},

		stats: &BlockchainStats{},

		orphans:           make(map[utils.Hash]*Block),
		prevHashToOrphans: make(map[utils.Hash][]utils.Hash),
		cache:             make(map[utils.Hash]*network.BlockHeader),
	}
}

var (
	blocksBucket    = []byte("blocks")
	heightsBucket   = []byte("heights")
	statsBucket     = []byte("stats")
	followingBucket = []byte("following")
	utxosBucket     = []byte("utxos")
	stxojBucket     = []byte("stxoj")
)

func (blockchain *Blockchain) Load() error {
	var err error

	blockchain.db, err = bolt.Open(filepath.Join(blockchain.config.DataDir, "data.db"), 0600, nil)
	if err != nil {
		return errors.Wrap(err, "error opening the blockchain database")
	}

	shouldStoreGenesisBlock := false

	blockchain.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(blocksBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the blocks bucket")
		}

		genesisBlockInDb := b.Get(genesisBlock.Hash()[:])
		if genesisBlockInDb == nil {
			shouldStoreGenesisBlock = true
		}

		b, err = tx.CreateBucketIfNotExists(statsBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the stats bucket")
		}

		bestBlockHashInDb := b.Get([]byte("bestBlockHash"))
		if bestBlockHashInDb == nil {
			b.Put([]byte("bestBlockHash"), genesisBlock.Hash()[:])
			b.Put([]byte("bestBlockHeight"), []byte{0, 0, 0, 0})
		} else {
			blockchain.stats.BestBlockHash = *utils.NewHash(bestBlockHashInDb)
			blockchain.stats.BestBlockHeight = binary.BigEndian.Uint32(b.Get([]byte("bestBlockHeight")))
		}

		_, err = tx.CreateBucketIfNotExists(heightsBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the heights bucket")
		}

		_, err = tx.CreateBucketIfNotExists(followingBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the following bucket")
		}

		_, err = tx.CreateBucketIfNotExists(utxosBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the utxos bucket")
		}

		_, err = tx.CreateBucketIfNotExists(stxojBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the stxoj bucket")
		}

		return nil
	})

	if shouldStoreGenesisBlock {
		blockchain.storeBlock(&genesisBlock)
	}

	return nil
}

func (blockchain *Blockchain) Stop() error {
	log.Debug("blockchain shutting down")
	defer log.Debug("blockchain shutdown complete")

	return blockchain.db.Close()
}

func (blockchain *Blockchain) fetchUtxos(tx *Tx) (*Utxos, []*network.Outpoint, error) {
	utxos := newUtxos()
	var missings []*network.Outpoint

	err := blockchain.db.View(func(btx *bolt.Tx) error {
		b := btx.Bucket(utxosBucket)

		for _, input := range tx.Msg.Inputs {
			outpoint := input.PreviousOutput
			outpointBytes, err := outpoint.MarshalBinary()
			if err != nil {
				return errors.Wrap(err, "error marshaling an outpoint")
			}

			utxoBytes := b.Get(outpointBytes)
			if utxoBytes == nil {
				missings = append(missings, outpoint)
				continue
			}

			utxo := new(UtxoEntry)
			err = utxo.UnmarshalBinary(utxoBytes)
			if err != nil {
				return err
			}

			utxos.AddEntry(outpoint, utxo)
		}

		return nil
	})

	return utxos, missings, err
}

func (blockchain *Blockchain) FetchUtxos(tx *Tx) (*Utxos, []*network.Outpoint, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.fetchUtxos(tx)
}

func (blockchain *Blockchain) fetchAllUtxos() (*Utxos, error) {
	utxos := newUtxos()

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(utxosBucket)

		b.ForEach(func(outpointBytes, utxoBytes []byte) error {
			outpoint := new(network.Outpoint)
			err := outpoint.UnmarshalBinary(outpointBytes)
			if err != nil {
				return err
			}

			utxo := new(UtxoEntry)
			err = utxo.UnmarshalBinary(utxoBytes)
			if err != nil {
				return err
			}

			utxos.AddEntry(outpoint, utxo)

			return nil
		})

		return nil
	})

	return utxos, err
}

func (blockchain *Blockchain) FetchAllUtxos() (*Utxos, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.fetchAllUtxos()
}

func (blockchain *Blockchain) storeBlock(block *Block) error {
	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(blocksBucket)

		blockBytes, err := block.MarshalBinary()
		if err != nil {
			return err
		}

		b.Put(block.Hash()[:], blockBytes)

		return nil
	})

	return err
}

func (blockchain *Blockchain) findBlockByHash(hash *utils.Hash) (*Block, error) {
	var block *Block

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(blocksBucket)

		blockBytes := b.Get(hash[:])
		if blockBytes == nil {
			return nil
		}

		block = NewBlock()

		block.UnmarshalBinary(blockBytes)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error finding a block")
	}

	return block, nil
}

func (blockchain *Blockchain) FindBlockByHash(hash *utils.Hash) (*Block, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.findBlockByHash(hash)
}

func (blockchain *Blockchain) findBlockHeaderByHash(hash *utils.Hash) (*network.BlockHeader, error) {
	header, ok := blockchain.cache[*hash]
	if ok {
		return header, nil
	}

	block, err := blockchain.findBlockByHash(hash)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}

	header = block.Msg.Header

	blockchain.cache[*hash] = header

	return header, nil
}

func (blockchain *Blockchain) FindBlockHeaderByHash(hash *utils.Hash) (*network.BlockHeader, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.findBlockHeaderByHash(hash)
}

func (blockchain *Blockchain) findBestBlock() (*Block, error) {
	var bestBlockHash *utils.Hash

	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(statsBucket)
		bestBlockHash = utils.NewHash(b.Get([]byte("bestBlockHash")))

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error finding the best block hash")
	}

	block, err := blockchain.findBlockByHash(bestBlockHash)
	if err != nil {
		return nil, errors.Wrap(err, "error finding the best block")
	}

	return block, nil
}

func (blockchain *Blockchain) FindBestBlock() (*Block, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.findBestBlock()
}

func (blockchain *Blockchain) findBlockHashesStartingAt(hash *utils.Hash) ([]*utils.Hash, error) {
	var hashes []*utils.Hash

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(followingBucket)

		for {
			hashBytes := b.Get(hash[:])
			if hashBytes == nil {
				break
			}

			hash = utils.NewHash(hashBytes)

			hashes = append(hashes, hash)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(nil, "error finding all the hashes")
	}

	return hashes, nil
}

func (blockchain *Blockchain) FindBlockHashesStartingAt(hash *utils.Hash) ([]*utils.Hash, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.findBlockHashesStartingAt(hash)
}

func (blockchain *Blockchain) findAncestorHeader(block *Block, nth uint32) (*network.BlockHeader, error) {
	current := block.Msg.Header
	targetHeight := block.Msg.Header.Height - nth

	var err error

	for current.Height != targetHeight {
		current, err = blockchain.findBlockHeaderByHash(current.HashPrevBlock) // TODO: maybe there is a race condition here
		if err != nil {
			return nil, err
		}
	}

	return current, nil
}

func (blockchain *Blockchain) calcNextBlockDifficulty(block *Block, nextBlock *Block) (uint32, error) {
	if nextBlock.Msg.Header.Height < consensus.BLOCKS_PER_RETARGET {
		return block.Msg.Header.Bits, nil
	}

	lastRetargetBlockHeader, err := blockchain.findAncestorHeader(nextBlock, consensus.BLOCKS_PER_RETARGET)
	if err != nil {
		return 0, err
	}

	realizedTimespan := uint64(nextBlock.Msg.Header.Timestamp.Unix() - lastRetargetBlockHeader.Timestamp.Unix())
	if realizedTimespan < consensus.MIN_RETARGET_TIMESPAN {
		realizedTimespan = consensus.MIN_RETARGET_TIMESPAN
	} else if realizedTimespan > consensus.MAX_RETARGET_TIMESPAN {
		realizedTimespan = consensus.MAX_RETARGET_TIMESPAN
	}

	targetBefore := utils.BitsToBig(block.Msg.Header.Bits)
	target := new(big.Int).Mul(targetBefore, big.NewInt(int64(realizedTimespan)))
	target.Div(target, big.NewInt(consensus.BLOCKS_MEAN_TIMESPAN*consensus.BLOCKS_PER_RETARGET))

	return utils.BigToBits(target), nil
}

func (blockchain *Blockchain) CalcNextBlockDifficulty(block *Block, nextBlock *Block) (uint32, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.calcNextBlockDifficulty(block, nextBlock)
}

func (blockchain *Blockchain) validateBlock(block *Block) (error, error) {
	parentBlock, err := blockchain.findBlockByHash(block.Msg.Header.HashPrevBlock)
	if parentBlock.Msg.Header.Height+1 != block.Msg.Header.Height {
		return errors.New("height is not equal to the previous block height + 1"), nil
	}

	nextBits, err := blockchain.calcNextBlockDifficulty(parentBlock, block)
	if err != nil {
		return nil, err
	}
	if block.Msg.Header.Bits != nextBits {
		return errors.New("the difficulty is invalid"), nil
	}

	if block.Txs[0].Msg.Flags[0] != strconv.Itoa(int(block.Msg.Header.Height)) {
		return errors.New("the first flag of the coinbase is not the block height"), nil
	}

	var totalFees uint64

	for _, tx := range block.Txs[1:] { // TODO: huuuum
		utxos, notfound, err := blockchain.fetchUtxos(tx)
		if err != nil {
			return nil, err
		}
		if len(notfound) != 0 {
			return errors.New("a tx try to spend an spent / unknown tx"), nil
		}

		valid, fees, err := blockchain.validateTx(tx, block, utxos)
		if err == nil {
			return nil, nil
		}
		if !valid {
			return errors.New("a tx is invalid"), nil
		}

		totalFees += fees
	}

	var totalCoinbaseAmount uint64

	for _, output := range block.Txs[0].Msg.Outputs {
		totalCoinbaseAmount += output.Value
	}

	if totalCoinbaseAmount != totalFees+block.CalcBlockSubsidy() {
		return errors.New("coinbase value is invalid"), nil
	}

	return nil, nil
}

func (blockchain *Blockchain) ValidateBlock(block *Block) (error, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.validateBlock(block)
}

func (blockchain *Blockchain) validateTx(tx *Tx, block *Block, utxos *Utxos) (bool, uint64, error) {
	var inputsSum uint64

	for _, input := range tx.Msg.Inputs {
		utxo := utxos.FindEntry(input.PreviousOutput)

		if utxo.IsCoinBase() {
			if block.Msg.Header.Height-utxo.blockHeight < consensus.COINBASE_MATURITY {
				return false, 0, nil
			}
		}

		inputsSum += utxo.Amount()
	}

	var outputsSum uint64

	for _, output := range tx.Msg.Outputs {
		outputsSum += output.Value
	}

	if inputsSum <= outputsSum {
		return false, 0, nil
	}

	for _, input := range tx.Msg.Inputs {
		utxo := utxos.FindEntry(input.PreviousOutput)
		script := scripts.NewScript(tx.Msg, input, utxo.Amount(), utxo.Script(), input.Script)
		valid, err := script.Validate()
		if err != nil {
			return false, 0, err
		}
		if !valid {
			return false, 0, nil
		}
	}

	return false, inputsSum - outputsSum, nil
}

func (blockchain *Blockchain) Validatetx(tx *Tx, block *Block, utxos *Utxos) (bool, uint64, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.Validatetx(tx, block, utxos)
}

func (blockchain *Blockchain) storeUtxos(utxos *Utxos) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(utxosBucket)
		if err != nil {
			return err
		}

		b, err := tx.CreateBucket(utxosBucket)
		if err != nil {
			return err
		}

		for outpoint, entry := range utxos.entries {
			outpointBytes, err := outpoint.MarshalBinary()
			if err != nil {
				return err
			}

			entryBytes, err := entry.MarshalBinary()
			if err != nil {
				return err
			}

			err = b.Put(outpointBytes, entryBytes)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (blockchain *Blockchain) storeFollowing(block *utils.Hash, nextBlock *utils.Hash) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(followingBucket)

		return b.Put(block[:], nextBlock[:])
	})
}

func (blockchain *Blockchain) removeFollowing(block *utils.Hash) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(followingBucket)

		return b.Delete(block[:])
	})
}

func (blockchain *Blockchain) removeStxojEntry(blockHash *utils.Hash) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(stxojBucket)

		return b.Delete(blockHash[:])
	})
}

func (blockchain *Blockchain) storeStxojEntry(blockHash *utils.Hash, entry []*UtxoEntry) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(stxojBucket)

		buf := bytes.NewBuffer(nil)

		for _, stxo := range entry {
			err := WriteUtxoEntry(buf, stxo)
			if err != nil {
				return err
			}
		}

		return b.Put(blockHash[:], buf.Bytes())
	})
}

func (blockchain *Blockchain) fetchStxojEntry(blockHash *utils.Hash) ([]*UtxoEntry, error) {
	var utxoEntries []*UtxoEntry

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(stxojBucket)

		buf := bytes.NewBuffer(b.Get(blockHash[:]))

		for buf.Len() != 0 {
			utxoEntry, err := ReadUtxoEntry(buf)
			if err != nil {
				return err
			}

			utxoEntries = append(utxoEntries, utxoEntry)
		}

		return nil
	})

	return utxoEntries, err
}

func (blockchain *Blockchain) FetchStxojEntry(blockHash *utils.Hash) ([]*UtxoEntry, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.fetchStxojEntry(blockHash)
}

func (blockchain *Blockchain) storeLongestChain(blockHash *utils.Hash) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(statsBucket)

		return b.Put([]byte("longestChain"), blockHash[:])
	})
}

func (blockchain *Blockchain) storeBlockHashAtHeight(blockHash *utils.Hash, height uint32) error {
	return blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(heightsBucket)

		heightBytes := make([]byte, 4)

		binary.BigEndian.PutUint32(heightBytes, height)

		return b.Put(heightBytes, blockHash[:])
	})
}

func (blockchain *Blockchain) getBlockHashAtHeight(height uint32) (*utils.Hash, error) {
	var hash *utils.Hash

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(heightsBucket)

		heightBytes := make([]byte, 4)

		binary.BigEndian.PutUint32(heightBytes, height)

		hash = utils.NewHash(b.Get(heightBytes))

		return nil
	})

	return hash, err
}

func (blockchain *Blockchain) GetBlockHashAtHeight(height uint32) (*utils.Hash, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.getBlockHashAtHeight(height)
}

func (blockchain *Blockchain) pushBlock(block *Block) error {
	longestChain, err := blockchain.findBestBlock()
	if err != nil {
		return err
	}

	if !longestChain.Hash().IsEqual(block.Msg.Header.HashPrevBlock) {
		return errors.New("block must extend the best chain")
	}

	utxos, err := blockchain.fetchAllUtxos()
	if err != nil {
		return err
	}

	var stxoj []*UtxoEntry

	for _, tx := range block.Txs {
		for _, input := range tx.Msg.Inputs {
			removedEntry := utxos.RemoveEntry(input.PreviousOutput)
			stxoj = append(stxoj, removedEntry)
		}

		for index, output := range tx.Msg.Outputs {
			utxos.AddEntry(&network.Outpoint{
				Hash:  *tx.Hash(),
				Index: uint32(index),
			}, &UtxoEntry{
				amount:      output.Value,
				script:      output.Script,
				blockHeight: block.Msg.Header.Height,
				coinBase:    tx.IsCoinBase(),
			})
		}
	}

	err = blockchain.storeLongestChain(block.Hash())
	if err != nil {
		return err
	}

	err = blockchain.storeFollowing(longestChain.Hash(), block.Hash())
	if err != nil {
		return err
	}

	err = blockchain.storeStxojEntry(block.Hash(), stxoj)
	if err != nil {
		return err
	}

	err = blockchain.storeUtxos(utxos)
	if err != nil {
		return err
	}

	err = blockchain.storeBlockHashAtHeight(block.Hash(), block.Msg.Header.Height)
	if err != nil {
		return err
	}

	err = blockchain.storeBlockHashAtHeight(block.Hash(), block.Msg.Header.Height)
	if err != nil {
		return err
	}

	err = blockchain.storeFollowing(block.Msg.Header.HashPrevBlock, block.Hash())
	if err != nil {
		return err
	}

	if blockchain.config.OnPushedBlock != nil {
		blockchain.config.OnPushedBlock(block)
	}

	return nil
}

func (blockchain *Blockchain) popBlock(block *Block) error {
	longestChain, err := blockchain.findBestBlock()
	if err != nil {
		return err
	}

	if longestChain.Hash() != block.Hash() {
		return errors.New("block must be the best block")
	}

	utxos, err := blockchain.fetchAllUtxos()
	if err != nil {
		return err
	}

	stxoj, err := blockchain.fetchStxojEntry(block.Hash())
	if err != nil {
		return err
	}

	var currentStxoId uint64

	for _, tx := range block.Txs {
		for index := range tx.Msg.Outputs {
			utxos.RemoveEntry(&network.Outpoint{
				Hash:  *tx.Hash(),
				Index: uint32(index),
			})
		}

		for _, input := range tx.Msg.Inputs {
			stxo := stxoj[currentStxoId]
			currentStxoId++

			entry := utxos.FindEntry(input.PreviousOutput)
			if entry == nil {
				entry = newUtxoEntry()
				utxos.AddEntry(input.PreviousOutput, entry)
			}

			entry.amount = stxo.amount
			entry.script = stxo.script
			entry.blockHeight = stxo.blockHeight
			entry.coinBase = stxo.coinBase
		}
	}

	err = blockchain.storeUtxos(utxos)
	if err != nil {
		return err
	}

	err = blockchain.removeStxojEntry(block.Hash())
	if err != nil {
		return err
	}

	err = blockchain.removeFollowing(block.Msg.Header.HashPrevBlock)
	if err != nil {
		return err
	}

	err = blockchain.storeLongestChain(block.Msg.Header.HashPrevBlock)
	if err != nil {
		return err
	}

	blockchain.config.OnPoppedBlock(block)

	return nil
}

func (blockchain *Blockchain) findForkingBlock(blockA *Block, blockB *Block) (*Block, error) {
	var err error

	// If blockB height is superior to blockA height, we walk back to the blockA height
	for blockB.Msg.Header.Height > blockA.Msg.Header.Height {
		blockB, err = blockchain.findBlockByHash(blockB.Msg.Header.HashPrevBlock)
		if err != nil {
			return nil, err
		}
	}

	// Same with blockA
	for blockA.Msg.Header.Height > blockB.Msg.Header.Height {
		blockA, err = blockchain.findBlockByHash(blockB.Msg.Header.HashPrevBlock)
		if err != nil {
			return nil, err
		}
	}

	// While the two blocks are different, we walk back of one block
	for blockA != nil && !blockB.Hash().IsEqual(blockA.Hash()) {
		blockA, err = blockchain.findBlockByHash(blockA.Msg.Header.HashPrevBlock)
		if err != nil {
			return nil, err
		}

		blockB, err = blockchain.findBlockByHash(blockB.Msg.Header.HashPrevBlock)
		if err != nil {
			return nil, err
		}
	}

	return blockA, nil
}

func (blockchain *Blockchain) acceptBlock(block *Block) (bool, error) {
	err := blockchain.storeBlock(block)
	if err != nil {
		return false, err
	}

	bestBlock, err := blockchain.findBestBlock()
	if err != nil {
		return false, err
	}

	if bestBlock.Hash().IsEqual(block.Msg.Header.HashPrevBlock) {
		err = blockchain.pushBlock(block)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	forkingBlock, err := blockchain.findForkingBlock(bestBlock, block)
	if err != nil {
		return false, err
	}
	if forkingBlock == nil {
		return false, errors.New("can't find the forking block")
	}

	return true, nil
}

func (blockchain *Blockchain) addOrphan(block *Block) {
	blockchain.orphans[*block.Hash()] = block
	blockchain.prevHashToOrphans[*block.Msg.Header.HashPrevBlock] = append(blockchain.prevHashToOrphans[*block.Msg.Header.HashPrevBlock], *block.Hash())
}

func (blockchain *Blockchain) processOrphans(acceptedBlock *Block) error {
	for _, blockHash := range blockchain.prevHashToOrphans[*acceptedBlock.Hash()] {
		block, err := blockchain.findOrphan(&blockHash)
		if err != nil {
			return err
		}

		valid, err := blockchain.validateBlock(block)
		if err != nil {
			return err
		}
		if valid != nil {
			blockchain.removeOrphan(block.Hash())
		}

		_, err = blockchain.acceptBlock(block)
		if err != nil {
			return err
		}

		err = blockchain.processOrphans(block)
		if err != nil {
			return err
		}
	}

	return nil
}

func (blockchain *Blockchain) findOrphan(hash *utils.Hash) (*Block, error) {
	return blockchain.orphans[*hash], nil
}

func (blockchain *Blockchain) FindOrphan(hash *utils.Hash) (*Block, error) {
	blockchain.lock.RLock()
	defer blockchain.lock.RUnlock()

	return blockchain.findOrphan(hash)
}

func (blockchain *Blockchain) removeOrphan(hash *utils.Hash) {
	for _, blockHash := range blockchain.prevHashToOrphans[*hash] {
		blockchain.removeOrphan(&blockHash)
	}

	delete(blockchain.prevHashToOrphans, *hash)
	delete(blockchain.orphans, *hash)
}

func (blockchain *Blockchain) ProcessBlock(block *Block) (error, error) {
	blockchain.lock.Lock()
	defer blockchain.lock.Unlock()

	header, err := blockchain.findBlockHeaderByHash(block.Hash())
	if err != nil {
		return nil, err
	}
	if header != nil {
		return &ValidationError{"a block with the same hash exist"}, nil
	}

	orphanBlock, err := blockchain.findOrphan(block.Hash())
	if err != nil {
		return nil, nil
	}
	if orphanBlock != nil {
		return &ValidationError{"an orphan block with the same hash exist"}, nil
	}

	if valid := block.IsSane(); valid != nil {
		return errors.Wrap(valid, "the block is not sane"), nil
	}

	header, err = blockchain.findBlockHeaderByHash(block.Msg.Header.HashPrevBlock)
	if err != nil {
		return nil, err
	}
	if header == nil {
		log.WithField("hash", block.Msg.Header.Hash().String()).Info("accepted orphan block")

		blockchain.addOrphan(block)

		return nil, nil
	}

	valid, err := blockchain.validateBlock(block)
	if err != nil {
		return nil, err
	}
	if valid != nil {
		return errors.Wrap(valid, "the block is invalid"), nil
	}

	_, err = blockchain.acceptBlock(block)
	if err != nil {
		return nil, err
	}

	err = blockchain.processOrphans(block)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
