package main

import (
	"fmt"
	"math"
)

func main() {
	numberBinary()
	separator()
	convertTypeInt()
	separator()
	numberHexadecimal()
	separator()
	numberMaxFloat()
}

func separator() {
	fmt.Println("--------------------------------")
}

func numberBinary() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<1) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}

func convertTypeInt() {
	var apples int32 = 1
	var oranges int16 = 2
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)
}

func numberHexadecimal() {
	o := 0777
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

func numberMaxFloat() {
	float32 := math.MaxFloat32
	float64 := math.MaxFloat64
	fmt.Printf("%g\n%g\n", float32, float64)

	separator()

	var z = float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	// nan := math.NaN()
	// fmt.Println(nan == nan, nan < nan, nan > nan)

}
