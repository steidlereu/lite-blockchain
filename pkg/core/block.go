/*
Copyright 2022 steidler.eu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
*/

package core

import "hash"

type Block struct {
	Hash         hash.Hash64
	PreviousHash hash.Hash64
	Payload      interface{}
	timestamp    int64
	nonce        int
}
