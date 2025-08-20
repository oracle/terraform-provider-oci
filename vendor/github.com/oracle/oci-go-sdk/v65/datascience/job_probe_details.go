// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// JobProbeDetails The probe indicates whether the application within the job run has started.
type JobProbeDetails interface {
}

type jobprobedetails struct {
	JsonData          []byte
	JobProbeCheckType string `json:"jobProbeCheckType"`
}

// UnmarshalJSON unmarshals json
func (m *jobprobedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobprobedetails jobprobedetails
	s := struct {
		Model Unmarshalerjobprobedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobProbeCheckType = s.Model.JobProbeCheckType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobprobedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobProbeCheckType {
	case "EXEC":
		mm := JobExecProbeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for JobProbeDetails: %s.", m.JobProbeCheckType)
		return *m, nil
	}
}

func (m jobprobedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobprobedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobProbeDetailsJobProbeCheckTypeEnum Enum with underlying type: string
type JobProbeDetailsJobProbeCheckTypeEnum string

// Set of constants representing the allowable values for JobProbeDetailsJobProbeCheckTypeEnum
const (
	JobProbeDetailsJobProbeCheckTypeExec JobProbeDetailsJobProbeCheckTypeEnum = "EXEC"
)

var mappingJobProbeDetailsJobProbeCheckTypeEnum = map[string]JobProbeDetailsJobProbeCheckTypeEnum{
	"EXEC": JobProbeDetailsJobProbeCheckTypeExec,
}

var mappingJobProbeDetailsJobProbeCheckTypeEnumLowerCase = map[string]JobProbeDetailsJobProbeCheckTypeEnum{
	"exec": JobProbeDetailsJobProbeCheckTypeExec,
}

// GetJobProbeDetailsJobProbeCheckTypeEnumValues Enumerates the set of values for JobProbeDetailsJobProbeCheckTypeEnum
func GetJobProbeDetailsJobProbeCheckTypeEnumValues() []JobProbeDetailsJobProbeCheckTypeEnum {
	values := make([]JobProbeDetailsJobProbeCheckTypeEnum, 0)
	for _, v := range mappingJobProbeDetailsJobProbeCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobProbeDetailsJobProbeCheckTypeEnumStringValues Enumerates the set of values in String for JobProbeDetailsJobProbeCheckTypeEnum
func GetJobProbeDetailsJobProbeCheckTypeEnumStringValues() []string {
	return []string{
		"EXEC",
	}
}

// GetMappingJobProbeDetailsJobProbeCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobProbeDetailsJobProbeCheckTypeEnum(val string) (JobProbeDetailsJobProbeCheckTypeEnum, bool) {
	enum, ok := mappingJobProbeDetailsJobProbeCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
