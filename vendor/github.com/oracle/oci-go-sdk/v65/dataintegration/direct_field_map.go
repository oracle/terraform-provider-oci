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

// DirectFieldMap The information about a field map.
type DirectFieldMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Deprecated - Reference to a typed object.
	SourceTypedObject *string `mandatory:"false" json:"sourceTypedObject"`

	// Deprecated - Reference to a typed object.
	TargetTypedObject *string `mandatory:"false" json:"targetTypedObject"`

	SourceScopeReference *ScopeReference `mandatory:"false" json:"sourceScopeReference"`

	TargetScopeReference *ScopeReference `mandatory:"false" json:"targetScopeReference"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

// GetDescription returns Description
func (m DirectFieldMap) GetDescription() *string {
	return m.Description
}

func (m DirectFieldMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DirectFieldMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DirectFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectFieldMap DirectFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDirectFieldMap
	}{
		"DIRECT_FIELD_MAP",
		(MarshalTypeDirectFieldMap)(m),
	}

	return json.Marshal(&s)
}
