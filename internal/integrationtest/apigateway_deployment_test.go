// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v58/apigateway"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

var (
	DeploymentRequiredOnlyResource = DeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuth)

	DeploymentResourceConfig = DeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuth)

	deploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment.id}`},
	}

	deploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentDataSourceFilterRepresentation}}
	deploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_deployment.test_deployment.id}`}},
	}

	deploymentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"gateway_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"path_prefix":    acctest.Representation{RepType: acctest.Required, Create: `/v1`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"specification":  acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRepresentation},
	}
	deploymentSpecificationRepresentation = map[string]interface{}{
		"logging_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationLoggingPoliciesRepresentation},
		"request_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesRepresentation},
		"routes":           acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRepresentation},
	}
	deploymentSpecificationLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationLoggingPoliciesAccessLogRepresentation},
		"execution_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationLoggingPoliciesExecutionLogRepresentation},
	}
	deploymentSpecificationRequestPoliciesRepresentation = map[string]interface{}{
		"authentication": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesAuthenticationRepresentation},
		"cors":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesCorsRepresentation},
		"mutual_tls":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesMutualTlsRepresentation},
		"rate_limiting":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesRateLimitingRepresentation},
	}
	deploymentSpecificationRoutesRepresentation = map[string]interface{}{
		"backend":           acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesBackendRepresentation},
		"path":              acctest.Representation{RepType: acctest.Required, Create: `/hello`, Update: `/world`},
		"logging_policies":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesLoggingPoliciesRepresentation},
		"methods":           acctest.Representation{RepType: acctest.Required, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"request_policies":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesRepresentation},
		"response_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesResponsePoliciesRepresentation},
	}
	deploymentSpecificationLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"log_level":  acctest.Representation{RepType: acctest.Optional, Create: `INFO`, Update: `WARN`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_AUTHENTICATION`, Update: `CUSTOM_AUTHENTICATION`},
		"audiences":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"function_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
		"is_anonymous_access_allowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"issuers":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"max_clock_skew_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"public_keys":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation},
		"token_auth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `Bearer`, Update: `Bearer`},
		"token_header":                acctest.Representation{RepType: acctest.Optional, Create: `Authorization`, Update: `Authorization`},
		"verify_claims":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation},
	}
	deploymentSpecificationRequestPoliciesAuthorizeScopeRepresentation = map[string]interface{}{
		"allowed_scope": acctest.Representation{RepType: acctest.Optional, Create: []string{`cors`}},
	}
	deploymentSpecificationRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              acctest.Representation{RepType: acctest.Required, Create: []string{`https://www.oracle.org`}, Update: []string{`*`}},
		"allowed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `500`},
	}
	deploymentSpecificationRequestPoliciesMutualTlsRepresentation = map[string]interface{}{
		"allowed_sans":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedSans`}, Update: []string{`allowedSans2`}},
		"is_verified_certificate_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRequestPoliciesRateLimitingRepresentation = map[string]interface{}{
		"rate_in_requests_per_second": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"rate_key":                    acctest.Representation{RepType: acctest.Required, Create: `CLIENT_IP`, Update: `TOTAL`},
	}
	deploymentSpecificationRoutesBackendRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `HTTP_BACKEND`, Update: `HTTP_BACKEND`},
		"url":  acctest.Representation{RepType: acctest.Required, Create: `https://api.weather.gov`, Update: `https://www.oracle.com`},
	}
	deploymentSpecificationRoutesLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation},
		"execution_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesRepresentation = map[string]interface{}{
		"authorization":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation},
		"body_validation":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation},
		"cors":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesCorsRepresentation},
		"header_transformations":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation},
		"header_validations":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation},
		"query_parameter_transformations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation},
		"query_parameter_validations":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation},
		"response_cache_lookup":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation},
	}
	deploymentSpecificationRoutesResponsePoliciesRepresentation = map[string]interface{}{
		"header_transformations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation},
		"response_cache_store":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation},
	}
	deploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `REMOTE_JWKS`, Update: `STATIC_KEYS`},
		"is_ssl_verify_disabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"keys":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation},
		"max_cache_duration_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"uri":                         acctest.Representation{RepType: acctest.Optional, Create: `https://oracle.com/jwks.json`, Update: `https://oracle.com/jwkstest.json`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"values":      acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	deploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"log_level":  acctest.Representation{RepType: acctest.Optional, Create: `INFO`, Update: `WARN`},
	}

	deploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATION_ONLY`, Update: `ANONYMOUS`},
	}
	deploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation = map[string]interface{}{
		"content":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation},
		"required":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              acctest.Representation{RepType: acctest.Required, Create: []string{`*`}, Update: []string{`*`}},
		"allowed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `500`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation = map[string]interface{}{
		"headers":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation = map[string]interface{}{
		"filter_query_parameters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation},
		"set_query_parameters":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation = map[string]interface{}{
		"parameters":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	deploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation = map[string]interface{}{
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `SIMPLE_LOOKUP_POLICY`},
		"cache_key_additions":        acctest.Representation{RepType: acctest.Optional, Create: []string{`request.query[Foo]`}, Update: []string{`request.query[Accept]`}},
		"is_enabled":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_private_caching_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	deploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation = map[string]interface{}{
		"time_to_live_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `FIXED_TTL_STORE_POLICY`},
	}
	deploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation = map[string]interface{}{
		"media_type":      acctest.Representation{RepType: acctest.Required, Create: `*/*`, Update: `application/json`},
		"validation_type": acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	deploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameA`, Update: `nameA2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `nameB`, Update: `nameB2`},
	}
	deploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameC`, Update: `nameC2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `nameD`, Update: `nameD2`},
	}
	deploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameE`, Update: `nameE2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	deploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation = map[string]interface{}{
		"format":  acctest.Representation{RepType: acctest.Required, Create: `PEM`, Update: `JSON_WEB_KEY`},
		"alg":     acctest.Representation{RepType: acctest.Optional, Create: `alg`, Update: `RS256`},
		"e":       acctest.Representation{RepType: acctest.Optional, Create: `e`, Update: `AQAB`},
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"key_ops": acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{`verify`}},
		"kid":     acctest.Representation{RepType: acctest.Optional, Create: `kid`, Update: `master_key`},
		"kty":     acctest.Representation{RepType: acctest.Optional, Create: `kty`, Update: `RSA`},
		"n":       acctest.Representation{RepType: acctest.Optional, Create: `n`, Update: `n2`},
		"use":     acctest.Representation{RepType: acctest.Optional, Create: `use`, Update: `sig`},
	}

	DeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", acctest.Required, acctest.Create, gatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRegionalRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies

	deploymentRepresentationCustomAuth = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_apigateway_deployment.test_deployment"
	datasourceName := "data.oci_apigateway_deployments.test_deployments"
	singularDatasourceName := "data.oci_apigateway_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentation), "apigateway", "deployment", t)

	acctest.ResourceTest(t, testAccCheckApigatewayDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuth),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationCustomAuth),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + imageVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationCustomAuth, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuth),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployments", "test_deployments", acctest.Optional, acctest.Update, deploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuth),
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
		// verify singular datasource
		{
			Config: config + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceConfig,
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
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "specification.0.routes.0.request_policies.0.body_validation.0.content", map[string]string{
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation),
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DeploymentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_deployment" {
			noResourceFound = false
			request := oci_apigateway.GetDeploymentRequest{}

			tmp := rs.Primary.ID
			request.DeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ApigatewayDeployment") {
		resource.AddTestSweepers("ApigatewayDeployment", &resource.Sweeper{
			Name:         "ApigatewayDeployment",
			Dependencies: acctest.DependencyGraph["deployment"],
			F:            sweepApigatewayDeploymentResource,
		})
	}
}

func sweepApigatewayDeploymentResource(compartment string) error {
	deploymentClient := acctest.GetTestClients(&schema.ResourceData{}).DeploymentClient()
	deploymentIds, err := getDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentId := range deploymentIds {
		if ok := acctest.SweeperDefaultResourceId[deploymentId]; !ok {
			deleteDeploymentRequest := oci_apigateway.DeleteDeploymentRequest{}

			deleteDeploymentRequest.DeploymentId = &deploymentId

			deleteDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")
			_, error := deploymentClient.DeleteDeployment(context.Background(), deleteDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting Deployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deploymentId, deploymentSweepWaitCondition, time.Duration(3*time.Minute),
				deploymentSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deploymentClient := acctest.GetTestClients(&schema.ResourceData{}).DeploymentClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentId", id)
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

func deploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DeploymentClient().GetDeployment(context.Background(), oci_apigateway.GetDeploymentRequest{
		DeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
