package xmodule

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/jsonbase"
	"github.com/Cyako/Cyako.go/statistics"
)

type Test struct {
}

func (t Test) PrintStat(c *cyako.Ctx) {
	stat := c.Middleware["Statistics"].(statistics.Statistics)
	if c.Error.Fatal == nil {
		stat.Get()
	}
}

func (t Test) TestJSONBaseLoad(c *cyako.Ctx) {
	j := c.Middleware["JSONBase"].(jsonbase.JSONBase)
	if c.Error.Fatal == nil {
		j.Load(c)
	}
}

func (t Test) TestJSONBaseSave(c *cyako.Ctx) {
	j := c.Middleware["JSONBase"].(jsonbase.JSONBase)
	if c.Error.Fatal == nil {
		j.Save(c)
	}
}

func init() {
	var m Test
	cyako.LoadModule(m)
}
