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

// AuditArtifactVersionSummary Summary of an Artifact Version, returned as a result of a List operation.
type AuditArtifactVersionSummary struct {

	// The complete Package URL (purl) with the version
	Purl *string `mandatory:"true" json:"purl"`

	// Time at which the artifact version has been published in the ecosystem (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimePublished *common.SDKTime `mandatory:"true" json:"timePublished"`
}

func (m AuditArtifactVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditArtifactVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
