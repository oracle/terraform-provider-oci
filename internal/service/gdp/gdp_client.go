// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package gdp

import (
	"strings"

	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
)

// getGdpClient returns an operation-local client when the commercial endpoint
// must be selected. Provider metadata can be shared by concurrent in-process
// reconciliations, so request-specific endpoint selection must not mutate the
// cached provider client.
func getGdpClient(m interface{}, useCommercialEndpoint bool) *oci_gdp.GuardedDataPipelineClient {
	defaultClient := m.(*client.OracleClients).GuardedDataPipelineClient()
	if !useCommercialEndpoint {
		return defaultClient
	}

	commercialClient := *defaultClient
	commercialClient.Host = strings.Replace(defaultClient.Host, "gdp", commercialSubdomain, 1)
	return &commercialClient
}
