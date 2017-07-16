package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

func stu() {
	mm2 := map[string]int{"golang": 3, "python": 1, "csharp": 2}
	fmt.Printf("%d, %d, %d\n", mm2["python"], mm2["golang"], mm2["csharp"])
}

func chanstu() {
	ch2 := make(chan string, 1)

	go func() {
		ch2 <- ""
	}()

	var value string = "数据"
	value = value + ""
	fmt.Println(value)
}

func stuchan3() {
	var mychannel = make(chan int, 0)
	var number = 6
	type Sender chan<- int
	type Receiver <-chan int
	go func() {
		var sender Sender = mychannel
		sender <- number
		fmt.Println("Sent!")
	}()

	go func() {
		var receiver Receiver = mychannel
		fmt.Println("Received!", <-receiver)
	}()

	time.Sleep(time.Second)
}

// type EmployeeIdGenerator func(company string, department string, sn uint32) string
type EmployeeIdGenerator func(company string, department string, sn uint32) string

var company = "Gophers"
var sn uint32

// func generateId(generator EmployeeIdGenerator, department string) (string, bool){
// 	if generator == nil {
// 		return "", false
// 	}
// 	newSn := atomic.AddUint32(&sn, 1)
// 	return generator(company, department)
// }

// 生成员工ID
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	// 这是一条 if 语句，我们会在下一章讲解它。
	// 若员工ID生成器不可用，则无法生成员工ID，应直接返回。
	if generator == nil {
		return "", false
	}
	// 使用代码包 sync/atomic 中提供的原子操作函数可以保证并发安全。
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

type Person struct {
    Name    string
	Gender  string
	Age     uint8
	Address string
}

func (p *Person) Move(newaddr string) string {
    oldaddr := p.Address
    p.Address = newaddr
    return oldaddr
}

func main() {

	// myconf := new(conf.Config)
	// myconf.InitConfig("config.ini")
	// fmt.Println(myconf.Read("default", "path"))
	// fmt.Println(myconf.Mymap)
	// fmt.Printf("Run Appliction!\r\n")
	// fmt.Println("------------------------------------")
	// stu()
	// fmt.Println("------------------------------------")
	// chanstu()
	// fmt.Println("------------------------------------")

	//stuchan3()

	// var result = func(a int) string { return strconv.Itoa(a) }(1)
	// fmt.Println(result)

	var generator EmployeeIdGenerator

	generator = func(company string, department string, sn uint32) string {

		return appendSn(company+"-"+department+"-", sn)

	}
	fmt.Println(generateId(generator, "RD"))

}
