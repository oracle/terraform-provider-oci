// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsKeyStore The details of the key store.
type DatabaseToolsKeyStore struct {

	// The key store type.
	KeyStoreType KeyStoreTypeEnum `mandatory:"false" json:"keyStoreType,omitempty"`

	KeyStoreContent DatabaseToolsKeyStoreContent `mandatory:"false" json:"keyStoreContent"`

	KeyStorePassword DatabaseToolsKeyStorePassword `mandatory:"false" json:"keyStorePassword"`
}

func (m DatabaseToolsKeyStore) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsKeyStore) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		KeyStoreType     KeyStoreTypeEnum              `json:"keyStoreType"`
		KeyStoreContent  databasetoolskeystorecontent  `json:"keyStoreContent"`
		KeyStorePassword databasetoolskeystorepassword `json:"keyStorePassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.KeyStoreType = model.KeyStoreType

	nn, e = model.KeyStoreContent.UnmarshalPolymorphicJSON(model.KeyStoreContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.KeyStoreContent = nn.(DatabaseToolsKeyStoreContent)
	} else {
		m.KeyStoreContent = nil
	}

	nn, e = model.KeyStorePassword.UnmarshalPolymorphicJSON(model.KeyStorePassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.KeyStorePassword = nn.(DatabaseToolsKeyStorePassword)
	} else {
		m.KeyStorePassword = nil
	}

	return
}
