package UBTDapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)


func (contract *Contract)putImage(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	var image Images

/*
type Images struct {

	ObjectType string `json:"object_type"`
	ContractID string `json:"contract_id"`
	ImageName string `json:"image_name"`
	TimeStamp string    `json:"time_stamp"`
	ImageHash string  `json:"image_hash"`
	ImageUser string   `json:"image_user"`
	ImageSender string `json:"image_sender"`
	ImageAccepter string `json:"image_accepter"`

	ImageTopic string `json:"image_topic"`
	ImageLevel string `json:"image_level"`
	ImagekeyWords string `json:"image_key_words"`
	ImagePhotoDate string `json:"image_photo_date"`
	ImageHeight string `json:"image_height"`
	ImageWidth string `json:"image_width"`
	ImageSize string `json:"image_size"`


	ImageSourceSet string `json:"image_source_set"`
	ImageUploadSet string `json:"image_upload_set"`

	ImageDealPrice string `json:"image_deal_price"`  //token
	ImageThumbnail string `json:"image_thumbnail"`   //0.4 alpha 10%size


}
 */
	image.ObjectType = "imagedemo"
	image.ContractID = IDRandomCreate()
	image.TimeStamp = strconv.FormatInt(time.Now().UnixNano(),10)
	image.ImageUser = args[0]
	image.ImageSender = args[1]
	image.ImageAccepter = args[2]
	image.ImageName = args[3]
	image.ImageHash = args[4]
	image.ImageTopic = args[5]
	image.ImageLevel = args[6]
	image.ImagekeyWords = args[7]
	image.ImagePhotoDate = args[8]
	image.ImageHeight = args[9]
	image.ImageWidth = args[10]
	image.ImageSize = args[11]
	image.ImageSourceSet = args[12]
	image.ImageUploadSet = args[13]
	image.ImageDealPrice = args[14]
	image.ImageThumbnail = args[15]


	imageJsonbytes,err := json.Marshal(&image)
	if err != nil {
		return shim.Error("json bytes error")
	}
	err = stub.PutState(image.ContractID,imageJsonbytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("put image data success!"))
}

func (contract *Contract)queryMyContract(stub shim.ChaincodeStubInterface, args []string) peer.Response  {

	if len(args) != 1 {
		return shim.Error("parameter must be username")
	}
	querystr := fmt.Sprintf("{\"selector\":{\"image_user\":\"%s\"}}",args[0])
	resultsIterator, err := stub.GetQueryResult(querystr)
	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResult.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"data\":")
		buffer.WriteString(string(queryResult.Value))
		buffer.WriteString("}")

		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")

	fmt.Printf("- getContractUser queryResult:\n%s\n",buffer.String())
	return shim.Success(buffer.Bytes())

}

func (contract *Contract)queryDetailTx(stub shim.ChaincodeStubInterface, args []string) peer.Response  {

	if len(args) != 1 {
		return shim.Error("parameter must be contractID")
	}

	contractID := args[0]
	con,err := stub.GetState(contractID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + contractID + "\"}"
		return shim.Error(jsonResp)
	}
	if con == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + contractID + "\"}"
		return shim.Error(jsonResp)
	}
	txID := stub.GetTxID()
	jsonResp := fmt.Sprintf("{\"contractID\":\"%s\",\"txID\":\"%s\",\"data\":%s}",contractID,txID,string(con))
	fmt.Printf("Query Response:%s\n",jsonResp)
	return shim.Success([]byte(jsonResp))

}