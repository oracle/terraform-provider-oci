// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ProjectSummary Metadata about the project.
type ProjectSummary struct {

	// Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// When the project was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the project.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Human-friendly name for the project, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// When the project was created, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail which can provide actionable information if creation failed.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ProjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProjectLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProjectLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
