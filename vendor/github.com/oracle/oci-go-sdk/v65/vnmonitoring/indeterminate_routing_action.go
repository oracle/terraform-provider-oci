// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IndeterminateRoutingAction Defines the routing action taken on a traffic node where the routing action is INDETERMINATE.
type IndeterminateRoutingAction struct {

	// The type of the routing support for the traffic flow.
	ActionType RoutingActionActionTypeEnum `mandatory:"true" json:"actionType"`
}

// GetActionType returns ActionType
func (m IndeterminateRoutingAction) GetActionType() RoutingActionActionTypeEnum {
	return m.ActionType
}

func (m IndeterminateRoutingAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndeterminateRoutingAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetRoutingActionActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IndeterminateRoutingAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIndeterminateRoutingAction IndeterminateRoutingAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeIndeterminateRoutingAction
	}{
		"INDETERMINATE",
		(MarshalTypeIndeterminateRoutingAction)(m),
	}

	return json.Marshal(&s)
}
