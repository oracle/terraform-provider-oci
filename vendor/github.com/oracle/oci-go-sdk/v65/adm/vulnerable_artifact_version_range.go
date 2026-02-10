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

// VulnerableArtifactVersionRange Vulnerable artifact version range.
type VulnerableArtifactVersionRange struct {

	// The version immediately after the last affected version. Versions up to, but not including this version, are vulnerable.
	VersionEndExcluding *string `mandatory:"false" json:"versionEndExcluding"`

	// Marks the latest version that is affected by the vulnerability. This version and all preceding versions, going back to versionStartExcluding or versionStartIncluding, are considered vulnerable.
	VersionEndIncluding *string `mandatory:"false" json:"versionEndIncluding"`

	// The first version affected by the vulnerability. This version and those following it are considered vulnerable until versionEndExcluding or versionEndIncluding is reached.
	VersionStartIncluding *string `mandatory:"false" json:"versionStartIncluding"`

	// The version immediately before the start of affected versions. The specified version is not affected, but versions immediately after are, up to versionStartIncluding or beyond, if not otherwise defined.
	VersionStartExcluding *string `mandatory:"false" json:"versionStartExcluding"`
}

func (m VulnerableArtifactVersionRange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VulnerableArtifactVersionRange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
