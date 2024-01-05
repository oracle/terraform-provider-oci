// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	ApigatewayDeploymentRequiredOnlyResource = DeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuthWithTokenHeader) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Required, acctest.Create, deploymentRepresentationDynamicAuth) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Required, acctest.Create, deploymentRepresentationOidc)

	ApigatewayDeploymentResourceConfig = DeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuthWithTokenHeader) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Optional, acctest.Update, deploymentRepresentationDynamicAuth) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Optional, acctest.Update, deploymentRepresentationOidc)

	ApigatewayDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment.id}`},
	}
	ApigatewayDeploymentWithDynamicAuthenticationSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment_with_dynamic_auth.id}`},
	}
	ApigatewayDeploymentWithOidcSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment_with_oidc.id}`},
	}

	ApigatewayDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentDataSourceFilterRepresentation},
	}

	ApigatewayDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_deployment.test_deployment.id}`}},
	}

	ApigatewayDeploymentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"gateway_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"path_prefix":    acctest.Representation{RepType: acctest.Required, Create: `/v1`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"specification":  acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentIgnoreChangesDeploymentRepresentation},
	}

	ApigatewayDeploymentIgnoreChangesDeploymentRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	ApigatewayDeploymentSpecificationRepresentation = map[string]interface{}{
		"logging_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesRepresentation},
		"request_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesRepresentation},
		"routes":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRepresentation},
	}
	ApigatewayDeploymentSpecificationTokenAuthRepresentation = map[string]interface{}{
		"logging_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesRepresentation},
		"request_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthRepresentation},
		"routes":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRepresentation},
	}
	ApigatewayDeploymentSpecificationLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesAccessLogRepresentation},
		"execution_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesExecutionLogRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesRepresentation = map[string]interface{}{
		"authentication":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationRepresentation},
		"cors":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesCorsRepresentation},
		"dynamic_authentication": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationRepresentation},
		"mutual_tls":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesMutualTlsRepresentation},
		"rate_limiting":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesRateLimitingRepresentation},
		"usage_plans":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesUsagePlansRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthRepresentation = map[string]interface{}{
		"authentication": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRepresentation = map[string]interface{}{
		"backend":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesBackendRepresentation},
		"path":              acctest.Representation{RepType: acctest.Required, Create: `/hello`, Update: `/world`},
		"logging_policies":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesLoggingPoliciesRepresentation},
		"methods":           acctest.Representation{RepType: acctest.Required, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"request_policies":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesRepresentation},
		"response_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesRepresentation},
	}
	ApigatewayDeploymentSpecificationLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"log_level":  acctest.Representation{RepType: acctest.Optional, Create: `INFO`, Update: `WARN`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_AUTHENTICATION`, Update: `CUSTOM_AUTHENTICATION`},
		"audiences":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"function_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
		"is_anonymous_access_allowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"issuers":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"max_clock_skew_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"public_keys":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation},
		"token_auth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `Bearer`, Update: `Bearer`},
		"token_header":                acctest.Representation{RepType: acctest.Optional, Create: `Authorization`, Update: `Authorization`},
		"verify_claims":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation},
		"cache_key":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`foo`}, Update: nil},
		"parameters":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"foo": "request.headers[abc]"}, Update: map[string]string{"foo": "request.headers[abc]", "bar": "request.query[def]"}},
		"validation_failure_policy":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyRepresentation},
		"validation_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `TOKEN_AUTHENTICATION`, Update: `TOKEN_AUTHENTICATION`},
		"token_auth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `Bearer`, Update: `Bearer`},
		"token_header":                acctest.Representation{RepType: acctest.Optional, Create: `Authorization`, Update: `Authorization`},
		"is_anonymous_access_allowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"validation_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicyRepresentation},
		"validation_failure_policy":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationFailurePolicyRepresentation},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              acctest.Representation{RepType: acctest.Required, Create: []string{`https://www.oracle.org`}, Update: []string{`*`}},
		"allowed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `500`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationRepresentation = map[string]interface{}{
		"authentication_servers": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersRepresentation},
		"selection_source":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationSelectionSourceRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesMutualTlsRepresentation = map[string]interface{}{
		"allowed_sans":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedSans`}, Update: []string{`allowedSans2`}},
		"is_verified_certificate_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesRateLimitingRepresentation = map[string]interface{}{
		"rate_in_requests_per_second": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"rate_key":                    acctest.Representation{RepType: acctest.Required, Create: `CLIENT_IP`, Update: `TOTAL`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesUsagePlansRepresentation = map[string]interface{}{
		"token_locations": acctest.Representation{RepType: acctest.Required, Create: []string{`request.headers[apiKeyLocation]`}, Update: []string{`request.path[apiKeyLocation]`}},
	}
	ApigatewayDeploymentSpecificationRoutesBackendRepresentation = map[string]interface{}{
		"type":             acctest.Representation{RepType: acctest.Required, Create: `DYNAMIC_ROUTING_BACKEND`},
		"routing_backends": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendRepresentation},
		"selection_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesBackendSelectionSourceRepresentation},
	}

	ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendRepresentation = map[string]interface{}{
		"key":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendKeyRepresentation},
		"backend": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendBackendRepresentation},
	}

	ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendKeyRepresentation = map[string]interface{}{
		"type":       acctest.Representation{RepType: acctest.Required, Create: `ANY_OF`},
		"values":     acctest.Representation{RepType: acctest.Required, Create: []string{`abc`, `def`}, Update: []string{`xyz`}},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `key1`, Update: `key2`},
		"is_default": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ApigatewayDeploymentSpecificationRoutesBackendRoutingBackendBackendRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `HTTP_BACKEND`, Update: `HTTP_BACKEND`},
		"url":  acctest.Representation{RepType: acctest.Required, Create: `https://api.weather.gov`, Update: `https://www.oracle.com`},
	}

	ApigatewayDeploymentSpecificationRoutesLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation},
		"execution_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesRepresentation = map[string]interface{}{
		"authorization":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation},
		"body_validation":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation},
		"cors":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesCorsRepresentation},
		"header_transformations":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation},
		"header_validations":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation},
		"query_parameter_transformations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation},
		"query_parameter_validations":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation},
		"response_cache_lookup":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesRepresentation = map[string]interface{}{
		"header_transformations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation},
		"response_cache_store":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationPublicKeysRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `REMOTE_JWKS`, Update: `STATIC_KEYS`},
		"is_ssl_verify_disabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"keys":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation},
		"max_cache_duration_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"uri":                         acctest.Representation{RepType: acctest.Optional, Create: `https://oracle.com/jwks.json`, Update: `https://oracle.com/jwkstest.json`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"values":      acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `MODIFY_RESPONSE`, Update: `MODIFY_RESPONSE`},
		"client_details":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyClientDetailsRepresentation},
		"fallback_redirect_path":             acctest.Representation{RepType: acctest.Optional, Create: `fallbackRedirectPath`, Update: `fallbackRedirectPath2`},
		"logout_path":                        acctest.Representation{RepType: acctest.Optional, Create: `logoutPath`, Update: `logoutPath2`},
		"max_expiry_duration_in_hours":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"response_code":                      acctest.Representation{RepType: acctest.Optional, Create: `210`, Update: `220`},
		"response_header_transformations":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRepresentation},
		"response_message":                   acctest.Representation{RepType: acctest.Optional, Create: `responseMessage`, Update: `responseMessage2`},
		"response_type":                      acctest.Representation{RepType: acctest.Optional, Create: `CODE`},
		"scopes":                             acctest.Representation{RepType: acctest.Optional, Create: []string{`scopes`}, Update: []string{`scopes2`}},
		"source_uri_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicySourceUriDetailsRepresentation},
		"use_cookies_for_intermediate_steps": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_cookies_for_session":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_pkce":                           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationFailurePolicyRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `OAUTH2`, Update: `OAUTH2`},
		"client_details":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationFailurePolicyClientDetailsRepresentation},
		"source_uri_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicySourceUriDetailsRepresentation},
		"scopes":                             acctest.Representation{RepType: acctest.Optional, Create: []string{`openid`, `scopes1`}, Update: []string{`openid`, `scopes2`}},
		"max_expiry_duration_in_hours":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"use_cookies_for_intermediate_steps": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_cookies_for_session":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_pkce":                           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"fallback_redirect_path":             acctest.Representation{RepType: acctest.Optional, Create: `/fallbackRedirectPath`, Update: `/fallbackRedirectPath2`},
		"response_type":                      acctest.Representation{RepType: acctest.Optional, Create: `CODE`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `STATIC_KEYS`, Update: `REMOTE_JWKS`},
		"additional_validation_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyAdditionalValidationPolicyRepresentation},
		"client_details":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyClientDetailsRepresentation},
		"is_ssl_verify_disabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"keys":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyKeysRepresentation},
		"max_cache_duration_in_hours":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"source_uri_details":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicySourceUriDetailsRepresentation},
		"uri":                          acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicyRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `REMOTE_DISCOVERY`, Update: `REMOTE_DISCOVERY`},
		"client_details":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicyClientDetailsRepresentation},
		"source_uri_details":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicySourceUriDetailsRepresentation},
		"is_ssl_verify_disabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_cache_duration_in_hours":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"additional_validation_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyAdditionalValidationPolicyRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersRepresentation = map[string]interface{}{
		"authentication_server_detail": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailRepresentation},
		"key":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersKeyRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationSelectionSourceRepresentation = map[string]interface{}{
		"selector": acctest.Representation{RepType: acctest.Required, Create: `request.headers[tenant]`, Update: `request.subdomain[oracle.com]`},
		"type":     acctest.Representation{RepType: acctest.Required, Create: `SINGLE`},
	}

	ApigatewayDeploymentSpecificationRoutesBackendSelectionSourceRepresentation = map[string]interface{}{
		"selector": acctest.Representation{RepType: acctest.Required, Create: `request.headers[route]`, Update: `request.subdomain[oracle.com]`},
		"type":     acctest.Representation{RepType: acctest.Required, Create: `SINGLE`},
	}
	ApigatewayDeploymentSpecificationRoutesLoggingPoliciesAccessLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRoutesLoggingPoliciesExecutionLogRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"log_level":  acctest.Representation{RepType: acctest.Optional, Create: `INFO`, Update: `WARN`},
	}

	ApigatewayDeploymentSpecificationRoutesRequestPoliciesAuthorizationRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATION_ONLY`, Update: `ANONYMOUS`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesBodyValidationRepresentation = map[string]interface{}{
		"content":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation},
		"required":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              acctest.Representation{RepType: acctest.Required, Create: []string{`*`}, Update: []string{`*`}},
		"allowed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `500`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderValidationsRepresentation = map[string]interface{}{
		"headers":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsRepresentation = map[string]interface{}{
		"filter_query_parameters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation},
		"set_query_parameters":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsRepresentation = map[string]interface{}{
		"parameters":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation},
		"validation_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENFORCING`, Update: `PERMISSIVE`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesResponseCacheLookupRepresentation = map[string]interface{}{
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `SIMPLE_LOOKUP_POLICY`},
		"cache_key_additions":        acctest.Representation{RepType: acctest.Optional, Create: []string{`request.query[Foo]`}, Update: []string{`request.query[Accept]`}},
		"is_enabled":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_private_caching_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesResponseCacheStoreRepresentation = map[string]interface{}{
		"time_to_live_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `FIXED_TTL_STORE_POLICY`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyClientDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `CUSTOM`},
		"client_id":                    acctest.Representation{RepType: acctest.Optional, Create: `example_client_id`},
		"client_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.client_secret_id}`},
		"client_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationFailurePolicyClientDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `VALIDATION_BLOCK`},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersRepresentation},
		"rename_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicySourceUriDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `VALIDATION_BLOCK`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyAdditionalValidationPolicyRepresentation = map[string]interface{}{
		"audiences":     acctest.Representation{RepType: acctest.Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"issuers":       acctest.Representation{RepType: acctest.Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"verify_claims": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyAdditionalValidationPolicyVerifyClaimsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyClientDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `CUSTOM`},
		"client_id":                    acctest.Representation{RepType: acctest.Optional, Create: `example_client_id`},
		"client_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.client_secret_id}`},
		"client_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicyClientDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`, Update: `CUSTOM`},
		"client_id":                    acctest.Representation{RepType: acctest.Optional, Create: `example_client_id`},
		"client_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.client_secret_id}`},
		"client_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyKeysRepresentation = map[string]interface{}{
		"format":  acctest.Representation{RepType: acctest.Required, Create: `JSON_WEB_KEY`, Update: `PEM`},
		"alg":     acctest.Representation{RepType: acctest.Optional, Create: `alg`, Update: `alg2`},
		"e":       acctest.Representation{RepType: acctest.Optional, Create: `e`, Update: `e2`},
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"key_ops": acctest.Representation{RepType: acctest.Optional, Create: []string{`keyOps`}, Update: []string{`keyOps2`}},
		"kid":     acctest.Representation{RepType: acctest.Optional, Create: `kid`, Update: `kid2`},
		"kty":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`},
		"n":       acctest.Representation{RepType: acctest.Optional, Create: `n`, Update: `n2`},
		"use":     acctest.Representation{RepType: acctest.Optional, Create: `sig`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicySourceUriDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DISCOVERY_URI`, Update: `VALIDATION_BLOCK`},
		"uri":  acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesTokenAuthenticationValidationPolicySourceUriDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DISCOVERY_URI`, Update: `DISCOVERY_URI`},
		"uri":  acctest.Representation{RepType: acctest.Optional, Create: `https://oracle.com/discovery`, Update: `https://oracle.com/discovery2`},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_AUTHENTICATION`, Update: `CUSTOM_AUTHENTICATION`},
		"audiences":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"function_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
		"is_anonymous_access_allowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"issuers":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"max_clock_skew_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"public_keys":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailPublicKeysRepresentation},
		"token_auth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `tokenAuthScheme`, Update: `tokenAuthScheme2`},
		"token_header":                acctest.Representation{RepType: acctest.Optional, Create: `tokenHeader`, Update: `tokenHeader2`},
		"token_query_param":           acctest.Representation{RepType: acctest.Optional, Create: `tokenQueryParam`, Update: `tokenQueryParam2`},
		"validation_failure_policy":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyRepresentation},
		"validation_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyRepresentation},
		"verify_claims":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailVerifyClaimsRepresentation},
	}

	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersKeyRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"is_default": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"type":       acctest.Representation{RepType: acctest.Optional, Create: `ANY_OF`},
		"values":     acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesBodyValidationContentRepresentation = map[string]interface{}{
		"media_type":      acctest.Representation{RepType: acctest.Required, Create: `*/*`, Update: `application/json`},
		"validation_type": acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderValidationsHeadersRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterValidationsParametersRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `BLOCK`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationPolicyAdditionalValidationPolicyVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"values":      acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailPublicKeysRepresentation = map[string]interface{}{
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `STATIC_KEYS`, Update: `REMOTE_JWKS`},
		"is_ssl_verify_disabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"keys":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailPublicKeysKeysRepresentation},
		"max_cache_duration_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"uri":                         acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `MODIFY_RESPONSE`, Update: `MODIFY_RESPONSE`},
		"client_details":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyClientDetailsRepresentation},
		"fallback_redirect_path":             acctest.Representation{RepType: acctest.Optional, Create: `fallbackRedirectPath`, Update: `fallbackRedirectPath2`},
		"logout_path":                        acctest.Representation{RepType: acctest.Optional, Create: `logoutPath`, Update: `logoutPath2`},
		"max_expiry_duration_in_hours":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"response_code":                      acctest.Representation{RepType: acctest.Optional, Create: `200`, Update: `404`},
		"response_header_transformations":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRepresentation},
		"response_message":                   acctest.Representation{RepType: acctest.Optional, Create: `responseMessage`, Update: `responseMessage2`},
		"response_type":                      acctest.Representation{RepType: acctest.Optional, Create: `CODE`},
		"scopes":                             acctest.Representation{RepType: acctest.Optional, Create: []string{`scopes`}, Update: []string{`scopes2`}},
		"source_uri_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicySourceUriDetailsRepresentation},
		"use_cookies_for_intermediate_steps": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_cookies_for_session":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"use_pkce":                           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `STATIC_KEYS`, Update: `REMOTE_JWKS`},
		"additional_validation_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyAdditionalValidationPolicyRepresentation},
		"client_details":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyClientDetailsRepresentation},
		"is_ssl_verify_disabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"keys":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyKeysRepresentation},
		"max_cache_duration_in_hours":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"source_uri_details":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicySourceUriDetailsRepresentation},
		"uri":                          acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"values":      acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameA`, Update: `nameA2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `nameB`, Update: `nameB2`},
	}
	ApigatewayDeploymentSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameC`, Update: `nameC2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `nameD`, Update: `nameD2`},
	}
	ApigatewayDeploymentSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameE`, Update: `name2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `namefilter`, Update: `namefilter2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersItemsRepresentation = map[string]interface{}{
		"from": acctest.Representation{RepType: acctest.Optional, Create: `from`, Update: `from2`},
		"to":   acctest.Representation{RepType: acctest.Optional, Create: `to`, Update: `to2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `APPEND`},
		"name":      acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"values":    acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailPublicKeysKeysRepresentation = map[string]interface{}{
		"format":  acctest.Representation{RepType: acctest.Required, Create: `JSON_WEB_KEY`, Update: `PEM`},
		"alg":     acctest.Representation{RepType: acctest.Optional, Create: `alg`, Update: `alg2`},
		"e":       acctest.Representation{RepType: acctest.Optional, Create: `e`, Update: `e2`},
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"key_ops": acctest.Representation{RepType: acctest.Optional, Create: []string{`keyOps`}, Update: []string{`keyOps2`}},
		"kid":     acctest.Representation{RepType: acctest.Optional, Create: `kid`, Update: `kid2`},
		"kty":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`},
		"n":       acctest.Representation{RepType: acctest.Optional, Create: `n`, Update: `n2`},
		"use":     acctest.Representation{RepType: acctest.Optional, Create: `sig`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyClientDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `CUSTOM`},
		"client_id":                    acctest.Representation{RepType: acctest.Optional, Create: `example_client_id`},
		"client_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.client_secret_id}`},
		"client_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsFilterHeadersRepresentation},
		"rename_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRenameHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsSetHeadersRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicySourceUriDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DISCOVERY_URI`, Update: `VALIDATION_BLOCK`},
		"uri":  acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyAdditionalValidationPolicyRepresentation = map[string]interface{}{
		"audiences":     acctest.Representation{RepType: acctest.Optional, Create: []string{`audiences`}, Update: []string{`audiences2`}},
		"issuers":       acctest.Representation{RepType: acctest.Optional, Create: []string{`issuers`}, Update: []string{`issuers2`}},
		"verify_claims": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyAdditionalValidationPolicyVerifyClaimsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyClientDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `VALIDATION_BLOCK`, Update: `CUSTOM`},
		"client_id":                    acctest.Representation{RepType: acctest.Optional, Create: `example_client_id`},
		"client_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.client_secret_id}`},
		"client_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyKeysRepresentation = map[string]interface{}{
		"format":  acctest.Representation{RepType: acctest.Required, Create: `JSON_WEB_KEY`, Update: `PEM`},
		"alg":     acctest.Representation{RepType: acctest.Optional, Create: `alg`, Update: `alg2`},
		"e":       acctest.Representation{RepType: acctest.Optional, Create: `e`, Update: `e2`},
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"key_ops": acctest.Representation{RepType: acctest.Optional, Create: []string{`keyOps`}, Update: []string{`keyOps2`}},
		"kid":     acctest.Representation{RepType: acctest.Optional, Create: `kid`, Update: `kid2`},
		"kty":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`},
		"n":       acctest.Representation{RepType: acctest.Optional, Create: `n`, Update: `n2`},
		"use":     acctest.Representation{RepType: acctest.Optional, Create: `sig`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicySourceUriDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DISCOVERY_URI`, Update: `VALIDATION_BLOCK`},
		"uri":  acctest.Representation{RepType: acctest.Optional, Create: `uri`, Update: `uri2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsFilterHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsFilterHeadersItemsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Optional, Create: `BLOCK`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRenameHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRenameHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsSetHeadersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsSetHeadersItemsRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationPolicyAdditionalValidationPolicyVerifyClaimsRepresentation = map[string]interface{}{
		"is_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"values":      acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `namefilter`, Update: `namefilter2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsRenameHeadersItemsRepresentation = map[string]interface{}{
		"from": acctest.Representation{RepType: acctest.Optional, Create: `from`, Update: `from2`},
		"to":   acctest.Representation{RepType: acctest.Optional, Create: `to`, Update: `to2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesDynamicAuthenticationAuthenticationServersAuthenticationServerDetailValidationFailurePolicyResponseHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `APPEND`},
		"name":      acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"values":    acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationPublicKeysKeysRepresentation = map[string]interface{}{
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

	gatewayCaBundlesRepresentationWithCABundle = acctest.GetUpdatedRepresentationCopy(
		"ca_bundle_id",
		acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_ca_bundle.test_ca_bundle_dep.id}`},
		gatewayCaBundlesRepresentation)

	gatewayRepresentationWithCABundle = acctest.GetUpdatedRepresentationCopy(
		"ca_bundles",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: gatewayCaBundlesRepresentationWithCABundle},
		ApigatewayRepresentation)

	DeploymentResourceGatewayDependency = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", acctest.Required, acctest.Create, gatewayRepresentationWithCABundle)

	DeploymentResourceDependenciesWithoutCABundle = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRegionalRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies

	DeploymentResourceDependencies = DeploymentResourceGatewayDependency + DeploymentResourceDependenciesWithoutCABundle +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle_dep", acctest.Optional, acctest.Create, caBundleRepresentation)

	deploymentRepresentationCustomAuthWithTokenHeader = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		// A deployment specification cannot have both auth and dynamic auth and thus exclude the dynamic auth request policy
		"specification.request_policies.dynamic_authentication",
		"specification.request_policies.authentication.audiences",
		"specification.request_policies.authentication.issuers",
		"specification.request_policies.authentication.max_clock_skew_in_seconds",
		"specification.request_policies.authentication.public_keys",
		"specification.request_policies.authentication.token_auth_scheme",
		"specification.request_policies.authentication.verify_claims",
		"specification.request_policies.authentication.validation_failure_policy",
		"specification.request_policies.authentication.validation_policy",
		"specification.request_policies.authentication.parameters",
		"specification.request_policies.authentication.cache_key",
	}, ApigatewayDeploymentRepresentation)

	deploymentRepresentationCustomAuthWithParamCacheKey = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		// A deployment specification cannot have both auth and dynamic auth and thus exclude the dynamic auth request policy
		"specification.request_policies.dynamic_authentication",
		"specification.request_policies.authentication.audiences",
		"specification.request_policies.authentication.issuers",
		"specification.request_policies.authentication.max_clock_skew_in_seconds",
		"specification.request_policies.authentication.public_keys",
		"specification.request_policies.authentication.token_auth_scheme",
		"specification.request_policies.authentication.verify_claims",
		"specification.request_policies.authentication.validation_policy",
		"specification.request_policies.authentication.validation_failure_policy.client_details",
		"specification.request_policies.authentication.validation_failure_policy.fallback_redirect_path",
		"specification.request_policies.authentication.validation_failure_policy.logout_path",
		"specification.request_policies.authentication.validation_failure_policy.max_expiry_duration_in_hours",
		"specification.request_policies.authentication.validation_failure_policy.scopes",
		"specification.request_policies.authentication.validation_failure_policy.source_uri_details",
		"specification.request_policies.authentication.validation_failure_policy.use_cookies_for_intermediate_steps",
		"specification.request_policies.authentication.validation_failure_policy.use_cookies_for_session",
		"specification.request_policies.authentication.validation_failure_policy.use_pkce",
		"specification.request_policies.authentication.validation_failure_policy.response_type",
		"specification.request_policies.authentication.token_header",
	}, ApigatewayDeploymentRepresentation)

	deploymentRepresentationRequestBasedAuthCustomAuth = acctest.GetUpdatedRepresentationCopy(
		"path_prefix",
		acctest.Representation{RepType: acctest.Required, Create: `/v2`},
		deploymentRepresentationCustomAuthWithParamCacheKey)

	// Creating a new deployment for dynamic authentication as same deployment cant have both authentication and
	// dynamic authentication policies
	deploymentRepresentationWithNewPathPrefix3 = acctest.GetUpdatedRepresentationCopy(
		"path_prefix",
		acctest.Representation{RepType: acctest.Required, Create: `/v3`},
		ApigatewayDeploymentRepresentation)

	deploymentRepresentationDynamicAuth = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.audiences",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.issuers",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.max_clock_skew_in_seconds",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.public_keys",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.token_auth_scheme",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.verify_claims",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.token_query_param",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.validation_policy",
		"specification.request_policies.dynamic_authentication.authentication_servers.authentication_server_detail.validation_failure_policy",
	}, deploymentRepresentationWithNewPathPrefix3)

	deploymentRepresentationOidc = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"gateway_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_gateway.test_gateway.id}`},
		"path_prefix":    acctest.Representation{RepType: acctest.Required, Create: `/v4`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"specification":  acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentSpecificationTokenAuthRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayDeploymentIgnoreChangesDeploymentRepresentation},
	}
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

	clientSectetId := utils.GetEnvSettingWithBlankDefault("client_secret_id")
	clientSecretIdVariableStr := fmt.Sprintf("variable \"client_secret_id\" { default = \"%s\" }\n", clientSectetId)

	resourceName := "oci_apigateway_deployment.test_deployment"
	resourceNameRba := "oci_apigateway_deployment.test_deployment_rba"
	datasourceName := "data.oci_apigateway_deployments.test_deployments"
	singularDatasourceName := "data.oci_apigateway_deployment.test_deployment"

	resourceNameWithDynamicAuth := "oci_apigateway_deployment.test_deployment_with_dynamic_auth"
	singularDatasourceNameWithDynamicAuth := "data.oci_apigateway_deployment.test_deployment_with_dynamic_auth"

	resourceNameWithOidc := "oci_apigateway_deployment.test_deployment_with_oidc"
	singularDatasourceNameWithOidc := "data.oci_apigateway_deployment.test_deployment_with_oidc"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+imageVariableStr+clientSecretIdVariableStr+DeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationCustomAuthWithTokenHeader)+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth)+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Optional, acctest.Create, deploymentRepresentationDynamicAuth)+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Optional, acctest.Create, deploymentRepresentationOidc), "apigateway", "deployment", t)

	acctest.ResourceTest(t, testAccCheckApigatewayDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuthWithTokenHeader) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Required, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Required, acctest.Create, deploymentRepresentationDynamicAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Required, acctest.Create, deploymentRepresentationOidc),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "gateway_id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "path_prefix", "/v3"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.#", "1"),

				resource.TestCheckResourceAttr(resourceNameRba, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRba, "gateway_id"),
				resource.TestCheckResourceAttr(resourceNameRba, "path_prefix", "/v2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.#", "1"),

				resource.TestCheckResourceAttr(resourceNameWithOidc, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "gateway_id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "path_prefix", "/v4"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.#", "1"),

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
			Config: config + compartmentIdVariableStr + imageVariableStr + clientSecretIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationCustomAuthWithTokenHeader) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Optional, acctest.Create, deploymentRepresentationDynamicAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Optional, acctest.Create, deploymentRepresentationOidc),

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
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.dynamic_authentication.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "10"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "CLIENT_IP"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.0.token_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.is_default", "false"),
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
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "path_prefix", "/v3"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.token_header", "tokenHeader"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.is_default", "false"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.name", "name"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.selector", "request.headers[tenant]"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.type", "SINGLE"),

				// Test the auth part for request based auth
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.cache_key.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.cache_key.0", "foo"),
				resource.TestCheckResourceAttrSet(resourceNameRba, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_code", "210"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.0.name", "namefilter"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.from", "from"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.to", "to"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_message", "responseMessage"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "MODIFY_RESPONSE"),
				// Test the  parts for oidc
				resource.TestCheckResourceAttr(resourceNameWithOidc, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "path_prefix", "/v4"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.type", "TOKEN_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.type", "REMOTE_DISCOVERY"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.max_cache_duration_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_id", "example_client_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_version_number", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.type", "DISCOVERY_URI"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.uri", "https://oracle.com/discovery"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.is_required", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.key", "key"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "OAUTH2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.scopes.#", "2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.max_expiry_duration_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_intermediate_steps", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_session", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_pkce", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.fallback_redirect_path", "/fallbackRedirectPath"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_type", "CODE"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.0.type", "VALIDATION_BLOCK"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.0.type", "VALIDATION_BLOCK"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + imageVariableStr + clientSecretIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationCustomAuthWithTokenHeader, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationRequestBasedAuthCustomAuth, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationDynamicAuth, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationOidc, map[string]interface{}{
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.0.token_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.is_default", "false"),
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
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "path_prefix", "/v3"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.token_header", "tokenHeader"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.is_default", "false"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.name", "name"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.selector", "request.headers[tenant]"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.type", "SINGLE"),

				// Test the auth part for request based auth
				resource.TestCheckResourceAttr(resourceNameRba, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameRba, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.cache_key.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.cache_key.0", "foo"),
				resource.TestCheckResourceAttrSet(resourceNameRba, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_code", "210"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.0.name", "namefilter"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.from", "from"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.to", "to"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.if_exists", "OVERWRITE"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_message", "responseMessage"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "MODIFY_RESPONSE"),

				// Test the  parts for oidc
				resource.TestCheckResourceAttr(resourceNameWithOidc, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "path_prefix", "/v4"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.type", "TOKEN_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.type", "REMOTE_DISCOVERY"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.is_ssl_verify_disabled", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.max_cache_duration_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_id", "example_client_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_version_number", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.type", "DISCOVERY_URI"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.uri", "https://oracle.com/discovery"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.is_required", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.key", "key"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "OAUTH2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.scopes.#", "2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.max_expiry_duration_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_intermediate_steps", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_session", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_pkce", "false"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.fallback_redirect_path", "/fallbackRedirectPath"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_type", "CODE"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.0.type", "VALIDATION_BLOCK"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.0.type", "VALIDATION_BLOCK"),

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
			Config: config + compartmentIdVariableStr + imageVariableStr + clientSecretIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuthWithTokenHeader) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Update, deploymentRepresentationRequestBasedAuthCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Optional, acctest.Update, deploymentRepresentationDynamicAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Optional, acctest.Update, deploymentRepresentationOidc),

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
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.dynamic_authentication.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.request_policies.0.usage_plans.0.token_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.is_default", "true"),
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
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "path_prefix", "/v3"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.token_header", "tokenHeader2"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.is_default", "true"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.selector", "request.subdomain[oracle.com]"),
				resource.TestCheckResourceAttr(resourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.type", "SINGLE"),

				// Test the auth part for request based auth
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameRba, "specification.0.request_policies.0.authentication.0.function_id"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "true"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.parameters.%", "2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_code", "220"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.items.0.name", "namefilter2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.filter_headers.0.type", "BLOCK"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.from", "from2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.rename_headers.0.items.0.to", "to2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.if_exists", "APPEND"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_message", "responseMessage2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "MODIFY_RESPONSE"),

				// Test the  parts for oidc
				resource.TestCheckResourceAttr(resourceNameWithOidc, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "endpoint"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "path_prefix", "/v4"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.type", "TOKEN_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.type", "REMOTE_DISCOVERY"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.is_ssl_verify_disabled", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.max_cache_duration_in_hours", "11"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_id", "example_client_id"),
				resource.TestCheckResourceAttrSet(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_id"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_version_number", "2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.type", "DISCOVERY_URI"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.uri", "https://oracle.com/discovery2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.is_required", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "OAUTH2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.scopes.#", "2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.max_expiry_duration_in_hours", "11"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_intermediate_steps", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_session", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_pkce", "true"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.fallback_redirect_path", "/fallbackRedirectPath2"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_type", "CODE"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.0.type", "VALIDATION_BLOCK"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.0.type", "VALIDATION_BLOCK"),

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
			Config: config + imageVariableStr + clientSecretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployments", "test_deployments", acctest.Optional, acctest.Update, ApigatewayDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuthWithTokenHeader),
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
			Config: config + imageVariableStr + clientSecretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, ApigatewayDeploymentSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_dynamic_auth", acctest.Required, acctest.Create, ApigatewayDeploymentWithDynamicAuthenticationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_with_oidc", acctest.Required, acctest.Create, ApigatewayDeploymentWithOidcSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayDeploymentResourceConfig,
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
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.authentication.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_methods.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.exposed_headers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.is_allow_credentials_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.cors.0.max_age_in_seconds", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.dynamic_authentication.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.0.allowed_sans.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.mutual_tls.0.is_verified_certificate_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_in_requests_per_second", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.rate_limiting.0.rate_key", "TOTAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.usage_plans.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.request_policies.0.usage_plans.0.token_locations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.name", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.backend.0.routing_backends.0.key.0.is_default", "true"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.header_transformations.0.set_headers.0.items.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.time_to_live_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.routes.0.response_policies.0.response_cache_store.0.type", "FIXED_TTL_STORE_POLICY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithDynamicAuth, "endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithDynamicAuth, "gateway_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithDynamicAuth, "id"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "path_prefix", "/v3"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.function_id"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.token_header", "tokenHeader2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.authentication_server_detail.0.type", "CUSTOM_AUTHENTICATION"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.is_default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.type", "ANY_OF"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.authentication_servers.0.key.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.selector", "request.subdomain[oracle.com]"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithDynamicAuth, "specification.0.request_policies.0.dynamic_authentication.0.selection_source.0.type", "SINGLE"),

				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithOidc, "endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithOidc, "gateway_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithOidc, "id"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "path_prefix", "/v4"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.type", "TOKEN_AUTHENTICATION"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_auth_scheme", "Bearer"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.token_header", "Authorization"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.is_anonymous_access_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.type", "REMOTE_DISCOVERY"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.is_ssl_verify_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.max_cache_duration_in_hours", "11"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.type", "CUSTOM"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_id", "example_client_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.client_details.0.client_secret_version_number", "2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.type", "DISCOVERY_URI"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.source_uri_details.0.uri", "https://oracle.com/discovery2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.audiences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.issuers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.is_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_policy.0.additional_validation_policy.0.verify_claims.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.type", "OAUTH2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.scopes.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.max_expiry_duration_in_hours", "11"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_intermediate_steps", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_cookies_for_session", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.use_pkce", "true"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.fallback_redirect_path", "/fallbackRedirectPath2"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.response_type", "CODE"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.client_details.0.type", "VALIDATION_BLOCK"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameWithOidc, "specification.0.request_policies.0.authentication.0.validation_failure_policy.0.source_uri_details.0.type", "VALIDATION_BLOCK"),
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
