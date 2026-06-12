// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiHostedApplicationRequiredOnlyResource = GenerativeAiHostedApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Required, acctest.Create, GenerativeAiHostedApplicationRepresentation)

	GenerativeAiHostedApplicationResourceConfig = GenerativeAiHostedApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationRepresentation)

	GenerativeAiHostedApplicationSingularDataSourceRepresentation = map[string]interface{}{
		"hosted_application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_hosted_application.test_hosted_application.id}`},
	}

	GenerativeAiHostedApplicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_hosted_application.test_hosted_application.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationDataSourceFilterRepresentation}}
	GenerativeAiHostedApplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_hosted_application.test_hosted_application.id}`}},
	}

	GenerativeAiHostedApplicationRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"environment_variables": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationEnvironmentVariablesRepresentation},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"inbound_auth_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationInboundAuthConfigRepresentation},
		"networking_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationNetworkingConfigRepresentation},
		"scaling_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationScalingConfigRepresentation},
		"storage_configs":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationStorageConfigsRepresentation},
	}
	GenerativeAiHostedApplicationEnvironmentVariablesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `PLAINTEXT`, Update: `VAULT`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `{\"dummyKey\": \"dummyValue\"}`},
	}
	GenerativeAiHostedApplicationInboundAuthConfigRepresentation = map[string]interface{}{
		"inbound_auth_config_type": acctest.Representation{RepType: acctest.Required, Create: `IDCS_AUTH_CONFIG`},
		"idcs_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationInboundAuthConfigIdcsConfigRepresentation},
	}
	GenerativeAiHostedApplicationNetworkingConfigRepresentation = map[string]interface{}{
		"inbound_networking_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationNetworkingConfigInboundNetworkingConfigRepresentation},
		"outbound_networking_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationNetworkingConfigOutboundNetworkingConfigRepresentation},
	}
	GenerativeAiHostedApplicationScalingConfigRepresentation = map[string]interface{}{
		"scaling_type":                 acctest.Representation{RepType: acctest.Required, Create: `CPU`, Update: `MEMORY`},
		"max_replica":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_replica":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"target_concurrency_threshold": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"target_cpu_threshold":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"target_memory_threshold":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"target_rps_threshold":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	GenerativeAiHostedApplicationStorageConfigsRepresentation = map[string]interface{}{
		"environment_variable_key": acctest.Representation{RepType: acctest.Required, Create: `environmentVariableKey`},
		"storage_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_storage.test_storage.id}`},
	}
	GenerativeAiHostedApplicationInboundAuthConfigIdcsConfigRepresentation = map[string]interface{}{
		"domain_url": acctest.Representation{RepType: acctest.Required, Create: `domainUrl`, Update: `domainUrl2`},
		"scope":      acctest.Representation{RepType: acctest.Required, Create: `scope`, Update: `scope2`},
		"audience":   acctest.Representation{RepType: acctest.Optional, Create: `audience`, Update: `audience2`},
	}
	GenerativeAiHostedApplicationNetworkingConfigInboundNetworkingConfigRepresentation = map[string]interface{}{
		"endpoint_mode":       acctest.Representation{RepType: acctest.Required, Create: `PUBLIC`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id}`},
	}
	GenerativeAiHostedApplicationNetworkingConfigOutboundNetworkingConfigRepresentation = map[string]interface{}{
		"network_mode":     acctest.Representation{RepType: acctest.Required, Create: `MANAGED`},
		"custom_subnet_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"nsg_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
	}

	GenerativeAiHostedApplicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Required, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiHostedApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiHostedApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_hosted_application.test_hosted_application"
	datasourceName := "data.oci_generative_ai_hosted_applications.test_hosted_applications"
	singularDatasourceName := "data.oci_generative_ai_hosted_application.test_hosted_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiHostedApplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationRepresentation), "generativeai", "hostedApplication", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiHostedApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Required, acctest.Create, GenerativeAiHostedApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\": \"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.domain_url", "domainUrl"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.scope", "scope"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.inbound_auth_config_type", "IDCS_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.inbound_networking_config.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.outbound_networking_config.0.custom_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.0.environment_variable_key", "environmentVariableKey"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_configs.0.storage_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiHostedApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiHostedApplicationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\": \"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.domain_url", "domainUrl"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.scope", "scope"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.inbound_auth_config_type", "IDCS_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.inbound_networking_config.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.outbound_networking_config.0.custom_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.0.environment_variable_key", "environmentVariableKey"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_configs.0.storage_id"),

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
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "VAULT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\": \"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.audience", "audience2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.domain_url", "domainUrl2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.idcs_config.0.scope", "scope2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_auth_config.0.inbound_auth_config_type", "IDCS_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.inbound_networking_config.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "networking_config.0.outbound_networking_config.0.custom_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "11"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "11"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "MEMORY"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage_configs.0.environment_variable_key", "environmentVariableKey"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_configs.0.storage_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_applications", "test_hosted_applications", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "hosted_application_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hosted_application_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Required, acctest.Create, GenerativeAiHostedApplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hosted_application_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.type", "VAULT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.value", "{\"dummyKey\": \"dummyValue\"}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.0.idcs_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.0.idcs_config.0.audience", "audience2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.0.idcs_config.0.domain_url", "domainUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.0.idcs_config.0.scope", "scope2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_auth_config.0.inbound_auth_config_type", "IDCS_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.max_replica", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.min_replica", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.scaling_type", "MEMORY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_concurrency_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_cpu_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_memory_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_rps_threshold", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_configs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_configs.0.environment_variable_key", "environmentVariableKey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiHostedApplicationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiHostedApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_hosted_application" {
			noResourceFound = false
			request := oci_generative_ai.GetHostedApplicationRequest{}

			tmp := rs.Primary.ID
			request.HostedApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetHostedApplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.HostedApplicationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiHostedApplication") {
		resource.AddTestSweepers("GenerativeAiHostedApplication", &resource.Sweeper{
			Name:         "GenerativeAiHostedApplication",
			Dependencies: acctest.DependencyGraph["hostedApplication"],
			F:            sweepGenerativeAiHostedApplicationResource,
		})
	}
}

func sweepGenerativeAiHostedApplicationResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	hostedApplicationIds, err := getGenerativeAiHostedApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, hostedApplicationId := range hostedApplicationIds {
		if ok := acctest.SweeperDefaultResourceId[hostedApplicationId]; !ok {
			deleteHostedApplicationRequest := oci_generative_ai.DeleteHostedApplicationRequest{}

			deleteHostedApplicationRequest.HostedApplicationId = &hostedApplicationId

			deleteHostedApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteHostedApplication(context.Background(), deleteHostedApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting HostedApplication %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostedApplicationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostedApplicationId, GenerativeAiHostedApplicationSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiHostedApplicationSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiHostedApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HostedApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listHostedApplicationsRequest := oci_generative_ai.ListHostedApplicationsRequest{}
	listHostedApplicationsRequest.CompartmentId = &compartmentId
	listHostedApplicationsRequest.LifecycleState = oci_generative_ai.HostedApplicationBaseLifecycleStateActive
	listHostedApplicationsResponse, err := generativeAiClient.ListHostedApplications(context.Background(), listHostedApplicationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HostedApplication list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, hostedApplication := range listHostedApplicationsResponse.Items {
		id := *hostedApplication.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HostedApplicationId", id)
	}
	return resourceIds, nil
}

func GenerativeAiHostedApplicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if hostedApplicationResponse, ok := response.Response.(oci_generative_ai.GetHostedApplicationResponse); ok {
		return hostedApplicationResponse.LifecycleState != oci_generative_ai.HostedApplicationLifecycleStateDeleted
	}
	return false
}

func GenerativeAiHostedApplicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetHostedApplication(context.Background(), oci_generative_ai.GetHostedApplicationRequest{
		HostedApplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
