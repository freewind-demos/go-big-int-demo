package main

import (
	"fmt"
	"math/big"
)

func main() {
	bi := big.NewInt(1)
	fmt.Println(bi.Bytes())

	bi = big.NewInt(256 + 1)
	fmt.Println(bi.Bytes())

	bi = big.NewInt(256*256 + 256 + 1)
	fmt.Println(bi.Bytes())

	bi = big.NewInt(256*256*256 + 256*256 + 256 + 1)
	fmt.Println(bi.Bytes())

	bi = big.NewInt(256*256*256*256 + 256*256*256 + 256*256 + 256 + 1)
	fmt.Println(bi.Bytes())

	fmt.Println("-------------------------------------------")

	fmt.Println(256*256*256 + 256*256 + 256 + 1)
	fmt.Println(big.NewInt(256*256*256 + 256*256 + 256 + 1).Bits())

	fmt.Println(256 * 256 * 256 * 256 * 256 * 256 * 256)
	fmt.Println(big.NewInt(256 * 256 * 256 * 256 * 256 * 256 * 256).Bits())

	// both: constant 18446744073709551616 overflows int64
	// fmt.Println(big.NewInt(256 * 256 * 256 * 256 * 256 * 256 * 256 * 256))
	// fmt.Println(big.NewInt(256 * 256 * 256 * 256 * 256 * 256 * 256 * 256).Bits())
}
