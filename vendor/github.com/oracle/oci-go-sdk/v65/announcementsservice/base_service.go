// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseService Object representing a single service.
type BaseService interface {

	// ID of the service object.
	GetId() *string

	// Name of the service represented by this object.
	GetServiceName() *string

	// Short name of the team to whom this service object is related.
	GetShortName() *string

	// Team name to which this service object is related.
	GetTeamName() *string

	// The platform type this service object is related to.
	GetPlatformType() PlatformTypeEnum

	// Name of the comms manager team that manages Notifications to this service.
	GetCommsManagerName() CommsManagerNameEnum

	// The list of realms where this service is not available to be used.
	GetExcludedRealms() []string

	// The list of previously used names for this service object.
	GetPreviousServiceNames() []string

	// The date and time when the service object was created.
	GetTimeCreated() *common.SDKTime

	// The date and time when the service object was updated.
	GetTimeUpdated() *common.SDKTime

	// Current state of the service object.
	GetLifecycleState() BaseServiceLifecycleStateEnum
}

type baseservice struct {
	JsonData             []byte
	TimeCreated          *common.SDKTime               `mandatory:"false" json:"timeCreated"`
	TimeUpdated          *common.SDKTime               `mandatory:"false" json:"timeUpdated"`
	LifecycleState       BaseServiceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	Id                   *string                       `mandatory:"true" json:"id"`
	ServiceName          *string                       `mandatory:"true" json:"serviceName"`
	ShortName            *string                       `mandatory:"true" json:"shortName"`
	TeamName             *string                       `mandatory:"true" json:"teamName"`
	PlatformType         PlatformTypeEnum              `mandatory:"true" json:"platformType"`
	CommsManagerName     CommsManagerNameEnum          `mandatory:"true" json:"commsManagerName"`
	ExcludedRealms       []string                      `mandatory:"true" json:"excludedRealms"`
	PreviousServiceNames []string                      `mandatory:"true" json:"previousServiceNames"`
	Type                 string                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseservice) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseservice baseservice
	s := struct {
		Model Unmarshalerbaseservice
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ServiceName = s.Model.ServiceName
	m.ShortName = s.Model.ShortName
	m.TeamName = s.Model.TeamName
	m.PlatformType = s.Model.PlatformType
	m.CommsManagerName = s.Model.CommsManagerName
	m.ExcludedRealms = s.Model.ExcludedRealms
	m.PreviousServiceNames = s.Model.PreviousServiceNames
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseservice) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "Service":
		mm := Service{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ServiceSummary":
		mm := ServiceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NotificationsSummary":
		mm := NotificationsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BaseService: %s.", m.Type)
		return *m, nil
	}
}

// GetTimeCreated returns TimeCreated
func (m baseservice) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m baseservice) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m baseservice) GetLifecycleState() BaseServiceLifecycleStateEnum {
	return m.LifecycleState
}

// GetId returns Id
func (m baseservice) GetId() *string {
	return m.Id
}

// GetServiceName returns ServiceName
func (m baseservice) GetServiceName() *string {
	return m.ServiceName
}

// GetShortName returns ShortName
func (m baseservice) GetShortName() *string {
	return m.ShortName
}

// GetTeamName returns TeamName
func (m baseservice) GetTeamName() *string {
	return m.TeamName
}

// GetPlatformType returns PlatformType
func (m baseservice) GetPlatformType() PlatformTypeEnum {
	return m.PlatformType
}

// GetCommsManagerName returns CommsManagerName
func (m baseservice) GetCommsManagerName() CommsManagerNameEnum {
	return m.CommsManagerName
}

// GetExcludedRealms returns ExcludedRealms
func (m baseservice) GetExcludedRealms() []string {
	return m.ExcludedRealms
}

// GetPreviousServiceNames returns PreviousServiceNames
func (m baseservice) GetPreviousServiceNames() []string {
	return m.PreviousServiceNames
}

func (m baseservice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m baseservice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPlatformTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCommsManagerNameEnum(string(m.CommsManagerName)); !ok && m.CommsManagerName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommsManagerName: %s. Supported values are: %s.", m.CommsManagerName, strings.Join(GetCommsManagerNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseServiceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseServiceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseServiceLifecycleStateEnum Enum with underlying type: string
type BaseServiceLifecycleStateEnum string

// Set of constants representing the allowable values for BaseServiceLifecycleStateEnum
const (
	BaseServiceLifecycleStateActive  BaseServiceLifecycleStateEnum = "ACTIVE"
	BaseServiceLifecycleStateDeleted BaseServiceLifecycleStateEnum = "DELETED"
)

var mappingBaseServiceLifecycleStateEnum = map[string]BaseServiceLifecycleStateEnum{
	"ACTIVE":  BaseServiceLifecycleStateActive,
	"DELETED": BaseServiceLifecycleStateDeleted,
}

var mappingBaseServiceLifecycleStateEnumLowerCase = map[string]BaseServiceLifecycleStateEnum{
	"active":  BaseServiceLifecycleStateActive,
	"deleted": BaseServiceLifecycleStateDeleted,
}

// GetBaseServiceLifecycleStateEnumValues Enumerates the set of values for BaseServiceLifecycleStateEnum
func GetBaseServiceLifecycleStateEnumValues() []BaseServiceLifecycleStateEnum {
	values := make([]BaseServiceLifecycleStateEnum, 0)
	for _, v := range mappingBaseServiceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseServiceLifecycleStateEnumStringValues Enumerates the set of values in String for BaseServiceLifecycleStateEnum
func GetBaseServiceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingBaseServiceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseServiceLifecycleStateEnum(val string) (BaseServiceLifecycleStateEnum, bool) {
	enum, ok := mappingBaseServiceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
