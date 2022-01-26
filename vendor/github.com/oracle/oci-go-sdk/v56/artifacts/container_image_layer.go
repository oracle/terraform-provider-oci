// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ContainerImageLayer The container image layer metadata.
type ContainerImageLayer struct {

	// The sha256 digest of the image layer.
	Digest *string `mandatory:"true" json:"digest"`

	// The size of the layer in bytes.
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// An RFC 3339 timestamp indicating when the layer was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m ContainerImageLayer) String() string {
	return common.PointerString(m)
}
