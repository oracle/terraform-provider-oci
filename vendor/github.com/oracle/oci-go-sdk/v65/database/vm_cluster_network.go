// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmClusterNetwork The VM cluster network.
type VmClusterNetwork struct {

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
	// CREATING - The resource is being created
	// REQUIRES_VALIDATION - The resource is created and may not be usable until it is validated.
	// VALIDATING - The resource is being validated and not available to use.
	// VALIDATED - The resource is validated and is available for consumption by VM cluster.
	// VALIDATION_FAILED - The resource validation has failed and might require user input to be corrected.
	// UPDATING - The resource is being updated and not available to use.
	// ALLOCATED - The resource is is currently being used by VM cluster.
	// TERMINATING - The resource is being deleted and not available to use.
	// TERMINATED - The resource is deleted and unavailable.
	// FAILED - The resource is in a failed state due to validation or other errors.
	// NEEDS_ATTENTION - The resource is in needs attention state as some of it's child nodes are not validated
	//                   and unusable by VM cluster.
	LifecycleState VmClusterNetworkLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m VmClusterNetwork) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterNetwork) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmClusterNetworkLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterNetworkLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterNetworkLifecycleStateEnum Enum with underlying type: string
type VmClusterNetworkLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterNetworkLifecycleStateEnum
const (
	VmClusterNetworkLifecycleStateCreating           VmClusterNetworkLifecycleStateEnum = "CREATING"
	VmClusterNetworkLifecycleStateRequiresValidation VmClusterNetworkLifecycleStateEnum = "REQUIRES_VALIDATION"
	VmClusterNetworkLifecycleStateValidating         VmClusterNetworkLifecycleStateEnum = "VALIDATING"
	VmClusterNetworkLifecycleStateValidated          VmClusterNetworkLifecycleStateEnum = "VALIDATED"
	VmClusterNetworkLifecycleStateValidationFailed   VmClusterNetworkLifecycleStateEnum = "VALIDATION_FAILED"
	VmClusterNetworkLifecycleStateUpdating           VmClusterNetworkLifecycleStateEnum = "UPDATING"
	VmClusterNetworkLifecycleStateAllocated          VmClusterNetworkLifecycleStateEnum = "ALLOCATED"
	VmClusterNetworkLifecycleStateTerminating        VmClusterNetworkLifecycleStateEnum = "TERMINATING"
	VmClusterNetworkLifecycleStateTerminated         VmClusterNetworkLifecycleStateEnum = "TERMINATED"
	VmClusterNetworkLifecycleStateFailed             VmClusterNetworkLifecycleStateEnum = "FAILED"
	VmClusterNetworkLifecycleStateNeedsAttention     VmClusterNetworkLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingVmClusterNetworkLifecycleStateEnum = map[string]VmClusterNetworkLifecycleStateEnum{
	"CREATING":            VmClusterNetworkLifecycleStateCreating,
	"REQUIRES_VALIDATION": VmClusterNetworkLifecycleStateRequiresValidation,
	"VALIDATING":          VmClusterNetworkLifecycleStateValidating,
	"VALIDATED":           VmClusterNetworkLifecycleStateValidated,
	"VALIDATION_FAILED":   VmClusterNetworkLifecycleStateValidationFailed,
	"UPDATING":            VmClusterNetworkLifecycleStateUpdating,
	"ALLOCATED":           VmClusterNetworkLifecycleStateAllocated,
	"TERMINATING":         VmClusterNetworkLifecycleStateTerminating,
	"TERMINATED":          VmClusterNetworkLifecycleStateTerminated,
	"FAILED":              VmClusterNetworkLifecycleStateFailed,
	"NEEDS_ATTENTION":     VmClusterNetworkLifecycleStateNeedsAttention,
}

var mappingVmClusterNetworkLifecycleStateEnumLowerCase = map[string]VmClusterNetworkLifecycleStateEnum{
	"creating":            VmClusterNetworkLifecycleStateCreating,
	"requires_validation": VmClusterNetworkLifecycleStateRequiresValidation,
	"validating":          VmClusterNetworkLifecycleStateValidating,
	"validated":           VmClusterNetworkLifecycleStateValidated,
	"validation_failed":   VmClusterNetworkLifecycleStateValidationFailed,
	"updating":            VmClusterNetworkLifecycleStateUpdating,
	"allocated":           VmClusterNetworkLifecycleStateAllocated,
	"terminating":         VmClusterNetworkLifecycleStateTerminating,
	"terminated":          VmClusterNetworkLifecycleStateTerminated,
	"failed":              VmClusterNetworkLifecycleStateFailed,
	"needs_attention":     VmClusterNetworkLifecycleStateNeedsAttention,
}

// GetVmClusterNetworkLifecycleStateEnumValues Enumerates the set of values for VmClusterNetworkLifecycleStateEnum
func GetVmClusterNetworkLifecycleStateEnumValues() []VmClusterNetworkLifecycleStateEnum {
	values := make([]VmClusterNetworkLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterNetworkLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterNetworkLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterNetworkLifecycleStateEnum
func GetVmClusterNetworkLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"REQUIRES_VALIDATION",
		"VALIDATING",
		"VALIDATED",
		"VALIDATION_FAILED",
		"UPDATING",
		"ALLOCATED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingVmClusterNetworkLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterNetworkLifecycleStateEnum(val string) (VmClusterNetworkLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterNetworkLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
