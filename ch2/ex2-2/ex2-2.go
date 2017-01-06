// A general-purpose unit-conversion program

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl-ex/ch2/ex2-1/tempconv"
	"gopl-ex/ch2/ex2-2/lenconv"
	"gopl-ex/ch2/ex2-2/wtconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		ft := lenconv.Feet(t)
		m := lenconv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			ft, lenconv.FToM(ft), m, lenconv.MToF(m))

		lb := wtconv.Pounds(t)
		kg := wtconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n",
			lb, wtconv.LbToKg(lb), kg, wtconv.KgToLb(kg))

	}
}
