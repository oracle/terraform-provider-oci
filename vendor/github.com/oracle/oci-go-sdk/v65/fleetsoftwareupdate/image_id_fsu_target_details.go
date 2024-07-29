// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ImageIdFsuTargetDetails The Database service Database Software Image resource is required as goal version for the Exadata Fleet Update Maintenance Cycle
// when IMAGE_ID type is selected.
// Specify a Database Software Image of type DATABASE_IMAGE for a DB Collection.
// Specify a Database Software Image of type GRID_IMAGE for a GI Collection.
type ImageIdFsuTargetDetails struct {

	// Target database software image OCID.
	SoftwareImageId *string `mandatory:"true" json:"softwareImageId"`

	// Prefix name used for new DB home resources created as part of the Stage Action.
	// Format: <specified_prefix>_<timestamp>
	// If not specified, a default OCI DB home resource will be generated for the new DB home resources created.
	NewHomePrefix *string `mandatory:"false" json:"newHomePrefix"`

	// Goal home policy to use when Staging the Goal Version during patching.
	// CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version.
	// USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.
	//   If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.
	//   If more than one existing home for the selected image is found, then the home with the least number of databases will be used.
	//   If multiple homes have the least number of databases, then a home will be selected at random.
	HomePolicy FsuGoalVersionDetailsHomePolicyEnum `mandatory:"false" json:"homePolicy,omitempty"`
}

// GetHomePolicy returns HomePolicy
func (m ImageIdFsuTargetDetails) GetHomePolicy() FsuGoalVersionDetailsHomePolicyEnum {
	return m.HomePolicy
}

// GetNewHomePrefix returns NewHomePrefix
func (m ImageIdFsuTargetDetails) GetNewHomePrefix() *string {
	return m.NewHomePrefix
}

func (m ImageIdFsuTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageIdFsuTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFsuGoalVersionDetailsHomePolicyEnum(string(m.HomePolicy)); !ok && m.HomePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HomePolicy: %s. Supported values are: %s.", m.HomePolicy, strings.Join(GetFsuGoalVersionDetailsHomePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImageIdFsuTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageIdFsuTargetDetails ImageIdFsuTargetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeImageIdFsuTargetDetails
	}{
		"IMAGE_ID",
		(MarshalTypeImageIdFsuTargetDetails)(m),
	}

	return json.Marshal(&s)
}
