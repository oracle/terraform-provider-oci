// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAddonDetails The data to update a KafkaClusterAddon.
type UpdateAddonDetails interface {

	// A unique user-friendly name. Avoid entering confidential information.
	GetDescription() *string
}

type updateaddondetails struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	AddonType   string  `json:"addonType"`
}

// UnmarshalJSON unmarshals json
func (m *updateaddondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateaddondetails updateaddondetails
	s := struct {
		Model Unmarshalerupdateaddondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.AddonType = s.Model.AddonType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateaddondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AddonType {
	case "PUBLICCONNECTIVITY":
		mm := UpdatePublicConnectivityAddonDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateAddonDetails: %s.", m.AddonType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m updateaddondetails) GetDescription() *string {
	return m.Description
}

func (m updateaddondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateaddondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
