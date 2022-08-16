package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type student struct {
	firstName, lastName, email string
	idNum                      int
	phoneNumber                int64
}

func addStudent(firstName, lastName, email string, id int, phoneNumber int64, s *[]student) {
	temp := new(student)
	temp.idNum = id
	temp.phoneNumber = phoneNumber
	temp.firstName = firstName
	temp.lastName = lastName
	temp.email = email
	*s = append(*s, *temp)
}

func deleteStudent(id int, s []student) {
	for i := 0; i < len(s); i++ {
		if s[i].idNum == id {
			s = append(s[:i], s[i+1])
		}
	}
}

func printStudents(s []student) {
	fmt.Println("************ LIST OF STUDENTS IN THE DATABASE ************")
	for i := 0; i < len(s); i++ {
		fmt.Println(i+1, " FirstName: ", s[i].firstName, " LastName: ", s[i].lastName, " email: ", s[i].email, " id#: ", s[i].idNum, " Phone#: ", s[i].phoneNumber)
	}
}

func main() {
	var db []student // array of type student
	p := student{
		firstName:   "safik",
		lastName:    "sunesara",
		email:       "ss123@ymail.com",
		idNum:       -1,
		phoneNumber: 8322111234,
	}
	db = append(db, p)
	var input int
	fmt.Println(" Welcome to Student Management System ")
	for ok := true; ok; ok = (input != 5) {
		fmt.Println(" Please select from the following options")
		fmt.Println("1. Add Student")
		fmt.Println("2. Update Student")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Print all Students")
		fmt.Println("5. Quit")

		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			break
		}
		switch input {
		case 1:

			fmt.Println("Please enter student firstName: ")
			temp := bufio.NewScanner(os.Stdin)
			temp.Scan()
			firstName := temp.Text()

			fmt.Println("Please enter student lastName: ")
			temp = bufio.NewScanner(os.Stdin)
			temp.Scan()
			lastName := temp.Text()

			fmt.Println("Please enter student email: ")
			temp = bufio.NewScanner(os.Stdin)
			temp.Scan()
			email := temp.Text()

			fmt.Println("Please enter student id: ")
			temp2 := bufio.NewScanner(os.Stdin)
			temp.Scan()
			dummy, _ := strconv.Atoi(temp.Text())
			idNum := dummy

			fmt.Println("Please enter student phonenumber: ")
			temp2 = bufio.NewScanner(os.Stdin)
			temp2.Scan()
			dummy2, _ := strconv.ParseInt(temp2.Text(), 10, 64)
			phoneNumber := dummy2
			addStudent(firstName, lastName, email, idNum, phoneNumber, &db)
			fmt.Println("New Student Added Successfully To The DataBase!!")
			break
		case 2:
			fmt.Println("2")
			break
		case 3:
			fmt.Println("3")
			break
		case 4:
			printStudents(db)
			break
		case 5:
			fmt.Println("System Shutting Down in")
			fmt.Println("3sec...")
			time.Sleep(1 * time.Second)
			fmt.Println("2sec...")
			time.Sleep(1 * time.Second)
			fmt.Println("1sec...")
			time.Sleep(1 * time.Second)
			fmt.Println("Exited Successfully!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Input!!")

		}

	}
}
