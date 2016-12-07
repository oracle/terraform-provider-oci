package baremetal

type resourceName string

type InstanceActions string
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

	SDKVersion = "20160918"

	identityServiceAPI        = "https://identity.us-az-phoenix-1.OracleIaaS.com"
	identityServiceAPIVersion = SDKVersion

	coreServiceAPI        = "https://core.us-az-phoenix-1.OracleIaaS.com"
	coreServiceAPIVersion = SDKVersion

	// Header Keys
	headerRetryToken     = "opc-retry-token"
	headerOPCNextPage    = "opc-next-page"
	headerIfMatch        = "If-Match"
	headerETag           = "ETag"
	headerOPCRequestID   = "opc-request-id"
	headerBytesRemaining = "opc-bytes-remaining"

	// Actions that can be applied to compute instances
	actionStart InstanceActions = "START"
	actionStop  InstanceActions = "STOP"
	actionReset InstanceActions = "RESET"

	// Network entity types for routing rules
	networkEntityVnic                      NetworkEntityType = "VNIC"
	networkEntityInternetGateway           NetworkEntityType = "INTERNET_GATEWAY"
	networkEntityDynamicallyRoutingGateway NetworkEntityType = "DYNAMICALLY_ROUTING_GATEWAY"

	// DB Resources
	resourceDBSystemShapes resourceName = "dbSystemShapes"

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
	resourceSecurityLists            resourceName = "securityLists"
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
