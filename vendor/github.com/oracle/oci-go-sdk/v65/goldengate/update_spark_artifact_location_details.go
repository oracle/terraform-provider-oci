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

// UpdateSparkArtifactLocationDetails The information to update a location used to share Spark application artifacts with a remote cluster.
type UpdateSparkArtifactLocationDetails interface {

	// Controls whether a staged Spark application artifact replaces an existing file or object
	// with the same name at the configured destination. The artifact is typically a GGFAB
	// pipeline package or one of its dependency files.
	// When set to `true`, the existing file or object is replaced with the newly staged version.
	// When set to `false`, staging fails or skips the upload rather than replacing it. This
	// setting does not replace the entire target directory or prefix. It allows an updated
	// pipeline or dependency to be deployed when artifact names remain the same across versions.
	GetShouldOverwriteExisting() *bool
}

type updatesparkartifactlocationdetails struct {
	JsonData                []byte
	ShouldOverwriteExisting *bool  `mandatory:"false" json:"shouldOverwriteExisting"`
	LocationType            string `json:"locationType"`
}

// UnmarshalJSON unmarshals json
func (m *updatesparkartifactlocationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatesparkartifactlocationdetails updatesparkartifactlocationdetails
	s := struct {
		Model Unmarshalerupdatesparkartifactlocationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ShouldOverwriteExisting = s.Model.ShouldOverwriteExisting
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatesparkartifactlocationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OCI_OBJECT_STORAGE":
		mm := UpdateOciObjectStorageSparkArtifactLocationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateSparkArtifactLocationDetails: %s.", m.LocationType)
		return *m, nil
	}
}

// GetShouldOverwriteExisting returns ShouldOverwriteExisting
func (m updatesparkartifactlocationdetails) GetShouldOverwriteExisting() *bool {
	return m.ShouldOverwriteExisting
}

func (m updatesparkartifactlocationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatesparkartifactlocationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
