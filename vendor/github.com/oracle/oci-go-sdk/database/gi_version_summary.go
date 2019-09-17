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

// GiVersionSummary The Oracle Grid Infrastructure (GI) version.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type GiVersionSummary struct {

	// A valid Oracle Grid Infrastructure (GI) software version.
	Version *string `mandatory:"true" json:"version"`
}

func (m GiVersionSummary) String() string {
	return common.PointerString(m)
}
