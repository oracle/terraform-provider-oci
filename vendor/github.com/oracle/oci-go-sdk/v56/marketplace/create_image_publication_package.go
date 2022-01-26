// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateImagePublicationPackage An object for creating an image publication package.
type CreateImagePublicationPackage struct {

	// The package version.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	OperatingSystem *OperatingSystem `mandatory:"true" json:"operatingSystem"`

	// The end user license agreeement (EULA) that consumers of this listing must accept.
	Eula []Eula `mandatory:"true" json:"eula"`

	// The unique identifier for the base image of the publication.
	ImageId *string `mandatory:"false" json:"imageId"`
}

//GetPackageVersion returns PackageVersion
func (m CreateImagePublicationPackage) GetPackageVersion() *string {
	return m.PackageVersion
}

//GetOperatingSystem returns OperatingSystem
func (m CreateImagePublicationPackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

//GetEula returns Eula
func (m CreateImagePublicationPackage) GetEula() []Eula {
	return m.Eula
}

func (m CreateImagePublicationPackage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateImagePublicationPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateImagePublicationPackage CreateImagePublicationPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeCreateImagePublicationPackage
	}{
		"IMAGE",
		(MarshalTypeCreateImagePublicationPackage)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateImagePublicationPackage) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ImageId         *string          `json:"imageId"`
		PackageVersion  *string          `json:"packageVersion"`
		OperatingSystem *OperatingSystem `json:"operatingSystem"`
		Eula            []eula           `json:"eula"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ImageId = model.ImageId

	m.PackageVersion = model.PackageVersion

	m.OperatingSystem = model.OperatingSystem

	m.Eula = make([]Eula, len(model.Eula))
	for i, n := range model.Eula {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Eula[i] = nn.(Eula)
		} else {
			m.Eula[i] = nil
		}
	}

	return
}
