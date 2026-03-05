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

// ScheduleIntervalConfig The interval schedule config.
type ScheduleIntervalConfig struct {

	// The interval of frequency.
	Interval *int `mandatory:"true" json:"interval"`

	// The schedule starting date time, if null, System set the time when schedule is created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The schedule end date time, if null, the schedule will never expire.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The type of frequency
	Frequency ScheduleIntervalConfigFrequencyEnum `mandatory:"true" json:"frequency"`

	// The state of the schedule. The state can be either ENABLED or DISABLED.
	State ScheduleConfigStateEnum `mandatory:"false" json:"state,omitempty"`
}

// GetState returns State
func (m ScheduleIntervalConfig) GetState() ScheduleConfigStateEnum {
	return m.State
}

// GetTimeStart returns TimeStart
func (m ScheduleIntervalConfig) GetTimeStart() *common.SDKTime {
	return m.TimeStart
}

// GetTimeEnd returns TimeEnd
func (m ScheduleIntervalConfig) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m ScheduleIntervalConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleIntervalConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleIntervalConfigFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetScheduleIntervalConfigFrequencyEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScheduleConfigStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetScheduleConfigStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleIntervalConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleIntervalConfig ScheduleIntervalConfig
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScheduleIntervalConfig
	}{
		"INTERVAL",
		(MarshalTypeScheduleIntervalConfig)(m),
	}

	return json.Marshal(&s)
}

// ScheduleIntervalConfigFrequencyEnum Enum with underlying type: string
type ScheduleIntervalConfigFrequencyEnum string

// Set of constants representing the allowable values for ScheduleIntervalConfigFrequencyEnum
const (
	ScheduleIntervalConfigFrequencyHourly ScheduleIntervalConfigFrequencyEnum = "HOURLY"
	ScheduleIntervalConfigFrequencyDaily  ScheduleIntervalConfigFrequencyEnum = "DAILY"
)

var mappingScheduleIntervalConfigFrequencyEnum = map[string]ScheduleIntervalConfigFrequencyEnum{
	"HOURLY": ScheduleIntervalConfigFrequencyHourly,
	"DAILY":  ScheduleIntervalConfigFrequencyDaily,
}

var mappingScheduleIntervalConfigFrequencyEnumLowerCase = map[string]ScheduleIntervalConfigFrequencyEnum{
	"hourly": ScheduleIntervalConfigFrequencyHourly,
	"daily":  ScheduleIntervalConfigFrequencyDaily,
}

// GetScheduleIntervalConfigFrequencyEnumValues Enumerates the set of values for ScheduleIntervalConfigFrequencyEnum
func GetScheduleIntervalConfigFrequencyEnumValues() []ScheduleIntervalConfigFrequencyEnum {
	values := make([]ScheduleIntervalConfigFrequencyEnum, 0)
	for _, v := range mappingScheduleIntervalConfigFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleIntervalConfigFrequencyEnumStringValues Enumerates the set of values in String for ScheduleIntervalConfigFrequencyEnum
func GetScheduleIntervalConfigFrequencyEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
	}
}

// GetMappingScheduleIntervalConfigFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleIntervalConfigFrequencyEnum(val string) (ScheduleIntervalConfigFrequencyEnum, bool) {
	enum, ok := mappingScheduleIntervalConfigFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
