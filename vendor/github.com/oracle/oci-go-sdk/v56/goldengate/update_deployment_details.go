// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateDeploymentDetails The information to use to update a Deployment.
type UpdateDeploymentDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of Network Security Group (https://docs.cloud.oracle.com/Content/Network/Concepts/networksecuritygroups.htm) OCIDs used to define network access for a deployment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	OggData *UpdateOggDeploymentDetails `mandatory:"false" json:"oggData"`
}

func (m UpdateDeploymentDetails) String() string {
	return common.PointerString(m)
}
