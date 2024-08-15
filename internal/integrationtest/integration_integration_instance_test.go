// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_integration "github.com/oracle/oci-go-sdk/v65/integration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

type IdcsAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

var (
	IntegrationIntegrationInstanceRequiredOnlyResource = IntegrationIntegrationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Required, acctest.Create, integrationInstanceRepresentation)

	IntegrationIntegrationInstanceResourceConfig = IntegrationIntegrationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation)

	IntegrationintegrationInstanceSingularDataSourceRepresentation = map[string]interface{}{
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

	IntegrationIntegrationInstanceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"integration_instance_type":  acctest.Representation{RepType: acctest.Required, Create: `${var.instance_type}`},
		"is_byol":                    acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"message_packs":              acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"alternate_custom_endpoints": acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model":          acctest.Representation{RepType: acctest.Optional, Create: `UCM`},
		// STANDARD or ENTERPRISE only
		// "custom_endpoint":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceCustomEndpointRepresentation},
		// "defined_tags":                      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		// "idcs_at":                   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.domain_id}`},
		"is_file_server_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// STANDARD or ENTERPRISE only
		// "network_endpoint_details":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsRepresentation},
		"shape":                             acctest.Representation{RepType: acctest.Optional, Create: `DEVELOPMENT`},
		"enable_process_automation_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"extend_data_retention_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	integrationInstanceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"integration_instance_type": acctest.Representation{RepType: acctest.Required, Create: `${var.instance_type}`},
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `DEVELOPMENT`},
		"is_byol":                   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"message_packs":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// Not supported yet
		// "alternate_custom_endpoints": acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model": acctest.Representation{RepType: acctest.Optional, Create: `UCM`}, // STANDARD or ENTERPRISE only
		// "custom_endpoint":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceCustomEndpointRepresentation},
		// "defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},

		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		// "idcs_at":                   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"domain_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.domain_id}`},
		"is_file_server_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// STANDARD or ENTERPRISE only
		// "network_endpoint_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentationAgain},
	}

	integrationInstanceRepresentationIdcsAt = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"integration_instance_type": acctest.Representation{RepType: acctest.Required, Create: `${var.instance_type}`},
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `DEVELOPMENT`},
		"is_byol":                   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"message_packs":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// Not supported yet
		// "alternate_custom_endpoints": acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model": acctest.Representation{RepType: acctest.Optional, Create: `UCM`}, // STANDARD or ENTERPRISE only
		// "custom_endpoint":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceCustomEndpointRepresentation},
		// "defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},

		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_at":                   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"is_file_server_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// STANDARD or ENTERPRISE only
		// "network_endpoint_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentationAgain},
		// "network_endpoint_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: integrationInstanceNetworkEndpointDetailsRepresentation},
	}
	integrationPrivateEndpointRepresentation = map[string]interface{}{
		"integration_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_integration_integration_instance.test_integration_instance.id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"nsg_ids":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}},
	}

	ignoreDefinedTagsDifferencesRepresentationAgain = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
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

	integrationInstanceNetworkEndpointDetailsAllowlistedHttpVcnsRepresentation = map[string]interface{}{
		"id":              acctest.Representation{RepType: acctest.Required, Create: `${var.allow_listed_http_vcn}`},
		"allowlisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{`172.16.0.239/32`}},
	}

	IntegrationIntegrationInstanceResourceDependencies = "" /* DefinedTagsDependencies + KmsVaultIdVariableStr + integrationInstanceVcnRepresentation*/
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

	instanceType := utils.GetEnvSettingWithBlankDefault("instance_type")
	instanceTypeVariableStr := fmt.Sprintf("variable \"instance_type\" { default = \"%s\" }\n", instanceType)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("oci_vault_secret_id")
	vaultSecretIdStr := fmt.Sprintf("variable \"oci_vault_secret_id\" { default = \"%s\" }\n", vaultSecretId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_id")
	nsgIdStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	domainIdVariable := utils.GetEnvSettingWithBlankDefault("domain_id")
	domainIdVariableStr := fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainIdVariable)

	resourceName := "oci_integration_integration_instance.test_integration_instance"
	datasourceName := "data.oci_integration_integration_instances.test_integration_instances"
	singularDatasourceName := "data.oci_integration_integration_instance.test_integration_instance"
	privateEndpointResourceName := "oci_integration_private_endpoint_outbound_connection.integration_private_endpoint"
	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+instanceTypeVariableStr+compartmentIdVariableStr+IntegrationIntegrationInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Create, integrationInstanceRepresentation), "integration", "integrationInstance", t)

	acctest.ResourceTest(t, testAccCheckIntegrationIntegrationInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + instanceTypeVariableStr + compartmentIdVariableStr + subnetIdStr + domainIdVariableStr + nsgIdStr + idcsAccessTokenVariableStr() + IntegrationIntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Required, acctest.Create, integrationInstanceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_private_endpoint_outbound_connection", "integration_private_endpoint", acctest.Required, acctest.Create, integrationPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "data_retention_period"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},

				resource.TestCheckResourceAttrSet(privateEndpointResourceName, "integration_instance_id"),
				resource.TestCheckResourceAttr(privateEndpointResourceName, "subnet_id", utils.GetEnvSettingWithBlankDefault("subnet_id")),
				resource.TestCheckResourceAttr(privateEndpointResourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(privateEndpointResourceName, "nsg_ids.0", utils.GetEnvSettingWithBlankDefault("nsg_id")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, privateEndpointResourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + instanceTypeVariableStr + compartmentIdVariableStr +
				tagVariablesStr() +
				domainIdVariableStr +
				vaultSecretIdStr +
				IntegrationIntegrationInstanceResourceDependencies +
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
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"), // 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "data_retention_period"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "0"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				// acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
				// 	"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				// },
				// 	[]string{}),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "shape", "DEVELOPMENT"),
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
			Config: config + instanceTypeVariableStr + compartmentIdVariableStr +
				tagVariablesStr() +
				domainIdVariableStr +
				compartmentIdUVariableStr +
				vaultSecretIdStr +
				IntegrationIntegrationInstanceResourceDependencies +
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
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_retention_period"),
				// resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "0"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				// acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
				// 	"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				// },
				// 	[]string{}),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "shape", "DEVELOPMENT"),

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
			Config: config + instanceTypeVariableStr + compartmentIdVariableStr + tagVariablesStr() + domainIdVariableStr + vaultSecretIdStr + IntegrationIntegrationInstanceResourceDependencies +
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
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "2"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "0"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				// acctest.CheckResourceSetContainsElementWithProperties(resourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
				// 	"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				// },
				// 	[]string{}),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "shape", "DEVELOPMENT"),
				resource.TestCheckResourceAttr(resourceName, "shape", "DEVELOPMENT"),
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
			Config: config + instanceTypeVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_integration_integration_instances", "test_integration_instances", acctest.Optional, acctest.Update, integrationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + tagVariablesStr() + domainIdVariableStr + vaultSecretIdStr + IntegrationIntegrationInstanceResourceDependencies +
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
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.data_retention_period"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.custom_endpoint.0.alias"),
				// resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.custom_endpoint.0.certificate_secret_id"),
				// resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.custom_endpoint.0.certificate_secret_version"),
				// resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.instance_url"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_byol", "true"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.message_packs", "2"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.#", "0"),
				// resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				// acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "integration_instances.0.network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
				// 	"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				// },
				// 	[]string{}),
				// resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(datasourceName, "integration_instances.0.shape", "DEVELOPMENT"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "integration_instances.0.time_updated"),
			),
		},

		// verify singular datasource
		{
			Config: config + instanceTypeVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, IntegrationintegrationInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + tagVariablesStr() + domainIdVariableStr + vaultSecretIdStr + IntegrationIntegrationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Update, integrationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "integration_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attachments.#", "0"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.#", "1"),
				// CheckResourceSetContainsElementWithProperties(singularDatasourceName, "alternate_custom_endpoints", map[string]string{
				// 	"hostname": "hostname2",
				// },
				// 	[]string{
				// 		"certificate_secret_version",
				// 	}),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_retention_period"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.#", "0"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_endpoint.0.alias"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_endpoint.0.certificate_secret_version"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.0.hostname", "hostname2-updated.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_byol", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_file_server_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message_packs", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.#", "0"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns.#", "1"),
				// acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns", map[string]string{
				// 	"id": utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn"),
				// },
				// 	[]string{}),
				// resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.allowlisted_http_vcns.0.id",
				// 	utils.GetEnvSettingWithBlankDefault("allow_listed_http_vcn")),
				// resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.is_integration_vcn_allowlisted", "false"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.network_endpoint_type", "PUBLIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "DEVELOPMENT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + IntegrationIntegrationInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"domain_id",
				"idcs_at",
			},
			ResourceName: resourceName,
		},
	})
}

func TestIntegrationIntegrationInstanceResource_idcsAt(t *testing.T) {
	httpreplay.SetScenario("TestIntegrationIntegrationInstanceResource_idcsAt")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestIntegrationIntegrationInstanceResource_idcsAt") {
		t.Skip("Skipping suppressed TestIntegrationIntegrationInstanceResource_idcsAt")
	}

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	instanceType := utils.GetEnvSettingWithBlankDefault("instance_type")
	instanceTypeVariableStr := fmt.Sprintf("variable \"instance_type\" { default = \"%s\" }\n", instanceType)

	resourceName := "oci_integration_integration_instance.test_integration_instance"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+instanceTypeVariableStr+compartmentIdVariableStr+IntegrationIntegrationInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Optional, acctest.Create, integrationInstanceRepresentationIdcsAt), "integration", "integrationInstance", t)

	acctest.ResourceTest(t, testAccCheckIntegrationIntegrationInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + instanceTypeVariableStr + compartmentIdVariableStr + idcsAccessTokenVariableStr() +
				acctest.GenerateResourceFromRepresentationMap("oci_integration_integration_instance", "test_integration_instance", acctest.Required, acctest.Create, integrationInstanceRepresentationIdcsAt),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "integration_instance_type", utils.GetEnvSettingWithBlankDefault("instance_type")),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_packs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
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
	integrationInstanceIds, err := getIntegrationIntegrationInstanceIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &integrationInstanceId, IntegrationintegrationInstancesSweepWaitCondition, time.Duration(3*time.Minute),
				IntegrationintegrationInstancesSweepResponseFetchOperation, "integration", true)
		}
	}
	return nil
}

func getIntegrationIntegrationInstanceIds(compartment string) ([]string, error) {
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

func IntegrationintegrationInstancesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if integrationInstanceResponse, ok := response.Response.(oci_integration.GetIntegrationInstanceResponse); ok {
		return integrationInstanceResponse.LifecycleState != oci_integration.IntegrationInstanceLifecycleStateDeleted
	}
	return false
}

func IntegrationintegrationInstancesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
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
	// Generate new IDCS token each time calling the creation of instance to avoid expired token
	endpoint := utils.GetEnvSettingWithBlankDefault("idcs_auth_endpoint")

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("scope", "urn:opc:idm:__myscopes__")
	data.Set("username", utils.GetEnvSettingWithBlankDefault("idcs_username"))
	data.Set("password", utils.GetEnvSettingWithBlankDefault("idcs_password"))

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	req.SetBasicAuth(
		utils.GetEnvSettingWithBlankDefault("idcs_auth_username"),
		utils.GetEnvSettingWithBlankDefault("idcs_auth_password"))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var result = new(IdcsAccessToken)

	err = json.Unmarshal(body, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.AccessToken)
	return result.AccessToken
}
