package filter

import (
	"github.com/elimity-com/scim"
	filter "github.com/scim2/filter-parser/v2"
)

func evaluateNotExpression(expression *filter.NotExpression, resource scim.Resource) bool {
	return !evaluateResource(expression.Expression, resource)
}
