// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationDependencyRecommendationSummary An application dependency with the recommended version that does not contain any CVE.
// Each application dependency has a property specifying multiple node identifiers on which which this current node depends.
type ApplicationDependencyRecommendationSummary struct {

	// Unique Group Artifact Version (GAV) identifier in the format _Group:Artifact:Version_, e.g. org.graalvm.nativeimage:svm:21.1.0.
	Gav *string `mandatory:"true" json:"gav"`

	// Unique node identifier of an application dependency with an associated Recommendation, e.g. nodeId1.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// List of (application dependencies) node identifiers from which this node depends.
	ApplicationDependencyNodeIds []string `mandatory:"true" json:"applicationDependencyNodeIds"`

	// Package URL defined in https://github.com/package-url/purl-spec, e.g. pkg:maven/org.graalvm.nativeimage/svm@21.1.0
	Purl *string `mandatory:"false" json:"purl"`

	// Recommended application dependency in "group:artifact:version" (GAV) format, e.g. org.graalvm.nativeimage:svm:21.2.0.
	RecommendedGav *string `mandatory:"false" json:"recommendedGav"`

	// Recommended application dependency in PURL format, e.g. pkg:maven/org.graalvm.nativeimage/svm@21.2.0
	RecommendedPurl *string `mandatory:"false" json:"recommendedPurl"`
}

func (m ApplicationDependencyRecommendationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationDependencyRecommendationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
