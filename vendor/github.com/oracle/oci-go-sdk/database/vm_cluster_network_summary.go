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

// VmClusterNetworkSummary Details of the VM cluster network.
type VmClusterNetworkSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// The user-friendly name for the VM cluster network. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The SCAN details.
	Scans []ScanDetails `mandatory:"false" json:"scans"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns []string `mandatory:"false" json:"dns"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntp []string `mandatory:"false" json:"ntp"`

	// Details of the client and backup networks.
	VmNetworks []VmNetworkDetails `mandatory:"false" json:"vmNetworks"`

	// The current state of the VM cluster network.
	LifecycleState VmClusterNetworkSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time when the VM cluster network was created.
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

func (m VmClusterNetworkSummary) String() string {
	return common.PointerString(m)
}

// VmClusterNetworkSummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterNetworkSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterNetworkSummaryLifecycleStateEnum
const (
	VmClusterNetworkSummaryLifecycleStateCreating           VmClusterNetworkSummaryLifecycleStateEnum = "CREATING"
	VmClusterNetworkSummaryLifecycleStateRequiresValidation VmClusterNetworkSummaryLifecycleStateEnum = "REQUIRES_VALIDATION"
	VmClusterNetworkSummaryLifecycleStateValidating         VmClusterNetworkSummaryLifecycleStateEnum = "VALIDATING"
	VmClusterNetworkSummaryLifecycleStateValidated          VmClusterNetworkSummaryLifecycleStateEnum = "VALIDATED"
	VmClusterNetworkSummaryLifecycleStateValidationFailed   VmClusterNetworkSummaryLifecycleStateEnum = "VALIDATION_FAILED"
	VmClusterNetworkSummaryLifecycleStateUpdating           VmClusterNetworkSummaryLifecycleStateEnum = "UPDATING"
	VmClusterNetworkSummaryLifecycleStateAllocated          VmClusterNetworkSummaryLifecycleStateEnum = "ALLOCATED"
	VmClusterNetworkSummaryLifecycleStateTerminating        VmClusterNetworkSummaryLifecycleStateEnum = "TERMINATING"
	VmClusterNetworkSummaryLifecycleStateTerminated         VmClusterNetworkSummaryLifecycleStateEnum = "TERMINATED"
	VmClusterNetworkSummaryLifecycleStateFailed             VmClusterNetworkSummaryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterNetworkSummaryLifecycleState = map[string]VmClusterNetworkSummaryLifecycleStateEnum{
	"CREATING":            VmClusterNetworkSummaryLifecycleStateCreating,
	"REQUIRES_VALIDATION": VmClusterNetworkSummaryLifecycleStateRequiresValidation,
	"VALIDATING":          VmClusterNetworkSummaryLifecycleStateValidating,
	"VALIDATED":           VmClusterNetworkSummaryLifecycleStateValidated,
	"VALIDATION_FAILED":   VmClusterNetworkSummaryLifecycleStateValidationFailed,
	"UPDATING":            VmClusterNetworkSummaryLifecycleStateUpdating,
	"ALLOCATED":           VmClusterNetworkSummaryLifecycleStateAllocated,
	"TERMINATING":         VmClusterNetworkSummaryLifecycleStateTerminating,
	"TERMINATED":          VmClusterNetworkSummaryLifecycleStateTerminated,
	"FAILED":              VmClusterNetworkSummaryLifecycleStateFailed,
}

// GetVmClusterNetworkSummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterNetworkSummaryLifecycleStateEnum
func GetVmClusterNetworkSummaryLifecycleStateEnumValues() []VmClusterNetworkSummaryLifecycleStateEnum {
	values := make([]VmClusterNetworkSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterNetworkSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
