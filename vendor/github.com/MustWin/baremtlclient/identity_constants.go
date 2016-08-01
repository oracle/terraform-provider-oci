package baremtlsdk

type resourceName string

const (
	ResourceCreated  = "CREATED"
	ResourceCreating = "CREATING"

	identityServiceAPI        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion = "v1"

	// Header Keys
	headerOPCIdempotencyToken = "opc-idempotency-token"
	headerOPCNextPage         = "opc-next-page"
	headerIfMatch             = "if-match"
	headerOPCRequestID        = "opc-request-id"

	// URL Query Keys
	queryCompartmentID = "compartmentId"
	queryGroupID       = "groupId"
	queryPage          = "page"
	queryLimit         = "limit"
	queryUserID        = "userId"

	// Resources
	resourceAvailabilityDomains  resourceName = "availabilityDomains"
	resourceCompartments         resourceName = "compartments"
	resourceGroups               resourceName = "groups"
	resourcePolicies             resourceName = "policies"
	resourceUsers                resourceName = "users"
	resourceUserGroupMemberships resourceName = "userGroupMemberships"

	apiKeys    = "apiKeys"
	uiPassword = "uiPassword"
)
