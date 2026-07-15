// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciFssExportDetails An OCI File Storage Service (FSS) Export. Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/Export/ for more details.
type CreateOciFssExportDetails interface {
}

type createocifssexportdetails struct {
	JsonData         []byte
	OciFssExportType string `json:"ociFssExportType"`
}

// UnmarshalJSON unmarshals json
func (m *createocifssexportdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateocifssexportdetails createocifssexportdetails
	s := struct {
		Model Unmarshalercreateocifssexportdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OciFssExportType = s.Model.OciFssExportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createocifssexportdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OciFssExportType {
	case "OCID":
		mm := CreateOciFssExportIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateOciFssExportDetails: %s.", m.OciFssExportType)
		return *m, nil
	}
}

func (m createocifssexportdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createocifssexportdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOciFssExportDetailsOciFssExportTypeEnum Enum with underlying type: string
type CreateOciFssExportDetailsOciFssExportTypeEnum string

// Set of constants representing the allowable values for CreateOciFssExportDetailsOciFssExportTypeEnum
const (
	CreateOciFssExportDetailsOciFssExportTypeOcid CreateOciFssExportDetailsOciFssExportTypeEnum = "OCID"
)

var mappingCreateOciFssExportDetailsOciFssExportTypeEnum = map[string]CreateOciFssExportDetailsOciFssExportTypeEnum{
	"OCID": CreateOciFssExportDetailsOciFssExportTypeOcid,
}

var mappingCreateOciFssExportDetailsOciFssExportTypeEnumLowerCase = map[string]CreateOciFssExportDetailsOciFssExportTypeEnum{
	"ocid": CreateOciFssExportDetailsOciFssExportTypeOcid,
}

// GetCreateOciFssExportDetailsOciFssExportTypeEnumValues Enumerates the set of values for CreateOciFssExportDetailsOciFssExportTypeEnum
func GetCreateOciFssExportDetailsOciFssExportTypeEnumValues() []CreateOciFssExportDetailsOciFssExportTypeEnum {
	values := make([]CreateOciFssExportDetailsOciFssExportTypeEnum, 0)
	for _, v := range mappingCreateOciFssExportDetailsOciFssExportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOciFssExportDetailsOciFssExportTypeEnumStringValues Enumerates the set of values in String for CreateOciFssExportDetailsOciFssExportTypeEnum
func GetCreateOciFssExportDetailsOciFssExportTypeEnumStringValues() []string {
	return []string{
		"OCID",
	}
}

// GetMappingCreateOciFssExportDetailsOciFssExportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOciFssExportDetailsOciFssExportTypeEnum(val string) (CreateOciFssExportDetailsOciFssExportTypeEnum, bool) {
	enum, ok := mappingCreateOciFssExportDetailsOciFssExportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
