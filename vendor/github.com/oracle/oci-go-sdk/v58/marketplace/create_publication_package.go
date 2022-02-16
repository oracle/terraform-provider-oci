// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreatePublicationPackage A base object for creating a publication package.
type CreatePublicationPackage interface {

	// The package version.
	GetPackageVersion() *string

	GetOperatingSystem() *OperatingSystem

	// The end user license agreeement (EULA) that consumers of this listing must accept.
	GetEula() []Eula
}

type createpublicationpackage struct {
	JsonData        []byte
	PackageVersion  *string          `mandatory:"true" json:"packageVersion"`
	OperatingSystem *OperatingSystem `mandatory:"true" json:"operatingSystem"`
	Eula            json.RawMessage  `mandatory:"true" json:"eula"`
	PackageType     string           `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *createpublicationpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatepublicationpackage createpublicationpackage
	s := struct {
		Model Unmarshalercreatepublicationpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PackageVersion = s.Model.PackageVersion
	m.OperatingSystem = s.Model.OperatingSystem
	m.Eula = s.Model.Eula
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createpublicationpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "IMAGE":
		mm := CreateImagePublicationPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetPackageVersion returns PackageVersion
func (m createpublicationpackage) GetPackageVersion() *string {
	return m.PackageVersion
}

//GetOperatingSystem returns OperatingSystem
func (m createpublicationpackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

//GetEula returns Eula
func (m createpublicationpackage) GetEula() json.RawMessage {
	return m.Eula
}

func (m createpublicationpackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createpublicationpackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
