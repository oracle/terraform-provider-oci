// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportInventoryDetails Details for importing assets from a file.
type ImportInventoryDetails interface {

	// The OCID of the compartmentId that resources import.
	GetCompartmentId() *string

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type importinventorydetails struct {
	JsonData      []byte
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	ResourceType  string                            `json:"resourceType"`
}

// UnmarshalJSON unmarshals json
func (m *importinventorydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimportinventorydetails importinventorydetails
	s := struct {
		Model Unmarshalerimportinventorydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ResourceType = s.Model.ResourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *importinventorydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResourceType {
	case "ASSET":
		mm := ImportInventoryViaAssetsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ImportInventoryDetails: %s.", m.ResourceType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m importinventorydetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m importinventorydetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m importinventorydetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m importinventorydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m importinventorydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportInventoryDetailsResourceTypeEnum Enum with underlying type: string
type ImportInventoryDetailsResourceTypeEnum string

// Set of constants representing the allowable values for ImportInventoryDetailsResourceTypeEnum
const (
	ImportInventoryDetailsResourceTypeAsset ImportInventoryDetailsResourceTypeEnum = "ASSET"
)

var mappingImportInventoryDetailsResourceTypeEnum = map[string]ImportInventoryDetailsResourceTypeEnum{
	"ASSET": ImportInventoryDetailsResourceTypeAsset,
}

var mappingImportInventoryDetailsResourceTypeEnumLowerCase = map[string]ImportInventoryDetailsResourceTypeEnum{
	"asset": ImportInventoryDetailsResourceTypeAsset,
}

// GetImportInventoryDetailsResourceTypeEnumValues Enumerates the set of values for ImportInventoryDetailsResourceTypeEnum
func GetImportInventoryDetailsResourceTypeEnumValues() []ImportInventoryDetailsResourceTypeEnum {
	values := make([]ImportInventoryDetailsResourceTypeEnum, 0)
	for _, v := range mappingImportInventoryDetailsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetImportInventoryDetailsResourceTypeEnumStringValues Enumerates the set of values in String for ImportInventoryDetailsResourceTypeEnum
func GetImportInventoryDetailsResourceTypeEnumStringValues() []string {
	return []string{
		"ASSET",
	}
}

// GetMappingImportInventoryDetailsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportInventoryDetailsResourceTypeEnum(val string) (ImportInventoryDetailsResourceTypeEnum, bool) {
	enum, ok := mappingImportInventoryDetailsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
