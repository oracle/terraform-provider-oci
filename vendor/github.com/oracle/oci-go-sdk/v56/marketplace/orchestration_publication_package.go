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

// OrchestrationPublicationPackage A publication package for stack publications.
type OrchestrationPublicationPackage struct {

	// The ID of the listing that the specified package belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The package version.
	Version *string `mandatory:"true" json:"version"`

	// A description of the package.
	Description *string `mandatory:"false" json:"description"`

	// The unique identifier for the package resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The date and time the publication package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	// A link to the stack resource.
	ResourceLink *string `mandatory:"false" json:"resourceLink"`

	// A list of variables for the stack resource.
	Variables []OrchestrationVariable `mandatory:"false" json:"variables"`
}

//GetDescription returns Description
func (m OrchestrationPublicationPackage) GetDescription() *string {
	return m.Description
}

//GetListingId returns ListingId
func (m OrchestrationPublicationPackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m OrchestrationPublicationPackage) GetVersion() *string {
	return m.Version
}

//GetResourceId returns ResourceId
func (m OrchestrationPublicationPackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m OrchestrationPublicationPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetOperatingSystem returns OperatingSystem
func (m OrchestrationPublicationPackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

func (m OrchestrationPublicationPackage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OrchestrationPublicationPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOrchestrationPublicationPackage OrchestrationPublicationPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeOrchestrationPublicationPackage
	}{
		"ORCHESTRATION",
		(MarshalTypeOrchestrationPublicationPackage)(m),
	}

	return json.Marshal(&s)
}
