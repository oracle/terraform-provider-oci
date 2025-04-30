// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdcsInfoDetails Information for IDCS access
type IdcsInfoDetails struct {

	// URL for the location of the IDCS Application (used by IDCS APIs)
	IdcsAppLocationUrl *string `mandatory:"true" json:"idcsAppLocationUrl"`

	// The IDCS application display name associated with the instance
	IdcsAppDisplayName *string `mandatory:"true" json:"idcsAppDisplayName"`

	// The IDCS application ID associated with the instance
	IdcsAppId *string `mandatory:"true" json:"idcsAppId"`

	// The IDCS application name associated with the instance
	IdcsAppName *string `mandatory:"true" json:"idcsAppName"`

	// The URL used as the primary audience for visual builder flows in this instance
	// type: string
	InstancePrimaryAudienceUrl *string `mandatory:"true" json:"instancePrimaryAudienceUrl"`
}

func (m IdcsInfoDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdcsInfoDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
