// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ExadataInfrastructureSummary Details of the Exadata infrastructure.
type ExadataInfrastructureSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the Exadata infrastructure.
	LifecycleState ExadataInfrastructureSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Exadata infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

	// The total number of CPU cores available.
	MaxCpuCount *int `mandatory:"false" json:"maxCpuCount"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The total memory available in GBs.
	MaxMemoryInGBs *int `mandatory:"false" json:"maxMemoryInGBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The total local node storage available in GBs.
	MaxDbNodeStorageInGBs *int `mandatory:"false" json:"maxDbNodeStorageInGBs"`

	// Size, in terabytes, of the DATA disk group.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The total available DATA disk group size.
	MaxDataStorageInTBs *float64 `mandatory:"false" json:"maxDataStorageInTBs"`

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

func (m ExadataInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ExadataInfrastructureSummaryLifecycleStateEnum Enum with underlying type: string
type ExadataInfrastructureSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureSummaryLifecycleStateEnum
const (
	ExadataInfrastructureSummaryLifecycleStateCreating           ExadataInfrastructureSummaryLifecycleStateEnum = "CREATING"
	ExadataInfrastructureSummaryLifecycleStateRequiresActivation ExadataInfrastructureSummaryLifecycleStateEnum = "REQUIRES_ACTIVATION"
	ExadataInfrastructureSummaryLifecycleStateActivating         ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVATING"
	ExadataInfrastructureSummaryLifecycleStateActive             ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVE"
	ExadataInfrastructureSummaryLifecycleStateActivationFailed   ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVATION_FAILED"
	ExadataInfrastructureSummaryLifecycleStateFailed             ExadataInfrastructureSummaryLifecycleStateEnum = "FAILED"
	ExadataInfrastructureSummaryLifecycleStateUpdating           ExadataInfrastructureSummaryLifecycleStateEnum = "UPDATING"
	ExadataInfrastructureSummaryLifecycleStateDeleting           ExadataInfrastructureSummaryLifecycleStateEnum = "DELETING"
	ExadataInfrastructureSummaryLifecycleStateDeleted            ExadataInfrastructureSummaryLifecycleStateEnum = "DELETED"
	ExadataInfrastructureSummaryLifecycleStateDisconnected       ExadataInfrastructureSummaryLifecycleStateEnum = "DISCONNECTED"
)

var mappingExadataInfrastructureSummaryLifecycleState = map[string]ExadataInfrastructureSummaryLifecycleStateEnum{
	"CREATING":            ExadataInfrastructureSummaryLifecycleStateCreating,
	"REQUIRES_ACTIVATION": ExadataInfrastructureSummaryLifecycleStateRequiresActivation,
	"ACTIVATING":          ExadataInfrastructureSummaryLifecycleStateActivating,
	"ACTIVE":              ExadataInfrastructureSummaryLifecycleStateActive,
	"ACTIVATION_FAILED":   ExadataInfrastructureSummaryLifecycleStateActivationFailed,
	"FAILED":              ExadataInfrastructureSummaryLifecycleStateFailed,
	"UPDATING":            ExadataInfrastructureSummaryLifecycleStateUpdating,
	"DELETING":            ExadataInfrastructureSummaryLifecycleStateDeleting,
	"DELETED":             ExadataInfrastructureSummaryLifecycleStateDeleted,
	"DISCONNECTED":        ExadataInfrastructureSummaryLifecycleStateDisconnected,
}

// GetExadataInfrastructureSummaryLifecycleStateEnumValues Enumerates the set of values for ExadataInfrastructureSummaryLifecycleStateEnum
func GetExadataInfrastructureSummaryLifecycleStateEnumValues() []ExadataInfrastructureSummaryLifecycleStateEnum {
	values := make([]ExadataInfrastructureSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExadataInfrastructureSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
