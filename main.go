package main

import (
	"first-golang-demo/helper"
	"fmt"
	"os"
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
		cart []string
	}

	// Embeded  or Has-A or Composition relationship

	// Processor struct
	type Processor struct{
		p_name string
		p_price int
	}

	// Memory struct
	type Memory struct{
		m_name string
		m_price int
	}

	//Computer struct
	type Computer struct{
		Processor
		Memory
		brand string
	}

	// pointer with struct

	type Foo struct{
		name string	
	}

	type Rectangle struct{
		width, height int
	}

	// Creating wait group for synchronization
	var wg = sync.WaitGroup{}

	//========================================== MAIN FUNCTION =============================================
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
		//or arrayA := [...]string{"Gandia", "Arturo", "Alicia"}

		fmt.Printf("Whole array %v \n", arrayA)

		arrayB := []int{1,2,3,4,5,6}
		fmt.Printf("Whole int array %v \n", arrayB)
		//========= Array copying ==========
		/* 
		Copying the whole array. Changes to copy array will not affect the original
		Begin Index - inclusive
		End Index - Exclusive
		*/
		arrayCopy := arrayB
		fmt.Printf("Copy array %v \n", arrayCopy)

		a := arrayB[:]
		fmt.Printf("Array A %v \n", a)
        // copy from 2 nd index to end 
		b := arrayB[2:]
		fmt.Printf("Array B %v \n", b)
		// copy till 4 th index 
		c := arrayB[:4]
		fmt.Printf("Array C %v \n", c)
		// copy from 2nd to 5th index 
		d := arrayB[2:5]
		fmt.Printf("Array D %v \n", d)

		// Assigning the reference only. Changes to reference will affect the original array
		arrayRef := &arrayB
		fmt.Printf("Reference array %v \n", arrayRef)

		// Multi Dimensional Array
		var s [2][2]int
		s[0][0] = 1
		s[0][1] = 2
		fmt.Printf("2D array %v \n", s)

        var identityMatrix = [3][3] int {
			{1,0,0},
			{1,0,1},
			{0,0,1},
		}
		fmt.Printf("multi dimensional  array %v \n", identityMatrix)
         
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
	
		/*
			Slice internally use an array i.e it hold reference of an array 
			So changing copy of a slice actually changes the original slice as well
			when we copy a slice it actually copy the pointer location as well
			All array operation is applicable for slice also
			Size and capacity can be different for slice
			Duplicate allowed in slice
			append only make changes to copy slice not the original
		*/
		var ab =[]int{1,2,3}
		fmt.Printf("Slice AB %v\n", ab)
		bc := ab
		fmt.Printf("Slice BC %v\n", bc)
		bc[0]=6
		fmt.Printf("Slice AB after update to index 0%v\n", ab)
		fmt.Printf("Slice BC after update to index 0%v\n", bc)
		fmt.Printf("Capacity of slice ab %v\n", cap(ab))

		// Slice using make
		cd := make([]int,1,3)
		fmt.Printf("Capacity of slice cd %v\n", cap(cd))
		fmt.Printf("Size of slice cd %v\n", len(cd))
		cd = append(cd, 1)
		cd = append(cd, 2)
		cd = append(cd, 2)
		cd = append(cd, 2)
		fmt.Printf("Content of slice cd %v\n", cd)
		// if we declare size of slice other than 0 then it will prefill the slice with default value based on slice type

		x := []int{1,2,3}
		fmt.Printf("Content of slice x %v\n", x)
		//append only make changes to copy slice not the original
		y:=x
		y= append(x,4)	
		fmt.Printf("Content of slice y %v\n", y)
		fmt.Printf("Content of slice x %v\n", x)

		y = append(x[1:3],4)	
		fmt.Printf("Content of slice y %v\n", y)
		// spread and append
		z := append(x,y...)
		fmt.Printf("Content of slice z %v\n", z)



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

		for i,j := 0, 0; i <5; i, j =i+1, j+1{
			fmt.Println(i,j)
		}
		//OR
		i,j := 0, 0;
		for i<5 {
			fmt.Println(i,j)
			i, j = i+1, j+1
		}

		myMaps := make(map[string]string,0)
		myMaps["name"] ="Rajib"
		myMaps["age"] = "11"

		if _, ok := myMaps["email"]; ok{
             fmt.Println("Name exits in the map")
		}else{
			fmt.Println("Name does not exits in the map")
		}

		//OR
		
		_, ok := myMaps["email"]
		if ok {
			fmt.Println("Name exits in the map")
		}else{
				fmt.Println("Name does not exits in the map")
		}
        // Iterate over map
		for k,v := range myMaps{
			fmt.Println(k,v)
		}
		// Iterate over slice
		sliceIt := make([]int,0)
		sliceIt = append(sliceIt, 10)
		sliceIt = append(sliceIt, 20)

		for k,v := range sliceIt{
			fmt.Println(k,v)
		}
		// Iterate over array
		arrayIt := [6]int{1,2,3,4,5}
		for k,v := range arrayIt{
			fmt.Println(k,v)
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

		// switch with Interface

		var op interface{} = 8
		switch op.(type) {
		case int:
			fmt.Println("Int type")
		case string:
			fmt.Println("String type")
		default:
			fmt.Println("No choice found")
		}
		// switch with fallthrough
		//fallthrough will execute the next statement also

		switch 1 {
		case 1:
			fmt.Println("1")
			fallthrough
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")	
		default:
			fmt.Println("No choice found")
		}
		
		//================== MAP ================================= 
		/*
			Map :unique set of key and values
			Not an ordered data structure
		*/
		
		var myMap = make(map[string]string)
		myMap["firstname"] = "Rajib"
		myMap["lastName"] = "Rath"
		//deletion by key
		delete(myMap, "firstname")
		fmt.Printf("Type of myMap %T \n", myMap)
		fmt.Printf("Length of myMap %v \n", len(myMap))
		fmt.Printf("FirstName %v \n", myMap["firstname"])
		

		//OR
		myMap1 := map[string]string {
			"firstName": "Rajib",
			"LastName": "Rath",
		}
		fmt.Printf("Data of myMap1 %v \n", myMap1)
		fmt.Printf("Type of myMap1 %T \n", myMap1)

		//Empty map
		var emptyMap = map[int]int{}
		fmt.Printf("Length of emptyMap %v \n", len(emptyMap))

		//Array is valid map key type but not the slice
		var testMap = map[[2]string]string{}
		fmt.Printf("Data of testMap %v \n", testMap)

		m:=map[[1]int]string{}
		fmt.Printf("Data of m %v \n", m)

		// ,ok sysntax to check availability of key in the map
		commaOkMapTest := make(map[string]int)
		commaOkMapTest["mango"] = 1
		commaOkMapTest["apple"] = 2
		// will return 0 even if key absent so to avoid this use ,ok
        fmt.Printf("When key absent%v\n",commaOkMapTest["nuts"]) 

		_, oks := commaOkMapTest["nuts"]
        fmt.Printf("When key absent with comma ok -->%v\n", oks) 

		// Map copy
		/*
		When copying new map will have a pointer pointing to old One
		So making any changes to copy map also reflects in old map 
		*/
        // copy map
		testMapOne := make(map[string]int)
		testMapOne["mango"] = 1
		testMapOne["apple"] = 2
		copyMapByValue :=commaOkMapTest
        fmt.Printf("copyMapByValue content%v\n",copyMapByValue) 
		fmt.Printf("testMapOne content before any change to copyMap%v\n",commaOkMapTest) 
		copyMapByValue["apple"] = 8
		fmt.Printf("copyMapByValue content after update %v\n",copyMapByValue) 
		fmt.Printf("testMapOne content after changes to copyMap%v\n",commaOkMapTest) 

		//List or slice  of map
		var listOfMap =make([]map[string]string,0)
		fmt.Printf("Type of myMap1 %T \n", listOfMap)
		listOfMap = append(listOfMap, myMap)

		for _, mapData := range listOfMap{
			fmt.Println("Iterating over list of map")
			fmt.Printf("Data of map %v \n", mapData)
			fmt.Printf("FirstName %v \n", mapData["firstname"])
		}

		// ########### Struct operation #############
         /*
		  Struct is value type not based on reference 
		  i.e. changing value of copy struct does not affect original struct
		  copy by reference ensure that change into copy also reflects in original
		 */

		// slice of struct
		var listOfStruct =make([]UserData,0)
		fmt.Printf("Type of listOfStruct %T \n", listOfStruct)
		fmt.Printf("Content of listOfStruct %v \n", listOfStruct)

		// Initializing struct
		var userdata = UserData{
			firstName: "Rajib",
			lastName: "Rath",
			age: 25,
			isActive: true,
		}
		//updating struct data
		userdata.age = 26

		//copy struct by value
		userDataCopy := userdata
		fmt.Printf("Type of userDataCopy %T \n", userDataCopy)


		fmt.Printf("Type of userdata %T \n", userdata)
		fmt.Printf("FirstName from userData %v \n", userdata.firstName)
		var userdata1 = UserData{
			firstName: "Hello",
			lastName: "Worls",
			age: 15,
			isActive: false,
			cart: []string{
				"Apple",
				"Mango",
			},
		}
		listOfStruct = append(listOfStruct, userdata)
		listOfStruct = append(listOfStruct, userdata1)
		fmt.Printf("Size of listOfStruct %v \n", len(listOfStruct))
		
		for _, structData := range listOfStruct{
			fmt.Printf("Type of structData %T \n", structData)
			fmt.Printf("FirstName from structData %v \n", structData.firstName)
			fmt.Printf("cart from structData %v \n", structData.cart)

		}

		// Copy of struct by reference
		//change to reference reflects in original
		copyStructByReference := & userdata
		fmt.Printf("Content of copyStructByReference %v \n", copyStructByReference)
		fmt.Printf("Content of userdata before update to copyStructByReference %v \n", userdata)
        copyStructByReference.age = 0
		fmt.Printf("Content of copyStructByReference after update %v \n", copyStructByReference)
		fmt.Printf("Content of userdata after update to copyStructByReference %v \n", userdata)

		//Embeding in golang

        computerStruct := Computer{}
		computerStruct.brand = "DELL"
		computerStruct.m_name = "TOSHIBA"
		computerStruct.p_name = "INTEL"
		computerStruct.m_price= 100
		computerStruct.p_price= 1000

		fmt.Printf("Content of ComputerStruct %v \n", computerStruct)

		//OR

		computerStructOne := Computer{

			Processor: Processor{
				p_name: "AMD",
				p_price: 1000,
			},
			Memory: Memory{
				m_name: "SAMSUNG",
				m_price: 11111,
			},
		}
		fmt.Printf("Content of computerStructOne %v \n", computerStructOne)

		//====================== Defer, Panic and Recover =========================================
		/*
		 defer - means we are defering from execution of a statement at that moment.
		 But the statement will get executed just before returning the value
		 defer act in LIFO order
		*/
		// op will be 1 3 2
		fmt.Println(1)   
		defer fmt.Println(2)   
		fmt.Println(3) 
		// op will be 3 2 1
		defer fmt.Println(1)   
		defer fmt.Println(2)   
		defer fmt.Println(3) 
		// panic will throw the error if any
		myFile := createFile("myFile.txt")
		fmt.Printf("Type of file %T", myFile)

		//================== Pointers ======================

		var obx int = 1
		var obx_ptr *int = &obx
		fmt.Printf("Value of obx %v\n", obx)
		fmt.Printf("Memory location of obx_ptr %v\n", obx_ptr)
		fmt.Printf("Value of obx_ptr %v\n", *obx_ptr)
		//OR
		var obj = 1
		var obj_ptr = &obj
		fmt.Printf("Value of obj %v\n", obj)
		fmt.Printf("Memory location of obj_ptr %v\n", obj_ptr)
		fmt.Printf("Value of obj_ptr %v\n", *obj_ptr)
		// pointer with struct
		// foo_obj := Foo{}
		// foo_obj.name ="Hello Pointer"
		// fmt.Println(foo_obj) //{Hello Pointer}

		var foo_ptr *Foo
		fmt.Println(foo_ptr) // print default value nil
		foo_ptr = new(Foo)
		fmt.Println(&foo_ptr) //0xc00000e040
		fmt.Println(*&(foo_ptr).name) // empty
		foo_ptr.name ="Hi"
		fmt.Println(foo_ptr.name) // Hi

		//================== Go Function ===================================

		// By default go function parameter are pass by value
		msg := "Hi"
		writeMessage(msg)  // Hi
		fmt.Println(msg) // Hi

		writeMessage2(msg)  // Hello
		fmt.Println(msg) // Hi
		// Pass parameter by reference
		fmt.Printf("Before value change to the reference: %v\n",msg)
		writeMessageByReference(&msg)
		fmt.Printf("After value change to the reference: %v\n",msg)

		// Multi param function or variadic function
		multiParam(1,2,3)
		multiParam(2,3)
        // Function with error return
		value, err :=funcWithReturnWIthError()
		if(err != nil){
            fmt.Printf("Error msg is :%v\n",err)
		}
		fmt.Print(value)

		//Anonymous function declaration and invocation
		func(){
			fmt.Println("Inside anonymous function")
		}() // calling function by ()
		
		//Anonymous function declaration and assignment
		annoFunc := func(){
			fmt.Println("Inside anonymous function")
		}
		annoFunc()

		//Function with struct
	
		rect := Rectangle{
			width: 2,
			height: 3,
		}
	    ara := rect.area()
		fmt.Printf("Area of Rectangle is :%v\n", ara)
		//================== Caling functions ==================
		greetuser()
		greetuserByName("Rajib")

		returnValue :=funcWithReturn()
		fmt.Printf("Return value from function is..........%v\n", returnValue)
		
		returnValue1, returnValue2 := funcWithMultipleReturn()
		fmt.Printf("First return value from function is..........%v\n", returnValue1)
		fmt.Printf("Second return value from function is..........%v\n", returnValue2)
	
		// Interface and strut Use
		var circleCalc geometry = Circle{
			radius: 3,
		}
		fmt.Println(circleCalc.area())
		fmt.Println(circleCalc.perimeter())
		// OR
		circleCalc2 := Circle{
			radius: 3,
		}
		fmt.Println(circleCalc2.area())
		fmt.Println(circleCalc2.perimeter())


		//================== Go concurency ========================== 

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

	// ================== function Declaration================================  
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
	// function with return value and error

	func funcWithReturnWIthError() (string, error){
		if(true){
			return "Returning from function", fmt.Errorf("Some error occurred")
		}else{
			return "Returning from function", nil
		}
	}
	// function create file
	func createFile(path string) *os.File{

		fmt.Println("Creating file ")
		file , err := os.Create(path)
		if err != nil {
			panic(err)
		}
		return file
	}
	// pass by value or passing copy 
     func writeMessage(msg string){
		fmt.Println(msg)
		msg = "Hello"
	 }
	 // pass by value or passing copy 
	 func writeMessage2(msg string){
		msg = "Hello"
		fmt.Println(msg)
	 }
	 // pass by reference
	 func writeMessageByReference(msg *string){
		fmt.Println(*msg)
		*msg = "Hello from ptr"
	 }

	// function with multiple param
	func multiParam(values ...int){
		sum := 0
		for _,v := range values{
			sum = sum + v
		}
		fmt.Printf("Sum of input values %v\n", sum)
	}

	//================== INTERFACE ==================
	// Interface declaration 
	type geometry interface{
		area() float64
		perimeter() float64
	}
	
	type Circle struct{
		radius float64
	}

	func (rect Rectangle) area() int{
		return rect.height * rect.width;
    }

	func (c Circle) area() float64{
	   return c.radius * c.radius * 3.141
    }

    func (c Circle) perimeter() float64{
	return 2 * c.radius * 3.141
    }
	