/*
Copyright 2022 steidler.eu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
*/

package block

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Hash         hash.Hash
	PreviousHash hash.Hash
	Payload      string
	Timestamp    time.Time
	Nonce        int
	Complexity   int
}

func Init(payload string, complexity int) *Block {

	if !json.Valid([]byte(payload)) {
		panic("Block payload is not a valid json")
	}

	block := Block{}
	block.Timestamp = time.Now()
	block.Nonce = 0
	block.Payload = payload
	block.Complexity = complexity

	return &block
}

func New(payload string, complexity int, previousHash hash.Hash) *Block {
	block := Init(payload, complexity)
	block.PreviousHash = previousHash

	return block
}

func (block *Block) Calculate() []byte {
	block.Hash = sha512.New()

	if block.PreviousHash != nil {
		block.Hash.Write(block.PreviousHash.Sum(nil))
	}
	block.Hash.Write([]byte(block.Timestamp.String()))
	block.Hash.Write([]byte(strconv.Itoa(block.Nonce)))
	block.Hash.Write([]byte(block.Payload))

	return block.Hash.Sum(nil)
}

func (block *Block) Mine() {
	hashResult := block.Calculate()
	complexityZero := strings.Repeat("0", block.Complexity)

	for hex.EncodeToString(hashResult)[0:block.Complexity] != complexityZero {
		block.Nonce++
		hashResult = block.Calculate()
	}
}

func (block *Block) Valid() bool {
	if block.Hash == nil {
		return false
	}

	hashResult := block.Hash.Sum(nil)
	complexityZero := strings.Repeat("0", block.Complexity)

	return hex.EncodeToString(hashResult)[0:block.Complexity] == complexityZero
}

func (block *Block) HashToString(h hash.Hash) string {
	if h == nil {
		return "none"
	}

	hashResult := h.Sum(nil)
	return hex.EncodeToString(hashResult)
}

func (block Block) String() string {
	return fmt.Sprintf("\nBlock\nPayload: %v\nTime: %v\nComplexity: %v\nNonce: %v\nPrev. Hash: %v\nCurrent Hash: %v", block.Payload, block.Timestamp.String(), block.Complexity, block.Nonce, block.HashToString(block.PreviousHash), block.HashToString(block.Hash))
}
