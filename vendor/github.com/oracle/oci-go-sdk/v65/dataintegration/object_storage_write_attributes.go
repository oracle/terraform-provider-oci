// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageWriteAttributes Properties to configure writing to Object Storage.
type ObjectStorageWriteAttributes struct {

	// Specifies whether to write output to single-file or not.
	WriteToSingleFile *bool `mandatory:"false" json:"writeToSingleFile"`
}

func (m ObjectStorageWriteAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageWriteAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageWriteAttributes) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageWriteAttributes ObjectStorageWriteAttributes
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeObjectStorageWriteAttributes
	}{
		"OBJECT_STORAGE_WRITE_ATTRIBUTE",
		(MarshalTypeObjectStorageWriteAttributes)(m),
	}

	return json.Marshal(&s)
}
