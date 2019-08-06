package model

import (
	"github.com/go-xorm/xorm"
	"strings"
)

type Criteria struct {
	ANDCondition []string
	ANDBind      []interface{}
	ORCondition  []string
	ORBind       []interface{}
	OrderBy      []string
	GroupBy      []string
	Limit        []int
}

func NewCriteria() *Criteria {
	return &Criteria{}
}

func (c Criteria) Apply(engine *xorm.Engine) *xorm.Session {

	session := engine.NewSession()
	if len(c.ANDCondition) > 0 {
		session = session.Where(strings.Join(c.ANDCondition, " AND "), c.ANDBind...)
	}

	if len(c.ORCondition) > 0 {
		session = session.Or(strings.Join(c.ORCondition, " AND "), c.ORBind)
	}
	session = session.OrderBy(strings.Join(c.OrderBy, ","))
	return session.GroupBy(strings.Join(c.GroupBy, ","))
}
