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

// OdspInsight Details for a single ODSP insight.
type OdspInsight struct {

	// Type of Insight collected for the database system.
	InsightType OdspInsightInsightTypeEnum `mandatory:"true" json:"insightType"`

	// Retention period for Insight data, in days. Current supported value is 7 days. the system default is 7 days.
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`
}

func (m OdspInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdspInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdspInsightInsightTypeEnum(string(m.InsightType)); !ok && m.InsightType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightType: %s. Supported values are: %s.", m.InsightType, strings.Join(GetOdspInsightInsightTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdspInsightInsightTypeEnum Enum with underlying type: string
type OdspInsightInsightTypeEnum string

// Set of constants representing the allowable values for OdspInsightInsightTypeEnum
const (
	OdspInsightInsightTypeQueryInsight OdspInsightInsightTypeEnum = "QUERY_INSIGHT"
)

var mappingOdspInsightInsightTypeEnum = map[string]OdspInsightInsightTypeEnum{
	"QUERY_INSIGHT": OdspInsightInsightTypeQueryInsight,
}

var mappingOdspInsightInsightTypeEnumLowerCase = map[string]OdspInsightInsightTypeEnum{
	"query_insight": OdspInsightInsightTypeQueryInsight,
}

// GetOdspInsightInsightTypeEnumValues Enumerates the set of values for OdspInsightInsightTypeEnum
func GetOdspInsightInsightTypeEnumValues() []OdspInsightInsightTypeEnum {
	values := make([]OdspInsightInsightTypeEnum, 0)
	for _, v := range mappingOdspInsightInsightTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOdspInsightInsightTypeEnumStringValues Enumerates the set of values in String for OdspInsightInsightTypeEnum
func GetOdspInsightInsightTypeEnumStringValues() []string {
	return []string{
		"QUERY_INSIGHT",
	}
}

// GetMappingOdspInsightInsightTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdspInsightInsightTypeEnum(val string) (OdspInsightInsightTypeEnum, bool) {
	enum, ok := mappingOdspInsightInsightTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
