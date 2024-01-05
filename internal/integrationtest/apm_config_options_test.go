// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ConfigOptionsRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Required, acctest.Create, configOptionsRepresentation)

	ConfigDataResourceOptions = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Required, acctest.Create, configOptionsSingularDataSourceRepresentation)

	configOptionsSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_options.id}`},
	}

	configOptionsDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeOptions, Update: configTypeOptions},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: configOptionsDataSourceFilterRepresentation},
	}

	configOptionsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_options.id}`}},
	}

	configTypeOptions = "OPTIONS"

	configOptionsRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Required, Create: configTypeOptions, Update: configTypeOptions},
		"defined_tags": acctest.Representation{RepType: acctest.Optional,
			Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`,
			Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"group":         acctest.Representation{RepType: acctest.Required, Create: `group`, Update: `group2`},
		"options":       acctest.Representation{RepType: acctest.Required, Create: `{\"dummyKey\": \"dummyValue\"}`},
	}
)

// issue-routing-tag: apm_config/default
func TestApmConfigOptionsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigOptionsResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	optionsResourceName := "oci_apm_config_config.test_options"
	datasourceOptionsName := "data.oci_apm_config_configs.test_optionss"
	singularOptionsDatasourceName := "data.oci_apm_config_config.test_options"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmConfigConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Optional, acctest.Create, configOptionsRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigOptionsDestroy,
		Steps: []resource.TestStep{
			// Step 1: verify create Options
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Required, acctest.Create, configOptionsRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(optionsResourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "id"),
					resource.TestCheckResourceAttr(optionsResourceName, "config_type", configTypeOptions),
					resource.TestCheckResourceAttr(optionsResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(optionsResourceName, "group", "group"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "etag"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, optionsResourceName, "id")
						return err
					},
				),
			},
			// Step 2: delete Options
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies,
			},
			// Step 3: create Options with optionals
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Optional, acctest.Create, configOptionsRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(optionsResourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "id"),
					resource.TestCheckResourceAttr(optionsResourceName, "config_type", configTypeOptions),
					resource.TestCheckResourceAttr(optionsResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(optionsResourceName, "group", "group"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "options"),
					resource.TestCheckResourceAttr(optionsResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(optionsResourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(optionsResourceName, "etag"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, optionsResourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, optionsResourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// Step 4: delete Options before next create
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies,
			},
			// Step 5: verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_optionss", acctest.Optional, acctest.Update, configOptionsDataSourceRepresentation) +
					compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Optional, acctest.Update, configOptionsRepresentation) +
					ConfigDataResourceOptions,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceOptionsName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceOptionsName, "config_type", configTypeOptions),
					resource.TestCheckResourceAttr(datasourceOptionsName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceOptionsName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceOptionsName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 6: verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Required, acctest.Create, configOptionsSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_options", acctest.Optional, acctest.Update, configOptionsRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularOptionsDatasourceName, "config_type", configTypeOptions),
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularOptionsDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularOptionsDatasourceName, "group", "group2"),
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "options"),
					resource.TestCheckResourceAttr(singularOptionsDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularOptionsDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularOptionsDatasourceName, "updated_by"),
				),
			},
			// Step 7 verify resource import
			{
				Config:            config + ConfigOptionsRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
					"opc_dry_run",
				},
				ResourceName: optionsResourceName,
			},
		},
	})
}

func testAccCheckApmConfigOptionsDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("ApmConfigOptions") {
		resource.AddTestSweepers("ApmConfigOptions", &resource.Sweeper{
			Name:         "ApmConfigOptions",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigOptionsResource,
		})
	}
}

func sweepApmConfigOptionsResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getOptionsIds(compartment)
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

func getOptionsIds(compartment string) ([]string, error) {
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
