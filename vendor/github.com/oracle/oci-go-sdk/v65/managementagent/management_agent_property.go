// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementAgentProperty Property item in name/value pair, with optional unit type.
type ManagementAgentProperty struct {

	// Name of the property
	Name *string `mandatory:"true" json:"name"`

	// Values of the property
	Values []string `mandatory:"true" json:"values"`

	// Unit for the property
	Units PropertyUnitsEnum `mandatory:"false" json:"units,omitempty"`
}

func (m ManagementAgentProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPropertyUnitsEnum(string(m.Units)); !ok && m.Units != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Units: %s. Supported values are: %s.", m.Units, strings.Join(GetPropertyUnitsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
