// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccmDemandSignalResourcePropertiesSummary A summary model for the Occm demand signal resource properties.
type OccmDemandSignalResourcePropertiesSummary struct {

	// The name of demand signal resource's property.
	PropertyName *string `mandatory:"true" json:"propertyName"`

	// Default value of demand signal resource's property.
	PropertyValue *string `mandatory:"true" json:"propertyValue"`

	// This will indicate if demand signal resource's property is editable.
	IsEditable *bool `mandatory:"true" json:"isEditable"`

	// The minimum value of demand signal resource's property. This is an optional parameter.
	PropertyMinValue *int64 `mandatory:"false" json:"propertyMinValue"`

	// The maximum value of demand signal resource's property. This is an optional parameter.
	PropertyMaxValue *int64 `mandatory:"false" json:"propertyMaxValue"`

	// Predefined options for demand signal resource's property. This is an optional parameter.
	PropertyOptions []OccmDemandSignalResourcePropertyOptionSummary `mandatory:"false" json:"propertyOptions"`

	// Unit for demand signal resource's property.
	PropertyUnit *string `mandatory:"false" json:"propertyUnit"`
}

func (m OccmDemandSignalResourcePropertiesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalResourcePropertiesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
