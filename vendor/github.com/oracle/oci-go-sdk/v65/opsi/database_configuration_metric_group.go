// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseConfigurationMetricGroup Supported configuration metric groups for database capacity planning service.
type DatabaseConfigurationMetricGroup interface {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	GetTimeCollected() *common.SDKTime
}

type databaseconfigurationmetricgroup struct {
	JsonData      []byte
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`
	MetricName    string          `json:"metricName"`
}

// UnmarshalJSON unmarshals json
func (m *databaseconfigurationmetricgroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseconfigurationmetricgroup databaseconfigurationmetricgroup
	s := struct {
		Model Unmarshalerdatabaseconfigurationmetricgroup
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeCollected = s.Model.TimeCollected
	m.MetricName = s.Model.MetricName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseconfigurationmetricgroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricName {
	case "DB_OS_CONFIG_INSTANCE":
		mm := DbosConfigInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_EXTERNAL_INSTANCE":
		mm := DbExternalInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_PARAMETERS":
		mm := DbParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_CONNECTION_STATUS":
		mm := DbConnectionStatus{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_RESOURCE_ALLOCATION":
		mm := HostAllocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADATA_CELL_CONFIG":
		mm := ExadataCellConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_EXTERNAL_PROPERTIES":
		mm := DbExternalProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASM_ENTITY":
		mm := ExadataAsmEntity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseConfigurationMetricGroup: %s.", m.MetricName)
		return *m, nil
	}
}

// GetTimeCollected returns TimeCollected
func (m databaseconfigurationmetricgroup) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m databaseconfigurationmetricgroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databaseconfigurationmetricgroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseConfigurationMetricGroupMetricNameEnum Enum with underlying type: string
type DatabaseConfigurationMetricGroupMetricNameEnum string

// Set of constants representing the allowable values for DatabaseConfigurationMetricGroupMetricNameEnum
const (
	DatabaseConfigurationMetricGroupMetricNameDbExternalProperties   DatabaseConfigurationMetricGroupMetricNameEnum = "DB_EXTERNAL_PROPERTIES"
	DatabaseConfigurationMetricGroupMetricNameDbExternalInstance     DatabaseConfigurationMetricGroupMetricNameEnum = "DB_EXTERNAL_INSTANCE"
	DatabaseConfigurationMetricGroupMetricNameDbOsConfigInstance     DatabaseConfigurationMetricGroupMetricNameEnum = "DB_OS_CONFIG_INSTANCE"
	DatabaseConfigurationMetricGroupMetricNameDbParameters           DatabaseConfigurationMetricGroupMetricNameEnum = "DB_PARAMETERS"
	DatabaseConfigurationMetricGroupMetricNameDbConnectionStatus     DatabaseConfigurationMetricGroupMetricNameEnum = "DB_CONNECTION_STATUS"
	DatabaseConfigurationMetricGroupMetricNameHostResourceAllocation DatabaseConfigurationMetricGroupMetricNameEnum = "HOST_RESOURCE_ALLOCATION"
	DatabaseConfigurationMetricGroupMetricNameAsmEntity              DatabaseConfigurationMetricGroupMetricNameEnum = "ASM_ENTITY"
	DatabaseConfigurationMetricGroupMetricNameExadataCellConfig      DatabaseConfigurationMetricGroupMetricNameEnum = "EXADATA_CELL_CONFIG"
)

var mappingDatabaseConfigurationMetricGroupMetricNameEnum = map[string]DatabaseConfigurationMetricGroupMetricNameEnum{
	"DB_EXTERNAL_PROPERTIES":   DatabaseConfigurationMetricGroupMetricNameDbExternalProperties,
	"DB_EXTERNAL_INSTANCE":     DatabaseConfigurationMetricGroupMetricNameDbExternalInstance,
	"DB_OS_CONFIG_INSTANCE":    DatabaseConfigurationMetricGroupMetricNameDbOsConfigInstance,
	"DB_PARAMETERS":            DatabaseConfigurationMetricGroupMetricNameDbParameters,
	"DB_CONNECTION_STATUS":     DatabaseConfigurationMetricGroupMetricNameDbConnectionStatus,
	"HOST_RESOURCE_ALLOCATION": DatabaseConfigurationMetricGroupMetricNameHostResourceAllocation,
	"ASM_ENTITY":               DatabaseConfigurationMetricGroupMetricNameAsmEntity,
	"EXADATA_CELL_CONFIG":      DatabaseConfigurationMetricGroupMetricNameExadataCellConfig,
}

var mappingDatabaseConfigurationMetricGroupMetricNameEnumLowerCase = map[string]DatabaseConfigurationMetricGroupMetricNameEnum{
	"db_external_properties":   DatabaseConfigurationMetricGroupMetricNameDbExternalProperties,
	"db_external_instance":     DatabaseConfigurationMetricGroupMetricNameDbExternalInstance,
	"db_os_config_instance":    DatabaseConfigurationMetricGroupMetricNameDbOsConfigInstance,
	"db_parameters":            DatabaseConfigurationMetricGroupMetricNameDbParameters,
	"db_connection_status":     DatabaseConfigurationMetricGroupMetricNameDbConnectionStatus,
	"host_resource_allocation": DatabaseConfigurationMetricGroupMetricNameHostResourceAllocation,
	"asm_entity":               DatabaseConfigurationMetricGroupMetricNameAsmEntity,
	"exadata_cell_config":      DatabaseConfigurationMetricGroupMetricNameExadataCellConfig,
}

// GetDatabaseConfigurationMetricGroupMetricNameEnumValues Enumerates the set of values for DatabaseConfigurationMetricGroupMetricNameEnum
func GetDatabaseConfigurationMetricGroupMetricNameEnumValues() []DatabaseConfigurationMetricGroupMetricNameEnum {
	values := make([]DatabaseConfigurationMetricGroupMetricNameEnum, 0)
	for _, v := range mappingDatabaseConfigurationMetricGroupMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConfigurationMetricGroupMetricNameEnumStringValues Enumerates the set of values in String for DatabaseConfigurationMetricGroupMetricNameEnum
func GetDatabaseConfigurationMetricGroupMetricNameEnumStringValues() []string {
	return []string{
		"DB_EXTERNAL_PROPERTIES",
		"DB_EXTERNAL_INSTANCE",
		"DB_OS_CONFIG_INSTANCE",
		"DB_PARAMETERS",
		"DB_CONNECTION_STATUS",
		"HOST_RESOURCE_ALLOCATION",
		"ASM_ENTITY",
		"EXADATA_CELL_CONFIG",
	}
}

// GetMappingDatabaseConfigurationMetricGroupMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConfigurationMetricGroupMetricNameEnum(val string) (DatabaseConfigurationMetricGroupMetricNameEnum, bool) {
	enum, ok := mappingDatabaseConfigurationMetricGroupMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
