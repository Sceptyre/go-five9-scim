package filter

import (
	"github.com/elimity-com/scim"
	filter "github.com/scim2/filter-parser/v2"
)

func evaluateLogicalExpression(expression *filter.LogicalExpression, resource scim.Resource) bool {
	switch expression.Operator {
	case filter.AND:
		return evaluateResource(expression.Left, resource) && evaluateResource(expression.Right, resource)
	case filter.OR:
		return evaluateResource(expression.Left, resource) || evaluateResource(expression.Right, resource)
	}

	return false
}
