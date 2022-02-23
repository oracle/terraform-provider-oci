// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// ReferenceArtifactSummary Represents Reference details of a data asset.
type ReferenceArtifactSummary struct {

	// The type of the ReferenceInfo.
	ModelType *string `mandatory:"true" json:"modelType"`

	// unique id of service which is referencing dcms artifact.
	ServiceArtifactId *string `mandatory:"true" json:"serviceArtifactId"`

	// Generated key that can be used in API calls to identify referenceinfo.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description of the referenceInfo.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// unique id of dcms artifact that is getting registered.
	DcmsArtifactId *string `mandatory:"false" json:"dcmsArtifactId"`

	// count of how many times a dcms artifact has been registered by a service.
	ReferenceCount *int `mandatory:"false" json:"referenceCount"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m ReferenceArtifactSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferenceArtifactSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
