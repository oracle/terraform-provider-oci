// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GgcsArtifactsDetail Details required to deploy artifacts in the GGCS deployment.
type GgcsArtifactsDetail struct {

	// Instance id of the exisitng GGCS instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Object storage root path containing GGCS artifacts.
	ArtifactObjectStoragePath *string `mandatory:"false" json:"artifactObjectStoragePath"`

	// Ggcs user details to be created or updated.
	Users []GgcsUserDetail `mandatory:"false" json:"users"`

	// Source Detail to configure existing or new datasource.
	Sources []GgcsSourceDetail `mandatory:"false" json:"sources"`

	// Target Detail to configure existing or new datasource.
	Targets []GgcsTargetDetail `mandatory:"false" json:"targets"`
}

func (m GgcsArtifactsDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GgcsArtifactsDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
