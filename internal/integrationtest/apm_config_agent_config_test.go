// Copyright (c) 2025, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ApmConfigAgentConfigRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Required, acctest.Create, configAgentConfigRepresentation)

	ApmConfigAgentConfigResource = ApmConfigConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Optional, acctest.Update, configAgentConfigRepresentation)

	ApmConfigAgentConfigDataResource = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Required, acctest.Create, ApmConfigAgentconfigSingularDataSourceRepresentation)

	ApmConfigAgentconfigSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_agent_config.id}`},
	}

	ApmConfigAgentConfigDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeAgentConfig, Update: configTypeAgentConfig},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmConfigAgentConfigFilterDataSourceFilterRepresentation},
	}

	ApmConfigAgentConfigFilterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_agent_config.id}`}},
	}

	configTypeAgentConfig = "AGENT"

	configAgentConfigRepresentation = map[string]interface{}{
		"apm_domain_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":                       acctest.Representation{RepType: acctest.Required, Create: configTypeAgentConfig, Update: configTypeAgentConfig},
		"config":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: configAgentConfigConfigRepresentation},
		"match_agents_with_attribute_value": acctest.Representation{RepType: acctest.Required, Create: "value1"},
		"overrides":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: configAgentConfigOverridesRepresentation},
		"defined_tags": acctest.Representation{RepType: acctest.Optional,
			Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`,
			Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	configAgentConfigConfigRepresentation = map[string]interface{}{
		"config_map": acctest.RepresentationGroup{RepType: acctest.Required, Group: configAgentConfigConfigFileRepresentation},
	}

	configAgentConfigConfigFileRepresentation = map[string]interface{}{
		"file_name": acctest.Representation{RepType: acctest.Required, Create: "example1", Update: "example2"},
		"body": acctest.Representation{RepType: acctest.Required, Create: "Y29tLm9yYWNsZS5hcG0uYWdlbnQudHJhY2VyLmVuYWJsZS5qZnIgPSB7eyBpc0pmckVuYWJsZWQgfCBkZWZhdWx0IGZhbHNlIH19",
			Update: "Y29tLm9yYWNsZS5hcG0uYWdlbnQuY2lyY3VpdC5icmVha2VyLmVuYWJsZSA9IHt7IGlzQ2lyY3VpdEJyZWFrZXJFbmFibGVkIHwgZGVmYXVsdCB0cnVlIH19"},
		"content_type": acctest.Representation{RepType: acctest.Required, Create: "charset=utf-8"},
	}

	configAgentConfigOverridesRepresentation = map[string]interface{}{
		"override_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: configAgentConfigOverrideListRepresentation},
	}

	configAgentConfigOverrideListRepresentation = map[string]interface{}{
		"agent_filter": acctest.Representation{RepType: acctest.Optional, Create: "Component='Server'", Update: "Component='Browser'"},
		"override_map": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"isJfrEnabled": "true"}, Update: map[string]string{"isCircuitBreakerEnabled": "false"}},
	}
)

// issue-routing-tag: apm_config/default
func TestApmConfigAgentConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigAgentConfigResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_config_config.test_agent_config"
	datasourceName := "data.oci_apm_config_configs.test_agent_configs"
	singularDatasourceName := "data.oci_apm_config_config.test_agent_config"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmConfigConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Optional, acctest.Create, configAgentConfigRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigAgentConfigDestroy,
		Steps: []resource.TestStep{
			// Find these steps in the test log easily with "Executing step (number)"
			// Step 1: verify create Agent Config
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Required, acctest.Create, configAgentConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeAgentConfig),
					resource.TestCheckResourceAttr(resourceName, "config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "match_agents_with_attribute_value", "value1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						utils.Logf("This is error")
						return err
					},
				),
			},
			// Step 2: delete Agent Config before next create
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies,
			},
			// Step 3: verify create Agent Config with optionals
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Optional, acctest.Create, configAgentConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeAgentConfig),
					resource.TestCheckResourceAttr(resourceName, "config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config.0.config_map.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config.0.config_map.0.file_name", "example1"),
					resource.TestCheckResourceAttr(resourceName, "overrides.0.override_list.0.agent_filter", "Component='Server'"),
					resource.TestCheckResourceAttr(resourceName, "match_agents_with_attribute_value", "value1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
			// Step 4: verify updates to Agent Config updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_agent_config", acctest.Optional, acctest.Update, configAgentConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeAgentConfig),
					resource.TestCheckResourceAttr(resourceName, "config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "config.0.config_map.0.file_name", "example2"),
					resource.TestCheckResourceAttr(resourceName, "overrides.0.override_list.0.agent_filter", "Component='Browser'"),
					resource.TestCheckResourceAttr(resourceName, "match_agents_with_attribute_value", "value1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 5: verify datasource (Agent Config)
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_agent_configs", acctest.Optional, acctest.Update, ApmConfigAgentConfigDataSourceRepresentation) +
					compartmentIdVariableStr + ApmConfigAgentConfigResource + ApmConfigAgentConfigDataResource,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "config_type", configTypeAgentConfig),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 6: verify singular datasource (Agent Config)
			{
				Config: config + compartmentIdVariableStr + ApmConfigAgentConfigResource + ApmConfigAgentConfigDataResource,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_type", configTypeAgentConfig),
					resource.TestCheckResourceAttr(singularDatasourceName, "config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config.0.config_map.0.file_name", "example2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "overrides.0.override_list.0.agent_filter", "Component='Browser'"),
					resource.TestCheckResourceAttr(singularDatasourceName, "match_agents_with_attribute_value", "value1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				),
			},
			// Step 7: remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ApmConfigAgentConfigResource,
			},
			// Step 8: verify resource import
			{
				Config:            config + ApmConfigAgentConfigRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
					"opc_dry_run",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckApmConfigAgentConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ConfigClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_config_config" {
			noResourceFound = false
			request := oci_apm_config.GetConfigRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp, _ := parseConfigCompositeId(rs.Primary.ID)
			request.ConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")

			_, err := client.GetConfig(context.Background(), request)

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

func parseConfigCompositeId(compositeId string) (configId string, err error) {
	parts := strings.Split(compositeId, "/")

	match, _ := regexp.MatchString("configs/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	configId, _ = url.PathUnescape(parts[1])

	return
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ApmConfigAgentConfig") {
		resource.AddTestSweepers("ApmConfigAgentConfig", &resource.Sweeper{
			Name:         "ApmConfigAgentConfig",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigAgentConfigResource,
		})
	}
}

func sweepApmConfigAgentConfigResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getAgentConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, configId := range configIds {
		if ok := acctest.SweeperDefaultResourceId[configId]; !ok {
			deleteConfigRequest := oci_apm_config.DeleteConfigRequest{}

			deleteConfigRequest.ConfigId = &configId

			deleteConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")
			_, error := configClient.DeleteConfig(context.Background(), deleteConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting Config %s %s, It is possible that the resource is already deleted. Please verify manually \n", configId, error)
				continue
			}
		}
	}
	return nil
}

func getAgentConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()

	listConfigsRequest := oci_apm_config.ListConfigsRequest{}
	//listConfigsRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Config resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listConfigsRequest.ApmDomainId = &apmDomainId

		listConfigsResponse, err := configClient.ListConfigs(context.Background(), listConfigsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Config list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, config := range listConfigsResponse.Items {
			id := *config.GetId()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConfigId", id)
		}

	}
	return resourceIds, nil
}
