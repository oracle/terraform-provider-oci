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

// FsuGoalVersionDetails Goal version or image details for the Exadata Fleet Update Cycle.
type FsuGoalVersionDetails interface {

	// Goal home policy to use when Staging the Goal Version during patching.
	// CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version.
	// USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.
	//   If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.
	//   If more than one existing home for the selected image is found, then the home with the least number of databases will be used.
	//   If multiple homes have the least number of databases, then a home will be selected at random.
	GetHomePolicy() FsuGoalVersionDetailsHomePolicyEnum

	// Prefix name used for new DB home resources created as part of the Stage Action.
	// Format: <specified_prefix>_<timestamp>
	// If not specified, a default OCI DB home resource will be generated for the new DB home resources created.
	GetNewHomePrefix() *string
}

type fsugoalversiondetails struct {
	JsonData      []byte
	HomePolicy    FsuGoalVersionDetailsHomePolicyEnum `mandatory:"false" json:"homePolicy,omitempty"`
	NewHomePrefix *string                             `mandatory:"false" json:"newHomePrefix"`
	Type          string                              `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsugoalversiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsugoalversiondetails fsugoalversiondetails
	s := struct {
		Model Unmarshalerfsugoalversiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.HomePolicy = s.Model.HomePolicy
	m.NewHomePrefix = s.Model.NewHomePrefix
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsugoalversiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VERSION":
		mm := VersionFsuTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE_ID":
		mm := ImageIdFsuTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuGoalVersionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetHomePolicy returns HomePolicy
func (m fsugoalversiondetails) GetHomePolicy() FsuGoalVersionDetailsHomePolicyEnum {
	return m.HomePolicy
}

// GetNewHomePrefix returns NewHomePrefix
func (m fsugoalversiondetails) GetNewHomePrefix() *string {
	return m.NewHomePrefix
}

func (m fsugoalversiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsugoalversiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFsuGoalVersionDetailsHomePolicyEnum(string(m.HomePolicy)); !ok && m.HomePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HomePolicy: %s. Supported values are: %s.", m.HomePolicy, strings.Join(GetFsuGoalVersionDetailsHomePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FsuGoalVersionDetailsHomePolicyEnum Enum with underlying type: string
type FsuGoalVersionDetailsHomePolicyEnum string

// Set of constants representing the allowable values for FsuGoalVersionDetailsHomePolicyEnum
const (
	FsuGoalVersionDetailsHomePolicyCreateNew   FsuGoalVersionDetailsHomePolicyEnum = "CREATE_NEW"
	FsuGoalVersionDetailsHomePolicyUseExisting FsuGoalVersionDetailsHomePolicyEnum = "USE_EXISTING"
)

var mappingFsuGoalVersionDetailsHomePolicyEnum = map[string]FsuGoalVersionDetailsHomePolicyEnum{
	"CREATE_NEW":   FsuGoalVersionDetailsHomePolicyCreateNew,
	"USE_EXISTING": FsuGoalVersionDetailsHomePolicyUseExisting,
}

var mappingFsuGoalVersionDetailsHomePolicyEnumLowerCase = map[string]FsuGoalVersionDetailsHomePolicyEnum{
	"create_new":   FsuGoalVersionDetailsHomePolicyCreateNew,
	"use_existing": FsuGoalVersionDetailsHomePolicyUseExisting,
}

// GetFsuGoalVersionDetailsHomePolicyEnumValues Enumerates the set of values for FsuGoalVersionDetailsHomePolicyEnum
func GetFsuGoalVersionDetailsHomePolicyEnumValues() []FsuGoalVersionDetailsHomePolicyEnum {
	values := make([]FsuGoalVersionDetailsHomePolicyEnum, 0)
	for _, v := range mappingFsuGoalVersionDetailsHomePolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuGoalVersionDetailsHomePolicyEnumStringValues Enumerates the set of values in String for FsuGoalVersionDetailsHomePolicyEnum
func GetFsuGoalVersionDetailsHomePolicyEnumStringValues() []string {
	return []string{
		"CREATE_NEW",
		"USE_EXISTING",
	}
}

// GetMappingFsuGoalVersionDetailsHomePolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuGoalVersionDetailsHomePolicyEnum(val string) (FsuGoalVersionDetailsHomePolicyEnum, bool) {
	enum, ok := mappingFsuGoalVersionDetailsHomePolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FsuGoalVersionDetailsTypeEnum Enum with underlying type: string
type FsuGoalVersionDetailsTypeEnum string

// Set of constants representing the allowable values for FsuGoalVersionDetailsTypeEnum
const (
	FsuGoalVersionDetailsTypeVersion FsuGoalVersionDetailsTypeEnum = "VERSION"
	FsuGoalVersionDetailsTypeImageId FsuGoalVersionDetailsTypeEnum = "IMAGE_ID"
)

var mappingFsuGoalVersionDetailsTypeEnum = map[string]FsuGoalVersionDetailsTypeEnum{
	"VERSION":  FsuGoalVersionDetailsTypeVersion,
	"IMAGE_ID": FsuGoalVersionDetailsTypeImageId,
}

var mappingFsuGoalVersionDetailsTypeEnumLowerCase = map[string]FsuGoalVersionDetailsTypeEnum{
	"version":  FsuGoalVersionDetailsTypeVersion,
	"image_id": FsuGoalVersionDetailsTypeImageId,
}

// GetFsuGoalVersionDetailsTypeEnumValues Enumerates the set of values for FsuGoalVersionDetailsTypeEnum
func GetFsuGoalVersionDetailsTypeEnumValues() []FsuGoalVersionDetailsTypeEnum {
	values := make([]FsuGoalVersionDetailsTypeEnum, 0)
	for _, v := range mappingFsuGoalVersionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuGoalVersionDetailsTypeEnumStringValues Enumerates the set of values in String for FsuGoalVersionDetailsTypeEnum
func GetFsuGoalVersionDetailsTypeEnumStringValues() []string {
	return []string{
		"VERSION",
		"IMAGE_ID",
	}
}

// GetMappingFsuGoalVersionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuGoalVersionDetailsTypeEnum(val string) (FsuGoalVersionDetailsTypeEnum, bool) {
	enum, ok := mappingFsuGoalVersionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
