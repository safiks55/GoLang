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

func deleteStudent(id int, s *[]student) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i].idNum == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}

func updateStudent(input int, id int, info string, s *[]student) {
	switch {
	case input == 1:
		for i := 0; i < len(*s); i++ {
			if (*s)[i].idNum == id {
				(*s)[i].firstName = info
			}
		}
		break
	case input == 2:
		for i := 0; i < len(*s); i++ {
			if (*s)[i].idNum == id {
				(*s)[i].lastName = info
			}
		}
		break
	case input == 3:
		for i := 0; i < len(*s); i++ {
			if (*s)[i].idNum == id {
				(*s)[i].email = info
			}
		}
		break
	case input == 4:
		temp, _ := strconv.Atoi(info)
		for i := 0; i < len(*s); i++ {
			if (*s)[i].idNum == id {
				(*s)[i].idNum = temp
			}
		}
		break
	case input == 5:
		temp2, _ := strconv.ParseInt(info, 10, 64)
		for i := 0; i < len(*s); i++ {
			if (*s)[i].idNum == id {
				(*s)[i].phoneNumber = temp2
			}
		}
		break
	}
}

func printStudents(s []student) {
	fmt.Println("************ LIST OF STUDENTS IN THE DATABASE ************")
	if len(s) == 0 {
		fmt.Println("Database is empty")
	} else {
		for i := 0; i < len(s); i++ {
			fmt.Println(i+1, ":", "FirstName:", s[i].firstName, "LastName:", s[i].lastName, "email:", s[i].email, "id#:", s[i].idNum, "Phone#:", s[i].phoneNumber)
		}
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
	var input, input2 int
	fmt.Println(" Welcome to Student Management System ")
	for ok := true; ok; ok = input != 5 {
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
			for ok := true; ok; ok = input2 != 7 {
				fmt.Println("Please select from the following options:")
				fmt.Println("1. Update firstName")
				fmt.Println("2. Update lastName")
				fmt.Println("3. Update email")
				fmt.Println("4. Update idNum")
				fmt.Println("5. Update email")
				fmt.Println("6. Return to Main Menu")

				n, err := fmt.Scanln(&input2)
				if n < 1 || err != nil {
					fmt.Println("invalid input")
					break
				}
				switch input2 {
				case 1:
					fmt.Println("Please enter the student id: ")
					in := bufio.NewScanner(os.Stdin)
					in.Scan()
					id, _ := strconv.Atoi(in.Text())

					fmt.Println("Please enter the new firstName: ")
					temp := bufio.NewScanner(os.Stdin)
					temp.Scan()
					fName := temp.Text()
					updateStudent(input2, id, fName, &db)
					break
				case 2:
					fmt.Println("Please enter the student id: ")
					in := bufio.NewScanner(os.Stdin)
					in.Scan()
					id, _ := strconv.Atoi(in.Text())

					fmt.Println("Please enter the new lastName: ")
					temp := bufio.NewScanner(os.Stdin)
					temp.Scan()
					lName := temp.Text()
					updateStudent(input2, id, lName, &db)

					break
				case 3:
					fmt.Println("Please enter the student id: ")
					in := bufio.NewScanner(os.Stdin)
					in.Scan()
					id, _ := strconv.Atoi(in.Text())

					fmt.Println("Please enter the new email: ")
					temp := bufio.NewScanner(os.Stdin)
					temp.Scan()
					email := temp.Text()
					updateStudent(input2, id, email, &db)

					break
				case 4:
					fmt.Println("Please enter the student id: ")
					in := bufio.NewScanner(os.Stdin)
					in.Scan()
					id, _ := strconv.Atoi(in.Text())

					fmt.Println("Please enter the new idNumber: ")
					temp := bufio.NewScanner(os.Stdin)
					temp.Scan()
					idNum := temp.Text()
					updateStudent(input2, id, idNum, &db)
					break
				case 5:
					fmt.Println("Please enter the student id: ")
					in := bufio.NewScanner(os.Stdin)
					in.Scan()
					id, _ := strconv.Atoi(in.Text())

					fmt.Println("Please enter the new phoneNumber: ")
					temp := bufio.NewScanner(os.Stdin)
					temp.Scan()
					phoneNumber := temp.Text()
					updateStudent(input2, id, phoneNumber, &db)
					break
				case 6:
					fmt.Println("Returning to main menu...")
					time.Sleep(2 * time.Second)
					break
				}
			}
			break
		case 3:
			fmt.Println("Please enter the student id number: ")
			in := bufio.NewScanner(os.Stdin)
			in.Scan()
			id, _ := strconv.Atoi(in.Text())
			deleteStudent(id, &db)
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
