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

// UpdateHostDumpTransferDetailsKindEnum Enum with underlying type: string
type UpdateHostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for UpdateHostDumpTransferDetailsKindEnum
const (
	UpdateHostDumpTransferDetailsKindCurl   UpdateHostDumpTransferDetailsKindEnum = "CURL"
	UpdateHostDumpTransferDetailsKindOciCli UpdateHostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingUpdateHostDumpTransferDetailsKind = map[string]UpdateHostDumpTransferDetailsKindEnum{
	"CURL":    UpdateHostDumpTransferDetailsKindCurl,
	"OCI_CLI": UpdateHostDumpTransferDetailsKindOciCli,
}

// GetUpdateHostDumpTransferDetailsKindEnumValues Enumerates the set of values for UpdateHostDumpTransferDetailsKindEnum
func GetUpdateHostDumpTransferDetailsKindEnumValues() []UpdateHostDumpTransferDetailsKindEnum {
	values := make([]UpdateHostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingUpdateHostDumpTransferDetailsKind {
		values = append(values, v)
	}
	return values
}
