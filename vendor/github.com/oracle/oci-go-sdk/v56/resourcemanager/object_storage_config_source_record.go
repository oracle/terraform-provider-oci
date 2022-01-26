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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ObjectStorageConfigSourceRecord Metadata about the Object Storage configuration source.
type ObjectStorageConfigSourceRecord struct {

	// The name of the bucket's region.
	// Example: `PHX`
	Region *string `mandatory:"true" json:"region"`

	// The Object Storage namespace that contains the bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket that contains the Terraform configuration files.
	BucketName *string `mandatory:"true" json:"bucketName"`
}

func (m ObjectStorageConfigSourceRecord) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageConfigSourceRecord) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageConfigSourceRecord ObjectStorageConfigSourceRecord
	s := struct {
		DiscriminatorParam string `json:"configSourceRecordType"`
		MarshalTypeObjectStorageConfigSourceRecord
	}{
		"OBJECT_STORAGE_CONFIG_SOURCE",
		(MarshalTypeObjectStorageConfigSourceRecord)(m),
	}

	return json.Marshal(&s)
}
