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

// AllowedSecurityActionDetails Defines details for the security action taken on allowed traffic.
type AllowedSecurityActionDetails struct {

	// If true, the allowed security configuration details are incomplete.
	IsRestrictedOrPartial *bool `mandatory:"true" json:"isRestrictedOrPartial"`

	AllowedSecurityConfiguration AllowedSecurityConfiguration `mandatory:"false" json:"allowedSecurityConfiguration"`
}

func (m AllowedSecurityActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllowedSecurityActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AllowedSecurityActionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AllowedSecurityConfiguration allowedsecurityconfiguration `json:"allowedSecurityConfiguration"`
		IsRestrictedOrPartial        *bool                        `json:"isRestrictedOrPartial"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.AllowedSecurityConfiguration.UnmarshalPolymorphicJSON(model.AllowedSecurityConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AllowedSecurityConfiguration = nn.(AllowedSecurityConfiguration)
	} else {
		m.AllowedSecurityConfiguration = nil
	}

	m.IsRestrictedOrPartial = model.IsRestrictedOrPartial

	return
}
