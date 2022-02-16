// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Control Plane API
//
// Use the Application Performance Monitoring Control Plane API to perform operations such as creating, updating,
// deleting and listing APM domains and monitoring the progress of these operations using the work request APIs.
//

package apmcontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ApmDomainSummary Summary of an APM domain.
type ApmDomainSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Display name of the APM domain, which can be updated.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment corresponding to the APM domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of the APM domain.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the APM domain.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates if this is an Always Free resource.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The time the APM domain was created, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the APM domain was updated, expressed in RFC 3339 timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ApmDomainSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApmDomainSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
