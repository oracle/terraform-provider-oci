// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateJavaLicenseAcceptanceRecordDetails The attributes to create a new JavaLicenseAcceptanceRecord.
type CreateJavaLicenseAcceptanceRecordDetails struct {

	// The tenancy OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user accepting the license.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// License type for the Java version.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// Status of license acceptance.
	LicenseAcceptanceStatus LicenseAcceptanceStatusEnum `mandatory:"true" json:"licenseAcceptanceStatus"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJavaLicenseAcceptanceRecordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJavaLicenseAcceptanceRecordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseAcceptanceStatusEnum(string(m.LicenseAcceptanceStatus)); !ok && m.LicenseAcceptanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseAcceptanceStatus: %s. Supported values are: %s.", m.LicenseAcceptanceStatus, strings.Join(GetLicenseAcceptanceStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
