package main

import "regexp"

import "fmt"

func main() {
	reZip := regexp.MustCompile(`[0-9]{5}(?:-[0-9]{4})`)
	reSSN := regexp.MustCompile(`[0-9]{3}-([0-9]{2})-[0-9]{4}`)
	rePhone := regexp.MustCompile(`\(([0-9]{3}\))\s([0-9]{3}\-[0-9]{4})`)

	fmt.Printf("%s\n", reZip.FindString("Can it find the 12345-6789 zip code?"))
	fmt.Printf("%s\n", reSSN.FindString("Can it find the 123-45-6789 SSN?"))
	fmt.Printf("%s\n", rePhone.FindString("Can it find the (123) 456-7890 phone number?"))
}
