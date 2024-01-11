// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageDataset **ObjectStorageDataset**
// The dataset residing in OCI Object Storage.
type ObjectStorageDataset struct {

	// The Object Storage namespace.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The Object Storage bucket name.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage object name.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m ObjectStorageDataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageDataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageDataset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageDataset ObjectStorageDataset
	s := struct {
		DiscriminatorParam string `json:"datasetType"`
		MarshalTypeObjectStorageDataset
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageDataset)(m),
	}

	return json.Marshal(&s)
}
