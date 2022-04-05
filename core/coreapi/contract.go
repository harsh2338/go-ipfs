// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreapi

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
	FileIPFSCID     string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSet         [][16]byte
	HashSetLength   *big.Int
	Language        *big.Int
}

// PlagiarismContractMetaData contains all meta data concerning the PlagiarismContract contract.
var PlagiarismContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"CheckingPlagiarism\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileIPFSCID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"CodeFileUploadEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"CodeSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPlagiarismContract.CodeFile[]\",\"name\":\"codefiles\",\"type\":\"tuple[]\"}],\"name\":\"GetFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"plagiarisedResult\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"}],\"name\":\"PlagiarismResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"addPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"l1\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes16[]\",\"name\":\"l2\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"l1OriLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2OriLength\",\"type\":\"uint256\"}],\"name\":\"calculateSimilarityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"_hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"checkIfPlagiarised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"doesUserHavePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fileCountByLanguage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filesMapByLanguage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"list1\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes16[]\",\"name\":\"list2\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"setLength1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setLength2\",\"type\":\"uint256\"}],\"name\":\"getCommonElementsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fileIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"getFileByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"internalType\":\"structPlagiarismContract.CodeFile\",\"name\":\"records\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFilesUploadedByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileIPFSCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"codeAuthor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeUploaded\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"internalType\":\"structPlagiarismContract.CodeFile[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasPermissionsMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"removePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_fileIPFSCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_codeFingerPrint\",\"type\":\"string\"},{\"internalType\":\"bytes16[]\",\"name\":\"_hashSet\",\"type\":\"bytes16[]\"},{\"internalType\":\"uint256\",\"name\":\"hashSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"language\",\"type\":\"uint256\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFilesIndexByLanguage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5836e03d": "addPermission(string)",
		"0d665ad1": "calculateSimilarityScore(bytes16[],bytes16[],uint256,uint256)",
		"d11e44d6": "checkIfPlagiarised(string,bytes16[],uint256,uint256)",
		"a89796cc": "doesUserHavePermission(string,address)",
		"0641e449": "fileCountByLanguage(uint256)",
		"60ff2aac": "filesMapByLanguage(uint256,uint256)",
		"9ecefadc": "getCommonElementsCount(bytes16[],bytes16[],uint256,uint256)",
		"ee7e64f7": "getFileByIndex(uint256,uint256)",
		"bab50cc9": "getFileCount()",
		"51887b71": "getFilesUploadedByUser()",
		"9e199566": "hasPermissionsMap(string,address)",
		"4217649f": "removePermission(string)",
		"31eb7a1a": "threshold(uint256)",
		"7f9d4b4a": "totalFileCount()",
		"ca749bec": "uploadFile(uint256,string,string,string,string,bytes16[],uint256,uint256)",
		"97f42b1a": "userFilesIndexByLanguage(address,uint256,uint256)",
	},
	Bin: "0x60e06040526000608081815260a082905260c082905262000023919060036200006b565b50600060015560408051606081018252601e8082526020820181905291810191909152620000569060029060036200006b565b503480156200006457600080fd5b50620000d7565b828054828255906000526020600020908101928215620000ae579160200282015b82811115620000ae578251829060ff169055916020019190600101906200008c565b50620000bc929150620000c0565b5090565b5b80821115620000bc5760008155600101620000c1565b611c4180620000e76000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806397f42b1a11610097578063bab50cc911610066578063bab50cc914610247578063ca749bec1461024f578063d11e44d614610262578063ee7e64f71461027557600080fd5b806397f42b1a146101c05780639e199566146101d35780639ecefadc14610221578063a89796cc1461023457600080fd5b806351887b71116100d357806351887b71146101665780635836e03d1461017b57806360ff2aac1461018e5780637f9d4b4a146101b757600080fd5b80630641e449146101055780630d665ad11461012b57806331eb7a1a1461013e5780634217649f14610151575b600080fd5b610118610113366004611386565b610295565b6040519081526020015b60405180910390f35b61011861013936600461147d565b6102b6565b61011861014c366004611386565b610305565b61016461015f366004611561565b610315565b005b61016e610354565b6040516101229190611712565b610164610189366004611561565b6105a3565b6101a161019c366004611774565b6105b5565b6040516101229a99989796959493929190611796565b61011860015481565b6101186101ce366004611840565b61083b565b6102116101e1366004611873565b81516020818401810180516006825292820194820194909420919093529091526000908152604090205460ff1681565b6040519015158152602001610122565b61011861022f36600461147d565b610879565b610211610242366004611873565b61091f565b600154610118565b61016461025d3660046118c1565b610963565b6102116102703660046119b0565b610cf9565b610288610283366004611774565b610df0565b60405161012291906119ea565b600081815481106102a557600080fd5b600091825260209091200154905081565b6000806102c7868688518851610879565b90506000816102d68587611a1a565b6102e09190611a32565b9050806102ee836064611a49565b6102f89190611a68565b925050505b949350505050565b600281815481106102a557600080fd5b60006006826040516103279190611a8a565b908152604080516020928190038301902033600090815292529020805460ff191691151591909117905550565b3360009081526004602090815260408083208380529091528082205460018352818320546002845291832054606093919291816103918585611a1a565b61039b9190611a1a565b67ffffffffffffffff8111156103b3576103b361139f565b6040519080825280602002602001820160405280156103ec57816020015b6103d96111e1565b8152602001906001900390816103d15790505b5090506000805b8581101561047c573360009081526004602090815260408083208380529091529020805461043e91908390811061042c5761042c611aa6565b90600052602060002001546000610df0565b83838151811061045057610450611aa6565b6020026020010181905250818061046690611abc565b925050808061047490611abc565b9150506103f3565b5060005b8481101561050a5733600090815260046020908152604080832060018452909152902080546104cc9190839081106104ba576104ba611aa6565b90600052602060002001546001610df0565b8383815181106104de576104de611aa6565b602002602001018190525081806104f490611abc565b925050808061050290611abc565b915050610480565b5060005b8381101561059857336000908152600460209081526040808320600284529091529020805461055a91908390811061054857610548611aa6565b90600052602060002001546002610df0565b83838151811061056c5761056c611aa6565b6020026020010181905250818061058290611abc565b925050808061059090611abc565b91505061050e565b509095945050505050565b60016006826040516103279190611a8a565b6003602090815260009283526040808420909152908252902080546001820154600283018054929391926105e890611ad5565b80601f016020809104026020016040519081016040528092919081815260200182805461061490611ad5565b80156106615780601f1061063657610100808354040283529160200191610661565b820191906000526020600020905b81548152906001019060200180831161064457829003601f168201915b50505050509080600301805461067690611ad5565b80601f01602080910402602001604051908101604052809291908181526020018280546106a290611ad5565b80156106ef5780601f106106c4576101008083540402835291602001916106ef565b820191906000526020600020905b8154815290600101906020018083116106d257829003601f168201915b50505050509080600401805461070490611ad5565b80601f016020809104026020016040519081016040528092919081815260200182805461073090611ad5565b801561077d5780601f106107525761010080835404028352916020019161077d565b820191906000526020600020905b81548152906001019060200180831161076057829003601f168201915b505050506005830154600684015460078501805494956001600160a01b0390931694919350906107ac90611ad5565b80601f01602080910402602001604051908101604052809291908181526020018280546107d890611ad5565b80156108255780601f106107fa57610100808354040283529160200191610825565b820191906000526020600020905b81548152906001019060200180831161080857829003601f168201915b50505050509080600901549080600a015490508a565b6004602052826000526040600020602052816000526040600020818154811061086357600080fd5b9060005260206000200160009250925050505481565b600080805b848110156109155760005b84811015610902578681815181106108a3576108a3611aa6565b60200260200101516001600160801b0319168883815181106108c7576108c7611aa6565b60200260200101516001600160801b031916036108f057826108e881611abc565b935050610902565b806108fa81611abc565b915050610889565b508061090d81611abc565b91505061087e565b5095945050505050565b60006006836040516109319190611a8a565b908152604080519182900360209081019092206001600160a01b0385166000908152925290205460ff16905092915050565b7f90a4de5c450d2bbdd8a9564fb9646ea0df44cf5dbb8e4bed1c6f806b7b37e57e42878360405161099693929190611b0f565b60405180910390a16005876040516109ae9190611a8a565b9081526040519081900360200190205460ff161515600103610a095760408051600181524260208201527f33baf730752a6c778112e3ee4c2e8bd128e07d02117675aa7c7feec6dd8b51db91015b60405180910390a1610cef565b610a1586848484610cf9565b15610a505760408051600181524260208201527f33baf730752a6c778112e3ee4c2e8bd128e07d02117675aa7c7feec6dd8b51db91016109fc565b610a59876105a3565b60008181548110610a6c57610a6c611aa6565b60009182526020822001805491610a8283611abc565b909155505060018054906000610a9783611abc565b91905055506001600588604051610aae9190611a8a565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506040518061016001604052806001548152602001898152602001888152602001878152602001868152602001336001600160a01b0316815260200142815260200185815260200184815260200183815260200182815250600360008381526020019081526020016000206000808481548110610b5157610b51611aa6565b9060005260206000200154815260200190815260200160002060008201518160000155602082015181600101556040820151816002019080519060200190610b9a929190611244565b5060608201518051610bb6916003840191602090910190611244565b5060808201518051610bd2916004840191602090910190611244565b5060a08201516005820180546001600160a01b0319166001600160a01b0390921691909117905560c0820151600682015560e08201518051610c1e916007840191602090910190611244565b506101008201518051610c3b9160088401916020909101906112c8565b50610120820151600982015561014090910151600a909101553360009081526004602090815260408083208484529091528120815490919083908110610c8357610c83611aa6565b6000918252602080832090910154835460018181018655948452919092200155546040517f3f57d78cc64cf92f6e3395a2500653142d55d8ccce3c2822cc104fc2f7aa7fbb91610ce6918b908b908b908b90339042908d908d908d908d90611b38565b60405180910390a15b5050505050505050565b600080600090507f07ecc4d2c84ccebeededec1ffd783fb95395d34ea2aec93c939cd704166d409b428785604051610d3393929190611b0f565b60405180910390a160015b60008481548110610d5157610d51611aa6565b90600052602060002001548111610de3576000610d6e8286611146565b6000868152600360209081526040808320868452909152902060090154909150610d9a888389846102b6565b935060028681548110610daf57610daf611aa6565b9060005260206000200154841115610dce5760019450505050506102fd565b50508080610ddb90611abc565b915050610d3e565b5060009695505050505050565b610df86111e1565b600360008381526020019081526020016000206000848152602001908152602001600020604051806101600160405290816000820154815260200160018201548152602001600282018054610e4c90611ad5565b80601f0160208091040260200160405190810160405280929190818152602001828054610e7890611ad5565b8015610ec55780601f10610e9a57610100808354040283529160200191610ec5565b820191906000526020600020905b815481529060010190602001808311610ea857829003601f168201915b50505050508152602001600382018054610ede90611ad5565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0a90611ad5565b8015610f575780601f10610f2c57610100808354040283529160200191610f57565b820191906000526020600020905b815481529060010190602001808311610f3a57829003601f168201915b50505050508152602001600482018054610f7090611ad5565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9c90611ad5565b8015610fe95780601f10610fbe57610100808354040283529160200191610fe9565b820191906000526020600020905b815481529060010190602001808311610fcc57829003601f168201915b505050918352505060058201546001600160a01b031660208201526006820154604082015260078201805460609092019161102390611ad5565b80601f016020809104026020016040519081016040528092919081815260200182805461104f90611ad5565b801561109c5780601f106110715761010080835404028352916020019161109c565b820191906000526020600020905b81548152906001019060200180831161107f57829003601f168201915b505050505081526020016008820180548060200260200160405190810160405280929190818152602001828054801561112157602002820191906000526020600020906000905b82829054906101000a900460801b6001600160801b03191681526020019060100190602082600f010492830192600103820291508084116110e35790505b5050505050815260200160098201548152602001600a82015481525050905092915050565b60008181526003602090815260408083208584528252918290206008018054835181840281018401909452808452606093928301828280156111d457602002820191906000526020600020906000905b82829054906101000a900460801b6001600160801b03191681526020019060100190602082600f010492830192600103820291508084116111965790505b5050505050905092915050565b604051806101600160405280600081526020016000815260200160608152602001606081526020016060815260200160006001600160a01b0316815260200160008152602001606081526020016060815260200160008152602001600081525090565b82805461125090611ad5565b90600052602060002090601f01602090048101928261127257600085556112b8565b82601f1061128b57805160ff19168380011785556112b8565b828001600101855582156112b8579182015b828111156112b857825182559160200191906001019061129d565b506112c4929150611371565b5090565b828054828255906000526020600020906001016002900481019282156112b85791602002820160005b8382111561133557835183826101000a8154816001600160801b03021916908360801c02179055509260200192601001602081600f010492830192600103026112f1565b80156113685782816101000a8154906001600160801b030219169055601001602081600f01049283019260010302611335565b50506112c49291505b5b808211156112c45760008155600101611372565b60006020828403121561139857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156113de576113de61139f565b604052919050565b600082601f8301126113f757600080fd5b8135602067ffffffffffffffff8211156114135761141361139f565b8160051b6114228282016113b5565b928352848101820192828101908785111561143c57600080fd5b83870192505b848310156114725782356001600160801b0319811681146114635760008081fd5b82529183019190830190611442565b979650505050505050565b6000806000806080858703121561149357600080fd5b843567ffffffffffffffff808211156114ab57600080fd5b6114b7888389016113e6565b955060208701359150808211156114cd57600080fd5b506114da878288016113e6565b949794965050505060408301359260600135919050565b600082601f83011261150257600080fd5b813567ffffffffffffffff81111561151c5761151c61139f565b61152f601f8201601f19166020016113b5565b81815284602083860101111561154457600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561157357600080fd5b813567ffffffffffffffff81111561158a57600080fd5b6102fd848285016114f1565b60005b838110156115b1578181015183820152602001611599565b838111156115c0576000848401525b50505050565b600081518084526115de816020860160208601611596565b601f01601f19169290920160200192915050565b600081518084526020808501945080840160005b8381101561162c5781516001600160801b03191687529582019590820190600101611606565b509495945050505050565b600061016082518452602083015160208501526040830151816040860152611661828601826115c6565b9150506060830151848203606086015261167b82826115c6565b9150506080830151848203608086015261169582826115c6565b91505060a08301516116b260a08601826001600160a01b03169052565b5060c083015160c085015260e083015184820360e08601526116d482826115c6565b91505061010080840151858303828701526116ef83826115f2565b610120868101519088015261014095860151959096019490945250929392505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561176757603f19888603018452611755858351611637565b94509285019290850190600101611739565b5092979650505050505050565b6000806040838503121561178757600080fd5b50508035926020909101359150565b60006101408c83528b60208401528060408401526117b68184018c6115c6565b905082810360608401526117ca818b6115c6565b905082810360808401526117de818a6115c6565b6001600160a01b03891660a085015260c0840188905283810360e0850152905061180881876115c6565b6101008401959095525050610120015298975050505050505050565b80356001600160a01b038116811461183b57600080fd5b919050565b60008060006060848603121561185557600080fd5b61185e84611824565b95602085013595506040909401359392505050565b6000806040838503121561188657600080fd5b823567ffffffffffffffff81111561189d57600080fd5b6118a9858286016114f1565b9250506118b860208401611824565b90509250929050565b600080600080600080600080610100898b0312156118de57600080fd5b88359750602089013567ffffffffffffffff808211156118fd57600080fd5b6119098c838d016114f1565b985060408b013591508082111561191f57600080fd5b61192b8c838d016114f1565b975060608b013591508082111561194157600080fd5b61194d8c838d016114f1565b965060808b013591508082111561196357600080fd5b61196f8c838d016114f1565b955060a08b013591508082111561198557600080fd5b506119928b828c016113e6565b93505060c0890135915060e089013590509295985092959890939650565b600080600080608085870312156119c657600080fd5b843567ffffffffffffffff808211156119de57600080fd5b6114b7888389016114f1565b6020815260006119fd6020830184611637565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60008219821115611a2d57611a2d611a04565b500190565b600082821015611a4457611a44611a04565b500390565b6000816000190483118215151615611a6357611a63611a04565b500290565b600082611a8557634e487b7160e01b600052601260045260246000fd5b500490565b60008251611a9c818460208701611596565b9190910192915050565b634e487b7160e01b600052603260045260246000fd5b600060018201611ace57611ace611a04565b5060010190565b600181811c90821680611ae957607f821691505b602082108103611b0957634e487b7160e01b600052602260045260246000fd5b50919050565b838152606060208201526000611b2860608301856115c6565b9050826040830152949350505050565b60006101608d835260208d81850152816040850152611b598285018e6115c6565b91508382036060850152611b6d828d6115c6565b91508382036080850152611b81828c6115c6565b6001600160a01b038b1660a086015260c085018a905284810360e08601529150611bab82896115c6565b848103610100860152875180825282890193509082019060005b81811015611beb5784516001600160801b03191683529383019391830191600101611bc5565b50506101208501969096525050506101400152999850505050505050505056fea2646970667358221220cce77ca7fc4a48c645cbcb373e00dab65ecf7b594132deddc872c253c4bc11af64736f6c634300080d0033",
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

// DoesUserHavePermission is a free data retrieval call binding the contract method 0xa89796cc.
//
// Solidity: function doesUserHavePermission(string cid, address userAddress) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCaller) DoesUserHavePermission(opts *bind.CallOpts, cid string, userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "doesUserHavePermission", cid, userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesUserHavePermission is a free data retrieval call binding the contract method 0xa89796cc.
//
// Solidity: function doesUserHavePermission(string cid, address userAddress) view returns(bool)
func (_PlagiarismContract *PlagiarismContractSession) DoesUserHavePermission(cid string, userAddress common.Address) (bool, error) {
	return _PlagiarismContract.Contract.DoesUserHavePermission(&_PlagiarismContract.CallOpts, cid, userAddress)
}

// DoesUserHavePermission is a free data retrieval call binding the contract method 0xa89796cc.
//
// Solidity: function doesUserHavePermission(string cid, address userAddress) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCallerSession) DoesUserHavePermission(cid string, userAddress common.Address) (bool, error) {
	return _PlagiarismContract.Contract.DoesUserHavePermission(&_PlagiarismContract.CallOpts, cid, userAddress)
}

// FileCountByLanguage is a free data retrieval call binding the contract method 0x0641e449.
//
// Solidity: function fileCountByLanguage(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) FileCountByLanguage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "fileCountByLanguage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FileCountByLanguage is a free data retrieval call binding the contract method 0x0641e449.
//
// Solidity: function fileCountByLanguage(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) FileCountByLanguage(arg0 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.FileCountByLanguage(&_PlagiarismContract.CallOpts, arg0)
}

// FileCountByLanguage is a free data retrieval call binding the contract method 0x0641e449.
//
// Solidity: function fileCountByLanguage(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) FileCountByLanguage(arg0 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.FileCountByLanguage(&_PlagiarismContract.CallOpts, arg0)
}

// FilesMapByLanguage is a free data retrieval call binding the contract method 0x60ff2aac.
//
// Solidity: function filesMapByLanguage(uint256 , uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength, uint256 language)
func (_PlagiarismContract *PlagiarismContractCaller) FilesMapByLanguage(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSCID     string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
	Language        *big.Int
}, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "filesMapByLanguage", arg0, arg1)

	outstruct := new(struct {
		FileId          *big.Int
		FileSize        *big.Int
		FileIPFSCID     string
		FileName        string
		FileDescription string
		CodeAuthor      common.Address
		TimeUploaded    *big.Int
		CodeFingerPrint string
		HashSetLength   *big.Int
		Language        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FileSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FileIPFSCID = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.FileName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.FileDescription = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CodeAuthor = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TimeUploaded = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CodeFingerPrint = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.HashSetLength = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Language = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FilesMapByLanguage is a free data retrieval call binding the contract method 0x60ff2aac.
//
// Solidity: function filesMapByLanguage(uint256 , uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength, uint256 language)
func (_PlagiarismContract *PlagiarismContractSession) FilesMapByLanguage(arg0 *big.Int, arg1 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSCID     string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
	Language        *big.Int
}, error) {
	return _PlagiarismContract.Contract.FilesMapByLanguage(&_PlagiarismContract.CallOpts, arg0, arg1)
}

// FilesMapByLanguage is a free data retrieval call binding the contract method 0x60ff2aac.
//
// Solidity: function filesMapByLanguage(uint256 , uint256 ) view returns(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, uint256 hashSetLength, uint256 language)
func (_PlagiarismContract *PlagiarismContractCallerSession) FilesMapByLanguage(arg0 *big.Int, arg1 *big.Int) (struct {
	FileId          *big.Int
	FileSize        *big.Int
	FileIPFSCID     string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSetLength   *big.Int
	Language        *big.Int
}, error) {
	return _PlagiarismContract.Contract.FilesMapByLanguage(&_PlagiarismContract.CallOpts, arg0, arg1)
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

// GetFileByIndex is a free data retrieval call binding the contract method 0xee7e64f7.
//
// Solidity: function getFileByIndex(uint256 _fileIndex, uint256 language) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256) records)
func (_PlagiarismContract *PlagiarismContractCaller) GetFileByIndex(opts *bind.CallOpts, _fileIndex *big.Int, language *big.Int) (PlagiarismContractCodeFile, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getFileByIndex", _fileIndex, language)

	if err != nil {
		return *new(PlagiarismContractCodeFile), err
	}

	out0 := *abi.ConvertType(out[0], new(PlagiarismContractCodeFile)).(*PlagiarismContractCodeFile)

	return out0, err

}

// GetFileByIndex is a free data retrieval call binding the contract method 0xee7e64f7.
//
// Solidity: function getFileByIndex(uint256 _fileIndex, uint256 language) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256) records)
func (_PlagiarismContract *PlagiarismContractSession) GetFileByIndex(_fileIndex *big.Int, language *big.Int) (PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFileByIndex(&_PlagiarismContract.CallOpts, _fileIndex, language)
}

// GetFileByIndex is a free data retrieval call binding the contract method 0xee7e64f7.
//
// Solidity: function getFileByIndex(uint256 _fileIndex, uint256 language) view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256) records)
func (_PlagiarismContract *PlagiarismContractCallerSession) GetFileByIndex(_fileIndex *big.Int, language *big.Int) (PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFileByIndex(&_PlagiarismContract.CallOpts, _fileIndex, language)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) GetFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) GetFileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.GetFileCount(&_PlagiarismContract.CallOpts)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) GetFileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.GetFileCount(&_PlagiarismContract.CallOpts)
}

// GetFilesUploadedByUser is a free data retrieval call binding the contract method 0x51887b71.
//
// Solidity: function getFilesUploadedByUser() view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[])
func (_PlagiarismContract *PlagiarismContractCaller) GetFilesUploadedByUser(opts *bind.CallOpts) ([]PlagiarismContractCodeFile, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "getFilesUploadedByUser")

	if err != nil {
		return *new([]PlagiarismContractCodeFile), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlagiarismContractCodeFile)).(*[]PlagiarismContractCodeFile)

	return out0, err

}

// GetFilesUploadedByUser is a free data retrieval call binding the contract method 0x51887b71.
//
// Solidity: function getFilesUploadedByUser() view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[])
func (_PlagiarismContract *PlagiarismContractSession) GetFilesUploadedByUser() ([]PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFilesUploadedByUser(&_PlagiarismContract.CallOpts)
}

// GetFilesUploadedByUser is a free data retrieval call binding the contract method 0x51887b71.
//
// Solidity: function getFilesUploadedByUser() view returns((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[])
func (_PlagiarismContract *PlagiarismContractCallerSession) GetFilesUploadedByUser() ([]PlagiarismContractCodeFile, error) {
	return _PlagiarismContract.Contract.GetFilesUploadedByUser(&_PlagiarismContract.CallOpts)
}

// HasPermissionsMap is a free data retrieval call binding the contract method 0x9e199566.
//
// Solidity: function hasPermissionsMap(string , address ) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCaller) HasPermissionsMap(opts *bind.CallOpts, arg0 string, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "hasPermissionsMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermissionsMap is a free data retrieval call binding the contract method 0x9e199566.
//
// Solidity: function hasPermissionsMap(string , address ) view returns(bool)
func (_PlagiarismContract *PlagiarismContractSession) HasPermissionsMap(arg0 string, arg1 common.Address) (bool, error) {
	return _PlagiarismContract.Contract.HasPermissionsMap(&_PlagiarismContract.CallOpts, arg0, arg1)
}

// HasPermissionsMap is a free data retrieval call binding the contract method 0x9e199566.
//
// Solidity: function hasPermissionsMap(string , address ) view returns(bool)
func (_PlagiarismContract *PlagiarismContractCallerSession) HasPermissionsMap(arg0 string, arg1 common.Address) (bool, error) {
	return _PlagiarismContract.Contract.HasPermissionsMap(&_PlagiarismContract.CallOpts, arg0, arg1)
}

// Threshold is a free data retrieval call binding the contract method 0x31eb7a1a.
//
// Solidity: function threshold(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) Threshold(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x31eb7a1a.
//
// Solidity: function threshold(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) Threshold(arg0 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.Threshold(&_PlagiarismContract.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x31eb7a1a.
//
// Solidity: function threshold(uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) Threshold(arg0 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.Threshold(&_PlagiarismContract.CallOpts, arg0)
}

// TotalFileCount is a free data retrieval call binding the contract method 0x7f9d4b4a.
//
// Solidity: function totalFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) TotalFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "totalFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFileCount is a free data retrieval call binding the contract method 0x7f9d4b4a.
//
// Solidity: function totalFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) TotalFileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.TotalFileCount(&_PlagiarismContract.CallOpts)
}

// TotalFileCount is a free data retrieval call binding the contract method 0x7f9d4b4a.
//
// Solidity: function totalFileCount() view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) TotalFileCount() (*big.Int, error) {
	return _PlagiarismContract.Contract.TotalFileCount(&_PlagiarismContract.CallOpts)
}

// UserFilesIndexByLanguage is a free data retrieval call binding the contract method 0x97f42b1a.
//
// Solidity: function userFilesIndexByLanguage(address , uint256 , uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCaller) UserFilesIndexByLanguage(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlagiarismContract.contract.Call(opts, &out, "userFilesIndexByLanguage", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFilesIndexByLanguage is a free data retrieval call binding the contract method 0x97f42b1a.
//
// Solidity: function userFilesIndexByLanguage(address , uint256 , uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractSession) UserFilesIndexByLanguage(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.UserFilesIndexByLanguage(&_PlagiarismContract.CallOpts, arg0, arg1, arg2)
}

// UserFilesIndexByLanguage is a free data retrieval call binding the contract method 0x97f42b1a.
//
// Solidity: function userFilesIndexByLanguage(address , uint256 , uint256 ) view returns(uint256)
func (_PlagiarismContract *PlagiarismContractCallerSession) UserFilesIndexByLanguage(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _PlagiarismContract.Contract.UserFilesIndexByLanguage(&_PlagiarismContract.CallOpts, arg0, arg1, arg2)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5836e03d.
//
// Solidity: function addPermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractTransactor) AddPermission(opts *bind.TransactOpts, cid string) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "addPermission", cid)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5836e03d.
//
// Solidity: function addPermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractSession) AddPermission(cid string) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.AddPermission(&_PlagiarismContract.TransactOpts, cid)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5836e03d.
//
// Solidity: function addPermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractTransactorSession) AddPermission(cid string) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.AddPermission(&_PlagiarismContract.TransactOpts, cid)
}

// CheckIfPlagiarised is a paid mutator transaction binding the contract method 0xd11e44d6.
//
// Solidity: function checkIfPlagiarised(string _fileName, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns(bool)
func (_PlagiarismContract *PlagiarismContractTransactor) CheckIfPlagiarised(opts *bind.TransactOpts, _fileName string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "checkIfPlagiarised", _fileName, _hashSet, hashSetLength, language)
}

// CheckIfPlagiarised is a paid mutator transaction binding the contract method 0xd11e44d6.
//
// Solidity: function checkIfPlagiarised(string _fileName, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns(bool)
func (_PlagiarismContract *PlagiarismContractSession) CheckIfPlagiarised(_fileName string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.CheckIfPlagiarised(&_PlagiarismContract.TransactOpts, _fileName, _hashSet, hashSetLength, language)
}

// CheckIfPlagiarised is a paid mutator transaction binding the contract method 0xd11e44d6.
//
// Solidity: function checkIfPlagiarised(string _fileName, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns(bool)
func (_PlagiarismContract *PlagiarismContractTransactorSession) CheckIfPlagiarised(_fileName string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.CheckIfPlagiarised(&_PlagiarismContract.TransactOpts, _fileName, _hashSet, hashSetLength, language)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x4217649f.
//
// Solidity: function removePermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractTransactor) RemovePermission(opts *bind.TransactOpts, cid string) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "removePermission", cid)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x4217649f.
//
// Solidity: function removePermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractSession) RemovePermission(cid string) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.RemovePermission(&_PlagiarismContract.TransactOpts, cid)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x4217649f.
//
// Solidity: function removePermission(string cid) returns()
func (_PlagiarismContract *PlagiarismContractTransactorSession) RemovePermission(cid string) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.RemovePermission(&_PlagiarismContract.TransactOpts, cid)
}

// UploadFile is a paid mutator transaction binding the contract method 0xca749bec.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSCID, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns()
func (_PlagiarismContract *PlagiarismContractTransactor) UploadFile(opts *bind.TransactOpts, _fileSize *big.Int, _fileIPFSCID string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.contract.Transact(opts, "uploadFile", _fileSize, _fileIPFSCID, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength, language)
}

// UploadFile is a paid mutator transaction binding the contract method 0xca749bec.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSCID, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns()
func (_PlagiarismContract *PlagiarismContractSession) UploadFile(_fileSize *big.Int, _fileIPFSCID string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.UploadFile(&_PlagiarismContract.TransactOpts, _fileSize, _fileIPFSCID, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength, language)
}

// UploadFile is a paid mutator transaction binding the contract method 0xca749bec.
//
// Solidity: function uploadFile(uint256 _fileSize, string _fileIPFSCID, string _fileName, string _fileDescription, string _codeFingerPrint, bytes16[] _hashSet, uint256 hashSetLength, uint256 language) returns()
func (_PlagiarismContract *PlagiarismContractTransactorSession) UploadFile(_fileSize *big.Int, _fileIPFSCID string, _fileName string, _fileDescription string, _codeFingerPrint string, _hashSet [][16]byte, hashSetLength *big.Int, language *big.Int) (*types.Transaction, error) {
	return _PlagiarismContract.Contract.UploadFile(&_PlagiarismContract.TransactOpts, _fileSize, _fileIPFSCID, _fileName, _fileDescription, _codeFingerPrint, _hashSet, hashSetLength, language)
}

// PlagiarismContractCheckingPlagiarismIterator is returned from FilterCheckingPlagiarism and is used to iterate over the raw logs and unpacked data for CheckingPlagiarism events raised by the PlagiarismContract contract.
type PlagiarismContractCheckingPlagiarismIterator struct {
	Event *PlagiarismContractCheckingPlagiarism // Event containing the contract specifics and raw log

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
func (it *PlagiarismContractCheckingPlagiarismIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlagiarismContractCheckingPlagiarism)
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
		it.Event = new(PlagiarismContractCheckingPlagiarism)
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
func (it *PlagiarismContractCheckingPlagiarismIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlagiarismContractCheckingPlagiarismIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlagiarismContractCheckingPlagiarism represents a CheckingPlagiarism event raised by the PlagiarismContract contract.
type PlagiarismContractCheckingPlagiarism struct {
	TimeUploaded *big.Int
	FileName     string
	Language     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCheckingPlagiarism is a free log retrieval operation binding the contract event 0x07ecc4d2c84ccebeededec1ffd783fb95395d34ea2aec93c939cd704166d409b.
//
// Solidity: event CheckingPlagiarism(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterCheckingPlagiarism(opts *bind.FilterOpts) (*PlagiarismContractCheckingPlagiarismIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "CheckingPlagiarism")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractCheckingPlagiarismIterator{contract: _PlagiarismContract.contract, event: "CheckingPlagiarism", logs: logs, sub: sub}, nil
}

// WatchCheckingPlagiarism is a free log subscription operation binding the contract event 0x07ecc4d2c84ccebeededec1ffd783fb95395d34ea2aec93c939cd704166d409b.
//
// Solidity: event CheckingPlagiarism(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) WatchCheckingPlagiarism(opts *bind.WatchOpts, sink chan<- *PlagiarismContractCheckingPlagiarism) (event.Subscription, error) {

	logs, sub, err := _PlagiarismContract.contract.WatchLogs(opts, "CheckingPlagiarism")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlagiarismContractCheckingPlagiarism)
				if err := _PlagiarismContract.contract.UnpackLog(event, "CheckingPlagiarism", log); err != nil {
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

// ParseCheckingPlagiarism is a log parse operation binding the contract event 0x07ecc4d2c84ccebeededec1ffd783fb95395d34ea2aec93c939cd704166d409b.
//
// Solidity: event CheckingPlagiarism(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) ParseCheckingPlagiarism(log types.Log) (*PlagiarismContractCheckingPlagiarism, error) {
	event := new(PlagiarismContractCheckingPlagiarism)
	if err := _PlagiarismContract.contract.UnpackLog(event, "CheckingPlagiarism", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	FileIPFSCID     string
	FileName        string
	FileDescription string
	CodeAuthor      common.Address
	TimeUploaded    *big.Int
	CodeFingerPrint string
	HashSet         [][16]byte
	HashSetLength   *big.Int
	Language        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCodeFileUploadEvent is a free log retrieval operation binding the contract event 0x3f57d78cc64cf92f6e3395a2500653142d55d8ccce3c2822cc104fc2f7aa7fbb.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterCodeFileUploadEvent(opts *bind.FilterOpts) (*PlagiarismContractCodeFileUploadEventIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "CodeFileUploadEvent")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractCodeFileUploadEventIterator{contract: _PlagiarismContract.contract, event: "CodeFileUploadEvent", logs: logs, sub: sub}, nil
}

// WatchCodeFileUploadEvent is a free log subscription operation binding the contract event 0x3f57d78cc64cf92f6e3395a2500653142d55d8ccce3c2822cc104fc2f7aa7fbb.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength, uint256 language)
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

// ParseCodeFileUploadEvent is a log parse operation binding the contract event 0x3f57d78cc64cf92f6e3395a2500653142d55d8ccce3c2822cc104fc2f7aa7fbb.
//
// Solidity: event CodeFileUploadEvent(uint256 fileId, uint256 fileSize, string fileIPFSCID, string fileName, string fileDescription, address codeAuthor, uint256 timeUploaded, string codeFingerPrint, bytes16[] hashSet, uint256 hashSetLength, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) ParseCodeFileUploadEvent(log types.Log) (*PlagiarismContractCodeFileUploadEvent, error) {
	event := new(PlagiarismContractCodeFileUploadEvent)
	if err := _PlagiarismContract.contract.UnpackLog(event, "CodeFileUploadEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlagiarismContractCodeSubmittedIterator is returned from FilterCodeSubmitted and is used to iterate over the raw logs and unpacked data for CodeSubmitted events raised by the PlagiarismContract contract.
type PlagiarismContractCodeSubmittedIterator struct {
	Event *PlagiarismContractCodeSubmitted // Event containing the contract specifics and raw log

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
func (it *PlagiarismContractCodeSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlagiarismContractCodeSubmitted)
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
		it.Event = new(PlagiarismContractCodeSubmitted)
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
func (it *PlagiarismContractCodeSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlagiarismContractCodeSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlagiarismContractCodeSubmitted represents a CodeSubmitted event raised by the PlagiarismContract contract.
type PlagiarismContractCodeSubmitted struct {
	TimeUploaded *big.Int
	FileName     string
	Language     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCodeSubmitted is a free log retrieval operation binding the contract event 0x90a4de5c450d2bbdd8a9564fb9646ea0df44cf5dbb8e4bed1c6f806b7b37e57e.
//
// Solidity: event CodeSubmitted(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterCodeSubmitted(opts *bind.FilterOpts) (*PlagiarismContractCodeSubmittedIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "CodeSubmitted")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractCodeSubmittedIterator{contract: _PlagiarismContract.contract, event: "CodeSubmitted", logs: logs, sub: sub}, nil
}

// WatchCodeSubmitted is a free log subscription operation binding the contract event 0x90a4de5c450d2bbdd8a9564fb9646ea0df44cf5dbb8e4bed1c6f806b7b37e57e.
//
// Solidity: event CodeSubmitted(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) WatchCodeSubmitted(opts *bind.WatchOpts, sink chan<- *PlagiarismContractCodeSubmitted) (event.Subscription, error) {

	logs, sub, err := _PlagiarismContract.contract.WatchLogs(opts, "CodeSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlagiarismContractCodeSubmitted)
				if err := _PlagiarismContract.contract.UnpackLog(event, "CodeSubmitted", log); err != nil {
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

// ParseCodeSubmitted is a log parse operation binding the contract event 0x90a4de5c450d2bbdd8a9564fb9646ea0df44cf5dbb8e4bed1c6f806b7b37e57e.
//
// Solidity: event CodeSubmitted(uint256 timeUploaded, string fileName, uint256 language)
func (_PlagiarismContract *PlagiarismContractFilterer) ParseCodeSubmitted(log types.Log) (*PlagiarismContractCodeSubmitted, error) {
	event := new(PlagiarismContractCodeSubmitted)
	if err := _PlagiarismContract.contract.UnpackLog(event, "CodeSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlagiarismContractGetFilesEventIterator is returned from FilterGetFilesEvent and is used to iterate over the raw logs and unpacked data for GetFilesEvent events raised by the PlagiarismContract contract.
type PlagiarismContractGetFilesEventIterator struct {
	Event *PlagiarismContractGetFilesEvent // Event containing the contract specifics and raw log

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
func (it *PlagiarismContractGetFilesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlagiarismContractGetFilesEvent)
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
		it.Event = new(PlagiarismContractGetFilesEvent)
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
func (it *PlagiarismContractGetFilesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlagiarismContractGetFilesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlagiarismContractGetFilesEvent represents a GetFilesEvent event raised by the PlagiarismContract contract.
type PlagiarismContractGetFilesEvent struct {
	Codefiles []PlagiarismContractCodeFile
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGetFilesEvent is a free log retrieval operation binding the contract event 0x5dcc3d1516b39e7bac4a3d16089018742f5746f671dd966fd7271b32bd476c37.
//
// Solidity: event GetFilesEvent((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[] codefiles)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterGetFilesEvent(opts *bind.FilterOpts) (*PlagiarismContractGetFilesEventIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "GetFilesEvent")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractGetFilesEventIterator{contract: _PlagiarismContract.contract, event: "GetFilesEvent", logs: logs, sub: sub}, nil
}

// WatchGetFilesEvent is a free log subscription operation binding the contract event 0x5dcc3d1516b39e7bac4a3d16089018742f5746f671dd966fd7271b32bd476c37.
//
// Solidity: event GetFilesEvent((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[] codefiles)
func (_PlagiarismContract *PlagiarismContractFilterer) WatchGetFilesEvent(opts *bind.WatchOpts, sink chan<- *PlagiarismContractGetFilesEvent) (event.Subscription, error) {

	logs, sub, err := _PlagiarismContract.contract.WatchLogs(opts, "GetFilesEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlagiarismContractGetFilesEvent)
				if err := _PlagiarismContract.contract.UnpackLog(event, "GetFilesEvent", log); err != nil {
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

// ParseGetFilesEvent is a log parse operation binding the contract event 0x5dcc3d1516b39e7bac4a3d16089018742f5746f671dd966fd7271b32bd476c37.
//
// Solidity: event GetFilesEvent((uint256,uint256,string,string,string,address,uint256,string,bytes16[],uint256,uint256)[] codefiles)
func (_PlagiarismContract *PlagiarismContractFilterer) ParseGetFilesEvent(log types.Log) (*PlagiarismContractGetFilesEvent, error) {
	event := new(PlagiarismContractGetFilesEvent)
	if err := _PlagiarismContract.contract.UnpackLog(event, "GetFilesEvent", log); err != nil {
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
	TimeUploaded      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPlagiarismResult is a free log retrieval operation binding the contract event 0x33baf730752a6c778112e3ee4c2e8bd128e07d02117675aa7c7feec6dd8b51db.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult, uint256 timeUploaded)
func (_PlagiarismContract *PlagiarismContractFilterer) FilterPlagiarismResult(opts *bind.FilterOpts) (*PlagiarismContractPlagiarismResultIterator, error) {

	logs, sub, err := _PlagiarismContract.contract.FilterLogs(opts, "PlagiarismResult")
	if err != nil {
		return nil, err
	}
	return &PlagiarismContractPlagiarismResultIterator{contract: _PlagiarismContract.contract, event: "PlagiarismResult", logs: logs, sub: sub}, nil
}

// WatchPlagiarismResult is a free log subscription operation binding the contract event 0x33baf730752a6c778112e3ee4c2e8bd128e07d02117675aa7c7feec6dd8b51db.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult, uint256 timeUploaded)
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

// ParsePlagiarismResult is a log parse operation binding the contract event 0x33baf730752a6c778112e3ee4c2e8bd128e07d02117675aa7c7feec6dd8b51db.
//
// Solidity: event PlagiarismResult(bool plagiarisedResult, uint256 timeUploaded)
func (_PlagiarismContract *PlagiarismContractFilterer) ParsePlagiarismResult(log types.Log) (*PlagiarismContractPlagiarismResult, error) {
	event := new(PlagiarismContractPlagiarismResult)
	if err := _PlagiarismContract.contract.UnpackLog(event, "PlagiarismResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
