package chapter1

type Blockchain struct {
	Block []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Block[len(bc.Block)-1]
	block := NewBlock(data,preBlock.Hash)
	bc.Block = append(bc.Block, block)
}

func NewBlockchain() *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}