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

// GiGoalSoftwareComponentDetails Details of goal version for 'GI' component in an 'EXADB_STACK' type Exadata Fleet Update Collection.
type GiGoalSoftwareComponentDetails struct {
	GoalVersionDetails GiGoalVersionDetails `mandatory:"true" json:"goalVersionDetails"`

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
	HomePolicy GiGoalSoftwareComponentDetailsHomePolicyEnum `mandatory:"false" json:"homePolicy,omitempty"`
}

func (m GiGoalSoftwareComponentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiGoalSoftwareComponentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGiGoalSoftwareComponentDetailsHomePolicyEnum(string(m.HomePolicy)); !ok && m.HomePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HomePolicy: %s. Supported values are: %s.", m.HomePolicy, strings.Join(GetGiGoalSoftwareComponentDetailsHomePolicyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiGoalSoftwareComponentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiGoalSoftwareComponentDetails GiGoalSoftwareComponentDetails
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeGiGoalSoftwareComponentDetails
	}{
		"GI",
		(MarshalTypeGiGoalSoftwareComponentDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GiGoalSoftwareComponentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HomePolicy         GiGoalSoftwareComponentDetailsHomePolicyEnum `json:"homePolicy"`
		NewHomePrefix      *string                                      `json:"newHomePrefix"`
		GoalVersionDetails gigoalversiondetails                         `json:"goalVersionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HomePolicy = model.HomePolicy

	m.NewHomePrefix = model.NewHomePrefix

	nn, e = model.GoalVersionDetails.UnmarshalPolymorphicJSON(model.GoalVersionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GoalVersionDetails = nn.(GiGoalVersionDetails)
	} else {
		m.GoalVersionDetails = nil
	}

	return
}

// GiGoalSoftwareComponentDetailsHomePolicyEnum Enum with underlying type: string
type GiGoalSoftwareComponentDetailsHomePolicyEnum string

// Set of constants representing the allowable values for GiGoalSoftwareComponentDetailsHomePolicyEnum
const (
	GiGoalSoftwareComponentDetailsHomePolicyCreateNew   GiGoalSoftwareComponentDetailsHomePolicyEnum = "CREATE_NEW"
	GiGoalSoftwareComponentDetailsHomePolicyUseExisting GiGoalSoftwareComponentDetailsHomePolicyEnum = "USE_EXISTING"
)

var mappingGiGoalSoftwareComponentDetailsHomePolicyEnum = map[string]GiGoalSoftwareComponentDetailsHomePolicyEnum{
	"CREATE_NEW":   GiGoalSoftwareComponentDetailsHomePolicyCreateNew,
	"USE_EXISTING": GiGoalSoftwareComponentDetailsHomePolicyUseExisting,
}

var mappingGiGoalSoftwareComponentDetailsHomePolicyEnumLowerCase = map[string]GiGoalSoftwareComponentDetailsHomePolicyEnum{
	"create_new":   GiGoalSoftwareComponentDetailsHomePolicyCreateNew,
	"use_existing": GiGoalSoftwareComponentDetailsHomePolicyUseExisting,
}

// GetGiGoalSoftwareComponentDetailsHomePolicyEnumValues Enumerates the set of values for GiGoalSoftwareComponentDetailsHomePolicyEnum
func GetGiGoalSoftwareComponentDetailsHomePolicyEnumValues() []GiGoalSoftwareComponentDetailsHomePolicyEnum {
	values := make([]GiGoalSoftwareComponentDetailsHomePolicyEnum, 0)
	for _, v := range mappingGiGoalSoftwareComponentDetailsHomePolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetGiGoalSoftwareComponentDetailsHomePolicyEnumStringValues Enumerates the set of values in String for GiGoalSoftwareComponentDetailsHomePolicyEnum
func GetGiGoalSoftwareComponentDetailsHomePolicyEnumStringValues() []string {
	return []string{
		"CREATE_NEW",
		"USE_EXISTING",
	}
}

// GetMappingGiGoalSoftwareComponentDetailsHomePolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiGoalSoftwareComponentDetailsHomePolicyEnum(val string) (GiGoalSoftwareComponentDetailsHomePolicyEnum, bool) {
	enum, ok := mappingGiGoalSoftwareComponentDetailsHomePolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
