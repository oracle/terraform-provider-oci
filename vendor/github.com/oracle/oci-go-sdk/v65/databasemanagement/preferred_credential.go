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

// PreferredCredential The details of the preferred credential.
type PreferredCredential interface {

	// The name of the preferred credential.
	GetCredentialName() *string

	// The status of the preferred credential.
	GetStatus() PreferredCredentialStatusEnum

	// Indicates whether the preferred credential is accessible.
	GetIsAccessible() *bool
}

type preferredcredential struct {
	JsonData       []byte
	CredentialName *string                       `mandatory:"false" json:"credentialName"`
	Status         PreferredCredentialStatusEnum `mandatory:"false" json:"status,omitempty"`
	IsAccessible   *bool                         `mandatory:"false" json:"isAccessible"`
	Type           string                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *preferredcredential) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpreferredcredential preferredcredential
	s := struct {
		Model Unmarshalerpreferredcredential
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialName = s.Model.CredentialName
	m.Status = s.Model.Status
	m.IsAccessible = s.Model.IsAccessible
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *preferredcredential) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "BASIC":
		mm := BasicPreferredCredential{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PreferredCredential: %s.", m.Type)
		return *m, nil
	}
}

// GetCredentialName returns CredentialName
func (m preferredcredential) GetCredentialName() *string {
	return m.CredentialName
}

// GetStatus returns Status
func (m preferredcredential) GetStatus() PreferredCredentialStatusEnum {
	return m.Status
}

// GetIsAccessible returns IsAccessible
func (m preferredcredential) GetIsAccessible() *bool {
	return m.IsAccessible
}

func (m preferredcredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m preferredcredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPreferredCredentialStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPreferredCredentialStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
