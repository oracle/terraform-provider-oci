// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	deploymentRepresentationJwt = acctest.GetUpdatedRepresentationCopy(
		"specification.request_policies.authentication.type",
		acctest.Representation{RepType: acctest.Required, Create: `JWT_AUTHENTICATION`, Update: `JWT_AUTHENTICATION`},
		ApigatewayDeploymentRepresentation)
	deploymentRepresentationWithMtlsDisabled = acctest.GetUpdatedRepresentationCopy(
		"specification.request_policies.mutual_tls.is_verified_certificate_required",
		acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		deploymentRepresentationJwt)
	deploymentRepresentationJwtRemoteJWKS = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication.function_id",
		"specification.request_policies.authentication.public_keys.keys",
	}, deploymentRepresentationWithMtlsDisabled)
	deploymentRepresentationJwtStaticKeys = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication.function_id",
		"specification.request_policies.authentication.public_keys.uri",
		"specification.request_policies.authentication.public_keys.max_cache_duration_in_hours",
		"specification.request_policies.authentication.public_keys.is_ssl_verify_disabled",
		"specification.request_policies.authentication.public_keys.keys.key",
	}, deploymentRepresentationWithMtlsDisabled)

	DeploymentResourceConfigCommon = DeploymentResourceDependenciesWithoutCABundle + acctest.GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", acctest.Required, acctest.Create, ApigatewayRepresentation)

	DeploymentResourceConfigJwt = DeploymentResourceConfigCommon +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationJwtStaticKeys)
)

// issue-routing-tag: apigateway/default
func TestResourceApigatewayDeploymentResourceJwt_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayDeploymentResourceJwt_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_apigateway_deployment.test_deployment"
	datasourceName := "data.oci_apigateway_deployments.test_deployments"
	singularDatasourceName := "data.oci_apigateway_deployment.test_deployment"

	var resId string

	acctest.ResourceTest(t, testAccCheckApigatewayDeploymentDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceConfigCommon +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationJwtRemoteJWKS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceConfigCommon,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + DeploymentResourceConfigCommon +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationJwtRemoteJWKS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.access_log.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.0.log_level", "INFO"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "JWT_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.max_clock_skew_in_seconds", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.max_cache_duration_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.type", "REMOTE_JWKS"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.uri", "https://oracle.com/jwks.json"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.is_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "CLIENT_IP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.url", "https://api.weather.gov"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.connect_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.read_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.send_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.0.selector", "request.headers[route]"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.0.type", "SINGLE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.access_log.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.log_level", "INFO"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.path", "/hello"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.authorization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.authorization.0.type", "AUTHENTICATION_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "600"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + DeploymentResourceConfigCommon +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationJwtStaticKeys),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.access_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.logging_policies.0.execution_log.0.log_level", "WARN"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "JWT_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.max_clock_skew_in_seconds", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.type", "STATIC_KEYS"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.format", "JSON_WEB_KEY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.kty", "RSA"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.n", "n2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.e", "AQAB"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.alg", "RS256"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.use", "sig"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.kid", "master_key"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.key_ops.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.is_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.url", "https://www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.0.selector", "request.subdomain[oracle.com]"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.selection_source.0.type", "SINGLE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.access_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.log_level", "WARN"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.path", "/world"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.authorization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.authorization.0.type", "ANONYMOUS"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "500"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + imageVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployments", "test_deployments", acctest.Optional, acctest.Update, ApigatewayDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceConfigCommon +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationJwtStaticKeys),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "gateway_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_collection.0.endpoint"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_collection.0.gateway_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_collection.0.id"),
			),
		},

		//verify singular datasource
		{
			Config: config + imageVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, ApigatewayDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceConfigJwt,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.0.access_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.0.execution_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.logging_policies.0.execution_log.0.log_level", "WARN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.type", "JWT_AUTHENTICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.max_clock_skew_in_seconds", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.type", "STATIC_KEYS"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.format", "JSON_WEB_KEY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.kty", "RSA"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.n", "n2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.e", "AQAB"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.alg", "RS256"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.use", "sig"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.kid", "master_key"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.public_keys.0.keys.0.key_ops.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.is_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.backend.0.url", "https://www.oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.selection_source.0.selector", "request.subdomain[oracle.com]"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.selection_source.0.type", "SINGLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.access_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.log_level", "WARN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.path", "/world"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + ApigatewayDeploymentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
			},
			ResourceName: resourceName,
		},
	})
}
