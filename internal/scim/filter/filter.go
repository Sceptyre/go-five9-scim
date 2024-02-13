package filter

import (
	"github.com/elimity-com/scim"
	filter "github.com/scim2/filter-parser/v2"
)

var KeyMap map[string]string

func Filter(expression filter.Expression, resources []scim.Resource) []scim.Resource {
	output := []scim.Resource{}

	firstResource := resources[0]

	KeyMap = extractMapKeysAsMap(firstResource.Attributes)

	for _, resource := range resources {
		if evaluateResource(expression, resource) {
			output = append(output, resource)
		}
	}

	return output
}
