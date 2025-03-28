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

// RemoveFsuCollectionTargetsDetails Remove targets from a Exadata Fleet Update Collection.
type RemoveFsuCollectionTargetsDetails interface {
}

type removefsucollectiontargetsdetails struct {
	JsonData        []byte
	RemovalStrategy string `json:"removalStrategy"`
}

// UnmarshalJSON unmarshals json
func (m *removefsucollectiontargetsdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerremovefsucollectiontargetsdetails removefsucollectiontargetsdetails
	s := struct {
		Model Unmarshalerremovefsucollectiontargetsdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RemovalStrategy = s.Model.RemovalStrategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *removefsucollectiontargetsdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RemovalStrategy {
	case "TARGET_IDS":
		mm := TargetIdsRemoveTargetsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RemoveFsuCollectionTargetsDetails: %s.", m.RemovalStrategy)
		return *m, nil
	}
}

func (m removefsucollectiontargetsdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m removefsucollectiontargetsdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum Enum with underlying type: string
type RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum string

// Set of constants representing the allowable values for RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum
const (
	RemoveFsuCollectionTargetsDetailsRemovalStrategyTargetIds RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum = "TARGET_IDS"
)

var mappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnum = map[string]RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum{
	"TARGET_IDS": RemoveFsuCollectionTargetsDetailsRemovalStrategyTargetIds,
}

var mappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumLowerCase = map[string]RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum{
	"target_ids": RemoveFsuCollectionTargetsDetailsRemovalStrategyTargetIds,
}

// GetRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumValues Enumerates the set of values for RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum
func GetRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumValues() []RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum {
	values := make([]RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum, 0)
	for _, v := range mappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumStringValues Enumerates the set of values in String for RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum
func GetRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumStringValues() []string {
	return []string{
		"TARGET_IDS",
	}
}

// GetMappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnum(val string) (RemoveFsuCollectionTargetsDetailsRemovalStrategyEnum, bool) {
	enum, ok := mappingRemoveFsuCollectionTargetsDetailsRemovalStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
