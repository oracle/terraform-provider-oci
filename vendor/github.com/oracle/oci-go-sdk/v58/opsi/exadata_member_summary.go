// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExadataMemberSummary Lists name, display name and type of exadata member.
type ExadataMemberSummary struct {

	// Name of exadata member target
	Name *string `mandatory:"true" json:"name"`

	// Display Name of exadata member target
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Entity type of exadata member target
	EntityType ExadataMemberSummaryEntityTypeEnum `mandatory:"true" json:"entityType"`
}

func (m ExadataMemberSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataMemberSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataMemberSummaryEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetExadataMemberSummaryEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataMemberSummaryEntityTypeEnum Enum with underlying type: string
type ExadataMemberSummaryEntityTypeEnum string

// Set of constants representing the allowable values for ExadataMemberSummaryEntityTypeEnum
const (
	ExadataMemberSummaryEntityTypeDatabase         ExadataMemberSummaryEntityTypeEnum = "DATABASE"
	ExadataMemberSummaryEntityTypeIlomServer       ExadataMemberSummaryEntityTypeEnum = "ILOM_SERVER"
	ExadataMemberSummaryEntityTypePdu              ExadataMemberSummaryEntityTypeEnum = "PDU"
	ExadataMemberSummaryEntityTypeStorageServer    ExadataMemberSummaryEntityTypeEnum = "STORAGE_SERVER"
	ExadataMemberSummaryEntityTypeClusterAsm       ExadataMemberSummaryEntityTypeEnum = "CLUSTER_ASM"
	ExadataMemberSummaryEntityTypeInfinibandSwitch ExadataMemberSummaryEntityTypeEnum = "INFINIBAND_SWITCH"
	ExadataMemberSummaryEntityTypeEthernetSwitch   ExadataMemberSummaryEntityTypeEnum = "ETHERNET_SWITCH"
)

var mappingExadataMemberSummaryEntityTypeEnum = map[string]ExadataMemberSummaryEntityTypeEnum{
	"DATABASE":          ExadataMemberSummaryEntityTypeDatabase,
	"ILOM_SERVER":       ExadataMemberSummaryEntityTypeIlomServer,
	"PDU":               ExadataMemberSummaryEntityTypePdu,
	"STORAGE_SERVER":    ExadataMemberSummaryEntityTypeStorageServer,
	"CLUSTER_ASM":       ExadataMemberSummaryEntityTypeClusterAsm,
	"INFINIBAND_SWITCH": ExadataMemberSummaryEntityTypeInfinibandSwitch,
	"ETHERNET_SWITCH":   ExadataMemberSummaryEntityTypeEthernetSwitch,
}

// GetExadataMemberSummaryEntityTypeEnumValues Enumerates the set of values for ExadataMemberSummaryEntityTypeEnum
func GetExadataMemberSummaryEntityTypeEnumValues() []ExadataMemberSummaryEntityTypeEnum {
	values := make([]ExadataMemberSummaryEntityTypeEnum, 0)
	for _, v := range mappingExadataMemberSummaryEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataMemberSummaryEntityTypeEnumStringValues Enumerates the set of values in String for ExadataMemberSummaryEntityTypeEnum
func GetExadataMemberSummaryEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"ILOM_SERVER",
		"PDU",
		"STORAGE_SERVER",
		"CLUSTER_ASM",
		"INFINIBAND_SWITCH",
		"ETHERNET_SWITCH",
	}
}

// GetMappingExadataMemberSummaryEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataMemberSummaryEntityTypeEnum(val string) (ExadataMemberSummaryEntityTypeEnum, bool) {
	mappingExadataMemberSummaryEntityTypeEnumIgnoreCase := make(map[string]ExadataMemberSummaryEntityTypeEnum)
	for k, v := range mappingExadataMemberSummaryEntityTypeEnum {
		mappingExadataMemberSummaryEntityTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExadataMemberSummaryEntityTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
