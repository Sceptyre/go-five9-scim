package scim

import (
	"github.com/Sceptyre/go-five9-scim/internal/scim/resourcetypes"

	"github.com/elimity-com/scim"
)

var ScimServer = &scim.Server{
	ResourceTypes: []scim.ResourceType{resourcetypes.UserResourceType},
}
