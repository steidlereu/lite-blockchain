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
	"steidler.eu/lite-blockchain/pkg/core"
)

func main() {
	fmt.Println("");
	fmt.Println("Hello World");
	fmt.Println("");

	test := core.Block{}
	test.Calculate()
}

