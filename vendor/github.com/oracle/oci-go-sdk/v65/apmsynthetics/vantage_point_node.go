// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VantagePointNode Details of the vantage point node.
type VantagePointNode struct {

	// Name of the vantage point node.
	Name *string `mandatory:"true" json:"name"`

	// ID of the vantage point node.
	Id *string `mandatory:"false" json:"id"`

	// Display name of the vantage point node.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Geographical information of the vantage point node.
	GeoInfo *string `mandatory:"false" json:"geoInfo"`

	// Outgoing links from the vantage point node.
	OutgoingLinks []string `mandatory:"false" json:"outgoingLinks"`
}

func (m VantagePointNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VantagePointNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
