// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryDetailsSummary Summarized Discovery details.
type DiscoveryDetailsSummary struct {

	// Exadata Fleet Update Discovery type.
	Type DiscoveryTypesEnum `mandatory:"true" json:"type"`

	// Exadata service type for the target resource members.
	ServiceType DiscoveryServiceTypesEnum `mandatory:"true" json:"serviceType"`

	// Criteria used for Exadata Fleet Update Discovery.
	Criteria DiscoveryCriteriaEnum `mandatory:"false" json:"criteria,omitempty"`
}

func (m DiscoveryDetailsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryDetailsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDiscoveryTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetDiscoveryServiceTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveryCriteriaEnum(string(m.Criteria)); !ok && m.Criteria != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Criteria: %s. Supported values are: %s.", m.Criteria, strings.Join(GetDiscoveryCriteriaEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
