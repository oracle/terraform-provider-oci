// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionsRuntimeSummary Summary of the FunctionsRuntime.
type FunctionsRuntimeSummary struct {

	// The OCID of the FunctionsRuntime that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A brief descriptive name for the FunctionsRuntime. The FunctionsRuntime name must be unique, and not match any existing
	//   FunctionsRuntime.
	Name *string `mandatory:"true" json:"name"`

	// The operating system of the FunctionsRuntime. This is the OS that the FunctionsRuntime provides for execution of customer payloads.
	Os *string `mandatory:"true" json:"os"`

	// The programming language of the FunctionsRuntime. This is the language that the FunctionsRuntime provides for execution of customer payloads.
	Language *string `mandatory:"true" json:"language"`

	// The time when the FunctionsRuntime will be deprecated. An RFC3339 formatted datetime string.
	TimeDeprecated *common.SDKTime `mandatory:"true" json:"timeDeprecated"`

	// The time when the FunctionsRuntime will be decommissioned. An RFC3339 formatted datetime string.
	TimeDecommissioned *common.SDKTime `mandatory:"true" json:"timeDecommissioned"`

	// The current state of the FunctionsRuntime resource.
	LifecycleState FunctionsRuntimeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the current FunctionsRuntimeVersion for this FunctionsRuntime.
	CurrentFunctionsRuntimeVersionId *string `mandatory:"true" json:"currentFunctionsRuntimeVersionId"`

	// The time when the FunctionsRuntime was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the FunctionsRuntime was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Metadata for the FunctionsRuntime Resource.
	Metadata *string `mandatory:"false" json:"metadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FunctionsRuntimeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionsRuntimeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFunctionsRuntimeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFunctionsRuntimeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
