// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCacheUser An OCI cache user is required to connect to an OCI cache cluster.
type OciCacheUser struct {

	// OCI Cache user unique ID.
	Id *string `mandatory:"true" json:"id"`

	// OCI Cache user name.
	Name *string `mandatory:"true" json:"name"`

	// OCI Cache user compartment ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	AuthenticationMode AuthenticationMode `mandatory:"true" json:"authenticationMode"`

	// ACL string of OCI cache user.
	AclString *string `mandatory:"true" json:"aclString"`

	// OCI Cache user status. ON enables and OFF disables the OCI cache user to login to the cluster.
	Status OciCacheUserStatusEnum `mandatory:"true" json:"status"`

	// OCI Cache user lifecycle state.
	LifecycleState OciCacheUserLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of OCI cache user.
	Description *string `mandatory:"false" json:"description"`

	// The date and time, when the OCI cache user was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time, when the OCI cache user was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OciCacheUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheUserStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOciCacheUserStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheUserLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheUserLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *OciCacheUser) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                           `json:"description"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		Id                 *string                           `json:"id"`
		Name               *string                           `json:"name"`
		CompartmentId      *string                           `json:"compartmentId"`
		AuthenticationMode authenticationmode                `json:"authenticationMode"`
		AclString          *string                           `json:"aclString"`
		Status             OciCacheUserStatusEnum            `json:"status"`
		LifecycleState     OciCacheUserLifecycleStateEnum    `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.AuthenticationMode.UnmarshalPolymorphicJSON(model.AuthenticationMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthenticationMode = nn.(AuthenticationMode)
	} else {
		m.AuthenticationMode = nil
	}

	m.AclString = model.AclString

	m.Status = model.Status

	m.LifecycleState = model.LifecycleState

	return
}

// OciCacheUserStatusEnum Enum with underlying type: string
type OciCacheUserStatusEnum string

// Set of constants representing the allowable values for OciCacheUserStatusEnum
const (
	OciCacheUserStatusOn  OciCacheUserStatusEnum = "ON"
	OciCacheUserStatusOff OciCacheUserStatusEnum = "OFF"
)

var mappingOciCacheUserStatusEnum = map[string]OciCacheUserStatusEnum{
	"ON":  OciCacheUserStatusOn,
	"OFF": OciCacheUserStatusOff,
}

var mappingOciCacheUserStatusEnumLowerCase = map[string]OciCacheUserStatusEnum{
	"on":  OciCacheUserStatusOn,
	"off": OciCacheUserStatusOff,
}

// GetOciCacheUserStatusEnumValues Enumerates the set of values for OciCacheUserStatusEnum
func GetOciCacheUserStatusEnumValues() []OciCacheUserStatusEnum {
	values := make([]OciCacheUserStatusEnum, 0)
	for _, v := range mappingOciCacheUserStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheUserStatusEnumStringValues Enumerates the set of values in String for OciCacheUserStatusEnum
func GetOciCacheUserStatusEnumStringValues() []string {
	return []string{
		"ON",
		"OFF",
	}
}

// GetMappingOciCacheUserStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheUserStatusEnum(val string) (OciCacheUserStatusEnum, bool) {
	enum, ok := mappingOciCacheUserStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OciCacheUserLifecycleStateEnum Enum with underlying type: string
type OciCacheUserLifecycleStateEnum string

// Set of constants representing the allowable values for OciCacheUserLifecycleStateEnum
const (
	OciCacheUserLifecycleStateCreating OciCacheUserLifecycleStateEnum = "CREATING"
	OciCacheUserLifecycleStateUpdating OciCacheUserLifecycleStateEnum = "UPDATING"
	OciCacheUserLifecycleStateActive   OciCacheUserLifecycleStateEnum = "ACTIVE"
	OciCacheUserLifecycleStateDeleting OciCacheUserLifecycleStateEnum = "DELETING"
	OciCacheUserLifecycleStateDeleted  OciCacheUserLifecycleStateEnum = "DELETED"
	OciCacheUserLifecycleStateFailed   OciCacheUserLifecycleStateEnum = "FAILED"
)

var mappingOciCacheUserLifecycleStateEnum = map[string]OciCacheUserLifecycleStateEnum{
	"CREATING": OciCacheUserLifecycleStateCreating,
	"UPDATING": OciCacheUserLifecycleStateUpdating,
	"ACTIVE":   OciCacheUserLifecycleStateActive,
	"DELETING": OciCacheUserLifecycleStateDeleting,
	"DELETED":  OciCacheUserLifecycleStateDeleted,
	"FAILED":   OciCacheUserLifecycleStateFailed,
}

var mappingOciCacheUserLifecycleStateEnumLowerCase = map[string]OciCacheUserLifecycleStateEnum{
	"creating": OciCacheUserLifecycleStateCreating,
	"updating": OciCacheUserLifecycleStateUpdating,
	"active":   OciCacheUserLifecycleStateActive,
	"deleting": OciCacheUserLifecycleStateDeleting,
	"deleted":  OciCacheUserLifecycleStateDeleted,
	"failed":   OciCacheUserLifecycleStateFailed,
}

// GetOciCacheUserLifecycleStateEnumValues Enumerates the set of values for OciCacheUserLifecycleStateEnum
func GetOciCacheUserLifecycleStateEnumValues() []OciCacheUserLifecycleStateEnum {
	values := make([]OciCacheUserLifecycleStateEnum, 0)
	for _, v := range mappingOciCacheUserLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheUserLifecycleStateEnumStringValues Enumerates the set of values in String for OciCacheUserLifecycleStateEnum
func GetOciCacheUserLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOciCacheUserLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheUserLifecycleStateEnum(val string) (OciCacheUserLifecycleStateEnum, bool) {
	enum, ok := mappingOciCacheUserLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
