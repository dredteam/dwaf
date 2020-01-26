package main

import "fmt"

type RuleMetadata struct {
	name        string
	description string
	cve         string
	cwe         string
	cpe         string
}

type RuleAction string

type RuleHeader struct {
	name  string
	value string
}

type RuleHeaders []RuleHeader

type Rule struct {
	method   string
	resource string
	headers  RuleHeaders
	action   RuleAction
	metadata RuleMetadata
	enabled  bool
}

type Rules []Rule

func main() {
	fmt.Printf("%v", r)
}
