/*
Copyright 2022 steidler.eu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
*/

package main

import (
	"fmt"

	"steidler.eu/lite-blockchain/pkg/block"
	"steidler.eu/lite-blockchain/pkg/chain"
)

func main() {
	fmt.Println("");
	fmt.Println("Hello World");
	fmt.Println("");

	exampleBlock := block.Init(`{"name":"John", "age":30, "car":null}`, 4)
	exampleBlock.Mine()
	fmt.Println(exampleBlock.Valid());

	exampleBlock2 := block.Init(`{"name":"John", "age":30, "car":null}`, 4)
	fmt.Println(exampleBlock2.Valid());

	//exampleBlock3 := block.Init("Blabla", 4)
	//fmt.Println(exampleBlock3.Valid());

	exampleChain := chain.New()
	exampleChain.Add(`{"name":"John", "age":30, "car":null}`, 4)
	exampleChain.Add(`{"name":"John", "age":30, "car":null}`, 4)
	exampleChain.Add(`{"name":"John", "age":30, "car":null}`, 4)

	for _, c := range exampleChain.BlockChain {
        fmt.Println(c)
    }

	exampleChain.Valid()

	//fmt.Println(exampleChain);
}

