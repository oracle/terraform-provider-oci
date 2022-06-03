// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete
// log groups, log objects, agent configurations, and log data models.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogDataModel Top level log data model resource object.
type LogDataModel struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	LogDataModelDetails *LogDataModelDetails `mandatory:"true" json:"logDataModelDetails"`

	// Lifecycle state of log data model.
	LifecycleState LogDataModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Validity state of log data model.
	ValidityState LogDataModelValidityStateEnum `mandatory:"true" json:"validityState"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"true" json:"timeLastModified"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Log data model template name.
	Template *string `mandatory:"false" json:"template"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LogDataModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogDataModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogDataModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogDataModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogDataModelValidityStateEnum(string(m.ValidityState)); !ok && m.ValidityState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidityState: %s. Supported values are: %s.", m.ValidityState, strings.Join(GetLogDataModelValidityStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
