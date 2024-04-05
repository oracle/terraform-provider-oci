// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlEndpoint The description of a SQL Endpoint.
type SqlEndpoint struct {

	// The provision identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The SQL Endpoint name, which can be changed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The version of SQL Endpoint.
	SqlEndpointVersion *string `mandatory:"true" json:"sqlEndpointVersion"`

	// The shape of the SQL Endpoint driver instance.
	DriverShape *string `mandatory:"true" json:"driverShape"`

	// The shape of the SQL Endpoint executor instance.
	ExecutorShape *string `mandatory:"true" json:"executorShape"`

	// The minimum number of executors.
	MinExecutorCount *int `mandatory:"true" json:"minExecutorCount"`

	// The maximum number of executors.
	MaxExecutorCount *int `mandatory:"true" json:"maxExecutorCount"`

	// The OCID of OCI Hive Metastore.
	MetastoreId *string `mandatory:"true" json:"metastoreId"`

	// The OCID of OCI Lake.
	LakeId *string `mandatory:"true" json:"lakeId"`

	// The warehouse bucket URI. It is a Oracle Cloud Infrastructure Object Storage bucket URI as defined here https://docs.oracle.com/en/cloud/paas/atp-cloud/atpud/object-storage-uris.html
	WarehouseBucketUri *string `mandatory:"true" json:"warehouseBucketUri"`

	// The description of the SQL Endpoint.
	Description *string `mandatory:"true" json:"description"`

	// The JDBC URL field. For example, jdbc:spark://{serviceFQDN}:443/default;SparkServerType=DFI
	JdbcEndpointUrl *string `mandatory:"false" json:"jdbcEndpointUrl"`

	// The time the Sql Endpoint was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Sql Endpoint was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Sql Endpoint.
	LifecycleState SqlEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the reason why the resource is in it's current state. Helps bubble up errors in state changes. For example, it can be used to provide actionable information for a resource in the Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// This token is used by Splat, and indicates that the service accepts the request, and that the request is currently being processed.
	LastAcceptedRequestToken *string `mandatory:"false" json:"lastAcceptedRequestToken"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	SparkAdvancedConfigurations map[string]string `mandatory:"false" json:"sparkAdvancedConfigurations"`

	// The SQL Endpoint message displayed as a banner to provide user with any action items required on the resource.
	BannerMessage *string `mandatory:"false" json:"bannerMessage"`

	NetworkConfiguration SqlEndpointNetworkConfiguration `mandatory:"false" json:"networkConfiguration"`
}

func (m SqlEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SqlEndpoint) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		JdbcEndpointUrl             *string                           `json:"jdbcEndpointUrl"`
		TimeCreated                 *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                 *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState              SqlEndpointLifecycleStateEnum     `json:"lifecycleState"`
		StateMessage                *string                           `json:"stateMessage"`
		DriverShapeConfig           *ShapeConfig                      `json:"driverShapeConfig"`
		ExecutorShapeConfig         *ShapeConfig                      `json:"executorShapeConfig"`
		LastAcceptedRequestToken    *string                           `json:"lastAcceptedRequestToken"`
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                  map[string]map[string]interface{} `json:"systemTags"`
		SparkAdvancedConfigurations map[string]string                 `json:"sparkAdvancedConfigurations"`
		BannerMessage               *string                           `json:"bannerMessage"`
		NetworkConfiguration        sqlendpointnetworkconfiguration   `json:"networkConfiguration"`
		Id                          *string                           `json:"id"`
		DisplayName                 *string                           `json:"displayName"`
		CompartmentId               *string                           `json:"compartmentId"`
		SqlEndpointVersion          *string                           `json:"sqlEndpointVersion"`
		DriverShape                 *string                           `json:"driverShape"`
		ExecutorShape               *string                           `json:"executorShape"`
		MinExecutorCount            *int                              `json:"minExecutorCount"`
		MaxExecutorCount            *int                              `json:"maxExecutorCount"`
		MetastoreId                 *string                           `json:"metastoreId"`
		LakeId                      *string                           `json:"lakeId"`
		WarehouseBucketUri          *string                           `json:"warehouseBucketUri"`
		Description                 *string                           `json:"description"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.JdbcEndpointUrl = model.JdbcEndpointUrl

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.StateMessage = model.StateMessage

	m.DriverShapeConfig = model.DriverShapeConfig

	m.ExecutorShapeConfig = model.ExecutorShapeConfig

	m.LastAcceptedRequestToken = model.LastAcceptedRequestToken

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.SparkAdvancedConfigurations = model.SparkAdvancedConfigurations

	m.BannerMessage = model.BannerMessage

	nn, e = model.NetworkConfiguration.UnmarshalPolymorphicJSON(model.NetworkConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkConfiguration = nn.(SqlEndpointNetworkConfiguration)
	} else {
		m.NetworkConfiguration = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.SqlEndpointVersion = model.SqlEndpointVersion

	m.DriverShape = model.DriverShape

	m.ExecutorShape = model.ExecutorShape

	m.MinExecutorCount = model.MinExecutorCount

	m.MaxExecutorCount = model.MaxExecutorCount

	m.MetastoreId = model.MetastoreId

	m.LakeId = model.LakeId

	m.WarehouseBucketUri = model.WarehouseBucketUri

	m.Description = model.Description

	return
}
