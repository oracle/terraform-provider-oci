// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateHostDumpTransferDetails Optional additional properties for dump transfer in source or target host. Default kind is CURL
type CreateHostDumpTransferDetails interface {

	// Directory path to OCI SSL wallet location on Db server node.
	GetWalletLocation() *string
}

type createhostdumptransferdetails struct {
	JsonData       []byte
	WalletLocation *string `mandatory:"false" json:"walletLocation"`
	Kind           string  `json:"kind"`
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
	m.WalletLocation = s.Model.WalletLocation
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
		common.Logf("Recieved unsupported enum value for CreateHostDumpTransferDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetWalletLocation returns WalletLocation
func (m createhostdumptransferdetails) GetWalletLocation() *string {
	return m.WalletLocation
}

func (m createhostdumptransferdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createhostdumptransferdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateHostDumpTransferDetailsKindEnum Enum with underlying type: string
type CreateHostDumpTransferDetailsKindEnum string

// Set of constants representing the allowable values for CreateHostDumpTransferDetailsKindEnum
const (
	CreateHostDumpTransferDetailsKindCurl   CreateHostDumpTransferDetailsKindEnum = "CURL"
	CreateHostDumpTransferDetailsKindOciCli CreateHostDumpTransferDetailsKindEnum = "OCI_CLI"
)

var mappingCreateHostDumpTransferDetailsKindEnum = map[string]CreateHostDumpTransferDetailsKindEnum{
	"CURL":    CreateHostDumpTransferDetailsKindCurl,
	"OCI_CLI": CreateHostDumpTransferDetailsKindOciCli,
}

var mappingCreateHostDumpTransferDetailsKindEnumLowerCase = map[string]CreateHostDumpTransferDetailsKindEnum{
	"curl":    CreateHostDumpTransferDetailsKindCurl,
	"oci_cli": CreateHostDumpTransferDetailsKindOciCli,
}

// GetCreateHostDumpTransferDetailsKindEnumValues Enumerates the set of values for CreateHostDumpTransferDetailsKindEnum
func GetCreateHostDumpTransferDetailsKindEnumValues() []CreateHostDumpTransferDetailsKindEnum {
	values := make([]CreateHostDumpTransferDetailsKindEnum, 0)
	for _, v := range mappingCreateHostDumpTransferDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateHostDumpTransferDetailsKindEnumStringValues Enumerates the set of values in String for CreateHostDumpTransferDetailsKindEnum
func GetCreateHostDumpTransferDetailsKindEnumStringValues() []string {
	return []string{
		"CURL",
		"OCI_CLI",
	}
}

// GetMappingCreateHostDumpTransferDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateHostDumpTransferDetailsKindEnum(val string) (CreateHostDumpTransferDetailsKindEnum, bool) {
	enum, ok := mappingCreateHostDumpTransferDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
