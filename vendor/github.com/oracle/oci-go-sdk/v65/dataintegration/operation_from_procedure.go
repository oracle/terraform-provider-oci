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

// OperationFromProcedure The operation object.
type OperationFromProcedure struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The model version of the object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The operation name.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

// GetMetadata returns Metadata
func (m OperationFromProcedure) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m OperationFromProcedure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperationFromProcedure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OperationFromProcedure) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOperationFromProcedure OperationFromProcedure
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOperationFromProcedure
	}{
		"PROCEDURE",
		(MarshalTypeOperationFromProcedure)(m),
	}

	return json.Marshal(&s)
}
