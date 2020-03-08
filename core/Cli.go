package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
	Bc *blockChain
}

func (cli *Cli) printUsage() {
	fmt.Println("Usage")
	fmt.Println("addblock -data BLOCK_DATA - add Block to the blockchain")
	fmt.Println("printchain - print all the blocks of the blockchain")
}

func (cli *Cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *Cli) addBlock(data string) {
	cli.Bc.AppendBlock(data)
	fmt.Println("Success!")
}

func (cli *Cli) printChains() {
	cli.Bc.Iterator()
}

/**
逻辑控制代码
*/
func (cli *Cli) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addbock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChains()
	}
}
