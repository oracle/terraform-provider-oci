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

// CreateSqlEndpointDetails The information about a new SQL Endpoint.
type CreateSqlEndpointDetails struct {

	// The identifier of the compartment used with the SQL Endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The SQL Endpoint name, which can be changed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The version of the SQL Endpoint.
	SqlEndpointVersion *string `mandatory:"true" json:"sqlEndpointVersion"`

	// The shape of the SQL Endpoint driver instance.
	DriverShape *string `mandatory:"true" json:"driverShape"`

	// The shape of the SQL Endpoint worker instance.
	ExecutorShape *string `mandatory:"true" json:"executorShape"`

	// The minimum number of executors.
	MinExecutorCount *int `mandatory:"true" json:"minExecutorCount"`

	// The maximum number of executors.
	MaxExecutorCount *int `mandatory:"true" json:"maxExecutorCount"`

	// Metastore OCID
	MetastoreId *string `mandatory:"true" json:"metastoreId"`

	// OCI lake OCID
	LakeId *string `mandatory:"true" json:"lakeId"`

	// The warehouse bucket URI. It is a Oracle Cloud Infrastructure Object Storage bucket URI as defined here https://docs.oracle.com/en/cloud/paas/atp-cloud/atpud/object-storage-uris.html
	WarehouseBucketUri *string `mandatory:"true" json:"warehouseBucketUri"`

	NetworkConfiguration SqlEndpointNetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// The description of CreateSQLEndpointDetails.
	Description *string `mandatory:"false" json:"description"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	SparkAdvancedConfigurations map[string]string `mandatory:"false" json:"sparkAdvancedConfigurations"`
}

func (m CreateSqlEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSqlEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSqlEndpointDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                 *string                           `json:"description"`
		DriverShapeConfig           *ShapeConfig                      `json:"driverShapeConfig"`
		ExecutorShapeConfig         *ShapeConfig                      `json:"executorShapeConfig"`
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
		SparkAdvancedConfigurations map[string]string                 `json:"sparkAdvancedConfigurations"`
		CompartmentId               *string                           `json:"compartmentId"`
		DisplayName                 *string                           `json:"displayName"`
		SqlEndpointVersion          *string                           `json:"sqlEndpointVersion"`
		DriverShape                 *string                           `json:"driverShape"`
		ExecutorShape               *string                           `json:"executorShape"`
		MinExecutorCount            *int                              `json:"minExecutorCount"`
		MaxExecutorCount            *int                              `json:"maxExecutorCount"`
		MetastoreId                 *string                           `json:"metastoreId"`
		LakeId                      *string                           `json:"lakeId"`
		WarehouseBucketUri          *string                           `json:"warehouseBucketUri"`
		NetworkConfiguration        sqlendpointnetworkconfiguration   `json:"networkConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DriverShapeConfig = model.DriverShapeConfig

	m.ExecutorShapeConfig = model.ExecutorShapeConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SparkAdvancedConfigurations = model.SparkAdvancedConfigurations

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.SqlEndpointVersion = model.SqlEndpointVersion

	m.DriverShape = model.DriverShape

	m.ExecutorShape = model.ExecutorShape

	m.MinExecutorCount = model.MinExecutorCount

	m.MaxExecutorCount = model.MaxExecutorCount

	m.MetastoreId = model.MetastoreId

	m.LakeId = model.LakeId

	m.WarehouseBucketUri = model.WarehouseBucketUri

	nn, e = model.NetworkConfiguration.UnmarshalPolymorphicJSON(model.NetworkConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkConfiguration = nn.(SqlEndpointNetworkConfiguration)
	} else {
		m.NetworkConfiguration = nil
	}

	return
}
