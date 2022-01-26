// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreatePrivateApplicationPackage A base object for creating a private application package.
type CreatePrivateApplicationPackage interface {

	// The package version.
	GetVersion() *string
}

type createprivateapplicationpackage struct {
	JsonData    []byte
	Version     *string `mandatory:"true" json:"version"`
	PackageType string  `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *createprivateapplicationpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateprivateapplicationpackage createprivateapplicationpackage
	s := struct {
		Model Unmarshalercreateprivateapplicationpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Version = s.Model.Version
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createprivateapplicationpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "STACK":
		mm := CreatePrivateApplicationStackPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetVersion returns Version
func (m createprivateapplicationpackage) GetVersion() *string {
	return m.Version
}

func (m createprivateapplicationpackage) String() string {
	return common.PointerString(m)
}
