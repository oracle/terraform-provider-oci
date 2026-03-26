// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChargebackPlanReport A chargeback plan report that allows Ops Insights services to showcase chargeback costs.
type ChargebackPlanReport struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Chargeback plan report.
	ReportId *string `mandatory:"true" json:"reportId"`

	// The chargeback plan report name.
	ReportName *string `mandatory:"true" json:"reportName"`

	// Defines the type of resource (example: EXADATA, HOST)
	ResourceType ChargebackPlanReportResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Chargeback plan report.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The date and time the chargeback plan was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time chargeback plan was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	ReportProperties *ReportPropertyDetails `mandatory:"true" json:"reportProperties"`
}

func (m ChargebackPlanReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChargebackPlanReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChargebackPlanReportResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetChargebackPlanReportResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChargebackPlanReportResourceTypeEnum Enum with underlying type: string
type ChargebackPlanReportResourceTypeEnum string

// Set of constants representing the allowable values for ChargebackPlanReportResourceTypeEnum
const (
	ChargebackPlanReportResourceTypeExadataInsight  ChargebackPlanReportResourceTypeEnum = "EXADATA_INSIGHT"
	ChargebackPlanReportResourceTypeDatabaseInsight ChargebackPlanReportResourceTypeEnum = "DATABASE_INSIGHT"
	ChargebackPlanReportResourceTypeHostInsight     ChargebackPlanReportResourceTypeEnum = "HOST_INSIGHT"
)

var mappingChargebackPlanReportResourceTypeEnum = map[string]ChargebackPlanReportResourceTypeEnum{
	"EXADATA_INSIGHT":  ChargebackPlanReportResourceTypeExadataInsight,
	"DATABASE_INSIGHT": ChargebackPlanReportResourceTypeDatabaseInsight,
	"HOST_INSIGHT":     ChargebackPlanReportResourceTypeHostInsight,
}

var mappingChargebackPlanReportResourceTypeEnumLowerCase = map[string]ChargebackPlanReportResourceTypeEnum{
	"exadata_insight":  ChargebackPlanReportResourceTypeExadataInsight,
	"database_insight": ChargebackPlanReportResourceTypeDatabaseInsight,
	"host_insight":     ChargebackPlanReportResourceTypeHostInsight,
}

// GetChargebackPlanReportResourceTypeEnumValues Enumerates the set of values for ChargebackPlanReportResourceTypeEnum
func GetChargebackPlanReportResourceTypeEnumValues() []ChargebackPlanReportResourceTypeEnum {
	values := make([]ChargebackPlanReportResourceTypeEnum, 0)
	for _, v := range mappingChargebackPlanReportResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChargebackPlanReportResourceTypeEnumStringValues Enumerates the set of values in String for ChargebackPlanReportResourceTypeEnum
func GetChargebackPlanReportResourceTypeEnumStringValues() []string {
	return []string{
		"EXADATA_INSIGHT",
		"DATABASE_INSIGHT",
		"HOST_INSIGHT",
	}
}

// GetMappingChargebackPlanReportResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChargebackPlanReportResourceTypeEnum(val string) (ChargebackPlanReportResourceTypeEnum, bool) {
	enum, ok := mappingChargebackPlanReportResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
