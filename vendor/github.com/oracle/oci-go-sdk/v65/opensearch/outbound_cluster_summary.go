// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OutboundClusterSummary Contains details of a Outbound cluster
type OutboundClusterSummary struct {

	// Name of the Outbound cluster. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the Outbound cluster
	SeedClusterId *string `mandatory:"true" json:"seedClusterId"`

	// Sets the time interval between regular application-level ping messages that are sent to try and keep outbound cluster connections alive. If set to -1, application-level ping messages to this outbound cluster are not sent. If unset, application-level ping messages are sent according to the global transport.ping_schedule setting, which defaults to -1 meaning that pings are not sent.
	PingSchedule *string `mandatory:"false" json:"pingSchedule"`

	// Flag to indicate whether to skip the Outbound cluster during cross cluster search, if it is unavailable
	IsSkipUnavailable *bool `mandatory:"false" json:"isSkipUnavailable"`

	// Mode for the cross cluster connection
	Mode CccModeEnum `mandatory:"false" json:"mode,omitempty"`
}

func (m OutboundClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OutboundClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCccModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetCccModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
