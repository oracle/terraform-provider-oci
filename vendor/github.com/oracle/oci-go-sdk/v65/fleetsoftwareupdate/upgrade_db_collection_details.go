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

// UpgradeDbCollectionDetails Details of supported upgrade options for DB collection.
type UpgradeDbCollectionDetails struct {

	// Enables or disables time zone upgrade.
	IsTimeZoneUpgrade *bool `mandatory:"false" json:"isTimeZoneUpgrade"`

	// Enables or disables the recompilation of invalid objects.
	IsRecompileInvalidObjects *bool `mandatory:"false" json:"isRecompileInvalidObjects"`

	// Service drain timeout specified in seconds.
	MaxDrainTimeoutInSeconds *int `mandatory:"false" json:"maxDrainTimeoutInSeconds"`
}

func (m UpgradeDbCollectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpgradeDbCollectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpgradeDbCollectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpgradeDbCollectionDetails UpgradeDbCollectionDetails
	s := struct {
		DiscriminatorParam string `json:"collectionType"`
		MarshalTypeUpgradeDbCollectionDetails
	}{
		"DB",
		(MarshalTypeUpgradeDbCollectionDetails)(m),
	}

	return json.Marshal(&s)
}
