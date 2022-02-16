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

// HostDumpTransferDetails Optional additional properties for dump transfer in source or target host. Default kind is CURL
type HostDumpTransferDetails interface {
}

type hostdumptransferdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *hostdumptransferdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostdumptransferdetails hostdumptransferdetails
	s := struct {
		Model Unmarshalerhostdumptransferdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostdumptransferdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "OCI_CLI":
		mm := OciCliDumpTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CURL":
		mm := CurlTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m hostdumptransferdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m hostdumptransferdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostDumpTransferDetailsKindEnum Enum with underlying type: string
type HostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for HostDumpTransferDetailsKindEnum
const (
	HostDumpTransferDetailsKindCurl   HostDumpTransferDetailsKindEnum = "CURL"
	HostDumpTransferDetailsKindOciCli HostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingHostDumpTransferDetailsKindEnum = map[string]HostDumpTransferDetailsKindEnum{
	"CURL":    HostDumpTransferDetailsKindCurl,
	"OCI_CLI": HostDumpTransferDetailsKindOciCli,
}

// GetHostDumpTransferDetailsKindEnumValues Enumerates the set of values for HostDumpTransferDetailsKindEnum
func GetHostDumpTransferDetailsKindEnumValues() []HostDumpTransferDetailsKindEnum {
	values := make([]HostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingHostDumpTransferDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetHostDumpTransferDetailsKindEnumStringValues Enumerates the set of values in String for HostDumpTransferDetailsKindEnum
func GetHostDumpTransferDetailsKindEnumStringValues() []string {
	return []string{
		"CURL",
		"OCI_CLI",
	}
}

// GetMappingHostDumpTransferDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostDumpTransferDetailsKindEnum(val string) (HostDumpTransferDetailsKindEnum, bool) {
	mappingHostDumpTransferDetailsKindEnumIgnoreCase := make(map[string]HostDumpTransferDetailsKindEnum)
	for k, v := range mappingHostDumpTransferDetailsKindEnum {
		mappingHostDumpTransferDetailsKindEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHostDumpTransferDetailsKindEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
