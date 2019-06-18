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
	"fmt"
	"github.com/assetsadapterstore/tivalue-adapter/tivalue_addrdec"
	"github.com/blocktree/go-owcrypt"

)

type AddressDecoder struct {
	wm *WalletManager //钱包管理者
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *AddressDecoder {
	decoder := AddressDecoder{}
	decoder.wm = wm
	return &decoder
}

//PrivateKeyToWIF 私钥转WIF
func (decoder *AddressDecoder) PrivateKeyToWIF(priv []byte, isTestnet bool) (string, error) {
	return "", fmt.Errorf("PrivateKeyToWIF not implemented")
}

//PublicKeyToAddress 公钥转地址
func (decoder *AddressDecoder) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {
	tivalue_addrdec.Default.IsTestNet = decoder.wm.Config.IsTestNet
	xpub, err := owcrypt.CURVE25519_convert_Ed_to_X(pub)
	if err != nil {
		return "", fmt.Errorf("Invalid public key!")
	}

	pkHash := owcrypt.Hash(owcrypt.Hash(xpub, 32, owcrypt.HASH_ALG_BLAKE2B), 32, owcrypt.HASH_ALG_KECCAK256)[:20]
	return tivalue_addrdec.Default.AddressEncode(pkHash)
}

//RedeemScriptToAddress 多重签名赎回脚本转地址
func (decoder *AddressDecoder) RedeemScriptToAddress(pubs [][]byte, required uint64, isTestnet bool) (string, error) {
	return "", fmt.Errorf("WIFToPrivateKey not implemented")
}

//WIFToPrivateKey WIF转私钥
func (decoder *AddressDecoder) WIFToPrivateKey(wif string, isTestnet bool) ([]byte, error) {
	return nil, fmt.Errorf("WIFToPrivateKey not implemented")
}
