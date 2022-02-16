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

// UpdateGoldenGateHub Details about Oracle GoldenGate Microservices.
type UpdateGoldenGateHub struct {
	RestAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"restAdminCredentials"`

	SourceDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"sourceDbAdminCredentials"`

	SourceContainerDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"sourceContainerDbAdminCredentials"`

	TargetDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"targetDbAdminCredentials"`

	// Oracle GoldenGate hub's REST endpoint.
	// Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA
	Url *string `mandatory:"false" json:"url"`

	// Name of GoldenGate deployment to operate on source database
	SourceMicroservicesDeploymentName *string `mandatory:"false" json:"sourceMicroservicesDeploymentName"`

	// Name of GoldenGate deployment to operate on target database
	TargetMicroservicesDeploymentName *string `mandatory:"false" json:"targetMicroservicesDeploymentName"`

	// OCID of GoldenGate compute instance. An empty value will remove the stored computeId.
	ComputeId *string `mandatory:"false" json:"computeId"`
}

func (m UpdateGoldenGateHub) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGoldenGateHub) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
