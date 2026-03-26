// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataInfrastructureUsageMetrics The list of aggregated metrics for Exadata infrastructures in the fleet.
type ExadataInfrastructureUsageMetrics struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	InfrastructureId *string `mandatory:"false" json:"infrastructureId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Exadata infrastructure resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The Exadata infrastructure deployment type.
	DeploymentType ExadataInfrastructureDeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// The display name of the Exadata infrastructure.
	InfrastructureName *string `mandatory:"false" json:"infrastructureName"`

	// The lifecycle state of the Exadata infrastructure.
	State ExadataInfrastructureLifecycleStateValuesStateEnum `mandatory:"false" json:"state,omitempty"`

	// The number of Database Systems created on the Exadata infrastructure.
	NumberOfDbSystems *int `mandatory:"false" json:"numberOfDbSystems"`

	// The size of the Exadata infrastructure.
	RackSize ExadataInfrastructureUsageMetricsRackSizeEnum `mandatory:"false" json:"rackSize,omitempty"`

	// The number of storage server for the Exadata infrastructure.
	StorageServerCount *int `mandatory:"false" json:"storageServerCount"`

	// A list of the health metrics like CPU, Storage, and Memory.
	Metrics []ExadataFleetMetricDefinition `mandatory:"false" json:"metrics"`
}

func (m ExadataInfrastructureUsageMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureUsageMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataInfrastructureDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetExadataInfrastructureDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureLifecycleStateValuesStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetExadataInfrastructureLifecycleStateValuesStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureUsageMetricsRackSizeEnum(string(m.RackSize)); !ok && m.RackSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RackSize: %s. Supported values are: %s.", m.RackSize, strings.Join(GetExadataInfrastructureUsageMetricsRackSizeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureUsageMetricsRackSizeEnum Enum with underlying type: string
type ExadataInfrastructureUsageMetricsRackSizeEnum string

// Set of constants representing the allowable values for ExadataInfrastructureUsageMetricsRackSizeEnum
const (
	ExadataInfrastructureUsageMetricsRackSizeFull    ExadataInfrastructureUsageMetricsRackSizeEnum = "FULL"
	ExadataInfrastructureUsageMetricsRackSizeHalf    ExadataInfrastructureUsageMetricsRackSizeEnum = "HALF"
	ExadataInfrastructureUsageMetricsRackSizeQuarter ExadataInfrastructureUsageMetricsRackSizeEnum = "QUARTER"
	ExadataInfrastructureUsageMetricsRackSizeEighth  ExadataInfrastructureUsageMetricsRackSizeEnum = "EIGHTH"
	ExadataInfrastructureUsageMetricsRackSizeOther   ExadataInfrastructureUsageMetricsRackSizeEnum = "OTHER"
)

var mappingExadataInfrastructureUsageMetricsRackSizeEnum = map[string]ExadataInfrastructureUsageMetricsRackSizeEnum{
	"FULL":    ExadataInfrastructureUsageMetricsRackSizeFull,
	"HALF":    ExadataInfrastructureUsageMetricsRackSizeHalf,
	"QUARTER": ExadataInfrastructureUsageMetricsRackSizeQuarter,
	"EIGHTH":  ExadataInfrastructureUsageMetricsRackSizeEighth,
	"OTHER":   ExadataInfrastructureUsageMetricsRackSizeOther,
}

var mappingExadataInfrastructureUsageMetricsRackSizeEnumLowerCase = map[string]ExadataInfrastructureUsageMetricsRackSizeEnum{
	"full":    ExadataInfrastructureUsageMetricsRackSizeFull,
	"half":    ExadataInfrastructureUsageMetricsRackSizeHalf,
	"quarter": ExadataInfrastructureUsageMetricsRackSizeQuarter,
	"eighth":  ExadataInfrastructureUsageMetricsRackSizeEighth,
	"other":   ExadataInfrastructureUsageMetricsRackSizeOther,
}

// GetExadataInfrastructureUsageMetricsRackSizeEnumValues Enumerates the set of values for ExadataInfrastructureUsageMetricsRackSizeEnum
func GetExadataInfrastructureUsageMetricsRackSizeEnumValues() []ExadataInfrastructureUsageMetricsRackSizeEnum {
	values := make([]ExadataInfrastructureUsageMetricsRackSizeEnum, 0)
	for _, v := range mappingExadataInfrastructureUsageMetricsRackSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureUsageMetricsRackSizeEnumStringValues Enumerates the set of values in String for ExadataInfrastructureUsageMetricsRackSizeEnum
func GetExadataInfrastructureUsageMetricsRackSizeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
		"OTHER",
	}
}

// GetMappingExadataInfrastructureUsageMetricsRackSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureUsageMetricsRackSizeEnum(val string) (ExadataInfrastructureUsageMetricsRackSizeEnum, bool) {
	enum, ok := mappingExadataInfrastructureUsageMetricsRackSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
