// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSqlEndpointDetails The information about all updatable parameters of a SQL Endpoint.
type UpdateSqlEndpointDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The SQL Endpoint name, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of CreateSQLEndpointDetails.
	Description *string `mandatory:"false" json:"description"`

	// The shape of the SQL Endpoint driver instance.
	DriverShape *string `mandatory:"false" json:"driverShape"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	// The shape of the SQL Endpoint worker instance.
	ExecutorShape *string `mandatory:"false" json:"executorShape"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// The minimum number of executors.
	MinExecutorCount *int `mandatory:"false" json:"minExecutorCount"`

	// The maximum number of executors.
	MaxExecutorCount *int `mandatory:"false" json:"maxExecutorCount"`

	// Metastore OCID
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// OCI lake OCID
	LakeId *string `mandatory:"false" json:"lakeId"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	SparkAdvancedConfigurations map[string]string `mandatory:"false" json:"sparkAdvancedConfigurations"`
}

func (m UpdateSqlEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSqlEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
