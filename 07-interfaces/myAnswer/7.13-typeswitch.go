// Question 7.13 — Type switch
//
// stringify(v any) string using switch v := v.(type) for int, string, and default.
// Print a few cases from main.
//
// Extra: sqlLiteral(v any) string — quote strings, print numbers, "NULL" for nil.
//
// You will need: any, type switch, fmt.
//
// Run:
//   go run 7.13-typeswitch.go

package main

import "fmt"

func stringify(v any) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case nil:
		return "NULL"
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func main() {
	fmt.Println("int:", stringify(3))
	fmt.Println("string:", stringify("hello"))
	fmt.Println("bool:", stringify(true))
	fmt.Println("bool:", stringify(false))
	fmt.Println("nil:", stringify(nil))
}
