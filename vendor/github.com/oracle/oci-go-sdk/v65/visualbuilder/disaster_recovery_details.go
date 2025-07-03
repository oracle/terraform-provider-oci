// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryDetails Disaster recovery details of the vb instance
type DisasterRecoveryDetails struct {

	// DR region type for the vb instance URL
	RegionType DisasterRecoveryDetailsRegionTypeEnum `mandatory:"false" json:"regionType,omitempty"`

	// DR state for the vb instance URL
	DisasterRecoveryState DisasterRecoveryDetailsDisasterRecoveryStateEnum `mandatory:"false" json:"disasterRecoveryState,omitempty"`

	// OCI region of the peer Vb Instance.
	PeerRegion *string `mandatory:"false" json:"peerRegion"`

	// Display name of the peer Vb Instance.
	PeerVbName *string `mandatory:"false" json:"peerVbName"`

	// Ocid of the peer Vb Instance.
	PeerVbId *string `mandatory:"false" json:"peerVbId"`

	// The time the VbInstance DR state was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DisasterRecoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisasterRecoveryDetailsRegionTypeEnum(string(m.RegionType)); !ok && m.RegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegionType: %s. Supported values are: %s.", m.RegionType, strings.Join(GetDisasterRecoveryDetailsRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryDetailsDisasterRecoveryStateEnum(string(m.DisasterRecoveryState)); !ok && m.DisasterRecoveryState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryState: %s. Supported values are: %s.", m.DisasterRecoveryState, strings.Join(GetDisasterRecoveryDetailsDisasterRecoveryStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisasterRecoveryDetailsRegionTypeEnum Enum with underlying type: string
type DisasterRecoveryDetailsRegionTypeEnum string

// Set of constants representing the allowable values for DisasterRecoveryDetailsRegionTypeEnum
const (
	DisasterRecoveryDetailsRegionTypePrimary DisasterRecoveryDetailsRegionTypeEnum = "PRIMARY"
	DisasterRecoveryDetailsRegionTypeRemote  DisasterRecoveryDetailsRegionTypeEnum = "REMOTE"
)

var mappingDisasterRecoveryDetailsRegionTypeEnum = map[string]DisasterRecoveryDetailsRegionTypeEnum{
	"PRIMARY": DisasterRecoveryDetailsRegionTypePrimary,
	"REMOTE":  DisasterRecoveryDetailsRegionTypeRemote,
}

var mappingDisasterRecoveryDetailsRegionTypeEnumLowerCase = map[string]DisasterRecoveryDetailsRegionTypeEnum{
	"primary": DisasterRecoveryDetailsRegionTypePrimary,
	"remote":  DisasterRecoveryDetailsRegionTypeRemote,
}

// GetDisasterRecoveryDetailsRegionTypeEnumValues Enumerates the set of values for DisasterRecoveryDetailsRegionTypeEnum
func GetDisasterRecoveryDetailsRegionTypeEnumValues() []DisasterRecoveryDetailsRegionTypeEnum {
	values := make([]DisasterRecoveryDetailsRegionTypeEnum, 0)
	for _, v := range mappingDisasterRecoveryDetailsRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryDetailsRegionTypeEnumStringValues Enumerates the set of values in String for DisasterRecoveryDetailsRegionTypeEnum
func GetDisasterRecoveryDetailsRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"REMOTE",
	}
}

// GetMappingDisasterRecoveryDetailsRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryDetailsRegionTypeEnum(val string) (DisasterRecoveryDetailsRegionTypeEnum, bool) {
	enum, ok := mappingDisasterRecoveryDetailsRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DisasterRecoveryDetailsDisasterRecoveryStateEnum Enum with underlying type: string
type DisasterRecoveryDetailsDisasterRecoveryStateEnum string

// Set of constants representing the allowable values for DisasterRecoveryDetailsDisasterRecoveryStateEnum
const (
	DisasterRecoveryDetailsDisasterRecoveryStateActive          DisasterRecoveryDetailsDisasterRecoveryStateEnum = "ACTIVE"
	DisasterRecoveryDetailsDisasterRecoveryStateStandby         DisasterRecoveryDetailsDisasterRecoveryStateEnum = "STANDBY"
	DisasterRecoveryDetailsDisasterRecoveryStateSnapshotStandby DisasterRecoveryDetailsDisasterRecoveryStateEnum = "SNAPSHOT_STANDBY"
)

var mappingDisasterRecoveryDetailsDisasterRecoveryStateEnum = map[string]DisasterRecoveryDetailsDisasterRecoveryStateEnum{
	"ACTIVE":           DisasterRecoveryDetailsDisasterRecoveryStateActive,
	"STANDBY":          DisasterRecoveryDetailsDisasterRecoveryStateStandby,
	"SNAPSHOT_STANDBY": DisasterRecoveryDetailsDisasterRecoveryStateSnapshotStandby,
}

var mappingDisasterRecoveryDetailsDisasterRecoveryStateEnumLowerCase = map[string]DisasterRecoveryDetailsDisasterRecoveryStateEnum{
	"active":           DisasterRecoveryDetailsDisasterRecoveryStateActive,
	"standby":          DisasterRecoveryDetailsDisasterRecoveryStateStandby,
	"snapshot_standby": DisasterRecoveryDetailsDisasterRecoveryStateSnapshotStandby,
}

// GetDisasterRecoveryDetailsDisasterRecoveryStateEnumValues Enumerates the set of values for DisasterRecoveryDetailsDisasterRecoveryStateEnum
func GetDisasterRecoveryDetailsDisasterRecoveryStateEnumValues() []DisasterRecoveryDetailsDisasterRecoveryStateEnum {
	values := make([]DisasterRecoveryDetailsDisasterRecoveryStateEnum, 0)
	for _, v := range mappingDisasterRecoveryDetailsDisasterRecoveryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryDetailsDisasterRecoveryStateEnumStringValues Enumerates the set of values in String for DisasterRecoveryDetailsDisasterRecoveryStateEnum
func GetDisasterRecoveryDetailsDisasterRecoveryStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"STANDBY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingDisasterRecoveryDetailsDisasterRecoveryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryDetailsDisasterRecoveryStateEnum(val string) (DisasterRecoveryDetailsDisasterRecoveryStateEnum, bool) {
	enum, ok := mappingDisasterRecoveryDetailsDisasterRecoveryStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
