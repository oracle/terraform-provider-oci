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

// KerberosAuthDetails Kerberos Authentication details for the database system.
type KerberosAuthDetails interface {
}

type kerberosauthdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *kerberosauthdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkerberosauthdetails kerberosauthdetails
	s := struct {
		Model Unmarshalerkerberosauthdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *kerberosauthdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "DISABLED":
		mm := DisabledKerberosAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENABLED":
		mm := EnabledKerberosAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KerberosAuthDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m kerberosauthdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m kerberosauthdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KerberosAuthDetailsKindEnum Enum with underlying type: string
type KerberosAuthDetailsKindEnum string

// Set of constants representing the allowable values for KerberosAuthDetailsKindEnum
const (
	KerberosAuthDetailsKindEnabled  KerberosAuthDetailsKindEnum = "ENABLED"
	KerberosAuthDetailsKindDisabled KerberosAuthDetailsKindEnum = "DISABLED"
)

var mappingKerberosAuthDetailsKindEnum = map[string]KerberosAuthDetailsKindEnum{
	"ENABLED":  KerberosAuthDetailsKindEnabled,
	"DISABLED": KerberosAuthDetailsKindDisabled,
}

var mappingKerberosAuthDetailsKindEnumLowerCase = map[string]KerberosAuthDetailsKindEnum{
	"enabled":  KerberosAuthDetailsKindEnabled,
	"disabled": KerberosAuthDetailsKindDisabled,
}

// GetKerberosAuthDetailsKindEnumValues Enumerates the set of values for KerberosAuthDetailsKindEnum
func GetKerberosAuthDetailsKindEnumValues() []KerberosAuthDetailsKindEnum {
	values := make([]KerberosAuthDetailsKindEnum, 0)
	for _, v := range mappingKerberosAuthDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetKerberosAuthDetailsKindEnumStringValues Enumerates the set of values in String for KerberosAuthDetailsKindEnum
func GetKerberosAuthDetailsKindEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingKerberosAuthDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKerberosAuthDetailsKindEnum(val string) (KerberosAuthDetailsKindEnum, bool) {
	enum, ok := mappingKerberosAuthDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
