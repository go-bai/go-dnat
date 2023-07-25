package cmd

import (
	"time"

	"github.com/go-bai/go-dnat/db"
	"github.com/go-bai/go-dnat/db/rulemodel"
	"github.com/go-bai/go-dnat/iptables"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli/v2"
)

func Append(cCtx *cli.Context) error {
	iface, port, dest, comment := cCtx.String("i"), cCtx.Int("p"), cCtx.String("d"), cCtx.String("m")

	rule := &rulemodel.Rule{
		Iface:     iface,
		Port:      port,
		Dest:      dest,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	err := db.Tx(func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec("INSERT INTO rule (iface, port, dest, comment, created_at) VALUES (:iface, :port, :dest, :comment, :created_at)", rule)
		if err != nil {
			return err
		}

		if err := iptables.AppendRule(iface, port, dest); err != nil {
			return err
		}
		return nil
	})

	return err
}

func Delete(cCtx *cli.Context) error {
	id := cCtx.Int64("id")
	rule, err := rulemodel.Get(id)
	if err != nil {
		return err
	}

	err = db.Tx(func(tx *sqlx.Tx) error {
		if _, err := tx.Exec("DELETE FROM rule WHERE id = $1", id); err != nil {
			return err
		}

		if err := iptables.DeleteRule(rule.Iface, rule.Port, rule.Dest); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func List(cCtx *cli.Context) error {
	rules, err := rulemodel.List()
	if err != nil {
		return err
	}

	printTable(rules)
	return nil
}

func Get(cCtx *cli.Context) error {
	id := cCtx.Int64("id")
	rule, err := rulemodel.Get(id)
	if err != nil {
		return err
	}

	printTable([]*rulemodel.Rule{rule})
	return nil
}

func Sync(cCtx *cli.Context) error {
	rules, err := rulemodel.List()
	if err != nil {
		return err
	}

	for _, rule := range rules {
		if err := iptables.AppendRule(rule.Iface, rule.Port, rule.Dest); err != nil {
			return err
		}
	}
	return nil
}

func Masquerade(cCtx *cli.Context) error {
	append, delete, iface := cCtx.Bool("append"), cCtx.Bool("delete"), cCtx.String("o")
	switch {
	case append:
		if err := iptables.AppendMasqueradeRule(iface); err != nil {
			return err
		}
	case delete:
		if err := iptables.DeleteMasqueradeRule(iface); err != nil {
			return err
		}
	}
	return nil
}
