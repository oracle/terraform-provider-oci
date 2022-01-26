// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// OperatingSystem Operating System of the platform on which the Java Runtime was reported.
type OperatingSystem struct {

	// The operating system type, such as Windows or Linux
	Family OsFamilyEnum `mandatory:"true" json:"family"`

	// The name of the operating system as provided by the Java system property os.name.
	Name *string `mandatory:"true" json:"name"`

	// The version of the operating system as provided by the Java system property os.version.
	Version *string `mandatory:"true" json:"version"`

	// The architecture of the operating system as provided by the Java system property os.arch.
	Architecture *string `mandatory:"true" json:"architecture"`
}

func (m OperatingSystem) String() string {
	return common.PointerString(m)
}
