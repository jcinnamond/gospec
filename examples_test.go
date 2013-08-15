package gospec

import (
	"fmt"
)

func Example() {
	// This should return an error as 1 != 2
 	if err := Verify(1).Equals(2); err != nil {
		fmt.Println(err)
	}

	// This should return nil
 	if err := Verify(1).Equals(1); err != nil {
		fmt.Println(err)
	}
	// Output: expected 2, got 1
}

func ExampleVerify() {
	a := 4 / 2
 	if err := Verify(a).Equals(2); err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSpec_Equals() {
 	if err := Verify(1).Equals(2); err != nil {
		fmt.Println(err)
	}

 	if err := Verify(1).Equals(2.3); err != nil {
		fmt.Println(err)
	}

 	if err := Verify("baz").Equals("bar"); err != nil {
		fmt.Println(err)
	}
	// Output:
	// expected 2, got 1
	// expected 2.3, got 1
	// expected bar, got baz
}

func ExampleSpec_DoesNot() {
	// This should return nil
	if err := Verify(1).DoesNot().Equal(2); err != nil {
		fmt.Println(err)
	}

	// This should return an error as 1 = 1
 	if err := Verify(1).DoesNot().Equal(1); err != nil {
		fmt.Println(err)
	}
	// Output: expected not 1, got 1
}

func ExampleSpec_EqualsWithPrecision() {
	// This should return an error as 1.123 != 1.124
	if err := Verify(1.12311).EqualsWithPrecision(1.12411, 3); err != nil {
		fmt.Println(err)
	}

	// This should return an error as 2.0 != 3.0
	if err := Verify(2.0).EqualsWithPrecision(3.0, 2); err != nil {
		fmt.Println(err)
	}

	// This should return nil
 	if err := Verify(1.12311).EqualsWithPrecision(1.12411, 2); err != nil {
		fmt.Println(err)
	}
	// Output:
	// expected 1.124, got 1.123
	// expected 3.00, got 2.00
}

func ExampleSpec_EqualsWithPrecision_inverted() {
	// This should return an error as 1.123 = 1.123
	if err := Verify(1.1236).DoesNot().EqualWithPrecision(1.124, 3); err != nil {
		fmt.Println(err)
	}

	// This should return nil
	if err := Verify(1.1231).DoesNot().EqualWithPrecision(1.124, 3); err != nil {
		fmt.Println(err)
	}
	// Output: expected not 1.124, got 1.124
}
