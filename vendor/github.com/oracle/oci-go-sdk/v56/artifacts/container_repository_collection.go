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

// ContainerRepositoryCollection List of container repository results.
type ContainerRepositoryCollection struct {

	// Total number of layers.
	LayerCount *int `mandatory:"true" json:"layerCount"`

	// Total storage in bytes consumed by layers.
	LayersSizeInBytes *int64 `mandatory:"true" json:"layersSizeInBytes"`

	// Total number of images.
	ImageCount *int `mandatory:"true" json:"imageCount"`

	// Collection of container repositories.
	Items []ContainerRepositorySummary `mandatory:"true" json:"items"`

	// Estimated number of remaining results.
	RemainingItemsCount *int `mandatory:"true" json:"remainingItemsCount"`

	// Total number of repositories.
	RepositoryCount *int `mandatory:"true" json:"repositoryCount"`
}

func (m ContainerRepositoryCollection) String() string {
	return common.PointerString(m)
}
