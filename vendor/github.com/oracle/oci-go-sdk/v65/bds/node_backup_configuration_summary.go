// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// NodeBackupConfigurationSummary The information about the NodeBackupConfiguration.
type NodeBackupConfigurationSummary struct {

	// The id of the NodeBackupConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the bdsInstance which is the parent resource id.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	LevelTypeDetails LevelTypeDetails `mandatory:"true" json:"levelTypeDetails"`

	// The state of the NodeBackupConfiguration.
	LifecycleState NodeBackupConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the NodeBackupConfiguration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the NodeBackupConfiguration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m NodeBackupConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeBackupConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeBackupConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeBackupConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *NodeBackupConfigurationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id               *string                                   `json:"id"`
		BdsInstanceId    *string                                   `json:"bdsInstanceId"`
		DisplayName      *string                                   `json:"displayName"`
		LevelTypeDetails leveltypedetails                          `json:"levelTypeDetails"`
		LifecycleState   NodeBackupConfigurationLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                           `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                           `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.BdsInstanceId = model.BdsInstanceId

	m.DisplayName = model.DisplayName

	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
