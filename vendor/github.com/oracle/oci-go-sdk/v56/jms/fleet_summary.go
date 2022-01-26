// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FleetSummary The summary of the Fleet.
// A Fleet is the primary collection with which users interact when using Java Management Service.
type FleetSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Fleet. The displayName must be unique for Fleets in the same compartment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Fleet's description.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The approximate count of all unique Java Runtimes in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateJreCount *int `mandatory:"true" json:"approximateJreCount"`

	// The approximate count of all unique Java Installations in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateInstallationCount *int `mandatory:"true" json:"approximateInstallationCount"`

	// The approximate count of all unique applications in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateApplicationCount *int `mandatory:"true" json:"approximateApplicationCount"`

	// The approximate count of all unique managed instances in the Fleet in the past seven days.
	// This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateManagedInstanceCount *int `mandatory:"true" json:"approximateManagedInstanceCount"`

	// The creation date and time of the Fleet (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the Fleet.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FleetSummary) String() string {
	return common.PointerString(m)
}
