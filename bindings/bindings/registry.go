package bindings

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/wemixkanvas/kanvas/bindings/solc"
)

var layouts = make(map[string]*solc.StorageLayout)

var deployedBytecodes = make(map[string]string)

func GetStorageLayout(name string) (*solc.StorageLayout, error) {
	layout := layouts[name]
	if layout == nil {
		return nil, errors.New("storage layout not found")
	}
	return layout, nil
}

func GetDeployedBytecode(name string) ([]byte, error) {
	bc := deployedBytecodes[name]
	if bc == "" {
		return nil, fmt.Errorf("deployed bytecode %s not found", name)
	}

	return common.FromHex(bc), nil
}
