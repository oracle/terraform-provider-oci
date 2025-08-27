// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Provision A FamProvision is a description of a FamProvision.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Provision struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the FamProvision.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableTenancyId *string `mandatory:"true" json:"tfVariableTenancyId"`

	// A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableRegionId *string `mandatory:"true" json:"tfVariableRegionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the FamProvision was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the FamProvision.
	LifecycleState ProvisionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item.
	PackageCatalogItemId *string `mandatory:"true" json:"packageCatalogItemId"`

	// A OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item to a file with key/value pairs to set up variables for createStack API.
	ConfigCatalogItemId *string `mandatory:"true" json:"configCatalogItemId"`

	// A display Name of the Catalog Item in the Catalog.
	PackageCatalogItemDisplayName *string `mandatory:"true" json:"packageCatalogItemDisplayName"`

	// A listing ID of the Catalog Item in the Catalog.
	PackageCatalogItemListingId *string `mandatory:"true" json:"packageCatalogItemListingId"`

	// A listing version of the Catalog Item in the Catalog.
	PackageCatalogItemListingVersion *string `mandatory:"true" json:"packageCatalogItemListingVersion"`

	// A display Name of the Catalog Item in the Catalog.
	ConfigCatalogItemDisplayName *string `mandatory:"true" json:"configCatalogItemDisplayName"`

	// A listing ID of the Catalog Item in the Catalog.
	ConfigCatalogItemListingId *string `mandatory:"true" json:"configCatalogItemListingId"`

	// A listing version of the Catalog Item in the Catalog.
	ConfigCatalogItemListingVersion *string `mandatory:"true" json:"configCatalogItemListingVersion"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A description of the provision.
	ProvisionDescription *string `mandatory:"false" json:"provisionDescription"`

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCurrentUserId *string `mandatory:"false" json:"tfVariableCurrentUserId"`

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCompartmentId *string `mandatory:"false" json:"tfVariableCompartmentId"`

	// The date and time the FamProvision was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the FamProvision in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the RMS Stack.
	StackId *string `mandatory:"false" json:"stackId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the RMS APPLY Job.
	RmsApplyJobId *string `mandatory:"false" json:"rmsApplyJobId"`

	// Outputs from the Terraform Apply job
	TfOutputs []JobExecutionDetails `mandatory:"false" json:"tfOutputs"`

	// The deployed resources and their summary
	DeployedResources []DeployedResourceDetails `mandatory:"false" json:"deployedResources"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Provision) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Provision) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProvisionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProvisionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProvisionLifecycleStateEnum Enum with underlying type: string
type ProvisionLifecycleStateEnum string

// Set of constants representing the allowable values for ProvisionLifecycleStateEnum
const (
	ProvisionLifecycleStateCreating ProvisionLifecycleStateEnum = "CREATING"
	ProvisionLifecycleStateUpdating ProvisionLifecycleStateEnum = "UPDATING"
	ProvisionLifecycleStateActive   ProvisionLifecycleStateEnum = "ACTIVE"
	ProvisionLifecycleStateDeleting ProvisionLifecycleStateEnum = "DELETING"
	ProvisionLifecycleStateDeleted  ProvisionLifecycleStateEnum = "DELETED"
	ProvisionLifecycleStateFailed   ProvisionLifecycleStateEnum = "FAILED"
)

var mappingProvisionLifecycleStateEnum = map[string]ProvisionLifecycleStateEnum{
	"CREATING": ProvisionLifecycleStateCreating,
	"UPDATING": ProvisionLifecycleStateUpdating,
	"ACTIVE":   ProvisionLifecycleStateActive,
	"DELETING": ProvisionLifecycleStateDeleting,
	"DELETED":  ProvisionLifecycleStateDeleted,
	"FAILED":   ProvisionLifecycleStateFailed,
}

var mappingProvisionLifecycleStateEnumLowerCase = map[string]ProvisionLifecycleStateEnum{
	"creating": ProvisionLifecycleStateCreating,
	"updating": ProvisionLifecycleStateUpdating,
	"active":   ProvisionLifecycleStateActive,
	"deleting": ProvisionLifecycleStateDeleting,
	"deleted":  ProvisionLifecycleStateDeleted,
	"failed":   ProvisionLifecycleStateFailed,
}

// GetProvisionLifecycleStateEnumValues Enumerates the set of values for ProvisionLifecycleStateEnum
func GetProvisionLifecycleStateEnumValues() []ProvisionLifecycleStateEnum {
	values := make([]ProvisionLifecycleStateEnum, 0)
	for _, v := range mappingProvisionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProvisionLifecycleStateEnumStringValues Enumerates the set of values in String for ProvisionLifecycleStateEnum
func GetProvisionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingProvisionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProvisionLifecycleStateEnum(val string) (ProvisionLifecycleStateEnum, bool) {
	enum, ok := mappingProvisionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
