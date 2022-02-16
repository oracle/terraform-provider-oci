// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateHostDumpTransferDetails Optional additional properties for dump transfer in source or target host. Default kind is CURL
type UpdateHostDumpTransferDetails interface {
}

type updatehostdumptransferdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *updatehostdumptransferdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatehostdumptransferdetails updatehostdumptransferdetails
	s := struct {
		Model Unmarshalerupdatehostdumptransferdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatehostdumptransferdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "CURL":
		mm := UpdateCurlTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_CLI":
		mm := UpdateOciCliDumpTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updatehostdumptransferdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatehostdumptransferdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateHostDumpTransferDetailsKindEnum Enum with underlying type: string
type UpdateHostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for UpdateHostDumpTransferDetailsKindEnum
const (
	UpdateHostDumpTransferDetailsKindCurl   UpdateHostDumpTransferDetailsKindEnum = "CURL"
	UpdateHostDumpTransferDetailsKindOciCli UpdateHostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingUpdateHostDumpTransferDetailsKindEnum = map[string]UpdateHostDumpTransferDetailsKindEnum{
	"CURL":    UpdateHostDumpTransferDetailsKindCurl,
	"OCI_CLI": UpdateHostDumpTransferDetailsKindOciCli,
}

// GetUpdateHostDumpTransferDetailsKindEnumValues Enumerates the set of values for UpdateHostDumpTransferDetailsKindEnum
func GetUpdateHostDumpTransferDetailsKindEnumValues() []UpdateHostDumpTransferDetailsKindEnum {
	values := make([]UpdateHostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingUpdateHostDumpTransferDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHostDumpTransferDetailsKindEnumStringValues Enumerates the set of values in String for UpdateHostDumpTransferDetailsKindEnum
func GetUpdateHostDumpTransferDetailsKindEnumStringValues() []string {
	return []string{
		"CURL",
		"OCI_CLI",
	}
}

// GetMappingUpdateHostDumpTransferDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHostDumpTransferDetailsKindEnum(val string) (UpdateHostDumpTransferDetailsKindEnum, bool) {
	mappingUpdateHostDumpTransferDetailsKindEnumIgnoreCase := make(map[string]UpdateHostDumpTransferDetailsKindEnum)
	for k, v := range mappingUpdateHostDumpTransferDetailsKindEnum {
		mappingUpdateHostDumpTransferDetailsKindEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateHostDumpTransferDetailsKindEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
