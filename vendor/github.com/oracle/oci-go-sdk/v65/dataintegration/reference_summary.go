// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReferenceSummary This is the reference summary information.
type ReferenceSummary struct {

	// The reference's key, key of the object that is being used by a published object or its dependents.
	Key *string `mandatory:"false" json:"key"`

	// The name of reference object.
	Name *string `mandatory:"false" json:"name"`

	// The identifier of reference object.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The identifier path of reference object.
	IdentifierPath *string `mandatory:"false" json:"identifierPath"`

	// The description of reference object.
	Description *string `mandatory:"false" json:"description"`

	// The type of reference object.
	Type ReferenceSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The target object referenced. References are made to data assets and child references are made to connections. The type defining this reference is mentioned in the property type.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The aggregator of reference object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// List of published objects where this is used.
	UsedBy []ReferenceUsedBy `mandatory:"false" json:"usedBy"`

	// List of references that are dependent on this reference.
	ChildReferences []ChildReference `mandatory:"false" json:"childReferences"`
}

func (m ReferenceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferenceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReferenceSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetReferenceSummaryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferenceSummaryTypeEnum Enum with underlying type: string
type ReferenceSummaryTypeEnum string

// Set of constants representing the allowable values for ReferenceSummaryTypeEnum
const (
	ReferenceSummaryTypeOracleDataAsset              ReferenceSummaryTypeEnum = "ORACLE_DATA_ASSET"
	ReferenceSummaryTypeOracleObjectStorageDataAsset ReferenceSummaryTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	ReferenceSummaryTypeOracleAtpDataAsset           ReferenceSummaryTypeEnum = "ORACLE_ATP_DATA_ASSET"
	ReferenceSummaryTypeOracleAdwcDataAsset          ReferenceSummaryTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	ReferenceSummaryTypeMysqlDataAsset               ReferenceSummaryTypeEnum = "MYSQL_DATA_ASSET"
	ReferenceSummaryTypeGenericJdbcDataAsset         ReferenceSummaryTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	ReferenceSummaryTypeFusionAppDataAsset           ReferenceSummaryTypeEnum = "FUSION_APP_DATA_ASSET"
	ReferenceSummaryTypeAmazonS3DataAsset            ReferenceSummaryTypeEnum = "AMAZON_S3_DATA_ASSET"
	ReferenceSummaryTypeSchema                       ReferenceSummaryTypeEnum = "SCHEMA"
	ReferenceSummaryTypeIntegrationTask              ReferenceSummaryTypeEnum = "INTEGRATION_TASK"
	ReferenceSummaryTypeDataLoaderTask               ReferenceSummaryTypeEnum = "DATA_LOADER_TASK"
	ReferenceSummaryTypeSqlTask                      ReferenceSummaryTypeEnum = "SQL_TASK"
	ReferenceSummaryTypeOciDataflowTask              ReferenceSummaryTypeEnum = "OCI_DATAFLOW_TASK"
	ReferenceSummaryTypePipelineTask                 ReferenceSummaryTypeEnum = "PIPELINE_TASK"
	ReferenceSummaryTypeRestTask                     ReferenceSummaryTypeEnum = "REST_TASK"
)

var mappingReferenceSummaryTypeEnum = map[string]ReferenceSummaryTypeEnum{
	"ORACLE_DATA_ASSET":                ReferenceSummaryTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": ReferenceSummaryTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            ReferenceSummaryTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           ReferenceSummaryTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 ReferenceSummaryTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          ReferenceSummaryTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            ReferenceSummaryTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             ReferenceSummaryTypeAmazonS3DataAsset,
	"SCHEMA":                           ReferenceSummaryTypeSchema,
	"INTEGRATION_TASK":                 ReferenceSummaryTypeIntegrationTask,
	"DATA_LOADER_TASK":                 ReferenceSummaryTypeDataLoaderTask,
	"SQL_TASK":                         ReferenceSummaryTypeSqlTask,
	"OCI_DATAFLOW_TASK":                ReferenceSummaryTypeOciDataflowTask,
	"PIPELINE_TASK":                    ReferenceSummaryTypePipelineTask,
	"REST_TASK":                        ReferenceSummaryTypeRestTask,
}

var mappingReferenceSummaryTypeEnumLowerCase = map[string]ReferenceSummaryTypeEnum{
	"oracle_data_asset":                ReferenceSummaryTypeOracleDataAsset,
	"oracle_object_storage_data_asset": ReferenceSummaryTypeOracleObjectStorageDataAsset,
	"oracle_atp_data_asset":            ReferenceSummaryTypeOracleAtpDataAsset,
	"oracle_adwc_data_asset":           ReferenceSummaryTypeOracleAdwcDataAsset,
	"mysql_data_asset":                 ReferenceSummaryTypeMysqlDataAsset,
	"generic_jdbc_data_asset":          ReferenceSummaryTypeGenericJdbcDataAsset,
	"fusion_app_data_asset":            ReferenceSummaryTypeFusionAppDataAsset,
	"amazon_s3_data_asset":             ReferenceSummaryTypeAmazonS3DataAsset,
	"schema":                           ReferenceSummaryTypeSchema,
	"integration_task":                 ReferenceSummaryTypeIntegrationTask,
	"data_loader_task":                 ReferenceSummaryTypeDataLoaderTask,
	"sql_task":                         ReferenceSummaryTypeSqlTask,
	"oci_dataflow_task":                ReferenceSummaryTypeOciDataflowTask,
	"pipeline_task":                    ReferenceSummaryTypePipelineTask,
	"rest_task":                        ReferenceSummaryTypeRestTask,
}

// GetReferenceSummaryTypeEnumValues Enumerates the set of values for ReferenceSummaryTypeEnum
func GetReferenceSummaryTypeEnumValues() []ReferenceSummaryTypeEnum {
	values := make([]ReferenceSummaryTypeEnum, 0)
	for _, v := range mappingReferenceSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferenceSummaryTypeEnumStringValues Enumerates the set of values in String for ReferenceSummaryTypeEnum
func GetReferenceSummaryTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATA_ASSET",
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		"ORACLE_ATP_DATA_ASSET",
		"ORACLE_ADWC_DATA_ASSET",
		"MYSQL_DATA_ASSET",
		"GENERIC_JDBC_DATA_ASSET",
		"FUSION_APP_DATA_ASSET",
		"AMAZON_S3_DATA_ASSET",
		"SCHEMA",
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"PIPELINE_TASK",
		"REST_TASK",
	}
}

// GetMappingReferenceSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferenceSummaryTypeEnum(val string) (ReferenceSummaryTypeEnum, bool) {
	enum, ok := mappingReferenceSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
