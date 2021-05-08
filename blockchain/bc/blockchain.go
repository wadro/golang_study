package bc

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

const (
	dbPath = "./tmp/blocks"
)

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.WithDir(dbPath)
	opts.WithValueDir(dbPath)

	// opts.Dir = dbPath
	// opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound { // lh -> last hash
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved?")
			err = txn.Set(genesis.Hash, genesis.Serialize()) // hash: key for the genesis block
			Handle(err)                                      // only for the err 🠕(uparrow)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err) // only for the err 🠕(uparrow)

			err = item.Value(func(b []byte) error {
				lastHash = b
				return err
			})
			return err
		}
	})

	Handle(err) // db.Update function's err

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)

		// lastHash, err = item.Value()
		err = item.Value(func(b []byte) error {
			lastHash = b
			return err
		})

		return err
	})
	Handle(err) // db.View err

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)

		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err) // db.Update err
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err) //  only for the err 🠕(uparrow)

		// encodedBlock, err := item.Value()
		err = item.Value(func(b []byte) error {
			encodedBlock = b
			return err
		})

		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err) // db.View err

	iter.CurrentHash = block.PrevHash

	return block
}
