// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ActivateExadataInfrastructureDetails The activation details for the Exadata infrastructure.
type ActivateExadataInfrastructureDetails struct {

	// The activation zip file.
	ActivationFile []byte `mandatory:"true" json:"activationFile"`
}

func (m ActivateExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}
