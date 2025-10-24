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

// DataflowArtifactsDetail Detail to deploy Artifacts for Dataflow service.
type DataflowArtifactsDetail struct {

	// Instance id of the existing Dataflow Instance for artifact deployment.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Contains the main file (py/jar) along with parameters & configuration to be passed to the DataFlow run.
	Execute *string `mandatory:"true" json:"execute"`

	// Contains the archive from object storage bucket which can be added as dependency to data flow application.
	ArchiveUri *string `mandatory:"false" json:"archiveUri"`
}

func (m DataflowArtifactsDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataflowArtifactsDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
