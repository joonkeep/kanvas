// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/wemixkanvas/kanvas/bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_uint256)1001_storage\"}],\"types\":{\"t_array(t_uint256)1001_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c80637f46ddb2116100505780637f46ddb2146100bd578063927ede2d14610109578063aa5574521461013057600080fd5b80633687011a1461007757806354fd4d501461008c578063761f4493146100aa575b600080fd5b61008a61008536600461116e565b610143565b005b6100946101ef565b6040516100a1919061126b565b60405180910390f35b61008a6100b836600461127e565b610292565b6100e47f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100a1565b6100e47f000000000000000000000000000000000000000000000000000000000000000081565b61008a61013e366004611316565b6107f9565b333b156101d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6101e786863333888888886108b5565b505050505050565b606061021a7f0000000000000000000000000000000000000000000000000000000000000000610e53565b6102437f0000000000000000000000000000000000000000000000000000000000000000610e53565b61026c7f0000000000000000000000000000000000000000000000000000000000000000610e53565b60405160200161027e9392919061138d565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156103b057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610374573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103989190611403565b73ffffffffffffffffffffffffffffffffffffffff16145b61043c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016101ce565b3073ffffffffffffffffffffffffffffffffffffffff8816036104e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016101ce565b61050b877f74259ebf00000000000000000000000000000000000000000000000000000000610f90565b610597576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e740000000000000000000060648201526084016101ce565b8673ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611403565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146106e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204b616e766173204d696e7461626c6520455243373231206c6f60648201527f63616c20746f6b656e0000000000000000000000000000000000000000000000608482015260a4016101ce565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b15801561075657600080fd5b505af115801561076a573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107e89493929190611469565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff851661089c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016101ce565b6108ac87873388888888886108b5565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff8716610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016101ce565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa1580156109c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e79190611403565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610aa1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e6572000060648201526084016101ce565b60008873ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b129190611403565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610bcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c756500000000000000000060648201526084016101ce565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610c3f57600080fd5b505af1158015610c53573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610c7e97969594939291906114a9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610d93907f00000000000000000000000000000000000000000000000000000000000000009085908a90600401611506565b600060405180830381600087803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610e3f9493929190611469565b60405180910390a450505050505050505050565b606081600003610e9657505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610ec05780610eaa8161157a565b9150610eb99050600a836115e1565b9150610e9a565b60008167ffffffffffffffff811115610edb57610edb6115f5565b6040519080825280601f01601f191660200182016040528015610f05576020820181803683370190505b5090505b8415610f8857610f1a600183611624565b9150610f27600a8661163b565b610f3290603061164f565b60f81b818381518110610f4757610f47611667565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610f81600a866115e1565b9450610f09565b949350505050565b6000610f9b83610fb3565b8015610fac5750610fac8383611018565b9392505050565b6000610fdf827f01ffc9a700000000000000000000000000000000000000000000000000000000611018565b80156110125750611010827fffffffff00000000000000000000000000000000000000000000000000000000611018565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156110d0575060208210155b80156110dc5750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461110957600080fd5b50565b803563ffffffff8116811461112057600080fd5b919050565b60008083601f84011261113757600080fd5b50813567ffffffffffffffff81111561114f57600080fd5b60208301915083602082850101111561116757600080fd5b9250929050565b60008060008060008060a0878903121561118757600080fd5b8635611192816110e7565b955060208701356111a2816110e7565b9450604087013593506111b76060880161110c565b9250608087013567ffffffffffffffff8111156111d357600080fd5b6111df89828a01611125565b979a9699509497509295939492505050565b60005b8381101561120c5781810151838201526020016111f4565b8381111561121b576000848401525b50505050565b600081518084526112398160208601602086016111f1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610fac6020830184611221565b600080600080600080600060c0888a03121561129957600080fd5b87356112a4816110e7565b965060208801356112b4816110e7565b955060408801356112c4816110e7565b945060608801356112d4816110e7565b93506080880135925060a088013567ffffffffffffffff8111156112f757600080fd5b6113038a828b01611125565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561133157600080fd5b873561133c816110e7565b9650602088013561134c816110e7565b9550604088013561135c816110e7565b9450606088013593506113716080890161110c565b925060a088013567ffffffffffffffff8111156112f757600080fd5b6000845161139f8184602089016111f1565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516113db816001850160208a016111f1565b600192019182015283516113f68160028401602088016111f1565b0160020195945050505050565b60006020828403121561141557600080fd5b8151610fac816110e7565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061149f606083018486611420565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526114f960c083018486611420565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006115356060830185611221565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115ab576115ab61154b565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826115f0576115f06115b2565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156116365761163661154b565b500390565b60008261164a5761164a6115b2565b500690565b600082198211156116625761166261154b565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
}