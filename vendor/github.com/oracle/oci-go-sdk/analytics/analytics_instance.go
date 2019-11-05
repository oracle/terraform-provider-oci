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

// AnalyticsInstance Analytics Instance metadata.
type AnalyticsInstance struct {

	// The resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of an instance.
	LifecycleState AnalyticsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Analytics feature set.
	FeatureSet FeatureSetEnum `mandatory:"true" json:"featureSet"`

	Capacity *Capacity `mandatory:"true" json:"capacity"`

	// The date and time the instance was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Optional description.
	Description *string `mandatory:"false" json:"description"`

	// The license used for the service.
	LicenseType LicenseTypeEnum `mandatory:"false" json:"licenseType,omitempty"`

	// Email address receiving notifications.
	EmailNotification *string `mandatory:"false" json:"emailNotification"`

	// URL of the Analytics service.
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The date and time the instance was last updated (in the format defined by RFC3339).
	// This timestamp represents updates made through this API. External events do not
	// influence it.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m AnalyticsInstance) String() string {
	return common.PointerString(m)
}
