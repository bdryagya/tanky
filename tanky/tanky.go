package tanky

import "fmt"

// Tanky is struct
type Tanky struct {
	Capacity int
	Level    int
}

// Fill fills water
func (t *Tanky) Fill() {
	var f int
	fmt.Println("fill water in tanky: ")
	fmt.Scan(&f)
	if t.Capacity == 0 {
		t.Level += f
		if t.Level > 500 {
			fmt.Println("can't fill")
			t.Level -= f
		} else if t.Level <= 500 {
			fmt.Println("water level is: ", t.Level, "liter")
		} else {
			fmt.Println("can't fill")
		}
	} else {
		fmt.Println("tanky is full you can't fill into it")
	}
}

// Drain drains water
func (t *Tanky) Drain() {
	var d int
	fmt.Println("drain water from tanky:")
	fmt.Scan(&d)
	if t.Level != 0 {
		if d > 500 {
			fmt.Println("insufficient ")
		} else {
			t.Level = t.Level - d
			if t.Level < 0 {
				t.Level += d
				fmt.Println("can't drain")
			} else if t.Level <= 500 && t.Level >= 0 {
				fmt.Println("water level is: ", t.Level)
			} else {
				fmt.Println("can't drain")
				fmt.Println(t.Level)
			}
		}
	} else {
		fmt.Println("tanky is already empty,so no need to drain")
	}
}
