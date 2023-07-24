package iptables

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/coreos/go-iptables/iptables"
)

var comment = "go-dnat"

func AppendRule(iface string, port int, dest string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	if err := ipt.AppendUnique("nat", "PREROUTING",
		"-i", iface, "-p", "tcp", "--dport", strconv.Itoa(port), "-j", "DNAT", "--to-destination", dest,
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	if err := ipt.AppendUnique("nat", "PREROUTING",
		"-i", iface, "-p", "udp", "--dport", strconv.Itoa(port), "-j", "DNAT", "--to-destination", dest,
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	return nil
}

func AppendMasqueradeRule(iface string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	if err := ipt.AppendUnique("nat", "POSTROUTING",
		"-o", iface, "-p", "tcp", "-j", "MASQUERADE",
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	if err := ipt.AppendUnique("nat", "POSTROUTING",
		"-o", iface, "-p", "udp", "-j", "MASQUERADE",
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	return nil
}

func DeleteRule(iface string, port int, dest string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	if err := ipt.DeleteIfExists("nat", "PREROUTING",
		"-i", iface, "-p", "tcp", "--dport", strconv.Itoa(port), "-j", "DNAT", "--to-destination", dest,
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	if err := ipt.DeleteIfExists("nat", "PREROUTING",
		"-i", iface, "-p", "udp", "--dport", strconv.Itoa(port), "-j", "DNAT", "--to-destination", dest,
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	return nil
}

func DeleteMasqueradeRule(iface string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	if err := ipt.DeleteIfExists("nat", "POSTROUTING",
		"-o", iface, "-p", "tcp", "-j", "MASQUERADE",
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	if err := ipt.DeleteIfExists("nat", "POSTROUTING",
		"-o", iface, "-p", "udp", "-j", "MASQUERADE",
		"-m", "comment", "--comment", comment); err != nil {
		return err
	}

	return nil
}

func ListRules() ([]string, error) {
	ipt, err := iptables.New()
	if err != nil {
		return nil, err
	}

	rules, err := ipt.ListWithCounters("nat", "PREROUTING")
	if err != nil {
		return nil, err
	}

	dnatRuleList := make([]string, 0)
	for _, rule := range rules {
		if strings.Contains(rule, fmt.Sprintf("-m comment --comment %s", comment)) {
			dnatRuleList = append(dnatRuleList, rule)
		}
	}

	return dnatRuleList, nil
}
