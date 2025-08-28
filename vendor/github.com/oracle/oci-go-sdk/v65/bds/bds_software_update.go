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

// BdsSoftwareUpdate Details about the given BDS type of software update. Previously known as Micro Service Patch.
type BdsSoftwareUpdate struct {

	// Unique identifier of a given software update
	SoftwareUpdateKey *string `mandatory:"true" json:"softwareUpdateKey"`

	// The version of the software update.
	SoftwareUpdateVersion *string `mandatory:"true" json:"softwareUpdateVersion"`

	// The time when the software update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The due date for the software update. Big Data Service will be updated automatically after this date.
	TimeDue *common.SDKTime `mandatory:"true" json:"timeDue"`

	// The lifecycle state of the software update.
	LifecycleState SoftwareUpdateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetSoftwareUpdateKey returns SoftwareUpdateKey
func (m BdsSoftwareUpdate) GetSoftwareUpdateKey() *string {
	return m.SoftwareUpdateKey
}

// GetSoftwareUpdateVersion returns SoftwareUpdateVersion
func (m BdsSoftwareUpdate) GetSoftwareUpdateVersion() *string {
	return m.SoftwareUpdateVersion
}

// GetTimeReleased returns TimeReleased
func (m BdsSoftwareUpdate) GetTimeReleased() *common.SDKTime {
	return m.TimeReleased
}

// GetLifecycleState returns LifecycleState
func (m BdsSoftwareUpdate) GetLifecycleState() SoftwareUpdateLifecycleStateEnum {
	return m.LifecycleState
}

func (m BdsSoftwareUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsSoftwareUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwareUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareUpdateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BdsSoftwareUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBdsSoftwareUpdate BdsSoftwareUpdate
	s := struct {
		DiscriminatorParam string `json:"softwareUpdateType"`
		MarshalTypeBdsSoftwareUpdate
	}{
		"BDS",
		(MarshalTypeBdsSoftwareUpdate)(m),
	}

	return json.Marshal(&s)
}
