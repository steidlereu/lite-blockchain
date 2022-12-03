/*
Copyright 2022 steidler.eu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
*/

package core

import (
	"crypto/sha512"
	"fmt"
	"hash"
	//"strconv"
	"time"
)

type ChainBlock interface {

}

type Block struct {
	Hash         hash.Hash
	PreviousHash hash.Hash
	Payload      string
	timestamp    time.Time
	nonce        int
}

func (b Block) Calculate() {

	sha512 := sha512.New()

	if b.PreviousHash != nil {
		sha512.Write(b.PreviousHash.Sum(nil))
	}
	//sha512.Write([]byte(b.timestamp.String()))
	//sha512.Write([]byte(strconv.Itoa(b.nonce)))
	//sha512.Write([]byte(b.Payload))

	sha512.Write([]byte("Hello"))
	sha512.Write([]byte("World"))

	fmt.Printf("%x", sha512.Sum(nil))

}