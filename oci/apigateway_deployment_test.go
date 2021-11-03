// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v54/apigateway"
	"github.com/oracle/oci-go-sdk/v54/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeploymentRequiredOnlyResource = DeploymentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Required, Create, deploymentRepresentationCustomAuth)

	DeploymentResourceConfig = DeploymentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Update, deploymentRepresentationCustomAuth)

	deploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": Representation{RepType: Required, Create: `${oci_apigateway_deployment.test_deployment.id}`},
	}

	deploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":     Representation{RepType: Optional, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, deploymentDataSourceFilterRepresentation}}
	deploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_apigateway_deployment.test_deployment.id}`}},
	}

	deploymentRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"gateway_id":     Representation{RepType: Required, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"path_prefix":    Representation{RepType: Required, Create: `/v1`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"specification":  RepresentationGroup{Required, deploymentSpecificationRepresentation},
	}
	deploymentSpecificationRepresentation = map[string]interface{}{
		"logging_policies": RepresentationGroup{Optional, deploymentSpecificationLoggingPoliciesRepresentation},
		"request_policies": RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesRepresentation},
		"routes":           RepresentationGroup{Required, deploymentSpecificationRoutesRepresentation},
	}
	deploymentSpecificationLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    RepresentationGroup{Optional, deploymentSpecificationLoggingPoliciesAccessLogRepresentation},
		"execution_log": RepresentationGroup{Optional, deploymentSpecificationLoggingPoliciesExecutionLogRepresentation},
	}
	deploymentSpecificationRequestPoliciesRepresentation = map[string]interface{}{
		"authentication": RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesAuthenticationRepresentation},
		"cors":           RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesCorsRepresentation},
		"mutual_tls":     RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesMutualTlsRepresentation},
		"rate_limiting":  RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesRateLimitingRepresentation},
	}
	deploymentSpecificationRoutesRepresentation = map[string]interface{}{
		"backend":           RepresentationGroup{Required, deploymentSpecificationRoutesBackendRepresentation},
		"path":              Representation{RepType: Required, Create: `/hello`, Update: `/world`},
		"logging_policies":  RepresentationGroup{Optional, deploymentSpecificationRoutesLoggingPoliciesRepresentation},
		"methods":           Representation{RepType: Required, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"request_policies":  RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesRepresentation},
		"response_policies": RepresentationGroup{Optional, deploymentSpecificationRoutesResponsePoliciesRepresentation},
	}
	deploymentSpecificationLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"log_level":  Representation{RepType: Optional, Create: `INFO`, Update: `WARN`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationRepresentation = map[string]interface{}{
		"type":                        Representation{RepType: Required, Create: `CUSTOM_AUTHENTICATION`, Update: `CUSTOM_AUTHENTICATION`},
		"audiences":                   Representation{RepType: Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"function_id":                 Representation{RepType: Optional, Create: `${oci_functions_function.test_function.id}`},
		"is_anonymous_access_allowed": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"issuers":                     Representation{RepType: Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"max_clock_skew_in_seconds":   Representation{RepType: Optional, Create: `1.0`, Update: `2.0`},
		"public_keys":                 RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation},
		"token_auth_scheme":           Representation{RepType: Optional, Create: `Bearer`, Update: `Bearer`},
		"token_header":                Representation{RepType: Optional, Create: `Authorization`, Update: `Authorization`},
		"verify_claims":               RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation},
	}
	deploymentSpecificationRequestPoliciesAuthorizeScopeRepresentation = map[string]interface{}{
		"allowed_scope": Representation{RepType: Optional, Create: []string{`cors`}},
	}
	deploymentSpecificationRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              Representation{RepType: Required, Create: []string{`https://www.oracle.org`}, Update: []string{`*`}},
		"allowed_headers":              Representation{RepType: Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              Representation{RepType: Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              Representation{RepType: Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           Representation{RepType: Optional, Create: `600`, Update: `500`},
	}
	deploymentSpecificationRequestPoliciesMutualTlsRepresentation = map[string]interface{}{
		"allowed_sans":                     Representation{RepType: Optional, Create: []string{`allowedSans`}, Update: []string{`allowedSans2`}},
		"is_verified_certificate_required": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRequestPoliciesRateLimitingRepresentation = map[string]interface{}{
		"rate_in_requests_per_second": Representation{RepType: Required, Create: `10`, Update: `11`},
		"rate_key":                    Representation{RepType: Required, Create: `CLIENT_IP`, Update: `TOTAL`},
	}
	deploymentSpecificationRoutesBackendRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `HTTP_BACKEND`, Update: `HTTP_BACKEND`},
		"url":  Representation{RepType: Required, Create: `https://api.weather.gov`, Update: `https://www.oracle.com`},
	}
	deploymentSpecificationRoutesLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    RepresentationGroup{Optional, deploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation},
		"execution_log": RepresentationGroup{Optional, deploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesRepresentation = map[string]interface{}{
		"authorization":                   RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation},
		"body_validation":                 RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation},
		"cors":                            RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesCorsRepresentation},
		"header_transformations":          RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation},
		"header_validations":              RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation},
		"query_parameter_transformations": RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation},
		"query_parameter_validations":     RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation},
		"response_cache_lookup":           RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation},
	}
	deploymentSpecificationRoutesResponsePoliciesRepresentation = map[string]interface{}{
		"header_transformations": RepresentationGroup{Optional, deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation},
		"response_cache_store":   RepresentationGroup{Optional, deploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation},
	}
	deploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation = map[string]interface{}{
		"type":                        Representation{RepType: Required, Create: `REMOTE_JWKS`, Update: `STATIC_KEYS`},
		"is_ssl_verify_disabled":      Representation{RepType: Optional, Create: `false`, Update: `true`},
		"keys":                        RepresentationGroup{Optional, deploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation},
		"max_cache_duration_in_hours": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"uri":                         Representation{RepType: Optional, Create: `https://oracle.com/jwks.json`, Update: `https://oracle.com/jwkstest.json`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"key":         Representation{RepType: Optional, Create: `key`, Update: `key2`},
		"values":      Representation{RepType: Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	deploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"log_level":  Representation{RepType: Optional, Create: `INFO`, Update: `WARN`},
	}

	deploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation = map[string]interface{}{
		"type": Representation{RepType: Optional, Create: `AUTHENTICATION_ONLY`, Update: `ANONYMOUS`},
	}
	deploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation = map[string]interface{}{
		"content":         RepresentationGroup{Required, deploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation},
		"required":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"validation_mode": Representation{RepType: Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              Representation{RepType: Required, Create: []string{`*`}, Update: []string{`*`}},
		"allowed_headers":              Representation{RepType: Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              Representation{RepType: Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              Representation{RepType: Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           Representation{RepType: Optional, Create: `600`, Update: `500`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation = map[string]interface{}{
		"headers":         RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation},
		"validation_mode": Representation{RepType: Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation = map[string]interface{}{
		"filter_query_parameters": RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation},
		"set_query_parameters":    RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation = map[string]interface{}{
		"parameters":      RepresentationGroup{Optional, deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation},
		"validation_mode": Representation{RepType: Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation = map[string]interface{}{
		"type":                       Representation{RepType: Required, Create: `SIMPLE_LOOKUP_POLICY`},
		"cache_key_additions":        Representation{RepType: Optional, Create: []string{`request.query[Foo]`}, Update: []string{`request.query[Accept]`}},
		"is_enabled":                 Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_private_caching_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": RepresentationGroup{Optional, deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    RepresentationGroup{Optional, deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	deploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation = map[string]interface{}{
		"time_to_live_in_seconds": Representation{RepType: Required, Create: `10`, Update: `11`},
		"type":                    Representation{RepType: Required, Create: `FIXED_TTL_STORE_POLICY`},
	}
	deploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation = map[string]interface{}{
		"media_type":      Representation{RepType: Required, Create: `*/*`, Update: `application/json`},
		"validation_type": Representation{RepType: Required, Create: `NONE`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  Representation{RepType: Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation = map[string]interface{}{
		"name":     Representation{RepType: Required, Create: `name`, Update: `name2`},
		"required": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation},
		"type":  Representation{RepType: Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation = map[string]interface{}{
		"name":     Representation{RepType: Required, Create: `name`, Update: `name2`},
		"required": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  Representation{RepType: Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `name`, Update: `name2`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      Representation{RepType: Required, Create: `nameA`, Update: `nameA2`},
		"values":    Representation{RepType: Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": Representation{RepType: Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `nameB`, Update: `nameB2`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation = map[string]interface{}{
		"name":      Representation{RepType: Required, Create: `nameC`, Update: `nameC2`},
		"values":    Representation{RepType: Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": Representation{RepType: Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `nameD`, Update: `nameD2`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      Representation{RepType: Required, Create: `nameE`, Update: `nameE2`},
		"values":    Representation{RepType: Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": Representation{RepType: Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation = map[string]interface{}{
		"format":  Representation{RepType: Required, Create: `PEM`, Update: `JSON_WEB_KEY`},
		"alg":     Representation{RepType: Optional, Create: `alg`, Update: `RS256`},
		"e":       Representation{RepType: Optional, Create: `e`, Update: `AQAB`},
		"key":     Representation{RepType: Optional, Create: `key`, Update: `key2`},
		"key_ops": Representation{RepType: Optional, Create: []string{}, Update: []string{`verify`}},
		"kid":     Representation{RepType: Optional, Create: `kid`, Update: `master_key`},
		"kty":     Representation{RepType: Optional, Create: `kty`, Update: `RSA`},
		"n":       Representation{RepType: Optional, Create: `n`, Update: `n2`},
		"use":     Representation{RepType: Optional, Create: `use`, Update: `sig`},
	}

	DeploymentResourceDependencies = GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Required, Create, gatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRegionalRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies

	deploymentRepresentationCustomAuth = GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication.audiences",
		"specification.request_policies.authentication.issuers",
		"specification.request_policies.authentication.max_clock_skew_in_seconds",
		"specification.request_policies.authentication.public_keys",
		"specification.request_policies.authentication.token_auth_scheme",
		"specification.request_policies.authentication.verify_claims",
	}, deploymentRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_apigateway_deployment.test_deployment"
	datasourceName := "data.oci_apigateway_deployments.test_deployments"
	singularDatasourceName := "data.oci_apigateway_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DeploymentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Create, deploymentRepresentation), "apigateway", "deployment", t)

	ResourceTest(t, testAccCheckApigatewayDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Required, Create, deploymentRepresentationCustomAuth),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + DeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
				GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Create, deploymentRepresentationCustomAuth),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "CLIENT_IP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.connect_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.read_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.send_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.url", "https://api.weather.gov"),
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
					"media_type":      "*/*",
					"validation_type": "NONE",
				}, []string{}),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameA"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.0.name", "nameB"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.name", "nameC"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.cache_key_additions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_private_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.type", "SIMPLE_LOOKUP_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.0.name", "nameD"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + imageVariableStr + DeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
				GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Create,
					RepresentationCopyWithNewProperties(deploymentRepresentationCustomAuth, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "CLIENT_IP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.connect_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.read_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.send_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.url", "https://api.weather.gov"),
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
					"media_type":      "*/*",
					"validation_type": "NONE",
				}, []string{}),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameA"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.0.name", "nameB"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.name", "nameC"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.validation_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.cache_key_additions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_private_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.type", "SIMPLE_LOOKUP_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.0.name", "nameD"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + DeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
				GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Update, deploymentRepresentationCustomAuth),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.connect_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.read_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(resourceName, "specification.0.routes.0.backend.0.send_timeout_in_seconds"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.url", "https://www.oracle.com"),
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
					"media_type":      "application/json",
					"validation_type": "NONE",
				}, []string{}),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameA2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.header_validations.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.0.name", "nameB2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.name", "nameC2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.cache_key_additions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_private_caching_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.type", "SIMPLE_LOOKUP_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.0.name", "nameD2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameE2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + imageVariableStr +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_apigateway_deployments", "test_deployments", Optional, Update, deploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Optional, Update, deploymentRepresentationCustomAuth),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
		// verify singular datasource
		{
			Config: config + imageVariableStr +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Required, Create, deploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "specification.0.routes.0.backend.0.connect_timeout_in_seconds"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "specification.0.routes.0.backend.0.read_timeout_in_seconds"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "specification.0.routes.0.backend.0.send_timeout_in_seconds"),

				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.type", "HTTP_BACKEND"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.url", "https://www.oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.access_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.access_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.logging_policies.0.execution_log.0.log_level", "WARN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.path", "/world"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.body_validation.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
					"media_type":      "application/json",
					"validation_type": "NONE",
				}, []string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.body_validation.0.required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.body_validation.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.filter_headers.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameA2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_validations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_validations.0.headers.0.required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.header_validations.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.filter_query_parameters.0.items.0.name", "nameB2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.name", "nameC2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_transformations.0.set_query_parameters.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.parameters.0.required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.query_parameter_validations.0.validation_mode", "PERMISSIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.cache_key_additions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.is_private_caching_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.request_policies.0.response_cache_lookup.0.type", "SIMPLE_LOOKUP_POLICY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.filter_headers.0.items.0.name", "nameD2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.if_exists", "SKIP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "nameE2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + DeploymentResourceConfig +
				GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApigatewayDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).deploymentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_deployment" {
			noResourceFound = false
			request := oci_apigateway.GetDeploymentRequest{}

			tmp := rs.Primary.ID
			request.DeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "apigateway")

			response, err := client.GetDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.DeploymentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("ApigatewayDeployment") {
		resource.AddTestSweepers("ApigatewayDeployment", &resource.Sweeper{
			Name:         "ApigatewayDeployment",
			Dependencies: DependencyGraph["deployment"],
			F:            sweepApigatewayDeploymentResource,
		})
	}
}

func sweepApigatewayDeploymentResource(compartment string) error {
	deploymentClient := GetTestClients(&schema.ResourceData{}).deploymentClient()
	deploymentIds, err := getDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentId := range deploymentIds {
		if ok := SweeperDefaultResourceId[deploymentId]; !ok {
			deleteDeploymentRequest := oci_apigateway.DeleteDeploymentRequest{}

			deleteDeploymentRequest.DeploymentId = &deploymentId

			deleteDeploymentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "apigateway")
			_, error := deploymentClient.DeleteDeployment(context.Background(), deleteDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting Deployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &deploymentId, deploymentSweepWaitCondition, time.Duration(3*time.Minute),
				deploymentSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getDeploymentIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deploymentClient := GetTestClients(&schema.ResourceData{}).deploymentClient()

	listDeploymentsRequest := oci_apigateway.ListDeploymentsRequest{}
	listDeploymentsRequest.CompartmentId = &compartmentId
	listDeploymentsRequest.LifecycleState = oci_apigateway.DeploymentLifecycleStateActive
	listDeploymentsResponse, err := deploymentClient.ListDeployments(context.Background(), listDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Deployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployment := range listDeploymentsResponse.Items {
		id := *deployment.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentId", id)
	}
	return resourceIds, nil
}

func deploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deploymentResponse, ok := response.Response.(oci_apigateway.GetDeploymentResponse); ok {
		return deploymentResponse.LifecycleState != oci_apigateway.DeploymentLifecycleStateDeleted
	}
	return false
}

func deploymentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.deploymentClient().GetDeployment(context.Background(), oci_apigateway.GetDeploymentRequest{
		DeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
