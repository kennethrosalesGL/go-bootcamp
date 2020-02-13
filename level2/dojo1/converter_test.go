package main
import (
	"fmt"
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	expected_data := []string{"DOG", "CAT", "SHE3P", "SNAKE"}
	validateTranform(expected_data, applyTransform(animals(), strings.ToUpper), t)
}

func TestCapitalize(t *testing.T) {
	expected_data := []string{"Dog", "Cat", "She3p", "Snake"}
	validateTranform(expected_data, applyTransform(animals(), capitalize), t)
}

func TestReverse(t *testing.T) {
	expected_data := []string{"GoD", "tac", "P3ehs", "EKANS"}
	validateTranform(expected_data, applyTransform(animals(), reverse), t)
}
func TestEcho(t *testing.T) {
	expected_data := []string{"DoG", "cat", "she3P", "SNAKE"}
	validateTranform(expected_data, applyTransform(animals(), echo), t)
}
func validateTranform(expected_data []string, transformed []string, t *testing.T) {
	if len(expected_data) != len(transformed) {
		t.Errorf("Transformed slice has different length than the expected")
	}
	for i, item := range expected_data {
		if item != transformed[i] {
			t.Errorf("Value not transformed correctly")
		}
	}
}
func animals() []string {
	return []string{"DoG", "cat", "she3P", "SNAKE"}
}
func capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}
func echo(s string) string {
	fmt.Println(s)
	return s
}
func reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	  }
	  return 
}
