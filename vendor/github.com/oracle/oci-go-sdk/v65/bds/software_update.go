// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareUpdate Details about the given software update.
type SoftwareUpdate interface {

	// Unique identifier of a given software update
	GetSoftwareUpdateKey() *string

	// The version of the software update.
	GetSoftwareUpdateVersion() *string

	// The time when the software update was released.
	GetTimeReleased() *common.SDKTime

	// The lifecycle state of the software update.
	GetLifecycleState() SoftwareUpdateLifecycleStateEnum
}

type softwareupdate struct {
	JsonData              []byte
	SoftwareUpdateKey     *string                          `mandatory:"true" json:"softwareUpdateKey"`
	SoftwareUpdateVersion *string                          `mandatory:"true" json:"softwareUpdateVersion"`
	TimeReleased          *common.SDKTime                  `mandatory:"true" json:"timeReleased"`
	LifecycleState        SoftwareUpdateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	SoftwareUpdateType    string                           `json:"softwareUpdateType"`
}

// UnmarshalJSON unmarshals json
func (m *softwareupdate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersoftwareupdate softwareupdate
	s := struct {
		Model Unmarshalersoftwareupdate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SoftwareUpdateKey = s.Model.SoftwareUpdateKey
	m.SoftwareUpdateVersion = s.Model.SoftwareUpdateVersion
	m.TimeReleased = s.Model.TimeReleased
	m.LifecycleState = s.Model.LifecycleState
	m.SoftwareUpdateType = s.Model.SoftwareUpdateType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *softwareupdate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SoftwareUpdateType {
	case "BDS":
		mm := BdsSoftwareUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SoftwareUpdate: %s.", m.SoftwareUpdateType)
		return *m, nil
	}
}

// GetSoftwareUpdateKey returns SoftwareUpdateKey
func (m softwareupdate) GetSoftwareUpdateKey() *string {
	return m.SoftwareUpdateKey
}

// GetSoftwareUpdateVersion returns SoftwareUpdateVersion
func (m softwareupdate) GetSoftwareUpdateVersion() *string {
	return m.SoftwareUpdateVersion
}

// GetTimeReleased returns TimeReleased
func (m softwareupdate) GetTimeReleased() *common.SDKTime {
	return m.TimeReleased
}

// GetLifecycleState returns LifecycleState
func (m softwareupdate) GetLifecycleState() SoftwareUpdateLifecycleStateEnum {
	return m.LifecycleState
}

func (m softwareupdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m softwareupdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareUpdateLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareUpdateLifecycleStateEnum Enum with underlying type: string
type SoftwareUpdateLifecycleStateEnum string

// Set of constants representing the allowable values for SoftwareUpdateLifecycleStateEnum
const (
	SoftwareUpdateLifecycleStateWaiting    SoftwareUpdateLifecycleStateEnum = "WAITING"
	SoftwareUpdateLifecycleStateInProgress SoftwareUpdateLifecycleStateEnum = "IN_PROGRESS"
	SoftwareUpdateLifecycleStateSucceeded  SoftwareUpdateLifecycleStateEnum = "SUCCEEDED"
	SoftwareUpdateLifecycleStateFailed     SoftwareUpdateLifecycleStateEnum = "FAILED"
	SoftwareUpdateLifecycleStateCanceled   SoftwareUpdateLifecycleStateEnum = "CANCELED"
)

var mappingSoftwareUpdateLifecycleStateEnum = map[string]SoftwareUpdateLifecycleStateEnum{
	"WAITING":     SoftwareUpdateLifecycleStateWaiting,
	"IN_PROGRESS": SoftwareUpdateLifecycleStateInProgress,
	"SUCCEEDED":   SoftwareUpdateLifecycleStateSucceeded,
	"FAILED":      SoftwareUpdateLifecycleStateFailed,
	"CANCELED":    SoftwareUpdateLifecycleStateCanceled,
}

var mappingSoftwareUpdateLifecycleStateEnumLowerCase = map[string]SoftwareUpdateLifecycleStateEnum{
	"waiting":     SoftwareUpdateLifecycleStateWaiting,
	"in_progress": SoftwareUpdateLifecycleStateInProgress,
	"succeeded":   SoftwareUpdateLifecycleStateSucceeded,
	"failed":      SoftwareUpdateLifecycleStateFailed,
	"canceled":    SoftwareUpdateLifecycleStateCanceled,
}

// GetSoftwareUpdateLifecycleStateEnumValues Enumerates the set of values for SoftwareUpdateLifecycleStateEnum
func GetSoftwareUpdateLifecycleStateEnumValues() []SoftwareUpdateLifecycleStateEnum {
	values := make([]SoftwareUpdateLifecycleStateEnum, 0)
	for _, v := range mappingSoftwareUpdateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareUpdateLifecycleStateEnumStringValues Enumerates the set of values in String for SoftwareUpdateLifecycleStateEnum
func GetSoftwareUpdateLifecycleStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

// GetMappingSoftwareUpdateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareUpdateLifecycleStateEnum(val string) (SoftwareUpdateLifecycleStateEnum, bool) {
	enum, ok := mappingSoftwareUpdateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SoftwareUpdateSoftwareUpdateTypeEnum Enum with underlying type: string
type SoftwareUpdateSoftwareUpdateTypeEnum string

// Set of constants representing the allowable values for SoftwareUpdateSoftwareUpdateTypeEnum
const (
	SoftwareUpdateSoftwareUpdateTypeBds SoftwareUpdateSoftwareUpdateTypeEnum = "BDS"
)

var mappingSoftwareUpdateSoftwareUpdateTypeEnum = map[string]SoftwareUpdateSoftwareUpdateTypeEnum{
	"BDS": SoftwareUpdateSoftwareUpdateTypeBds,
}

var mappingSoftwareUpdateSoftwareUpdateTypeEnumLowerCase = map[string]SoftwareUpdateSoftwareUpdateTypeEnum{
	"bds": SoftwareUpdateSoftwareUpdateTypeBds,
}

// GetSoftwareUpdateSoftwareUpdateTypeEnumValues Enumerates the set of values for SoftwareUpdateSoftwareUpdateTypeEnum
func GetSoftwareUpdateSoftwareUpdateTypeEnumValues() []SoftwareUpdateSoftwareUpdateTypeEnum {
	values := make([]SoftwareUpdateSoftwareUpdateTypeEnum, 0)
	for _, v := range mappingSoftwareUpdateSoftwareUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareUpdateSoftwareUpdateTypeEnumStringValues Enumerates the set of values in String for SoftwareUpdateSoftwareUpdateTypeEnum
func GetSoftwareUpdateSoftwareUpdateTypeEnumStringValues() []string {
	return []string{
		"BDS",
	}
}

// GetMappingSoftwareUpdateSoftwareUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareUpdateSoftwareUpdateTypeEnum(val string) (SoftwareUpdateSoftwareUpdateTypeEnum, bool) {
	enum, ok := mappingSoftwareUpdateSoftwareUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
