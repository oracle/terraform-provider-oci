// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArtifactVersionSummary Summary of an Artifact Version, returned as a result of a List operation.
type ArtifactVersionSummary struct {

	// Versionless Package URL (purl) used to identify an Artifact (e.g. pkg:maven/com.oracle.oci.sdk/oci-java-sdk-core).
	ArtifactPurl *string `mandatory:"true" json:"artifactPurl"`

	// Unique identifier of the Artifact Version.
	Id *string `mandatory:"true" json:"id"`

	// Ecosystem which publishes the artifact (e.g. Debian:11, Maven).
	Ecosystem *string `mandatory:"true" json:"ecosystem"`

	// Time at which the artifact version has been published in the ecosystem (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimePublished *common.SDKTime `mandatory:"true" json:"timePublished"`

	// Version of the artifact (e.g. 3.5.2).
	Version *string `mandatory:"false" json:"version"`

	// Version of the ecosystem, such as Fedora, that this artifact belongs to (e.g. 24, main, stable).
	EcosystemVersion *string `mandatory:"false" json:"ecosystemVersion"`

	// Architecture of the platform for which the artifact version is compiled (e.g. ALL, X86_64, AARCH64).
	Architecture *string `mandatory:"false" json:"architecture"`
}

func (m ArtifactVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArtifactVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
