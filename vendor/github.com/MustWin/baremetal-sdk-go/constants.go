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
	queryAction                 = "action"
	queryAvailabilityDomain     = "availabilityDomain"
	queryCompartmentID          = "compartmentId"
	queryCpeID                  = "cpeId"
	queryDrgID                  = "drgId"
	queryGroupID                = "groupId"
	queryImageID                = "imageId"
	queryInstanceID             = "instanceId"
	queryLength                 = "length "
	queryLimit                  = "limit"
	queryOffset                 = "offset"
	queryOperatingSystem        = "operatingSystem"
	queryOperatingSystemVersion = "operatingSystemVersion"
	queryPage                   = "page"
	queryUserID                 = "userId"
	queryVcnID                  = "vcn"
	queryVnicID                 = "vnicId"
	queryVolumeID               = "volumeId"

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
	resourceDHCPOptions              resourceName = "dhcps"
	resourceDrgAttachments           resourceName = "drgAttachments"
	resourceDrgs                     resourceName = "drgs"
	resourceImages                   resourceName = "images"
	resourceInstanceConsoleHistories resourceName = "instanceConsoleHistories"
	resourceInstances                resourceName = "instances"
	resourceInternetGateways         resourceName = "internetGateways"
	resourceIPSecConnections         resourceName = "ipsecConnections"
	resourceRouteTables              resourceName = "routeTables"
	resourceShapes                   resourceName = "shapes"
	resourceSubnets                  resourceName = "subnets"
	resourceVirtualNetworks          resourceName = "vcns"
	resourceVnics                    resourceName = "vnics"
	resourceVnicAttachments          resourceName = "vnicAttachments"
	resourceVolumes                  resourceName = "volumes"
	resourceVolumeAttachments        resourceName = "volumeAttachments"
	resourceVolumeBackups            resourceName = "volumeBackups"

	apiKeys      = "apiKeys"
	uiPassword   = "uiPassword"
	deviceConfig = "deviceConfig"
	deviceStatus = "deviceStatus"
	dataURLPart  = "data"
)
