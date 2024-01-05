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

// NewInstallationSite The properties of a new Java installation site.
type NewInstallationSite struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The release version of the Java Runtime.
	ReleaseVersion *string `mandatory:"true" json:"releaseVersion"`

	// Artifact content type for the Java version.
	ArtifactContentType ArtifactContentTypeEnum `mandatory:"false" json:"artifactContentType,omitempty"`

	// Custom path to install new Java installation site.
	InstallationPath *string `mandatory:"false" json:"installationPath"`

	// Flag to install headless or headful Java installation. Only valid for Oracle Linux in OCI.
	HeadlessMode *bool `mandatory:"false" json:"headlessMode"`

	// Forces the installation request even if a more recent release is already present in the host.
	ForceInstall *bool `mandatory:"false" json:"forceInstall"`
}

func (m NewInstallationSite) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NewInstallationSite) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArtifactContentTypeEnum(string(m.ArtifactContentType)); !ok && m.ArtifactContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArtifactContentType: %s. Supported values are: %s.", m.ArtifactContentType, strings.Join(GetArtifactContentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
