// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apigateway_api", ApigatewayApiDataSource())
	tfresource.RegisterDatasource("oci_apigateway_api_content", ApigatewayApiContentDataSource())
	tfresource.RegisterDatasource("oci_apigateway_api_deployment_specification", ApigatewayApiDeploymentSpecificationDataSource())
	tfresource.RegisterDatasource("oci_apigateway_api_validation", ApigatewayApiValidationDataSource())
	tfresource.RegisterDatasource("oci_apigateway_apis", ApigatewayApisDataSource())
	tfresource.RegisterDatasource("oci_apigateway_certificate", ApigatewayCertificateDataSource())
	tfresource.RegisterDatasource("oci_apigateway_certificates", ApigatewayCertificatesDataSource())
	tfresource.RegisterDatasource("oci_apigateway_deployment", ApigatewayDeploymentDataSource())
	tfresource.RegisterDatasource("oci_apigateway_deployments", ApigatewayDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_apigateway_gateway", ApigatewayGatewayDataSource())
	tfresource.RegisterDatasource("oci_apigateway_gateways", ApigatewayGatewaysDataSource())
	tfresource.RegisterDatasource("oci_apigateway_subscriber", ApigatewaySubscriberDataSource())
	tfresource.RegisterDatasource("oci_apigateway_subscribers", ApigatewaySubscribersDataSource())
	tfresource.RegisterDatasource("oci_apigateway_usage_plan", ApigatewayUsagePlanDataSource())
	tfresource.RegisterDatasource("oci_apigateway_usage_plans", ApigatewayUsagePlansDataSource())
}
