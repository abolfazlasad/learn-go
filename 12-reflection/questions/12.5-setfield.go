// Question 12.5 — Setting values with reflect
//
// Show that reflect.ValueOf(x).CanSet() is false, but
// reflect.ValueOf(&x).Elem().CanSet() is true. Set x to a new int through reflect.
//
// Extra: setField(ptr any, name string, val any) error that sets an exported
// struct field by name. Error on unknown or unexported fields.
//
// You will need: reflect.Value, CanSet, SetInt / Set, Elem.
//
// Run:
//   go run 12.5-setfield.go

package main

func main() {
	// TODO: write your solution
}
