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

// UpgradeDeploymentCurrentReleaseDetails Definiton of the additional attributes for a Current Release upgrade.
type UpgradeDeploymentCurrentReleaseDetails struct {
}

func (m UpgradeDeploymentCurrentReleaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpgradeDeploymentCurrentReleaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpgradeDeploymentCurrentReleaseDetails UpgradeDeploymentCurrentReleaseDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpgradeDeploymentCurrentReleaseDetails
	}{
		"CURRENT_RELEASE",
		(MarshalTypeUpgradeDeploymentCurrentReleaseDetails)(m),
	}

	return json.Marshal(&s)
}
