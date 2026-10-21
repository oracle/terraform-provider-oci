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

// OciObjectStorageDatabricksEngineArtifactLocationSummary Summary of an OCI Object Storage location used to share Databricks application artifacts.
type OciObjectStorageDatabricksEngineArtifactLocationSummary struct {
	Configuration DatabricksEngineOciObjectStorageConfig `mandatory:"true" json:"configuration"`

	// Controls whether a staged application artifact replaces an existing file or object with
	// the same name at the configured destination. The artifact is typically a GGFAB pipeline
	// package or one of its dependency files.
	// When set to `true`, the existing file or object is replaced with the newly staged version.
	// When set to `false`, staging fails or skips the upload rather than replacing it. This
	// setting does not replace the entire target directory or prefix. It allows an updated
	// pipeline or dependency to be deployed when artifact names remain the same across versions.
	ShouldOverwriteExisting *bool `mandatory:"false" json:"shouldOverwriteExisting"`
}

// GetShouldOverwriteExisting returns ShouldOverwriteExisting
func (m OciObjectStorageDatabricksEngineArtifactLocationSummary) GetShouldOverwriteExisting() *bool {
	return m.ShouldOverwriteExisting
}

func (m OciObjectStorageDatabricksEngineArtifactLocationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciObjectStorageDatabricksEngineArtifactLocationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciObjectStorageDatabricksEngineArtifactLocationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciObjectStorageDatabricksEngineArtifactLocationSummary OciObjectStorageDatabricksEngineArtifactLocationSummary
	s := struct {
		DiscriminatorParam string `json:"locationType"`
		MarshalTypeOciObjectStorageDatabricksEngineArtifactLocationSummary
	}{
		"OCI_OBJECT_STORAGE",
		(MarshalTypeOciObjectStorageDatabricksEngineArtifactLocationSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OciObjectStorageDatabricksEngineArtifactLocationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ShouldOverwriteExisting *bool                                  `json:"shouldOverwriteExisting"`
		Configuration           databricksengineociobjectstorageconfig `json:"configuration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ShouldOverwriteExisting = model.ShouldOverwriteExisting

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(DatabricksEngineOciObjectStorageConfig)
	} else {
		m.Configuration = nil
	}

	return
}
