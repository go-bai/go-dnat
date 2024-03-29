package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-bai/go-dnat/db/rulemodel"
	"github.com/rodaine/table"
)

func printTable(rules []*rulemodel.Rule) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Iface", "Port", "Dest", "Comment", "CreatedAt")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	fmt.Println()
	for _, rule := range rules {
		tbl.AddRow(rule.ID, rule.Iface, rule.Port, rule.Dest, rule.Comment, rule.CreatedAt.Format(time.RFC3339))
	}
	tbl.Print()
	fmt.Println()
}
