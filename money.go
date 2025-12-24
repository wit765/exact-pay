package main

import (
	"errors"
	"fmt"
)

var (
	errCanNotPay      = errors.New("can not pay")
	errNotEnoughMoney = errors.New("not enough money")
)

// FindSolution determines whether it is possible to pay the exact amount using the given number of 10-yuan, 5-yuan, and 2-yuan bills in one transaction (no change required).
//
// Parameters:
//
//	a: number of 10-yuan bills
//	b: number of 5-yuan bills
//	c: number of 2-yuan bills
//	amount: total amount to pay
//
// Returns:
//
//	x: number of 10-yuan bills used
//	y: number of 5-yuan bills used
//	z: number of 2-yuan bills used
//	err: error if payment is not possible, otherwise nil
//
// Algorithm:
//  1. Use a 5-yuan bill first if the amount is odd (if possible), otherwise only need to handle even amounts.
//  2. Use as many 10-yuan bills as possible, then 5-yuan bills, finally 2-yuan bills.
//  3. If any step cannot satisfy the requirement, return an error.
//
// Examples:
//
//	FindSolution(10, 10, 10, 27) => 2, 1, 1, nil
//	FindSolution(10, 0, 10, 16)  => 1, 0, 3, nil
//	FindSolution(1, 10, 1, 62)   => 0, 0, 0, err
func FindSolution(a, b, c, amount uint) (uint, uint, uint, error) {
	var (
		x uint = 0 // number of 10-yuan bills used
		y uint = 0 // number of 5-yuan bills used
		z uint = 0 // number of 2-yuan bills used
	)

	// Handle the case of paying 0 yuan
	if amount == 0 {
		return 0, 0, 0, nil
	}

	// Check if total available money is sufficient
	if a*10+b*5+c*2 < amount {
		return 0, 0, 0, errNotEnoughMoney
	}

	// If the amount is odd, use one 5-yuan bill if possible
	if amount%2 != 0 {
		if amount < 5 || b == 0 {
			return 0, 0, 0, errCanNotPay
		}
		amount -= 5
		b -= 1
		y += 1
	}

	// Use as many 10-yuan bills as possible
	if amount >= 10 {
		need := min(a, amount/10)
		amount -= need * 10
		x += need
		a -= need
	}

	// Use as many 5-yuan bills as possible (in pairs, if amount >= 10)
	if amount >= 10 {
		need := min(b/2, amount/10)
		amount -= need * 10
		y += need * 2
		b -= need * 2
	}

	// Use 2-yuan bills for the remaining amount
	need := amount / 2
	if c >= need {
		amount -= need * 2
		z += need
		c -= need
	}

	// If there is still remaining amount, payment is not possible
	if amount > 0 {
		return 0, 0, 0, errCanNotPay
	}

	return x, y, z, nil
}

func main() {
	cases := []struct {
		a, b, c, amount uint
	}{
		{10, 10, 10, 27}, // 2 x 10 + 1 x 5 + 1 x 2
		{10, 0, 10, 16},  // No 5-yuan, even amount possible
		{1, 10, 1, 62},   // No 5-yuan, odd amount impossible
	}
	for _, cs := range cases {
		x, y, z, err := FindSolution(cs.a, cs.b, cs.c, cs.amount)
		fmt.Printf("Input: a=%d, b=%d, c=%d, amount=%d\n", cs.a, cs.b, cs.c, cs.amount)
		if err != nil {
			fmt.Printf("Output: Cannot pay the exact amount, err=%v\n", err)
		} else {
			fmt.Printf("Output: %d x 10-yuan + %d x 5-yuan + %d x 2-yuan = %d\n", x, y, z, cs.amount)
		}
		fmt.Println()
	}
}
