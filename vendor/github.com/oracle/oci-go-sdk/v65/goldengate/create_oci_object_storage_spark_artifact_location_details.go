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

// CreateOciObjectStorageSparkArtifactLocationDetails The information about an OCI Object Storage location used to share Spark application artifacts.
type CreateOciObjectStorageSparkArtifactLocationDetails struct {
	Configuration CreateSparkOciObjectStorageConfigDetails `mandatory:"true" json:"configuration"`

	// Controls whether a staged Spark application artifact replaces an existing file or object
	// with the same name at the configured destination. The artifact is typically a GGFAB
	// pipeline package or one of its dependency files.
	// When set to `true`, the existing file or object is replaced with the newly staged version.
	// When set to `false`, staging fails or skips the upload rather than replacing it. This
	// setting does not replace the entire target directory or prefix. It allows an updated
	// pipeline or dependency to be deployed when artifact names remain the same across versions.
	ShouldOverwriteExisting *bool `mandatory:"false" json:"shouldOverwriteExisting"`
}

// GetShouldOverwriteExisting returns ShouldOverwriteExisting
func (m CreateOciObjectStorageSparkArtifactLocationDetails) GetShouldOverwriteExisting() *bool {
	return m.ShouldOverwriteExisting
}

func (m CreateOciObjectStorageSparkArtifactLocationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciObjectStorageSparkArtifactLocationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciObjectStorageSparkArtifactLocationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciObjectStorageSparkArtifactLocationDetails CreateOciObjectStorageSparkArtifactLocationDetails
	s := struct {
		DiscriminatorParam string `json:"locationType"`
		MarshalTypeCreateOciObjectStorageSparkArtifactLocationDetails
	}{
		"OCI_OBJECT_STORAGE",
		(MarshalTypeCreateOciObjectStorageSparkArtifactLocationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOciObjectStorageSparkArtifactLocationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ShouldOverwriteExisting *bool                                    `json:"shouldOverwriteExisting"`
		Configuration           createsparkociobjectstorageconfigdetails `json:"configuration"`
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
		m.Configuration = nn.(CreateSparkOciObjectStorageConfigDetails)
	} else {
		m.Configuration = nil
	}

	return
}
