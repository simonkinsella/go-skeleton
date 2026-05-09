package luhn_test

import (
	"testing"

	"github.com/simonkinsella/go-skeleton/internal/luhn"
)

type Scenario struct {
	name     string
	input    string
	expected bool
}

var Scenarios []Scenario

func TestLunn(t *testing.T) {
	addInvalidScenarios()
	addValidScenarios()

	sut := luhn.NewChecker()
	for _, scenario := range Scenarios {
		got := sut.Check(scenario.input)
		if got != scenario.expected {
			t.Errorf("Scenario '%s', input:'%s' - expected %v, got %v", scenario.name, scenario.input, scenario.expected, got)
		}
	}
}

func addScenario(name string, input string, expected bool) {
	Scenarios = append(Scenarios, Scenario{name, input, expected})
}

func addInvalidScenarios() {
	addScenario("empty input", "", false)
	addScenario("not a number", "not a number", false)
	addScenario("zero", "0", false)
	addScenario("single digit", "2", false)
	addScenario("single digit", "7", false)
	addScenario("two digit", "35", false)
	addScenario("three digit", "876", false)
	addScenario("four digit", "8229", false)
	addScenario("five digit", "82885", false)
	addScenario("six digit", "517216", false)
	addScenario("seven digit", "5121568", false)
	addScenario("eight digit", "76117299", false)
	addScenario("nine digit", "951864751", false)
	addScenario("ten digit", "4365512614", false)
}

func addValidScenarios() {
	addScenario("two digit", "26", true)
	addScenario("two digit", "34", true)
	addScenario("three digit", "018", true)
	addScenario("three digit", "976", true)
	addScenario("four digit", "8219", true)
	addScenario("four digit", "5439", true)
	addScenario("five digit", "81885", true)
	addScenario("five digit", "28704", true)
	addScenario("six digit", "993782", true)
	addScenario("six digit", "056705", true)
	addScenario("seven digit", "5127568", true)
	addScenario("seven digit", "5545009", true)
	addScenario("eight digit", "76127299", true)
	addScenario("eight digit", "18523373", true)
	addScenario("nine digit", "950864751", true)
	addScenario("nine digit", "760943266", true)
	addScenario("ten digit", "4365512674", true)
	addScenario("ten digit", "5363215632", true)
}
