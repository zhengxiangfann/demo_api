package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const middle = "====="

//Config ini 文件结构体
type Config struct {
	Mymap  map[string]string
	strect string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strect = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strect) == 0 {
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		first := strings.TrimSpace(s[:index])
		if len(first) == 0 {
			continue
		}

		second := strings.TrimSpace(s[index+1:])

		// if len(second) == 0 {
		// 	continue
		// }

		pos := strings.Index(second, "\t#")

		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "#")

		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "//")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strect + middle + first
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c *Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
