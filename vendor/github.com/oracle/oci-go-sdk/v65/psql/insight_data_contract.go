// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InsightDataContract Describes the response data format returned for an insight type.
type InsightDataContract struct {

	// Indicates the structure of the insight data payload.
	Kind InsightDataContractKindEnum `mandatory:"true" json:"kind"`

	// Optional unit associated with numeric values.
	Unit *string `mandatory:"false" json:"unit"`
}

func (m InsightDataContract) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InsightDataContract) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInsightDataContractKindEnum(string(m.Kind)); !ok && m.Kind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kind: %s. Supported values are: %s.", m.Kind, strings.Join(GetInsightDataContractKindEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InsightDataContractKindEnum Enum with underlying type: string
type InsightDataContractKindEnum string

// Set of constants representing the allowable values for InsightDataContractKindEnum
const (
	InsightDataContractKindTimeSeries InsightDataContractKindEnum = "TIME_SERIES"
	InsightDataContractKindTable      InsightDataContractKindEnum = "TABLE"
	InsightDataContractKindSummary    InsightDataContractKindEnum = "SUMMARY"
)

var mappingInsightDataContractKindEnum = map[string]InsightDataContractKindEnum{
	"TIME_SERIES": InsightDataContractKindTimeSeries,
	"TABLE":       InsightDataContractKindTable,
	"SUMMARY":     InsightDataContractKindSummary,
}

var mappingInsightDataContractKindEnumLowerCase = map[string]InsightDataContractKindEnum{
	"time_series": InsightDataContractKindTimeSeries,
	"table":       InsightDataContractKindTable,
	"summary":     InsightDataContractKindSummary,
}

// GetInsightDataContractKindEnumValues Enumerates the set of values for InsightDataContractKindEnum
func GetInsightDataContractKindEnumValues() []InsightDataContractKindEnum {
	values := make([]InsightDataContractKindEnum, 0)
	for _, v := range mappingInsightDataContractKindEnum {
		values = append(values, v)
	}
	return values
}

// GetInsightDataContractKindEnumStringValues Enumerates the set of values in String for InsightDataContractKindEnum
func GetInsightDataContractKindEnumStringValues() []string {
	return []string{
		"TIME_SERIES",
		"TABLE",
		"SUMMARY",
	}
}

// GetMappingInsightDataContractKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInsightDataContractKindEnum(val string) (InsightDataContractKindEnum, bool) {
	enum, ok := mappingInsightDataContractKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
