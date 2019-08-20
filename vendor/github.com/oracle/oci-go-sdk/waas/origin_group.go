// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
