// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DecryptionProfile Decryption Profile used on the firewall policy rules.
type DecryptionProfile interface {
}

type decryptionprofile struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *decryptionprofile) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdecryptionprofile decryptionprofile
	s := struct {
		Model Unmarshalerdecryptionprofile
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *decryptionprofile) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SSL_INBOUND_INSPECTION":
		mm := SslInboundInspectionProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SSL_FORWARD_PROXY":
		mm := SslForwardProxyProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m decryptionprofile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m decryptionprofile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DecryptionProfileTypeEnum Enum with underlying type: string
type DecryptionProfileTypeEnum string

// Set of constants representing the allowable values for DecryptionProfileTypeEnum
const (
	DecryptionProfileTypeInboundInspection DecryptionProfileTypeEnum = "SSL_INBOUND_INSPECTION"
	DecryptionProfileTypeForwardProxy      DecryptionProfileTypeEnum = "SSL_FORWARD_PROXY"
)

var mappingDecryptionProfileTypeEnum = map[string]DecryptionProfileTypeEnum{
	"SSL_INBOUND_INSPECTION": DecryptionProfileTypeInboundInspection,
	"SSL_FORWARD_PROXY":      DecryptionProfileTypeForwardProxy,
}

var mappingDecryptionProfileTypeEnumLowerCase = map[string]DecryptionProfileTypeEnum{
	"ssl_inbound_inspection": DecryptionProfileTypeInboundInspection,
	"ssl_forward_proxy":      DecryptionProfileTypeForwardProxy,
}

// GetDecryptionProfileTypeEnumValues Enumerates the set of values for DecryptionProfileTypeEnum
func GetDecryptionProfileTypeEnumValues() []DecryptionProfileTypeEnum {
	values := make([]DecryptionProfileTypeEnum, 0)
	for _, v := range mappingDecryptionProfileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDecryptionProfileTypeEnumStringValues Enumerates the set of values in String for DecryptionProfileTypeEnum
func GetDecryptionProfileTypeEnumStringValues() []string {
	return []string{
		"SSL_INBOUND_INSPECTION",
		"SSL_FORWARD_PROXY",
	}
}

// GetMappingDecryptionProfileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDecryptionProfileTypeEnum(val string) (DecryptionProfileTypeEnum, bool) {
	enum, ok := mappingDecryptionProfileTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
