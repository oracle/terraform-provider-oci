// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateUcVolumeDatabricksEngineArtifactLocationDetails The information to update a Unity Catalog volume location used to share Databricks application artifacts.
type UpdateUcVolumeDatabricksEngineArtifactLocationDetails struct {

	// Controls whether a staged application artifact replaces an existing file or object with
	// the same name at the configured destination. The artifact is typically a GGFAB pipeline
	// package or one of its dependency files.
	// When set to `true`, the existing file or object is replaced with the newly staged version.
	// When set to `false`, staging fails or skips the upload rather than replacing it. This
	// setting does not replace the entire target directory or prefix. It allows an updated
	// pipeline or dependency to be deployed when artifact names remain the same across versions.
	ShouldOverwriteExisting *bool `mandatory:"false" json:"shouldOverwriteExisting"`

	// Unity Catalog volume artifact path.
	Path *string `mandatory:"false" json:"path"`
}

// GetShouldOverwriteExisting returns ShouldOverwriteExisting
func (m UpdateUcVolumeDatabricksEngineArtifactLocationDetails) GetShouldOverwriteExisting() *bool {
	return m.ShouldOverwriteExisting
}

func (m UpdateUcVolumeDatabricksEngineArtifactLocationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateUcVolumeDatabricksEngineArtifactLocationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateUcVolumeDatabricksEngineArtifactLocationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateUcVolumeDatabricksEngineArtifactLocationDetails UpdateUcVolumeDatabricksEngineArtifactLocationDetails
	s := struct {
		DiscriminatorParam string `json:"locationType"`
		MarshalTypeUpdateUcVolumeDatabricksEngineArtifactLocationDetails
	}{
		"UC_VOLUME",
		(MarshalTypeUpdateUcVolumeDatabricksEngineArtifactLocationDetails)(m),
	}

	return json.Marshal(&s)
}
