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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ExadataInsightResourceStatisticsAggregation Contains resource details and current statistics
type ExadataInsightResourceStatisticsAggregation interface {
}

type exadatainsightresourcestatisticsaggregation struct {
	JsonData            []byte
	ExadataResourceType string `json:"exadataResourceType"`
}

// UnmarshalJSON unmarshals json
func (m *exadatainsightresourcestatisticsaggregation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexadatainsightresourcestatisticsaggregation exadatainsightresourcestatisticsaggregation
	s := struct {
		Model Unmarshalerexadatainsightresourcestatisticsaggregation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExadataResourceType = s.Model.ExadataResourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *exadatainsightresourcestatisticsaggregation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ExadataResourceType {
	case "STORAGE_SERVER":
		mm := ExadataStorageServerStatisticsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST":
		mm := ExadataHostStatisticsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := ExadataDatabaseStatisticsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISKGROUP":
		mm := ExadataDiskgroupStatisticsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m exadatainsightresourcestatisticsaggregation) String() string {
	return common.PointerString(m)
}

// ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum Enum with underlying type: string
type ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum string

// Set of constants representing the allowable values for ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum
const (
	ExadataInsightResourceStatisticsAggregationExadataResourceTypeDatabase      ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum = "DATABASE"
	ExadataInsightResourceStatisticsAggregationExadataResourceTypeHost          ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum = "HOST"
	ExadataInsightResourceStatisticsAggregationExadataResourceTypeStorageServer ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum = "STORAGE_SERVER"
	ExadataInsightResourceStatisticsAggregationExadataResourceTypeDiskgroup     ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum = "DISKGROUP"
)

var mappingExadataInsightResourceStatisticsAggregationExadataResourceType = map[string]ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum{
	"DATABASE":       ExadataInsightResourceStatisticsAggregationExadataResourceTypeDatabase,
	"HOST":           ExadataInsightResourceStatisticsAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": ExadataInsightResourceStatisticsAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      ExadataInsightResourceStatisticsAggregationExadataResourceTypeDiskgroup,
}

// GetExadataInsightResourceStatisticsAggregationExadataResourceTypeEnumValues Enumerates the set of values for ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum
func GetExadataInsightResourceStatisticsAggregationExadataResourceTypeEnumValues() []ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum {
	values := make([]ExadataInsightResourceStatisticsAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingExadataInsightResourceStatisticsAggregationExadataResourceType {
		values = append(values, v)
	}
	return values
}
