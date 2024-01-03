// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MonitoredResourceAssociationSummary Summary of the monitored resource association.
type MonitoredResourceAssociationSummary struct {

	// Association type between source and destination resources.
	AssociationType *string `mandatory:"true" json:"associationType"`

	// Source Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SourceResourceId *string `mandatory:"true" json:"sourceResourceId"`

	// Destination Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DestinationResourceId *string `mandatory:"true" json:"destinationResourceId"`

	// The association creation time. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	SourceResourceDetails *AssociationResourceDetails `mandatory:"false" json:"sourceResourceDetails"`

	DestinationResourceDetails *AssociationResourceDetails `mandatory:"false" json:"destinationResourceDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MonitoredResourceAssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResourceAssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
