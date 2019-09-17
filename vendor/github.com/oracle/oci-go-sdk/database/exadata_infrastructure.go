// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ExadataInfrastructure ExadataInfrastructure
type ExadataInfrastructure struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the Exadata infrastructure.
	LifecycleState ExadataInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Exadata infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

	// Size, in terabytes, of the DATA disk group.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The IP address for the first control plane server.
	CloudControlPlaneServer1 *string `mandatory:"false" json:"cloudControlPlaneServer1"`

	// The IP address for the second control plane server.
	CloudControlPlaneServer2 *string `mandatory:"false" json:"cloudControlPlaneServer2"`

	// The netmask for the control plane network.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The gateway for the control plane network.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The CIDR block for the Exadata administration network.
	AdminNetworkCIDR *string `mandatory:"false" json:"adminNetworkCIDR"`

	// The CIDR block for the Exadata InfiniBand interconnect.
	InfiniBandNetworkCIDR *string `mandatory:"false" json:"infiniBandNetworkCIDR"`

	// The corporate network proxy for access to the control plane network.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServer []string `mandatory:"false" json:"dnsServer"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServer []string `mandatory:"false" json:"ntpServer"`

	// The date and time the Exadata infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ExadataInfrastructureLifecycleStateEnum Enum with underlying type: string
type ExadataInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureLifecycleStateEnum
const (
	ExadataInfrastructureLifecycleStateCreating           ExadataInfrastructureLifecycleStateEnum = "CREATING"
	ExadataInfrastructureLifecycleStateRequiresActivation ExadataInfrastructureLifecycleStateEnum = "REQUIRES_ACTIVATION"
	ExadataInfrastructureLifecycleStateActivating         ExadataInfrastructureLifecycleStateEnum = "ACTIVATING"
	ExadataInfrastructureLifecycleStateActive             ExadataInfrastructureLifecycleStateEnum = "ACTIVE"
	ExadataInfrastructureLifecycleStateActivationFailed   ExadataInfrastructureLifecycleStateEnum = "ACTIVATION_FAILED"
	ExadataInfrastructureLifecycleStateFailed             ExadataInfrastructureLifecycleStateEnum = "FAILED"
	ExadataInfrastructureLifecycleStateUpdating           ExadataInfrastructureLifecycleStateEnum = "UPDATING"
	ExadataInfrastructureLifecycleStateDeleting           ExadataInfrastructureLifecycleStateEnum = "DELETING"
	ExadataInfrastructureLifecycleStateDeleted            ExadataInfrastructureLifecycleStateEnum = "DELETED"
	ExadataInfrastructureLifecycleStateOffline            ExadataInfrastructureLifecycleStateEnum = "OFFLINE"
)

var mappingExadataInfrastructureLifecycleState = map[string]ExadataInfrastructureLifecycleStateEnum{
	"CREATING":            ExadataInfrastructureLifecycleStateCreating,
	"REQUIRES_ACTIVATION": ExadataInfrastructureLifecycleStateRequiresActivation,
	"ACTIVATING":          ExadataInfrastructureLifecycleStateActivating,
	"ACTIVE":              ExadataInfrastructureLifecycleStateActive,
	"ACTIVATION_FAILED":   ExadataInfrastructureLifecycleStateActivationFailed,
	"FAILED":              ExadataInfrastructureLifecycleStateFailed,
	"UPDATING":            ExadataInfrastructureLifecycleStateUpdating,
	"DELETING":            ExadataInfrastructureLifecycleStateDeleting,
	"DELETED":             ExadataInfrastructureLifecycleStateDeleted,
	"OFFLINE":             ExadataInfrastructureLifecycleStateOffline,
}

// GetExadataInfrastructureLifecycleStateEnumValues Enumerates the set of values for ExadataInfrastructureLifecycleStateEnum
func GetExadataInfrastructureLifecycleStateEnumValues() []ExadataInfrastructureLifecycleStateEnum {
	values := make([]ExadataInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingExadataInfrastructureLifecycleState {
		values = append(values, v)
	}
	return values
}
