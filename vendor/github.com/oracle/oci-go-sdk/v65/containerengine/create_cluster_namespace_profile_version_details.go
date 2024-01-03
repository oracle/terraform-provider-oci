// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateClusterNamespaceProfileVersionDetails The information about a new version of a Cluster Namespace Profile.
type CreateClusterNamespaceProfileVersionDetails struct {

	// A name for the Cluster Namespace Profile Version. Names are unique across versions in a Cluster Namespace Profile Profiles.
	Name *string `mandatory:"true" json:"name"`

	// OCID of compartment owning the Cluster Namespace Profile Version.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the ClusterNamespaceProfile
	ClusterNamespaceProfileId *string `mandatory:"true" json:"clusterNamespaceProfileId"`

	// Name of the ClusterRole to bind to the admin account in the resulting namespace.
	AdminClusterRoleName *string `mandatory:"true" json:"adminClusterRoleName"`

	// Description of the resource. It can be changed after creation.
	Description *string `mandatory:"false" json:"description"`

	// List of Kubernetes labels to apply to the resulting namespace.
	FixedNamespaceLabels []NamespaceLabel `mandatory:"false" json:"fixedNamespaceLabels"`

	// List of Kubernetes annotations to apply to the resulting namespace.
	FixedNamespaceAnnotations []NamespaceAnnotation `mandatory:"false" json:"fixedNamespaceAnnotations"`

	// List of Kubernetes labels that may be specified via Cluster Namespaces.
	AllowedNamespaceLabels []AllowedNamespaceLabel `mandatory:"false" json:"allowedNamespaceLabels"`

	// List of Kubernetes annotations that may be specified via Cluster Namespaces.
	AllowedNamespaceAnnotations []AllowedNamespaceAnnotation `mandatory:"false" json:"allowedNamespaceAnnotations"`

	// List of Kubernetes labels that must be specified via Cluster Namespaces.
	RequiredNamespaceLabels []RequiredNamespaceLabel `mandatory:"false" json:"requiredNamespaceLabels"`

	// List of Kubernetes annotations that must be specified via Cluster Namespaces.
	RequiredNamespaceAnnotations []RequiredNamespaceAnnotation `mandatory:"false" json:"requiredNamespaceAnnotations"`

	// If set to true, the Cluster Namespace Profile Version is not consumable by new Cluster Namespace configurations.
	IsDeprecated *bool `mandatory:"false" json:"isDeprecated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateClusterNamespaceProfileVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateClusterNamespaceProfileVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
