// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DropTablespaceDetails The details required to drop a tablespace.
type DropTablespaceDetails struct {
	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// Specifies whether all the contents of the tablespace being dropped should be dropped.
	IsIncludingContents *bool `mandatory:"false" json:"isIncludingContents"`

	// Specifies whether all the associated data files of the tablespace being dropped should be dropped.
	IsDroppingDataFiles *bool `mandatory:"false" json:"isDroppingDataFiles"`

	// Specifies whether all the constraints on the tablespace being dropped should be dropped.
	IsCascadeConstraints *bool `mandatory:"false" json:"isCascadeConstraints"`
}

func (m DropTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DropTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DropTablespaceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsIncludingContents  *bool                            `json:"isIncludingContents"`
		IsDroppingDataFiles  *bool                            `json:"isDroppingDataFiles"`
		IsCascadeConstraints *bool                            `json:"isCascadeConstraints"`
		CredentialDetails    tablespaceadmincredentialdetails `json:"credentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsIncludingContents = model.IsIncludingContents

	m.IsDroppingDataFiles = model.IsDroppingDataFiles

	m.IsCascadeConstraints = model.IsCascadeConstraints

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(TablespaceAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	return
}
