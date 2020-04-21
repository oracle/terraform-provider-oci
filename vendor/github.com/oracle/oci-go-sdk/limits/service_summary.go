// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service limits APIs
//
// APIs that interact with the resource limits of a specific resource type
//

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceSummary A specific OCI service supported by resource limits.
type ServiceSummary struct {

	// The service name. Use this when calling the other APIs.
	Name *string `mandatory:"false" json:"name"`

	// The friendly service name.
	Description *string `mandatory:"false" json:"description"`
}

func (m ServiceSummary) String() string {
	return common.PointerString(m)
}
