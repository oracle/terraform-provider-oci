// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProblemEndpointSummary Summary information for endpoints associated with a problem (Problem object).
type ProblemEndpointSummary struct {

	// Unique identifier for problem endpoint.
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier for sighting associated with the endpoint
	SightingId *string `mandatory:"true" json:"sightingId"`

	// Unique identifier for problem associated with the endpoint
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Unique identifier for the sighting type associated with the endpoint
	SightingType *string `mandatory:"true" json:"sightingType"`

	// Display name of the sighting type
	SightingTypeDisplayName *string `mandatory:"true" json:"sightingTypeDisplayName"`

	// IP address of the endpoint
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Type of IP address for the endpoint
	IpAddressType *string `mandatory:"true" json:"ipAddressType"`

	// Time when activities were last detected
	TimeLastDetected *common.SDKTime `mandatory:"true" json:"timeLastDetected"`

	// IP address classification type for the endpoint
	IpClassificationType *string `mandatory:"false" json:"ipClassificationType"`

	// Country of the endpoint
	Country *string `mandatory:"false" json:"country"`

	// Latitude of the endpoint
	Latitude *float64 `mandatory:"false" json:"latitude"`

	// Longitude of the endpoint
	Longitude *float64 `mandatory:"false" json:"longitude"`

	// ASN number of the endpoint
	AsnNumber *string `mandatory:"false" json:"asnNumber"`

	// Regions where activities were performed from this IP address
	Regions []string `mandatory:"false" json:"regions"`

	// List of services where activities were performed from this IP address
	Services []string `mandatory:"false" json:"services"`
}

func (m ProblemEndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProblemEndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
