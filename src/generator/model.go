package generator

import "github.com/Velocidex/ordereddict"

type Check struct {
	Id        string  `json:"Id,omitempty"`
	Title     string  `json:"Title,omitempty"`
	Condition string  `json:"Condition,omitempty"`
	Rules     []*Test `json:"Rules,omitempty"`

	// Once we manually verify the check this will be set
	// true. Further SCA imports will use the existing rules.
	Verified bool `json:"Verified"`

	// Any VQL that should be run to remediate this check (bring into
	// compliance)
	Remediate string `json:"Remediate,omitempty"`
}

type Test struct {
	Type             string `json:"Type,omitempty"`
	Name             string `json:"Name,omitempty"`
	ColumnExpression string `json:"ColumnExpression,omitempty"`
	WhereExpression  string `json:"WhereExpression,omitempty"`
	OriginalTest     string `json:"OriginalTest,omitempty"`
	Error            string `json:"Error,omitempty"`

	Env *ordereddict.Dict `json:"Env,omitempty"`
}

type Rules struct {
	Source string   `json:"Source,omitempty"`
	Type   string   `json:"Type,omitempty"`
	Checks []*Check `json:"Checks,omitempty"`

	// Contains any VQL that should be added to the export section.
	Export string `json:"Export,omitempty"`
}
