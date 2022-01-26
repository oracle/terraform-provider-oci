// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateLogDetails Update log object properties.
type UpdateLogDetails struct {

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Log retention duration in 30-day increments (30, 60, 90 and so on).
	RetentionDuration *int `mandatory:"false" json:"retentionDuration"`

	Configuration *UpdateConfigurationDetails `mandatory:"false" json:"configuration"`
}

func (m UpdateLogDetails) String() string {
	return common.PointerString(m)
}
