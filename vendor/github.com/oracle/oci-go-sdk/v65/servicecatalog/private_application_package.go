// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateApplicationPackage A base object for all types of private application packages.
type PrivateApplicationPackage interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application package.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted.
	GetPrivateApplicationId() *string

	// The package version.
	GetVersion() *string

	// The date and time the private application package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-27T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The display name of the package.
	GetDisplayName() *string
}

type privateapplicationpackage struct {
	JsonData             []byte
	DisplayName          *string         `mandatory:"false" json:"displayName"`
	Id                   *string         `mandatory:"true" json:"id"`
	PrivateApplicationId *string         `mandatory:"true" json:"privateApplicationId"`
	Version              *string         `mandatory:"true" json:"version"`
	TimeCreated          *common.SDKTime `mandatory:"true" json:"timeCreated"`
	PackageType          string          `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *privateapplicationpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprivateapplicationpackage privateapplicationpackage
	s := struct {
		Model Unmarshalerprivateapplicationpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.PrivateApplicationId = s.Model.PrivateApplicationId
	m.Version = s.Model.Version
	m.TimeCreated = s.Model.TimeCreated
	m.DisplayName = s.Model.DisplayName
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *privateapplicationpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "STACK":
		mm := PrivateApplicationStackPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PrivateApplicationPackage: %s.", m.PackageType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m privateapplicationpackage) GetDisplayName() *string {
	return m.DisplayName
}

// GetId returns Id
func (m privateapplicationpackage) GetId() *string {
	return m.Id
}

// GetPrivateApplicationId returns PrivateApplicationId
func (m privateapplicationpackage) GetPrivateApplicationId() *string {
	return m.PrivateApplicationId
}

// GetVersion returns Version
func (m privateapplicationpackage) GetVersion() *string {
	return m.Version
}

// GetTimeCreated returns TimeCreated
func (m privateapplicationpackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m privateapplicationpackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m privateapplicationpackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
