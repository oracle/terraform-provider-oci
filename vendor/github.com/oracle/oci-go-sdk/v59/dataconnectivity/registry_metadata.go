// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// RegistryMetadata Information about the object and its parent.
type RegistryMetadata struct {

	// The owning object's key for this object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	// The registry version.
	RegistryVersion *int `mandatory:"false" json:"registryVersion"`

	// The identifying key for the object.
	Key *string `mandatory:"false" json:"key"`

	// Specifies whether this object is a favorite or not.
	IsFavorite *bool `mandatory:"false" json:"isFavorite"`

	// The id of the user who created the object.
	CreatedByUserId *string `mandatory:"false" json:"createdByUserId"`

	// The name of the user who created the object.
	CreatedByUserName *string `mandatory:"false" json:"createdByUserName"`

	// The id of the user who updated the object.
	UpdatedByUserId *string `mandatory:"false" json:"updatedByUserId"`

	// The name of the user who updated the object.
	UpdatedByUserName *string `mandatory:"false" json:"updatedByUserName"`

	// The date and time that the object was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time that the object was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m RegistryMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegistryMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
