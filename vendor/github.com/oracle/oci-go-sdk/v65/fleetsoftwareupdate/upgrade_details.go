// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpgradeDetails Details of supported upgrade options for DB or GI collection.
type UpgradeDetails interface {
}

type upgradedetails struct {
	JsonData       []byte
	CollectionType string `json:"collectionType"`
}

// UnmarshalJSON unmarshals json
func (m *upgradedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupgradedetails upgradedetails
	s := struct {
		Model Unmarshalerupgradedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CollectionType = s.Model.CollectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *upgradedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CollectionType {
	case "DB":
		mm := UpgradeDbCollectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GI":
		mm := UpgradeGiCollectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpgradeDetails: %s.", m.CollectionType)
		return *m, nil
	}
}

func (m upgradedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m upgradedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
