# chain5j-solc

通过Go调用`soljson.js`来实现合约编译。

`soljson.js`通过[solc-bin repository](https://github.com/ethereum/solc-bin) 可获取。

## 使用

Example:

```go
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
```

## LICENSE
Please refer to [LICENSE](LICENSE) file.

Copyright@2021 chain5j