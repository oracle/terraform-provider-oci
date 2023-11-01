// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricExtensionSummary Summary information about metric extension resources
type MetricExtensionSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of metric extension.
	Id *string `mandatory:"true" json:"id"`

	// Metric Extension Resource name.
	Name *string `mandatory:"true" json:"name"`

	// Resource type to which Metric Extension applies
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the metric extension.
	Status MetricExtensionLifeCycleDetailsEnum `mandatory:"true" json:"status"`

	// Metric Extension resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the metric extension.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the metric extension
	LifecycleState MetricExtensionLifeCycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Metric Extension creation time. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Metric Extension updation time. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type of possible collection methods.
	CollectionMethod MetricExtensionCollectionMethodsEnum `mandatory:"false" json:"collectionMethod,omitempty"`

	// Count of resources on which this metric extension is enabled.
	EnabledOnResourcesCount *int `mandatory:"false" json:"enabledOnResourcesCount"`

	// The URI path that the user can do a GET on to access the metric extension metadata
	ResourceUri *string `mandatory:"false" json:"resourceUri"`
}

func (m MetricExtensionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricExtensionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricExtensionLifeCycleDetailsEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMetricExtensionLifeCycleDetailsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMetricExtensionLifeCycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMetricExtensionLifeCycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMetricExtensionCollectionMethodsEnum(string(m.CollectionMethod)); !ok && m.CollectionMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionMethod: %s. Supported values are: %s.", m.CollectionMethod, strings.Join(GetMetricExtensionCollectionMethodsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
