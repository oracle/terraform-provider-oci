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

// ExadbStackFsuGoalVersionDetails Details of goal version for an 'EXADB_STACK' type Exadata Fleet Update Collection.
// Currently, components allowed in an Exadata software stack are 'GUEST_OS' and 'GI'.
// At least two distinct component types are required for an Exadata software stack.
type ExadbStackFsuGoalVersionDetails struct {

	// Details of goal versions for components in an Exadata software stack.
	Components []GoalSoftwareComponentDetails `mandatory:"true" json:"components"`

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
func (m ExadbStackFsuGoalVersionDetails) GetHomePolicy() FsuGoalVersionDetailsHomePolicyEnum {
	return m.HomePolicy
}

// GetNewHomePrefix returns NewHomePrefix
func (m ExadbStackFsuGoalVersionDetails) GetNewHomePrefix() *string {
	return m.NewHomePrefix
}

func (m ExadbStackFsuGoalVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbStackFsuGoalVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFsuGoalVersionDetailsHomePolicyEnum(string(m.HomePolicy)); !ok && m.HomePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HomePolicy: %s. Supported values are: %s.", m.HomePolicy, strings.Join(GetFsuGoalVersionDetailsHomePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExadbStackFsuGoalVersionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadbStackFsuGoalVersionDetails ExadbStackFsuGoalVersionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExadbStackFsuGoalVersionDetails
	}{
		"EXADB_STACK",
		(MarshalTypeExadbStackFsuGoalVersionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExadbStackFsuGoalVersionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HomePolicy    FsuGoalVersionDetailsHomePolicyEnum `json:"homePolicy"`
		NewHomePrefix *string                             `json:"newHomePrefix"`
		Components    []goalsoftwarecomponentdetails      `json:"components"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HomePolicy = model.HomePolicy

	m.NewHomePrefix = model.NewHomePrefix

	m.Components = make([]GoalSoftwareComponentDetails, len(model.Components))
	for i, n := range model.Components {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Components[i] = nn.(GoalSoftwareComponentDetails)
		} else {
			m.Components[i] = nil
		}
	}
	return
}
