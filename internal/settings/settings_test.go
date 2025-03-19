package settings

import (
	"testing"
)

func TestBindInterface(t *testing.T) {
	args := []string{"-interface", "127.0.0.1"}

	config := Recover(args)

	if config.Interface != "127.0.0.1" {
		t.Errorf("expected '127.0.0.1' interface but got '%v'", config.Interface)
	}
}

func TestBindInterfaceDefault(t *testing.T) {
	args := []string{}

	config := Recover(args)

	if config.Interface != "0.0.0.0" {
		t.Errorf("expected Interface '0.0.0.0' but got '%v'", config.Interface)
	}
}

func TestBindPort(t *testing.T) {
	args := []string{"-port", "18201"}

	config := Recover(args)

	if config.Port != 18201 {
		t.Errorf("expected Port '18201' but got '%v'", config.Port)
	}
}

func TestBindPortDefault(t *testing.T) {
	args := []string{}

	config := Recover(args)

	if config.Port != 8080 {
		t.Errorf("expected Port '8080' but got '%v'", config.Port)
	}
}
