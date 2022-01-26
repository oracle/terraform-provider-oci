// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_integration "github.com/oracle/oci-go-sdk/v56/integration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

type IdcsAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

var (
	IntegrationInstanceRequiredOnlyResource = IntegrationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Required, acctest.Create, integrationInstanceRepresentation)

	IntegrationInstanceResourceConfig = IntegrationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation)

	integrationInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"integration_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_integration_integration_instance.test_integration_instance.id}`},
	}

	integrationInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: integrationInstanceDataSourceFilterRepresentation},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
	integrationInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_integration_integration_instance.test_integration_instance.id}`}},
	}

	integrationInstanceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"integration_instance_type": acctest.Representation{RepType: acctest.Required, Create: `STANDARD`, Update: `ENTERPRISE`},
		"is_byol":                   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"message_packs":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// Not supported yet
		// "alternate_custom_endpoints": acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model":         acctest.Representation{RepType: acctest.Optional, Create: `UCM`},
		"custom_endpoint":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceCustomEndpointRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${var.oci_identity_tag_namespace}.${var.oci_identity_tag}", "value")}`, Update: `${map("${var.oci_identity_tag_namespace}.${var.oci_identity_tag}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_at":                   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"is_file_server_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"network_endpoint_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsRepresentation},
	}
	integrationInstanceAlternateCustomEndpointsRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `althostname.com`, Update: `althostname2.com`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.oci_vault_secret_id}`},
	}
	integrationInstanceCustomEndpointRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `hostname2.com`, Update: `hostname2-updated.com`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.oci_vault_secret_id}`},
	}
	integrationInstanceNetworkEndpointDetailsRepresentation = map[string]interface{}{
		"network_endpoint_type":          acctest.Representation{RepType: acctest.Required, Create: `PUBLIC`},
		"allowlisted_http_ips":           acctest.Representation{RepType: acctest.Optional, Create: []string{`172.16.0.239/32`}},
		"allowlisted_http_vcns":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsAllowlistedHttpVcnsRepresentation},
		"is_integration_vcn_allowlisted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	integrationInstanceVcnRepresentation = `resource "oci_core_vcn" "vcn" {
cidr_blocks    = ["10.0.0.0/16"]
dns_label      = "vcn"
compartment_id = var.compartment_id
display_name   = "vcn"
}`

	integrationInstanceNetworkEndpointDetailsAllowlistedHttpVcnsRepresentation = map[string]interface{}{
		"id":              acctest.Representation{RepType: acctest.Required, Create: `oci_core_vcn.vcn.id`},
		"allowlisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{`172.16.0.239/32`}},
	}

	IntegrationInstanceResourceDependencies = DefinedTagsDependencies + KmsVaultIdVariableStr + integrationInstanceVcnRepresentation
)

// issue-routing-tag: integration/default
func TestIntegrationIntegrationInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIntegrationIntegrationInstanceResource_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestIntegrationIntegrationInstanceResource_basic") {
		t.Skip("Skipping suppressed TestIntegrationIntegrationInstanceResource_basic")
	}

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("oci_vault_secret_id")
	vaultSecretIdStr := fmt.Sprintf("variable \"oci_vault_secret_id\" { default = \"%s\" }\n", vaultSecretId)

	resourceName := "oci_integration_integration_instance.test_integration_instance"
	datasourceName := "data.oci_integration_integration_instances.test_integration_instances"
	singularDatasourceName := "data.oci_integration_integration_instance.test_integration_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IntegrationInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Create, integrationInstanceRepresentation), "integration", "integrationInstance", t)

	acctest.ResourceTest(t, testAccCheckIntegrationIntegrationInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr() + IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Required, acctest.Create, integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr(),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				tagVariablesStr() +
				idcsAccessTokenVariableStr() +
				vaultSecretIdStr +
				IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_integration_integration_instance",
					"test_integration_instance",
					acctest.Optional,
					acctest.Create,
					integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(resourceName, "alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname",
				// },
				// 	[]string{
				// 		"certificate_secret_id",
				// 	}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
					"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

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
			Config: config + compartmentIdVariableStr +
				tagVariablesStr() +
				idcsAccessTokenVariableStr() +
				compartmentIdUVariableStr +
				vaultSecretIdStr +
				IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_integration_integration_instance",
					"test_integration_instance",
					acctest.Optional,
					acctest.Create,
					acctest.RepresentationCopyWithNewProperties(integrationInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(resourceName, "alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname",
				// },
				// 	[]string{
				// 		"certificate_secret_id",
				// 	}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
					"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),

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
			Config: config + compartmentIdVariableStr + tagVariablesStr() + idcsAccessTokenVariableStr() + vaultSecretIdStr + IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(resourceName, "alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname2",
				// },
				// 	[]string{
				// 		"certificate_secret_id",
				// 	}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", "ENTERPRISE"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "2"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
					"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_integration_integration_instances", "test_integration_instances", acctest.Optional, acctest.Update, integrationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + tagVariablesStr() + idcsAccessTokenVariableStr() + vaultSecretIdStr + IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "integration_instances.#", "1"),
				// resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(datasourceName, "integration_instances.0.alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname2",
				// },
				// 	[]string{
				// 		"certificate_secret_id",
				// 		"certificate_secret_version",
				// 	}),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.consumption_model", "UCM"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.custom_endpoint.0.certificate_secret_version"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.instance_url"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.integration_instance_type", "ENTERPRISE"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_byol", "true"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_file_server_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.message_packs", "2"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "integration_instances.0.network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
					"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				},
					[]string{}),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + tagVariablesStr() + idcsAccessTokenVariableStr() + vaultSecretIdStr + IntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(singularDatasourceName, "alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname2",
				// },
				// 	[]string{
				// 		"certificate_secret_version",
				// 	}),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_endpoint.0.certificate_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "integration_instance_type", "ENTERPRISE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_byol", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_file_server_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message_packs", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
					"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns.0.id",
					utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn")),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + tagVariablesStr() + idcsAccessTokenVariableStr() + vaultSecretIdStr + IntegrationInstanceResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"idcs_at",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIntegrationIntegrationInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IntegrationInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_integration_integration_instance" {
			noResourceFound = false
			request := oci_integration.GetIntegrationInstanceRequest{}

			tmp := rs.Primary.ID
			request.IntegrationInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "integration")

			response, err := client.GetIntegrationInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_integration.IntegrationInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IntegrationIntegrationInstance") {
		resource.AddTestSweepers("IntegrationIntegrationInstance", &resource.Sweeper{
			Name:         "IntegrationIntegrationInstance",
			Dependencies: acctest.DependencyGraph["integrationInstance"],
			F:            sweepIntegrationIntegrationInstanceResource,
		})
	}
}

func sweepIntegrationIntegrationInstanceResource(compartment string) error {
	integrationInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).IntegrationInstanceClient()
	integrationInstanceIds, err := getIntegrationInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, integrationInstanceId := range integrationInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[integrationInstanceId]; !ok {
			deleteIntegrationInstanceRequest := oci_integration.DeleteIntegrationInstanceRequest{}

			deleteIntegrationInstanceRequest.IntegrationInstanceId = &integrationInstanceId

			deleteIntegrationInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "integration")
			_, error := integrationInstanceClient.DeleteIntegrationInstance(context.Background(), deleteIntegrationInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting IntegrationInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", integrationInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &integrationInstanceId, integrationInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				integrationInstanceSweepResponseFetchOperation, "integration", true)
		}
	}
	return nil
}

func getIntegrationInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IntegrationInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	integrationInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).IntegrationInstanceClient()

	listIntegrationInstancesRequest := oci_integration.ListIntegrationInstancesRequest{}
	listIntegrationInstancesRequest.CompartmentId = &compartmentId
	listIntegrationInstancesRequest.LifecycleState = oci_integration.ListIntegrationInstancesLifecycleStateActive
	listIntegrationInstancesResponse, err := integrationInstanceClient.ListIntegrationInstances(context.Background(), listIntegrationInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IntegrationInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, integrationInstance := range listIntegrationInstancesResponse.Items {
		id := *integrationInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IntegrationInstanceId", id)
	}
	return resourceIds, nil
}

func integrationInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if integrationInstanceResponse, ok := response.Response.(oci_integration.GetIntegrationInstanceResponse); ok {
		return integrationInstanceResponse.LifecycleState != oci_integration.IntegrationInstanceLifecycleStateDeleted
	}
	return false
}

func integrationInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IntegrationInstanceClient().GetIntegrationInstance(context.Background(), oci_integration.GetIntegrationInstanceRequest{
		IntegrationInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func idcsAccessTokenVariableStr() string {
	return fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", readIdcsAccessToken())
}

func tagVariablesStr() string {
	vars := []string{
		fmt.Sprintf("variable \"oci_identity_tag_namespace\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("oci_identity_tag_namespace")),
		fmt.Sprintf("variable \"oci_identity_tag\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("oci_identity_tag")),
		fmt.Sprintf("variable \"allow_listed_http_vcn\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn")),
	}

	return strings.Join(vars[:], "\n")
}

func readIdcsAccessToken() string {
	return utils.GetEnvSettingWithBlankDefault("idcs_access_token")
}
