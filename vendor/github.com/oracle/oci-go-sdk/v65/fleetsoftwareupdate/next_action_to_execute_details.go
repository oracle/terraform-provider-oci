// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NextActionToExecuteDetails Details of the next Exadata Fleet Update Action to execute in a Maintenance Cycle.
type NextActionToExecuteDetails struct {

	// Type of Exadata Fleet Update Action
	Type DetailedActionTypesEnum `mandatory:"true" json:"type"`

	// The date and time the Exadata Fleet Update Action is expected to start. Null if no Action has been scheduled.
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeToStart *common.SDKTime `mandatory:"false" json:"timeToStart"`
}

func (m NextActionToExecuteDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NextActionToExecuteDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetailedActionTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDetailedActionTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
