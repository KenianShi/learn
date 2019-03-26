package chapter1

type Blockchain struct {
	Blocks 		[]*Block
}

func (bc *Blockchain) AddBlock(data string){
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].Hash
	newBlock := NewBlock(data,prevBlockHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",[]byte{})
}

func NewBlockchain()*Blockchain{
	return &Blockchain{Blocks:[]*Block{NewGenesisBlock()}}
}