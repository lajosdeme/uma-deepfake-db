// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DeepFakeDetectorABI is the input ABI used to generate the binding from.
const DeepFakeDetectorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"imageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"ImageAsserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"imageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"ImageAssertionResolved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"imageId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"}],\"name\":\"assertDataFor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"assertionDisputedCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertionLiveness\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"assertedTruthfully\",\"type\":\"bool\"}],\"name\":\"assertionResolvedCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCurrency\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getImage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oo\",\"outputs\":[{\"internalType\":\"contractOptimisticOracleV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DeepFakeDetectorBin is the compiled bytecode used for deploying new contracts.
var DeepFakeDetectorBin = "0x60e060405234801561001057600080fd5b507307865c6e87b9f70255377e024ace6630c1eaa37f608052739923d42ef695b5dd9911d05ac944d4caca3c4eab60a08190526040805163d509b01760e01b8152905163d509b017916004808201926020929091908290030181865afa15801561007e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100a291906100aa565b60c0526100c3565b6000602082840312156100bc57600080fd5b5051919050565b60805160a05160c0516115a26101306000396000818161017401526103c801526000818161011801528181610207015281816102d8015281816102ff015261062e01526000818160b8015281816101d901528181610281015281816102b601526103a401526115a26000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b75d4e241161005b578063b75d4e241461013a578063d448a4ec1461015c578063d509b0171461016f578063f1b156b21461019657600080fd5b8063096464cc1461008d57806320402e1d146100b35780636ced1ae9146100f25780638921a61414610113575b600080fd5b6100a061009b366004610f77565b6101a9565b6040519081526020015b60405180910390f35b6100da7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100aa565b610105610100366004611012565b61053f565b6040516100aa92919061107b565b6100da7f000000000000000000000000000000000000000000000000000000000000000081565b610143610e1081565b60405167ffffffffffffffff90911681526020016100aa565b61016d61016a366004611012565b50565b005b6100a07f000000000000000000000000000000000000000000000000000000000000000081565b61016d6101a43660046110a4565b610623565b60006001600160a01b038216156101c057816101c2565b335b604051634360af3d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529193506000917f00000000000000000000000000000000000000000000000000000000000000001690634360af3d90602401602060405180830381865afa15801561024e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027291906110d4565b90506102a96001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330846107d7565b6102fd6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000083610848565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636457c97986866103378a610967565b610340886109a8565b610349896109a8565b610352426109ff565b61035b306109a8565b6040516020016103719796959493929190611109565b60408051601f19818403018152908290526001600160e01b031960e084901b1682526103f29187903090600090610e10907f0000000000000000000000000000000000000000000000000000000000000000908a907f0000000000000000000000000000000000000000000000000000000000000000908590600401611250565b6020604051808303816000875af1158015610411573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043591906110d4565b9150604051806080016040528087815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050506001600160a01b0386166020808401919091526040928301829052858252818152919020825181559082015160018201906104bb908261135c565b50604082810151600290920180546060909401511515600160a01b026001600160a81b03199094166001600160a01b0393841617939093179092559051839185169088907fb3b5c304bc1ba781ea735c86e786e9ae6432b86b469fbc00ec8bb252bf54eda89061052e908a908a9061141c565b60405180910390a450949350505050565b600081815260208190526040812060020154606090600160a01b900460ff1661057a5750506040805160208101909152600080825292909150565b600083815260208190526040902060019081018054819061059a906112d4565b80601f01602080910402602001604051908101604052809291908181526020018280546105c6906112d4565b80156106135780601f106105e857610100808354040283529160200191610613565b820191906000526020600020905b8154815290600101906020018083116105f657829003601f168201915b5050505050905091509150915091565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461065857600080fd5b80156107a25760008281526020818152604080832060028101805460ff60a01b1916600160a01b1790558151608081019092528054825260018101805492939192918401916106a6906112d4565b80601f01602080910402602001604051908101604052809291908181526020018280546106d2906112d4565b801561071f5780601f106106f45761010080835404028352916020019161071f565b820191906000526020600020905b81548152906001019060200180831161070257829003601f168201915b5050509183525050600291909101546001600160a01b03808216602080850191909152600160a01b90920460ff16151560409384015283830151845192850151935194955087949116927f2cd47d03ca5eed86ed11475382af40bc0f2965dd65bc4ab337d45577ea5dfaf991610795919061144b565b60405180910390a4505050565b6000828152602081905260408120818155906107c16001830182610f29565b5060020180546001600160a81b03191690555050565b6040516001600160a01b03808516602483015283166044820152606481018290526108429085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610b2b565b50505050565b8015806108c25750604051636eb1769f60e11b81523060048201526001600160a01b03838116602483015284169063dd62ed3e90604401602060405180830381865afa15801561089c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c091906110d4565b155b6109325760405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527520746f206e6f6e2d7a65726f20616c6c6f77616e636560501b60648201526084015b60405180910390fd5b6040516001600160a01b03831660248201526044810182905261096290849063095ea7b360e01b9060640161080b565b505050565b6060610976608083901c610c00565b61097f83610c00565b6040805160208101939093528201526060015b6040516020818303038152906040529050919050565b60606109c06001600160801b03602084901c16610c00565b6109db8360601b6bffffffffffffffffffffffff1916610c00565b6040516020016109929291909182526001600160c01b031916602082015260280190565b606081600003610a265750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610a505780610a3a8161147b565b9150610a499050600a836114aa565b9150610a2a565b60008167ffffffffffffffff811115610a6b57610a6b6112be565b6040519080825280601f01601f191660200182016040528015610a95576020820181803683370190505b509050815b8515610b2257610aab6001826114cc565b90506000610aba600a886114aa565b610ac590600a6114e5565b610acf90886114cc565b610ada906030611504565b905060008160f81b905080848481518110610af757610af761151d565b60200101906001600160f81b031916908160001a905350610b19600a896114aa565b97505050610a9a565b50949350505050565b6000610b80826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610d999092919063ffffffff16565b9050805160001480610ba1575080806020019051810190610ba19190611533565b6109625760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610929565b6000808260001c9050806001600160801b03169050806801000000000000000002811777ffffffffffffffff0000000000000000ffffffffffffffff169050806401000000000281177bffffffff00000000ffffffff00000000ffffffff00000000ffffffff16905080620100000281177dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff169050806101000281177eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff1690508060100281177f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f16905060006008827f08080808080808080808080808080808080808080808080808080808080808081681610d1b57610d1b611494565b0460047f040404040404040404040404040404040404040404040404040404040404040484160460027f020202020202020202020202020202020202020202020202020202020202020285160417166027029091017f3030303030303030303030303030303030303030303030303030303030303030019392505050565b6060610da88484600085610db0565b949350505050565b606082471015610e115760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610929565b600080866001600160a01b03168587604051610e2d9190611550565b60006040518083038185875af1925050503d8060008114610e6a576040519150601f19603f3d011682016040523d82523d6000602084013e610e6f565b606091505b5091509150610e8087838387610e8b565b979650505050505050565b60608315610efa578251600003610ef3576001600160a01b0385163b610ef35760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610929565b5081610da8565b610da88383815115610f0f5781518083602001fd5b8060405162461bcd60e51b8152600401610929919061144b565b508054610f35906112d4565b6000825580601f10610f45575050565b601f01602090049060005260206000209081019061016a91905b80821115610f735760008155600101610f5f565b5090565b60008060008060608587031215610f8d57600080fd5b84359350602085013567ffffffffffffffff80821115610fac57600080fd5b818701915087601f830112610fc057600080fd5b813581811115610fcf57600080fd5b886020828501011115610fe157600080fd5b60209290920194509092505060408501356001600160a01b038116811461100757600080fd5b939692955090935050565b60006020828403121561102457600080fd5b5035919050565b60005b8381101561104657818101518382015260200161102e565b50506000910152565b6000815180845261106781602086016020860161102b565b601f01601f19169290920160200192915050565b8215158152604060208201526000610da8604083018461104f565b801515811461016a57600080fd5b600080604083850312156110b757600080fd5b8235915060208301356110c981611096565b809150509250929050565b6000602082840312156110e657600080fd5b5051919050565b600081516110ff81856020860161102b565b9290920192915050565b7f496d6167652061737365727465642061742075726c3a2000000000000000000081528688601783013760008782016f040ccdee440d2dac2ceca9288744060f60831b60178201528751611164816027840160208c0161102b565b70040c2dcc840c2e6e6cae4e8cae4744060f607b1b602792909101918201528651611196816038840160208b0161102b565b86519101906111ac816038840160208a0161102b565b6e01030ba103a34b6b2b9ba30b6b81d1608d1b6038929091019182015284516111dc81604784016020890161102b565b7f20696e207468652044617461417373657274657220636f6e7472616374206174604792909101918201526204060f60eb1b6067820152611242611223606a8301866110ed565b721034b9903737ba1030903232b2b83330b5b29760691b815260130190565b9a9950505050505050505050565b60006101208083526112648184018d61104f565b6001600160a01b039b8c166020850152998b1660408401525050958816606087015267ffffffffffffffff9490941660808601529190951660a084015260c083019490945260e08201939093526101000191909152919050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806112e857607f821691505b60208210810361130857634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561096257600081815260208120601f850160051c810160208610156113355750805b601f850160051c820191505b8181101561135457828155600101611341565b505050505050565b815167ffffffffffffffff811115611376576113766112be565b61138a8161138484546112d4565b8461130e565b602080601f8311600181146113bf57600084156113a75750858301515b600019600386901b1c1916600185901b178555611354565b600085815260208120601f198616915b828110156113ee578886015182559484019460019091019084016113cf565b508582101561140c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60208152600061145e602083018461104f565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60006001820161148d5761148d611465565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826114c757634e487b7160e01b600052601260045260246000fd5b500490565b818103818111156114df576114df611465565b92915050565b60008160001904831182151516156114ff576114ff611465565b500290565b60ff81811683821601908111156114df576114df611465565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561154557600080fd5b815161145e81611096565b6000825161156281846020870161102b565b919091019291505056fea264697066735822122022759d2ec01478f58482cfd77d95067e6904f299377122900a416de4b2e7c38464736f6c63430008100033"

// DeployDeepFakeDetector deploys a new Ethereum contract, binding an instance of DeepFakeDetector to it.
func DeployDeepFakeDetector(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeepFakeDetector, error) {
	parsed, err := abi.JSON(strings.NewReader(DeepFakeDetectorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeepFakeDetectorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeepFakeDetector{DeepFakeDetectorCaller: DeepFakeDetectorCaller{contract: contract}, DeepFakeDetectorTransactor: DeepFakeDetectorTransactor{contract: contract}, DeepFakeDetectorFilterer: DeepFakeDetectorFilterer{contract: contract}}, nil
}

// DeepFakeDetector is an auto generated Go binding around an Ethereum contract.
type DeepFakeDetector struct {
	DeepFakeDetectorCaller     // Read-only binding to the contract
	DeepFakeDetectorTransactor // Write-only binding to the contract
	DeepFakeDetectorFilterer   // Log filterer for contract events
}

// DeepFakeDetectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeepFakeDetectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeepFakeDetectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeepFakeDetectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeepFakeDetectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeepFakeDetectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeepFakeDetectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeepFakeDetectorSession struct {
	Contract     *DeepFakeDetector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeepFakeDetectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeepFakeDetectorCallerSession struct {
	Contract *DeepFakeDetectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DeepFakeDetectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeepFakeDetectorTransactorSession struct {
	Contract     *DeepFakeDetectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DeepFakeDetectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeepFakeDetectorRaw struct {
	Contract *DeepFakeDetector // Generic contract binding to access the raw methods on
}

// DeepFakeDetectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeepFakeDetectorCallerRaw struct {
	Contract *DeepFakeDetectorCaller // Generic read-only contract binding to access the raw methods on
}

// DeepFakeDetectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeepFakeDetectorTransactorRaw struct {
	Contract *DeepFakeDetectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeepFakeDetector creates a new instance of DeepFakeDetector, bound to a specific deployed contract.
func NewDeepFakeDetector(address common.Address, backend bind.ContractBackend) (*DeepFakeDetector, error) {
	contract, err := bindDeepFakeDetector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetector{DeepFakeDetectorCaller: DeepFakeDetectorCaller{contract: contract}, DeepFakeDetectorTransactor: DeepFakeDetectorTransactor{contract: contract}, DeepFakeDetectorFilterer: DeepFakeDetectorFilterer{contract: contract}}, nil
}

// NewDeepFakeDetectorCaller creates a new read-only instance of DeepFakeDetector, bound to a specific deployed contract.
func NewDeepFakeDetectorCaller(address common.Address, caller bind.ContractCaller) (*DeepFakeDetectorCaller, error) {
	contract, err := bindDeepFakeDetector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetectorCaller{contract: contract}, nil
}

// NewDeepFakeDetectorTransactor creates a new write-only instance of DeepFakeDetector, bound to a specific deployed contract.
func NewDeepFakeDetectorTransactor(address common.Address, transactor bind.ContractTransactor) (*DeepFakeDetectorTransactor, error) {
	contract, err := bindDeepFakeDetector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetectorTransactor{contract: contract}, nil
}

// NewDeepFakeDetectorFilterer creates a new log filterer instance of DeepFakeDetector, bound to a specific deployed contract.
func NewDeepFakeDetectorFilterer(address common.Address, filterer bind.ContractFilterer) (*DeepFakeDetectorFilterer, error) {
	contract, err := bindDeepFakeDetector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetectorFilterer{contract: contract}, nil
}

// bindDeepFakeDetector binds a generic wrapper to an already deployed contract.
func bindDeepFakeDetector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeepFakeDetectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeepFakeDetector *DeepFakeDetectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeepFakeDetector.Contract.DeepFakeDetectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeepFakeDetector *DeepFakeDetectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.DeepFakeDetectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeepFakeDetector *DeepFakeDetectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.DeepFakeDetectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeepFakeDetector *DeepFakeDetectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeepFakeDetector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeepFakeDetector *DeepFakeDetectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeepFakeDetector *DeepFakeDetectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.contract.Transact(opts, method, params...)
}

// AssertionLiveness is a free data retrieval call binding the contract method 0xb75d4e24.
//
// Solidity: function assertionLiveness() view returns(uint64)
func (_DeepFakeDetector *DeepFakeDetectorCaller) AssertionLiveness(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DeepFakeDetector.contract.Call(opts, &out, "assertionLiveness")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AssertionLiveness is a free data retrieval call binding the contract method 0xb75d4e24.
//
// Solidity: function assertionLiveness() view returns(uint64)
func (_DeepFakeDetector *DeepFakeDetectorSession) AssertionLiveness() (uint64, error) {
	return _DeepFakeDetector.Contract.AssertionLiveness(&_DeepFakeDetector.CallOpts)
}

// AssertionLiveness is a free data retrieval call binding the contract method 0xb75d4e24.
//
// Solidity: function assertionLiveness() view returns(uint64)
func (_DeepFakeDetector *DeepFakeDetectorCallerSession) AssertionLiveness() (uint64, error) {
	return _DeepFakeDetector.Contract.AssertionLiveness(&_DeepFakeDetector.CallOpts)
}

// DefaultCurrency is a free data retrieval call binding the contract method 0x20402e1d.
//
// Solidity: function defaultCurrency() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorCaller) DefaultCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeepFakeDetector.contract.Call(opts, &out, "defaultCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultCurrency is a free data retrieval call binding the contract method 0x20402e1d.
//
// Solidity: function defaultCurrency() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorSession) DefaultCurrency() (common.Address, error) {
	return _DeepFakeDetector.Contract.DefaultCurrency(&_DeepFakeDetector.CallOpts)
}

// DefaultCurrency is a free data retrieval call binding the contract method 0x20402e1d.
//
// Solidity: function defaultCurrency() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorCallerSession) DefaultCurrency() (common.Address, error) {
	return _DeepFakeDetector.Contract.DefaultCurrency(&_DeepFakeDetector.CallOpts)
}

// DefaultIdentifier is a free data retrieval call binding the contract method 0xd509b017.
//
// Solidity: function defaultIdentifier() view returns(bytes32)
func (_DeepFakeDetector *DeepFakeDetectorCaller) DefaultIdentifier(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeepFakeDetector.contract.Call(opts, &out, "defaultIdentifier")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DefaultIdentifier is a free data retrieval call binding the contract method 0xd509b017.
//
// Solidity: function defaultIdentifier() view returns(bytes32)
func (_DeepFakeDetector *DeepFakeDetectorSession) DefaultIdentifier() ([32]byte, error) {
	return _DeepFakeDetector.Contract.DefaultIdentifier(&_DeepFakeDetector.CallOpts)
}

// DefaultIdentifier is a free data retrieval call binding the contract method 0xd509b017.
//
// Solidity: function defaultIdentifier() view returns(bytes32)
func (_DeepFakeDetector *DeepFakeDetectorCallerSession) DefaultIdentifier() ([32]byte, error) {
	return _DeepFakeDetector.Contract.DefaultIdentifier(&_DeepFakeDetector.CallOpts)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(bytes32 assertionId) view returns(bool, string)
func (_DeepFakeDetector *DeepFakeDetectorCaller) GetImage(opts *bind.CallOpts, assertionId [32]byte) (bool, string, error) {
	var out []interface{}
	err := _DeepFakeDetector.contract.Call(opts, &out, "getImage", assertionId)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(bytes32 assertionId) view returns(bool, string)
func (_DeepFakeDetector *DeepFakeDetectorSession) GetImage(assertionId [32]byte) (bool, string, error) {
	return _DeepFakeDetector.Contract.GetImage(&_DeepFakeDetector.CallOpts, assertionId)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(bytes32 assertionId) view returns(bool, string)
func (_DeepFakeDetector *DeepFakeDetectorCallerSession) GetImage(assertionId [32]byte) (bool, string, error) {
	return _DeepFakeDetector.Contract.GetImage(&_DeepFakeDetector.CallOpts, assertionId)
}

// Oo is a free data retrieval call binding the contract method 0x8921a614.
//
// Solidity: function oo() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorCaller) Oo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeepFakeDetector.contract.Call(opts, &out, "oo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oo is a free data retrieval call binding the contract method 0x8921a614.
//
// Solidity: function oo() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorSession) Oo() (common.Address, error) {
	return _DeepFakeDetector.Contract.Oo(&_DeepFakeDetector.CallOpts)
}

// Oo is a free data retrieval call binding the contract method 0x8921a614.
//
// Solidity: function oo() view returns(address)
func (_DeepFakeDetector *DeepFakeDetectorCallerSession) Oo() (common.Address, error) {
	return _DeepFakeDetector.Contract.Oo(&_DeepFakeDetector.CallOpts)
}

// AssertDataFor is a paid mutator transaction binding the contract method 0x096464cc.
//
// Solidity: function assertDataFor(bytes32 imageId, string imageUrl, address asserter) returns(bytes32 assertionId)
func (_DeepFakeDetector *DeepFakeDetectorTransactor) AssertDataFor(opts *bind.TransactOpts, imageId [32]byte, imageUrl string, asserter common.Address) (*types.Transaction, error) {
	return _DeepFakeDetector.contract.Transact(opts, "assertDataFor", imageId, imageUrl, asserter)
}

// AssertDataFor is a paid mutator transaction binding the contract method 0x096464cc.
//
// Solidity: function assertDataFor(bytes32 imageId, string imageUrl, address asserter) returns(bytes32 assertionId)
func (_DeepFakeDetector *DeepFakeDetectorSession) AssertDataFor(imageId [32]byte, imageUrl string, asserter common.Address) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertDataFor(&_DeepFakeDetector.TransactOpts, imageId, imageUrl, asserter)
}

// AssertDataFor is a paid mutator transaction binding the contract method 0x096464cc.
//
// Solidity: function assertDataFor(bytes32 imageId, string imageUrl, address asserter) returns(bytes32 assertionId)
func (_DeepFakeDetector *DeepFakeDetectorTransactorSession) AssertDataFor(imageId [32]byte, imageUrl string, asserter common.Address) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertDataFor(&_DeepFakeDetector.TransactOpts, imageId, imageUrl, asserter)
}

// AssertionDisputedCallback is a paid mutator transaction binding the contract method 0xd448a4ec.
//
// Solidity: function assertionDisputedCallback(bytes32 assertionId) returns()
func (_DeepFakeDetector *DeepFakeDetectorTransactor) AssertionDisputedCallback(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _DeepFakeDetector.contract.Transact(opts, "assertionDisputedCallback", assertionId)
}

// AssertionDisputedCallback is a paid mutator transaction binding the contract method 0xd448a4ec.
//
// Solidity: function assertionDisputedCallback(bytes32 assertionId) returns()
func (_DeepFakeDetector *DeepFakeDetectorSession) AssertionDisputedCallback(assertionId [32]byte) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertionDisputedCallback(&_DeepFakeDetector.TransactOpts, assertionId)
}

// AssertionDisputedCallback is a paid mutator transaction binding the contract method 0xd448a4ec.
//
// Solidity: function assertionDisputedCallback(bytes32 assertionId) returns()
func (_DeepFakeDetector *DeepFakeDetectorTransactorSession) AssertionDisputedCallback(assertionId [32]byte) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertionDisputedCallback(&_DeepFakeDetector.TransactOpts, assertionId)
}

// AssertionResolvedCallback is a paid mutator transaction binding the contract method 0xf1b156b2.
//
// Solidity: function assertionResolvedCallback(bytes32 assertionId, bool assertedTruthfully) returns()
func (_DeepFakeDetector *DeepFakeDetectorTransactor) AssertionResolvedCallback(opts *bind.TransactOpts, assertionId [32]byte, assertedTruthfully bool) (*types.Transaction, error) {
	return _DeepFakeDetector.contract.Transact(opts, "assertionResolvedCallback", assertionId, assertedTruthfully)
}

// AssertionResolvedCallback is a paid mutator transaction binding the contract method 0xf1b156b2.
//
// Solidity: function assertionResolvedCallback(bytes32 assertionId, bool assertedTruthfully) returns()
func (_DeepFakeDetector *DeepFakeDetectorSession) AssertionResolvedCallback(assertionId [32]byte, assertedTruthfully bool) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertionResolvedCallback(&_DeepFakeDetector.TransactOpts, assertionId, assertedTruthfully)
}

// AssertionResolvedCallback is a paid mutator transaction binding the contract method 0xf1b156b2.
//
// Solidity: function assertionResolvedCallback(bytes32 assertionId, bool assertedTruthfully) returns()
func (_DeepFakeDetector *DeepFakeDetectorTransactorSession) AssertionResolvedCallback(assertionId [32]byte, assertedTruthfully bool) (*types.Transaction, error) {
	return _DeepFakeDetector.Contract.AssertionResolvedCallback(&_DeepFakeDetector.TransactOpts, assertionId, assertedTruthfully)
}

// DeepFakeDetectorImageAssertedIterator is returned from FilterImageAsserted and is used to iterate over the raw logs and unpacked data for ImageAsserted events raised by the DeepFakeDetector contract.
type DeepFakeDetectorImageAssertedIterator struct {
	Event *DeepFakeDetectorImageAsserted // Event containing the contract specifics and raw log

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
func (it *DeepFakeDetectorImageAssertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeepFakeDetectorImageAsserted)
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
		it.Event = new(DeepFakeDetectorImageAsserted)
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
func (it *DeepFakeDetectorImageAssertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeepFakeDetectorImageAssertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeepFakeDetectorImageAsserted represents a ImageAsserted event raised by the DeepFakeDetector contract.
type DeepFakeDetectorImageAsserted struct {
	ImageId     [32]byte
	ImageUrl    string
	Asserter    common.Address
	AssertionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterImageAsserted is a free log retrieval operation binding the contract event 0xb3b5c304bc1ba781ea735c86e786e9ae6432b86b469fbc00ec8bb252bf54eda8.
//
// Solidity: event ImageAsserted(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) FilterImageAsserted(opts *bind.FilterOpts, imageId [][32]byte, asserter []common.Address, assertionId [][32]byte) (*DeepFakeDetectorImageAssertedIterator, error) {

	var imageIdRule []interface{}
	for _, imageIdItem := range imageId {
		imageIdRule = append(imageIdRule, imageIdItem)
	}

	var asserterRule []interface{}
	for _, asserterItem := range asserter {
		asserterRule = append(asserterRule, asserterItem)
	}
	var assertionIdRule []interface{}
	for _, assertionIdItem := range assertionId {
		assertionIdRule = append(assertionIdRule, assertionIdItem)
	}

	logs, sub, err := _DeepFakeDetector.contract.FilterLogs(opts, "ImageAsserted", imageIdRule, asserterRule, assertionIdRule)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetectorImageAssertedIterator{contract: _DeepFakeDetector.contract, event: "ImageAsserted", logs: logs, sub: sub}, nil
}

// WatchImageAsserted is a free log subscription operation binding the contract event 0xb3b5c304bc1ba781ea735c86e786e9ae6432b86b469fbc00ec8bb252bf54eda8.
//
// Solidity: event ImageAsserted(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) WatchImageAsserted(opts *bind.WatchOpts, sink chan<- *DeepFakeDetectorImageAsserted, imageId [][32]byte, asserter []common.Address, assertionId [][32]byte) (event.Subscription, error) {

	var imageIdRule []interface{}
	for _, imageIdItem := range imageId {
		imageIdRule = append(imageIdRule, imageIdItem)
	}

	var asserterRule []interface{}
	for _, asserterItem := range asserter {
		asserterRule = append(asserterRule, asserterItem)
	}
	var assertionIdRule []interface{}
	for _, assertionIdItem := range assertionId {
		assertionIdRule = append(assertionIdRule, assertionIdItem)
	}

	logs, sub, err := _DeepFakeDetector.contract.WatchLogs(opts, "ImageAsserted", imageIdRule, asserterRule, assertionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeepFakeDetectorImageAsserted)
				if err := _DeepFakeDetector.contract.UnpackLog(event, "ImageAsserted", log); err != nil {
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

// ParseImageAsserted is a log parse operation binding the contract event 0xb3b5c304bc1ba781ea735c86e786e9ae6432b86b469fbc00ec8bb252bf54eda8.
//
// Solidity: event ImageAsserted(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) ParseImageAsserted(log types.Log) (*DeepFakeDetectorImageAsserted, error) {
	event := new(DeepFakeDetectorImageAsserted)
	if err := _DeepFakeDetector.contract.UnpackLog(event, "ImageAsserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeepFakeDetectorImageAssertionResolvedIterator is returned from FilterImageAssertionResolved and is used to iterate over the raw logs and unpacked data for ImageAssertionResolved events raised by the DeepFakeDetector contract.
type DeepFakeDetectorImageAssertionResolvedIterator struct {
	Event *DeepFakeDetectorImageAssertionResolved // Event containing the contract specifics and raw log

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
func (it *DeepFakeDetectorImageAssertionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeepFakeDetectorImageAssertionResolved)
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
		it.Event = new(DeepFakeDetectorImageAssertionResolved)
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
func (it *DeepFakeDetectorImageAssertionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeepFakeDetectorImageAssertionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeepFakeDetectorImageAssertionResolved represents a ImageAssertionResolved event raised by the DeepFakeDetector contract.
type DeepFakeDetectorImageAssertionResolved struct {
	ImageId     [32]byte
	ImageUrl    string
	Asserter    common.Address
	AssertionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterImageAssertionResolved is a free log retrieval operation binding the contract event 0x2cd47d03ca5eed86ed11475382af40bc0f2965dd65bc4ab337d45577ea5dfaf9.
//
// Solidity: event ImageAssertionResolved(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) FilterImageAssertionResolved(opts *bind.FilterOpts, imageId [][32]byte, asserter []common.Address, assertionId [][32]byte) (*DeepFakeDetectorImageAssertionResolvedIterator, error) {

	var imageIdRule []interface{}
	for _, imageIdItem := range imageId {
		imageIdRule = append(imageIdRule, imageIdItem)
	}

	var asserterRule []interface{}
	for _, asserterItem := range asserter {
		asserterRule = append(asserterRule, asserterItem)
	}
	var assertionIdRule []interface{}
	for _, assertionIdItem := range assertionId {
		assertionIdRule = append(assertionIdRule, assertionIdItem)
	}

	logs, sub, err := _DeepFakeDetector.contract.FilterLogs(opts, "ImageAssertionResolved", imageIdRule, asserterRule, assertionIdRule)
	if err != nil {
		return nil, err
	}
	return &DeepFakeDetectorImageAssertionResolvedIterator{contract: _DeepFakeDetector.contract, event: "ImageAssertionResolved", logs: logs, sub: sub}, nil
}

// WatchImageAssertionResolved is a free log subscription operation binding the contract event 0x2cd47d03ca5eed86ed11475382af40bc0f2965dd65bc4ab337d45577ea5dfaf9.
//
// Solidity: event ImageAssertionResolved(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) WatchImageAssertionResolved(opts *bind.WatchOpts, sink chan<- *DeepFakeDetectorImageAssertionResolved, imageId [][32]byte, asserter []common.Address, assertionId [][32]byte) (event.Subscription, error) {

	var imageIdRule []interface{}
	for _, imageIdItem := range imageId {
		imageIdRule = append(imageIdRule, imageIdItem)
	}

	var asserterRule []interface{}
	for _, asserterItem := range asserter {
		asserterRule = append(asserterRule, asserterItem)
	}
	var assertionIdRule []interface{}
	for _, assertionIdItem := range assertionId {
		assertionIdRule = append(assertionIdRule, assertionIdItem)
	}

	logs, sub, err := _DeepFakeDetector.contract.WatchLogs(opts, "ImageAssertionResolved", imageIdRule, asserterRule, assertionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeepFakeDetectorImageAssertionResolved)
				if err := _DeepFakeDetector.contract.UnpackLog(event, "ImageAssertionResolved", log); err != nil {
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

// ParseImageAssertionResolved is a log parse operation binding the contract event 0x2cd47d03ca5eed86ed11475382af40bc0f2965dd65bc4ab337d45577ea5dfaf9.
//
// Solidity: event ImageAssertionResolved(bytes32 indexed imageId, string imageUrl, address indexed asserter, bytes32 indexed assertionId)
func (_DeepFakeDetector *DeepFakeDetectorFilterer) ParseImageAssertionResolved(log types.Log) (*DeepFakeDetectorImageAssertionResolved, error) {
	event := new(DeepFakeDetectorImageAssertionResolved)
	if err := _DeepFakeDetector.contract.UnpackLog(event, "ImageAssertionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
