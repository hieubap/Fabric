package main

import (
"fmt"
"encoding/json"
"strconv"	
"github.com/hyperledger/fabric-contract-api-go/contractapi"
"packageCar"
"os"
//"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// *********************************************************
type SmartContract struct {
        contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
		Key    string
        Make   string `json:"make"`
        Model  string `json:"model"`
        Colour string `json:"colour"`
        Owner  string `json:"owner"`
        Status bool   `json:"status"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
        Key    string `json:"Key"`
        Record *Car
}

func (s *SmartContract) AddCarOfPackage(ctx contractapi.TransactionContextInterface,p packageCar.PackageCar) error {
	for _, car := range p.Cars {
                carAsBytes, _ := json.Marshal(car)
                err := ctx.GetStub().PutState(car.Key, carAsBytes)

                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }
	return nil
}

func (s *SmartContract) Transport(ctx contractapi.TransactionContextInterface,p packageCar.PackageCar){
	p.Transport()
}

func (s *SmartContract) HandOver(ctx contractapi.TransactionContextInterface,p packageCar.PackageCar,own string) error{
	p.ChangePackageOwner(own)
	
	for _, car := range p.Cars {
                carAsBytes, _ := json.Marshal(car)
                err := ctx.GetStub().PutState(car.Key, carAsBytes)

                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }
	
	return nil
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
//	p := packageCar.PackageCar{}
//	p.InitPackage()
//	p.ShowAll()
	
//	car := Car{
//                Make:   "make",
//                Model:  "model",
//                Colour: "colour",
//                Owner:  "owner",
//        }
//
//    carAsBytes, _ := json.Marshal(car)
//        
//	transactioncontext := new(contractapi.TransactionContext)
//	
//	if carAsBytes == nil{
//		fmt.Println("error")
//	}
//	if transactioncontext.GetStub() == nil{
//		fmt.Println("error")
//	}
//	
//	transactioncontext.GetStub().PutState("STATE",carAsBytes)
//	
//	carAsBytes,_ = transactioncontext.GetStub().GetState("STATE")
//	car1 := new(Car)
//	
//	json.Unmarshal(carAsBytes,car1)
//	car2 := *car1
//	fmt.Println(car2)
	
//	s := SmartContract{}
//	s.InitLedger(transactioncontext)
	
	
//	chaincode, err := contractapi.NewChaincode(new(SmartContract))
//
//        if err != nil {
//                fmt.Printf("Error create fabcar chaincode: %s", err.Error())
//                return
//        }
//
//        if err := chaincode.Start(); err != nil {
//                fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
//        }

//	os.Setenv("CORE_CHAINCODE_ID_NAME","E:/WorkspaceEclipse/Fabcar/src/Fabcar")
//	if os.Getenv("CORE_CHAINCODE_ID_NAME") == ""{
//		fmt.Println("null")
//	}
	fmt.Println(os.Getenv("CORE_CHAINCODE_ID_NAME"))
	
	
}