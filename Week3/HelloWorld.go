package main

/*
#include <stdio.h>
void printHelloWorld() {
	printf("Hello, World!\n");
}
*/

import "C"

func main() {
	C.printHelloWorld()
}
