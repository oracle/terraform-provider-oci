// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Link link between 2 nodes
type Link struct {

	// id of Link
	Id *string `mandatory:"true" json:"id"`

	// source node id
	Source *string `mandatory:"false" json:"source"`

	// destination node id
	Destination *string `mandatory:"false" json:"destination"`

	// number of times this link is repeated
	RepeatCount *int `mandatory:"false" json:"repeatCount"`

	// average packet loss
	ForwardingLoss *float64 `mandatory:"false" json:"forwardingLoss"`

	// difference of packet response time between source and destination in milliseconds
	DelayInMilliseconds *float64 `mandatory:"false" json:"delayInMilliseconds"`

	// minimum delay in milliseconds
	MinDelayInMilliseconds *float64 `mandatory:"false" json:"minDelayInMilliseconds"`

	// maximum delay in milliseconds
	MaxDelayInMilliseconds *float64 `mandatory:"false" json:"maxDelayInMilliseconds"`
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
