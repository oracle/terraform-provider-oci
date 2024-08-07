// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaRelease Metadata associated with a specific release of Java. Includes the artifact details.
type JavaRelease struct {

	// Java release version identifier.
	ReleaseVersion *string `mandatory:"true" json:"releaseVersion"`

	// Java release family identifier.
	FamilyVersion *string `mandatory:"true" json:"familyVersion"`

	// The security status of the Java version.
	SecurityStatus JreSecurityStatusEnum `mandatory:"true" json:"securityStatus"`

	// Release category of the Java version.
	ReleaseType ReleaseTypeEnum `mandatory:"true" json:"releaseType"`

	// License type for the Java version.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// The release date of the Java version (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	ReleaseDate *common.SDKTime `mandatory:"true" json:"releaseDate"`

	// Release notes associated with the Java version.
	ReleaseNotesUrl *string `mandatory:"true" json:"releaseNotesUrl"`

	// Artifact content types for the Java version.
	ArtifactContentTypes []ArtifactContentTypeEnum `mandatory:"true" json:"artifactContentTypes"`

	// Parent Java release version identifier. This is applicable for BPR releases.
	ParentReleaseVersion *string `mandatory:"false" json:"parentReleaseVersion"`

	FamilyDetails *JavaFamily `mandatory:"false" json:"familyDetails"`

	LicenseDetails *JavaLicense `mandatory:"false" json:"licenseDetails"`

	// List of My Oracle Support(MoS) patches available for this release.
	// This information is only available for `BPR` release type.
	MosPatches []PatchDetail `mandatory:"false" json:"mosPatches"`

	// The number of days since this release has been under the security baseline.
	DaysUnderSecurityBaseline *int `mandatory:"false" json:"daysUnderSecurityBaseline"`

	// List of Java artifacts.
	Artifacts []JavaArtifact `mandatory:"false" json:"artifacts"`
}

func (m JavaRelease) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaRelease) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJreSecurityStatusEnum(string(m.SecurityStatus)); !ok && m.SecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityStatus: %s. Supported values are: %s.", m.SecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReleaseTypeEnum(string(m.ReleaseType)); !ok && m.ReleaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReleaseType: %s. Supported values are: %s.", m.ReleaseType, strings.Join(GetReleaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	for _, val := range m.ArtifactContentTypes {
		if _, ok := GetMappingArtifactContentTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArtifactContentTypes: %s. Supported values are: %s.", val, strings.Join(GetArtifactContentTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
