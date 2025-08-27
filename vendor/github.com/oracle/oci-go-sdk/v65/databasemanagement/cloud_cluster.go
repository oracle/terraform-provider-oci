// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudCluster The details of a cloud cluster.
type CloudCluster struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud cluster. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud cluster.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the cluster is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud cluster.
	LifecycleState CloudClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud cluster was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud cluster was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
	CloudConnectorId *string `mandatory:"false" json:"cloudConnectorId"`

	// The directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// Indicates whether the cluster is Oracle Flex Cluster or not.
	IsFlexCluster *bool `mandatory:"false" json:"isFlexCluster"`

	// The additional details of the cloud cluster defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The list of network address configurations of the cloud cluster.
	NetworkConfigurations []CloudClusterNetworkConfiguration `mandatory:"false" json:"networkConfigurations"`

	// The list of Virtual IP (VIP) configurations of the cloud cluster.
	VipConfigurations []CloudClusterVipConfiguration `mandatory:"false" json:"vipConfigurations"`

	// The list of Single Client Access Name (SCAN) configurations of the cloud cluster.
	ScanConfigurations []CloudClusterScanListenerConfiguration `mandatory:"false" json:"scanConfigurations"`

	// The location of the Oracle Cluster Registry (OCR).
	OcrFileLocation *string `mandatory:"false" json:"ocrFileLocation"`

	// The cluster version.
	Version *string `mandatory:"false" json:"version"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudClusterLifecycleStateEnum Enum with underlying type: string
type CloudClusterLifecycleStateEnum string

// Set of constants representing the allowable values for CloudClusterLifecycleStateEnum
const (
	CloudClusterLifecycleStateCreating     CloudClusterLifecycleStateEnum = "CREATING"
	CloudClusterLifecycleStateNotConnected CloudClusterLifecycleStateEnum = "NOT_CONNECTED"
	CloudClusterLifecycleStateActive       CloudClusterLifecycleStateEnum = "ACTIVE"
	CloudClusterLifecycleStateInactive     CloudClusterLifecycleStateEnum = "INACTIVE"
	CloudClusterLifecycleStateUpdating     CloudClusterLifecycleStateEnum = "UPDATING"
	CloudClusterLifecycleStateDeleting     CloudClusterLifecycleStateEnum = "DELETING"
	CloudClusterLifecycleStateDeleted      CloudClusterLifecycleStateEnum = "DELETED"
	CloudClusterLifecycleStateFailed       CloudClusterLifecycleStateEnum = "FAILED"
)

var mappingCloudClusterLifecycleStateEnum = map[string]CloudClusterLifecycleStateEnum{
	"CREATING":      CloudClusterLifecycleStateCreating,
	"NOT_CONNECTED": CloudClusterLifecycleStateNotConnected,
	"ACTIVE":        CloudClusterLifecycleStateActive,
	"INACTIVE":      CloudClusterLifecycleStateInactive,
	"UPDATING":      CloudClusterLifecycleStateUpdating,
	"DELETING":      CloudClusterLifecycleStateDeleting,
	"DELETED":       CloudClusterLifecycleStateDeleted,
	"FAILED":        CloudClusterLifecycleStateFailed,
}

var mappingCloudClusterLifecycleStateEnumLowerCase = map[string]CloudClusterLifecycleStateEnum{
	"creating":      CloudClusterLifecycleStateCreating,
	"not_connected": CloudClusterLifecycleStateNotConnected,
	"active":        CloudClusterLifecycleStateActive,
	"inactive":      CloudClusterLifecycleStateInactive,
	"updating":      CloudClusterLifecycleStateUpdating,
	"deleting":      CloudClusterLifecycleStateDeleting,
	"deleted":       CloudClusterLifecycleStateDeleted,
	"failed":        CloudClusterLifecycleStateFailed,
}

// GetCloudClusterLifecycleStateEnumValues Enumerates the set of values for CloudClusterLifecycleStateEnum
func GetCloudClusterLifecycleStateEnumValues() []CloudClusterLifecycleStateEnum {
	values := make([]CloudClusterLifecycleStateEnum, 0)
	for _, v := range mappingCloudClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudClusterLifecycleStateEnumStringValues Enumerates the set of values in String for CloudClusterLifecycleStateEnum
func GetCloudClusterLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudClusterLifecycleStateEnum(val string) (CloudClusterLifecycleStateEnum, bool) {
	enum, ok := mappingCloudClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
