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

// BulkDeleteSqlFirewallAllowedSqlsDetails Details of the allowed SQLs to be deleted from the SQL firewall policy.
type BulkDeleteSqlFirewallAllowedSqlsDetails struct {

	// The OCID of the SQL firewall policy whose allowed SQLs needs to be deleted.
	SqlFirewallPolicyId *string `mandatory:"true" json:"sqlFirewallPolicyId"`

	Selection SelectionDetails `mandatory:"true" json:"selection"`
}

func (m BulkDeleteSqlFirewallAllowedSqlsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkDeleteSqlFirewallAllowedSqlsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BulkDeleteSqlFirewallAllowedSqlsDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SqlFirewallPolicyId *string          `json:"sqlFirewallPolicyId"`
		Selection           selectiondetails `json:"selection"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SqlFirewallPolicyId = model.SqlFirewallPolicyId

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
