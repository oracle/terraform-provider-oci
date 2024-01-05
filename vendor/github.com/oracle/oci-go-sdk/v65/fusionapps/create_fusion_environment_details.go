// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFusionEnvironmentDetails The configuration details of the FusionEnvironment. For more information about these fields, see Managing Environments (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/manage-environment.htm).
type CreateFusionEnvironmentDetails struct {

	// FusionEnvironment Identifier can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique identifier (OCID) of the compartment where the Fusion Environment is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier (OCID) of the Fusion Environment Family that the Fusion Environment belongs to.
	FusionEnvironmentFamilyId *string `mandatory:"true" json:"fusionEnvironmentFamilyId"`

	// The type of environment. Valid values are Production, Test, or Development.
	FusionEnvironmentType FusionEnvironmentFusionEnvironmentTypeEnum `mandatory:"true" json:"fusionEnvironmentType"`

	CreateFusionEnvironmentAdminUserDetails *CreateFusionEnvironmentAdminUserDetails `mandatory:"true" json:"createFusionEnvironmentAdminUserDetails"`

	MaintenancePolicy *MaintenancePolicy `mandatory:"false" json:"maintenancePolicy"`

	// byok kms keyId
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// DNS prefix.
	DnsPrefix *string `mandatory:"false" json:"dnsPrefix"`

	// Language packs.
	AdditionalLanguagePacks []string `mandatory:"false" json:"additionalLanguagePacks"`

	// Rules.
	Rules []Rule `mandatory:"false" json:"rules"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFusionEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFusionEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFusionEnvironmentFusionEnvironmentTypeEnum(string(m.FusionEnvironmentType)); !ok && m.FusionEnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FusionEnvironmentType: %s. Supported values are: %s.", m.FusionEnvironmentType, strings.Join(GetFusionEnvironmentFusionEnvironmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFusionEnvironmentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MaintenancePolicy                       *MaintenancePolicy                         `json:"maintenancePolicy"`
		KmsKeyId                                *string                                    `json:"kmsKeyId"`
		DnsPrefix                               *string                                    `json:"dnsPrefix"`
		AdditionalLanguagePacks                 []string                                   `json:"additionalLanguagePacks"`
		Rules                                   []rule                                     `json:"rules"`
		FreeformTags                            map[string]string                          `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}          `json:"definedTags"`
		DisplayName                             *string                                    `json:"displayName"`
		CompartmentId                           *string                                    `json:"compartmentId"`
		FusionEnvironmentFamilyId               *string                                    `json:"fusionEnvironmentFamilyId"`
		FusionEnvironmentType                   FusionEnvironmentFusionEnvironmentTypeEnum `json:"fusionEnvironmentType"`
		CreateFusionEnvironmentAdminUserDetails *CreateFusionEnvironmentAdminUserDetails   `json:"createFusionEnvironmentAdminUserDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MaintenancePolicy = model.MaintenancePolicy

	m.KmsKeyId = model.KmsKeyId

	m.DnsPrefix = model.DnsPrefix

	m.AdditionalLanguagePacks = make([]string, len(model.AdditionalLanguagePacks))
	copy(m.AdditionalLanguagePacks, model.AdditionalLanguagePacks)
	m.Rules = make([]Rule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(Rule)
		} else {
			m.Rules[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.FusionEnvironmentFamilyId = model.FusionEnvironmentFamilyId

	m.FusionEnvironmentType = model.FusionEnvironmentType

	m.CreateFusionEnvironmentAdminUserDetails = model.CreateFusionEnvironmentAdminUserDetails

	return
}
