// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ArtifactVersionAuditResultSummary Audit results (license identifiers) for a given artifact version.
// Each artifact version is uniquely defined by a nodeId.
type ArtifactVersionAuditResultSummary struct {

	// Unique identifier of an artifact version, for example nodeId1.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// Package URL identifier, e.g. pkg:maven/org.graalvm.nativeimage/svm@21.1.0
	Purl *string `mandatory:"false" json:"purl"`

	// List of artifact version ids on which this artifact version depends, each identified by its nodeId.
	ApplicationDependencyNodeIds []string `mandatory:"false" json:"applicationDependencyNodeIds"`

	// Time at which the artifact version has been published in the ecosystem (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimePublished *common.SDKTime `mandatory:"false" json:"timePublished"`

	// List of licenses for the artifact version.
	Licenses []License `mandatory:"false" json:"licenses"`

	// Url of the repository where the source code is hosted.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// Url to the source code location of the artifactVersion.
	SourceCodeUrl *string `mandatory:"false" json:"sourceCodeUrl"`

	// Url to the primary public website of the project.
	ProjectUrl *string `mandatory:"false" json:"projectUrl"`

	// Potential issues in the artifact identifiers (purls) provided by the user, that the user is alerted about. ADM supports the following warnings:
	// - MISSING_VERSION: Missing version
	// - INCORRECT_FORMAT_VERSION: Version can not be parsed according to the ecosystem
	// - UNKNOWN_ECOSYSTEM: The purl's ecosystem is unknown to ADM
	// - INCORRECT_FORMAT_EPOCH: Epoch qualifier can not be parsed according to the ecosystem
	// - MISSING_DISTRO: The distro qualifier is required for the ecosystem, but was not provided in the purl
	// - UNKNOWN_DISTRO: ADM does not have data for the provided distro value for the given ecosystem
	// - MISSING_ARCH: The arch qualifier is required for the ecosystem and distro, but was not provided in the purl
	// - UNKNOWN_ARCH: ADM does not have data for the provided arch value for the given ecosystem and distro
	// - EPOCH_VALUE_MISMATCH: Epoch value in qualifier is not the same as the epoch value in version
	// - INVALID_PURL: The provided PURL could not be parsed
	Warnings []string `mandatory:"false" json:"warnings"`
}

func (m ArtifactVersionAuditResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArtifactVersionAuditResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
