// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DataSafeConfiguration A Data Safe configuration for a tenancy and region.
type DataSafeConfiguration struct {

	// Indicates if Data Safe is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The URL of the Data Safe service.
	Url *string `mandatory:"false" json:"url"`

	// The OCID of the tenancy used to enable Data Safe.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The date and time Data Safe was enabled, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnabled *common.SDKTime `mandatory:"false" json:"timeEnabled"`

	// The current state of Data Safe.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DataSafeConfiguration) String() string {
	return common.PointerString(m)
}
