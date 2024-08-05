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

// JavaArtifact Information about a binary artifact of Java.
type JavaArtifact struct {

	// Unique identifier for the artifact.
	ArtifactId *int64 `mandatory:"true" json:"artifactId"`

	// Description of the binary artifact. Typically includes the OS, architecture, and installer type.
	ArtifactDescription *string `mandatory:"true" json:"artifactDescription"`

	// Product content type of this artifact.
	ArtifactContentType ArtifactContentTypeEnum `mandatory:"true" json:"artifactContentType"`

	// Approximate compressed file size in bytes.
	ApproximateFileSizeInBytes *int64 `mandatory:"true" json:"approximateFileSizeInBytes"`

	// SHA256 checksum of the artifact.
	Sha256 *string `mandatory:"true" json:"sha256"`

	// The target Operating System family for the artifact.
	OsFamily *string `mandatory:"true" json:"osFamily"`

	// The target Operating System architecture for the artifact.
	Architecture *string `mandatory:"true" json:"architecture"`

	// The package type(typically the file extension) of the artifact.
	PackageType *string `mandatory:"true" json:"packageType"`

	// The endpoint that returns a short-lived artifact download URL in the response payload.
	// This download url can then be used for downloading the artifact.
	// See this API (https://docs.oracle.com/en-us/iaas/api/#/en/jms-java-download/20230601/DownloadUrl/GenerateArtifactDownloadUrl) for more details.
	DownloadUrl *string `mandatory:"true" json:"downloadUrl"`

	// The endpoint for downloading this artifact from command line, automatically in scripts and dockerfiles.
	// Depending on the context, this can point to the archive or latest update release version artifact in the specified family.
	ScriptDownloadUrl *string `mandatory:"true" json:"scriptDownloadUrl"`

	// The URL for retrieving the checksum for the artifact.
	// Depending on the context, this can point to the checksum of the archive or latest update release version artifact.
	ScriptChecksumUrl *string `mandatory:"true" json:"scriptChecksumUrl"`

	// The file name of the artifact.
	ArtifactFileName *string `mandatory:"false" json:"artifactFileName"`

	// Additional information about the package type.
	PackageTypeDetail *string `mandatory:"false" json:"packageTypeDetail"`
}

func (m JavaArtifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaArtifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactContentTypeEnum(string(m.ArtifactContentType)); !ok && m.ArtifactContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArtifactContentType: %s. Supported values are: %s.", m.ArtifactContentType, strings.Join(GetArtifactContentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
