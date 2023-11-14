// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbSystemDetails The information about new DbSystem.
type CreateDbSystemDetails struct {

	// DbSystem display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Version of DbSystem software.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	StorageDetails StorageDetails `mandatory:"true" json:"storageDetails"`

	// Shape of DbInstance. This name should match from with one of the available shapes from /shapes API.
	Shape *string `mandatory:"true" json:"shape"`

	NetworkDetails *NetworkDetails `mandatory:"true" json:"networkDetails"`

	// Description of a DbSystem. This field should be input by the user.
	Description *string `mandatory:"false" json:"description"`

	// Type of the DbSystem.
	SystemType DbSystemSystemTypeEnum `mandatory:"false" json:"systemType,omitempty"`

	// Configuration identifier
	ConfigId *string `mandatory:"false" json:"configId"`

	// The total number of OCPUs available to each DbInstance.
	InstanceOcpuCount *int `mandatory:"false" json:"instanceOcpuCount"`

	// The total amount of memory available to each DbInstance, in gigabytes.
	InstanceMemorySizeInGBs *int `mandatory:"false" json:"instanceMemorySizeInGBs"`

	// Count of DbInstances to be created in the DbSystem.
	InstanceCount *int `mandatory:"false" json:"instanceCount"`

	// Details of DbInstances to be created. Optional parameter.
	// If specified, its size must match instanceCount.
	InstancesDetails []CreateDbInstanceDetails `mandatory:"false" json:"instancesDetails"`

	Credentials *Credentials `mandatory:"false" json:"credentials"`

	ManagementPolicy *ManagementPolicyDetails `mandatory:"false" json:"managementPolicy"`

	Source SourceDetails `mandatory:"false" json:"source"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbSystemSystemTypeEnum(string(m.SystemType)); !ok && m.SystemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SystemType: %s. Supported values are: %s.", m.SystemType, strings.Join(GetDbSystemSystemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDbSystemDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                           `json:"description"`
		SystemType              DbSystemSystemTypeEnum            `json:"systemType"`
		ConfigId                *string                           `json:"configId"`
		InstanceOcpuCount       *int                              `json:"instanceOcpuCount"`
		InstanceMemorySizeInGBs *int                              `json:"instanceMemorySizeInGBs"`
		InstanceCount           *int                              `json:"instanceCount"`
		InstancesDetails        []CreateDbInstanceDetails         `json:"instancesDetails"`
		Credentials             *Credentials                      `json:"credentials"`
		ManagementPolicy        *ManagementPolicyDetails          `json:"managementPolicy"`
		Source                  sourcedetails                     `json:"source"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		DisplayName             *string                           `json:"displayName"`
		CompartmentId           *string                           `json:"compartmentId"`
		DbVersion               *string                           `json:"dbVersion"`
		StorageDetails          storagedetails                    `json:"storageDetails"`
		Shape                   *string                           `json:"shape"`
		NetworkDetails          *NetworkDetails                   `json:"networkDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.SystemType = model.SystemType

	m.ConfigId = model.ConfigId

	m.InstanceOcpuCount = model.InstanceOcpuCount

	m.InstanceMemorySizeInGBs = model.InstanceMemorySizeInGBs

	m.InstanceCount = model.InstanceCount

	m.InstancesDetails = make([]CreateDbInstanceDetails, len(model.InstancesDetails))
	copy(m.InstancesDetails, model.InstancesDetails)
	m.Credentials = model.Credentials

	m.ManagementPolicy = model.ManagementPolicy

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SourceDetails)
	} else {
		m.Source = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.DbVersion = model.DbVersion

	nn, e = model.StorageDetails.UnmarshalPolymorphicJSON(model.StorageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StorageDetails = nn.(StorageDetails)
	} else {
		m.StorageDetails = nil
	}

	m.Shape = model.Shape

	m.NetworkDetails = model.NetworkDetails

	return
}
