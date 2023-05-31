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
	iface, port, dest := cCtx.String("i"), cCtx.Int("p"), cCtx.String("d")

	rule := &rulemodel.Rule{
		Iface:     iface,
		Port:      port,
		Dest:      dest,
		CreatedAt: time.Now(),
	}

	err := db.Tx(func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec("INSERT INTO rule (iface, port, dest, created_at) VALUES (:iface, :port, :dest, :created_at)", rule)
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
	rule, err := rulemodel.Get(cCtx.Int64("id"))
	if err != nil {
		return err
	}

	printTable([]*rulemodel.Rule{rule})
	return nil
}
