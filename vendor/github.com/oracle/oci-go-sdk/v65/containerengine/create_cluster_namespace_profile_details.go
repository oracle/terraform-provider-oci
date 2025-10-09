// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateClusterNamespaceProfileDetails The information about new ClusterNamespaceProfile.
type CreateClusterNamespaceProfileDetails struct {

	// Name of the cluster namespace.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of compartment owning the cluster namespace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of the resource. It can be changed after creation.
	Description *string `mandatory:"false" json:"description"`

	// Suffix to append to the end of the namespaces generated from this Profile
	NamespaceSuffix *string `mandatory:"false" json:"namespaceSuffix"`

	ClusterNamespaceSchedulingPolicy ClusterNamespaceSchedulingPolicy `mandatory:"false" json:"clusterNamespaceSchedulingPolicy"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateClusterNamespaceProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateClusterNamespaceProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateClusterNamespaceProfileDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                           `json:"description"`
		NamespaceSuffix                  *string                           `json:"namespaceSuffix"`
		ClusterNamespaceSchedulingPolicy clusternamespaceschedulingpolicy  `json:"clusterNamespaceSchedulingPolicy"`
		FreeformTags                     map[string]string                 `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{} `json:"definedTags"`
		DisplayName                      *string                           `json:"displayName"`
		CompartmentId                    *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.NamespaceSuffix = model.NamespaceSuffix

	nn, e = model.ClusterNamespaceSchedulingPolicy.UnmarshalPolymorphicJSON(model.ClusterNamespaceSchedulingPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ClusterNamespaceSchedulingPolicy = nn.(ClusterNamespaceSchedulingPolicy)
	} else {
		m.ClusterNamespaceSchedulingPolicy = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	return
}
