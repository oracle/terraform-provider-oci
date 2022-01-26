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
	"github.com/oracle/oci-go-sdk/v56/common"
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

// HostDumpTransferDetailsKindEnum Enum with underlying type: string
type HostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for HostDumpTransferDetailsKindEnum
const (
	HostDumpTransferDetailsKindCurl   HostDumpTransferDetailsKindEnum = "CURL"
	HostDumpTransferDetailsKindOciCli HostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingHostDumpTransferDetailsKind = map[string]HostDumpTransferDetailsKindEnum{
	"CURL":    HostDumpTransferDetailsKindCurl,
	"OCI_CLI": HostDumpTransferDetailsKindOciCli,
}

// GetHostDumpTransferDetailsKindEnumValues Enumerates the set of values for HostDumpTransferDetailsKindEnum
func GetHostDumpTransferDetailsKindEnumValues() []HostDumpTransferDetailsKindEnum {
	values := make([]HostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingHostDumpTransferDetailsKind {
		values = append(values, v)
	}
	return values
}
