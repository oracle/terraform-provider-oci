// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// GoldenGateHub Details about Oracle GoldenGate Microservices.
type GoldenGateHub struct {
	RestAdminCredentials *AdminCredentials `mandatory:"true" json:"restAdminCredentials"`

	SourceDbAdminCredentials *AdminCredentials `mandatory:"true" json:"sourceDbAdminCredentials"`

	TargetDbAdminCredentials *AdminCredentials `mandatory:"true" json:"targetDbAdminCredentials"`

	// Oracle GoldenGate hub's REST endpoint.
	// Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA
	Url *string `mandatory:"true" json:"url"`

	// Name of GoldenGate deployment to operate on source database
	SourceMicroservicesDeploymentName *string `mandatory:"true" json:"sourceMicroservicesDeploymentName"`

	// Name of GoldenGate deployment to operate on target database
	TargetMicroservicesDeploymentName *string `mandatory:"true" json:"targetMicroservicesDeploymentName"`

	SourceContainerDbAdminCredentials *AdminCredentials `mandatory:"false" json:"sourceContainerDbAdminCredentials"`

	// OCID of GoldenGate compute instance.
	ComputeId *string `mandatory:"false" json:"computeId"`
}

func (m GoldenGateHub) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GoldenGateHub) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
