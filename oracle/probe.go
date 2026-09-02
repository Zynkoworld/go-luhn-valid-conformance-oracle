package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



import "strings"

func Valid(id string) bool {
	if len(strings.TrimSpace(id)) == 1 {
		return false
	}

	d := make([]int, 0, len(id))

	for _, r := range id {
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		d = append(d, int(r-'0'))
	}

	return sum(d)%10 == 0
}

func sum(d []int) (s int) {
	for i, x := range d {
		j := len(d) - i
		if j%2 == 0 {
			x *= 2
			if x > 9 {
				x -= 9
			}
		}
		s += x
	}
	return s
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []string{"1", "0", "059", "055 444 285", "055 444 286", "8273 1232 7352 0569", "1 2345 6789 1234 5678 9013", "095 245 88", "059a", "055-444-285", "055# 444$ 285", "0000 0", "091", "9999999999 9999999999 9999999999 9999999999", "055b 444 285", ":9"}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, Valid(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
