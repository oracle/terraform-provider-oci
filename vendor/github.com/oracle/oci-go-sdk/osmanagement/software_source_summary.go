// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SoftwareSourceSummary A software source contains a collection of packages
type SoftwareSourceSummary struct {

	// OCID for the Software Source
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name for the software source
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the Software Source
	RepoType *string `mandatory:"true" json:"repoType"`

	// Information specified by the user about the software source
	Description *string `mandatory:"false" json:"description"`

	// status of the software source.
	Status SoftwareSourceSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Number of packages
	Packages *int `mandatory:"false" json:"packages"`

	// The current state of the software source.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// OCID for the parent software source, if there is one
	ParentId *string `mandatory:"false" json:"parentId"`

	// Display name the parent software source, if there is one
	ParentName *string `mandatory:"false" json:"parentName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SoftwareSourceSummary) String() string {
	return common.PointerString(m)
}

// SoftwareSourceSummaryStatusEnum Enum with underlying type: string
type SoftwareSourceSummaryStatusEnum string

// Set of constants representing the allowable values for SoftwareSourceSummaryStatusEnum
const (
	SoftwareSourceSummaryStatusNormal      SoftwareSourceSummaryStatusEnum = "NORMAL"
	SoftwareSourceSummaryStatusUnreachable SoftwareSourceSummaryStatusEnum = "UNREACHABLE"
	SoftwareSourceSummaryStatusError       SoftwareSourceSummaryStatusEnum = "ERROR"
	SoftwareSourceSummaryStatusWarning     SoftwareSourceSummaryStatusEnum = "WARNING"
)

var mappingSoftwareSourceSummaryStatus = map[string]SoftwareSourceSummaryStatusEnum{
	"NORMAL":      SoftwareSourceSummaryStatusNormal,
	"UNREACHABLE": SoftwareSourceSummaryStatusUnreachable,
	"ERROR":       SoftwareSourceSummaryStatusError,
	"WARNING":     SoftwareSourceSummaryStatusWarning,
}

// GetSoftwareSourceSummaryStatusEnumValues Enumerates the set of values for SoftwareSourceSummaryStatusEnum
func GetSoftwareSourceSummaryStatusEnumValues() []SoftwareSourceSummaryStatusEnum {
	values := make([]SoftwareSourceSummaryStatusEnum, 0)
	for _, v := range mappingSoftwareSourceSummaryStatus {
		values = append(values, v)
	}
	return values
}
