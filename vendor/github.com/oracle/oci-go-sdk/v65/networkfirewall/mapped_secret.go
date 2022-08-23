// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MappedSecret Mapped secret used on the firewall policy rules.
type MappedSecret interface {

	// Type of the secrets mapped based on the policy.
	//  * `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
	//  * `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection.
	GetType() MappedSecretTypeEnum
}

type mappedsecret struct {
	JsonData []byte
	Type     MappedSecretTypeEnum `mandatory:"true" json:"type"`
	Source   string               `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *mappedsecret) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermappedsecret mappedsecret
	s := struct {
		Model Unmarshalermappedsecret
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *mappedsecret) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "OCI_VAULT":
		mm := VaultMappedSecret{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MappedSecret: %s.", m.Source)
		return *m, nil
	}
}

//GetType returns Type
func (m mappedsecret) GetType() MappedSecretTypeEnum {
	return m.Type
}

func (m mappedsecret) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m mappedsecret) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMappedSecretTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMappedSecretTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MappedSecretTypeEnum Enum with underlying type: string
type MappedSecretTypeEnum string

// Set of constants representing the allowable values for MappedSecretTypeEnum
const (
	MappedSecretTypeInboundInspection MappedSecretTypeEnum = "SSL_INBOUND_INSPECTION"
	MappedSecretTypeForwardProxy      MappedSecretTypeEnum = "SSL_FORWARD_PROXY"
)

var mappingMappedSecretTypeEnum = map[string]MappedSecretTypeEnum{
	"SSL_INBOUND_INSPECTION": MappedSecretTypeInboundInspection,
	"SSL_FORWARD_PROXY":      MappedSecretTypeForwardProxy,
}

var mappingMappedSecretTypeEnumLowerCase = map[string]MappedSecretTypeEnum{
	"ssl_inbound_inspection": MappedSecretTypeInboundInspection,
	"ssl_forward_proxy":      MappedSecretTypeForwardProxy,
}

// GetMappedSecretTypeEnumValues Enumerates the set of values for MappedSecretTypeEnum
func GetMappedSecretTypeEnumValues() []MappedSecretTypeEnum {
	values := make([]MappedSecretTypeEnum, 0)
	for _, v := range mappingMappedSecretTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedSecretTypeEnumStringValues Enumerates the set of values in String for MappedSecretTypeEnum
func GetMappedSecretTypeEnumStringValues() []string {
	return []string{
		"SSL_INBOUND_INSPECTION",
		"SSL_FORWARD_PROXY",
	}
}

// GetMappingMappedSecretTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedSecretTypeEnum(val string) (MappedSecretTypeEnum, bool) {
	enum, ok := mappingMappedSecretTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
