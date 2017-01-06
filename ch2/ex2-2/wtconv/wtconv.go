package wtconv

import "fmt"

type Pounds float64
type Kilograms float64

func (lb Pounds) String() string    { return fmt.Sprintf("%g lb", lb) }
func (kg Kilograms) String() string { return fmt.Sprintf("%g kg", kg) }

func LbToKg(lb Pounds) Kilograms { return Kilograms(lb * 0.45359237) }

func KgToLb(kg Kilograms) Pounds { return Pounds(kg / 0.45359237) }
