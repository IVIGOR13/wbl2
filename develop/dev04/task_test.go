package main

import (
	"reflect"
	"testing"
)

func TestTask(t *testing.T) {
	m := *findAnagram(&[]string{"пятка", "пятак", "листок", "столик", "Тяпка", "Слиток", "арваы"})
	if len(m) != 2 {
		t.Errorf("wrong length")
	}

	if !reflect.DeepEqual(m["пятка"], []string{"пятак", "пятка", "тяпка"}) {
		t.Errorf("wrong array")
	}

	if !reflect.DeepEqual(m["листок"], []string{"листок", "слиток", "столик"}) {
		t.Errorf("wrong array")
	}
}
