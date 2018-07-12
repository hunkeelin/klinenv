package klinenv

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("testing")
	config := NewAppConfig("testconfig")
	fmt.Println(config.Get("port"))
	configv2, err := NewAppConfigv2("testconfig")
	fmt.Println(configv2.Get("port"))
	fmt.Println(err)
}
