// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisassociateDrProtectionGroupDetails The details for disassociating this DR protection group from a peer DR protection group.
type DisassociateDrProtectionGroupDetails interface {
}

type disassociatedrprotectiongroupdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *disassociatedrprotectiongroupdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdisassociatedrprotectiongroupdetails disassociatedrprotectiongroupdetails
	s := struct {
		Model Unmarshalerdisassociatedrprotectiongroupdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *disassociatedrprotectiongroupdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DisassociateDrProtectionGroupDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DisassociateDrProtectionGroupDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m disassociatedrprotectiongroupdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m disassociatedrprotectiongroupdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisassociateDrProtectionGroupDetailsTypeEnum Enum with underlying type: string
type DisassociateDrProtectionGroupDetailsTypeEnum string

// Set of constants representing the allowable values for DisassociateDrProtectionGroupDetailsTypeEnum
const (
	DisassociateDrProtectionGroupDetailsTypeDefault DisassociateDrProtectionGroupDetailsTypeEnum = "DEFAULT"
)

var mappingDisassociateDrProtectionGroupDetailsTypeEnum = map[string]DisassociateDrProtectionGroupDetailsTypeEnum{
	"DEFAULT": DisassociateDrProtectionGroupDetailsTypeDefault,
}

var mappingDisassociateDrProtectionGroupDetailsTypeEnumLowerCase = map[string]DisassociateDrProtectionGroupDetailsTypeEnum{
	"default": DisassociateDrProtectionGroupDetailsTypeDefault,
}

// GetDisassociateDrProtectionGroupDetailsTypeEnumValues Enumerates the set of values for DisassociateDrProtectionGroupDetailsTypeEnum
func GetDisassociateDrProtectionGroupDetailsTypeEnumValues() []DisassociateDrProtectionGroupDetailsTypeEnum {
	values := make([]DisassociateDrProtectionGroupDetailsTypeEnum, 0)
	for _, v := range mappingDisassociateDrProtectionGroupDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDisassociateDrProtectionGroupDetailsTypeEnumStringValues Enumerates the set of values in String for DisassociateDrProtectionGroupDetailsTypeEnum
func GetDisassociateDrProtectionGroupDetailsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingDisassociateDrProtectionGroupDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisassociateDrProtectionGroupDetailsTypeEnum(val string) (DisassociateDrProtectionGroupDetailsTypeEnum, bool) {
	enum, ok := mappingDisassociateDrProtectionGroupDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
