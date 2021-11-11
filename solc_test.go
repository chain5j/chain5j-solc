package solc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestSloc_0_6_10(t *testing.T) {
	compiler, err := NewFromFile(filepath.Join("./solc-bin", "soljson-v0.6.10+commit.00c0fcaf.js"))
	if err != nil {
		t.Fatal(err)
	}
	solData, err := ioutil.ReadFile("./contracts/NFT-0.6.10.sol")
	if err != nil {
		t.Fatal(err)
	}
	input := &Input{
		Language: "Solidity",
		Sources: map[string]SourceIn{
			"NFT.sol": SourceIn{
				Content: string(solData),
			},
		},
		Settings: DefaultSetting,
	}

	output, _ := compiler.Compile(input)
	fmt.Printf("Bytecode: %v", output.Contracts["NFT.sol"]["ERC721Full"].EVM.Bytecode.Object)
}
func TestSloc_0_8_9(t *testing.T) {
	compiler, err := NewFromFile(filepath.Join("./solc-bin", "soljson-v0.8.9+commit.e5eed63a.js"))
	if err != nil {
		t.Fatal(err)
	}
	soljson, err := ioutil.ReadFile("./contracts/NFT.sol")
	if err != nil {
		t.Fatal(err)
	}
	input := &Input{
		Language: "Solidity",
		Sources: map[string]SourceIn{
			"NFT.sol": SourceIn{
				Content: string(soljson),
			},
		},
		Settings: DefaultSetting,
	}

	output, err := compiler.Compile(input)
	if err != nil {
		t.Fatal(err)
	}
	abis := output.Contracts["NFT.sol"]["ERC721Full"].ABI
	for _, abi := range abis {
		fmt.Printf("abi: %v", string(abi))
	}
	fmt.Printf("Bytecode: %v", output.Contracts["NFT.sol"]["ERC721Full"].EVM.Bytecode.Object)
}
