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

// CreatePeerTargetDatabaseDetails The details used to register the peer database of a database already registered in Data Safe.
type CreatePeerTargetDatabaseDetails struct {
	DatabaseDetails DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// The display name of the peer target database in Data Safe. The name is modifiable and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the peer target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the Data Guard Association resource in which the database being registered is considered as peer database to the primary database.
	DataguardAssociationId *string `mandatory:"false" json:"dataguardAssociationId"`

	TlsConfig *TlsConfig `mandatory:"false" json:"tlsConfig"`
}

func (m CreatePeerTargetDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePeerTargetDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreatePeerTargetDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string         `json:"displayName"`
		Description            *string         `json:"description"`
		DataguardAssociationId *string         `json:"dataguardAssociationId"`
		TlsConfig              *TlsConfig      `json:"tlsConfig"`
		DatabaseDetails        databasedetails `json:"databaseDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.DataguardAssociationId = model.DataguardAssociationId

	m.TlsConfig = model.TlsConfig

	nn, e = model.DatabaseDetails.UnmarshalPolymorphicJSON(model.DatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseDetails = nn.(DatabaseDetails)
	} else {
		m.DatabaseDetails = nil
	}

	return
}
