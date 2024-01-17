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

// PeerTargetDatabase The details of the peer target database in Data Safe.
type PeerTargetDatabase struct {

	// The display name of the peer target database in Data Safe.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The secondary key assigned for the peer target database in Data Safe.
	Key *int `mandatory:"true" json:"key"`

	// The OCID of the Data Guard Association resource in which the database associated to the peer target database is considered as peer database to the primary database.
	DataguardAssociationId *string `mandatory:"true" json:"dataguardAssociationId"`

	// The date and time of the peer target database registration in Data Safe.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	DatabaseDetails DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// The current state of the peer target database in Data Safe.
	LifecycleState TargetDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the peer target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// Role of the database associated to the peer target database.
	Role *string `mandatory:"false" json:"role"`

	// Unique name of the database associated to the peer target database.
	DatabaseUniqueName *string `mandatory:"false" json:"databaseUniqueName"`

	TlsConfig *TlsConfig `mandatory:"false" json:"tlsConfig"`

	// Details about the current state of the peer target database in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m PeerTargetDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeerTargetDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PeerTargetDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                          `json:"description"`
		Role                   *string                          `json:"role"`
		DatabaseUniqueName     *string                          `json:"databaseUniqueName"`
		TlsConfig              *TlsConfig                       `json:"tlsConfig"`
		LifecycleDetails       *string                          `json:"lifecycleDetails"`
		DisplayName            *string                          `json:"displayName"`
		Key                    *int                             `json:"key"`
		DataguardAssociationId *string                          `json:"dataguardAssociationId"`
		TimeCreated            *common.SDKTime                  `json:"timeCreated"`
		DatabaseDetails        databasedetails                  `json:"databaseDetails"`
		LifecycleState         TargetDatabaseLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Role = model.Role

	m.DatabaseUniqueName = model.DatabaseUniqueName

	m.TlsConfig = model.TlsConfig

	m.LifecycleDetails = model.LifecycleDetails

	m.DisplayName = model.DisplayName

	m.Key = model.Key

	m.DataguardAssociationId = model.DataguardAssociationId

	m.TimeCreated = model.TimeCreated

	nn, e = model.DatabaseDetails.UnmarshalPolymorphicJSON(model.DatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseDetails = nn.(DatabaseDetails)
	} else {
		m.DatabaseDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	return
}
