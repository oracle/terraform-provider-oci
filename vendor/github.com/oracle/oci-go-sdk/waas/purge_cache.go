// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PurgeCache The list of resources for cache purge. If a resources property is not provided, the purge targets all resources in a policy.
type PurgeCache struct {

	// A resource to purge, identified by either a hostless absolute path starting with a single slash (e.g., "/path/to/resource") or by a relative path in which the first component will be interpreted as a domain protected by this policy (e.g., "example.com/path/to/resource").
	Resources []string `mandatory:"false" json:"resources"`
}

func (m PurgeCache) String() string {
	return common.PointerString(m)
}
