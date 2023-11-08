// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaDownloadReport Details about a Java download report in a tenancy.
// The report is a file in Object Storage. It contains the download records in a specific format.
type JavaDownloadReport struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Java download report.
	Id *string `mandatory:"true" json:"id"`

	// Display name for the Java download report.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The file format of the Java download report.
	Format JavaDownloadReportFormatEnum `mandatory:"true" json:"format"`

	// Approximate size of the Java download report file in bytes.
	FileSizeInBytes *int64 `mandatory:"true" json:"fileSizeInBytes"`

	// The algorithm used for calculating the checksum.
	ChecksumType ChecksumTypeEnum `mandatory:"true" json:"checksumType"`

	// The checksum value of the Java download report file.
	ChecksumValue *string `mandatory:"true" json:"checksumValue"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the Java download report.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	CreatedBy *Principal `mandatory:"true" json:"createdBy"`

	// The time the Java download report was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Java download report.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m JavaDownloadReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaDownloadReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJavaDownloadReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetJavaDownloadReportFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingChecksumTypeEnum(string(m.ChecksumType)); !ok && m.ChecksumType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChecksumType: %s. Supported values are: %s.", m.ChecksumType, strings.Join(GetChecksumTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
