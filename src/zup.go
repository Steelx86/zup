package main

import (
	"fmt"
)

func createZup() Zup {
	return Zup{}
}

func readZupString(zupString string) Zup {
	return Zup{}
}

func (z *Zup) String() string {
	return fmt.Sprintf("%v", z)
}
