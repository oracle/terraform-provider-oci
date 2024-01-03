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

// Reference Reference contains application configuration information.
type Reference struct {

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
	Type ReferenceTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The new reference object to use instead of the original reference. For example, this can be a data asset reference.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The application key of the reference object.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// List of published objects where this is used.
	UsedBy []ReferenceUsedBy `mandatory:"false" json:"usedBy"`

	// List of references that are dependent on this reference.
	ChildReferences []ChildReference `mandatory:"false" json:"childReferences"`
}

func (m Reference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Reference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReferenceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetReferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferenceTypeEnum Enum with underlying type: string
type ReferenceTypeEnum string

// Set of constants representing the allowable values for ReferenceTypeEnum
const (
	ReferenceTypeOracleDataAsset              ReferenceTypeEnum = "ORACLE_DATA_ASSET"
	ReferenceTypeOracleObjectStorageDataAsset ReferenceTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	ReferenceTypeOracleAtpDataAsset           ReferenceTypeEnum = "ORACLE_ATP_DATA_ASSET"
	ReferenceTypeOracleAdwcDataAsset          ReferenceTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	ReferenceTypeMysqlDataAsset               ReferenceTypeEnum = "MYSQL_DATA_ASSET"
	ReferenceTypeGenericJdbcDataAsset         ReferenceTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	ReferenceTypeFusionAppDataAsset           ReferenceTypeEnum = "FUSION_APP_DATA_ASSET"
	ReferenceTypeAmazonS3DataAsset            ReferenceTypeEnum = "AMAZON_S3_DATA_ASSET"
	ReferenceTypeSchema                       ReferenceTypeEnum = "SCHEMA"
	ReferenceTypeIntegrationTask              ReferenceTypeEnum = "INTEGRATION_TASK"
	ReferenceTypeDataLoaderTask               ReferenceTypeEnum = "DATA_LOADER_TASK"
	ReferenceTypeSqlTask                      ReferenceTypeEnum = "SQL_TASK"
	ReferenceTypeOciDataflowTask              ReferenceTypeEnum = "OCI_DATAFLOW_TASK"
	ReferenceTypePipelineTask                 ReferenceTypeEnum = "PIPELINE_TASK"
	ReferenceTypeRestTask                     ReferenceTypeEnum = "REST_TASK"
)

var mappingReferenceTypeEnum = map[string]ReferenceTypeEnum{
	"ORACLE_DATA_ASSET":                ReferenceTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": ReferenceTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            ReferenceTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           ReferenceTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 ReferenceTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          ReferenceTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            ReferenceTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             ReferenceTypeAmazonS3DataAsset,
	"SCHEMA":                           ReferenceTypeSchema,
	"INTEGRATION_TASK":                 ReferenceTypeIntegrationTask,
	"DATA_LOADER_TASK":                 ReferenceTypeDataLoaderTask,
	"SQL_TASK":                         ReferenceTypeSqlTask,
	"OCI_DATAFLOW_TASK":                ReferenceTypeOciDataflowTask,
	"PIPELINE_TASK":                    ReferenceTypePipelineTask,
	"REST_TASK":                        ReferenceTypeRestTask,
}

var mappingReferenceTypeEnumLowerCase = map[string]ReferenceTypeEnum{
	"oracle_data_asset":                ReferenceTypeOracleDataAsset,
	"oracle_object_storage_data_asset": ReferenceTypeOracleObjectStorageDataAsset,
	"oracle_atp_data_asset":            ReferenceTypeOracleAtpDataAsset,
	"oracle_adwc_data_asset":           ReferenceTypeOracleAdwcDataAsset,
	"mysql_data_asset":                 ReferenceTypeMysqlDataAsset,
	"generic_jdbc_data_asset":          ReferenceTypeGenericJdbcDataAsset,
	"fusion_app_data_asset":            ReferenceTypeFusionAppDataAsset,
	"amazon_s3_data_asset":             ReferenceTypeAmazonS3DataAsset,
	"schema":                           ReferenceTypeSchema,
	"integration_task":                 ReferenceTypeIntegrationTask,
	"data_loader_task":                 ReferenceTypeDataLoaderTask,
	"sql_task":                         ReferenceTypeSqlTask,
	"oci_dataflow_task":                ReferenceTypeOciDataflowTask,
	"pipeline_task":                    ReferenceTypePipelineTask,
	"rest_task":                        ReferenceTypeRestTask,
}

// GetReferenceTypeEnumValues Enumerates the set of values for ReferenceTypeEnum
func GetReferenceTypeEnumValues() []ReferenceTypeEnum {
	values := make([]ReferenceTypeEnum, 0)
	for _, v := range mappingReferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferenceTypeEnumStringValues Enumerates the set of values in String for ReferenceTypeEnum
func GetReferenceTypeEnumStringValues() []string {
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

// GetMappingReferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferenceTypeEnum(val string) (ReferenceTypeEnum, bool) {
	enum, ok := mappingReferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
