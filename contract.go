package UBTDapp

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Contract struct {

}

func (contract *Contract)Invoke(stub shim.ChaincodeStubInterface) peer.Response  {

	re, args := stub.GetFunctionAndParameters()

	switch re {

	case "putimage":
		return  contract.putImage(stub,args)

	case "querycontract":
		return contract.queryMyContract(stub,args)
	case "querydetailTx":
		return contract.queryDetailTx(stub,args)

	default:
		return shim.Error("Unknown func type while invoke, please check!")
	}

	return shim.Error("Unknown func type while invoke, please check!")
}

func (contract *Contract)Init(stub shim.ChaincodeStubInterface) peer.Response  {

	args := stub.GetStringArgs()
	if len(args) != 0 {
		return shim.Error("init parameter error!")
	}


	return shim.Success(nil)
}