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
	oci_apm_config "github.com/oracle/oci-go-sdk/v56/apmconfig"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConfigDataResourceMetricGroup = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Required, acctest.Create, configMGroupSingularDataSourceRepresentation)

	configMGroupSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_metric_group.id}`},
	}

	configMGroupDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeMetricGroup, Update: configTypeMetricGroup},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: configMGroupDataSourceFilterRepresentation},
	}

	configMGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_metric_group.id}`}},
	}

	configTypeMetricGroup = "METRIC_GROUP"

	configMetricGroupRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Required, Create: configTypeMetricGroup, Update: configTypeMetricGroup},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"filter_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_apm_config_config.test_span_filter.id}`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `oracle_apm_monitoring`, Update: `oracle_apm_rum`},
		"metrics":       acctest.RepresentationGroup{RepType: acctest.Required, Group: configMetricsRepresentation},
	}

	configMetricsRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `ThreadCpuTime`, Update: `PageResponseTime`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"unit":        acctest.Representation{RepType: acctest.Optional, Create: `unit`, Update: `unit2`},
	}
)

// issue-routing-tag: apm_config/default
func TestApmConfigMetricGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigMetricGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	spanFilterResourceName := "oci_apm_config_config.test_span_filter"
	metricGroupResourceName := "oci_apm_config_config.test_metric_group"
	datasourceMGName := "data.oci_apm_config_configs.test_metric_groups"
	singularMGDatasourceName := "data.oci_apm_config_config.test_metric_group"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Optional, acctest.Create, configMetricGroupRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigMetricGroupDestroy,
		Steps: []resource.TestStep{
			// Find these steps in the test log easily with "Executing step (number)"
			// Step 0: create Span Filter - we need it to create a Metric Group
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_span_filter", acctest.Required, acctest.Create, configSpanFilterRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(spanFilterResourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(spanFilterResourceName, "id"),
					resource.TestCheckResourceAttr(spanFilterResourceName, "config_type", configTypeSpanFilter),
					resource.TestCheckResourceAttr(spanFilterResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(spanFilterResourceName, "filter_text", createFilterText),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, spanFilterResourceName, "id")
						return err
					},
				),
			},
			// Step 1: verify create Metric Group
			{
				Config: config + compartmentIdVariableStr + ConfigResourceSpanFilter + ConfigDataResourceSpanFilter +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Required, acctest.Create, configMetricGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "id"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "config_type", configTypeMetricGroup),
					resource.TestCheckResourceAttr(metricGroupResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "filter_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, metricGroupResourceName, "id")
						return err
					},
				),
			},
			// Step 2: delete Metric Group before next create
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies,
			},
			// Step 3: create Metric Group with optionals
			{
				Config: config + compartmentIdVariableStr + ConfigResourceSpanFilter + ConfigDataResourceSpanFilter +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Optional, acctest.Create, configMetricGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "id"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "config_type", configTypeMetricGroup),
					resource.TestCheckResourceAttr(metricGroupResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(metricGroupResourceName, "filter_id"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "metrics.#", "1"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "metrics.0.description", "description"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "metrics.0.name", "ThreadCpuTime"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "metrics.0.unit", "unit"),
					resource.TestCheckResourceAttr(metricGroupResourceName, "namespace", "oracle_apm_monitoring"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, metricGroupResourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, metricGroupResourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// Step 4: delete Metric Group before next create
			{
				Config: config + compartmentIdVariableStr + ConfigResourceDependencies,
			},
			// Step 5: verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_metric_groups", acctest.Optional, acctest.Update, configMGroupDataSourceRepresentation) +
					compartmentIdVariableStr + ConfigResourceSpanFilter + ConfigDataResourceSpanFilter +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Optional, acctest.Update, configMetricGroupRepresentation) +
					ConfigDataResourceMetricGroup,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceMGName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceMGName, "config_type", configTypeMetricGroup),
					resource.TestCheckResourceAttr(datasourceMGName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceMGName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceMGName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 6: verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Required, acctest.Create, configMGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ConfigResourceSpanFilter + ConfigDataResourceSpanFilter +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Optional, acctest.Update, configMetricGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularMGDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularMGDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "config_type", configTypeMetricGroup),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularMGDatasourceName, "filter_id"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "metrics.#", "1"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "metrics.0.description", "description2"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "metrics.0.name", "PageResponseTime"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "metrics.0.unit", "unit2"),
					resource.TestCheckResourceAttr(singularMGDatasourceName, "namespace", "oracle_apm_rum"),
				),
			},
			// Step 7 remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ConfigResourceSpanFilter + ConfigDataResourceSpanFilter +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_metric_group", acctest.Optional, acctest.Update, configMetricGroupRepresentation),
			},
			// Step 8 verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
					"opc_dry_run",
				},
				ResourceName: metricGroupResourceName,
			},
		},
	})
}

func testAccCheckApmConfigMetricGroupDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("ApmConfigMetricGroup") {
		resource.AddTestSweepers("ApmConfigMetricGroup", &resource.Sweeper{
			Name:         "ApmConfigMetricGroup",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigMetricGroupResource,
		})
	}
}

func sweepApmConfigMetricGroupResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getMetricGroupIds(compartment)
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

func getMetricGroupIds(compartment string) ([]string, error) {
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
