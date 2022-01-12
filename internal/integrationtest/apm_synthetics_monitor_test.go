// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MonitorRequiredOnlyResource = MonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, monitorRepresentation)

	MonitorResourceConfig = MonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, monitorRepresentation)

	monitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    Representation{RepType: Required, Create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	monitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `SCRIPTED_BROWSER`},
		"script_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"vantage_point": acctest.Representation{RepType: acctest.Optional, Create: `vantagePoint`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: monitorDataSourceFilterRepresentation}}
	monitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_monitor.test_monitor.id}`}},
	}

	monitorRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":               acctest.Representation{RepType: acctest.Required, Create: `SCRIPTED_BROWSER`},
		"repeat_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"vantage_points":             acctest.RepresentationGroup{RepType: acctest.Required, Group: monitorVantagePointsRepresentation},
		"configuration":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"script_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_apm_synthetics_script.test_script.id}`},
		"script_parameters":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorScriptParametersRepresentation},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                     acctest.Representation{RepType: acctest.Optional, Create: `target`, Update: `target2`},
		"timeout_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	monitorVantagePointsRepresentation = map[string]interface{}{}
	monitorConfigurationRepresentation = map[string]interface{}{
		"config_type":                       acctest.Representation{RepType: acctest.Optional, Create: `BROWSER_CONFIG`, Update: `SCRIPTED_BROWSER_CONFIG`},
		"is_certificate_validation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_failure_retried":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_redirection_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"network_configuration":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationNetworkConfigurationRepresentation},
		"req_authentication_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationReqAuthenticationDetailsRepresentation},
		"req_authentication_scheme":         acctest.Representation{RepType: acctest.Optional, Create: `OAUTH`, Update: `NONE`},
		"request_headers":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationRequestHeadersRepresentation},
		"request_method":                    acctest.Representation{RepType: acctest.Optional, Create: `GET`, Update: `POST`},
		"request_post_body":                 acctest.Representation{RepType: acctest.Optional, Create: `requestPostBody`, Update: `requestPostBody2`},
		"request_query_params":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationRequestQueryParamsRepresentation},
		"verify_response_codes":             acctest.Representation{RepType: acctest.Optional, Create: []string{`verifyResponseCodes`}, Update: []string{`verifyResponseCodes2`}},
		"verify_response_content":           acctest.Representation{RepType: acctest.Optional, Create: `verifyResponseContent`, Update: `verifyResponseContent2`},
		"verify_texts":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationVerifyTextsRepresentation},
	}
	monitorScriptParametersRepresentation = map[string]interface{}{
		"param_name":  acctest.Representation{RepType: acctest.Required, Create: `paramName`, Update: `paramName2`},
		"param_value": acctest.Representation{RepType: acctest.Required, Create: `paramValue`, Update: `paramValue2`},
	}
	monitorConfigurationNetworkConfigurationRepresentation = map[string]interface{}{
		"number_of_hops":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"probe_mode":        acctest.Representation{RepType: acctest.Optional, Create: `SACK`, Update: `SYN`},
		"probe_per_hop":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"protocol":          acctest.Representation{RepType: acctest.Optional, Create: `ICMP`, Update: `TCP`},
		"transmission_rate": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	monitorConfigurationReqAuthenticationDetailsRepresentation = map[string]interface{}{
		"auth_headers":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation},
		"auth_request_method":    acctest.Representation{RepType: acctest.Optional, Create: `GET`, Update: `POST`},
		"auth_request_post_body": acctest.Representation{RepType: acctest.Optional, Create: `authRequestPostBody`, Update: `authRequestPostBody2`},
		"auth_token":             acctest.Representation{RepType: acctest.Optional, Create: `authToken`, Update: `authToken2`},
		"auth_url":               acctest.Representation{RepType: acctest.Optional, Create: `authUrl`, Update: `authUrl2`},
		"auth_user_name":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
		"auth_user_password":     acctest.Representation{RepType: acctest.Optional, Create: `authUserPassword`, Update: `authUserPassword2`},
		"oauth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `BASIC`},
	}
	monitorConfigurationRequestHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `headerName`, Update: `headerName2`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `headerValue`, Update: `headerValue2`},
	}
	monitorConfigurationRequestQueryParamsRepresentation = map[string]interface{}{
		"param_name":  acctest.Representation{RepType: acctest.Optional, Create: `paramName`, Update: `paramName2`},
		"param_value": acctest.Representation{RepType: acctest.Optional, Create: `paramValue`, Update: `paramValue2`},
	}
	monitorConfigurationVerifyTextsRepresentation = map[string]interface{}{
		"text": acctest.Representation{RepType: acctest.Optional, Create: `text`, Update: `text2`},
	}
	monitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `headerName`, Update: `headerName2`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `headerValue`, Update: `headerValue2`},
	}

	MonitorResourceDependencies = GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
		GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, scriptRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, monitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, monitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, monitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SACK"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "ICMP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "headerName"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "headerValue"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_token", "authToken"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "authUrl"),
				resource.TestCheckResourceAttrSet(resourceName, "configuration.0.req_authentication_details.0.auth_user_name"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_user_password", "authUserPassword"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "headerName"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "headerValue"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_texts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_texts.0.text", "text"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "paramName"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "paramValue"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "target"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

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
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, monitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "headerName2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "headerValue2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_token", "authToken2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "authUrl2"),
				resource.TestCheckResourceAttrSet(resourceName, "configuration.0.req_authentication_details.0.auth_user_name"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_user_password", "authUserPassword2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "headerName2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "headerValue2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_texts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_texts.0.text", "text2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "true"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "target2"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, monitorDataSourceRepresentation) +
				compartmentIdVariableStr + MonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, monitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "monitorType"),
				resource.TestCheckResourceAttrSet(datasourceName, "script_id"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(datasourceName, "vantage_point", "vantagePoint"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, monitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_per_hop", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "headerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "headerValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_token", "authToken2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_url", "authUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_user_password", "authUserPassword2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "BASIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_scheme", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_name", "headerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_value", "headerValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_codes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_texts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_texts.0.text", "text2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "target2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + MonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmSyntheticsMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_monitor" {
			noResourceFound = false
			request := oci_apm_synthetics.GetMonitorRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")

			_, err := client.GetMonitor(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
		initDependencyGraph()
	}
	if !InSweeperExcludeList("ApmSyntheticsMonitor") {
		resource.AddTestSweepers("ApmSyntheticsMonitor", &resource.Sweeper{
			Name:         "ApmSyntheticsMonitor",
			Dependencies: DependencyGraph["monitor"],
			F:            sweepApmSyntheticsMonitorResource,
		})
	}
}

func sweepApmSyntheticsMonitorResource(compartment string) error {
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()
	monitorIds, err := getMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, monitorId := range monitorIds {
		if ok := SweeperDefaultResourceId[monitorId]; !ok {
			deleteMonitorRequest := oci_apm_synthetics.DeleteMonitorRequest{}

			deleteMonitorRequest.MonitorId = &monitorId

			deleteMonitorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteMonitor(context.Background(), deleteMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting Monitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitorId, error)
				continue
			}
		}
	}
	return nil
}

func getMonitorIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "MonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()

	listMonitorsRequest := oci_apm_synthetics.ListMonitorsRequest{}
	listMonitorsRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Monitor resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listMonitorsRequest.ApmDomainId = &apmDomainId

		listMonitorsResponse, err := apmSyntheticClient.ListMonitors(context.Background(), listMonitorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Monitor list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, monitor := range listMonitorsResponse.Items {
			id := *monitor.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitorId", id)
		}

	}
	return resourceIds, nil
}
