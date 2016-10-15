package main

import (
	"fmt"
	"strconv"
	"sort"
)



type product struct {
	name    string
	price   int
	}

type products []product

//Type Stringer() can be found in the fmt package. Every type that
// implements the stringer interface can be passed to fmt.Print,
// which will output the type's method String()
//type Stringer interface{
//	String() string
//}

var p product

// Implement Sort Interface, which requires three
// methods Len(), Swap() and Less()
func (g products) Len() int {
	return len(g)
}

func (g products) Swap(i,j int)  {g[i], g[j] = g[j], g[i]}

func (g products) Less(i, j int) bool {
    if g[i].price > g[j].price {
        return true
    }
    return false
}

func (p product) String() string {
	
	
	a:=p.name + ":" + strconv.Itoa(p.price) 
	return a
}



func main() {

	// The empty interface doesn't contain any method.
	// useful when we need to store any type.
	var a interface{}
	var i int = 5
	s := "Hello world"
	// These are legal statements
	a = i
	a = s

	p.name = "A Product name"
	fmt.Println(a, s, i, p.String())

	//pr :=new(products)

	pr := products{
		product{"a",1000},
		product{"b",500},
		product{"c",10000},
	}

	sort.Sort(pr)
	fmt.Println(pr)
}

/*
 Still to go over a bit more on interfaces
 Go over again pointers, references and dereferences
 */