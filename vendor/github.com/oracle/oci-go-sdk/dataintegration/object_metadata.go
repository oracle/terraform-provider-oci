// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ObjectMetadata A summary type containing information about the object including its key, name and when/who created/updated it
type ObjectMetadata struct {

	// The user that created the object.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The user that created the object.
	CreatedByName *string `mandatory:"false" json:"createdByName"`

	// The user that updated the object.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// The user that updated the object.
	UpdatedByName *string `mandatory:"false" json:"updatedByName"`

	// The date and time that the object was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time that the object was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The owning object key for this object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// The full path to identify this object.
	IdentifierPath *string `mandatory:"false" json:"identifierPath"`

	// infoFields
	InfoFields map[string]string `mandatory:"false" json:"infoFields"`

	// registryVersion
	RegistryVersion *int `mandatory:"false" json:"registryVersion"`

	// Labels are keywords or tags that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`
}

func (m ObjectMetadata) String() string {
	return common.PointerString(m)
}
