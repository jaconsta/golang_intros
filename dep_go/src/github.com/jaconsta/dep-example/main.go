package main

// Source: https://medium.freecodecamp.org/an-intro-to-dep-how-to-manage-your-golang-project-dependencies-7b07d84e7ba5

import humanize "github.com/dustin/go-humanize"
import accounting "github.com/leekchan/accounting"
import "math/big"
import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82862982)) // 83 MB
	fmt.Printf("You're my %s best friend.\n", humanize.Ordinal(193))
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491))

	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213213))
	fmt.Println(ac.FormatMoney(12345678))
	fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))
	fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))
	fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123)))
}
