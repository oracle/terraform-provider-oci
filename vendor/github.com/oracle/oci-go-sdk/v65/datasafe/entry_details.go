// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// EntryDetails Details specific to the security policy entry.
type EntryDetails interface {
}

type entrydetails struct {
	JsonData  []byte
	EntryType string `json:"entryType"`
}

// UnmarshalJSON unmarshals json
func (m *entrydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentrydetails entrydetails
	s := struct {
		Model Unmarshalerentrydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntryType = s.Model.EntryType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *entrydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntryType {
	case "FIREWALL_POLICY":
		mm := FirewallPolicyEntryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for EntryDetails: %s.", m.EntryType)
		return *m, nil
	}
}

func (m entrydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entrydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EntryDetailsEntryTypeEnum Enum with underlying type: string
type EntryDetailsEntryTypeEnum string

// Set of constants representing the allowable values for EntryDetailsEntryTypeEnum
const (
	EntryDetailsEntryTypeFirewallPolicy EntryDetailsEntryTypeEnum = "FIREWALL_POLICY"
)

var mappingEntryDetailsEntryTypeEnum = map[string]EntryDetailsEntryTypeEnum{
	"FIREWALL_POLICY": EntryDetailsEntryTypeFirewallPolicy,
}

var mappingEntryDetailsEntryTypeEnumLowerCase = map[string]EntryDetailsEntryTypeEnum{
	"firewall_policy": EntryDetailsEntryTypeFirewallPolicy,
}

// GetEntryDetailsEntryTypeEnumValues Enumerates the set of values for EntryDetailsEntryTypeEnum
func GetEntryDetailsEntryTypeEnumValues() []EntryDetailsEntryTypeEnum {
	values := make([]EntryDetailsEntryTypeEnum, 0)
	for _, v := range mappingEntryDetailsEntryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntryDetailsEntryTypeEnumStringValues Enumerates the set of values in String for EntryDetailsEntryTypeEnum
func GetEntryDetailsEntryTypeEnumStringValues() []string {
	return []string{
		"FIREWALL_POLICY",
	}
}

// GetMappingEntryDetailsEntryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntryDetailsEntryTypeEnum(val string) (EntryDetailsEntryTypeEnum, bool) {
	enum, ok := mappingEntryDetailsEntryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
