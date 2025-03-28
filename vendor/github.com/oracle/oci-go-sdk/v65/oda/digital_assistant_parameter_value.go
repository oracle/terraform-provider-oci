// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DigitalAssistantParameterValue Properties for configuring a Parameter in a Digital Assistant instance.
type DigitalAssistantParameterValue struct {

	// The Parameter name.  This must be unique within the parent resource.
	Name *string `mandatory:"true" json:"name"`

	// The value type.
	Type ParameterTypeEnum `mandatory:"true" json:"type"`

	// The current value.  The value will be interpreted based on the `type`.
	Value *string `mandatory:"true" json:"value"`
}

func (m DigitalAssistantParameterValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DigitalAssistantParameterValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingParameterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetParameterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
