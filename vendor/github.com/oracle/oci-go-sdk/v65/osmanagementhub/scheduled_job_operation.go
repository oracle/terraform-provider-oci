// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJobOperation Defines an operation that is performed by a scheduled job.
type ScheduledJobOperation struct {

	// The type of operation this scheduled job performs.
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// The types of update operations to stage updates for.
	StageUpdateTypes []StageUpdateTypesEnum `mandatory:"false" json:"stageUpdateTypes,omitempty"`

	// The names of the target packages. This parameter only applies when the scheduled job is for installing, updating, or removing packages.
	PackageNames []string `mandatory:"false" json:"packageNames"`

	// Unique identifier for the Windows update. This parameter only applies if the scheduled job is for installing Windows updates.
	// Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// For example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'.
	WindowsUpdateNames []string `mandatory:"false" json:"windowsUpdateNames"`

	ManageModuleStreamsDetails *ManageModuleStreamsInScheduledJobDetails `mandatory:"false" json:"manageModuleStreamsDetails"`

	SwitchModuleStreamsDetails *ModuleStreamDetails `mandatory:"false" json:"switchModuleStreamsDetails"`

	// The software source OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This parameter only applies when the scheduled job is for attaching/detaching or updating software sources.
	SoftwareSourceIds []string `mandatory:"false" json:"softwareSourceIds"`

	// The number of minutes the service waits for the reboot to complete. If the instance doesn't reboot within the
	// timeout, the service marks the reboot job as failed.
	RebootTimeoutInMins *int `mandatory:"false" json:"rebootTimeoutInMins"`

	VulnerabilityDetails VulnerabilityDetails `mandatory:"false" json:"vulnerabilityDetails"`

	InstallSnapDetails *InstallSnapDetails `mandatory:"false" json:"installSnapDetails"`

	RemoveSnapDetails *RemoveSnapDetails `mandatory:"false" json:"removeSnapDetails"`

	SwitchSnapChannelDetails *SwitchSnapChannelDetails `mandatory:"false" json:"switchSnapChannelDetails"`
}

func (m ScheduledJobOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledJobOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationTypesEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetOperationTypesEnumStringValues(), ",")))
	}

	for _, val := range m.StageUpdateTypes {
		if _, ok := GetMappingStageUpdateTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StageUpdateTypes: %s. Supported values are: %s.", val, strings.Join(GetStageUpdateTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduledJobOperation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StageUpdateTypes           []StageUpdateTypesEnum                    `json:"stageUpdateTypes"`
		PackageNames               []string                                  `json:"packageNames"`
		WindowsUpdateNames         []string                                  `json:"windowsUpdateNames"`
		ManageModuleStreamsDetails *ManageModuleStreamsInScheduledJobDetails `json:"manageModuleStreamsDetails"`
		SwitchModuleStreamsDetails *ModuleStreamDetails                      `json:"switchModuleStreamsDetails"`
		SoftwareSourceIds          []string                                  `json:"softwareSourceIds"`
		RebootTimeoutInMins        *int                                      `json:"rebootTimeoutInMins"`
		VulnerabilityDetails       vulnerabilitydetails                      `json:"vulnerabilityDetails"`
		InstallSnapDetails         *InstallSnapDetails                       `json:"installSnapDetails"`
		RemoveSnapDetails          *RemoveSnapDetails                        `json:"removeSnapDetails"`
		SwitchSnapChannelDetails   *SwitchSnapChannelDetails                 `json:"switchSnapChannelDetails"`
		OperationType              OperationTypesEnum                        `json:"operationType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StageUpdateTypes = make([]StageUpdateTypesEnum, len(model.StageUpdateTypes))
	copy(m.StageUpdateTypes, model.StageUpdateTypes)
	m.PackageNames = make([]string, len(model.PackageNames))
	copy(m.PackageNames, model.PackageNames)
	m.WindowsUpdateNames = make([]string, len(model.WindowsUpdateNames))
	copy(m.WindowsUpdateNames, model.WindowsUpdateNames)
	m.ManageModuleStreamsDetails = model.ManageModuleStreamsDetails

	m.SwitchModuleStreamsDetails = model.SwitchModuleStreamsDetails

	m.SoftwareSourceIds = make([]string, len(model.SoftwareSourceIds))
	copy(m.SoftwareSourceIds, model.SoftwareSourceIds)
	m.RebootTimeoutInMins = model.RebootTimeoutInMins

	nn, e = model.VulnerabilityDetails.UnmarshalPolymorphicJSON(model.VulnerabilityDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.VulnerabilityDetails = nn.(VulnerabilityDetails)
	} else {
		m.VulnerabilityDetails = nil
	}

	m.InstallSnapDetails = model.InstallSnapDetails

	m.RemoveSnapDetails = model.RemoveSnapDetails

	m.SwitchSnapChannelDetails = model.SwitchSnapChannelDetails

	m.OperationType = model.OperationType

	return
}
