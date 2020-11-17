package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Student struct {
		Name string `json:"studentName"`
		Age  int8
	}

	type ColorGroup struct {
		ID          int
		Name        string `json:"name"`
		Colors      []string
		note        string
		Student1    Student
		StudentList []Student
	}

	student12 := Student{
		Name: "liukai",
		Age:  12,
	}

	group := ColorGroup{
		ID:          1,
		Name:        "Reds",
		Colors:      []string{"Crimson", "Red", "Ruby", "Maroon"},
		Student1:    student12,
		StudentList: []Student{student12},
	}
	//Json Marshal：将数据编码成json字符串
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println("bbbbb")
	fmt.Println(string(b))

	group1 := ColorGroup{}
	group2 := ColorGroup{}
	group3 := ColorGroup{}
	//Json Unmarshal：将json字符串解码到相应的数据结构
	json.Unmarshal([]byte(string(b)), &group2)
	json.Unmarshal(b, &group1)
	json.Unmarshal([]byte(""), &group3)
	fmt.Println(group1.ID)
	fmt.Println(group1.Name)
	fmt.Println(group1.Colors)
	fmt.Println((group1.Colors)[0])
	fmt.Println((group1.StudentList)[0])
	fmt.Printf("%v", group1)
	fmt.Println()
	fmt.Println(group1.Student1.Age)
	fmt.Printf("%v", group2)
	fmt.Println()
	fmt.Printf("%v", group3)

}
