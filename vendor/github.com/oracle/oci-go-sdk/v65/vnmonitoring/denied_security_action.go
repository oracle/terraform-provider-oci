// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeniedSecurityAction Defines the security action taken on denied traffic.
type DeniedSecurityAction struct {
	DeniedSecurityActionDetails *DeniedSecurityActionDetails `mandatory:"false" json:"deniedSecurityActionDetails"`

	// Type of the `SecurityAction`.
	ActionType SecurityActionActionTypeEnum `mandatory:"true" json:"actionType"`
}

// GetActionType returns ActionType
func (m DeniedSecurityAction) GetActionType() SecurityActionActionTypeEnum {
	return m.ActionType
}

func (m DeniedSecurityAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeniedSecurityAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetSecurityActionActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeniedSecurityAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeniedSecurityAction DeniedSecurityAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeDeniedSecurityAction
	}{
		"DENIED",
		(MarshalTypeDeniedSecurityAction)(m),
	}

	return json.Marshal(&s)
}
