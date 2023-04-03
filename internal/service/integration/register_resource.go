// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integration

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_integration_integration_instance", IntegrationIntegrationInstanceResource())
	tfresource.RegisterResource("oci_integration_private_endpoint_outbound_connection", IntegrationPrivateEndpointOutboundConnectionResource())
}
