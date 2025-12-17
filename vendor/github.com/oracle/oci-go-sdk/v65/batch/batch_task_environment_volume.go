// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTaskEnvironmentVolume A base type for volume for batch task environment.
type BatchTaskEnvironmentVolume interface {
}

type batchtaskenvironmentvolume struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *batchtaskenvironmentvolume) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbatchtaskenvironmentvolume batchtaskenvironmentvolume
	s := struct {
		Model Unmarshalerbatchtaskenvironmentvolume
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *batchtaskenvironmentvolume) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NFS":
		mm := NfsVolume{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BatchTaskEnvironmentVolume: %s.", m.Type)
		return *m, nil
	}
}

func (m batchtaskenvironmentvolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m batchtaskenvironmentvolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchTaskEnvironmentVolumeTypeEnum Enum with underlying type: string
type BatchTaskEnvironmentVolumeTypeEnum string

// Set of constants representing the allowable values for BatchTaskEnvironmentVolumeTypeEnum
const (
	BatchTaskEnvironmentVolumeTypeNfs BatchTaskEnvironmentVolumeTypeEnum = "NFS"
)

var mappingBatchTaskEnvironmentVolumeTypeEnum = map[string]BatchTaskEnvironmentVolumeTypeEnum{
	"NFS": BatchTaskEnvironmentVolumeTypeNfs,
}

var mappingBatchTaskEnvironmentVolumeTypeEnumLowerCase = map[string]BatchTaskEnvironmentVolumeTypeEnum{
	"nfs": BatchTaskEnvironmentVolumeTypeNfs,
}

// GetBatchTaskEnvironmentVolumeTypeEnumValues Enumerates the set of values for BatchTaskEnvironmentVolumeTypeEnum
func GetBatchTaskEnvironmentVolumeTypeEnumValues() []BatchTaskEnvironmentVolumeTypeEnum {
	values := make([]BatchTaskEnvironmentVolumeTypeEnum, 0)
	for _, v := range mappingBatchTaskEnvironmentVolumeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskEnvironmentVolumeTypeEnumStringValues Enumerates the set of values in String for BatchTaskEnvironmentVolumeTypeEnum
func GetBatchTaskEnvironmentVolumeTypeEnumStringValues() []string {
	return []string{
		"NFS",
	}
}

// GetMappingBatchTaskEnvironmentVolumeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskEnvironmentVolumeTypeEnum(val string) (BatchTaskEnvironmentVolumeTypeEnum, bool) {
	enum, ok := mappingBatchTaskEnvironmentVolumeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
