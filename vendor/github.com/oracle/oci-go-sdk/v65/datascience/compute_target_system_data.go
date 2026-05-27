// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeTargetSystemData System data of the compute target.
type ComputeTargetSystemData interface {
}

type computetargetsystemdata struct {
	JsonData    []byte
	ComputeType string `json:"computeType"`
}

// UnmarshalJSON unmarshals json
func (m *computetargetsystemdata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercomputetargetsystemdata computetargetsystemdata
	s := struct {
		Model Unmarshalercomputetargetsystemdata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComputeType = s.Model.ComputeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *computetargetsystemdata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComputeType {
	case "MANAGED_COMPUTE_CLUSTER":
		mm := ManagedComputeClusterSystemData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ComputeTargetSystemData: %s.", m.ComputeType)
		return *m, nil
	}
}

func (m computetargetsystemdata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m computetargetsystemdata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeTargetSystemDataComputeTypeEnum Enum with underlying type: string
type ComputeTargetSystemDataComputeTypeEnum string

// Set of constants representing the allowable values for ComputeTargetSystemDataComputeTypeEnum
const (
	ComputeTargetSystemDataComputeTypeManagedComputeCluster ComputeTargetSystemDataComputeTypeEnum = "MANAGED_COMPUTE_CLUSTER"
)

var mappingComputeTargetSystemDataComputeTypeEnum = map[string]ComputeTargetSystemDataComputeTypeEnum{
	"MANAGED_COMPUTE_CLUSTER": ComputeTargetSystemDataComputeTypeManagedComputeCluster,
}

var mappingComputeTargetSystemDataComputeTypeEnumLowerCase = map[string]ComputeTargetSystemDataComputeTypeEnum{
	"managed_compute_cluster": ComputeTargetSystemDataComputeTypeManagedComputeCluster,
}

// GetComputeTargetSystemDataComputeTypeEnumValues Enumerates the set of values for ComputeTargetSystemDataComputeTypeEnum
func GetComputeTargetSystemDataComputeTypeEnumValues() []ComputeTargetSystemDataComputeTypeEnum {
	values := make([]ComputeTargetSystemDataComputeTypeEnum, 0)
	for _, v := range mappingComputeTargetSystemDataComputeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeTargetSystemDataComputeTypeEnumStringValues Enumerates the set of values in String for ComputeTargetSystemDataComputeTypeEnum
func GetComputeTargetSystemDataComputeTypeEnumStringValues() []string {
	return []string{
		"MANAGED_COMPUTE_CLUSTER",
	}
}

// GetMappingComputeTargetSystemDataComputeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeTargetSystemDataComputeTypeEnum(val string) (ComputeTargetSystemDataComputeTypeEnum, bool) {
	enum, ok := mappingComputeTargetSystemDataComputeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
