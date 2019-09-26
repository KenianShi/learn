package chapter2

type Blockchain struct {
	Blocks 		[]*Block
}

func (bc *Blockchain) AddBlock(data string){
	currentBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data,currentBlock.Hash)
	bc.Blocks = append(bc.Blocks, block)
}

func NewBlockchain()*Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
