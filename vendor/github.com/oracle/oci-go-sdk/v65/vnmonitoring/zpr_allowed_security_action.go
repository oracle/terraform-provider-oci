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

// ZprAllowedSecurityAction Defines the security action taken on zpr allowed traffic.
type ZprAllowedSecurityAction struct {
	AllowedSecurityActionDetails *ZprAllowedSecurityActionDetails `mandatory:"false" json:"allowedSecurityActionDetails"`

	// Type of the `SecurityAction`.
	ActionType SecurityActionActionTypeEnum `mandatory:"true" json:"actionType"`
}

// GetActionType returns ActionType
func (m ZprAllowedSecurityAction) GetActionType() SecurityActionActionTypeEnum {
	return m.ActionType
}

func (m ZprAllowedSecurityAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprAllowedSecurityAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetSecurityActionActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ZprAllowedSecurityAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeZprAllowedSecurityAction ZprAllowedSecurityAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeZprAllowedSecurityAction
	}{
		"ZPR_ALLOWED",
		(MarshalTypeZprAllowedSecurityAction)(m),
	}

	return json.Marshal(&s)
}
