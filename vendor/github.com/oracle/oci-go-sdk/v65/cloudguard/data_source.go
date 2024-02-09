// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSource Detailed information for a data source (DataSource resource).
type DataSource struct {

	// OCID for the data source
	Id *string `mandatory:"true" json:"id"`

	// Display name of the data source
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Possible type of dataSourceFeed Provider(LoggingQuery)
	DataSourceFeedProvider DataSourceFeedProviderEnum `mandatory:"true" json:"dataSourceFeedProvider"`

	// Compartment OCID of data source
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DataSourceDetails DataSourceDetails `mandatory:"false" json:"dataSourceDetails"`

	// The date and time the Data source was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the data source was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Enablement status of the data source
	Status DataSourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Information about the detector recipe and rule attached
	DataSourceDetectorMappingInfo []DataSourceMappingInfo `mandatory:"false" json:"dataSourceDetectorMappingInfo"`

	// Information about the region and status of query replication
	RegionStatusDetail []RegionStatusDetail `mandatory:"false" json:"regionStatusDetail"`

	// The current lifecycle state of the resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DataSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSourceFeedProviderEnum(string(m.DataSourceFeedProvider)); !ok && m.DataSourceFeedProvider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSourceFeedProvider: %s. Supported values are: %s.", m.DataSourceFeedProvider, strings.Join(GetDataSourceFeedProviderEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataSourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDataSourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataSourceDetails             datasourcedetails                 `json:"dataSourceDetails"`
		TimeCreated                   *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                   *common.SDKTime                   `json:"timeUpdated"`
		Status                        DataSourceStatusEnum              `json:"status"`
		DataSourceDetectorMappingInfo []DataSourceMappingInfo           `json:"dataSourceDetectorMappingInfo"`
		RegionStatusDetail            []RegionStatusDetail              `json:"regionStatusDetail"`
		LifecycleState                LifecycleStateEnum                `json:"lifecycleState"`
		Locks                         []ResourceLock                    `json:"locks"`
		FreeformTags                  map[string]string                 `json:"freeformTags"`
		DefinedTags                   map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                    map[string]map[string]interface{} `json:"systemTags"`
		Id                            *string                           `json:"id"`
		DisplayName                   *string                           `json:"displayName"`
		DataSourceFeedProvider        DataSourceFeedProviderEnum        `json:"dataSourceFeedProvider"`
		CompartmentId                 *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DataSourceDetails.UnmarshalPolymorphicJSON(model.DataSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSourceDetails = nn.(DataSourceDetails)
	} else {
		m.DataSourceDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Status = model.Status

	m.DataSourceDetectorMappingInfo = make([]DataSourceMappingInfo, len(model.DataSourceDetectorMappingInfo))
	copy(m.DataSourceDetectorMappingInfo, model.DataSourceDetectorMappingInfo)
	m.RegionStatusDetail = make([]RegionStatusDetail, len(model.RegionStatusDetail))
	copy(m.RegionStatusDetail, model.RegionStatusDetail)
	m.LifecycleState = model.LifecycleState

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.DataSourceFeedProvider = model.DataSourceFeedProvider

	m.CompartmentId = model.CompartmentId

	return
}
