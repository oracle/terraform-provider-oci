// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalCluster The details of an external cluster.
type ExternalCluster struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external cluster.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external cluster. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external cluster.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the cluster is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external cluster.
	LifecycleState ExternalClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external cluster was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external cluster was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external connector.
	ExternalConnectorId *string `mandatory:"false" json:"externalConnectorId"`

	// The directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// Indicates whether the cluster is Oracle Flex Cluster or not.
	IsFlexCluster *bool `mandatory:"false" json:"isFlexCluster"`

	// The additional details of the external cluster defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The list of network address configurations of the external cluster.
	NetworkConfigurations []ExternalClusterNetworkConfiguration `mandatory:"false" json:"networkConfigurations"`

	// The list of Virtual IP (VIP) configurations of the external cluster.
	VipConfigurations []ExternalClusterVipConfiguration `mandatory:"false" json:"vipConfigurations"`

	// The list of Single Client Access Name (SCAN) configurations of the external cluster.
	ScanConfigurations []ExternalClusterScanListenerConfiguration `mandatory:"false" json:"scanConfigurations"`

	// The location of the Oracle Cluster Registry (OCR).
	OcrFileLocation *string `mandatory:"false" json:"ocrFileLocation"`

	// The cluster version.
	Version *string `mandatory:"false" json:"version"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExternalCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalClusterLifecycleStateEnum Enum with underlying type: string
type ExternalClusterLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalClusterLifecycleStateEnum
const (
	ExternalClusterLifecycleStateCreating     ExternalClusterLifecycleStateEnum = "CREATING"
	ExternalClusterLifecycleStateNotConnected ExternalClusterLifecycleStateEnum = "NOT_CONNECTED"
	ExternalClusterLifecycleStateActive       ExternalClusterLifecycleStateEnum = "ACTIVE"
	ExternalClusterLifecycleStateInactive     ExternalClusterLifecycleStateEnum = "INACTIVE"
	ExternalClusterLifecycleStateUpdating     ExternalClusterLifecycleStateEnum = "UPDATING"
	ExternalClusterLifecycleStateDeleting     ExternalClusterLifecycleStateEnum = "DELETING"
	ExternalClusterLifecycleStateDeleted      ExternalClusterLifecycleStateEnum = "DELETED"
	ExternalClusterLifecycleStateFailed       ExternalClusterLifecycleStateEnum = "FAILED"
)

var mappingExternalClusterLifecycleStateEnum = map[string]ExternalClusterLifecycleStateEnum{
	"CREATING":      ExternalClusterLifecycleStateCreating,
	"NOT_CONNECTED": ExternalClusterLifecycleStateNotConnected,
	"ACTIVE":        ExternalClusterLifecycleStateActive,
	"INACTIVE":      ExternalClusterLifecycleStateInactive,
	"UPDATING":      ExternalClusterLifecycleStateUpdating,
	"DELETING":      ExternalClusterLifecycleStateDeleting,
	"DELETED":       ExternalClusterLifecycleStateDeleted,
	"FAILED":        ExternalClusterLifecycleStateFailed,
}

var mappingExternalClusterLifecycleStateEnumLowerCase = map[string]ExternalClusterLifecycleStateEnum{
	"creating":      ExternalClusterLifecycleStateCreating,
	"not_connected": ExternalClusterLifecycleStateNotConnected,
	"active":        ExternalClusterLifecycleStateActive,
	"inactive":      ExternalClusterLifecycleStateInactive,
	"updating":      ExternalClusterLifecycleStateUpdating,
	"deleting":      ExternalClusterLifecycleStateDeleting,
	"deleted":       ExternalClusterLifecycleStateDeleted,
	"failed":        ExternalClusterLifecycleStateFailed,
}

// GetExternalClusterLifecycleStateEnumValues Enumerates the set of values for ExternalClusterLifecycleStateEnum
func GetExternalClusterLifecycleStateEnumValues() []ExternalClusterLifecycleStateEnum {
	values := make([]ExternalClusterLifecycleStateEnum, 0)
	for _, v := range mappingExternalClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalClusterLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalClusterLifecycleStateEnum
func GetExternalClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalClusterLifecycleStateEnum(val string) (ExternalClusterLifecycleStateEnum, bool) {
	enum, ok := mappingExternalClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
