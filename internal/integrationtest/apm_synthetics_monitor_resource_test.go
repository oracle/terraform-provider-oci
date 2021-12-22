// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RestMonitorResourceConfig = MonitorResourceDependencies1 +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, restMonitorRepresentation)

	restMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	restMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `REST`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: restMonitorDataSourceFilterRepresentation},
	}
	restMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_monitor.test_monitor.display_name}`}},
	}

	restMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":               acctest.Representation{RepType: acctest.Required, Create: `REST`},
		"repeat_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":             acctest.Representation{RepType: acctest.Required, Create: []string{`OraclePublic-us-ashburn-1`}},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                     acctest.Representation{RepType: acctest.Optional, Create: `https://console.us-ashburn-1.oraclecloud.com`, Update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: restMonitorConfigurationRepresentation},
	}

	restMonitorConfigurationRepresentation = map[string]interface{}{
		"config_type":                       acctest.Representation{RepType: acctest.Optional, Create: `REST_CONFIG`},
		"is_certificate_validation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_failure_retried":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_redirection_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"req_authentication_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationReqAuthenticationDetailsRepresentation},
		"req_authentication_scheme":         acctest.Representation{RepType: acctest.Optional, Create: `OAUTH`},
		"request_headers":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationRequestHeadersRepresentation},
		"request_method":                    acctest.Representation{RepType: acctest.Optional, Create: `POST`},
		"request_post_body":                 acctest.Representation{RepType: acctest.Optional, Create: `requestPostBody`, Update: `requestPostBody2`},
		"request_query_params":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationRequestQueryParamsRepresentation},
		"verify_response_codes":             acctest.Representation{RepType: acctest.Optional, Create: []string{`200`, `300`, `400`}},
		"verify_response_content":           acctest.Representation{RepType: acctest.Optional, Create: `verifyResponseContent`, Update: `verifyResponseContent2`},
	}

	monitorConfigurationReqAuthenticationDetailsRepresentation = map[string]interface{}{
		"auth_headers":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation},
		"auth_request_method":    acctest.Representation{RepType: acctest.Optional, Create: `POST`},
		"auth_request_post_body": acctest.Representation{RepType: acctest.Optional, Create: `authRequestPostBody`, Update: `authRequestPostBody2`},
		"auth_url":               acctest.Representation{RepType: acctest.Optional, Create: `http://authUrl`, Update: `http://authUrl2`},
		"oauth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}
	monitorConfigurationRequestHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `content-type`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `json`},
	}
	monitorConfigurationRequestQueryParamsRepresentation = map[string]interface{}{
		"param_name":  acctest.Representation{RepType: acctest.Optional, Create: `paramName`, Update: `paramName2`},
		"param_value": acctest.Representation{RepType: acctest.Optional, Create: `paramValue`, Update: `paramValue2`},
	}

	monitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `content-type`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `json`},
	}

	MonitorResourceDependencies1 = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsMonitorResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitorResourceDependencies1+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, restMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies1 +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, restMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies1 +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, restMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, restMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + MonitorResourceDependencies1 +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, restMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, restMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RestMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + RestMonitorResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
			},
			ResourceName: resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
}
