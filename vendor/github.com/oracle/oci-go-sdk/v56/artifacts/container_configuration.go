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

// ContainerConfiguration Container configuration.
type ContainerConfiguration struct {

	// Whether to create a new container repository when a container is pushed to a new repository path.
	// Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush *bool `mandatory:"true" json:"isRepositoryCreatedOnFirstPush"`

	// The tenancy namespace used in the container repository path.
	Namespace *string `mandatory:"true" json:"namespace"`
}

func (m ContainerConfiguration) String() string {
	return common.PointerString(m)
}
