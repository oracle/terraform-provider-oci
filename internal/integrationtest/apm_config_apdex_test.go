// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apm_config "github.com/oracle/oci-go-sdk/v58/apmconfig"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConfigResourceApdex = ConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Optional, acctest.Update, configApdexRepresentation)

	ConfigDataResourceApdex = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Required, acctest.Create, configApdexSingularDataSourceRepresentation)

	configApdexSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_apdex.id}`},
	}

	configApdexDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeApdex, Update: configTypeApdex},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: configApdexDataSourceFilterRepresentation},
	}

	configApdexDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_apdex.id}`}},
	}

	configTypeApdex = "APDEX"

	createRulesFilterText = `kind='SERVER'`
	updateRulesFilterText = `kind='CLIENT'`

	configApdexRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Required, Create: configTypeApdex, Update: configTypeApdex},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"rules":         acctest.RepresentationGroup{RepType: acctest.Required, Group: configRulesRepresentation},
	}

	configRulesRepresentation = map[string]interface{}{
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter_text":              acctest.Representation{RepType: acctest.Required, Create: createRulesFilterText, Update: updateRulesFilterText},
		"is_apply_to_error_spans":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"priority":                 acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"satisfied_response_time":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"tolerating_response_time": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
)

// issue-routing-tag: apm_config/default
func TestApmConfigApdexResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigApdexResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_config_config.test_apdex"
	datasourceName := "data.oci_apm_config_configs.test_apdexes"
	singularDatasourceName := "data.oci_apm_config_config.test_apdex"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Optional, acctest.Create, configApdexRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigApdexDestroy,
		Steps: []resource.TestStep{
			// Find these steps in the test log easily with "Executing step (number)"
			// Step 0: verify create Apdex
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Required, acctest.Create, configApdexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeApdex),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.filter_text", createRulesFilterText),
					resource.TestCheckResourceAttr(resourceName, "rules.0.priority", "10"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// Step 1: delete Apdex before next create
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies,
			},
			// Step 2: verify create Apdex with optionals
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Optional, acctest.Create, configApdexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeApdex),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.filter_text", createRulesFilterText),
					resource.TestCheckResourceAttr(resourceName, "rules.0.is_apply_to_error_spans", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.priority", "10"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.satisfied_response_time", "10"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.tolerating_response_time", "10"),

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
			// Step 3: verify updates to Apdex updatable parameters
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_apdex", acctest.Optional, acctest.Update, configApdexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeApdex),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.filter_text", updateRulesFilterText),
					resource.TestCheckResourceAttr(resourceName, "rules.0.is_apply_to_error_spans", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.priority", "11"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.satisfied_response_time", "11"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.tolerating_response_time", "11"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 4: verify datasource (apdex)
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_apdexes", acctest.Optional, acctest.Update, configApdexDataSourceRepresentation) +
					compartmentIdVariableStr + ConfigResourceApdex + ConfigDataResourceApdex,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "config_type", configTypeApdex),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 5 verify singular datasource (apdex)
			{
				Config: config + compartmentIdVariableStr + ConfigResourceApdex + ConfigDataResourceApdex,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.filter_text", updateRulesFilterText),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.is_apply_to_error_spans", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.priority", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.satisfied_response_time", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.tolerating_response_time", "11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// Step 6 remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ConfigResourceApdex,
			},
			// Step 7 verify resource import
			{
				Config:            config,
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

func testAccCheckApmConfigApdexDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ConfigClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_config_config" {
			noResourceFound = false
			request := oci_apm_config.GetConfigRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ApmConfigApdex") {
		resource.AddTestSweepers("ApmConfigApdex", &resource.Sweeper{
			Name:         "ApmConfigApdex",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigApdexResource,
		})
	}
}

func sweepApmConfigApdexResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getApdexIds(compartment)
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

func getApdexIds(compartment string) ([]string, error) {
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
