// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmFencingPolicy Represents a cluster fencing policy.
type OlvmFencingPolicy struct {

	// Enable or disable fencing on this cluster.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	SkipIfConnectivityBroken *SkipIfConnectivityBroken `mandatory:"false" json:"skipIfConnectivityBroken"`

	// A flag indicating if fencing should be skipped if Gluster bricks are up and running in the host being fenced.
	IsSkipIfGlusterBricksUp *bool `mandatory:"false" json:"isSkipIfGlusterBricksUp"`

	// A flag indicating if fencing should be skipped if Gluster bricks are up and running and Gluster quorum will not be met without those bricks.
	IsSkipIfGlusterQuorumNotMet *bool `mandatory:"false" json:"isSkipIfGlusterQuorumNotMet"`

	SkipIfSdActive *SkipIfSdActive `mandatory:"false" json:"skipIfSdActive"`
}

func (m OlvmFencingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmFencingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
