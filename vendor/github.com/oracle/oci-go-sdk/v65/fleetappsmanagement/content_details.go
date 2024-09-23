// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContentDetails Content Source Details.
type ContentDetails interface {
}

type contentdetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *contentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontentdetails contentdetails
	s := struct {
		Model Unmarshalercontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *contentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_BUCKET":
		mm := ObjectStorageBucketContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ContentDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m contentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m contentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContentDetailsSourceTypeEnum Enum with underlying type: string
type ContentDetailsSourceTypeEnum string

// Set of constants representing the allowable values for ContentDetailsSourceTypeEnum
const (
	ContentDetailsSourceTypeObjectStorageBucket ContentDetailsSourceTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingContentDetailsSourceTypeEnum = map[string]ContentDetailsSourceTypeEnum{
	"OBJECT_STORAGE_BUCKET": ContentDetailsSourceTypeObjectStorageBucket,
}

var mappingContentDetailsSourceTypeEnumLowerCase = map[string]ContentDetailsSourceTypeEnum{
	"object_storage_bucket": ContentDetailsSourceTypeObjectStorageBucket,
}

// GetContentDetailsSourceTypeEnumValues Enumerates the set of values for ContentDetailsSourceTypeEnum
func GetContentDetailsSourceTypeEnumValues() []ContentDetailsSourceTypeEnum {
	values := make([]ContentDetailsSourceTypeEnum, 0)
	for _, v := range mappingContentDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContentDetailsSourceTypeEnumStringValues Enumerates the set of values in String for ContentDetailsSourceTypeEnum
func GetContentDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingContentDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContentDetailsSourceTypeEnum(val string) (ContentDetailsSourceTypeEnum, bool) {
	enum, ok := mappingContentDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
