// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkCreateSqlFirewallAllowedSqlsDetails The details used to append the violation logs as allowed SQLs
type BulkCreateSqlFirewallAllowedSqlsDetails struct {

	// The OCID of the SQL firewall policy where new allowed SQLs needs to be added.
	SqlFirewallPolicyId *string `mandatory:"true" json:"sqlFirewallPolicyId"`

	// The type of log to be added as an allowed sql.
	LogType LogTypeEnum `mandatory:"true" json:"logType"`

	Selection SelectionDetails `mandatory:"true" json:"selection"`
}

func (m BulkCreateSqlFirewallAllowedSqlsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCreateSqlFirewallAllowedSqlsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogTypeEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetLogTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BulkCreateSqlFirewallAllowedSqlsDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SqlFirewallPolicyId *string          `json:"sqlFirewallPolicyId"`
		LogType             LogTypeEnum      `json:"logType"`
		Selection           selectiondetails `json:"selection"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SqlFirewallPolicyId = model.SqlFirewallPolicyId

	m.LogType = model.LogType

	nn, e = model.Selection.UnmarshalPolymorphicJSON(model.Selection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Selection = nn.(SelectionDetails)
	} else {
		m.Selection = nil
	}

	return
}
