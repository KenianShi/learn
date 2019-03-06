package main

import (
	"sync"
	"log"
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"os"
	"fmt"
)

type Block struct {
	Index 		int
	Timestamp 	string
	BPM 		int
	Hash 		string
	PrevHash 	string
}

var Blockchain []Block

type Message struct {
	BPM 		int
}

var mutex = &sync.Mutex{}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Println("load error")
	//	log.Fatal(err)
	//}

	go func(){
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{Index:0,Timestamp:t.String(),BPM:0,Hash:calculateHash(genesisBlock),PrevHash:""}
		spew.Dump(genesisBlock)
		mutex.Lock()
		Blockchain = append(Blockchain,genesisBlock)
		mutex.Unlock()
	}()

	log.Fatal(run())

}

func calculateHash(block Block) string{
	record := strconv.Itoa(block.Index)+block.Timestamp+strconv.Itoa(block.BPM)+block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func run()error{
	mux := makeMuRouter()
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port : ",httpPort)
	s := http.Server{
		Addr:			":8082",
		Handler:		mux,
		ReadTimeout:	time.Second * 10,
		WriteTimeout:	time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe();err != nil {
		return err
	}
	return nil
}

func makeMuRouter() http.Handler{
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/",handlerGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/",handlerWriteBlockchain).Methods("POST")
	return muxRouter
}

func handlerGetBlockchain(w http.ResponseWriter,r *http.Request){
	bytes,err := json.MarshalIndent(Blockchain,"","")
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func handlerWriteBlockchain(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var msg Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg);err != nil {
		respondWithJson(w,r,http.StatusBadRequest,r.Body)
		return
	}
	fmt.Printf("%s \n",msg)
	defer r.Body.Close()
	mutex.Lock()
	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(prevBlock,msg.BPM)

	if isBlockValid(newBlock,prevBlock){
		Blockchain = append(Blockchain,newBlock)
		spew.Dump(Blockchain)
	}
	mutex.Unlock()
	respondWithJson(w,r,http.StatusCreated,newBlock)
}

func isBlockValid(newBlock,oldBlock Block) bool{
	if newBlock.Index != oldBlock.Index + 1 {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash{
		return false
	}
	return true
}

func generateBlock(oldBlock Block,BPM int)Block{
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}


func respondWithJson(w http.ResponseWriter,r *http.Request,code int,payload interface{}){
	response,err := json.MarshalIndent(payload,"","")
	if err != nil {
		w.WriteHeader(http.StatusInsufficientStorage)
		w.Write([]byte("HTTP 500:Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}


