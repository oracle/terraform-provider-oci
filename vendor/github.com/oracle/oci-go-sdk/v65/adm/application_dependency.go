// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ApplicationDependency An Application Dependency resource creates a Vulnerability Audit.
type ApplicationDependency struct {

	// Unique Group Artifact Version (GAV) identifier (Group:Artifact:Version), e.g. org.graalvm.nativeimage:svm:21.1.0.
	Gav *string `mandatory:"true" json:"gav"`

	// Unique identifier of an Application Dependency node, e.g. nodeId1.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// List of (Application Dependencies) node identifiers on which this node depends.
	ApplicationDependencyNodeIds []string `mandatory:"true" json:"applicationDependencyNodeIds"`
}

func (m ApplicationDependency) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationDependency) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
