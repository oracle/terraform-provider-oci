// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddServiceDetails The configuration details for adding new services to the existing Stack.
type AddServiceDetails struct {

	// List of templates to be added for the stack.
	StackTemplates []StackTemplateEnum `mandatory:"true" json:"stackTemplates"`

	// List of services to be added for the stack.
	Services []ServiceEnum `mandatory:"true" json:"services"`

	// ADB details if adb is included in the services to be added.
	Adb []AdbDetail `mandatory:"false" json:"adb"`

	// GGCS details if ggcs is included in the services to be added.
	Ggcs []GgcsDetail `mandatory:"false" json:"ggcs"`

	// DATAFLOW details if dataflow is included in the services to be added.
	Dataflow []DataflowDetail `mandatory:"false" json:"dataflow"`

	// Object Storage Details if object storage is included in services to be added.
	Objectstorage []ObjectStorageDetail `mandatory:"false" json:"objectstorage"`

	// GenAI Details if genai is included in services to be added.
	Genai []GenAiDetail `mandatory:"false" json:"genai"`
}

func (m AddServiceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddServiceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.StackTemplates {
		if _, ok := GetMappingStackTemplateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackTemplates: %s. Supported values are: %s.", val, strings.Join(GetStackTemplateEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Services {
		if _, ok := GetMappingServiceEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Services: %s. Supported values are: %s.", val, strings.Join(GetServiceEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
