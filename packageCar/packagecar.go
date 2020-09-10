package packageCar

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// *********************************************************
type PackageCar struct {
	Cars []Car
	Status bool
}

type Car struct {
        Key    string
        Make   string `json:"make"`
        Model  string `json:"model"`
        Colour string `json:"colour"`
        Owner  string `json:"owner"`
        Status bool   `json:"status"`
}

func (s *PackageCar) InitPackage(){
        cars := []Car{
                Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko",Status: false},
                Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad",Status: false},
                Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo",Status: false},
                Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max",Status: false},
                Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana",Status: false},
                Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel",Status: false},
                Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav",Status: false},
                Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari",Status: false},
                Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria",Status: false},
                Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro",Status: false},
        }
        
        s.Status = false
        i := len(s.Cars)

        for _, car := range cars {
                s.Cars = append(s.Cars,car)
                s.Cars[i].Key = "CAR"+strconv.Itoa(i)
                i++
        }
}

func (s *PackageCar) CreateCar(carNumber string, make string, model string, colour string, owner string) {
        car := Car{
	        	Key: carNumber,
                Make:   make,
                Model:  model,
                Colour: colour,
                Owner:  owner,
                Status: false,
        }
        s.Cars = append(s.Cars,car)       
}
func (s *PackageCar) AddCar(car Car) {
        s.Cars = append(s.Cars,car)       
}
// QueryCar returns the car stored in the world state with given id
func (s *PackageCar) QueryCar(id string) *Car {
	    
	    for i,car := range s.Cars{ // địa chỉ của car khác địa chỉ s.Cars[i]
	    	if (strings.Contains(car.Key,id) && utf8.RuneCountInString(car.Key) == utf8.RuneCountInString(id)) {
	    		return &s.Cars[i]
	    	}
	    }
	    fmt.Println("non car")
	    return &Car{}
}

func (s *PackageCar) QueryAllCars() ([]Car) {
        return s.Cars
}

func (s *PackageCar) ChangeCarOwner(carNumber string, newOwner string) {
        car := s.QueryCar(carNumber)

        car.Owner = newOwner // thay đổi chủ sở hữu
}

func (s *PackageCar) Transport(){
	s.Status = true
	for i,_ := range s.Cars{
	    	s.Cars[i].Status = true
	    } 
}

func (s *PackageCar) DoneHandOver(name string){
	s.Status = false
	s.ChangePackageOwner(name)
	
	for i,_ := range s.Cars{
	    	s.Cars[i].Status = false
	    }
}

func (s *PackageCar) ChangePackageOwner(name string){
	for i := range s.Cars{
		s.Cars[i].Owner = name
	}
}

//  ---------  printf -------
func (car *PackageCar)Show(id string){
	for _,car := range car.Cars{
	    	if strings.Contains(car.Key,id) && utf8.RuneCountInString(car.Key) == utf8.RuneCountInString(id) {
	    		fmt.Println(car)
	    		return
	    	}
	}
}
func (s *PackageCar)ShowAll(){
	for _,car := range s.Cars{
	  	fmt.Println(car)
	  	fmt.Println()
	}
}

//func main(){
//	fmt.Println("   *************************************************  tao package car  ******************** \n")
//	fmt.Println("-----------------  init package car --------------------- \n")
//	
//	pkg := PackageCar{}
//	pkg.ShowAll()
//	
//	fmt.Println("-----------------  test AddCar --------------------------- \n")
//	car := &Car{Key: "MyID",Make: "No-Make", Model: "No-Model", Colour: "No-Color", Owner: "Ngo-Quang-Hieu"};
//	pkg.AddCar(*car)
//	pkg.ShowAll()
//	
//	fmt.Println("-----------------  test func InitPackage ----------------- \n")
//	pkg.InitPackage()
//	pkg.ShowAll()
//	
//	fmt.Println("-----------------  test func CreateCar ------------------- \n")
//	pkg.CreateCar("A","B","C","D","E")
//	pkg.ShowAll()
//	
//	fmt.Println("-----------------  test func QueryCar -------------------- \n")
//	car2 := pkg.QueryCar("MyID")
//	fmt.Println(*car2)
//	
//	fmt.Println("-----------------  test func QueryAllCar ----------------- \n")
//	allcar := pkg.QueryAllCars()
//	fmt.Println(allcar)
//	
//	fmt.Println("-----------------  test func ChangeCarOwner -------------- \n")
//	pkg.ChangeCarOwner("MyID","Do-Duy-Hung")
//	car3 := pkg.QueryCar("MyID")
//	fmt.Println(*car3)
//	
//	fmt.Println("\n\n   *************************************************  van chuyen package car  ******************** \n")
//	pkg.Transport()
//	pkg.ShowAll()
//	fmt.Println(pkg.Status)
//	
//	fmt.Println("\n\n   *************************************************  ban giao package car  ******************** \n")
//	pkg.DoneHandOver("Nguyen Duc Thang")
//	pkg.ShowAll()
//	fmt.Printf("Status = %v",pkg.Status)
//	
//}