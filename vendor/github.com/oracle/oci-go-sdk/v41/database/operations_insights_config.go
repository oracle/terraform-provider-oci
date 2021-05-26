// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// OperationsInsightsConfig The configuration of Operations Insights for the external database
type OperationsInsightsConfig struct {

	// The status of Operations Insights
	OperationsInsightsStatus OperationsInsightsConfigOperationsInsightsStatusEnum `mandatory:"true" json:"operationsInsightsStatus"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	OperationsInsightsConnectorId *string `mandatory:"false" json:"operationsInsightsConnectorId"`
}

func (m OperationsInsightsConfig) String() string {
	return common.PointerString(m)
}

// OperationsInsightsConfigOperationsInsightsStatusEnum Enum with underlying type: string
type OperationsInsightsConfigOperationsInsightsStatusEnum string

// Set of constants representing the allowable values for OperationsInsightsConfigOperationsInsightsStatusEnum
const (
	OperationsInsightsConfigOperationsInsightsStatusEnabling        OperationsInsightsConfigOperationsInsightsStatusEnum = "ENABLING"
	OperationsInsightsConfigOperationsInsightsStatusEnabled         OperationsInsightsConfigOperationsInsightsStatusEnum = "ENABLED"
	OperationsInsightsConfigOperationsInsightsStatusDisabling       OperationsInsightsConfigOperationsInsightsStatusEnum = "DISABLING"
	OperationsInsightsConfigOperationsInsightsStatusNotEnabled      OperationsInsightsConfigOperationsInsightsStatusEnum = "NOT_ENABLED"
	OperationsInsightsConfigOperationsInsightsStatusFailedEnabling  OperationsInsightsConfigOperationsInsightsStatusEnum = "FAILED_ENABLING"
	OperationsInsightsConfigOperationsInsightsStatusFailedDisabling OperationsInsightsConfigOperationsInsightsStatusEnum = "FAILED_DISABLING"
)

var mappingOperationsInsightsConfigOperationsInsightsStatus = map[string]OperationsInsightsConfigOperationsInsightsStatusEnum{
	"ENABLING":         OperationsInsightsConfigOperationsInsightsStatusEnabling,
	"ENABLED":          OperationsInsightsConfigOperationsInsightsStatusEnabled,
	"DISABLING":        OperationsInsightsConfigOperationsInsightsStatusDisabling,
	"NOT_ENABLED":      OperationsInsightsConfigOperationsInsightsStatusNotEnabled,
	"FAILED_ENABLING":  OperationsInsightsConfigOperationsInsightsStatusFailedEnabling,
	"FAILED_DISABLING": OperationsInsightsConfigOperationsInsightsStatusFailedDisabling,
}

// GetOperationsInsightsConfigOperationsInsightsStatusEnumValues Enumerates the set of values for OperationsInsightsConfigOperationsInsightsStatusEnum
func GetOperationsInsightsConfigOperationsInsightsStatusEnumValues() []OperationsInsightsConfigOperationsInsightsStatusEnum {
	values := make([]OperationsInsightsConfigOperationsInsightsStatusEnum, 0)
	for _, v := range mappingOperationsInsightsConfigOperationsInsightsStatus {
		values = append(values, v)
	}
	return values
}
