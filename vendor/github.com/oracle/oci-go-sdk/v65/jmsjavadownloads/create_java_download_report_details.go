// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateJavaDownloadReportDetails Attributes to create a Java download report.
type CreateJavaDownloadReportDetails struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) here should be
	// the tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The format of the report that is generated.
	Format JavaDownloadReportFormatEnum `mandatory:"true" json:"format"`

	// The start time from when download records have to be included (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end time until when the download records have to be included (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The property to be used for sorting the records.
	SortBy JavaDownloadRecordSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order for the records.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJavaDownloadReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJavaDownloadReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJavaDownloadReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetJavaDownloadReportFormatEnumStringValues(), ",")))
	}

	if _, ok := GetMappingJavaDownloadRecordSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetJavaDownloadRecordSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
