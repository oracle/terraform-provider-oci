// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnalyticsInstance Analytics Instance metadata.
type AnalyticsInstance struct {

	// The resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of an instance.
	LifecycleState AnalyticsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Analytics feature set.
	FeatureSet FeatureSetEnum `mandatory:"true" json:"featureSet"`

	Capacity *Capacity `mandatory:"true" json:"capacity"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"true" json:"networkEndpointDetails"`

	// The date and time the instance was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Optional description.
	Description *string `mandatory:"false" json:"description"`

	// The license used for the service.
	LicenseType LicenseTypeEnum `mandatory:"false" json:"licenseType,omitempty"`

	// Email address receiving notifications.
	EmailNotification *string `mandatory:"false" json:"emailNotification"`

	// Map of PrivateAccessChannel unique identifier key as KEY and PrivateAccessChannel Object as VALUE.
	PrivateAccessChannels map[string]PrivateAccessChannel `mandatory:"false" json:"privateAccessChannels"`

	// Map of VanityUrl unique identifier key as KEY and VanityUrl Object as VALUE.
	VanityUrlDetails map[string]VanityUrlDetails `mandatory:"false" json:"vanityUrlDetails"`

	// URL of the Analytics service.
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"key": "value"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// OCID of the OCI Vault Key encrypting the customer data stored in this Analytics instance. A null value indicates Oracle managed default encryption.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The date and time the instance was last updated (in the format defined by RFC3339).
	// This timestamp represents updates made through this API. External events do not
	// influence it.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The feature set of an Analytics instance.
	FeatureBundle FeatureBundleEnum `mandatory:"false" json:"featureBundle,omitempty"`

	// Identity domain OCID.
	DomainId *string `mandatory:"false" json:"domainId"`
}

func (m AnalyticsInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyticsInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnalyticsInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAnalyticsInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFeatureSetEnum(string(m.FeatureSet)); !ok && m.FeatureSet != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureSet: %s. Supported values are: %s.", m.FeatureSet, strings.Join(GetFeatureSetEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFeatureBundleEnum(string(m.FeatureBundle)); !ok && m.FeatureBundle != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureBundle: %s. Supported values are: %s.", m.FeatureBundle, strings.Join(GetFeatureBundleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AnalyticsInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                             `json:"description"`
		LicenseType            LicenseTypeEnum                     `json:"licenseType"`
		EmailNotification      *string                             `json:"emailNotification"`
		PrivateAccessChannels  map[string]PrivateAccessChannel     `json:"privateAccessChannels"`
		VanityUrlDetails       map[string]VanityUrlDetails         `json:"vanityUrlDetails"`
		ServiceUrl             *string                             `json:"serviceUrl"`
		DefinedTags            map[string]map[string]interface{}   `json:"definedTags"`
		FreeformTags           map[string]string                   `json:"freeformTags"`
		SystemTags             map[string]map[string]interface{}   `json:"systemTags"`
		KmsKeyId               *string                             `json:"kmsKeyId"`
		TimeUpdated            *common.SDKTime                     `json:"timeUpdated"`
		FeatureBundle          FeatureBundleEnum                   `json:"featureBundle"`
		DomainId               *string                             `json:"domainId"`
		Id                     *string                             `json:"id"`
		Name                   *string                             `json:"name"`
		CompartmentId          *string                             `json:"compartmentId"`
		LifecycleState         AnalyticsInstanceLifecycleStateEnum `json:"lifecycleState"`
		FeatureSet             FeatureSetEnum                      `json:"featureSet"`
		Capacity               *Capacity                           `json:"capacity"`
		NetworkEndpointDetails networkendpointdetails              `json:"networkEndpointDetails"`
		TimeCreated            *common.SDKTime                     `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LicenseType = model.LicenseType

	m.EmailNotification = model.EmailNotification

	m.PrivateAccessChannels = model.PrivateAccessChannels

	m.VanityUrlDetails = model.VanityUrlDetails

	m.ServiceUrl = model.ServiceUrl

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	m.KmsKeyId = model.KmsKeyId

	m.TimeUpdated = model.TimeUpdated

	m.FeatureBundle = model.FeatureBundle

	m.DomainId = model.DomainId

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.FeatureSet = model.FeatureSet

	m.Capacity = model.Capacity

	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(NetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	return
}
