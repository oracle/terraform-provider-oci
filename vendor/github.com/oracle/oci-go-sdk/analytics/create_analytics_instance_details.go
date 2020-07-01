// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAnalyticsInstanceDetails Input payload to create an Anaytics instance.
type CreateAnalyticsInstanceDetails struct {

	// The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Analytics feature set.
	FeatureSet FeatureSetEnum `mandatory:"true" json:"featureSet"`

	Capacity *Capacity `mandatory:"true" json:"capacity"`

	// The license used for the service.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// Optional description.
	Description *string `mandatory:"false" json:"description"`

	// Email address receiving notifications.
	EmailNotification *string `mandatory:"false" json:"emailNotification"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`

	// IDCS access token identifying a stripe and service administrator user.
	IdcsAccessToken *string `mandatory:"false" json:"idcsAccessToken"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateAnalyticsInstanceDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAnalyticsInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                           `json:"description"`
		EmailNotification      *string                           `json:"emailNotification"`
		NetworkEndpointDetails networkendpointdetails            `json:"networkEndpointDetails"`
		IdcsAccessToken        *string                           `json:"idcsAccessToken"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		Name                   *string                           `json:"name"`
		CompartmentId          *string                           `json:"compartmentId"`
		FeatureSet             FeatureSetEnum                    `json:"featureSet"`
		Capacity               *Capacity                         `json:"capacity"`
		LicenseType            LicenseTypeEnum                   `json:"licenseType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.EmailNotification = model.EmailNotification

	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(NetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	m.IdcsAccessToken = model.IdcsAccessToken

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.FeatureSet = model.FeatureSet

	m.Capacity = model.Capacity

	m.LicenseType = model.LicenseType

	return
}
