	package carros

	// RuleSet é uma estrutura que armazena as regras de dependência e exclusão mútua
	type RuleSet struct {
		dependencia map[string]string
		conflito    map[string]string
	}

	// NewRuleSet retorna um conjunto de regras vazio
	func NewRuleSet() *RuleSet {
		return &RuleSet{
			dependencia: make(map[string]string),
			conflito:    make(map[string]string),
		}
	}

	// AddDep adiciona uma nova dependência entre A e B
	func (rs *RuleSet) AddDep(A, B string) {
		rs.dependencia[A] = B
	}

	// AddConflict adiciona um novo conflito entre A e B
	func (rs *RuleSet) AddConflict(A, B string) {
		rs.conflito[A] = B
		rs.conflito[B] = A
	}

	func (rs *RuleSet) IsCoherent() bool {
		visited := make(map[string]bool)
		inStack := make(map[string]bool)

		var hasCycle func(string) bool
		hasCycle = func(option string) bool {
			if inStack[option] {
				return true
			}
			if visited[option] {
				return false
			}

			visited[option] = true
			inStack[option] = true

			dependency, depExists := rs.dependencia[option]
			conflict, confExists := rs.conflito[option]

			if depExists && hasCycle(dependency) {
				return true
			}

			if confExists {
				if inStack[conflict] {
					return true
				}
				if hasCycle(conflict) {
					return true
				}
			}

			inStack[option] = false

			return false
		}

		for option := range rs.dependencia {
			if hasCycle(option) {
				return false
			}
		}

		return true
	}
