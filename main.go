package main

import (
	carros "Carros/Carros"
	"fmt"
)

func main() {
	ruleSet := carros.NewRuleSet()
	fmt.Println(ruleSet.IsCoherent())
}
