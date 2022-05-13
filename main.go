package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

const places = 2

func inc(decimal *[places]int) {
	for i := 0; i < len(decimal); i++ {
		decimal[i]++
		if decimal[i] < 10 {
			break
		}
		// to make loop exit tests efficient
		// allow most significant place to go out of range
		// instead of allocating an extra place
		if i < len(decimal)-1 {
			decimal[i] -= 10
		}
	}
}

func hasDecimal(decimal *[places]int) int {
	has := 0
	// omit leading zeros
	// 2 and 10 are not friend numbers
	// it's possible that it would be faster to use a slice
	// instead of the current fixed length array
	limit := len(decimal)
	for {
		i := limit - 1
		if i <= 0 || decimal[i] != 0 {
			break
		}
		limit = i
	}
	for i := 0; i < limit; i++ {
		has |= 1 << decimal[i]
	}
	return has
}

func decimal() int {
	// lsb decimal digits
	var pDecimal, qDecimal [places]int
	pDecimal[0] = 1
	count := 0
	for ; pDecimal[len(pDecimal)-1] < 10; inc(&pDecimal) {
		pHas := hasDecimal(&pDecimal)
		copy(qDecimal[:], pDecimal[:])
		inc(&qDecimal)
		for ; qDecimal[len(qDecimal)-1] < 10; inc(&qDecimal) {
			qHas := hasDecimal(&qDecimal)
			if pHas&qHas != 0 {
				count++
			}
		}
	}
	return count
}

func hasBinary(p int) int {
	has := 0
	// i tried passing in a *[places]byte
	// and passing that to AppendInt.
	// surprisingly, that was much slower.
	// also curious, benchmarking claims this isn't allocating
	ascii := strconv.Itoa(p)
	for i := 0; i < len(ascii); i++ {
		has |= 1 << (ascii[i] - '0')
	}
	return has
}

func binary() int {
	count := 0
	limit := int(math.Pow(10, places))
	for p := 1; p < limit; p++ {
		pHas := hasBinary(p)
		for q := p + 1; q < limit; q++ {
			qHas := hasBinary(q)
			if pHas&qHas != 0 {
				count++
			}
		}
	}
	return count
}

func mainError() (err error) {
	fmt.Println("decimal", decimal())
	// fmt.Println("binary", binary())
	return nil
}

func mainCode() int {
	err := mainError()
	if err == nil {
		return 0
	}
	fmt.Fprintf(os.Stderr, "%v: Error: %v\n", filepath.Base(os.Args[0]), err)
	return 1
}

func main() {
	os.Exit(mainCode())
}
