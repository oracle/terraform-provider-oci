// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RecalledData This is the information about recalled data
type RecalledData struct {

	// This is the end of the time range of the related data
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the start of the time range of the related data
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the time when the first recall operation was started for this RecalledData
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// This is the status of the recall
	Status RecalledDataStatusEnum `mandatory:"true" json:"status"`

	// This is the number of recall operations for this recall.  Note one RecalledData can be merged from the results
	// of several recall operations if the time duration of the results of the recall operations overlap.
	RecallCount *int `mandatory:"true" json:"recallCount"`

	// This is the size in bytes
	StorageUsageInBytes *int64 `mandatory:"true" json:"storageUsageInBytes"`
}

func (m RecalledData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecalledData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecalledDataStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRecalledDataStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecalledDataStatusEnum Enum with underlying type: string
type RecalledDataStatusEnum string

// Set of constants representing the allowable values for RecalledDataStatusEnum
const (
	RecalledDataStatusRecalled RecalledDataStatusEnum = "RECALLED"
	RecalledDataStatusPending  RecalledDataStatusEnum = "PENDING"
)

var mappingRecalledDataStatusEnum = map[string]RecalledDataStatusEnum{
	"RECALLED": RecalledDataStatusRecalled,
	"PENDING":  RecalledDataStatusPending,
}

// GetRecalledDataStatusEnumValues Enumerates the set of values for RecalledDataStatusEnum
func GetRecalledDataStatusEnumValues() []RecalledDataStatusEnum {
	values := make([]RecalledDataStatusEnum, 0)
	for _, v := range mappingRecalledDataStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRecalledDataStatusEnumStringValues Enumerates the set of values in String for RecalledDataStatusEnum
func GetRecalledDataStatusEnumStringValues() []string {
	return []string{
		"RECALLED",
		"PENDING",
	}
}

// GetMappingRecalledDataStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecalledDataStatusEnum(val string) (RecalledDataStatusEnum, bool) {
	mappingRecalledDataStatusEnumIgnoreCase := make(map[string]RecalledDataStatusEnum)
	for k, v := range mappingRecalledDataStatusEnum {
		mappingRecalledDataStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRecalledDataStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
