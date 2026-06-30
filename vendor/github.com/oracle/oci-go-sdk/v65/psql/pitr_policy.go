// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PitrPolicy Point-in-time recovery policy.
type PitrPolicy interface {
}

type pitrpolicy struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *pitrpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpitrpolicy pitrpolicy
	s := struct {
		Model Unmarshalerpitrpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pitrpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STANDARD":
		mm := StandardPitrPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NonePitrPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PitrPolicy: %s.", m.Kind)
		return *m, nil
	}
}

func (m pitrpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pitrpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PitrPolicyKindEnum Enum with underlying type: string
type PitrPolicyKindEnum string

// Set of constants representing the allowable values for PitrPolicyKindEnum
const (
	PitrPolicyKindStandard PitrPolicyKindEnum = "STANDARD"
	PitrPolicyKindNone     PitrPolicyKindEnum = "NONE"
)

var mappingPitrPolicyKindEnum = map[string]PitrPolicyKindEnum{
	"STANDARD": PitrPolicyKindStandard,
	"NONE":     PitrPolicyKindNone,
}

var mappingPitrPolicyKindEnumLowerCase = map[string]PitrPolicyKindEnum{
	"standard": PitrPolicyKindStandard,
	"none":     PitrPolicyKindNone,
}

// GetPitrPolicyKindEnumValues Enumerates the set of values for PitrPolicyKindEnum
func GetPitrPolicyKindEnumValues() []PitrPolicyKindEnum {
	values := make([]PitrPolicyKindEnum, 0)
	for _, v := range mappingPitrPolicyKindEnum {
		values = append(values, v)
	}
	return values
}

// GetPitrPolicyKindEnumStringValues Enumerates the set of values in String for PitrPolicyKindEnum
func GetPitrPolicyKindEnumStringValues() []string {
	return []string{
		"STANDARD",
		"NONE",
	}
}

// GetMappingPitrPolicyKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPitrPolicyKindEnum(val string) (PitrPolicyKindEnum, bool) {
	enum, ok := mappingPitrPolicyKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
