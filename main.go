package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"os/exec"
	"encoding/json"
)

var mainList [5]string

var i int



func main(){

	//check if the json file exists. if it doesn't, create it.
	CreateJson(mainList, "list.json")

	
	//Retrieve save data and store it in the array
	mainList = ReadArrayFromJson("list.json")

	for i == 0 {

		//clear screen
		clearScrean()

	
		// creating scanner to read user input
		scanner := bufio.NewScanner(os.Stdin)
		// print list to screen
		fmt.Println("Here is your list for today: ")
		

		fmt.Printf("1. %v \n2. %v \n3. %v \n4. %v \n5. %v \n", mainList[0], mainList[1], mainList[2],mainList[3],  mainList[4])
		

		//ask user what they want to do
		fmt.Println("Would you like to do? (Use commands like a (add), r (remove), or 'q' to quit)")
		scanner.Scan()
		mainInput := scanner.Text()
		

		//do things with input
		if mainInput == "a" {
			
			var n int
			
			//add new chore
			fmt.Println("Input new chore: ")
			scanner.Scan()
			input := scanner.Text()
			

			//put chore in a location.
			for n == 0 {
				fmt.Println("In which spot?")
				scanner.Scan()

				//string is converted to an int to work with array
				arrLocation, _ := strconv.ParseInt(scanner.Text(), 10, 64)

				if arrLocation > 5{
					fmt.Println("Invalid input. Please choose a number between 1 and 5")
				}else {
					mainList[arrLocation - 1] = input
					n = 1
				}
			}
		

		//quiting program
		}else if mainInput == "q" {

			//clear screen
			clearScrean()
			
			//save shit to json
			SaveArrayToJson(mainList)

			break 
		
		//removing chore
		}else if mainInput == "r" {
			var y int		
			for y == 0{
				//removing specific chore
				fmt.Println("Which one would you like to remove?")
				scanner.Scan()
				
				//string is converted into an int to work with array
				input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				
				if input>5 {
					fmt.Println("Invalid input. Please choose a number betweeb 1 and 5")
				}else {
					mainList[input - 1] = ""
					y = 1
				}
			}
		
		}else{
			fmt.Println("Invalid input")
		}
	}
}



func printList() {
	
}

func clearScrean(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}



func SaveArrayToJson(array [5]string){
	list, err := json.Marshal(array)
	if err!=nil {
		panic(err)
	}

	jsonFile, _ := os.Create("list.json")
	defer jsonFile.Close()

	jsonFile.Write(list)

}


func ReadArrayFromJson(file string) [5]string {
	var list [5]string

	fileS, _ := os.Open(file)
	defer fileS.Close()
	
	err := json.NewDecoder(fileS).Decode(&list)
	if err!=nil {
		panic(err)
	}

	return list
}

func CreateJson(array [5]string, file string) {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		list, err := json.Marshal(array)
		if err!=nil {
			panic(err)
		}

		jsonFile, _ := os.Create("list.json")
		defer jsonFile.Close()

		jsonFile.Write(list)
	}else {
	}
}




