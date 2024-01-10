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

// Link Details of the link between two nodes.
type Link struct {

	// ID of the link.
	Id *string `mandatory:"true" json:"id"`

	// ID of the source node.
	Source *string `mandatory:"false" json:"source"`

	// ID of the destination node.
	Destination *string `mandatory:"false" json:"destination"`

	// Number of times the link is repeated.
	RepeatCount *int `mandatory:"false" json:"repeatCount"`

	// Average packet loss.
	ForwardingLoss *float64 `mandatory:"false" json:"forwardingLoss"`

	// Difference of the packet response time between source and destination nodes, in milliseconds.
	DelayInMilliseconds *float64 `mandatory:"false" json:"delayInMilliseconds"`

	// Minimum delay in milliseconds.
	MinDelayInMilliseconds *float64 `mandatory:"false" json:"minDelayInMilliseconds"`

	// Maximum delay in milliseconds.
	MaxDelayInMilliseconds *float64 `mandatory:"false" json:"maxDelayInMilliseconds"`

	// List of all path IDs of which this link is part of.
	Paths []string `mandatory:"false" json:"paths"`
}

func (m Link) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Link) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
