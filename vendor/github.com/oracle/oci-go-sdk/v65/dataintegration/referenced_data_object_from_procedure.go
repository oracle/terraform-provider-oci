// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ReferencedDataObjectFromProcedure The input procedure object.
type ReferencedDataObjectFromProcedure struct {

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`
}

// GetModelVersion returns ModelVersion
func (m ReferencedDataObjectFromProcedure) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ReferencedDataObjectFromProcedure) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m ReferencedDataObjectFromProcedure) GetName() *string {
	return m.Name
}

// GetObjectVersion returns ObjectVersion
func (m ReferencedDataObjectFromProcedure) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetResourceName returns ResourceName
func (m ReferencedDataObjectFromProcedure) GetResourceName() *string {
	return m.ResourceName
}

// GetObjectStatus returns ObjectStatus
func (m ReferencedDataObjectFromProcedure) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetExternalKey returns ExternalKey
func (m ReferencedDataObjectFromProcedure) GetExternalKey() *string {
	return m.ExternalKey
}

func (m ReferencedDataObjectFromProcedure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferencedDataObjectFromProcedure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ReferencedDataObjectFromProcedure) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReferencedDataObjectFromProcedure ReferencedDataObjectFromProcedure
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeReferencedDataObjectFromProcedure
	}{
		"PROCEDURE",
		(MarshalTypeReferencedDataObjectFromProcedure)(m),
	}

	return json.Marshal(&s)
}
