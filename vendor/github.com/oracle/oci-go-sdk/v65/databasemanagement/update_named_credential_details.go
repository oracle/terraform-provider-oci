// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNamedCredentialDetails The details required to update a named credential.
type UpdateNamedCredentialDetails struct {

	// The information specified by the user about the named credential.
	Description *string `mandatory:"false" json:"description"`

	// The scope of the named credential.
	Scope NamedCredentialScopeEnum `mandatory:"false" json:"scope,omitempty"`

	Content NamedCredentialContent `mandatory:"false" json:"content"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource that
	// is associated to the named credential.
	AssociatedResource *string `mandatory:"false" json:"associatedResource"`
}

func (m UpdateNamedCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNamedCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNamedCredentialScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetNamedCredentialScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateNamedCredentialDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                  `json:"description"`
		Scope              NamedCredentialScopeEnum `json:"scope"`
		Content            namedcredentialcontent   `json:"content"`
		AssociatedResource *string                  `json:"associatedResource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Scope = model.Scope

	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(NamedCredentialContent)
	} else {
		m.Content = nil
	}

	m.AssociatedResource = model.AssociatedResource

	return
}
