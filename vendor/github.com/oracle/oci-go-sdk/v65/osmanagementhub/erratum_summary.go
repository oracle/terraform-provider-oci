// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ErratumSummary Important changes for software. This can include security advisories, bug fixes, or enhancements.
type ErratumSummary struct {

	// Advisory name.
	Name *string `mandatory:"true" json:"name"`

	// Summary description of the erratum.
	Synopsis *string `mandatory:"false" json:"synopsis"`

	// Date the erratum was issued, as described
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeIssued *common.SDKTime `mandatory:"false" json:"timeIssued"`

	// Most recent date the erratum was updated, as described
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type of the erratum.
	ClassificationType ClassificationTypesEnum `mandatory:"false" json:"classificationType,omitempty"`

	// List of CVEs applicable to this erratum.
	RelatedCves []string `mandatory:"false" json:"relatedCves"`

	// List of affected OS families.
	OsFamilies []OsFamilyEnum `mandatory:"false" json:"osFamilies"`

	// The severity advisory. Only valid for security type advisories.
	AdvisorySeverity AdvisorySeverityEnum `mandatory:"false" json:"advisorySeverity,omitempty"`
}

func (m ErratumSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ErratumSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClassificationTypesEnum(string(m.ClassificationType)); !ok && m.ClassificationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", m.ClassificationType, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAdvisorySeverityEnum(string(m.AdvisorySeverity)); !ok && m.AdvisorySeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisorySeverity: %s. Supported values are: %s.", m.AdvisorySeverity, strings.Join(GetAdvisorySeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
