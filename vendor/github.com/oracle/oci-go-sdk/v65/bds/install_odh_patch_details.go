// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstallOdhPatchDetails The reqeust body while installing a ODH patch to a cluster.
type InstallOdhPatchDetails struct {

	// The version of the ODH patch to be installed.
	Version *string `mandatory:"true" json:"version"`

	// The display name of the ODH patch to be installed.
	OdhPatchName *string `mandatory:"true" json:"odhPatchName"`

	// The patch Url of the ODH patch to be installed.
	PaUrl *string `mandatory:"true" json:"paUrl"`

	// The md5Hash of the ODH patch to be installed.
	Md5Hash *string `mandatory:"true" json:"md5Hash"`

	// Cluster Admin Password
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// The flag to check if the ODH patch can be installed.
	IsCompatibilityCheck *bool `mandatory:"false" json:"isCompatibilityCheck"`
}

func (m InstallOdhPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallOdhPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
