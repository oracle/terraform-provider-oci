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

// UpdatePeerTargetDatabaseDetails The details of the peer database used for updating the peer target database in Data Safe.
type UpdatePeerTargetDatabaseDetails struct {

	// The display name of the peer target database in Data Safe.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the peer target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	DatabaseDetails DatabaseDetails `mandatory:"false" json:"databaseDetails"`

	TlsConfig *TlsConfig `mandatory:"false" json:"tlsConfig"`
}

func (m UpdatePeerTargetDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePeerTargetDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdatePeerTargetDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string         `json:"displayName"`
		Description     *string         `json:"description"`
		DatabaseDetails databasedetails `json:"databaseDetails"`
		TlsConfig       *TlsConfig      `json:"tlsConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.DatabaseDetails.UnmarshalPolymorphicJSON(model.DatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseDetails = nn.(DatabaseDetails)
	} else {
		m.DatabaseDetails = nil
	}

	m.TlsConfig = model.TlsConfig

	return
}
