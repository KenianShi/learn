package main

import (
	"flag"
	"github.com/dgraph-io/badger"
	"fmt"
	"os"
	abci "github.com/tendermint/tendermint/abci/types"
	nm "github.com/tendermint/tendermint/node"
	cfg "github.com/tendermint/tendermint/config"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/proxy"
	"os/signal"
	"syscall"
	"shikenian/tendermint/kvstore"
)

var configFile string

func init() {
	flag.StringVar(&configFile,"config","/home/shikenian/.tendermint/config/config.toml","Path to config.toml")
}

func main() {
	db,err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		fmt.Fprintf(os.Stderr,"failed to open badger db: &v",err)
		os.Exit(1)
	}
	defer db.Close()

	app := kvstore.NewKVStoreApplication(db)
	fmt.Println("app创建成功")
	flag.Parse()

	node,err := newTendermint(app,configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v",err)
		os.Exit(2)
	}
	fmt.Println("node创建成功")
	node.Start()
	fmt.Println("node启动成功")
	defer func(){
		node.Stop()
		node.Wait()
	}()
	c := make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt,syscall.SIGTERM)
	<- c
	os.Exit(0)
}

func newTendermint(app abci.Application,configFile string)(*nm.Node,error){
	config := cfg.DefaultConfig()
	config.RootDir = filepath.Dir(filepath.Dir(configFile))
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil{
		return nil,errors.Wrap(err,"viper failed to read config file")
	}
	if err := viper.Unmarshal(config);err != nil {
		return nil,errors.Wrap(err,"viper failed to unmarshal config")
	}
	if err := config.ValidateBasic();err != nil {
		return nil,errors.Wrap(err,"config is invalid")
	}
	// create log
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	var err error
	logger,err = tmflags.ParseLogLevel(config.LogLevel,logger,cfg.DefaultLogLevel())
	if err != nil {
		return nil,errors.Wrap(err,"failed to parse log level")
	}

	fmt.Println("我也不知道自己在做什么")
	//fmt.Printf("%+v",config)
	fmt.Println(config.PrivValidatorKeyFile(),"++++++++++++",config.PrivValidatorState)
	pv := privval.LoadFilePV(config.PrivValidatorKeyFile(),config.PrivValidatorStateFile())  //todo

	nodeKey,err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil,errors.Wrap(err,"failed to load node's key")
	}
	fmt.Println("如果没有问题的话，接下来准备创建node")
	node,err := nm.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger,
	)
	if err != nil {
		return nil,errors.Wrap(err,"failed to create new Tendermint node")
	}
	return node,nil
}