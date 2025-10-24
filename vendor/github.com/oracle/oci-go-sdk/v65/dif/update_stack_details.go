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

// UpdateStackDetails Details to update a Stack.
type UpdateStackDetails struct {

	// email id to which the stack notifications would be sent.
	NotificationEmail *string `mandatory:"false" json:"notificationEmail"`

	// List of templates to be updated for the stack.
	StackTemplates []StackTemplateEnum `mandatory:"false" json:"stackTemplates,omitempty"`

	// List of services to be updated for the stack.
	Services []ServiceEnum `mandatory:"false" json:"services,omitempty"`

	// ADB details if adb is included in the services to be updated.
	Adb []AdbUpdateDetail `mandatory:"false" json:"adb"`

	// GGCS details if ggcs is included in the services to be updated.
	Ggcs []GgcsUpdateDetail `mandatory:"false" json:"ggcs"`

	// DATAFLOW details if dataflow is included in the services to be updated.
	Dataflow []DataflowUpdateDetail `mandatory:"false" json:"dataflow"`

	// Object Storage Details if object storage is included in services to be updated.
	Objectstorage []ObjectStorageUpdateDetail `mandatory:"false" json:"objectstorage"`

	// GenAI Details if genai is included in services to be updated.
	Genai []GenAiUpdateDetail `mandatory:"false" json:"genai"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateStackDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateStackDetails) ValidateEnumValue() (bool, error) {
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
