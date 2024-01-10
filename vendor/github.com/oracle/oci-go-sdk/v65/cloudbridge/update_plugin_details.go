// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePluginDetails The information to be updated.
type UpdatePluginDetails struct {

	// State to which the customer wants the plugin to move to.
	DesiredState PluginDesiredStateEnum `mandatory:"false" json:"desiredState,omitempty"`
}

func (m UpdatePluginDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePluginDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPluginDesiredStateEnum(string(m.DesiredState)); !ok && m.DesiredState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DesiredState: %s. Supported values are: %s.", m.DesiredState, strings.Join(GetPluginDesiredStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
