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

// JavaDownloadRecord A record of Java artifact download in a tenancy.
type JavaDownloadRecord struct {

	// Identifier of the source that downloaded the artifact.
	DownloadSourceId *string `mandatory:"true" json:"downloadSourceId"`

	// Timestamp of download.
	TimeDownloaded *common.SDKTime `mandatory:"true" json:"timeDownloaded"`

	// Type of download.
	DownloadType *string `mandatory:"true" json:"downloadType"`

	// The Java family version identifier.
	FamilyVersion *string `mandatory:"false" json:"familyVersion"`

	// The Java family display name.
	FamilyDisplayName *string `mandatory:"false" json:"familyDisplayName"`

	// The Java release version identifier.
	ReleaseVersion *string `mandatory:"false" json:"releaseVersion"`

	// The target Operating System family for the artifact.
	OsFamily *string `mandatory:"false" json:"osFamily"`

	// The target Operating System architecture for the artifact.
	Architecture *string `mandatory:"false" json:"architecture"`

	// The package type(typically the file extension) of the artifact.
	PackageType *string `mandatory:"false" json:"packageType"`

	// Additional information about the package type.
	PackageTypeDetail *string `mandatory:"false" json:"packageTypeDetail"`
}

func (m JavaDownloadRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaDownloadRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
