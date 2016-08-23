package baremetal

type resourceName string

type instanceActions string
type instanceStates string

const (
	// Identity States
	ResourceCreated  = "CREATED"
	ResourceCreating = "CREATING"

	// Core API States
	ResourceProvisioning       = "PROVISIONING"
	ResourceAvailable          = "AVAILABLE"
	ResourceTerminating        = "TERMINATING"
	ResourceTerminated         = "TERMINATED"
	ResourceAttaching          = "ATTACHING"
	ResourceAttached           = "ATTACHED"
	ResourceDetaching          = "DETACHING"
	ResourceDetached           = "DETACHED"
	ResourceRunning            = "RUNNING"
	ResourceStarting           = "STARTING"
	ResourceStopping           = "STOPPING"
	ResourceStopped            = "STOPPED"
	ResourceUp                 = "UP"
	ResourceDown               = "DOWN"
	ResourceDownForMaintenance = "DOWN_FOR_MAINTENANCE"

	identityServiceAPI        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion = "v1"

	coreServiceAPI        = "https://core.us-az-phoenix-1.OracleIaaS.com"
	coreServiceAPIVersion = "v1"

	// Header Keys
	headerOPCIdempotencyToken = "opc-idempotency-token"
	headerOPCNextPage         = "opc-next-page"
	headerIfMatch             = "If-Match"
	headerETag                = "ETag"
	headerOPCRequestID        = "opc-request-id"

	// URL Query Keys
	queryAvailabilityDomain = "availabilityDomain"
	queryCompartmentID      = "compartmentId"
	queryGroupID            = "groupId"
	queryImageID            = "imageId"
	queryInstanceID         = "instanceId"
	queryLimit              = "limit"
	queryPage               = "page"
	queryUserID             = "userId"
	queryVnicID             = "vnicId"
	queryAction             = "action"
	queryVcnID              = "vcn"
	queryDrgID              = "drgId"
	queryCpeID              = "cpeId"

	// Actions that can be applied to compute instances
	actionStart instanceActions = "START"
	actionStop  instanceActions = "STOP"
	actionReset instanceActions = "RESET"

	// Identity Resources
	resourceAvailabilityDomains  resourceName = "availabilityDomains"
	resourceCompartments         resourceName = "compartments"
	resourceGroups               resourceName = "groups"
	resourcePolicies             resourceName = "policies"
	resourceUsers                resourceName = "users"
	resourceUserGroupMemberships resourceName = "userGroupMemberships"

	// Core Resources
	resourceCustomerPremiseEquipment resourceName = "cpes"
	resourceShapes                   resourceName = "shapes"
	resourceVnicAttachments          resourceName = "vnicAttachments"
	resourceVirtualNetworks          resourceName = "vcns"
	resourceVolumes                  resourceName = "volumes"
	resourceVolumeAttachments        resourceName = "volumeAttachments"
	resourceInstances                resourceName = "instances"
	resourceSubnets                  resourceName = "subnets"
	resourceIPSecConnections         resourceName = "ipsecConnections"
	resourceDrgs                     resourceName = "drgs"

	apiKeys      = "apiKeys"
	uiPassword   = "uiPassword"
	deviceConfig = "deviceConfig"
	deviceStatus = "deviceStatus"
)
