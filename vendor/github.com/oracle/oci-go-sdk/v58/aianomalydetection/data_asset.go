// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DataAsset Description of DataAsset.
type DataAsset struct {

	// The Unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the DataAsset.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the DataAsset was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the Data Asset.
	LifecycleState DataAssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Unique project id which is created at project creation that is immutable on creation.
	ProjectId *string `mandatory:"true" json:"projectId"`

	DataSourceDetails DataSourceDetails `mandatory:"true" json:"dataSourceDetails"`

	// A short description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The time the the DataAsset was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of Private Endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

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

func (m DataAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataAssetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataAsset) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                           `json:"description"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		PrivateEndpointId *string                           `json:"privateEndpointId"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		CompartmentId     *string                           `json:"compartmentId"`
		DisplayName       *string                           `json:"displayName"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		LifecycleState    DataAssetLifecycleStateEnum       `json:"lifecycleState"`
		ProjectId         *string                           `json:"projectId"`
		DataSourceDetails datasourcedetails                 `json:"dataSourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.PrivateEndpointId = model.PrivateEndpointId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.ProjectId = model.ProjectId

	nn, e = model.DataSourceDetails.UnmarshalPolymorphicJSON(model.DataSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSourceDetails = nn.(DataSourceDetails)
	} else {
		m.DataSourceDetails = nil
	}

	return
}

// DataAssetLifecycleStateEnum Enum with underlying type: string
type DataAssetLifecycleStateEnum string

// Set of constants representing the allowable values for DataAssetLifecycleStateEnum
const (
	DataAssetLifecycleStateActive  DataAssetLifecycleStateEnum = "ACTIVE"
	DataAssetLifecycleStateDeleted DataAssetLifecycleStateEnum = "DELETED"
)

var mappingDataAssetLifecycleStateEnum = map[string]DataAssetLifecycleStateEnum{
	"ACTIVE":  DataAssetLifecycleStateActive,
	"DELETED": DataAssetLifecycleStateDeleted,
}

// GetDataAssetLifecycleStateEnumValues Enumerates the set of values for DataAssetLifecycleStateEnum
func GetDataAssetLifecycleStateEnumValues() []DataAssetLifecycleStateEnum {
	values := make([]DataAssetLifecycleStateEnum, 0)
	for _, v := range mappingDataAssetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataAssetLifecycleStateEnumStringValues Enumerates the set of values in String for DataAssetLifecycleStateEnum
func GetDataAssetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingDataAssetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataAssetLifecycleStateEnum(val string) (DataAssetLifecycleStateEnum, bool) {
	mappingDataAssetLifecycleStateEnumIgnoreCase := make(map[string]DataAssetLifecycleStateEnum)
	for k, v := range mappingDataAssetLifecycleStateEnum {
		mappingDataAssetLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataAssetLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
