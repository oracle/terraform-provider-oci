// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunbookImportAs Runbook import as type.
type RunbookImportAs interface {

	// ImportOptions for Runbook.
	GetImportOption() ImportRunbookDetailsImportOptionEnum

	// Version number.
	GetVersion() *string
}

type runbookimportas struct {
	JsonData     []byte
	Version      *string                              `mandatory:"false" json:"version"`
	ImportOption ImportRunbookDetailsImportOptionEnum `mandatory:"true" json:"importOption"`
	ImportType   string                               `json:"importType"`
}

// UnmarshalJSON unmarshals json
func (m *runbookimportas) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrunbookimportas runbookimportas
	s := struct {
		Model Unmarshalerrunbookimportas
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ImportOption = s.Model.ImportOption
	m.Version = s.Model.Version
	m.ImportType = s.Model.ImportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *runbookimportas) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ImportType {
	case "RUNBOOK_VERSION":
		mm := RunbookImportAsVersion{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUNBOOK":
		mm := RunbookImportAsRunbook{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RunbookImportAs: %s.", m.ImportType)
		return *m, nil
	}
}

// GetVersion returns Version
func (m runbookimportas) GetVersion() *string {
	return m.Version
}

// GetImportOption returns ImportOption
func (m runbookimportas) GetImportOption() ImportRunbookDetailsImportOptionEnum {
	return m.ImportOption
}

func (m runbookimportas) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m runbookimportas) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportRunbookDetailsImportOptionEnum(string(m.ImportOption)); !ok && m.ImportOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImportOption: %s. Supported values are: %s.", m.ImportOption, strings.Join(GetImportRunbookDetailsImportOptionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunbookImportAsImportTypeEnum Enum with underlying type: string
type RunbookImportAsImportTypeEnum string

// Set of constants representing the allowable values for RunbookImportAsImportTypeEnum
const (
	RunbookImportAsImportTypeRunbook        RunbookImportAsImportTypeEnum = "RUNBOOK"
	RunbookImportAsImportTypeRunbookVersion RunbookImportAsImportTypeEnum = "RUNBOOK_VERSION"
)

var mappingRunbookImportAsImportTypeEnum = map[string]RunbookImportAsImportTypeEnum{
	"RUNBOOK":         RunbookImportAsImportTypeRunbook,
	"RUNBOOK_VERSION": RunbookImportAsImportTypeRunbookVersion,
}

var mappingRunbookImportAsImportTypeEnumLowerCase = map[string]RunbookImportAsImportTypeEnum{
	"runbook":         RunbookImportAsImportTypeRunbook,
	"runbook_version": RunbookImportAsImportTypeRunbookVersion,
}

// GetRunbookImportAsImportTypeEnumValues Enumerates the set of values for RunbookImportAsImportTypeEnum
func GetRunbookImportAsImportTypeEnumValues() []RunbookImportAsImportTypeEnum {
	values := make([]RunbookImportAsImportTypeEnum, 0)
	for _, v := range mappingRunbookImportAsImportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookImportAsImportTypeEnumStringValues Enumerates the set of values in String for RunbookImportAsImportTypeEnum
func GetRunbookImportAsImportTypeEnumStringValues() []string {
	return []string{
		"RUNBOOK",
		"RUNBOOK_VERSION",
	}
}

// GetMappingRunbookImportAsImportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookImportAsImportTypeEnum(val string) (RunbookImportAsImportTypeEnum, bool) {
	enum, ok := mappingRunbookImportAsImportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
