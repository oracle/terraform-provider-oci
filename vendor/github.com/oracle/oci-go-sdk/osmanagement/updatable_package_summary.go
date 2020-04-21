// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdatablePackageSummary A software package available for update on a managed instance
type UpdatablePackageSummary struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique name for the package.
	Name *string `mandatory:"true" json:"name"`

	// Type of the update.
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package
	InstalledVersion *string `mandatory:"true" json:"installedVersion"`

	// Version of the package available for update
	AvailableVersion *string `mandatory:"true" json:"availableVersion"`

	// Unique name for the package available for update.
	AvailablePackageName *string `mandatory:"false" json:"availablePackageName"`

	// The architecture for which this package was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m UpdatablePackageSummary) String() string {
	return common.PointerString(m)
}
