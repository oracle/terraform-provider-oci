// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAiDataPlatformDetails The data to create a AiDataPlatform.
type CreateAiDataPlatformDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the AiDataPlatform in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The AiDataPlatform type.
	AiDataPlatformType *string `mandatory:"false" json:"aiDataPlatformType"`

	AttachAnalyticsDetails AttachAnalyticsDetails `mandatory:"false" json:"attachAnalyticsDetails"`

	// The flag to enable/disable AiFeatures for the instance.
	IsEnableAiFeature *bool `mandatory:"false" json:"isEnableAiFeature"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vector db Lakehouse 26ai.
	VectorDbId *string `mandatory:"false" json:"vectorDbId"`

	// The Vector DB Lakehouse 26ai ADMIN user password.
	VectorDbAdminCred *string `mandatory:"false" json:"vectorDbAdminCred"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI Vault secret holding the vector db Lakehouse 26ai Admin user password.
	VectorDbAdminSecretId *string `mandatory:"false" json:"vectorDbAdminSecretId"`

	// The [OCID] of the user Master Encryption Key to create resources in user tenancy while provisioning AiDataPlatform.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The [OCID] of the Vault Id to create resources in user tenancy while provisioning AiDataPlatform.
	KmsVaultId *string `mandatory:"false" json:"kmsVaultId"`

	// The name for the default workspace for the AiDataPlatform
	DefaultWorkspaceName *string `mandatory:"false" json:"defaultWorkspaceName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CreateAiDataPlatformDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAiDataPlatformDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateAiDataPlatformDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string                           `json:"displayName"`
		AiDataPlatformType     *string                           `json:"aiDataPlatformType"`
		AttachAnalyticsDetails attachanalyticsdetails            `json:"attachAnalyticsDetails"`
		IsEnableAiFeature      *bool                             `json:"isEnableAiFeature"`
		VectorDbId             *string                           `json:"vectorDbId"`
		VectorDbAdminCred      *string                           `json:"vectorDbAdminCred"`
		VectorDbAdminSecretId  *string                           `json:"vectorDbAdminSecretId"`
		KmsKeyId               *string                           `json:"kmsKeyId"`
		KmsVaultId             *string                           `json:"kmsVaultId"`
		DefaultWorkspaceName   *string                           `json:"defaultWorkspaceName"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		SystemTags             map[string]map[string]interface{} `json:"systemTags"`
		CompartmentId          *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.AiDataPlatformType = model.AiDataPlatformType

	nn, e = model.AttachAnalyticsDetails.UnmarshalPolymorphicJSON(model.AttachAnalyticsDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AttachAnalyticsDetails = nn.(AttachAnalyticsDetails)
	} else {
		m.AttachAnalyticsDetails = nil
	}

	m.IsEnableAiFeature = model.IsEnableAiFeature

	m.VectorDbId = model.VectorDbId

	m.VectorDbAdminCred = model.VectorDbAdminCred

	m.VectorDbAdminSecretId = model.VectorDbAdminSecretId

	m.KmsKeyId = model.KmsKeyId

	m.KmsVaultId = model.KmsVaultId

	m.DefaultWorkspaceName = model.DefaultWorkspaceName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	return
}
