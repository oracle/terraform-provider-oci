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

// OciFssExport An OCI File Storage Service (FSS) Export. Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/Export/ for more details.
type OciFssExport interface {
}

type ocifssexport struct {
	JsonData         []byte
	OciFssExportType string `json:"ociFssExportType"`
}

// UnmarshalJSON unmarshals json
func (m *ocifssexport) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerocifssexport ocifssexport
	s := struct {
		Model Unmarshalerocifssexport
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OciFssExportType = s.Model.OciFssExportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ocifssexport) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OciFssExportType {
	case "OCID":
		mm := OciFssExportId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OciFssExport: %s.", m.OciFssExportType)
		return *m, nil
	}
}

func (m ocifssexport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ocifssexport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciFssExportOciFssExportTypeEnum Enum with underlying type: string
type OciFssExportOciFssExportTypeEnum string

// Set of constants representing the allowable values for OciFssExportOciFssExportTypeEnum
const (
	OciFssExportOciFssExportTypeOcid OciFssExportOciFssExportTypeEnum = "OCID"
)

var mappingOciFssExportOciFssExportTypeEnum = map[string]OciFssExportOciFssExportTypeEnum{
	"OCID": OciFssExportOciFssExportTypeOcid,
}

var mappingOciFssExportOciFssExportTypeEnumLowerCase = map[string]OciFssExportOciFssExportTypeEnum{
	"ocid": OciFssExportOciFssExportTypeOcid,
}

// GetOciFssExportOciFssExportTypeEnumValues Enumerates the set of values for OciFssExportOciFssExportTypeEnum
func GetOciFssExportOciFssExportTypeEnumValues() []OciFssExportOciFssExportTypeEnum {
	values := make([]OciFssExportOciFssExportTypeEnum, 0)
	for _, v := range mappingOciFssExportOciFssExportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciFssExportOciFssExportTypeEnumStringValues Enumerates the set of values in String for OciFssExportOciFssExportTypeEnum
func GetOciFssExportOciFssExportTypeEnumStringValues() []string {
	return []string{
		"OCID",
	}
}

// GetMappingOciFssExportOciFssExportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciFssExportOciFssExportTypeEnum(val string) (OciFssExportOciFssExportTypeEnum, bool) {
	enum, ok := mappingOciFssExportOciFssExportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
