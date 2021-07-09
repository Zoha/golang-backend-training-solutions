package cad

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type CAD struct {
	cents int64
}

func Cents(n int64) CAD {
	return CAD{
		cents: n,
	}
}

func ParseCAD(s string) (CAD, error) {
	// we will return this error if parsing failed
	err := errors.New("passed string is not a valid money string")

	// if string does not contain any `$` or `¢` chars
	// then money string is not valid and we should return error
	containsAnyDolorSymbol := strings.ContainsAny(s, "$¢")
	if !containsAnyDolorSymbol {
		return Cents(0), err
	}

	// create a regex to parse two group variables
	// 1 -> contains any `-` , 2 -> the numbers part with `.` and `,`  for example 123,4.41
	moneyStringRegexp := regexp.MustCompile(`(?:CAD)?(?:(\-)|[\$\-\s])*([0-9,.]+)`)
	moneyRegexResult := moneyStringRegexp.FindStringSubmatch(s)

	// if there is no match we should return error
	if len(moneyRegexResult) == 0 {
		return Cents(0), err
	}

	// define variables from matches
	centsString := moneyRegexResult[2]
	isNegative := moneyRegexResult[1] != ""
	isInDolorFormat := false

	// if does not contains any ¢ or . then it is in dolor format
	// and we should multiply it by 100 (convert it to cents)
	containsAnyDot := strings.Contains(centsString, ".")
	containsAnyCentSymbol := strings.Contains(s, "¢")
	if !containsAnyDot && !containsAnyCentSymbol {
		isInDolorFormat = true
	}

	// remove all `.`s and ','s
	// because we want to convert the string to pure integer
	centsString = strings.ReplaceAll(centsString, ".", "")
	centsString = strings.ReplaceAll(centsString, ",", "")

	if centsString == "" {
		return Cents(0), nil
	}

	// parse the string and create int64 (cents)
	centsWithoutSeparator := strings.ReplaceAll(centsString, ",", "")
	cents, err := strconv.Atoi(centsWithoutSeparator)
	if err != nil {
		return Cents(0), err
	}
	// if is in dolor format we should convert it to cents
	if isInDolorFormat {
		cents *= 100
	}
	finalCents := int64(cents)

	// check being negative
	if isNegative {
		finalCents = -1 * finalCents
	}

	return Cents(finalCents), nil
}

func (receiver CAD) Abs() CAD {
	return Cents(int64(math.Abs(float64(receiver.AsCents()))))
}

func (receiver CAD) Add(other CAD) CAD {
	return Cents(receiver.AsCents() + other.AsCents())
}

func (receiver CAD) AsCents() int64 {
	return receiver.cents
}

func (receiver CAD) CanonicalForm() (dollars int64, cents int64) {
	asDolors := float64(receiver.AsCents()) / 100
	dolors := int64(math.Trunc(asDolors))
	return dolors, receiver.AsCents() - dolors*100
}

func (receiver CAD) Mul(scalar int64) CAD {
	return Cents(receiver.AsCents() * scalar)
}

func (receiver CAD) Sub(other CAD) CAD {
	return Cents(receiver.AsCents() - other.AsCents())
}

func (receiver CAD) GoString() string {
	return fmt.Sprintf("Cents(%d)", receiver.AsCents())
}

func (receiver CAD) String() string {
	dolors, cents := receiver.CanonicalForm()
	return fmt.Sprintf("CAD$%d.%d", dolors, cents)
}
