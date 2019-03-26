package chapter2

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain{
	return &Blockchain{Blocks:[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data string){
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data,preBlock.Hash)
	bc.Blocks = append(bc.Blocks,block)
}