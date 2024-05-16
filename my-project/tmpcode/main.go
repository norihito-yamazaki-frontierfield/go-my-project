package main

import (
	"fmt"
)

type MyInt int

type EnhancedInt int

// EnhancedInt のための String メソッド
func (e EnhancedInt) String() string {
	return fmt.Sprintf("EnhancedInt: %d", e)
}

type Float32 float32
type Float64 float64

// displayStringable は IntStringer 型制約を満たす任意の型 T の値を表示します
func displayStringable[T IntStringer](value T) {
	fmt.Println("IntStringer:", value.String())
}

// displayFloat は Float 型制約を満たす任意の型 T の値を表示します
func displayFloat[T Float](value T) {
	fmt.Printf("Float: %v\n", value)
}

// IntStringer は、基底型が int で String メソッドを実装する型のインターフェースです
type IntStringer interface {
	~int
	String() string
}

// Float は float32 または float64 の型制約を持つインターフェースです
type Float interface {
	~float32 | ~float64
}

func main() {
	var ei EnhancedInt = 42
	displayStringable(ei) // EnhancedInt は IntStringer を満たす

	var f32 Float32 = 3.14
	displayFloat(f32) // Float32 は Float を満たす
}
