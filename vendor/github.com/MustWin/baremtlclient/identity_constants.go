package baremtlclient

type resourceName string

const (
	ResourceCreated  = "CREATED"
	ResourceCreating = "CREATING"

	identityServiceAPI                        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion                 = "v1"
	resourceAvailabilityDomains  resourceName = "availabilityDomains"
	resourceCompartments         resourceName = "compartments"
	resourceGroups               resourceName = "groups"
	resourcePolicies             resourceName = "policies"
	resourceUsers                resourceName = "users"
	resourceUserGroupMemberships resourceName = "userGroupMemberships"
)
