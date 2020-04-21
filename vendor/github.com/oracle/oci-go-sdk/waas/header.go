// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Header An HTTP header with name and value.
type Header struct {

	// The name of the header.
	Name *string `mandatory:"true" json:"name"`

	// The value of the header.
	Value *string `mandatory:"true" json:"value"`
}

func (m Header) String() string {
	return common.PointerString(m)
}
