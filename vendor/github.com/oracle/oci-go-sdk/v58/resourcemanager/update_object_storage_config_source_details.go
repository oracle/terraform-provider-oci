// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateObjectStorageConfigSourceDetails Updates property details for the Object Storage bucket that contains Terraform configuration files.
type UpdateObjectStorageConfigSourceDetails struct {

	// The path of the directory from which to run terraform. If not specified, the the root will be used. This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The name of the bucket's region.
	// Example: `PHX`
	Region *string `mandatory:"false" json:"region"`

	// The Object Storage namespace that contains the bucket.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The name of the bucket that contains the Terraform configuration files.
	BucketName *string `mandatory:"false" json:"bucketName"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m UpdateObjectStorageConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m UpdateObjectStorageConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateObjectStorageConfigSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateObjectStorageConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateObjectStorageConfigSourceDetails UpdateObjectStorageConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeUpdateObjectStorageConfigSourceDetails
	}{
		"OBJECT_STORAGE_CONFIG_SOURCE",
		(MarshalTypeUpdateObjectStorageConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
