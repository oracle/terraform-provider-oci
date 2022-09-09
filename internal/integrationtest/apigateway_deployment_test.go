// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuth)

	ApigatewayDeploymentResourceConfig = DeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuth)

	ApigatewayDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment.id}`},
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
	ApigatewayDeploymentSpecificationLoggingPoliciesRepresentation = map[string]interface{}{
		"access_log":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesAccessLogRepresentation},
		"execution_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationLoggingPoliciesExecutionLogRepresentation},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesRepresentation = map[string]interface{}{
		"authentication": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationRepresentation},
		"cors":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesCorsRepresentation},
		"mutual_tls":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesMutualTlsRepresentation},
		"rate_limiting":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesRateLimitingRepresentation},
		"usage_plans":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesUsagePlansRepresentation},
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
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthorizeScopeRepresentation = map[string]interface{}{
		"allowed_scope": acctest.Representation{RepType: acctest.Optional, Create: []string{`cors`}},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesCorsRepresentation = map[string]interface{}{
		"allowed_origins":              acctest.Representation{RepType: acctest.Required, Create: []string{`https://www.oracle.org`}, Update: []string{`*`}},
		"allowed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"allowed_methods":              acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"exposed_headers":              acctest.Representation{RepType: acctest.Optional, Create: []string{`*`}, Update: []string{`*`, `Content-Type`}},
		"is_allow_credentials_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds":           acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `500`},
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
		"type":                            acctest.Representation{RepType: acctest.Required, Create: `MODIFY_RESPONSE`},
		"response_code":                   acctest.Representation{RepType: acctest.Optional, Create: `210`, Update: `220`},
		"response_header_transformations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRepresentation},
		"response_message":                acctest.Representation{RepType: acctest.Optional, Create: `responseMessage`, Update: `responseMessage2`},
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
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRepresentation = map[string]interface{}{
		"filter_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersRepresentation},
		"rename_headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersRepresentation},
		"set_headers":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersRepresentation},
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
		"name":      acctest.Representation{RepType: acctest.Required, Create: `nameE`, Update: `nameE2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `SKIP`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsFilterHeadersItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `namefilter`, Update: `namefilter2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsRenameHeadersItemsRepresentation = map[string]interface{}{
		"from": acctest.Representation{RepType: acctest.Required, Create: `from`, Update: `from2`},
		"to":   acctest.Representation{RepType: acctest.Required, Create: `to`, Update: `to2`},
	}
	ApigatewayDeploymentSpecificationRequestPoliciesAuthenticationValidationFailurePolicyResponseHeaderTransformationsSetHeadersItemsRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"values":    acctest.Representation{RepType: acctest.Required, Create: []string{`values`}, Update: []string{`values2`}},
		"if_exists": acctest.Representation{RepType: acctest.Optional, Create: `OVERWRITE`, Update: `APPEND`},
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

	deploymentRepresentationCustomAuth = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication.audiences",
		"specification.request_policies.authentication.issuers",
		"specification.request_policies.authentication.max_clock_skew_in_seconds",
		"specification.request_policies.authentication.public_keys",
		"specification.request_policies.authentication.token_auth_scheme",
		"specification.request_policies.authentication.verify_claims",
		"specification.request_policies.authentication.validation_failure_policy",
		"specification.request_policies.authentication.parameters",
		"specification.request_policies.authentication.cache_key",
	}, ApigatewayDeploymentRepresentation)

	deploymentRepresentationWithoutTokenHeaderCustomAuth = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
		"specification.request_policies.authentication.audiences",
		"specification.request_policies.authentication.issuers",
		"specification.request_policies.authentication.max_clock_skew_in_seconds",
		"specification.request_policies.authentication.public_keys",
		"specification.request_policies.authentication.token_auth_scheme",
		"specification.request_policies.authentication.verify_claims",
		"specification.request_policies.authentication.token_header",
	}, ApigatewayDeploymentRepresentation)

	deploymentRepresentationRequestBasedAuthCustomAuth = acctest.GetUpdatedRepresentationCopy(
		"path_prefix",
		acctest.Representation{RepType: acctest.Required, Create: `/v2`},
		deploymentRepresentationWithoutTokenHeaderCustomAuth)
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
	resourceNameRba := "oci_apigateway_deployment.test_deployment_rba"
	datasourceName := "data.oci_apigateway_deployments.test_deployments"
	singularDatasourceName := "data.oci_apigateway_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+imageVariableStr+DeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationCustomAuth)+acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth), "apigateway", "deployment", t)

	acctest.ResourceTest(t, testAccCheckApigatewayDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentationCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Required, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "path_prefix", "/v1"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),

				resource.TestCheckResourceAttr(resourceNameRba, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRba, "gateway_id"),
				resource.TestCheckResourceAttr(resourceNameRba, "path_prefix", "/v2"),
				resource.TestCheckResourceAttr(resourceNameRba, "specification.#", "1"),

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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create, deploymentRepresentationCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create, deploymentRepresentationRequestBasedAuthCustomAuth),
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationCustomAuth, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"specification.request_policies.authentication.parameters": acctest.Representation{RepType: acctest.Optional, Create: nil},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentRepresentationRequestBasedAuthCustomAuth, map[string]interface{}{
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Optional, acctest.Update, deploymentRepresentationCustomAuth) +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment_rba", acctest.Optional, acctest.Update, deploymentRepresentationRequestBasedAuthCustomAuth),
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployments", "test_deployments", acctest.Optional, acctest.Update, ApigatewayDeploymentDataSourceRepresentation) +
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
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, ApigatewayDeploymentSingularDataSourceRepresentation) +
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
