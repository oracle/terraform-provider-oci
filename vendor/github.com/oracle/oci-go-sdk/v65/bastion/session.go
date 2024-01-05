// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Session A bastion session resource. A bastion session lets authorized users connect to a target resource using a Secure Shell (SSH) for a predetermined amount of time.
type Session struct {

	// The unique identifier (OCID) of the session, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier (OCID) of the bastion that is hosting this session.
	BastionId *string `mandatory:"true" json:"bastionId"`

	// The name of the bastion that is hosting this session.
	BastionName *string `mandatory:"true" json:"bastionName"`

	TargetResourceDetails TargetResourceDetails `mandatory:"true" json:"targetResourceDetails"`

	KeyDetails *PublicKeyDetails `mandatory:"true" json:"keyDetails"`

	// The time the session was created. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the session.
	LifecycleState SessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The amount of time the session can remain active.
	SessionTtlInSeconds *int `mandatory:"true" json:"sessionTtlInSeconds"`

	// The name of the session.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The username that the session uses to connect to the target resource.
	BastionUserName *string `mandatory:"false" json:"bastionUserName"`

	// The connection message for the session.
	SshMetadata map[string]string `mandatory:"false" json:"sshMetadata"`

	// The type of the key used to connect to the session. PUB is a standard public key in OpenSSH format.
	KeyType SessionKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	// The public key of the bastion host. You can use this to verify that you're connecting to the correct bastion.
	BastionPublicHostKeyInfo *string `mandatory:"false" json:"bastionPublicHostKeyInfo"`

	// The time the session was updated. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current session state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m Session) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Session) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSessionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSessionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSessionKeyTypeEnum(string(m.KeyType)); !ok && m.KeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyType: %s. Supported values are: %s.", m.KeyType, strings.Join(GetSessionKeyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Session) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                   `json:"displayName"`
		BastionUserName          *string                   `json:"bastionUserName"`
		SshMetadata              map[string]string         `json:"sshMetadata"`
		KeyType                  SessionKeyTypeEnum        `json:"keyType"`
		BastionPublicHostKeyInfo *string                   `json:"bastionPublicHostKeyInfo"`
		TimeUpdated              *common.SDKTime           `json:"timeUpdated"`
		LifecycleDetails         *string                   `json:"lifecycleDetails"`
		Id                       *string                   `json:"id"`
		BastionId                *string                   `json:"bastionId"`
		BastionName              *string                   `json:"bastionName"`
		TargetResourceDetails    targetresourcedetails     `json:"targetResourceDetails"`
		KeyDetails               *PublicKeyDetails         `json:"keyDetails"`
		TimeCreated              *common.SDKTime           `json:"timeCreated"`
		LifecycleState           SessionLifecycleStateEnum `json:"lifecycleState"`
		SessionTtlInSeconds      *int                      `json:"sessionTtlInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.BastionUserName = model.BastionUserName

	m.SshMetadata = model.SshMetadata

	m.KeyType = model.KeyType

	m.BastionPublicHostKeyInfo = model.BastionPublicHostKeyInfo

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.BastionId = model.BastionId

	m.BastionName = model.BastionName

	nn, e = model.TargetResourceDetails.UnmarshalPolymorphicJSON(model.TargetResourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetResourceDetails = nn.(TargetResourceDetails)
	} else {
		m.TargetResourceDetails = nil
	}

	m.KeyDetails = model.KeyDetails

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SessionTtlInSeconds = model.SessionTtlInSeconds

	return
}

// SessionKeyTypeEnum Enum with underlying type: string
type SessionKeyTypeEnum string

// Set of constants representing the allowable values for SessionKeyTypeEnum
const (
	SessionKeyTypePub SessionKeyTypeEnum = "PUB"
)

var mappingSessionKeyTypeEnum = map[string]SessionKeyTypeEnum{
	"PUB": SessionKeyTypePub,
}

var mappingSessionKeyTypeEnumLowerCase = map[string]SessionKeyTypeEnum{
	"pub": SessionKeyTypePub,
}

// GetSessionKeyTypeEnumValues Enumerates the set of values for SessionKeyTypeEnum
func GetSessionKeyTypeEnumValues() []SessionKeyTypeEnum {
	values := make([]SessionKeyTypeEnum, 0)
	for _, v := range mappingSessionKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionKeyTypeEnumStringValues Enumerates the set of values in String for SessionKeyTypeEnum
func GetSessionKeyTypeEnumStringValues() []string {
	return []string{
		"PUB",
	}
}

// GetMappingSessionKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionKeyTypeEnum(val string) (SessionKeyTypeEnum, bool) {
	enum, ok := mappingSessionKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
