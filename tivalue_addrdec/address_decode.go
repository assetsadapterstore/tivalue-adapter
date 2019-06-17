package tivalue_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	alphabet = addressEncoder.VSYSAlphabet
)

var (

	//VSYS stuff
	TV_mainnetAddress = addressEncoder.AddressType{"base58", alphabet, "blake2b_and_keccak256_first_twenty", "blake2b_and_keccak256_first_twenty", 20, []byte{0x1D, 0x3B}, nil}
	TV_testnetAddress = addressEncoder.AddressType{"base58", alphabet, "blake2b_and_keccak256_first_twenty", "blake2b_and_keccak256_first_twenty", 20, []byte{0x1D, 0x54}, nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := TV_mainnetAddress
	if dec.IsTestNet {
		cfg = TV_testnetAddress
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := TV_mainnetAddress
	if dec.IsTestNet {
		cfg = TV_testnetAddress
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}
