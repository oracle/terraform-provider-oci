// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlOutputObjectStorageDetails Object Storage resource output details.
type ExecuteSqlOutputObjectStorageDetails struct {
	Object *ExecuteSqlObjectStorageLocation `mandatory:"true" json:"object"`

	// Defines how the result of commands in a script should be stored.
	// If the command does not match any template filter, the result will be inline.
	ResultDispositionTemplates []ExecuteSqlOutputResultDispositionTemplate `mandatory:"false" json:"resultDispositionTemplates"`

	// The time when the object becomes eligible for deletion, expressed as an RFC 3339 date-time string.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

// GetResultDispositionTemplates returns ResultDispositionTemplates
func (m ExecuteSqlOutputObjectStorageDetails) GetResultDispositionTemplates() []ExecuteSqlOutputResultDispositionTemplate {
	return m.ResultDispositionTemplates
}

func (m ExecuteSqlOutputObjectStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlOutputObjectStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlOutputObjectStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlOutputObjectStorageDetails ExecuteSqlOutputObjectStorageDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlOutputObjectStorageDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeExecuteSqlOutputObjectStorageDetails)(m),
	}

	return json.Marshal(&s)
}
