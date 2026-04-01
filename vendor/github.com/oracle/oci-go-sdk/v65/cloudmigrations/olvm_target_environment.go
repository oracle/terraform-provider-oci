// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmTargetEnvironment OLVM target enviroment
type OlvmTargetEnvironment struct {

	// Inventory asset id of the olvm cluster
	ClusterAssetId *string `mandatory:"true" json:"clusterAssetId"`

	// Inventory asset Id of the vnic profile
	VnicProfileAssetId *string `mandatory:"true" json:"vnicProfileAssetId"`

	// Target compartment identifier
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// OLVM OS type to inventory asset id of the template
	OlvmTemplates map[string]string `mandatory:"false" json:"olvmTemplates"`

	// Preferred VM shape type provided by the customer.
	PreferredShapeType VmTargetAssetPreferredShapeTypeEnum `mandatory:"false" json:"preferredShapeType,omitempty"`
}

// GetTargetCompartmentId returns TargetCompartmentId
func (m OlvmTargetEnvironment) GetTargetCompartmentId() *string {
	return m.TargetCompartmentId
}

func (m OlvmTargetEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmTargetEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmTargetAssetPreferredShapeTypeEnum(string(m.PreferredShapeType)); !ok && m.PreferredShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredShapeType: %s. Supported values are: %s.", m.PreferredShapeType, strings.Join(GetVmTargetAssetPreferredShapeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OlvmTargetEnvironment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOlvmTargetEnvironment OlvmTargetEnvironment
	s := struct {
		DiscriminatorParam string `json:"targetEnvironmentType"`
		MarshalTypeOlvmTargetEnvironment
	}{
		"OLVM_TARGET_ENV",
		(MarshalTypeOlvmTargetEnvironment)(m),
	}

	return json.Marshal(&s)
}
