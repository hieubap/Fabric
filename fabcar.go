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
        SizePackage int
}

// Car describes basic details of what makes up a car
type Car struct {
        Make   string `json:"make"`
        Model  string `json:"model"`
        Colour string `json:"colour"`
        Owner  string `json:"owner"`
        Status string `json:"status"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
        Key    string `json:"Key"`
        Record *Car
}
type PackageCar struct {
	Cars []Car
	Status string `json:"status"`
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error { return nil}

func (s *SmartContract) CreatePackageCar(ctx contractapi.TransactionContextInterface) (PackageCar,error){
	p := PackageCar{}
	cars := []Car{
               Car{Make: "Hanoi", Model: "A", Colour: "blue", Owner: "NoOwner",Status:"NoStatus"},
               Car{Make: "Hanoi", Model: "B", Colour: "red", Owner: "NoOwer",Status:"NoStatus"},
               Car{Make: "Hanoi", Model: "C", Colour: "green", Owner: "NoOwer",Status:"NoStatus"},
               Car{Make: "Hanoi", Model: "D", Colour: "yellow", Owner: "NoOwer",Status:"NoStatus"},
               Car{Make: "Hanoi", Model: "E", Colour: "black", Owner: "NoOwer",Status:"NoStatus"},
       }
	p.Cars = cars
	p.Status = "NoStatus"
	
	carAsBytes, _ := json.Marshal(p)
    _ = ctx.GetStub().PutState("PACKAGE"+strconv.Itoa(s.SizePackage), carAsBytes)
    s.SizePackage = s.SizePackage+1
    
    return p,nil
}

func (s *SmartContract) Transport(ctx contractapi.TransactionContextInterface,idPkg string) (*PackageCar,error){
	p := new(PackageCar)
	pkgAsbytes,err := ctx.GetStub().GetState(idPkg)
	if err != nil {
         return p,fmt.Errorf("Failed to read from world state. %s", err.Error())
    }

    if pkgAsbytes == nil {
         return p,fmt.Errorf("%s does not exist", idPkg)
    }
	_ = json.Unmarshal(pkgAsbytes,p)
	p.Status = "Transport"
	
	for i := range p.Cars{
		p.Cars[i].Status = "Transport"
	}
	pkgAsbytes, _ = json.Marshal(p)
	ctx.GetStub().PutState(idPkg,pkgAsbytes)
		
	return p,nil
}

func (s *SmartContract) HandOver(ctx contractapi.TransactionContextInterface,id string,own string) (*PackageCar,error){
	p := new(PackageCar)
	pkgAsbytes,err := ctx.GetStub().GetState(id)
	if err != nil {
         return p,fmt.Errorf("Failed to read from world state. %s", err.Error())
    }

    if pkgAsbytes == nil {
         return p,fmt.Errorf("%s does not exist", id)
    }
	_ = json.Unmarshal(pkgAsbytes,p)
	p.Status = "HandedOver"
	
	for i := range p.Cars{
		p.Cars[i].Status = "HandedOver"
		p.Cars[i].Owner = own
	}
	
	pkgAsbytes, _ = json.Marshal(p)
	ctx.GetStub().PutState(id,pkgAsbytes)
	
	return p,nil
}
func (s *SmartContract) QueryPkg(ctx contractapi.TransactionContextInterface,id string) (*PackageCar,error){
	p := new(PackageCar)
	pkgAsbytes,err := ctx.GetStub().GetState(id)
	if err != nil {
         return p,fmt.Errorf("Failed to read from world state. %s", err.Error())
    }

    if pkgAsbytes == nil {
         return p,fmt.Errorf("%s does not exist", id)
    }
	_ = json.Unmarshal(pkgAsbytes,p)
	
	return p,nil
}
func (s *SmartContract) TestFunction(ctx contractapi.TransactionContextInterface) (string, error){
	fmt.Println("in ra dong ket noi thanh cong")
	return "test thanh cong",nil
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
        cars := []Car{
                Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
                Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
                Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
                Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
                Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
                Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
                Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
                Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
                Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
                Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
        }
		
        for i, car := range cars {
                carAsBytes, _ := json.Marshal(car)
                err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }

        return nil
}

func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
        car := Car{
                Make:   make,
                Model:  model,
                Colour: colour,
                Owner:  owner,
        }

        carAsBytes, _ := json.Marshal(car)

        return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
        carAsBytes, err := ctx.GetStub().GetState(carNumber)

        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }

        if carAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", carNumber)
        }

        car := new(Car)
        _ = json.Unmarshal(carAsBytes, car)
        
        return car, nil
}

func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
        startKey := "CAR0"
        endKey := "CAR99"

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

                car := new(Car)
                _ = json.Unmarshal(queryResponse.Value, car)

                queryResult := QueryResult{Key: queryResponse.Key, Record: car}
                results = append(results, queryResult)
        }

        return results, nil
}

func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
        car, err := s.QueryCar(ctx, carNumber)

        if err != nil {
                return err
        }

        car.Owner = newOwner // thay đổi chủ sở hữu

        carAsBytes, _ := json.Marshal(car) // chuyển dữ liệu sang dạng byte

        return ctx.GetStub().PutState(carNumber, carAsBytes) // cập nhật lại
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