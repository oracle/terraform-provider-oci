// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Relation Description of Asset Relation.
type Relation struct {

	// The unqiue relation identifier.
	Key *string `mandatory:"true" json:"key"`

	// The asset ocid of the from-asset for relation.
	FromAssetId *string `mandatory:"true" json:"fromAssetId"`

	// The asset ocid of the to-asset for relation.
	ToAssetId *string `mandatory:"true" json:"toAssetId"`

	// The key of the relation from the external environment.
	ExternalRelationKey *string `mandatory:"true" json:"externalRelationKey"`

	// Type of relation.
	RelationType RelationRelationTypeEnum `mandatory:"true" json:"relationType"`

	// The time the relation was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m Relation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Relation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRelationRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetRelationRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RelationRelationTypeEnum Enum with underlying type: string
type RelationRelationTypeEnum string

// Set of constants representing the allowable values for RelationRelationTypeEnum
const (
	RelationRelationTypeAssociation RelationRelationTypeEnum = "ASSOCIATION"
	RelationRelationTypeDependency  RelationRelationTypeEnum = "DEPENDENCY"
	RelationRelationTypeComposition RelationRelationTypeEnum = "COMPOSITION"
)

var mappingRelationRelationTypeEnum = map[string]RelationRelationTypeEnum{
	"ASSOCIATION": RelationRelationTypeAssociation,
	"DEPENDENCY":  RelationRelationTypeDependency,
	"COMPOSITION": RelationRelationTypeComposition,
}

var mappingRelationRelationTypeEnumLowerCase = map[string]RelationRelationTypeEnum{
	"association": RelationRelationTypeAssociation,
	"dependency":  RelationRelationTypeDependency,
	"composition": RelationRelationTypeComposition,
}

// GetRelationRelationTypeEnumValues Enumerates the set of values for RelationRelationTypeEnum
func GetRelationRelationTypeEnumValues() []RelationRelationTypeEnum {
	values := make([]RelationRelationTypeEnum, 0)
	for _, v := range mappingRelationRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRelationRelationTypeEnumStringValues Enumerates the set of values in String for RelationRelationTypeEnum
func GetRelationRelationTypeEnumStringValues() []string {
	return []string{
		"ASSOCIATION",
		"DEPENDENCY",
		"COMPOSITION",
	}
}

// GetMappingRelationRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelationRelationTypeEnum(val string) (RelationRelationTypeEnum, bool) {
	enum, ok := mappingRelationRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
