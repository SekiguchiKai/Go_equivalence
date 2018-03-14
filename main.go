package main

import (
	"fmt"
	"reflect"
)

const bar = "==========================================================="

type Person struct {
	Name string
	Age  int
}

type StructIncludingSlice struct {
	Numbers []int
}

type StructIncludingPointer struct {
	Pointer *int
}

func main() {
	executeExperiment1()
	executeExperiment2()
	executeExperiment3()
	executeExperiment4()
	executeExperiment5()

}

func executeExperiment1() {
	fmt.Println(bar)

	p1 := Person{
		Name: "太郎",
		Age:  20,
	}

	p2 := Person{
		Name: "太郎",
		Age:  20,
	}

	fmt.Println("単純な構造体の比較")
	fmt.Printf("p1 == p2 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p2)
	fmt.Printf("&p1 == &p2 : 等値 : %t\n", &p1 == &p2)
	fmt.Printf("reflect.DeepEqual(p1, p2) : 等価 : %t\n", reflect.DeepEqual(p1, p2))
	fmt.Printf("reflect.DeepEqual(&p1, &p2)  : %t\n", reflect.DeepEqual(&p1, &p2))

	fmt.Println(bar)
	fmt.Println("単純な構造体の比較2")
	p3 := p1
	fmt.Printf("p1 == p3 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p3)
	fmt.Printf("&p1 == &p3 : 等値 : %t\n", &p1 == &p3)
	fmt.Printf("reflect.DeepEqual(p1, p3) : 等価 : %t\n", reflect.DeepEqual(p1, p3))
	fmt.Printf("reflect.DeepEqual(&p1, &p3) : %t\n", reflect.DeepEqual(&p1, &p3))
}

func executeExperiment2() {
	fmt.Println(bar)
	fmt.Println("構造体のポインタ型の比較(同一ポインタ)")

	p1 := Person{
		Name: "太郎",
		Age:  20,
	}

	p2 := Person{
		Name: "太郎",
		Age:  20,
	}

	p3 := &p1
	fmt.Printf("&p1 == p3 : 等値 : %t\n", &p1 == p3)
	fmt.Printf("reflect.DeepEqual(&p1, p3) : 等価 : %t\n", reflect.DeepEqual(&p1, p3))

	fmt.Println(bar)
	fmt.Println("構造体のポインタ型の比較(別ポインタ)")
	p4 := &p2
	fmt.Printf("&p1 == p5 : 等値 : %t\n", &p1 == p4)
	fmt.Printf("reflect.DeepEqual(&p1, p4) : 等価 : %t\n", reflect.DeepEqual(&p1, p4))

}

func executeExperiment3() {
	fmt.Println(bar)

	sa := []int{0, 1, 2, 3, 4, 5}
	sb := []int{0, 1, 2, 3, 4, 5}

	s1 := StructIncludingSlice{
		Numbers: sa,
	}

	s2 := StructIncludingSlice{
		Numbers: sb,
	}

	fmt.Println("フィールドにSliceが入っている構造体の比較")
	// fmt.Printf("等価(メモリが別でも値が一緒だったらOK) : %t\n", s1 == s2) // struct containing []int cannot be compared

	fmt.Println("s1 == s2 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == &s2 : 等値 : %t\n", &s1 == &s2)
	fmt.Printf("reflect.DeepEqual(s1, s2) : 等価 : %t\n", reflect.DeepEqual(s1, s2))
	fmt.Printf("reflect.DeepEqual(&s1, &s2) : %t\n", reflect.DeepEqual(&s1, &s2))

	fmt.Println(bar)
	fmt.Println("フィールドにSliceが入っている構造体の比較(Sliceが同じ物)")
	s3 := StructIncludingSlice{
		Numbers: sa,
	}

	fmt.Println("s1 == s3 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == &s3 : 等値 : %t\n", &s1 == &s3)
	fmt.Printf("reflect.DeepEqual(s1, s3) : 等価 : %t\n", reflect.DeepEqual(s1, s3))
	fmt.Printf("reflect.DeepEqual(&s1, &s3)  : %t\n", reflect.DeepEqual(&s1, &s3))

	fmt.Println(bar)
	fmt.Println("フィールドにSliceが入っている構造体の比較(同じアドレスの場合)")
	s4 := &s1

	fmt.Println("s1 == s4 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == s4 : 等値 : %t\n", &s1 == s4)
	fmt.Printf("reflect.DeepEqual(s1, s4) : 等価 : %t\n", reflect.DeepEqual(s1, s4))
	fmt.Printf("reflect.DeepEqual(&s1, s4)  : %t\n", reflect.DeepEqual(&s1, s4))

}

func executeExperiment4() {
	fmt.Println(bar)
	fmt.Println("StructIncludingPointerの実験")
	num1 := 1
	num2 := 1

	s1 := StructIncludingPointer{
		Pointer: &num1,
	}

	s2 := StructIncludingPointer{
		Pointer: &num2,
	}

	// プロパティにポインタ（スライス含む）が入っていると動作
	fmt.Println("フィールドに異なるアドレスを指す同じ数字が入っている構造体の比較")
	fmt.Printf("s1 == s2 : この場合、構造体のフィールドの等値での比較になる : %t\n", s1 == s2)
	fmt.Printf("&s1 == &s2 : %t\n", &s1 == &s2)
	fmt.Printf("reflect.DeepEqual(s1, s2) : 等価 : %t\n", reflect.DeepEqual(s1, s2))
	fmt.Printf("reflect.DeepEqual(&s1, &s2) : %t\n", reflect.DeepEqual(&s1, &s2))
}

func executeExperiment5() {
	fmt.Println(bar)

	numX := 2
	s3 := StructIncludingPointer{
		Pointer: &numX,
	}

	s4 := StructIncludingPointer{
		Pointer: &numX,
	}

	fmt.Println("フィールドに同一アドレスを指す同じ数字が入っている構造体の比較")
	fmt.Printf("s3 == s4 : 構造体のフィールドの等値での比較になる : %t\n", s3 == s4)
	fmt.Printf("&s3 == &s4 : %t\n", &s3 == &s4)
	fmt.Printf("reflect.DeepEqual(s3, s4) : 等価 : %t\n", reflect.DeepEqual(s3, s4))
	fmt.Printf("reflect.DeepEqual(&s3, &s4)  : %t\n", reflect.DeepEqual(&s3, &s4))

	fmt.Println(bar)
}
