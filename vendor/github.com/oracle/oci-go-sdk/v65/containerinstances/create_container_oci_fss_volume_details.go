// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateContainerOciFssVolumeDetails An OCI File Storage Service (FSS) File System that can be mounted to Containers within a Container Instance.
// Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/FileSystem for more details.
type CreateContainerOciFssVolumeDetails struct {

	// The name of the volume. This must be unique within a single container instance.
	Name *string `mandatory:"true" json:"name"`

	MountTarget CreateOciFssMountTargetDetails `mandatory:"true" json:"mountTarget"`

	Export CreateOciFssExportDetails `mandatory:"true" json:"export"`

	Security CreateOciFssSecurityDetails `mandatory:"false" json:"security"`

	MountCommand *CreateOciFssMountCommandDetails `mandatory:"false" json:"mountCommand"`

	// Specifies the network interface to be used for the OCI File Storage Service (FSS) volume.
	// This is a required parameter when a Container Instance is attached to more than one subnets.
	SubnetId *string `mandatory:"false" json:"subnetId"`
}

// GetName returns Name
func (m CreateContainerOciFssVolumeDetails) GetName() *string {
	return m.Name
}

func (m CreateContainerOciFssVolumeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerOciFssVolumeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateContainerOciFssVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateContainerOciFssVolumeDetails CreateContainerOciFssVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeCreateContainerOciFssVolumeDetails
	}{
		"OCI_FSS_FILE_SYSTEM",
		(MarshalTypeCreateContainerOciFssVolumeDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateContainerOciFssVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Security     createocifsssecuritydetails      `json:"security"`
		MountCommand *CreateOciFssMountCommandDetails `json:"mountCommand"`
		SubnetId     *string                          `json:"subnetId"`
		Name         *string                          `json:"name"`
		MountTarget  createocifssmounttargetdetails   `json:"mountTarget"`
		Export       createocifssexportdetails        `json:"export"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Security.UnmarshalPolymorphicJSON(model.Security.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Security = nn.(CreateOciFssSecurityDetails)
	} else {
		m.Security = nil
	}

	m.MountCommand = model.MountCommand

	m.SubnetId = model.SubnetId

	m.Name = model.Name

	nn, e = model.MountTarget.UnmarshalPolymorphicJSON(model.MountTarget.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.MountTarget = nn.(CreateOciFssMountTargetDetails)
	} else {
		m.MountTarget = nil
	}

	nn, e = model.Export.UnmarshalPolymorphicJSON(model.Export.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Export = nn.(CreateOciFssExportDetails)
	} else {
		m.Export = nil
	}

	return
}
