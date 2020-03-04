package acmutil_test

import (
	"fmt"
	"math"

	"github.com/Actinium-project/acmutil"
)

func ExampleAmount() {

	a := acmutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = acmutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = acmutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 ACM
	// 100,000,000 Satoshis: 1 ACM
	// 100,000 Satoshis: 0.001 ACM
}

func ExampleNewAmount() {
	amountOne, err := acmutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := acmutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := acmutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := acmutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 ACM
	// 0.01234567 ACM
	// 0 ACM
	// invalid litecoin amount
}

func ExampleAmount_unitConversions() {
	amount := acmutil.Amount(44433322211100)

	fmt.Println("Satoshi to kBTC:", amount.Format(acmutil.AmountKiloBTC))
	fmt.Println("Satoshi to ACM:", amount)
	fmt.Println("Satoshi to MilliBTC:", amount.Format(acmutil.AmountMilliBTC))
	fmt.Println("Satoshi to MicroBTC:", amount.Format(acmutil.AmountMicroBTC))
	fmt.Println("Satoshi to Satoshi:", amount.Format(acmutil.AmountSatoshi))

	// Output:
	// Satoshi to kBTC: 444.333222111 kBTC
	// Satoshi to ACM: 444333.222111 ACM
	// Satoshi to MilliBTC: 444333222.111 mBTC
	// Satoshi to MicroBTC: 444333222111 μBTC
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
