// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_apigateway_api", ApigatewayApiResource())
	tfresource.RegisterResource("oci_apigateway_certificate", ApigatewayCertificateResource())
	tfresource.RegisterResource("oci_apigateway_deployment", ApigatewayDeploymentResource())
	tfresource.RegisterResource("oci_apigateway_gateway", ApigatewayGatewayResource())
	tfresource.RegisterResource("oci_apigateway_subscriber", ApigatewaySubscriberResource())
	tfresource.RegisterResource("oci_apigateway_usage_plan", ApigatewayUsagePlanResource())
}
