// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ServiceEnvironmentCollection Collection of service environments.
// **Note:** Service URL formats may vary from the provided example.
type ServiceEnvironmentCollection struct {

	// Collection of items.
	Items []ServiceEnvironmentSummary `mandatory:"true" json:"items"`
}

func (m ServiceEnvironmentCollection) String() string {
	return common.PointerString(m)
}
