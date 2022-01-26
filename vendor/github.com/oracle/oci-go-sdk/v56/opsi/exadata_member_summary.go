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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingExadataMemberSummaryEntityType = map[string]ExadataMemberSummaryEntityTypeEnum{
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
	for _, v := range mappingExadataMemberSummaryEntityType {
		values = append(values, v)
	}
	return values
}
