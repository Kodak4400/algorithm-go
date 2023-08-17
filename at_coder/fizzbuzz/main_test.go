package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzz := FizzBuzz(15)
	if fizzBuzz != "FizzBuzz" {
		t.Errorf("%v != %v", fizzBuzz, "FizzBuzz")
	}

	fizz := FizzBuzz(3)
	if fizz != "Fizz" {
		t.Errorf("%v != %v", fizz, "Fizz!")
	}

	buzz := FizzBuzz(5)
	if buzz != "Buzz" {
		t.Errorf("%v != %v", buzz, "Buzz!")
	}

	num := FizzBuzz(2)
	if num != "2" {
		t.Errorf("%v != %v", num, "2")
	}
}
