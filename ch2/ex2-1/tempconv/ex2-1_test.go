package tempconv

import (
	"fmt"
	"testing"
)

func TestKelvin(t *testing.T) {
	var zeroK Kelvin = 0
	var boilingK Kelvin = 373.15

	var valueC Celsius = KToC(zeroK)
	var valueK Kelvin = CToK(BoilingC)

	if valueC != AbsoluteZeroC {
		t.Fatalf("Expecting %g, got %g instead", AbsoluteZeroC, valueC)
	}

	if valueK != boilingK {
		t.Fatalf("Expecting %g, got %g instead", boilingK, valueK)
	}

	fmt.Println("TestKelvin done!")
}

func TestFahrenheit(t *testing.T) {
	var boilingF Fahrenheit = 212.0
	var boilingK Kelvin = 373.15

	var valueK Kelvin = FToK(boilingF)
	if valueK != boilingK {
		t.Fatalf("Expecting %g, got %g instead", boilingK, valueK)
	}

	var valueC Celsius = FToC(boilingF)
	if valueC != BoilingC {
		t.Fatalf("Expecting %g, got %g instead", BoilingC, valueC)
	}

	var valueF Fahrenheit = KToF(boilingK)
	if valueF != boilingF {
		t.Fatalf("Expecting %g, got %g instead", boilingF, valueF)
	}

	fmt.Println("TestFahrenheit done!")

}
