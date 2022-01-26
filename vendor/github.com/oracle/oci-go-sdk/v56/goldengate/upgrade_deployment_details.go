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

// UpgradeDeploymentDetails The information about the Upgrade for a Deployment.
type UpgradeDeploymentDetails interface {
}

type upgradedeploymentdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *upgradedeploymentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupgradedeploymentdetails upgradedeploymentdetails
	s := struct {
		Model Unmarshalerupgradedeploymentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *upgradedeploymentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CURRENT_RELEASE":
		mm := UpgradeDeploymentCurrentReleaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m upgradedeploymentdetails) String() string {
	return common.PointerString(m)
}
