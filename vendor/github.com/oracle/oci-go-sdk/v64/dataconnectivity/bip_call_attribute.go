// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// BipCallAttribute The call attributes impl
type BipCallAttribute struct {
	StagingBucket *Schema `mandatory:"false" json:"stagingBucket"`

	// Parameter to set offset
	OffsetParameter *string `mandatory:"false" json:"offsetParameter"`

	// Parameter to fetch next set of rows
	FetchNextRowsParameter *string `mandatory:"false" json:"fetchNextRowsParameter"`

	StagingDataAsset *DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection *Connection `mandatory:"false" json:"stagingConnection"`

	// Prefix for the staging DataAsset
	StagingPrefix *string `mandatory:"false" json:"stagingPrefix"`
}

func (m BipCallAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BipCallAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BipCallAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBipCallAttribute BipCallAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeBipCallAttribute
	}{
		"BIPCALLATTRIBUTE",
		(MarshalTypeBipCallAttribute)(m),
	}

	return json.Marshal(&s)
}
