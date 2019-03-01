// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API.
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateJobDetails Updates the display name, free-form tags, and/or defined tag properties of the job.
type UpdateJobDetails struct {

	// The new display name to set.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateJobDetails) String() string {
	return common.PointerString(m)
}
