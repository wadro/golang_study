package bc

type BlockChain struct {
	Blocks []*Block
}
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0} // empty slice of byte, data string converted into slice of byte, previous Hash
	// block.DeriveHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
