package main

import (
	"fmt"
)

type IPAddr [4]byte

/**
 * This method will be overridden.
 * While using the Print method of fmt,
 * instead of printing an []IPAddr Array, it will print what the function returns.
 * In this case an IP like x.x.x.x
 */
func (ip IPAddr) String() string {
	numItems := len(ip)
	if numItems == 4 {
		return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	}
	return "0.0.0.0"
}

/**
 * Switch is also available for types
 */
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d is an %T.\n", v, v)
	case string:
		fmt.Printf("%q is a %T.\n", v, v)
	default:
		fmt.Printf("Undefined Type >> %T << with value >> %v <<.\n", v, v)
	}
}

func main() {
	/**
	 * Same as:
	 *
	 * i := 42
	 * f := float64(i)
	 * u := uint(f)
	 *
	 */
	var (
		i int     = 42
		f float64 = float64(i)
		u uint    = uint(f)
	)

	fmt.Println("--- Listing the types ---")
	const p = "%T(%v) %T(%v) %T(%v)\n"
	fmt.Printf(p, i, i, f, f, u, u)

	v1 := 42
	v2 := 42.0
	v3 := "string"
	v4 := 0.5i
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)
	fmt.Printf("v4 is of type %T\n", v4)

	fmt.Println("--- Using type switch ---")
	do(42)
	do("hello")
	do(true)

	fmt.Println("--- Using modified string conversion >fmt.Stringer::String()< ---")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}