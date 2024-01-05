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
	ApmConfigConfigRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Required, acctest.Create, configSpanFilterRepresentation)

	ApmConfigConfigResourceSpanFilter = ApmConfigConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Optional, acctest.Update, configSpanFilterRepresentation)

	ApmConfigConfigDataResourceSpanFilter = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Required, acctest.Create, ApmConfigconfigSingularDataSourceRepresentation)

	ApmConfigconfigSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_span_filter.id}`},
	}

	ApmConfigconfigSpanFilterDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeSpanFilter, Update: configTypeSpanFilter},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmConfigconfigSpanFilterDataSourceFilterRepresentation},
	}

	ApmConfigconfigSpanFilterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_span_filter.id}`}},
	}

	configTypeSpanFilter = "SPAN_FILTER"

	createFilterText = `kind='SERVER'`
	updateFilterText = `kind='CLIENT'`

	configSpanFilterRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Required, Create: configTypeSpanFilter, Update: configTypeSpanFilter},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"filter_text":   acctest.Representation{RepType: acctest.Required, Create: createFilterText, Update: updateFilterText},
		"defined_tags": acctest.Representation{RepType: acctest.Optional,
			Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`,
			Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ApmConfigConfigResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: apm_config/default
func TestApmConfigConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigConfigResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_config_config.test_span_filter"
	datasourceName := "data.oci_apm_config_configs.test_span_filters"
	singularDatasourceName := "data.oci_apm_config_config.test_span_filter"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmConfigConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Optional, acctest.Create, configSpanFilterRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigConfigDestroy,
		Steps: []resource.TestStep{
			// Find these steps in the test log easily with "Executing step (number)"
			// Step 0: verify create Span Filter
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Required, acctest.Create, configSpanFilterRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeSpanFilter),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "filter_text", createFilterText),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// Step 1: delete Span Filter before next create
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies,
			},
			// Step 2: verify create Span Filter with optionals
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Optional, acctest.Create, configSpanFilterRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeSpanFilter),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "filter_text", createFilterText),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),

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
			// Step 3: verify updates to Span Filter updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Optional, acctest.Update, configSpanFilterRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeSpanFilter),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_text", updateFilterText),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 4: verify datasource (span filter)
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_span_filters", acctest.Optional, acctest.Update, ApmConfigconfigSpanFilterDataSourceRepresentation) +
					compartmentIdVariableStr + ApmConfigConfigResourceSpanFilter + ApmConfigConfigDataResourceSpanFilter,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "config_type", configTypeSpanFilter),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 5 verify singular datasource (span filter)
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceSpanFilter + ApmConfigConfigDataResourceSpanFilter,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "filter_text", updateFilterText),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updated_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "in_use_by.0.items.#", "0"),
				),
			},
			// Step 6 remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceSpanFilter,
			},
			// Step 7 verify resource import
			{
				Config:            config + ApmConfigConfigRequiredOnlyResource,
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

func testAccCheckApmConfigConfigDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("ApmConfigConfig") {
		resource.AddTestSweepers("ApmConfigConfig", &resource.Sweeper{
			Name:         "ApmConfigConfig",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigConfigResource,
		})
	}
}

func sweepApmConfigConfigResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getConfigIds(compartment)
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

func getConfigIds(compartment string) ([]string, error) {
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
