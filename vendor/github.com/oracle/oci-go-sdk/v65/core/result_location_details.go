// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResultLocationDetails A base object for all types of upload locations for a firmware report.
type ResultLocationDetails interface {
}

type resultlocationdetails struct {
	JsonData     []byte
	LocationType string `json:"locationType"`
}

// UnmarshalJSON unmarshals json
func (m *resultlocationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresultlocationdetails resultlocationdetails
	s := struct {
		Model Unmarshalerresultlocationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resultlocationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageResultLocationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResultLocationDetails: %s.", m.LocationType)
		return *m, nil
	}
}

func (m resultlocationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resultlocationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResultLocationDetailsLocationTypeEnum Enum with underlying type: string
type ResultLocationDetailsLocationTypeEnum string

// Set of constants representing the allowable values for ResultLocationDetailsLocationTypeEnum
const (
	ResultLocationDetailsLocationTypeObjectStorage ResultLocationDetailsLocationTypeEnum = "OBJECT_STORAGE"
)

var mappingResultLocationDetailsLocationTypeEnum = map[string]ResultLocationDetailsLocationTypeEnum{
	"OBJECT_STORAGE": ResultLocationDetailsLocationTypeObjectStorage,
}

var mappingResultLocationDetailsLocationTypeEnumLowerCase = map[string]ResultLocationDetailsLocationTypeEnum{
	"object_storage": ResultLocationDetailsLocationTypeObjectStorage,
}

// GetResultLocationDetailsLocationTypeEnumValues Enumerates the set of values for ResultLocationDetailsLocationTypeEnum
func GetResultLocationDetailsLocationTypeEnumValues() []ResultLocationDetailsLocationTypeEnum {
	values := make([]ResultLocationDetailsLocationTypeEnum, 0)
	for _, v := range mappingResultLocationDetailsLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResultLocationDetailsLocationTypeEnumStringValues Enumerates the set of values in String for ResultLocationDetailsLocationTypeEnum
func GetResultLocationDetailsLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingResultLocationDetailsLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResultLocationDetailsLocationTypeEnum(val string) (ResultLocationDetailsLocationTypeEnum, bool) {
	enum, ok := mappingResultLocationDetailsLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
