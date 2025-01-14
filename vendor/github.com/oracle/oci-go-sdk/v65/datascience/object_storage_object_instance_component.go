// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageObjectInstanceComponent Reference to Object Storage Object
type ObjectStorageObjectInstanceComponent struct {

	// Name of instance component
	ComponentName *string `mandatory:"true" json:"componentName"`

	// OCID of Object Storage Object
	ObjectId *string `mandatory:"true" json:"objectId"`

	// Object Storage namespace
	Namespace *string `mandatory:"true" json:"namespace"`

	// Name of object storage bucket
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Name of referenced resource (generally resources do not have to have any name but most resources have name exposed as 'name' or 'displayName' field).
	Name *string `mandatory:"false" json:"name"`
}

// GetName returns Name
func (m ObjectStorageObjectInstanceComponent) GetName() *string {
	return m.Name
}

// GetComponentName returns ComponentName
func (m ObjectStorageObjectInstanceComponent) GetComponentName() *string {
	return m.ComponentName
}

func (m ObjectStorageObjectInstanceComponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageObjectInstanceComponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageObjectInstanceComponent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageObjectInstanceComponent ObjectStorageObjectInstanceComponent
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeObjectStorageObjectInstanceComponent
	}{
		"OBJECT_STORAGE_OBJECT",
		(MarshalTypeObjectStorageObjectInstanceComponent)(m),
	}

	return json.Marshal(&s)
}
