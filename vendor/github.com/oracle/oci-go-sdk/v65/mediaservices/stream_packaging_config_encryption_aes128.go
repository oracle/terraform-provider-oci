// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamPackagingConfigEncryptionAes128 AES128 encryption type (enabled by default).
type StreamPackagingConfigEncryptionAes128 struct {

	// The identifier of the customer managed Vault KMS symmetric encryption key (null if Oracle managed).
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

func (m StreamPackagingConfigEncryptionAes128) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamPackagingConfigEncryptionAes128) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamPackagingConfigEncryptionAes128) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamPackagingConfigEncryptionAes128 StreamPackagingConfigEncryptionAes128
	s := struct {
		DiscriminatorParam string `json:"algorithm"`
		MarshalTypeStreamPackagingConfigEncryptionAes128
	}{
		"AES128",
		(MarshalTypeStreamPackagingConfigEncryptionAes128)(m),
	}

	return json.Marshal(&s)
}
