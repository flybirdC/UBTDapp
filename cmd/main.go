package main

import (
	"fmt"
	"github.com/UBTDapp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(UBTDapp.Contract))
	if err != nil {
		fmt.Errorf("contract chaincode start error = %s",err)
	}
}

