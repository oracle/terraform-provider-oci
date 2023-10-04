// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroup The details of a DR protection group.
type DrProtectionGroup struct {

	// The OCID of the DR protection group.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the DR protection group.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the DR protection group.
	// Example: `EBS PHX Group`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The role of the DR protection group.
	Role DrProtectionGroupRoleEnum `mandatory:"true" json:"role"`

	// The date and time the DR protection group was created. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the DR protection group was updated. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the DR protection group.
	LifecycleState DrProtectionGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the peer DR protection group.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	PeerId *string `mandatory:"false" json:"peerId"`

	// The region of the peer DR protection group.
	// Example: `us-ashburn-1`
	PeerRegion *string `mandatory:"false" json:"peerRegion"`

	LogLocation *ObjectStorageLogLocation `mandatory:"false" json:"logLocation"`

	// A list of DR protection group members.
	Members []DrProtectionGroupMember `mandatory:"false" json:"members"`

	// A message describing the DR protection group's current state in more detail.
	LifeCycleDetails *string `mandatory:"false" json:"lifeCycleDetails"`

	// The current sub-state of the DR protection group.
	LifecycleSubState DrProtectionGroupLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DrProtectionGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrProtectionGroupRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDrProtectionGroupRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrProtectionGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDrProtectionGroupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDrProtectionGroupLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetDrProtectionGroupLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DrProtectionGroup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PeerId            *string                                `json:"peerId"`
		PeerRegion        *string                                `json:"peerRegion"`
		LogLocation       *ObjectStorageLogLocation              `json:"logLocation"`
		Members           []drprotectiongroupmember              `json:"members"`
		LifeCycleDetails  *string                                `json:"lifeCycleDetails"`
		LifecycleSubState DrProtectionGroupLifecycleSubStateEnum `json:"lifecycleSubState"`
		FreeformTags      map[string]string                      `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{}      `json:"definedTags"`
		SystemTags        map[string]map[string]interface{}      `json:"systemTags"`
		Id                *string                                `json:"id"`
		CompartmentId     *string                                `json:"compartmentId"`
		DisplayName       *string                                `json:"displayName"`
		Role              DrProtectionGroupRoleEnum              `json:"role"`
		TimeCreated       *common.SDKTime                        `json:"timeCreated"`
		TimeUpdated       *common.SDKTime                        `json:"timeUpdated"`
		LifecycleState    DrProtectionGroupLifecycleStateEnum    `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PeerId = model.PeerId

	m.PeerRegion = model.PeerRegion

	m.LogLocation = model.LogLocation

	m.Members = make([]DrProtectionGroupMember, len(model.Members))
	for i, n := range model.Members {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Members[i] = nn.(DrProtectionGroupMember)
		} else {
			m.Members[i] = nil
		}
	}
	m.LifeCycleDetails = model.LifeCycleDetails

	m.LifecycleSubState = model.LifecycleSubState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Role = model.Role

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}
