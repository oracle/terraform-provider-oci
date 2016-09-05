package baremetal

type resourceName string

type instanceActions string
type instanceStates string
type NetworkEntityType string

const (
	// Resource States
	ResourceActive             = "ACTIVE"
	ResourceAttached           = "ATTACHED"
	ResourceAttaching          = "ATTACHING"
	ResourceAvailable          = "AVAILABLE"
	ResourceCreated            = "CREATED"
	ResourceCreating           = "CREATING"
	ResourceCreatingImage      = "CREATING_IMAGE"
	ResourceDeleting           = "DELETING"
	ResourceDeleted            = "DELETED"
	ResourceDetached           = "DETACHED"
	ResourceDetaching          = "DETACHING"
	ResourceDisabled           = "DISABLED"
	ResourceDown               = "DOWN"
	ResourceDownForMaintenance = "DOWN_FOR_MAINTENANCE"
	ResourceFailed             = "FAILED"
	ResourceFaulty             = "FAULTY"
	ResourceGettingHistory     = "GETTING-HISTORY"
	ResourceInactive           = "INACTIVE"
	ResourceProvisioning       = "PROVISIONING"
	ResourceRequested          = "REQUESTED"
	ResourceRequestReceived    = "REQUEST_RECEIVED"
	ResourceRestoring          = "RESTORING"
	ResourceRunning            = "RUNNING"
	ResourceStarting           = "STARTING"
	ResourceStopped            = "STOPPED"
	ResourceStopping           = "STOPPING"
	ResourceSucceeded          = "SUCCEEDED"
	ResourceTerminated         = "TERMINATED"
	ResourceTerminating        = "TERMINATING"
	ResourceUp                 = "UP"

	identityServiceAPI        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion = "v1"

	coreServiceAPI        = "https://core.us-az-phoenix-1.OracleIaaS.com"
	coreServiceAPIVersion = "v1"

	// Header Keys
	headerRetryToken     = "opc-retry-token"
	headerOPCNextPage    = "opc-next-page"
	headerIfMatch        = "If-Match"
	headerETag           = "ETag"
	headerOPCRequestID   = "opc-request-id"
	headerBytesRemaining = "opc-bytes-remaining"

	// URL Query Keys
	queryAvailabilityDomain = "availabilityDomain"
	queryCompartmentID      = "compartmentId"
	queryGroupID            = "groupId"
	queryImageID            = "imageId"
	queryInstanceID         = "instanceId"
	queryLength             = "length "
	queryLimit              = "limit"
	queryOffset             = "offset"
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

	// Network entity types for routing rules
	networkEntityVnic                      NetworkEntityType = "VNIC"
	networkEntityInternetGateway           NetworkEntityType = "INTERNET_GATEWAY"
	networkEntityDynamicallyRoutingGateway NetworkEntityType = "DYNAMICALLY_ROUTING_GATEWAY"

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
	resourceVnics                    resourceName = "vnics"
	resourceVnicAttachments          resourceName = "vnicAttachments"
	resourceVirtualNetworks          resourceName = "vcns"
	resourceInstanceConsoleHistories resourceName = "instanceConsoleHistories"
	resourceVolumes                  resourceName = "volumes"
	resourceVolumeAttachments        resourceName = "volumeAttachments"
	resourceInstances                resourceName = "instances"
	resourceSubnets                  resourceName = "subnets"
	resourceIPSecConnections         resourceName = "ipsecConnections"
	resourceDrgs                     resourceName = "drgs"
	resourceDrgAttachments           resourceName = "drgAttachments"
	resourceInternetGateways         resourceName = "internetGateways"
	resourceRouteTables              resourceName = "routeTables"
	resourceVolumeBackups            resourceName = "volumeBackups"

	apiKeys      = "apiKeys"
	uiPassword   = "uiPassword"
	deviceConfig = "deviceConfig"
	deviceStatus = "deviceStatus"
	dataURLPart  = "data"
)
