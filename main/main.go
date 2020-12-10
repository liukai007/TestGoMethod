package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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
	student1 := Student{
		Name: "liukai1",
		Age:  1,
	}
	student2 := Student{
		Name: "liukai2",
		Age:  2,
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

	fmt.Println()
	fmt.Println("bb")
	fmt.Println("bb")
	fmt.Println("bb")
	student123 := []Student{student12, student1, student2}
	bb, err := json.Marshal(student123)
	fmt.Println(string(bb))
	st11 := student123[0]
	fmt.Println(st11.Age)
	sss1 := []Student{}
	json.Unmarshal(bb, &sss1)
	fmt.Println(sss1[0].Age)
	/*
		https://www.cnblogs.com/mafeng/p/10364787.html
		正则表达式 的例子


	*/

	str := "880218end"
	fmt.Println(str)
	match, _ := regexp.MatchString("\\d{1}", str) //六位连续的数字
	fmt.Println(match)                            //输出true
	reg := regexp.MustCompile("\\d{6}")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	fmt.Println(string(data)) //880218
	//go语言正则表达式判断是否为汉字
	matchChinese, _ := regexp.Match("[\u4e00-\u9fa5]", []byte("经度"))
	fmt.Println(matchChinese) //输出true
	//go语言正则表达式判断是否含有字符（大小写）
	matchChar, _ := regexp.Match("[a-zA-Z]", []byte("av132"))
	fmt.Println(matchChar) //输出false
	//go语言正则表达式判断是否含有以数字开头，不是为true
	matchDigit, _ := regexp.Match(`[^\d]`, []byte("as132"))
	fmt.Println(matchDigit) //输出true
	//go语言正则表达式判断是否含有为IP地址
	ip := "10.32.12.01"
	pattern := "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+"
	matchIp, _ := regexp.MatchString(pattern, ip)
	fmt.Println(matchIp) //输出true
	//go语言正则表达式判断是否包含某些字段
	id := "id=123;dfg"
	reg = regexp.MustCompile("id=[\\d]+")
	MEId := reg.FindString(id)
	fmt.Println(MEId) //输出id=123
	// 判断是否是个单词
	pattern = "^[A-Z]+$|(^[A-Z]?[a-z]+)$"
	char := "mAfeng"
	match, _ = regexp.Match(pattern, []byte(char))
	fmt.Println(match)

}
