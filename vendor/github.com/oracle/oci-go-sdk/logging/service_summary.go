// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceSummary Summary of Services that are integrated with public logging
type ServiceSummary struct {

	// Tenant OCID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Service id as set in Service Principal.
	ServicePrincipalName *string `mandatory:"true" json:"servicePrincipalName"`

	// Service endpoint.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// User friendly service name.
	Name *string `mandatory:"true" json:"name"`

	// Type of Resource that a Service provides.
	ResourceTypes []ResourceType `mandatory:"true" json:"resourceTypes"`

	// Apollo project namespace if any.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Service id.
	Id *string `mandatory:"false" json:"id"`
}

func (m ServiceSummary) String() string {
	return common.PointerString(m)
}
