package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
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

// func main() {

// 	// myconf := new(conf.Config)
// 	// myconf.InitConfig("config.ini")
// 	// fmt.Println(myconf.Read("default", "path"))
// 	// fmt.Println(myconf.Mymap)
// 	// fmt.Printf("Run Appliction!\r\n")
// 	// fmt.Println("------------------------------------")
// 	// stu()
// 	// fmt.Println("------------------------------------")
// 	// chanstu()
// 	// fmt.Println("------------------------------------")

// 	//stuchan3()

// 	// var result = func(a int) string { return strconv.Itoa(a) }(1)
// 	// fmt.Println(result)

// 	var generator EmployeeIdGenerator

// 	generator = func(company string, department string, sn uint32) string {

// 		return appendSn(company+"-"+department+"-", sn)

// 	}
// 	fmt.Println(generateId(generator, "RD"))

// }

// func main() {
// demo1()
// demo2()
// demo3()
// demo4()
// demoo5()
// demo6()
// demo7()

// nextInt := demo8()
// fmt.Println(nextInt())
// fmt.Println(nextInt())
// fmt.Println(nextInt())
// newInt := demo8()
// fmt.Println(newInt())

// i := 1
// fmt.Println("initial:", i)
// zeroval(i)
// fmt.Println("zeroval:", i)
// zeroptr(&i)
// fmt.Println("zeroptr", i)
// fmt.Println("zeroptr", &i)
// demo11()

// r := rect{width: 10, height: 5}
// fmt.Println("area:", r.area())
// fmt.Println("perim:", r.perim())

// 	r := rect{width: 3, height: 4}

// 	c := circle{radius: 5}
// 	measure(r)
// 	measure(c)

// }

func demo() {
	fmt.Println("demo")
	err, l1 := "111", "222"
	fmt.Println(l1 + err)
	err2, l1 := "3333", "444"
	fmt.Println(err2 + l1)

	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		break
	}

}

func demo1() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after nood")
	}

}

func demo2() {
	var a [5]int
	fmt.Println("set", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var two [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = i + j
		}
	}

	fmt.Println("two", two)

}

func demo3() {

	s := make([]int, 3)
	fmt.Println("emp", s)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("set", s)

	s = append(s, 4)
	s = append(s, 5)
	fmt.Println("append", s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)

}

func demo4() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k1")

	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

func demo5() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for _, c := range "go" {
		fmt.Print(c)
	}
}

func demo6() (int, int) {
	return 3, 7
}

func demo7(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func demo8() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func demo11() {
	type person struct {
		name string
		age  int
	}

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int {
// 	return r.width * r.height
// }

// func (r rect) perim() int {
// 	return 2*r.width + 2*r.height
// }

type geomery interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {

	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geomery) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	router := gin.Default()
	type Login struct {
		User     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password"`
	}
	router.POST("/postfrom", func(c *gin.Context) {
		var json Login
		res := c.BindJSON(&json)
		if res == nil {
			c.JSON(200, json)
			return
		}
		c.String(200, "ffff")
	})

	router.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		c_cp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note than you are using the copied context "c_cp", IMPORTANT
			log.Println("Done! in path " + c_cp.Request.URL.Path)
		}()
		c.String(200, "async")
	})

	router.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
		c.String(200, "sync")
	})

	router.Run(":9090")

}
