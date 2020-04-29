package main

import (
	"fmt"
	"testing"
)

var err = fmt.Errorf("something failed because this and that")

func BenchmarkSprint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprint("Something went wrong, more details: ", err)
    }
}

func BenchmarkSprintf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("Something went wrong, more details: %v", err)
    }
}