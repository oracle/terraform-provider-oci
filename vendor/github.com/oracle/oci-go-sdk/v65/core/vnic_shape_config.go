// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VnicShapeConfig Shape config of VNIC that will be used to allocate resource in the data plane once the VNIC is attached
type VnicShapeConfig struct {

	// It defines the percentage number of concurrent connections that can be tracked to the VNIC.
	PercentageOfConnTrack *int `mandatory:"false" json:"percentageOfConnTrack"`

	// It defines the bandwidthMbps for the shape.
	AggregateBandwidthBps *int64 `mandatory:"false" json:"aggregateBandwidthBps"`

	// VCNCP will use this flag to set the internet bandwidth for always free vnic.
	IsAlwaysFree *bool `mandatory:"false" json:"isAlwaysFree"`

	// An enum for different platforms (E3, A1, X5, etc).
	VnicShapeCategory VnicShapeCategoryEnum `mandatory:"false" json:"vnicShapeCategory,omitempty"`
}

func (m VnicShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VnicShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVnicShapeCategoryEnum(string(m.VnicShapeCategory)); !ok && m.VnicShapeCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShapeCategory: %s. Supported values are: %s.", m.VnicShapeCategory, strings.Join(GetVnicShapeCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
