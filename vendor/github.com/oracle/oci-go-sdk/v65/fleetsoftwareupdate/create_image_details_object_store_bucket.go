// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateImageDetailsObjectStoreBucket OCI Object Storage bucket that will contain the exa_map file to register the Exadata Fleet Update Image.
type CreateImageDetailsObjectStoreBucket struct {

	// Namespace name of the object store bucket.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Bucket name.
	BucketName *string `mandatory:"true" json:"bucketName"`
}

func (m CreateImageDetailsObjectStoreBucket) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateImageDetailsObjectStoreBucket) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateImageDetailsObjectStoreBucket) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateImageDetailsObjectStoreBucket CreateImageDetailsObjectStoreBucket
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCreateImageDetailsObjectStoreBucket
	}{
		"OBJECT_STORAGE_BUCKET",
		(MarshalTypeCreateImageDetailsObjectStoreBucket)(m),
	}

	return json.Marshal(&s)
}
