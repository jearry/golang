package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//
type TestDef struct {
	A int
	B string
}

func main() {
	var a, b int
	fmt.Println("请输入一个数字：")
	fmt.Scanf("%d", &a)

	fmt.Println("请输入另一个数字：")
	fmt.Scanf("%d", &b)

	fmt.Println("和是：", a+b, "积是：", a*b, "差是：", a-b, "商是：", a/b)
	c := 2.1 + 3.3i

	fmt.Printf("c: %s, %20.*f\n", fmt.Sprintf("%f+%f", real(c), imag(c)), 3, c)

	var v []rune = []rune("请输入一个数字：")
	fmt.Printf("%+q, %T, %#v\n", string(v), v, v)

	sv := TestDef{1, "2"}

	fmt.Printf("%+q\n", sv)

	r := strings.EqualFold("bA", "Ba")

	m := map[int]string{1: "1"}
	fmt.Println("r", r, "m", m)

	case1 := unicode.SpecialCase{
		{'a', 'z', [unicode.MaxCase]rune{'A' - 'a', 'Z' - 'z', 0}},
	}

	fmt.Println("upper:", strings.ToUpperSpecial(case1, "abc"))

	bt := []byte{'a', 'b', 'c'}

	fmt.Println("quote:", string(strconv.AppendQuote(bt, "abc")))

	bl, e1 := strconv.ParseBool("1")

	var b2 bool

	b3, e2 := fmt.Sscanf("1", "%t", &b2)

	fmt.Println("bool:", bl, b2, b3, e1, e2)

	mapt := make(map[string]string, 10)

	mapt["1"] = "a1"

    fmt.Println("mapt len", len(mapt))

	delete(mapt, "1")

    fmt.Println("mapt len", len(mapt))
}
