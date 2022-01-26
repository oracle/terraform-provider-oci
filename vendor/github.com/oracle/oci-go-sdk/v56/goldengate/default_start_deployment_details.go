// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DefaultStartDeploymentDetails Definiton of the additional attributes for default deployment start.
type DefaultStartDeploymentDetails struct {
}

func (m DefaultStartDeploymentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DefaultStartDeploymentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultStartDeploymentDetails DefaultStartDeploymentDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefaultStartDeploymentDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultStartDeploymentDetails)(m),
	}

	return json.Marshal(&s)
}
