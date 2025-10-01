// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkAnchor A NetworkAnchor is a description of a NetworkAnchor.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type NetworkAnchor struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI resource anchor Id (OCID).
	ResourceAnchorId *string `mandatory:"true" json:"resourceAnchorId"`

	// The date and time the NetworkAnchor was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the NetworkAnchor.
	LifecycleState NetworkAnchorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the NetworkAnchor was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the NetworkAnchor in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// AUTO_BIND - when passed compartment will be created on-behalf of customer and bind to this resource anchor
	// NO_AUTO_BIND - compartment will not be created and later customer can bind existing compartment.
	// to this resource anchor. This is for future use only
	SetupMode NetworkAnchorSetupModeEnum `mandatory:"false" json:"setupMode,omitempty"`

	// The CPG ID in which Network Anchor will be created.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	OciMetadataItem *OciNetworkMetadata `mandatory:"false" json:"ociMetadataItem"`

	CloudServiceProviderMetadataItem *CloudServiceProviderNetworkMetadataItem `mandatory:"false" json:"cloudServiceProviderMetadataItem"`
}

func (m NetworkAnchor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkAnchor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkAnchorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkAnchorLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNetworkAnchorSetupModeEnum(string(m.SetupMode)); !ok && m.SetupMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SetupMode: %s. Supported values are: %s.", m.SetupMode, strings.Join(GetNetworkAnchorSetupModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkAnchorLifecycleStateEnum Enum with underlying type: string
type NetworkAnchorLifecycleStateEnum string

// Set of constants representing the allowable values for NetworkAnchorLifecycleStateEnum
const (
	NetworkAnchorLifecycleStateCreating NetworkAnchorLifecycleStateEnum = "CREATING"
	NetworkAnchorLifecycleStateUpdating NetworkAnchorLifecycleStateEnum = "UPDATING"
	NetworkAnchorLifecycleStateActive   NetworkAnchorLifecycleStateEnum = "ACTIVE"
	NetworkAnchorLifecycleStateDeleting NetworkAnchorLifecycleStateEnum = "DELETING"
	NetworkAnchorLifecycleStateDeleted  NetworkAnchorLifecycleStateEnum = "DELETED"
	NetworkAnchorLifecycleStateFailed   NetworkAnchorLifecycleStateEnum = "FAILED"
)

var mappingNetworkAnchorLifecycleStateEnum = map[string]NetworkAnchorLifecycleStateEnum{
	"CREATING": NetworkAnchorLifecycleStateCreating,
	"UPDATING": NetworkAnchorLifecycleStateUpdating,
	"ACTIVE":   NetworkAnchorLifecycleStateActive,
	"DELETING": NetworkAnchorLifecycleStateDeleting,
	"DELETED":  NetworkAnchorLifecycleStateDeleted,
	"FAILED":   NetworkAnchorLifecycleStateFailed,
}

var mappingNetworkAnchorLifecycleStateEnumLowerCase = map[string]NetworkAnchorLifecycleStateEnum{
	"creating": NetworkAnchorLifecycleStateCreating,
	"updating": NetworkAnchorLifecycleStateUpdating,
	"active":   NetworkAnchorLifecycleStateActive,
	"deleting": NetworkAnchorLifecycleStateDeleting,
	"deleted":  NetworkAnchorLifecycleStateDeleted,
	"failed":   NetworkAnchorLifecycleStateFailed,
}

// GetNetworkAnchorLifecycleStateEnumValues Enumerates the set of values for NetworkAnchorLifecycleStateEnum
func GetNetworkAnchorLifecycleStateEnumValues() []NetworkAnchorLifecycleStateEnum {
	values := make([]NetworkAnchorLifecycleStateEnum, 0)
	for _, v := range mappingNetworkAnchorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAnchorLifecycleStateEnumStringValues Enumerates the set of values in String for NetworkAnchorLifecycleStateEnum
func GetNetworkAnchorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingNetworkAnchorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAnchorLifecycleStateEnum(val string) (NetworkAnchorLifecycleStateEnum, bool) {
	enum, ok := mappingNetworkAnchorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NetworkAnchorSetupModeEnum Enum with underlying type: string
type NetworkAnchorSetupModeEnum string

// Set of constants representing the allowable values for NetworkAnchorSetupModeEnum
const (
	NetworkAnchorSetupModeAutoBind   NetworkAnchorSetupModeEnum = "AUTO_BIND"
	NetworkAnchorSetupModeNoAutoBind NetworkAnchorSetupModeEnum = "NO_AUTO_BIND"
)

var mappingNetworkAnchorSetupModeEnum = map[string]NetworkAnchorSetupModeEnum{
	"AUTO_BIND":    NetworkAnchorSetupModeAutoBind,
	"NO_AUTO_BIND": NetworkAnchorSetupModeNoAutoBind,
}

var mappingNetworkAnchorSetupModeEnumLowerCase = map[string]NetworkAnchorSetupModeEnum{
	"auto_bind":    NetworkAnchorSetupModeAutoBind,
	"no_auto_bind": NetworkAnchorSetupModeNoAutoBind,
}

// GetNetworkAnchorSetupModeEnumValues Enumerates the set of values for NetworkAnchorSetupModeEnum
func GetNetworkAnchorSetupModeEnumValues() []NetworkAnchorSetupModeEnum {
	values := make([]NetworkAnchorSetupModeEnum, 0)
	for _, v := range mappingNetworkAnchorSetupModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAnchorSetupModeEnumStringValues Enumerates the set of values in String for NetworkAnchorSetupModeEnum
func GetNetworkAnchorSetupModeEnumStringValues() []string {
	return []string{
		"AUTO_BIND",
		"NO_AUTO_BIND",
	}
}

// GetMappingNetworkAnchorSetupModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAnchorSetupModeEnum(val string) (NetworkAnchorSetupModeEnum, bool) {
	enum, ok := mappingNetworkAnchorSetupModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
