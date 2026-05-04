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

// DatabaseToolsDatabaseApiGatewayConfigContent The content of a Database Tools database API gateway config global resource.
type DatabaseToolsDatabaseApiGatewayConfigContent struct {
	Global DatabaseToolsDatabaseApiGatewayConfigGlobal `mandatory:"true" json:"global"`

	// The content of the Database Tools database API gateway config pools each with API spec and auto API spec definitions.
	Pools []DatabaseToolsDatabaseApiGatewayConfigPoolContent `mandatory:"false" json:"pools"`
}

func (m DatabaseToolsDatabaseApiGatewayConfigContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsDatabaseApiGatewayConfigContent) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Pools  []databasetoolsdatabaseapigatewayconfigpoolcontent `json:"pools"`
		Global databasetoolsdatabaseapigatewayconfigglobal        `json:"global"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Pools = make([]DatabaseToolsDatabaseApiGatewayConfigPoolContent, len(model.Pools))
	for i, n := range model.Pools {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Pools[i] = nn.(DatabaseToolsDatabaseApiGatewayConfigPoolContent)
		} else {
			m.Pools[i] = nil
		}
	}
	nn, e = model.Global.UnmarshalPolymorphicJSON(model.Global.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Global = nn.(DatabaseToolsDatabaseApiGatewayConfigGlobal)
	} else {
		m.Global = nil
	}

	return
}
