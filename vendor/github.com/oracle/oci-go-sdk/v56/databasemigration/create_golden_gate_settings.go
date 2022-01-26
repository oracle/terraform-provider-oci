// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateGoldenGateSettings Optional settings for GoldenGate Microservices processes
type CreateGoldenGateSettings struct {
	Extract *CreateExtract `mandatory:"false" json:"extract"`

	Replicat *CreateReplicat `mandatory:"false" json:"replicat"`

	// ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
	AcceptableLag *int `mandatory:"false" json:"acceptableLag"`
}

func (m CreateGoldenGateSettings) String() string {
	return common.PointerString(m)
}
