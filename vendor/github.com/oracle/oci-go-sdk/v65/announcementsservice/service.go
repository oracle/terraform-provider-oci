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

// Service Summary of the service object.
type Service struct {

	// ID of the service object.
	Id *string `mandatory:"true" json:"id"`

	// Name of the service represented by this object.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Short name of the team to whom this service object is related.
	ShortName *string `mandatory:"true" json:"shortName"`

	// Team name to which this service object is related.
	TeamName *string `mandatory:"true" json:"teamName"`

	// The list of realms where this service is not available to be used.
	ExcludedRealms []string `mandatory:"true" json:"excludedRealms"`

	// The list of previously used names for this service object.
	PreviousServiceNames []string `mandatory:"true" json:"previousServiceNames"`

	// The date and time when the service object was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the service object was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The platform type this service object is related to.
	PlatformType PlatformTypeEnum `mandatory:"true" json:"platformType"`

	// Name of the comms manager team that manages Notifications to this service.
	CommsManagerName CommsManagerNameEnum `mandatory:"true" json:"commsManagerName"`

	// Current state of the service object.
	LifecycleState BaseServiceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m Service) GetId() *string {
	return m.Id
}

// GetServiceName returns ServiceName
func (m Service) GetServiceName() *string {
	return m.ServiceName
}

// GetShortName returns ShortName
func (m Service) GetShortName() *string {
	return m.ShortName
}

// GetTeamName returns TeamName
func (m Service) GetTeamName() *string {
	return m.TeamName
}

// GetPlatformType returns PlatformType
func (m Service) GetPlatformType() PlatformTypeEnum {
	return m.PlatformType
}

// GetCommsManagerName returns CommsManagerName
func (m Service) GetCommsManagerName() CommsManagerNameEnum {
	return m.CommsManagerName
}

// GetExcludedRealms returns ExcludedRealms
func (m Service) GetExcludedRealms() []string {
	return m.ExcludedRealms
}

// GetPreviousServiceNames returns PreviousServiceNames
func (m Service) GetPreviousServiceNames() []string {
	return m.PreviousServiceNames
}

// GetTimeCreated returns TimeCreated
func (m Service) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m Service) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m Service) GetLifecycleState() BaseServiceLifecycleStateEnum {
	return m.LifecycleState
}

func (m Service) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Service) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m Service) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeService Service
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeService
	}{
		"Service",
		(MarshalTypeService)(m),
	}

	return json.Marshal(&s)
}
