// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OmkArtifactsDetail Detail to deploy artifacts for OMK service.
type OmkArtifactsDetail struct {

	// Instance id of the existing OMK instance for artifact deployment.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// List of kubernetes secrets to create or update in the namespace-name of target cluster-namespace. Each entry source secret values from OCI vault.
	Secrets []SecretDetail `mandatory:"false" json:"secrets"`

	// Object storage path for the deployment manifest.
	ManifestObjectStoragePath *string `mandatory:"false" json:"manifestObjectStoragePath"`

	// Component overrides for stack specific parameters applied during artifact template rendering.
	ComponentValueOverrides []ComponentValueOverride `mandatory:"false" json:"componentValueOverrides"`
}

func (m OmkArtifactsDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OmkArtifactsDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
