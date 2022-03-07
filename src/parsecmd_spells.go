package nox

import (
	"context"

	"nox/v1/common/console"
	"nox/v1/common/things"
)

func init() {
	noxCmdList.Register(&console.Command{
		Token:  "spells",
		HelpID: "listspellshelp",
		Flags:  console.ClientServer,
		Func: func(ctx context.Context, c *console.Console, tokens []string) bool {
			const (
				allow    = "ok"
				disallow = "X"
				invalid  = "INVALID"
			)
			for i, sp := range noxServer.SpellDefs() {
				mana := sp.Def.ManaCost
				title := sp.Title
				ind := things.SpellID(i + 1)
				id := ind.String()
				astr := disallow
				if !sp.IsValid() {
					astr = invalid
				} else if sp.Enabled {
					astr = allow
				}
				c.Printf(console.ColorRed, "%3d\t%-40.40s\t%-40.40q\tcost: %-3d\t%-8s", ind, id, title, mana, astr)
			}
			return true
		},
	})
}
