// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutoScalingConfigurationDetails The representation of UpdateAutoScalingConfigurationDetails
type UpdateAutoScalingConfigurationDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// If the AutoScalingConfiguration is enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The minimum period of time between scaling actions. The default is 300 seconds.
	CoolDownInSeconds *int `mandatory:"false" json:"coolDownInSeconds"`
}

func (m UpdateAutoScalingConfigurationDetails) String() string {
	return common.PointerString(m)
}
