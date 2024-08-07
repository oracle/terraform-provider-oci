// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StageOutput Details of the generated artifact or report.
type StageOutput interface {

	// Name of stage step at which this output is generated.
	GetStepName() *string
}

type stageoutput struct {
	JsonData   []byte
	StepName   *string `mandatory:"true" json:"stepName"`
	OutputType string  `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *stageoutput) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstageoutput stageoutput
	s := struct {
		Model Unmarshalerstageoutput
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepName = s.Model.StepName
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *stageoutput) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "TEST_REPORT":
		mm := TestReportOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ARTIFACT":
		mm := GenericArtifacts{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StageOutput: %s.", m.OutputType)
		return *m, nil
	}
}

// GetStepName returns StepName
func (m stageoutput) GetStepName() *string {
	return m.StepName
}

func (m stageoutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m stageoutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StageOutputOutputTypeEnum Enum with underlying type: string
type StageOutputOutputTypeEnum string

// Set of constants representing the allowable values for StageOutputOutputTypeEnum
const (
	StageOutputOutputTypeArtifact   StageOutputOutputTypeEnum = "ARTIFACT"
	StageOutputOutputTypeTestReport StageOutputOutputTypeEnum = "TEST_REPORT"
)

var mappingStageOutputOutputTypeEnum = map[string]StageOutputOutputTypeEnum{
	"ARTIFACT":    StageOutputOutputTypeArtifact,
	"TEST_REPORT": StageOutputOutputTypeTestReport,
}

var mappingStageOutputOutputTypeEnumLowerCase = map[string]StageOutputOutputTypeEnum{
	"artifact":    StageOutputOutputTypeArtifact,
	"test_report": StageOutputOutputTypeTestReport,
}

// GetStageOutputOutputTypeEnumValues Enumerates the set of values for StageOutputOutputTypeEnum
func GetStageOutputOutputTypeEnumValues() []StageOutputOutputTypeEnum {
	values := make([]StageOutputOutputTypeEnum, 0)
	for _, v := range mappingStageOutputOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStageOutputOutputTypeEnumStringValues Enumerates the set of values in String for StageOutputOutputTypeEnum
func GetStageOutputOutputTypeEnumStringValues() []string {
	return []string{
		"ARTIFACT",
		"TEST_REPORT",
	}
}

// GetMappingStageOutputOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStageOutputOutputTypeEnum(val string) (StageOutputOutputTypeEnum, bool) {
	enum, ok := mappingStageOutputOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
