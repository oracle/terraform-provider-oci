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

// CreateHostDumpTransferDetails Optional additional properties for dump transfer in source or target host. Default kind is CURL
type CreateHostDumpTransferDetails interface {
}

type createhostdumptransferdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *createhostdumptransferdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatehostdumptransferdetails createhostdumptransferdetails
	s := struct {
		Model Unmarshalercreatehostdumptransferdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createhostdumptransferdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "OCI_CLI":
		mm := CreateOciCliDumpTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CURL":
		mm := CreateCurlTransferDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createhostdumptransferdetails) String() string {
	return common.PointerString(m)
}

// CreateHostDumpTransferDetailsKindEnum Enum with underlying type: string
type CreateHostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for CreateHostDumpTransferDetailsKindEnum
const (
	CreateHostDumpTransferDetailsKindCurl   CreateHostDumpTransferDetailsKindEnum = "CURL"
	CreateHostDumpTransferDetailsKindOciCli CreateHostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingCreateHostDumpTransferDetailsKind = map[string]CreateHostDumpTransferDetailsKindEnum{
	"CURL":    CreateHostDumpTransferDetailsKindCurl,
	"OCI_CLI": CreateHostDumpTransferDetailsKindOciCli,
}

// GetCreateHostDumpTransferDetailsKindEnumValues Enumerates the set of values for CreateHostDumpTransferDetailsKindEnum
func GetCreateHostDumpTransferDetailsKindEnumValues() []CreateHostDumpTransferDetailsKindEnum {
	values := make([]CreateHostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingCreateHostDumpTransferDetailsKind {
		values = append(values, v)
	}
	return values
}
