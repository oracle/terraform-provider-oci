// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service limits APIs
//
// APIs that interact with the resource limits of a specific resource type
//

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateQuotaDetails Request object for update quota operation.
type UpdateQuotaDetails struct {

	// The description you assign to the quota.
	Description *string `mandatory:"false" json:"description"`

	// An array of quota statements written in the declarative quota statement language.
	Statements []string `mandatory:"false" json:"statements"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateQuotaDetails) String() string {
	return common.PointerString(m)
}
