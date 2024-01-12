package carros_test

import (
	carros "Carros/Carros"
	"testing"
)

// TestDependsAA verifica se um conjunto de regras com dependência cíclica retorna incoerente
func TestDependsAA(t *testing.T) {
	ruleSet := carros.NewRuleSet()
	ruleSet.AddDep("a", "a")

	if ruleSet.IsCoherent() {
		t.Error("TestDependsAA: Falhou - O conjunto de regras deveria ser incoerente.")
	}
}

// TestDependsAB_BA verifica se um conjunto de regras com dependências mutuas retorna incoerente
func TestDependsAB_BA(t *testing.T) {
	ruleSet := carros.NewRuleSet()

	ruleSet.AddDep("a", "b")
	ruleSet.AddDep("b", "a")

	if ruleSet.IsCoherent() {
		t.Error("TestDependsAB_BA: Falhou - O conjunto de regras deveria ser incoerente.")
	}
}

// TestExclusiveAB verifica se um conjunto de regras com dependência e exclusão mútua retorna incoerente
func TestExclusiveAB(t *testing.T) {
	ruleSet := carros.NewRuleSet()

	ruleSet.AddDep("a", "b")
	ruleSet.AddConflict("a", "b")

	if !ruleSet.IsCoherent() {
		t.Error("TestExclusiveAB: Falhou - O conjunto de regras deveria ser incoerente.")
	}
}

// TestExclusiveAB_BC verifica se um conjunto de regras com dependências e exclusões mútuas retorna incoerente
func TestExclusiveAB_BC(t *testing.T) {
	ruleSet := carros.NewRuleSet()

	ruleSet.AddDep("a", "b")
	ruleSet.AddDep("b", "c")
	ruleSet.AddConflict("a", "c")

	if !ruleSet.IsCoherent() {
		t.Error("TestExclusiveAB_BC: Falhou - O conjunto de regras deveria ser incoerente.")
	}
}

// TestDeepDeps verifica se um conjunto de regras com dependências profundas e exclusão mútua retorna incoerente
func TestDeepDeps(t *testing.T) {
	ruleSet := carros.NewRuleSet()

	ruleSet.AddDep("a", "b")
	ruleSet.AddDep("b", "c")
	ruleSet.AddDep("c", "d")
	ruleSet.AddDep("d", "e")
	ruleSet.AddDep("a", "f")
	ruleSet.AddConflict("e", "f")

	if !ruleSet.IsCoherent() {
		t.Error("TestDeepDeps: Falhou - O conjunto de regras deveria ser incoerente.")
	}
}
