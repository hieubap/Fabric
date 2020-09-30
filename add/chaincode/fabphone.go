package main

import (
"fmt"
"encoding/json"
"strconv"	
"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// *********************************************************
type SmartContract struct {
        contractapi.Contract
        SizePhone int
        SizePackage int
}

// Car describes basic details of what makes up a car
/*
factory
store
customer

*/
type Phone struct {
        Name   string `json:"name"`
        Owner  string `json:"owner"`
        Price  int `json:"buy"`
        UserTime   int `json:"sell"`
        Status string `json:"status"`
}
type PackagePhone struct{
	Status string
	IdPhone []string
}
type PricePackage struct {
        Name    string `json:"name"`
        Price   int    `json:"salary"`
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) (error){
	phones := []Phone{
               Phone{Name: "Samsung J1", Owner: "Ngo Quang Hieu", Price: 0 , UserTime: 5 ,Status:"User"},
               Phone{Name: "Samsung J2", Owner: "Do Duy Hung", Price: 0 , UserTime: 20 ,Status:"User"},
               Phone{Name: "Samsung J3", Owner: "Store", Price: 5000 , UserTime: 0 ,Status:"NewProducts"},
               Phone{Name: "Samsung J4", Owner: "Store", Price: 6000 , UserTime: 0 ,Status:"NewProducts"},
               Phone{Name: "Samsung J5", Owner: "Store", Price: 7000 , UserTime: 0 ,Status:"NewProducts"},
               
       }
	start := s.SizePhone
	for i := range phones{
		phoneAsBytes, _ := json.Marshal(phones[i])
        err := ctx.GetStub().PutState("PHONE"+strconv.Itoa(i+start), phoneAsBytes)
		
        if err != nil {
                return fmt.Errorf("Failed to put to world state. %s", err.Error())
        }
        s.SizePhone = s.SizePhone + 1
	}
    
    return nil
}

func (s *SmartContract) CreatePackagePhone(ctx contractapi.TransactionContextInterface,name string,price int) (*PackagePhone,error){
	p := new(PackagePhone)
	phones := []Phone{
               Phone{Name: name , Owner: "factory", Price: 5000 , UserTime: 0 ,Status:"NewProducts"},
               Phone{Name: name , Owner: "factory", Price: 5000 , UserTime: 0 ,Status:"NewProducts"},
               Phone{Name: name , Owner: "factory", Price: 5000 , UserTime: 0 ,Status:"NewProducts"},
    }
	p.Status = "New"
	start := s.SizePhone
	for i := range phones{
		phoneAsBytes, _ := json.Marshal(phones[i])
        err := ctx.GetStub().PutState("PHONE"+strconv.Itoa(i+start), phoneAsBytes)
		p.IdPhone = append(p.IdPhone,"PHONE"+strconv.Itoa(i+start))
        if err != nil {
                return nil,fmt.Errorf("Failed to put to world state. %s", err.Error())
        }
        s.SizePhone = s.SizePhone + 1
	}
	phoneAsBytes, _ := json.Marshal(p)
    _ = ctx.GetStub().PutState("PACKAGE"+strconv.Itoa(s.SizePackage), phoneAsBytes)
    
    private := new(PricePackage)
    private.Name = name
    private.Price = price
    priAsbyte,_ := json.Marshal(private)
    ctx.GetStub().PutPrivateData("importprice","PACKAGE0",priAsbyte)
    
    s.SizePackage = s.SizePackage+1
    
    return p,nil
}
func (s *SmartContract) QueryPackagePhone(ctx contractapi.TransactionContextInterface,id string) (*PackagePhone,error){
	phoneAsBytes, err := ctx.GetStub().GetState(id)

        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }

        if phoneAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", id)
        }

        f := new(PackagePhone)
        _ = json.Unmarshal(phoneAsBytes, f)
        
        return f, nil
}
func (s *SmartContract) AddPhoneToPkg(ctx contractapi.TransactionContextInterface, idpkg string, idphone string) error {
	pkg, _ := ctx.GetStub().GetState(idpkg)
	tempPkg := new(PackagePhone)
	err1 := json.Unmarshal(pkg, tempPkg)
	if err1 != nil {
		return fmt.Errorf("Adding C2P Can not read data ! %s", err1.Error())
	}
	
	tempPkg.IdPhone = append(tempPkg.IdPhone, idphone)
	pkg, _ = json.Marshal(tempPkg)
	return ctx.GetStub().PutState(idpkg, pkg)
}

func (s *SmartContract) QueryPackagePrice(ctx contractapi.TransactionContextInterface, idpkg string) (*PricePackage,error) {
	pkg, _ := ctx.GetStub().GetPrivateData("importprice",idpkg)
	tempPkg := new(PricePackage)
	err1 := json.Unmarshal(pkg, tempPkg)
	
	if err1 != nil {
		return nil,fmt.Errorf("Adding C2P Can not read data ! %s", err1.Error())
	}
	
	return tempPkg,nil
}

func (s *SmartContract) HandOver(ctx contractapi.TransactionContextInterface,id string,own string) (*PackagePhone,error){
	pkgAsbytes,err := ctx.GetStub().GetState(id)
	if err != nil {
         return nil,fmt.Errorf("Failed to read from world state. %s", err.Error())
    }

    if pkgAsbytes == nil {
         return nil,fmt.Errorf("%s does not exist", id)
    }
	
	p := new(PackagePhone)
	_ = json.Unmarshal(pkgAsbytes,p)
	p.Status = "HandedOver"
	
	for i := range p.IdPhone{
		phone,_ := s.QueryPhone(ctx,p.IdPhone[i])
		
		phone.Owner = own
		
		phoneAsBytes, _ := json.Marshal(phone)
        err := ctx.GetStub().PutState(p.IdPhone[i], phoneAsBytes)
				
        if err != nil {
                return nil,fmt.Errorf("Failed to put to world state. %s", err.Error())
        }
	}
	
	pkgAsbytes, _ = json.Marshal(&p)
	ctx.GetStub().PutState(id,pkgAsbytes)
	
	return p,nil
}

func (s *SmartContract) ReSell(ctx contractapi.TransactionContextInterface,id string,price int,time int) (error){
	
	phoneAsbytes,err := ctx.GetStub().GetState(id)
	if err != nil {
         return fmt.Errorf("Failed to read from world state. %s", err.Error())
    }

    if phoneAsbytes == nil {
         return fmt.Errorf("%s does not exist", id)
    }
	
	p := new(Phone)
	_ = json.Unmarshal(phoneAsbytes,p)
	p.Status = "ReSell"
	p.UserTime = p.UserTime + time
	
	phoneAsbytes, _ = json.Marshal(&p)
	ctx.GetStub().PutState(id,phoneAsbytes)
	
	return nil
}

func (s *SmartContract) CreatePhone(ctx contractapi.TransactionContextInterface, id string, name string, owner string, price int,timeuse int,status string) error {
        phone := Phone{
                Name:   name,
                Owner:  owner,
                Price: price,
                UserTime:  timeuse,
                Status: "NewProduct",
        }

        phoneAsBytes, _ := json.Marshal(phone)

        return ctx.GetStub().PutState(id, phoneAsBytes)
}

func (s *SmartContract) Transaction(ctx contractapi.TransactionContextInterface, id string, owner string) error {
        phone,_ := s.QueryPhone(ctx,id)
        
        phone.Owner = owner
        phone.Status = "Using"
        
        phoneAsBytes, _ := json.Marshal(phone)

        return ctx.GetStub().PutState(id, phoneAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryPhone(ctx contractapi.TransactionContextInterface, id string) (*Phone, error) {
        phoneAsBytes, err := ctx.GetStub().GetState(id)

        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }

        if phoneAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", id)
        }

        f := new(Phone)
        _ = json.Unmarshal(phoneAsBytes, f)
        
        return f, nil
}
type QueryResult struct {
        Key    string `json:"Key"`
        Record *Phone
}
func (s *SmartContract) QueryAllPhones(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
        startKey := "PHONE0"
        endKey := "PHONE99"

        resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

        if err != nil {
                return nil, err
        }
        defer resultsIterator.Close()

        results := []QueryResult{}

        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()

                if err != nil {
                        return nil, err
                }

                phone := new(Phone)
                _ = json.Unmarshal(queryResponse.Value, phone)

                queryResult := QueryResult{Key: queryResponse.Key, Record: phone}
                results = append(results, queryResult)
        }

        return results, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

        if err != nil {
                fmt.Printf("Error create fabcar chaincode: %s", err.Error())
                return
        }
		if err := chaincode.Start(); err != nil {
                fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
        }
	
}