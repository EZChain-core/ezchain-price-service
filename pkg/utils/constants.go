package utils


type ConditionRules struct {
	Keys []string `json:"keys"`
	Operators []string `json:"operators"`
}

const (
	FEATURE_ORGANIZATION = "iam_core"
	FEATURE_USER_NAME = "user_experiment"
)

var CONDITION_RULES = &ConditionRules{
	Keys: []string{
		"IpSource",
		"Tag",
		"Secret",
		"UserAgent",
		"AccountExpire",
	},
	Operators: []string{
		"StringEqual",
		"StringNotEqual",
		"StringRegex",
	},
}

var PROJECT_ROLES = map[string]string{
	"Maintainer": "e7bc3dacfc40434bba37ae22332d2d3b",
	"Reader": "fefa546657ad46b1867c92ba470951fd",
	"Member": "3dd5349e5d8b463d9c4632080789ed7e",
}

var FULL_PROJECT_ROLES = map[string]string{
	"Maintainer": "e7bc3dacfc40434bba37ae22332d2d3b",
	"Reader": "fefa546657ad46b1867c92ba470951fd",
	"Member": "3dd5349e5d8b463d9c4632080789ed7e",
	"Owner": "359f3d93a9df47c4aeed1583fab806c3",
}