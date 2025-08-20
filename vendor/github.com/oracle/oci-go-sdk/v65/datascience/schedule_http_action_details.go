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

// ScheduleHttpActionDetails Schedule Http action details
type ScheduleHttpActionDetails interface {
}

type schedulehttpactiondetails struct {
	JsonData       []byte
	HttpActionType string `json:"httpActionType"`
}

// UnmarshalJSON unmarshals json
func (m *schedulehttpactiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschedulehttpactiondetails schedulehttpactiondetails
	s := struct {
		Model Unmarshalerschedulehttpactiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.HttpActionType = s.Model.HttpActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schedulehttpactiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.HttpActionType {
	case "CREATE_JOB_RUN":
		mm := CreateJobRunScheduleActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_ML_APPLICATION_PROVIDER_TRIGGER":
		mm := InvokeMlApplicationProviderTriggerScheduleActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREATE_PIPELINE_RUN":
		mm := CreatePipelineRunScheduleActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ScheduleHttpActionDetails: %s.", m.HttpActionType)
		return *m, nil
	}
}

func (m schedulehttpactiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m schedulehttpactiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleHttpActionDetailsHttpActionTypeEnum Enum with underlying type: string
type ScheduleHttpActionDetailsHttpActionTypeEnum string

// Set of constants representing the allowable values for ScheduleHttpActionDetailsHttpActionTypeEnum
const (
	ScheduleHttpActionDetailsHttpActionTypeCreateJobRun                       ScheduleHttpActionDetailsHttpActionTypeEnum = "CREATE_JOB_RUN"
	ScheduleHttpActionDetailsHttpActionTypeCreatePipelineRun                  ScheduleHttpActionDetailsHttpActionTypeEnum = "CREATE_PIPELINE_RUN"
	ScheduleHttpActionDetailsHttpActionTypeInvokeMlApplicationProviderTrigger ScheduleHttpActionDetailsHttpActionTypeEnum = "INVOKE_ML_APPLICATION_PROVIDER_TRIGGER"
)

var mappingScheduleHttpActionDetailsHttpActionTypeEnum = map[string]ScheduleHttpActionDetailsHttpActionTypeEnum{
	"CREATE_JOB_RUN":                         ScheduleHttpActionDetailsHttpActionTypeCreateJobRun,
	"CREATE_PIPELINE_RUN":                    ScheduleHttpActionDetailsHttpActionTypeCreatePipelineRun,
	"INVOKE_ML_APPLICATION_PROVIDER_TRIGGER": ScheduleHttpActionDetailsHttpActionTypeInvokeMlApplicationProviderTrigger,
}

var mappingScheduleHttpActionDetailsHttpActionTypeEnumLowerCase = map[string]ScheduleHttpActionDetailsHttpActionTypeEnum{
	"create_job_run":                         ScheduleHttpActionDetailsHttpActionTypeCreateJobRun,
	"create_pipeline_run":                    ScheduleHttpActionDetailsHttpActionTypeCreatePipelineRun,
	"invoke_ml_application_provider_trigger": ScheduleHttpActionDetailsHttpActionTypeInvokeMlApplicationProviderTrigger,
}

// GetScheduleHttpActionDetailsHttpActionTypeEnumValues Enumerates the set of values for ScheduleHttpActionDetailsHttpActionTypeEnum
func GetScheduleHttpActionDetailsHttpActionTypeEnumValues() []ScheduleHttpActionDetailsHttpActionTypeEnum {
	values := make([]ScheduleHttpActionDetailsHttpActionTypeEnum, 0)
	for _, v := range mappingScheduleHttpActionDetailsHttpActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleHttpActionDetailsHttpActionTypeEnumStringValues Enumerates the set of values in String for ScheduleHttpActionDetailsHttpActionTypeEnum
func GetScheduleHttpActionDetailsHttpActionTypeEnumStringValues() []string {
	return []string{
		"CREATE_JOB_RUN",
		"CREATE_PIPELINE_RUN",
		"INVOKE_ML_APPLICATION_PROVIDER_TRIGGER",
	}
}

// GetMappingScheduleHttpActionDetailsHttpActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleHttpActionDetailsHttpActionTypeEnum(val string) (ScheduleHttpActionDetailsHttpActionTypeEnum, bool) {
	enum, ok := mappingScheduleHttpActionDetailsHttpActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
