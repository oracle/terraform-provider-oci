// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MirrorConfiguration Mirror information used for the management station configuration.
type MirrorConfiguration struct {

	// Path to the data volume on the management station where software source mirrors are stored.
	Directory *string `mandatory:"true" json:"directory"`

	// Default mirror listening port for http.
	Port *string `mandatory:"true" json:"port"`

	// Default mirror listening port for https.
	Sslport *string `mandatory:"true" json:"sslport"`

	// Path to the SSL cerfificate.
	Sslcert *string `mandatory:"false" json:"sslcert"`

	// When enabled, the SSL certificate is verified whenever an instance installs or updates a package from a software source that is mirrored on the management station.
	IsSslverifyEnabled *bool `mandatory:"false" json:"isSslverifyEnabled"`
}

func (m MirrorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MirrorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
