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

// OriginGroup The representation of OriginGroup
type OriginGroup struct {

	// The list of objects containing origin references and additional properties.
	Origins []OriginGroupOrigins `mandatory:"false" json:"origins"`
}

func (m OriginGroup) String() string {
	return common.PointerString(m)
}
