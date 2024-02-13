package filter

import (
	"github.com/elimity-com/scim"
	filter "github.com/scim2/filter-parser/v2"
)

func evaluateResource(expression filter.Expression, resource scim.Resource) bool {
	switch v := expression.(type) {
	case *filter.LogicalExpression:
		return evaluateLogicalExpression(v, resource)
	case *filter.ValuePath:
		return evaluateValuePath(v, resource)
	case *filter.AttributeExpression:
		return evaluateAttributeExpression(v, resource)
	case *filter.NotExpression:
		return evaluateNotExpression(v, resource)
	}
	return false
}
