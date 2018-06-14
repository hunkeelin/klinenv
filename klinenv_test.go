package klinenv

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("testing")
	config := NewAppConfig("testconfig")
	fmt.Println(config)
}
