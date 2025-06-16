// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchTaskStepValidationsMessage Validation message
type PatchTaskStepValidationsMessage struct {

	// Start Time
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Validation Name
	ValidationName *string `mandatory:"false" json:"validationName"`

	// Status of the validation
	Result PatchTaskStepStatusEnum `mandatory:"false" json:"result,omitempty"`

	// Validation Name
	Message *string `mandatory:"false" json:"message"`
}

func (m PatchTaskStepValidationsMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTaskStepValidationsMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchTaskStepStatusEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetPatchTaskStepStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
