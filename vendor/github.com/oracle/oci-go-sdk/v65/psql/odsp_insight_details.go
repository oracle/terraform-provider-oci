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

// OdspInsightDetails ODSP Insight details for the database system.
type OdspInsightDetails interface {
}

type odspinsightdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *odspinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerodspinsightdetails odspinsightdetails
	s := struct {
		Model Unmarshalerodspinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *odspinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "DISABLED":
		mm := DisabledInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENABLED":
		mm := EnabledInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OdspInsightDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m odspinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m odspinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdspInsightDetailsKindEnum Enum with underlying type: string
type OdspInsightDetailsKindEnum string

// Set of constants representing the allowable values for OdspInsightDetailsKindEnum
const (
	OdspInsightDetailsKindEnabled  OdspInsightDetailsKindEnum = "ENABLED"
	OdspInsightDetailsKindDisabled OdspInsightDetailsKindEnum = "DISABLED"
)

var mappingOdspInsightDetailsKindEnum = map[string]OdspInsightDetailsKindEnum{
	"ENABLED":  OdspInsightDetailsKindEnabled,
	"DISABLED": OdspInsightDetailsKindDisabled,
}

var mappingOdspInsightDetailsKindEnumLowerCase = map[string]OdspInsightDetailsKindEnum{
	"enabled":  OdspInsightDetailsKindEnabled,
	"disabled": OdspInsightDetailsKindDisabled,
}

// GetOdspInsightDetailsKindEnumValues Enumerates the set of values for OdspInsightDetailsKindEnum
func GetOdspInsightDetailsKindEnumValues() []OdspInsightDetailsKindEnum {
	values := make([]OdspInsightDetailsKindEnum, 0)
	for _, v := range mappingOdspInsightDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetOdspInsightDetailsKindEnumStringValues Enumerates the set of values in String for OdspInsightDetailsKindEnum
func GetOdspInsightDetailsKindEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingOdspInsightDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdspInsightDetailsKindEnum(val string) (OdspInsightDetailsKindEnum, bool) {
	enum, ok := mappingOdspInsightDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
