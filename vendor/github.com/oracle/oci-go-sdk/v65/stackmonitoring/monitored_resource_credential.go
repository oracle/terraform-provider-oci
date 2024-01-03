// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoredResourceCredential Monitored Resource Credential Details.
type MonitoredResourceCredential interface {

	// The source type and source name combination, delimited with (.) separator.
	// {source type}.{source name} and source type max char limit is 63.
	GetSource() *string

	// The name of the credential, within the context of the source.
	GetName() *string

	// The type of the credential ( ex. JMXCreds,DBCreds).
	GetType() *string

	// The user-specified textual description of the credential.
	GetDescription() *string
}

type monitoredresourcecredential struct {
	JsonData       []byte
	Source         *string `mandatory:"false" json:"source"`
	Name           *string `mandatory:"false" json:"name"`
	Type           *string `mandatory:"false" json:"type"`
	Description    *string `mandatory:"false" json:"description"`
	CredentialType string  `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *monitoredresourcecredential) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitoredresourcecredential monitoredresourcecredential
	s := struct {
		Model Unmarshalermonitoredresourcecredential
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source
	m.Name = s.Model.Name
	m.Type = s.Model.Type
	m.Description = s.Model.Description
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitoredresourcecredential) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "EXISTING":
		mm := PreExistingCredentials{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENCRYPTED":
		mm := EncryptedCredentials{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLAINTEXT":
		mm := PlainTextCredentials{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MonitoredResourceCredential: %s.", m.CredentialType)
		return *m, nil
	}
}

// GetSource returns Source
func (m monitoredresourcecredential) GetSource() *string {
	return m.Source
}

// GetName returns Name
func (m monitoredresourcecredential) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m monitoredresourcecredential) GetType() *string {
	return m.Type
}

// GetDescription returns Description
func (m monitoredresourcecredential) GetDescription() *string {
	return m.Description
}

func (m monitoredresourcecredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m monitoredresourcecredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoredResourceCredentialCredentialTypeEnum Enum with underlying type: string
type MonitoredResourceCredentialCredentialTypeEnum string

// Set of constants representing the allowable values for MonitoredResourceCredentialCredentialTypeEnum
const (
	MonitoredResourceCredentialCredentialTypeExisting  MonitoredResourceCredentialCredentialTypeEnum = "EXISTING"
	MonitoredResourceCredentialCredentialTypePlaintext MonitoredResourceCredentialCredentialTypeEnum = "PLAINTEXT"
	MonitoredResourceCredentialCredentialTypeEncrypted MonitoredResourceCredentialCredentialTypeEnum = "ENCRYPTED"
)

var mappingMonitoredResourceCredentialCredentialTypeEnum = map[string]MonitoredResourceCredentialCredentialTypeEnum{
	"EXISTING":  MonitoredResourceCredentialCredentialTypeExisting,
	"PLAINTEXT": MonitoredResourceCredentialCredentialTypePlaintext,
	"ENCRYPTED": MonitoredResourceCredentialCredentialTypeEncrypted,
}

var mappingMonitoredResourceCredentialCredentialTypeEnumLowerCase = map[string]MonitoredResourceCredentialCredentialTypeEnum{
	"existing":  MonitoredResourceCredentialCredentialTypeExisting,
	"plaintext": MonitoredResourceCredentialCredentialTypePlaintext,
	"encrypted": MonitoredResourceCredentialCredentialTypeEncrypted,
}

// GetMonitoredResourceCredentialCredentialTypeEnumValues Enumerates the set of values for MonitoredResourceCredentialCredentialTypeEnum
func GetMonitoredResourceCredentialCredentialTypeEnumValues() []MonitoredResourceCredentialCredentialTypeEnum {
	values := make([]MonitoredResourceCredentialCredentialTypeEnum, 0)
	for _, v := range mappingMonitoredResourceCredentialCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredResourceCredentialCredentialTypeEnumStringValues Enumerates the set of values in String for MonitoredResourceCredentialCredentialTypeEnum
func GetMonitoredResourceCredentialCredentialTypeEnumStringValues() []string {
	return []string{
		"EXISTING",
		"PLAINTEXT",
		"ENCRYPTED",
	}
}

// GetMappingMonitoredResourceCredentialCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredResourceCredentialCredentialTypeEnum(val string) (MonitoredResourceCredentialCredentialTypeEnum, bool) {
	enum, ok := mappingMonitoredResourceCredentialCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
