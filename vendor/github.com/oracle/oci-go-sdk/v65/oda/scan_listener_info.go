// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScanListenerInfo Customer's Real Application Cluster (RAC)'s SCAN listener FQDN, port or list IPs and their ports.
type ScanListenerInfo struct {

	// FQDN of the customer's Real Application Cluster (RAC)'s SCAN listeners.
	ScanListenerFqdn *string `mandatory:"false" json:"scanListenerFqdn"`

	// A SCAN listener's IP of the customer's Real Application Cluster (RAC).
	ScanListenerIp *string `mandatory:"false" json:"scanListenerIp"`

	// The port that customer's Real Application Cluster (RAC)'s SCAN listeners are listening on.
	ScanListenerPort *int `mandatory:"false" json:"scanListenerPort"`
}

func (m ScanListenerInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScanListenerInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
