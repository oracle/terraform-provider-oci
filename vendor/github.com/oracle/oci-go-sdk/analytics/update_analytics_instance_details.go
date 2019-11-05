// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAnalyticsInstanceDetails Input payload to update an Analytics instance. Fields that are not provided
// will not be updated.
type UpdateAnalyticsInstanceDetails struct {

	// Optional description.
	Description *string `mandatory:"false" json:"description"`

	// Email address receiving notifications.
	EmailNotification *string `mandatory:"false" json:"emailNotification"`

	// The license used for the service.
	LicenseType LicenseTypeEnum `mandatory:"false" json:"licenseType,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateAnalyticsInstanceDetails) String() string {
	return common.PointerString(m)
}
