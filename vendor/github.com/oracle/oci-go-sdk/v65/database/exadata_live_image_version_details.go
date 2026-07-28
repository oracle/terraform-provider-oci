// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataLiveImageVersionDetails Details about the most recent live image version applied on the VM Cluster, if any. If a full OS update was applied, the fields would be blank.
type ExadataLiveImageVersionDetails struct {

	// The OS live update mode performed most recently on the VM Cluster.
	UpdateMode ExadataLiveImageVersionDetailsUpdateModeEnum `mandatory:"false" json:"updateMode,omitempty"`

	// Live Exadata Image Version of the Guest OS Update applied.
	Version *string `mandatory:"false" json:"version"`

	// Indicates whether OS updates that require node reboot are pending after the previous online update was applied.
	HasPendingUpdates *bool `mandatory:"false" json:"hasPendingUpdates"`

	// The release date and time for the applied Live Exadata Image OS version.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`
}

func (m ExadataLiveImageVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataLiveImageVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataLiveImageVersionDetailsUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetExadataLiveImageVersionDetailsUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataLiveImageVersionDetailsUpdateModeEnum Enum with underlying type: string
type ExadataLiveImageVersionDetailsUpdateModeEnum string

// Set of constants representing the allowable values for ExadataLiveImageVersionDetailsUpdateModeEnum
const (
	ExadataLiveImageVersionDetailsUpdateModeHighcvss   ExadataLiveImageVersionDetailsUpdateModeEnum = "ONLINE_HIGHCVSS"
	ExadataLiveImageVersionDetailsUpdateModeAllcvss    ExadataLiveImageVersionDetailsUpdateModeEnum = "ONLINE_ALLCVSS"
	ExadataLiveImageVersionDetailsUpdateModeAllUpdates ExadataLiveImageVersionDetailsUpdateModeEnum = "ONLINE_ALL_UPDATES"
)

var mappingExadataLiveImageVersionDetailsUpdateModeEnum = map[string]ExadataLiveImageVersionDetailsUpdateModeEnum{
	"ONLINE_HIGHCVSS":    ExadataLiveImageVersionDetailsUpdateModeHighcvss,
	"ONLINE_ALLCVSS":     ExadataLiveImageVersionDetailsUpdateModeAllcvss,
	"ONLINE_ALL_UPDATES": ExadataLiveImageVersionDetailsUpdateModeAllUpdates,
}

var mappingExadataLiveImageVersionDetailsUpdateModeEnumLowerCase = map[string]ExadataLiveImageVersionDetailsUpdateModeEnum{
	"online_highcvss":    ExadataLiveImageVersionDetailsUpdateModeHighcvss,
	"online_allcvss":     ExadataLiveImageVersionDetailsUpdateModeAllcvss,
	"online_all_updates": ExadataLiveImageVersionDetailsUpdateModeAllUpdates,
}

// GetExadataLiveImageVersionDetailsUpdateModeEnumValues Enumerates the set of values for ExadataLiveImageVersionDetailsUpdateModeEnum
func GetExadataLiveImageVersionDetailsUpdateModeEnumValues() []ExadataLiveImageVersionDetailsUpdateModeEnum {
	values := make([]ExadataLiveImageVersionDetailsUpdateModeEnum, 0)
	for _, v := range mappingExadataLiveImageVersionDetailsUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataLiveImageVersionDetailsUpdateModeEnumStringValues Enumerates the set of values in String for ExadataLiveImageVersionDetailsUpdateModeEnum
func GetExadataLiveImageVersionDetailsUpdateModeEnumStringValues() []string {
	return []string{
		"ONLINE_HIGHCVSS",
		"ONLINE_ALLCVSS",
		"ONLINE_ALL_UPDATES",
	}
}

// GetMappingExadataLiveImageVersionDetailsUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataLiveImageVersionDetailsUpdateModeEnum(val string) (ExadataLiveImageVersionDetailsUpdateModeEnum, bool) {
	enum, ok := mappingExadataLiveImageVersionDetailsUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
