package iptables

import (
	"fmt"
	"testing"
)

func TestAppendRule(t *testing.T) {
	if err := AppendRule("eth0", 9001, "192.168.3.27:9001"); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteRule(t *testing.T) {
	if err := DeleteRule("eth0", 9001, "192.168.3.27:9001"); err != nil {
		t.Fatal(err)
	}
}

func TestListRules(t *testing.T) {
	dnatRuleList, err := ListRules()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dnatRuleList)
}
