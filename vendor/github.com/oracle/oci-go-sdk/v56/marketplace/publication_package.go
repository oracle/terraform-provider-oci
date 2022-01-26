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

// PublicationPackage A base object for all types of publication packages.
type PublicationPackage interface {

	// The ID of the listing that the specified package belongs to.
	GetListingId() *string

	// The package version.
	GetVersion() *string

	// A description of the package.
	GetDescription() *string

	// The unique identifier for the package resource.
	GetResourceId() *string

	// The date and time the publication package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	GetOperatingSystem() *OperatingSystem
}

type publicationpackage struct {
	JsonData        []byte
	ListingId       *string          `mandatory:"true" json:"listingId"`
	Version         *string          `mandatory:"true" json:"version"`
	Description     *string          `mandatory:"false" json:"description"`
	ResourceId      *string          `mandatory:"false" json:"resourceId"`
	TimeCreated     *common.SDKTime  `mandatory:"false" json:"timeCreated"`
	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`
	PackageType     string           `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *publicationpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpublicationpackage publicationpackage
	s := struct {
		Model Unmarshalerpublicationpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ListingId = s.Model.ListingId
	m.Version = s.Model.Version
	m.Description = s.Model.Description
	m.ResourceId = s.Model.ResourceId
	m.TimeCreated = s.Model.TimeCreated
	m.OperatingSystem = s.Model.OperatingSystem
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *publicationpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "ORCHESTRATION":
		mm := OrchestrationPublicationPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE":
		mm := ImagePublicationPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetListingId returns ListingId
func (m publicationpackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m publicationpackage) GetVersion() *string {
	return m.Version
}

//GetDescription returns Description
func (m publicationpackage) GetDescription() *string {
	return m.Description
}

//GetResourceId returns ResourceId
func (m publicationpackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m publicationpackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetOperatingSystem returns OperatingSystem
func (m publicationpackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

func (m publicationpackage) String() string {
	return common.PointerString(m)
}
