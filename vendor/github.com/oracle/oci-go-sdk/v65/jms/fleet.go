// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Fleet A Fleet is the primary collection with which users interact when using Java Management Service.
type Fleet struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Fleet.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Fleet's description.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The approximate count of all unique Java Runtimes in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag.
	ApproximateJreCount *int `mandatory:"true" json:"approximateJreCount"`

	// The approximate count of all unique Java installations in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag.
	ApproximateInstallationCount *int `mandatory:"true" json:"approximateInstallationCount"`

	// The approximate count of all unique applications in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag.
	ApproximateApplicationCount *int `mandatory:"true" json:"approximateApplicationCount"`

	// The approximate count of all unique managed instances in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag.
	ApproximateManagedInstanceCount *int `mandatory:"true" json:"approximateManagedInstanceCount"`

	// The approximate count of all unique Java servers in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag.
	ApproximateJavaServerCount *int `mandatory:"true" json:"approximateJavaServerCount"`

	// The creation date and time of the Fleet (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the Fleet.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	InventoryLog *CustomLog `mandatory:"false" json:"inventoryLog"`

	OperationLog *CustomLog `mandatory:"false" json:"operationLog"`

	// Whether or not advanced features are enabled in this Fleet.
	// Deprecated, use `/fleets/{fleetId}/advanceFeatureConfiguration` API instead.
	IsAdvancedFeaturesEnabled *bool `mandatory:"false" json:"isAdvancedFeaturesEnabled"`

	// Whether or not export setting is enabled in this Fleet.
	IsExportSettingEnabled *bool `mandatory:"false" json:"isExportSettingEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Fleet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Fleet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
