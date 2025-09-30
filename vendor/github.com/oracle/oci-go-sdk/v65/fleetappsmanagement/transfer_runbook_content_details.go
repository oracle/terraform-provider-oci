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

// TransferRunbookContentDetails Content Source details.
type TransferRunbookContentDetails interface {
}

type transferrunbookcontentdetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *transferrunbookcontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertransferrunbookcontentdetails transferrunbookcontentdetails
	s := struct {
		Model Unmarshalertransferrunbookcontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *transferrunbookcontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_BUCKET":
		mm := TransferRunbookObjectStorageBucketContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PAR_URL":
		mm := TransferRunbookParUrlContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TransferRunbookContentDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m transferrunbookcontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m transferrunbookcontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferRunbookContentDetailsSourceTypeEnum Enum with underlying type: string
type TransferRunbookContentDetailsSourceTypeEnum string

// Set of constants representing the allowable values for TransferRunbookContentDetailsSourceTypeEnum
const (
	TransferRunbookContentDetailsSourceTypeObjectStorageBucket TransferRunbookContentDetailsSourceTypeEnum = "OBJECT_STORAGE_BUCKET"
	TransferRunbookContentDetailsSourceTypeParUrl              TransferRunbookContentDetailsSourceTypeEnum = "PAR_URL"
)

var mappingTransferRunbookContentDetailsSourceTypeEnum = map[string]TransferRunbookContentDetailsSourceTypeEnum{
	"OBJECT_STORAGE_BUCKET": TransferRunbookContentDetailsSourceTypeObjectStorageBucket,
	"PAR_URL":               TransferRunbookContentDetailsSourceTypeParUrl,
}

var mappingTransferRunbookContentDetailsSourceTypeEnumLowerCase = map[string]TransferRunbookContentDetailsSourceTypeEnum{
	"object_storage_bucket": TransferRunbookContentDetailsSourceTypeObjectStorageBucket,
	"par_url":               TransferRunbookContentDetailsSourceTypeParUrl,
}

// GetTransferRunbookContentDetailsSourceTypeEnumValues Enumerates the set of values for TransferRunbookContentDetailsSourceTypeEnum
func GetTransferRunbookContentDetailsSourceTypeEnumValues() []TransferRunbookContentDetailsSourceTypeEnum {
	values := make([]TransferRunbookContentDetailsSourceTypeEnum, 0)
	for _, v := range mappingTransferRunbookContentDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferRunbookContentDetailsSourceTypeEnumStringValues Enumerates the set of values in String for TransferRunbookContentDetailsSourceTypeEnum
func GetTransferRunbookContentDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
		"PAR_URL",
	}
}

// GetMappingTransferRunbookContentDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferRunbookContentDetailsSourceTypeEnum(val string) (TransferRunbookContentDetailsSourceTypeEnum, bool) {
	enum, ok := mappingTransferRunbookContentDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
