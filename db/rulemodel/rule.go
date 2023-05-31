package rulemodel

import (
	"time"

	"github.com/go-bai/go-dnat/db"
)

type Rule struct {
	ID        int64 `db:"id"`
	Iface     string
	Port      int
	Dest      string
	CreatedAt time.Time `db:"created_at"`
}

func Insert(rule *Rule) error {
	rule.CreatedAt = time.Now()
	_, err := db.Ins.NamedExec("INSERT INTO rule (iface, port, dest, created_at) VALUES (:iface, :port, :dest, :created_at)", rule)
	return err
}

func Get(id int64) (*Rule, error) {
	rule := &Rule{}
	if err := db.Ins.Get(rule, "SELECT * FROM rule WHERE id = $1 LIMIT 1", id); err != nil {
		return nil, err
	}
	return rule, nil
}

func Delete(id int64) error {
	if _, err := db.Ins.Exec("DELETE FROM rule WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}

func List() ([]*Rule, error) {
	rules := make([]*Rule, 0)
	if err := db.Ins.Select(&rules, "SELECT * FROM rule"); err != nil {
		return nil, err
	}
	return rules, nil
}
