// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaMigrationAnalysisTarget The target for the Java Migration Analysis
type JavaMigrationAnalysisTarget struct {

	// Name of the analysis project.
	AnalysisProjectName *string `mandatory:"true" json:"analysisProjectName"`

	// Object storage paths to the input files applications to be analysed.
	InputApplicationsObjectStoragePaths []string `mandatory:"true" json:"inputApplicationsObjectStoragePaths"`

	// Version of the target JDKs.
	TargetJdkVersions []string `mandatory:"true" json:"targetJdkVersions"`

	// Package prefixes to be included from the migration analysis. Either this or excludePackagePrefixes can be specified.
	IncludePackagePrefixes []string `mandatory:"false" json:"includePackagePrefixes"`

	// Package prefixes to be excluded from the migration analysis. Either this or includePackagePrefixes can be specified.
	ExcludePackagePrefixes []string `mandatory:"false" json:"excludePackagePrefixes"`
}

func (m JavaMigrationAnalysisTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaMigrationAnalysisTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
