// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaFamilySummary A summary of the Java release family information.
// A Java release family is typically a major version in the Java version identifier.
type JavaFamilySummary struct {

	// The Java release family identifier.
	FamilyVersion *string `mandatory:"true" json:"familyVersion"`

	// The display name of the release family.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// This indicates the support category for the Java release family.
	SupportType SupportTypeEnum `mandatory:"true" json:"supportType"`

	// The End of Support Life (EOSL) date of the Java release family (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	EndOfSupportLifeDate *common.SDKTime `mandatory:"true" json:"endOfSupportLifeDate"`

	// Link to access the documentation for the release.
	DocUrl *string `mandatory:"true" json:"docUrl"`

	// Latest Java release version in the family.
	LatestReleaseVersion *string `mandatory:"true" json:"latestReleaseVersion"`

	// Whether or not this Java release family is under active support.
	// Refer Java Support Roadmap (https://www.oracle.com/java/technologies/java-se-support-roadmap.html) for more details.
	IsSupportedVersion *bool `mandatory:"true" json:"isSupportedVersion"`
}

func (m JavaFamilySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaFamilySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSupportTypeEnum(string(m.SupportType)); !ok && m.SupportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportType: %s. Supported values are: %s.", m.SupportType, strings.Join(GetSupportTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
