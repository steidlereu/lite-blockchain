/*
Copyright 2022 steidler.eu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
*/

package chain

import (
	"fmt"

	"steidler.eu/lite-blockchain/pkg/block"
)


type Chain struct {
	BlockChain []block.Block
}

func New() *Chain {
	chain := Chain{}

	return &chain
}

func (chain *Chain) Add(payload string, raise int) {

	var tmpBlock *block.Block

	if chain.BlockChain == nil {
		tmpBlock = block.Init(payload, raise)
	} else {
		lastBlock := chain.BlockChain[len(chain.BlockChain)-1]
		tmpBlock = block.New(payload, raise, lastBlock.Hash)
	}

	tmpBlock.Mine()

	if tmpBlock.Valid() {
		chain.BlockChain = append(chain.BlockChain, *tmpBlock)
	}

}

func (chain *Chain) Valid() bool {

	isValid := true

	for i, c := range chain.BlockChain {
        fmt.Println(i,c)
    }

	return isValid
}