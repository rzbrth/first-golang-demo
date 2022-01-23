package main

import (
	"first-golang-demo/helper"
	"fmt"
	"sync"
)

//Global or package level variable
const conferenceTickets = 50 //constant

//==================STRUCT==================
//struct or structure help us create mixed data types
type UserData struct{
	firstName string
	lastName  string
	age uint
	isActive bool
}

// Creating wait group for synchronization
var wg = sync.WaitGroup{}

func main() {

	fmt.Println("########## Welcome to our", helper.ConferenceName, "booking application ##########")
	fmt.Printf("Hurry up! we have only %v tickets available\n", conferenceTickets) //formated printing
    //helper.ConferenceName is a package variable i.e. of package scope
	//variable decalration with type and separate assignment 
	var userName string 
	var bookedTickets int
	userName = "Rajib"
	bookedTickets =2
	fmt.Printf("%v has booked %v tickets\n", userName, bookedTickets)


	//variable decalration with type and assignment 
	var message string ="Hello ....."
	fmt.Printf("Printing custom message:- %v \n", message)
	
	//Variable declaration and assignment without mentioning type
	userRating :=  5.3
	fmt.Printf("User rating are %v \n", userRating)

	// Printing value and memory location
	fmt.Printf("Value of user name is %v \n", userName)
	fmt.Printf("Memory addrss of user name is %v \n", &userName)


	// ================== Arrays and Slices ===========================================

	//================== ARRAYS ==========================
	//Array declaration and separate assignment
	var arrayElement [10] string
	arrayElement[0] = "Hi"
    fmt.Printf("Whole array %v \n", arrayElement)
	fmt.Printf("First element of  array %v \n", arrayElement[0])
	fmt.Printf("Length of the array %v \n", len(arrayElement))
	fmt.Printf("Type of the array %T \n", arrayElement)
    
	//Array declaration and assignment
	var arrayA = []string{"Gandia", "Arturo", "Alicia"}
	//or var arrayA = [4]string{"Gandia", "Arturo", "Alicia"}
	//or arrayA := []string{"Gandia", "Arturo", "Alicia"}

	fmt.Printf("Whole array %v \n", arrayA)

	arrayB := []string{"Gandia", "Arturo", "Alicia"}
	fmt.Printf("Whole array %v \n", arrayB)

	//================== SLICES ===========================
	/*
	 It support adding element at the end of the slide .ie appending
	 Grow the slice if greater capacity is needed and return the
	 updated slice value.
	 It's like a dynamic list where we don't have to keep track of index
	*/

   //Slice declaration and separate assignment
	var sliceElement []string
	sliceElement = append(sliceElement,"Hello")
    fmt.Printf("Whole slice %v \n", sliceElement)
	fmt.Printf("First element of  slice %v \n", sliceElement[0])
	fmt.Printf("Length of the slice %v \n", len(sliceElement))
	fmt.Printf("Type of the slice %T \n", sliceElement)
    //OR
	var sliceC = make([]string,0)
	sliceC=append(sliceC, "aaa")
	sliceC=append(sliceC, "bbbb")

	fmt.Printf("Second element of  sliceC %v \n", sliceC[1])
	fmt.Printf("Type of the sliceC %T \n", sliceC)

	//OR
	sliceElementB:= []string{}
    fmt.Printf("Whole slice B%v \n", sliceElementB)
	sliceElementB = append(sliceElementB, "ooo")
    //Slice declaration and assignment
    var sliceA = []string{"Hi", "Hello"}
    sliceA = append(sliceA, "By")

   //OR
   sliceB := []string{"HI", "HELLO"}
   fmt.Printf("Whole slice B %v \n", sliceB)

   fmt.Printf("Whole slice %v \n", sliceA)
   fmt.Printf("Type of the slice %T \n", sliceA)
   //fmt.Printf("Element at index 1 %v \n", sliceA[3]) //index out of range [3] with length 3
   
   // ================== Loop and If-else ============================
   
   var times = 5;
   for{
	if times == 0 {
		break
	}
	fmt.Printf("Inside for loop %v times \n", times)
	times --	 
   }

  cars :=[]string{"TATA", "Ford", "Tesla", "Mahindra"}

  for index, car := range cars {
	  fmt.Printf("car details at car[%v] %v \n", index, car)
  }
   // Use Blank identifire i.e. _ to ignore the variable we don't want to use

  for _, car := range cars {
	fmt.Printf("car details %v \n", car)
}
count :=4
for i := 0; i < count; i++ {
	fmt.Printf("Counter in for loop --%v\n",i)
}


	//================== Asking for user input and User input validation ================== 
	var pincode string
	fmt.Println("Please enter your pincode")
	fmt.Scan(&pincode)

	if(len(pincode) < 4 || len(pincode) >=5){
       fmt.Printf("Bad input\n")
	}else{
		fmt.Printf("User pincode is %v \n", pincode)
	}

    //================== Switch statement ==================
	var carBrand string
	fmt.Printf("Enter car rand to get available car\n")
	fmt.Scan(&carBrand)

	switch carBrand {
	case "TATA":
			fmt.Printf("TATA car available\n")
	case "FORD":
		fmt.Printf("FORD car available\n")
	case "MAHINDRA", "TESLA":
		fmt.Printf("MAHINDRA or TESLA car available\n")
	default:
			fmt.Print("No car present\n")
	}

	//================== Caling functions ==================
	greetuser()
	greetuserByName("Rajib")

	returnValue :=funcWithReturn()
	fmt.Printf("Return value from function is..........%v\n", returnValue)
	
	returnValue1, returnValue2 := funcWithMultipleReturn()
	fmt.Printf("First return value from function is..........%v\n", returnValue1)
	fmt.Printf("Second return value from function is..........%v\n", returnValue2)
   
	//================== MAP ================================= 
	//Map :unique set of key and values
	// Not an ordered data structure
	var myMap = make(map[string]string)
	myMap["firstname"] = "Rajib"
	myMap["lastName"] = "Rath"
	//deletion by key
	delete(myMap, "firstname")
	fmt.Printf("Type of myMap %T \n", myMap)
	fmt.Printf("FirstName %v \n", myMap["firstname"])
	

	//OR
	myMap1 := map[string]string {
		"firstName": "Rajib",
		"LastName": "Rath",
	}
	fmt.Printf("Data of myMap1 %v \n", myMap1)
	fmt.Printf("Type of myMap1 %T \n", myMap1)

	//Array is valid map key type not the slice
	var testMap = map[[2]string]string{}
	fmt.Printf("Data of testMap %v \n", testMap)

	m:=map[[1]int]string{}
	fmt.Printf("Data of m %v \n", m)

    //List of map
	var listOfMap =make([]map[string]string,0)
	fmt.Printf("Type of myMap1 %T \n", listOfMap)
	listOfMap = append(listOfMap, myMap)

	for _, mapData := range listOfMap{
		fmt.Println("Iterating over list of map")
		fmt.Printf("Data of map %v \n", mapData)
	    fmt.Printf("FirstName %v \n", mapData["firstname"])
	}
    // List of struct
	var listOfStruct =make([]UserData,0)
	fmt.Printf("Type of listOfStruct %T \n", listOfStruct)

	// Initializing struct
    var userdata = UserData{
		firstName: "Rajib",
		lastName: "Rath",
		age: 25,
		isActive: true,
	}

	fmt.Printf("Type of userdata %T \n", userdata)
	fmt.Printf("FirstName from userData %v \n", userdata.firstName)
	var userdata1 = UserData{
		firstName: "Hello",
		lastName: "Worls",
		age: 15,
		isActive: false,
	}
	listOfStruct = append(listOfStruct, userdata)
	listOfStruct = append(listOfStruct, userdata1)
	fmt.Printf("Size of listOfStruct %v \n", len(listOfStruct))
    
	for _, structData := range listOfStruct{
		fmt.Printf("Type of structData %T \n", structData)
	    fmt.Printf("FirstName from structData %v \n", structData.firstName)
	}

	//================== Go concurency ================== 

	//calling function from other file of same package
    greetuserFromOutside()

	//calling function from other file of different package
	/*
	 if go mentioned then main thread exit before this thread is complete , 
	 so we need to make main thread to wait untill all threads are done by
	 using waitgroup of sync package
	 */
	wg.Add(2) 
	//We just need to make sure we are passing around a pointer of waitgroup so we do not get copies
	go helper.GreetuserFromOutside(&wg)  
	go greetuserFromOutsideWithSleep(&wg)

	//wait for wg.Done()
	wg.Wait()

}

// ================== function ================================  
func greetuser(){
	fmt.Print("Hello..........\n")
}
// function with parameter
func greetuserByName(name string){
	fmt.Printf("Hello..........%v\n", name)
}
// function with return value
func funcWithReturn() string{
    return "Returning from function"
}
// function with multiple return value
func funcWithMultipleReturn() (string, string){
    return "Returning from function", "Hello"
}


