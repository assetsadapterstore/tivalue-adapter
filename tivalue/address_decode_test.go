/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tivalue

import (
	"encoding/hex"
	"testing"
	"tivalue-adapter/tivalue_addrdec"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	hash, _ := hex.DecodeString("74de0805bd5eae0031a9f69140428bbfafed707c")
	address, _ := tivalue_addrdec.Default.AddressEncode(hash)
	t.Logf("address: %s", address)

}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	accountID := "tvBoZMpYv1uVt3JJRVhfXbTiPzV9vu3YSWo"
	pub, _ := tivalue_addrdec.Default.AddressDecode(accountID)
	t.Logf("hash: %s", hex.EncodeToString(pub))

}

func TestAddressDecoder_PublicKeyToAddress(t *testing.T) {
	pub, _ := hex.DecodeString("cc54844fbea5edd5620cb89dc489949685455c7e11facfd123dd50513681a1a3")
	address, _ := tw.GetAddressDecode().PublicKeyToAddress(pub, false)
	t.Logf("address: %s", address)
	//tvBoZMpYv1uVt3JJRVhfXbTiPzV9vu3YSWo
}