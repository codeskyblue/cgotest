package main

/*
#include <stdio.h>
*/
import "C"

func PrintHello() {
	C.puts(C.CString("Hello world"))
}

func main() {
	PrintHello()
}
