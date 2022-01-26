// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateOggDeploymentDetails Deployment Data for creating an OggDeployment
type CreateOggDeploymentDetails struct {

	// The name given to the GoldenGate service deployment. The name must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter.
	DeploymentName *string `mandatory:"true" json:"deploymentName"`

	// The GoldenGate deployment console username.
	AdminUsername *string `mandatory:"true" json:"adminUsername"`

	// The password associated with the GoldenGate deployment console username. The password must be 8 to 30 characters long and must contain at least 1 uppercase, 1 lowercase, 1 numeric, and 1 special character. Special characters such as ‘$’, ‘^’, or ‘?’ are not allowed.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// A PEM-encoded SSL certificate.
	Certificate *string `mandatory:"false" json:"certificate"`

	// A PEM-encoded private key.
	Key *string `mandatory:"false" json:"key"`
}

func (m CreateOggDeploymentDetails) String() string {
	return common.PointerString(m)
}
