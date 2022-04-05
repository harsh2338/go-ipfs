// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commands

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PlagiarismContractCodeFile is an auto generated low-level Go binding around an user-defined struct.
type PlagiarismContractCodeFile struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSHash    string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSet         [][16]byte
	HashSetLength   *big.Int
}

// PlagiarismContractMetaData contains all meta data concerning the PlagiarismContract contract.
var PlagiarismContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileIPFSHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"}],\"name\":\"CodeFileUploadEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"plagiarisedResult\",\"type\":\"bool\"}],\"name\":\"PlagiarismResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"l1\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes16[]\",\"name\":\"l2\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"l1OriLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2OriLength\",\"type\":\"uint256\"}],\"name\":\"calculateSimilarityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"_hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"}],\"name\":\"checkIfPlagiarised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filesMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"list1\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes16[]\",\"name\":\"list2\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"setLength1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setLength2\",\"type\":\"uint256\"}],\"name\":\"getCommonElementsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fileIndex\",\"type\":\"uint256\"}],\"name\":\"getFileByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"}],\"internalType\":\"structPlagiarismContract.CodeFile\",\"name\":\"records\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"_hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"oriHashSetLength\",\"type\":\"uint256\"}],\"name\":\"getMaximumSimilarityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_fileIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"_hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d665ad1": "calculateSimilarityScore(bytes16[],bytes16[],uint256,uint256)",
		"31631b64": "checkIfPlagiarised(bytes16[],uint256)",
		"43953519": "fileCount()",
		"c93bcd91": "filesMap(uint256)",
		"9ecefadc": "getCommonElementsCount(bytes16[],bytes16[],uint256,uint256)",
		"097694e1": "getFileByIndex(uint256)",
		"dbb373ee": "getMaximumSimilarityScore(bytes16[],uint256)",
		"371303c0": "inc()",
		"58ae2ae7": "uploadFile(uint256,string,string,string,string,bytes16[],uint256)",
	},
	Bin: "0x608060405260008055600060015534801561001957600080fd5b506113f3806100296000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80634395351911610066578063439535191461010f57806358ae2ae7146101185780639ecefadc1461012b578063c93bcd911461013e578063dbb373ee1461016657600080fd5b8063097694e1146100985780630d665ad1146100c157806331631b64146100e2578063371303c014610105575b600080fd5b6100ab6100a6366004610cdb565b610179565b6040516100b89190610d86565b60405180910390f35b6100d46100cf366004610f3f565b61050a565b6040519081526020016100b8565b6100f56100f0366004610fb3565b610557565b60405190151581526020016100b8565b61010d610588565b005b6100d460005481565b61010d610126366004611068565b61059f565b6100d4610139366004610f3f565b610781565b61015161014c366004610cdb565b610827565b6040516100b89998979695949392919061114d565b6100d4610174366004610fb3565b610a9e565b6101d8604051806101400160405280600081526020016000815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001600081526020016060815260200160608152602001600081525090565b6002600083815260200190815260200160002060405180610140016040529081600082015481526020016001820154815260200160028201805461021b906111d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610247906111d7565b80156102945780601f1061026957610100808354040283529160200191610294565b820191906000526020600020905b81548152906001019060200180831161027757829003601f168201915b505050505081526020016003820180546102ad906111d7565b80601f01602080910402602001604051908101604052809291908181526020018280546102d9906111d7565b80156103265780601f106102fb57610100808354040283529160200191610326565b820191906000526020600020905b81548152906001019060200180831161030957829003601f168201915b5050505050815260200160048201805461033f906111d7565b80601f016020809104026020016040519081016040528092919081815260200182805461036b906111d7565b80156103b85780601f1061038d576101008083540402835291602001916103b8565b820191906000526020600020905b81548152906001019060200180831161039b57829003601f168201915b505050918352505060058201546001600160a01b03166020820152600682015460408201526007820180546060909201916103f2906111d7565b80601f016020809104026020016040519081016040528092919081815260200182805461041e906111d7565b801561046b5780601f106104405761010080835404028352916020019161046b565b820191906000526020600020905b81548152906001019060200180831161044e57829003601f168201915b50505050508152602001600882018054806020026020016040519081016040528092919081815260200182805480156104f057602002820191906000526020600020906000905b82829054906101000a900460801b6001600160801b03191681526020019060100190602082600f010492830192600103820291508084116104b25790505b505050505081526020016009820154815250509050919050565b60008061051b868688518851610781565b905060008161052a8587611227565b610534919061123f565b905080610542836064611256565b61054c9190611275565b979650505050505050565b6000806105648484610a9e565b9050601e8082111561057b57600192505050610582565b6000925050505b92915050565b6001805490600061059883611297565b9190505550565b6105a98282610557565b156105e757604051600181527f711784307b11c5f98f3bc25170dfd6d229d797a0891cb2dd54b1fd406e60fc6d9060200160405180910390a1610778565b6000805490806105f683611297565b909155505060408051610140810182526000805480835260208084018c81528486018c8152606086018c9052608086018b90523360a08701524260c087015260e086018a9052610100860189905261012086018890529284526002808352959093208451815592516001840155905180519394929361067c938501929190910190610b99565b5060608201518051610698916003840191602090910190610b99565b50608082015180516106b4916004840191602090910190610b99565b5060a08201516005820180546001600160a01b0319166001600160a01b0390921691909117905560c0820151600682015560e08201518051610700916007840191602090910190610b99565b50610100820151805161071d916008840191602090910190610c1d565b5061012082015181600901559050507f2cb3b83cdd1534b844c5d0a1492ed9e146aba3d1a63aee4e5bd0d3f8e85212456000548888888833428a8a8a60405161076f9a999897969594939291906112b0565b60405180910390a15b50505050505050565b600080805b8481101561081d5760005b8481101561080a578681815181106107ab576107ab611381565b60200260200101516001600160801b0319168883815181106107cf576107cf611381565b60200260200101516001600160801b031916036107f857826107f081611297565b93505061080a565b8061080281611297565b915050610791565b508061081581611297565b915050610786565b5095945050505050565b600260208190526000918252604090912080546001820154928201805491939291610851906111d7565b80601f016020809104026020016040519081016040528092919081815260200182805461087d906111d7565b80156108ca5780601f1061089f576101008083540402835291602001916108ca565b820191906000526020600020905b8154815290600101906020018083116108ad57829003601f168201915b5050505050908060030180546108df906111d7565b80601f016020809104026020016040519081016040528092919081815260200182805461090b906111d7565b80156109585780601f1061092d57610100808354040283529160200191610958565b820191906000526020600020905b81548152906001019060200180831161093b57829003601f168201915b50505050509080600401805461096d906111d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610999906111d7565b80156109e65780601f106109bb576101008083540402835291602001916109e6565b820191906000526020600020905b8154815290600101906020018083116109c957829003601f168201915b505050506005830154600684015460078501805494956001600160a01b039093169491935090610a15906111d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610a41906111d7565b8015610a8e5780601f10610a6357610100808354040283529160200191610a8e565b820191906000526020600020905b815481529060010190602001808311610a7157829003601f168201915b5050505050908060090154905089565b6000808060015b6000548111610afd576000610ab982610b07565b600083815260026020526040902060090154909150610ada8883898461050a565b935084841115610ae8578394505b50508080610af590611297565b915050610aa5565b5090949350505050565b600081815260026020908152604091829020600801805483518184028101840190945280845260609392830182828015610b8d57602002820191906000526020600020906000905b82829054906101000a900460801b6001600160801b03191681526020019060100190602082600f01049283019260010382029150808411610b4f5790505b50505050509050919050565b828054610ba5906111d7565b90600052602060002090601f016020900481019282610bc75760008555610c0d565b82601f10610be057805160ff1916838001178555610c0d565b82800160010185558215610c0d579182015b82811115610c0d578251825591602001919060010190610bf2565b50610c19929150610cc6565b5090565b82805482825590600052602060002090600101600290048101928215610c0d5791602002820160005b83821115610c8a57835183826101000a8154816001600160801b03021916908360801c02179055509260200192601001602081600f01049283019260010302610c46565b8015610cbd5782816101000a8154906001600160801b030219169055601001602081600f01049283019260010302610c8a565b5050610c199291505b5b80821115610c195760008155600101610cc7565b600060208284031215610ced57600080fd5b5035919050565b6000815180845260005b81811015610d1a57602081850181015186830182015201610cfe565b81811115610d2c576000602083870101525b50601f01601f19169290920160200192915050565b600081518084526020808501945080840160005b83811015610d7b5781516001600160801b03191687529582019590820190600101610d55565b509495945050505050565b60208152815160208201526020820151604082015260006040830151610140806060850152610db9610160850183610cf4565b91506060850151601f1980868503016080870152610dd78483610cf4565b935060808701519150808685030160a0870152610df48483610cf4565b935060a08701519150610e1260c08701836001600160a01b03169052565b60c087015160e087015260e08701519150610100818786030181880152610e398584610cf4565b945080880151925050610120818786030181880152610e588584610d41565b970151959092019490945250929392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610eab57610eab610e6c565b604052919050565b600082601f830112610ec457600080fd5b8135602067ffffffffffffffff821115610ee057610ee0610e6c565b8160051b610eef828201610e82565b9283528481018201928281019087851115610f0957600080fd5b83870192505b8483101561054c5782356001600160801b031981168114610f305760008081fd5b82529183019190830190610f0f565b60008060008060808587031215610f5557600080fd5b843567ffffffffffffffff80821115610f6d57600080fd5b610f7988838901610eb3565b95506020870135915080821115610f8f57600080fd5b50610f9c87828801610eb3565b949794965050505060408301359260600135919050565b60008060408385031215610fc657600080fd5b823567ffffffffffffffff811115610fdd57600080fd5b610fe985828601610eb3565b95602094909401359450505050565b600082601f83011261100957600080fd5b813567ffffffffffffffff81111561102357611023610e6c565b611036601f8201601f1916602001610e82565b81815284602083860101111561104b57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060e0888a03121561108357600080fd5b87359650602088013567ffffffffffffffff808211156110a257600080fd5b6110ae8b838c01610ff8565b975060408a01359150808211156110c457600080fd5b6110d08b838c01610ff8565b965060608a01359150808211156110e657600080fd5b6110f28b838c01610ff8565b955060808a013591508082111561110857600080fd5b6111148b838c01610ff8565b945060a08a013591508082111561112a57600080fd5b506111378a828b01610eb3565b92505060c0880135905092959891949750929550565b60006101208b83528a602084015280604084015261116d8184018b610cf4565b90508281036060840152611181818a610cf4565b905082810360808401526111958189610cf4565b6001600160a01b03881660a085015260c0840187905283810360e085015290506111bf8186610cf4565b915050826101008301529a9950505050505050505050565b600181811c908216806111eb57607f821691505b60208210810361120b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561123a5761123a611211565b500190565b60008282101561125157611251611211565b500390565b600081600019048311821515161561127057611270611211565b500290565b60008261129257634e487b7160e01b600052601260045260246000fd5b500490565b6000600182016112a9576112a9611211565b5060010190565b60006101408c835260208c818501528160408501526112d18285018d610cf4565b915083820360608501526112e5828c610cf4565b915083820360808501526112f9828b610cf4565b6001600160a01b038a1660a086015260c0850189905284810360e086015291506113238288610cf4565b848103610100860152865180825282880193509082019060005b818110156113635784516001600160801b0319168352938301939183019160010161133d565b5050809350505050826101208301529b9a5050505050505050505050565b634e487b7160e01b600052603260045260246000fdfea26469706673582212208638058ac6c4981baa5865ff6ecdb32d93e8854f51d483b1ddff2fb31192c0d764736f6c637828302e382e31332d646576656c6f702e323032322e332e31362b636f6d6d69742e37323461663733660059",
}

// PlagiarismContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PlagiarismContractMetaData.ABI instead.
var PlagiarismContractABI = PlagiarismContractMetaData.ABI

// Deprecated: Use PlagiarismContractMetaData.Sigs instead.
// PlagiarismContractFuncSigs maps the 4-byte function signature to its string representation.
var PlagiarismContractFuncSigs = PlagiarismContractMetaData.Sigs

// PlagiarismContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PlagiarismContractMetaData.Bin instead.
var PlagiarismContractBin = PlagiarismContractMetaData.Bin

// DeployPlagiarismContract deploys a new Ethereum contract, binding an instance of PlagiarismContract to it.
func DeployPlagiarismContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlagiarismContract, error) {
	parsed, err := PlagiarismContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PlagiarismContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlagiarismContract{PlagiarismContractCaller: PlagiarismContractCaller{contract: contract}, PlagiarismContractTransactor: PlagiarismContractTransactor{contract: contract}, PlagiarismContractFilterer: PlagiarismContractFilterer{contract: contract}}, nil
}

// PlagiarismContract is an auto generated Go binding around an Ethereum contract.
type PlagiarismContract struct {
	PlagiarismContractCaller     // Read-only binding to the contract
	PlagiarismContractTransactor // Write-only binding to the contract
	PlagiarismContractFilterer   // Log filterer for contract events
}

// PlagiarismContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlagiarismContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlagiarismContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlagiarismContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlagiarismContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlagiarismContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlagiarismContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlagiarismContractSession struct {
	Contract     *PlagiarismContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlagiarismContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlagiarismContractCallerSession struct {
	Contract *PlagiarismContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PlagiarismContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlagiarismContractTransactorSession struct {
	Contract     *PlagiarismContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PlagiarismContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlagiarismContractRaw struct {
	Contract *PlagiarismContract // Generic contract binding to access the raw methods on
}

// PlagiarismContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlagiarismContractCallerRaw struct {
	Contract *PlagiarismContractCaller // Generic read-only contract binding to access the raw methods on
}

// PlagiarismContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlagiarismContractTransactorRaw struct {
	Contract *PlagiarismContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlagiarismContract creates a new instance of PlagiarismContract, bound to a specific deployed contract.
func NewPlagiarismContract(address common.Address, backend bind.ContractBackend) (*PlagiarismContract, error) {
	contract, err := bindPlagiarismContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlagiarismContract{PlagiarismContractCaller: PlagiarismContractCaller{contract: contract}, PlagiarismContractTransactor: PlagiarismContractTransactor{contract: contract}, PlagiarismContractFilterer: PlagiarismContractFilterer{contract: contract}}, nil
}

// NewPlagiarismContractCaller creates a new read-only instance of PlagiarismContract, bound to a specific deployed contract.
func NewPlagiarismContractCaller(address common.Address, caller bind.ContractCaller) (*PlagiarismContractCaller, error) {
	contract, err := bindPlagiarismContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractCaller{contract: contract}, nil
}

// NewPlagiarismContractTransactor creates a new write-only instance of PlagiarismContract, bound to a specific deployed contract.
func NewPlagiarismContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PlagiarismContractTransactor, error) {
	contract, err := bindPlagiarismContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractTransactor{contract: contract}, nil
}

// NewPlagiarismContractFilterer creates a new log filterer instance of PlagiarismContract, bound to a specific deployed contract.
func NewPlagiarismContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PlagiarismContractFilterer, error) {
	contract, err := bindPlagiarismContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractFilterer{contract: contract}, nil
}

// bindPlagiarismContract binds a generic wrapper to an already deployed contract.
func bindPlagiarismContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlagiarismContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlagiarismContract *PlagiarismContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlagiarismContract.Contract.PlagiarismContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlagiarismContract *PlagiarismContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.PlagiarismContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlagiarismContract *PlagiarismContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.PlagiarismContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlagiarismContract *PlagiarismContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlagiarismContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlagiarismContract *PlagiarismContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlagiarismContract *PlagiarismContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.contract.Transact(opts, method, params...)
}

// CalculateSimilarityScore is a free data retrieval call binding the contract method 0x0d665ad1.
//
// Solidity: function calculateSimilarityScore(bytes16[] l1, bytes16[] l2, uint256 l1OriLength, uint256 l2OriLength) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) CalculateSimilarityScore(opts *bind.CallOpts, l1 [][16]byte, l2 [][16]byte, l1OriLength *big.Int, l2OriLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "calculateSimilarityScore", l1, l2, l1OriLength, l2OriLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSimilarityScore is a free data retrieval call binding the contract method 0x0d665ad1.
//
// Solidity: function calculateSimilarityScore(bytes16[] l1, bytes16[] l2, uint256 l1OriLength, uint256 l2OriLength) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) CalculateSimilarityScore(l1 [][16]byte, l2 [][16]byte, l1OriLength *big.Int, l2OriLength *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.CalculateSimilarityScore(&_PlagiarismContract.CallOpts, l1, l2, l1OriLength, l2OriLength)
}

// CalculateSimilarityScore is a free data retrieval call binding the contract method 0x0d665ad1.
//
// Solidity: function calculateSimilarityScore(bytes16[] l1, bytes16[] l2, uint256 l1OriLength, uint256 l2OriLength) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) CalculateSimilarityScore(l1 [][16]byte, l2 [][16]byte, l1OriLength *big.Int, l2OriLength *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.CalculateSimilarityScore(&_PlagiarismContract.CallOpts, l1, l2, l1OriLength, l2OriLength)
}

// CheckIfPlagiarised is a free data retrieval call binding the contract method 0x31631b64.
//
// Solidity: function checkIfPlagiarised(bytes16[] _hashSet, uint256 hashSetLength) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCaller) CheckIfPlagiarised(opts *bind.CallOpts, _hashSet [][16]byte, hashSetLength *big.Int) (bool, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "checkIfPlagiarised", _hashSet, hashSetLength)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfPlagiarised is a free data retrieval call binding the contract method 0x31631b64.
//
// Solidity: function checkIfPlagiarised(bytes16[] _hashSet, uint256 hashSetLength) view returns(bool)
func (_PlagiarismContract *PlagiarismContractSession) CheckIfPlagiarised(_hashSet [][16]byte, hashSetLength *big.Int) (bool, error) {
	return _PlagiarismContract.Contract.CheckIfPlagiarised(&_PlagiarismContract.CallOpts, _hashSet, hashSetLength)
}

// CheckIfPlagiarised is a free data retrieval call binding the contract method 0x31631b64.
//
// Solidity: function checkIfPlagiarised(bytes16[] _hashSet, uint256 hashSetLength) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCallerSession) CheckIfPlagiarised(_hashSet [][16]byte, hashSetLength *big.Int) (bool, error) {
	return _PlagiarismContract.Contract.CheckIfPlagiarised(&_PlagiarismContract.CallOpts, _hashSet, hashSetLength)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) FileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "fileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) FileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.FileCount(&_PlagiarismContract.CallOpts)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) FileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.FileCount(&_PlagiarismContract.CallOpts)
}

// FilesMap is a free data retrieval call binding the contract method 0xc93bcd91.
//
// Solidity: function filesMap(uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractCaller) FilesMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSHash    string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
}, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "filesMap", arg0)

	outstruct := new(struct {
		FileId          *big.Int
		FileSize        *big.Int
		FileIPFSHash    string
		FileName        string
		FileDescription string
		CodeAuthor      common.Address
		TimeUploaded    *big.Int
		CodeFingerPrint string
		HashSetLength   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FileSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FileIPFSHash = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.FileName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.FileDescription = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CodeAuthor = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TimeUploaded = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CodeFingerPrint = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.HashSetLength = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FilesMap is a free data retrieval call binding the contract method 0xc93bcd91.
//
// Solidity: function filesMap(uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractSession) FilesMap(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSHash    string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
}, error) {
	return _PlagiarismContract.Contract.FilesMap(&_PlagiarismContract.CallOpts, arg0)
}

// FilesMap is a free data retrieval call binding the contract method 0xc93bcd91.
//
// Solidity: function filesMap(uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractCallerSession) FilesMap(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSHash    string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
}, error) {
	return _PlagiarismContract.Contract.FilesMap(&_PlagiarismContract.CallOpts, arg0)
}

// GetCommonElementsCount is a free data retrieval call binding the contract method 0x9ecefadc.
//
// Solidity: function getCommonElementsCount(bytes16[] list1, bytes16[] list2, uint256 setLength1, uint256 setLength2) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) GetCommonElementsCount(opts *bind.CallOpts, list1 [][16]byte, list2 [][16]byte, setLength1 *big.Int, setLength2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getCommonElementsCount", list1, list2, setLength1, setLength2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommonElementsCount is a free data retrieval call binding the contract method 0x9ecefadc.
//
// Solidity: function getCommonElementsCount(bytes16[] list1, bytes16[] list2, uint256 setLength1, uint256 setLength2) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) GetCommonElementsCount(list1 [][16]byte, list2 [][16]byte, setLength1 *big.Int, setLength2 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.GetCommonElementsCount(&_PlagiarismContract.CallOpts, list1, list2, setLength1, setLength2)
}

// GetCommonElementsCount is a free data retrieval call binding the contract method 0x9ecefadc.
//
// Solidity: function getCommonElementsCount(bytes16[] list1, bytes16[] list2, uint256 setLength1, uint256 setLength2) pure returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) GetCommonElementsCount(list1 [][16]byte, list2 [][16]byte, setLength1 *big.Int, setLength2 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.GetCommonElementsCount(&_PlagiarismContract.CallOpts, list1, list2, setLength1, setLength2)
}

// GetFileByIndex is a free data retrieval call binding the contract method 0x097694e1.
//
// Solidity: function getFileByIndex(uint256 _fileIndex) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256) records)
func (_PlagiarismContract *PlagiarismContractCaller) GetFileByIndex(opts *bind.CallOpts, _fileIndex *big.Int) (PlagiarismContractCodeFile, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getFileByIndex", _fileIndex)

	if err != nil {
		return *new(PlagiarismContractCodeFile), err
	}

	out0 := *abi.ConvertType(out[0], new(PlagiarismContractCodeFile)).(*PlagiarismContractCodeFile)

	return out0, err

}

// GetFileByIndex is a free data retrieval call binding the contract method 0x097694e1.
//
// Solidity: function getFileByIndex(uint256 _fileIndex) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256) records)
func (_PlagiarismContract *PlagiarismContractSession) GetFileByIndex(_fileIndex *big.Int) (PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFileByIndex(&_PlagiarismContract.CallOpts, _fileIndex)
}

// GetFileByIndex is a free data retrieval call binding the contract method 0x097694e1.
//
// Solidity: function getFileByIndex(uint256 _fileIndex) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256) records)
func (_PlagiarismContract *PlagiarismContractCallerSession) GetFileByIndex(_fileIndex *big.Int) (PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFileByIndex(&_PlagiarismContract.CallOpts, _fileIndex)
}

// GetMaximumSimilarityScore is a free data retrieval call binding the contract method 0xdbb373ee.
//
// Solidity: function getMaximumSimilarityScore(bytes16[] _hashSet, uint256 oriHashSetLength) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) GetMaximumSimilarityScore(opts *bind.CallOpts, _hashSet [][16]byte, oriHashSetLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getMaximumSimilarityScore", _hashSet, oriHashSetLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumSimilarityScore is a free data retrieval call binding the contract method 0xdbb373ee.
//
// Solidity: function getMaximumSimilarityScore(bytes16[] _hashSet, uint256 oriHashSetLength) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) GetMaximumSimilarityScore(_hashSet [][16]byte, oriHashSetLength *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.GetMaximumSimilarityScore(&_PlagiarismContract.CallOpts, _hashSet, oriHashSetLength)
}

// GetMaximumSimilarityScore is a free data retrieval call binding the contract method 0xdbb373ee.
//
// Solidity: function getMaximumSimilarityScore(bytes16[] _hashSet, uint256 oriHashSetLength) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) GetMaximumSimilarityScore(_hashSet [][16]byte, oriHashSetLength *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.GetMaximumSimilarityScore(&_PlagiarismContract.CallOpts, _hashSet, oriHashSetLength)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_PlagiarismContract *PlagiarismContractTransactor) Inc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "inc")
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_PlagiarismContract *PlagiarismContractSession) Inc() (*types.Transaction, error) {
	return _PlagiarismContract.Contract.Inc(&_PlagiarismContract.TransactOpts)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_PlagiarismContract *PlagiarismContractTransactorSession) Inc() (*types.Transaction, error) {
	return _PlagiarismContract.Contract.Inc(&_PlagiarismContract.TransactOpts)
}

// UploadFile is a paid mutator transaction binding the contract method 0x58ae2ae7.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSHash, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength) returns()
func (_PlagiarismContract *PlagiarismContractTransactor) UploadFile(opts *bind.TransactOpts, _fileSize *big.Int, _fileIPFSHash string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "uploadFile", _fileSize, _fileIPFSHash, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength)
}

// UploadFile is a paid mutator transaction binding the contract method 0x58ae2ae7.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSHash, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength) returns()
func (_PlagiarismContract *PlagiarismContractSession) UploadFile(_fileSize *big.Int, _fileIPFSHash string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.UploadFile(&_PlagiarismContract.TransactOpts, _fileSize, _fileIPFSHash, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength)
}

// UploadFile is a paid mutator transaction binding the contract method 0x58ae2ae7.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSHash, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength) returns()
func (_PlagiarismContract *PlagiarismContractTransactorSession) UploadFile(_fileSize *big.Int, _fileIPFSHash string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.UploadFile(&_PlagiarismContract.TransactOpts, _fileSize, _fileIPFSHash, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength)
}

// PlagiarismContractCodeFileUploadEventIterator is returned from FilterCodeFileUploadEvent and is used to iterate over the raw logs and unpacked data for CodeFileUploadEvent events raised by the PlagiarismContract contract.
type PlagiarismContractCodeFileUploadEventIterator struct {
	Event *PlagiarismContractCodeFileUploadEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlagiarismContractCodeFileUploadEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlagiarismContractCodeFileUploadEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlagiarismContractCodeFileUploadEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlagiarismContractCodeFileUploadEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlagiarismContractCodeFileUploadEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlagiarismContractCodeFileUploadEvent represents a CodeFileUploadEvent event raised by the PlagiarismContract contract.
type PlagiarismContractCodeFileUploadEvent struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSHash    string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSet         [][16]byte
	HashSetLength   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCodeFileUploadEvent is a free log retrieval operation binding the contract event 0x2cb3b83cdd1534b844c5d0a1492ed9e146aba3d1a63aee4e5bd0d3f8e8521245.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterCodeFileUploadEvent(opts *bind.FilterOpts) (*PlagiarismContractCodeFileUploadEventIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "CodeFileUploadEvent")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractCodeFileUploadEventIterator{contract: _PlagiarismContract.contract, event: "CodeFileUploadEvent", logs: logs, sub: sub}, nil
}

// WatchCodeFileUploadEvent is a free log subscription operation binding the contract event 0x2cb3b83cdd1534b844c5d0a1492ed9e146aba3d1a63aee4e5bd0d3f8e8521245.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractFilterer) WatchCodeFileUploadEvent(opts *bind.WatchOpts, sink chan<- *PlagiarismContractCodeFileUploadEvent) (event.Subscription, error) {

	logs, sub, err := _PlagiarismContract.contract.WatchLogs(opts, "CodeFileUploadEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlagiarismContractCodeFileUploadEvent)
				if err := _PlagiarismContract.contract.UnpackLog(event, "CodeFileUploadEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCodeFileUploadEvent is a log parse operation binding the contract event 0x2cb3b83cdd1534b844c5d0a1492ed9e146aba3d1a63aee4e5bd0d3f8e8521245.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSHash, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength)
func (_PlagiarismContract *PlagiarismContractFilterer) ParseCodeFileUploadEvent(log types.Log) (*PlagiarismContractCodeFileUploadEvent, error) {
	event := new(PlagiarismContractCodeFileUploadEvent)
	if err := _PlagiarismContract.contract.UnpackLog(event, "CodeFileUploadEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlagiarismContractPlagiarismResultIterator is returned from FilterPlagiarismResult and is used to iterate over the raw logs and unpacked data for PlagiarismResult events raised by the PlagiarismContract contract.
type PlagiarismContractPlagiarismResultIterator struct {
	Event *PlagiarismContractPlagiarismResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlagiarismContractPlagiarismResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlagiarismContractPlagiarismResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlagiarismContractPlagiarismResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlagiarismContractPlagiarismResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlagiarismContractPlagiarismResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlagiarismContractPlagiarismResult represents a PlagiarismResult event raised by the PlagiarismContract contract.
type PlagiarismContractPlagiarismResult struct {
	PlagiarisedResult bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPlagiarismResult is a free log retrieval operation binding the contract event 0x711784307b11c5f98f3bc25170dfd6d229d797a0891cb2dd54b1fd406e60fc6d.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterPlagiarismResult(opts *bind.FilterOpts) (*PlagiarismContractPlagiarismResultIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "PlagiarismResult")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractPlagiarismResultIterator{contract: _PlagiarismContract.contract, event: "PlagiarismResult", logs: logs, sub: sub}, nil
}

// WatchPlagiarismResult is a free log subscription operation binding the contract event 0x711784307b11c5f98f3bc25170dfd6d229d797a0891cb2dd54b1fd406e60fc6d.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult)
func (_PlagiarismContract *PlagiarismContractFilterer) WatchPlagiarismResult(opts *bind.WatchOpts, sink chan<- *PlagiarismContractPlagiarismResult) (event.Subscription, error) {

	logs, sub, err := _PlagiarismContract.contract.WatchLogs(opts, "PlagiarismResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlagiarismContractPlagiarismResult)
				if err := _PlagiarismContract.contract.UnpackLog(event, "PlagiarismResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlagiarismResult is a log parse operation binding the contract event 0x711784307b11c5f98f3bc25170dfd6d229d797a0891cb2dd54b1fd406e60fc6d.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult)
func (_PlagiarismContract *PlagiarismContractFilterer) ParsePlagiarismResult(log types.Log) (*PlagiarismContractPlagiarismResult, error) {
	event := new(PlagiarismContractPlagiarismResult)
	if err := _PlagiarismContract.contract.UnpackLog(event, "PlagiarismResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
