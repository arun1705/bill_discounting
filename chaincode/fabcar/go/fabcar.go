/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"bytes"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type invoice struct {
	InvoiceID   string    `json:"invoiceID"`
	Quantity    int    `json:"quantity"` //the fieldtags are needed to keep case from bouncing around
	Goods       string `json:"goods"`
	DueDate     string `json:"duedate"`
	InvoiceDate string `json:"invoicedate"`
	To          string `json:"to"`
	UnitPrice   int    `json:"unitprice"`
	TotalPrice  int    `json:"totalprice"`
	Discount   int    `json:"discount"`
	DiscountPrice int `json:"discountPrice"`
	Status      string `json:"status"`
}


// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {

	case "generateInvoice":
		//create a new purchaseOrder
		return t.generateInvoice(stub, args)
	case "readInvoiceDetails":
		//read a invoice details
		return t.readInvoiceDetails(stub, args)
	case "financeInvoice":
		//release finance details
		return t.financeInvoice(stub, args)
	case "getInvoiceByRange":
		//read a released financed details details
		return t.getInvoiceByRange(stub, args)
	case "getHistoryForInvoice":
		//read a released financed details details
		return t.getHistoryForInvoice(stub, args)

	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

// ============================================================
// generateInvoice - generate a new invoice, store into chaincode state
// ============================================================
func (t *SimpleChaincode) generateInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
   
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// ==== Input sanitation ====
	fmt.Println("- start generate invoice")
	if len(args[0]) == 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) == 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) == 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) == 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) == 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) == 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) == 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	if len(args[7]) == 0 {
		return shim.Error("8th argument must be a non-empty string")
	}
	if len(args[8]) == 0 {
		return shim.Error("9th argument must be a non-empty string")
	}


	invoiceID := args[0]

	goods := args[1]
	quantity, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	duedate := args[3]
	invoicedate := args[4]
	to := args[5]

	unitprice, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7th argument must be a numeric string")
	}
	totalprice, err := strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("8th argument must be a numeric string")
	}
	status := args[8]
	discount := 0
	discountprice := 0

	// ==== Create Invoice object and marshal to JSON ====

	invoice := &invoice{invoiceID, quantity, goods, duedate, invoicedate, to, unitprice, totalprice,discount,discountprice, status}
	invoiceJSONasBytes, err := json.Marshal(invoice)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Invoice to state ===
	err = stub.PutState(to , invoiceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== invoice saved. Return success ====
	fmt.Println("- end generate Invoice")
	return shim.Success(nil)
}

// ============================================================
// financeInvoice - settle a invoice with some discount on bill, store into chaincode state
// ============================================================
func (t *SimpleChaincode) financeInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) !=11  {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	// ==== Input sanitation ====
	fmt.Println("- start generate invoice")
	if len(args[0]) == 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) == 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) == 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) == 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) == 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) == 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) == 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	if len(args[7]) == 0 {
		return shim.Error("8th argument must be a non-empty string")
	}
	if len(args[8]) == 0 {
		return shim.Error("9th argument must be a non-empty string")
	}
	if len(args[9]) == 0 {
		return shim.Error("10th argument must be a non-empty string")
	}
	if len(args[10]) == 0 {
		return shim.Error("11th argument must be a non-empty string")
	}

	invoiceID := args[0]
	
	goods := args[1]
	quantity, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	duedate := args[3]
	invoicedate := args[4]
	to := args[5]

	unitprice, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7th argument must be a numeric string")
	}
	discount, err := strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("8th argument must be a numeric string")
	}
	totalprice, err := strconv.Atoi(args[8])
	if err != nil {
		return shim.Error("9th argument must be a numeric string")
	}
	discountprice, err := strconv.Atoi(args[9])
	if err != nil {
		return shim.Error("10th argument must be a numeric string")
	}
	status := args[10] 

	// ==== Save Invoice private details ====

	invoiceDetails := &invoice{invoiceID,quantity, goods,duedate,invoicedate,to,unitprice,totalprice, discount, discountprice,status}
	invoiceDetailsBytes, err := json.Marshal(invoiceDetails)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(to, invoiceDetailsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== invoice saved. Return success ====
	fmt.Println("- end finance invoice")
	return shim.Success(nil)
}

// ===============================================
// readInvoiveDetails - read a marble private details from chaincode state
// ===============================================
func (t *SimpleChaincode) readInvoiceDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var to string
	var jsonResp string
	var err error
	to = args[0]
	valAsbytes, err := stub.GetState(to) //get the marble private details from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get private details for " + to + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Ivoice details does not exist: " + to + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}




func (t *SimpleChaincode) getInvoiceByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getInvoiceByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


func (t *SimpleChaincode) getHistoryForInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	to := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", to)

	resultsIterator, err := stub.GetHistoryForKey(to)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForInvoice returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
