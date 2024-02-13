package filter

import (
	"strings"

	"github.com/elimity-com/scim"
	filter "github.com/scim2/filter-parser/v2"
)

func evaluateAttributeExpressionInt(attributeValue *int, comparisonValue *int, operator filter.CompareOperator) bool {
	switch operator {
	case filter.NE:
		return *attributeValue != *comparisonValue
	case filter.PR:
		return attributeValue != nil
	case filter.GT:
		return *attributeValue > *comparisonValue
	case filter.GE:
		return *attributeValue >= *comparisonValue
	case filter.LT:
		return *attributeValue < *comparisonValue
	case filter.LE:
		return *attributeValue <= *comparisonValue
	}

	return false
}

func evaluateAttributeExpressionStr(attributeValue *string, comparisonValue *string, operator filter.CompareOperator) bool {
	switch operator {
	case filter.CO:
		return strings.Contains(*attributeValue, *comparisonValue)
	case filter.EQ:
		return *attributeValue == *comparisonValue
	case filter.NE:
		return *attributeValue != *comparisonValue
	case filter.SW:
		return strings.HasPrefix(*attributeValue, *comparisonValue)
	case filter.EW:
		return strings.HasSuffix(*attributeValue, *comparisonValue)
	case filter.PR:
		return attributeValue != nil
	case filter.GT:
		return strings.Compare(*attributeValue, *comparisonValue) > 0
	case filter.GE:
		return strings.Compare(*attributeValue, *comparisonValue) >= 0
	case filter.LT:
		return strings.Compare(*attributeValue, *comparisonValue) < 0
	case filter.LE:
		return strings.Compare(*attributeValue, *comparisonValue) <= 0
	}

	return false
}

func evaluateAttributeExpressionBool(attributeValue *bool, comparisonValue *bool, operator filter.CompareOperator) bool {
	switch operator {
	case filter.EQ:
		return *attributeValue == *comparisonValue
	case filter.NE:
		return *attributeValue != *comparisonValue
	case filter.PR:
		return attributeValue != nil
	}

	return false
}

func extractSubAttributeValue[attributeType int | string | bool](subAttributePath string, parentAttributeValue map[string]interface{}) (*attributeType, bool) {
	splitSubAttributePath := strings.Split(subAttributePath, ".")
	currMap := parentAttributeValue

	for _, name := range splitSubAttributePath {
		currVal := currMap[name]

		// Attempt to cast value to the requested type
		valAsType, ok := currVal.(attributeType)
		if ok {
			return &valAsType, true
		}

		currMap = currVal.(map[string]interface{})
	}

	return nil, false
}

func extractAttributeValue[attributeType int | string | bool](attributePath filter.AttributePath, resource scim.Resource) (*attributeType, bool) {
	attributeName := KeyMap[strings.ToLower(attributePath.AttributeName)]

	if attributePath.SubAttribute == nil {
		return convertValue[attributeType](resource.Attributes[attributeName])
	}

	value, ok := extractSubAttributeValue[attributeType](
		KeyMap[strings.ToLower(attributeName+"."+*attributePath.SubAttribute)],
		resource.Attributes,
	)

	return value, ok
}

func evaluateAttributeExpression(expression *filter.AttributeExpression, resource scim.Resource) bool {
	// Attempt string
	valStr, ok := extractAttributeValue[string](expression.AttributePath, resource)
	if ok {
		comparisonValue, _ := convertValue[string](expression.CompareValue)
		return evaluateAttributeExpressionStr(valStr, comparisonValue, expression.Operator)
	}

	// Attempt int
	valInt, ok := extractAttributeValue[int](expression.AttributePath, resource)
	if ok {
		comparisonValue, _ := convertValue[int](expression.CompareValue)
		return evaluateAttributeExpressionInt(valInt, comparisonValue, expression.Operator)
	}

	// Attempt bool
	valBool, ok := extractAttributeValue[bool](expression.AttributePath, resource)
	if ok {
		comparisonValue, _ := convertValue[bool](expression.CompareValue)
		return evaluateAttributeExpressionBool(valBool, comparisonValue, expression.Operator)
	}

	return false
}
