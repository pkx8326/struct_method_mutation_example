package main

import "fmt"

type student struct {
	firstname string
	lastname  string
	score     float64
}

func main() {
	var studentdata student
	var nxoptn string
	for {

		c := menu()
		if c == "1" {
			studentdata = getdata()
			nxoptn = nextoptn()
		} else if c == "2" {
			studentdata.displaydata()
			nxoptn = nextoptn()
		} else if c == "3" {
			studentdata.cleardata()
			nxoptn = nextoptn()
		} else {
			nxoptn = "n"
		}
		if nxoptn == "N" || nxoptn == "n" {
			fmt.Println("Goodbye")
			break
		} else {
			continue
		}
	}
}

////////////////////////////////////////////////
func getdata() (studentdata student) { //This function is a struct method
	var studentfirstname, studentlastname string
	var studentscore float64

	fmt.Print("Please input the student's first name :")
	fmt.Scan(&studentfirstname)
	fmt.Print("Please input the student's last name :")
	fmt.Scan(&studentlastname)
	fmt.Print("Please input the student's score :")
	fmt.Scan(&studentscore)

	studentdata = student{
		firstname: studentfirstname,
		lastname:  studentlastname,
		score:     studentscore,
	}
	return
}

func (studentdata *student) displaydata() {
	fmt.Print("\nThe student's first name is ", (*studentdata).firstname)
	fmt.Print("\nThe student's last name is ", (*studentdata).lastname)
	fmt.Print("\nThe student's score is ", (*studentdata).score)
}

func (data *student) cleardata() { //This functions gets a pointer as an argument and return a de-referenced pointer
	fmt.Print("\n\n---Clearing Data---")
	(*data).firstname = "" //de-referenced pointer (value)
	(*data).lastname = ""  //de-referencing a value in a struct like this is not mandatory
	(*data).score = 0
	fmt.Print("\n")
}

func menu() (c string) {
	for {
		fmt.Println("Choose the following actions :")
		fmt.Println("1. Input data")
		fmt.Println("2. Display data")
		fmt.Println("3. Clear data")
		fmt.Println("4. Exit")
		fmt.Print("Your choice :")
		fmt.Scan(&c)

		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func nextoptn() (nxoptn string) {
	for {
		fmt.Print("\nNext operation? [Y/N] :")
		fmt.Scan(&nxoptn)
		if nxoptn == "Y" || nxoptn == "y" || nxoptn == "N" || nxoptn == "n" {
			return
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
