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

// MulticloudResourceSummary The multicloud resource, for eg. VMCluster, ExaInfra, and its attributes. The resource and network anchor that represents
type MulticloudResourceSummary struct {

	// The Id of the multicloud resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The date and time the subscription was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Endpoint used to retrieve displayName and lifeCycleState of the resource.
	ResourceDisplayName *string `mandatory:"false" json:"resourceDisplayName"`

	// What resource it refers to. Eg. VMCluster, ExaInfra, etc.
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Compartment name associated the resource.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// Compartment Id of the resource.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Resource Anchor name.
	VcnName *string `mandatory:"false" json:"vcnName"`

	// Id of the Virtual Cloud Network associated to the resource.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// Name of the network anchor associated to the resource.
	NetworkAnchorName *string `mandatory:"false" json:"networkAnchorName"`

	// OCID of the Network Anchor
	NetworkAnchorId *string `mandatory:"false" json:"networkAnchorId"`

	// Resource Id that comes from the Multi Cloud Control Plane
	CspResourceId *string `mandatory:"false" json:"cspResourceId"`

	// CSP Specific Additional Properties, AzureSubnetId for Azure
	CspAdditionalProperties map[string]string `mandatory:"false" json:"cspAdditionalProperties"`

	// The date and time the subscription was updated, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the multicloud resource.
	LifecycleState MulticloudResourceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MulticloudResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MulticloudResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMulticloudResourceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMulticloudResourceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MulticloudResourceSummaryLifecycleStateEnum Enum with underlying type: string
type MulticloudResourceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for MulticloudResourceSummaryLifecycleStateEnum
const (
	MulticloudResourceSummaryLifecycleStateActive   MulticloudResourceSummaryLifecycleStateEnum = "ACTIVE"
	MulticloudResourceSummaryLifecycleStateInactive MulticloudResourceSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingMulticloudResourceSummaryLifecycleStateEnum = map[string]MulticloudResourceSummaryLifecycleStateEnum{
	"ACTIVE":   MulticloudResourceSummaryLifecycleStateActive,
	"INACTIVE": MulticloudResourceSummaryLifecycleStateInactive,
}

var mappingMulticloudResourceSummaryLifecycleStateEnumLowerCase = map[string]MulticloudResourceSummaryLifecycleStateEnum{
	"active":   MulticloudResourceSummaryLifecycleStateActive,
	"inactive": MulticloudResourceSummaryLifecycleStateInactive,
}

// GetMulticloudResourceSummaryLifecycleStateEnumValues Enumerates the set of values for MulticloudResourceSummaryLifecycleStateEnum
func GetMulticloudResourceSummaryLifecycleStateEnumValues() []MulticloudResourceSummaryLifecycleStateEnum {
	values := make([]MulticloudResourceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingMulticloudResourceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudResourceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for MulticloudResourceSummaryLifecycleStateEnum
func GetMulticloudResourceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingMulticloudResourceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudResourceSummaryLifecycleStateEnum(val string) (MulticloudResourceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingMulticloudResourceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
