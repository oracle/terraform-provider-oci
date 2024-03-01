// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerCapabilities Linux Container capabilities to configure capabilities of container.
type ContainerCapabilities struct {

	// A list of additional configurable container capabilities.
	AddCapabilities []ContainerCapabilityTypeEnum `mandatory:"false" json:"addCapabilities,omitempty"`

	// A list of container capabilities that can be dropped.
	DropCapabilities []ContainerCapabilityTypeEnum `mandatory:"false" json:"dropCapabilities,omitempty"`
}

func (m ContainerCapabilities) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerCapabilities) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AddCapabilities {
		if _, ok := GetMappingContainerCapabilityTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AddCapabilities: %s. Supported values are: %s.", val, strings.Join(GetContainerCapabilityTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range m.DropCapabilities {
		if _, ok := GetMappingContainerCapabilityTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DropCapabilities: %s. Supported values are: %s.", val, strings.Join(GetContainerCapabilityTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
