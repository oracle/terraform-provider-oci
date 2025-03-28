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

// OutputLink Details about the outgoing data of an operator in a data flow design.
type OutputLink struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// Key of FlowPort reference
	Port *string `mandatory:"false" json:"port"`

	// The links from this output link to connect to other links in flow.
	ToLinks []string `mandatory:"false" json:"toLinks"`
}

// GetKey returns Key
func (m OutputLink) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m OutputLink) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m OutputLink) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetObjectStatus returns ObjectStatus
func (m OutputLink) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m OutputLink) GetDescription() *string {
	return m.Description
}

// GetPort returns Port
func (m OutputLink) GetPort() *string {
	return m.Port
}

func (m OutputLink) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OutputLink) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OutputLink) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOutputLink OutputLink
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOutputLink
	}{
		"OUTPUT_LINK",
		(MarshalTypeOutputLink)(m),
	}

	return json.Marshal(&s)
}
