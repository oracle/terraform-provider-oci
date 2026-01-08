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

// ArtifactVersionNoticeFile Content of a notice file and the file path to where it is located in the repository, associated to an artifact version.
type ArtifactVersionNoticeFile struct {

	// Unique identifier of an artifact version, for example nodeId1.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// Relative path to the root of the source code repository of the file where the notice text was identified.
	FilePath *string `mandatory:"true" json:"filePath"`

	// Full content of the notice file.
	Content *string `mandatory:"true" json:"content"`
}

func (m ArtifactVersionNoticeFile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArtifactVersionNoticeFile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
