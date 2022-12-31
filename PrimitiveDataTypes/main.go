package main

import "fmt"

// const block
const (
	first  = iota
	second = iota
)

// iota reset in different const block
const (
	third = iota
	fourth
)

func main() {
	// 1
	var i int
	i = 42
	fmt.Println(i)

	// 2
	var f float32 = 3.14
	fmt.Println(f)

	// 3
	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// Pointer
	var pFirstName *string = new(string)
	*pFirstName = "Arthur"
	fmt.Println(pFirstName, *pFirstName)

	// Defering
	pName := "Arthur"
	ptr := &pName
	fmt.Println(ptr, *ptr)
	pName = "Tricia"
	fmt.Println(ptr, *ptr)

	// constant
	const k int = 3
	fmt.Println(float32(k) + 1.2)

	// iota
	fmt.Println(first, second, third, fourth)

	// ========= Collections =========
	fmt.Println("========= Collections =========")
	// *** array: fix size entity
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// array 2
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// *** slice: not fix size
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

	// *** map
	m := map[string]int{"foo": 42, "a": 123}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println("After delete Foo: ", m)

	// *** struct: the only collection contain different types
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	u2 := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println("u2: ", u2)
}
