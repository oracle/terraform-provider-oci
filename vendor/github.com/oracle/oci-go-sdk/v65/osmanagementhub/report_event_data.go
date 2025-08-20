// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReportEventData Data related to the report event.
type ReportEventData struct {

	// status of the report event.
	ReportStatus ReportEventDataReportStatusEnum `mandatory:"true" json:"reportStatus"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m ReportEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReportEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportEventDataReportStatusEnum(string(m.ReportStatus)); !ok && m.ReportStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportStatus: %s. Supported values are: %s.", m.ReportStatus, strings.Join(GetReportEventDataReportStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportEventDataReportStatusEnum Enum with underlying type: string
type ReportEventDataReportStatusEnum string

// Set of constants representing the allowable values for ReportEventDataReportStatusEnum
const (
	ReportEventDataReportStatusStarted   ReportEventDataReportStatusEnum = "REPORT_STARTED"
	ReportEventDataReportStatusSucceeded ReportEventDataReportStatusEnum = "REPORT_SUCCEEDED"
	ReportEventDataReportStatusFailed    ReportEventDataReportStatusEnum = "REPORT_FAILED"
)

var mappingReportEventDataReportStatusEnum = map[string]ReportEventDataReportStatusEnum{
	"REPORT_STARTED":   ReportEventDataReportStatusStarted,
	"REPORT_SUCCEEDED": ReportEventDataReportStatusSucceeded,
	"REPORT_FAILED":    ReportEventDataReportStatusFailed,
}

var mappingReportEventDataReportStatusEnumLowerCase = map[string]ReportEventDataReportStatusEnum{
	"report_started":   ReportEventDataReportStatusStarted,
	"report_succeeded": ReportEventDataReportStatusSucceeded,
	"report_failed":    ReportEventDataReportStatusFailed,
}

// GetReportEventDataReportStatusEnumValues Enumerates the set of values for ReportEventDataReportStatusEnum
func GetReportEventDataReportStatusEnumValues() []ReportEventDataReportStatusEnum {
	values := make([]ReportEventDataReportStatusEnum, 0)
	for _, v := range mappingReportEventDataReportStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReportEventDataReportStatusEnumStringValues Enumerates the set of values in String for ReportEventDataReportStatusEnum
func GetReportEventDataReportStatusEnumStringValues() []string {
	return []string{
		"REPORT_STARTED",
		"REPORT_SUCCEEDED",
		"REPORT_FAILED",
	}
}

// GetMappingReportEventDataReportStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportEventDataReportStatusEnum(val string) (ReportEventDataReportStatusEnum, bool) {
	enum, ok := mappingReportEventDataReportStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
