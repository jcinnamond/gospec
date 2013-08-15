// Copyright 2013 John Cinnamond. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospec

import (
	"fmt"
	"errors"
)

type Spec struct {
	actual      interface{}
	inverted    bool
}

// An ExpectationError is returned when an expectation fails. It will
// yield a helpful description of the failed expectation.
type ExpectationError struct {
	spec *Spec
	expected interface{}
}

func (e *ExpectationError) Error() string {
	not := ""
	if e.spec.inverted {
		not = "not "
	}
	return fmt.Sprintf("expected %s%v, got %v", not, e.expected, e.spec.actual)
}

// Verify sets up a spec and stores the subject to test. Use this with
// an expectation matcher (e.g., `Equals') to check behaviour.
func Verify(actual interface{}) *Spec {
	return &Spec{actual: actual}
}

// DoesNot inverts the sense of an expectation.
func (spec *Spec) DoesNot() *Spec {
	spec.inverted = true
	return spec
}

// Equals is the simplest expectation matcher, returning an error
// if expected value does not match the subject passed to `Verify'.
func (spec *Spec) Equals(expected interface{}) (err error) {
	failed := spec.actual != expected
	if spec.inverted {
		failed = !failed
	}
	if failed {
		err = &ExpectationError{spec, expected}
	}
	return
}

// Alias for `Equals'.
func (spec *Spec) Equal(expected interface{}) error {
	return spec.Equals(expected)
}

// EqualsWithPrecision allows the comparison of two floats
// to a specified number of decimal places.
func (spec *Spec) EqualsWithPrecision(expected float64, precision int) error {
	v, ok := spec.actual.(float64)
	if ! ok {
		return errors.New((fmt.Sprintf("cannot compare a %T with precision", spec.actual)))
	}
	
	spec.actual = fmt.Sprintf("%0.*f", precision, float64(v))
	return spec.Equals(fmt.Sprintf("%0.*f", precision, expected))
}

// Alias for `EqualsWithPrecision'
func (spec *Spec) EqualWithPrecision(expected float64, precision int) error {
	return spec.EqualsWithPrecision(expected, precision)
}
