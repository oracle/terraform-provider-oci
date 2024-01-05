// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateJavaLicenseAcceptanceRecordDetails The attributes to create a new JavaLicenseAcceptanceRecord.
type CreateJavaLicenseAcceptanceRecordDetails struct {

	// The tenancy OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the user accepting the license.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// License type for the Java version.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// Status of license acceptance.
	LicenseAcceptanceStatus LicenseAcceptanceStatusEnum `mandatory:"true" json:"licenseAcceptanceStatus"`
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
