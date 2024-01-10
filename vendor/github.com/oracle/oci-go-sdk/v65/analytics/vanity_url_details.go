// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VanityUrlDetails Vanity url configuration details.
type VanityUrlDetails struct {

	// The vanity url unique identifier key.
	Key *string `mandatory:"false" json:"key"`

	// Description of the vanity url.
	Description *string `mandatory:"false" json:"description"`

	// List of urls supported by this vanity URL definition (max of 3).
	Urls []string `mandatory:"false" json:"urls"`

	// List of fully qualified hostnames supported by this vanity URL definition (max of 3).
	Hosts []string `mandatory:"false" json:"hosts"`

	// PEM certificate for HTTPS connections.
	PublicCertificate *string `mandatory:"false" json:"publicCertificate"`
}

func (m VanityUrlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VanityUrlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
