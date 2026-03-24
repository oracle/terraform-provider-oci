// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleConfig The Schedule configuration of a VectorStoreConnector to trigger a File Sync Operation.
type ScheduleConfig interface {

	// The state of the schedule. The state can be either ENABLED or DISABLED.
	GetState() ScheduleConfigStateEnum

	// The schedule starting date time, if null, System set the time when schedule is created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeStart() *common.SDKTime

	// The schedule end date time, if null, the schedule will never expire.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeEnd() *common.SDKTime
}

type scheduleconfig struct {
	JsonData   []byte
	State      ScheduleConfigStateEnum `mandatory:"false" json:"state,omitempty"`
	TimeStart  *common.SDKTime         `mandatory:"false" json:"timeStart"`
	TimeEnd    *common.SDKTime         `mandatory:"false" json:"timeEnd"`
	ConfigType string                  `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *scheduleconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduleconfig scheduleconfig
	s := struct {
		Model Unmarshalerscheduleconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.State = s.Model.State
	m.TimeStart = s.Model.TimeStart
	m.TimeEnd = s.Model.TimeEnd
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduleconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "CRON":
		mm := ScheduleCronConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTERVAL":
		mm := ScheduleIntervalConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ScheduleConfig: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetState returns State
func (m scheduleconfig) GetState() ScheduleConfigStateEnum {
	return m.State
}

// GetTimeStart returns TimeStart
func (m scheduleconfig) GetTimeStart() *common.SDKTime {
	return m.TimeStart
}

// GetTimeEnd returns TimeEnd
func (m scheduleconfig) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m scheduleconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduleconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScheduleConfigStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetScheduleConfigStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleConfigStateEnum Enum with underlying type: string
type ScheduleConfigStateEnum string

// Set of constants representing the allowable values for ScheduleConfigStateEnum
const (
	ScheduleConfigStateEnabled  ScheduleConfigStateEnum = "ENABLED"
	ScheduleConfigStateDisabled ScheduleConfigStateEnum = "DISABLED"
)

var mappingScheduleConfigStateEnum = map[string]ScheduleConfigStateEnum{
	"ENABLED":  ScheduleConfigStateEnabled,
	"DISABLED": ScheduleConfigStateDisabled,
}

var mappingScheduleConfigStateEnumLowerCase = map[string]ScheduleConfigStateEnum{
	"enabled":  ScheduleConfigStateEnabled,
	"disabled": ScheduleConfigStateDisabled,
}

// GetScheduleConfigStateEnumValues Enumerates the set of values for ScheduleConfigStateEnum
func GetScheduleConfigStateEnumValues() []ScheduleConfigStateEnum {
	values := make([]ScheduleConfigStateEnum, 0)
	for _, v := range mappingScheduleConfigStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleConfigStateEnumStringValues Enumerates the set of values in String for ScheduleConfigStateEnum
func GetScheduleConfigStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingScheduleConfigStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleConfigStateEnum(val string) (ScheduleConfigStateEnum, bool) {
	enum, ok := mappingScheduleConfigStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduleConfigConfigTypeEnum Enum with underlying type: string
type ScheduleConfigConfigTypeEnum string

// Set of constants representing the allowable values for ScheduleConfigConfigTypeEnum
const (
	ScheduleConfigConfigTypeInterval ScheduleConfigConfigTypeEnum = "INTERVAL"
	ScheduleConfigConfigTypeCron     ScheduleConfigConfigTypeEnum = "CRON"
)

var mappingScheduleConfigConfigTypeEnum = map[string]ScheduleConfigConfigTypeEnum{
	"INTERVAL": ScheduleConfigConfigTypeInterval,
	"CRON":     ScheduleConfigConfigTypeCron,
}

var mappingScheduleConfigConfigTypeEnumLowerCase = map[string]ScheduleConfigConfigTypeEnum{
	"interval": ScheduleConfigConfigTypeInterval,
	"cron":     ScheduleConfigConfigTypeCron,
}

// GetScheduleConfigConfigTypeEnumValues Enumerates the set of values for ScheduleConfigConfigTypeEnum
func GetScheduleConfigConfigTypeEnumValues() []ScheduleConfigConfigTypeEnum {
	values := make([]ScheduleConfigConfigTypeEnum, 0)
	for _, v := range mappingScheduleConfigConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleConfigConfigTypeEnumStringValues Enumerates the set of values in String for ScheduleConfigConfigTypeEnum
func GetScheduleConfigConfigTypeEnumStringValues() []string {
	return []string{
		"INTERVAL",
		"CRON",
	}
}

// GetMappingScheduleConfigConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleConfigConfigTypeEnum(val string) (ScheduleConfigConfigTypeEnum, bool) {
	enum, ok := mappingScheduleConfigConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
