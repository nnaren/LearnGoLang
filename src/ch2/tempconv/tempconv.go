// tempconv 包负责摄氏温度与华氏温度的转换
package tempconv

import "fmt"

type Celsius  float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreeZingC	Celsius = 0
	BoilingC	Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gC",c) }

func (c Fahrenheit) String() string { return fmt.Sprintf("%gF",c) }
