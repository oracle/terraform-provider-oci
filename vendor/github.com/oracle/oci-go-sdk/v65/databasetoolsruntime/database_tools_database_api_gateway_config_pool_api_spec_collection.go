// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection List of Database Tools database API gateway config API spec summary items.
type DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection struct {

	// Array of Database Tools database API gateway config API spec summary items.
	Items []DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummary `mandatory:"true" json:"items"`
}

func (m DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []databasetoolsdatabaseapigatewayconfigpoolapispecsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummary)
		} else {
			m.Items[i] = nil
		}
	}
	return
}
