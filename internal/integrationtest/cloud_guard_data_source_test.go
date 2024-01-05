// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"
)

var (
	CloudGuardDataSourceRequiredOnlyResource = CloudGuardDataSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Required, acctest.Create, CloudGuardDataSourceRepresentation)

	CloudGuardDataSourceResourceConfig = CloudGuardDataSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Update, CloudGuardDataSourceRepresentation)

	CloudGuardCloudGuardDataSourceSingularDataSourceRepresentation = map[string]interface{}{
		"data_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_data_source.test_data_source.id}`},
	}

	CloudGuardCloudGuardDataSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"data_source_feed_provider": acctest.Representation{RepType: acctest.Optional, Create: `LOGGINGQUERY`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"logging_query_type":        acctest.Representation{RepType: acctest.Optional, Create: `INSIGHT`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardDataSourceDataSourceFilterRepresentation}}
	CloudGuardDataSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_data_source.test_data_source.id}`}},
	}

	CloudGuardDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_feed_provider": acctest.Representation{RepType: acctest.Required, Create: `LOGGINGQUERY`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"data_source_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudGuardDataSourceDataSourceDetailsRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	CloudGuardDataSourceDataSourceDetailsRepresentation = map[string]interface{}{
		"data_source_feed_provider": acctest.Representation{RepType: acctest.Required, Create: `LOGGINGQUERY`},
		"additional_entities_count": acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `2`},
		"interval_in_minutes":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"logging_query_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudGuardDataSourceDataSourceDetailsLoggingQueryDetailsRepresentation},
		"logging_query_type":        acctest.Representation{RepType: acctest.Optional, Create: `INSIGHT`},
		"operator":                  acctest.Representation{RepType: acctest.Optional, Create: `GREATERTHANEQUALTO`, Update: `GREATERTHANEQUALTO`},
		"query":                     acctest.Representation{RepType: acctest.Optional, Create: `search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg01, data.resourceId as cg02`, Update: `search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg02, data.resourceId as cg01`},
		"query_start_time":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudGuardDataSourceDataSourceDetailsQueryStartTimeRepresentation},
		"regions":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`us-phoenix-1`}, Update: []string{`us-phoenix-1`}},
		"threshold":                 acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
	}
	CloudGuardDataSourceDataSourceDetailsLoggingQueryDetailsRepresentation = map[string]interface{}{
		"logging_query_type": acctest.Representation{RepType: acctest.Required, Create: `INSIGHT`},
		"key_entities_count": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
	}
	queryStartTime                                                    = time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond)
	CloudGuardDataSourceDataSourceDetailsQueryStartTimeRepresentation = map[string]interface{}{
		"start_policy_type": acctest.Representation{RepType: acctest.Required, Create: `ABSOLUTE_TIME_START_POLICY`},
		"query_start_time":  acctest.Representation{RepType: acctest.Optional, Create: queryStartTime.Format(time.RFC3339Nano)},
	}

	CloudGuardDataSourceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardDataSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardDataSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_guard_data_source.test_data_source"
	datasourceName := "data.oci_cloud_guard_data_sources.test_data_sources"
	singularDatasourceName := "data.oci_cloud_guard_data_source.test_data_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardDataSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Create, CloudGuardDataSourceRepresentation), "cloudguard", "dataSource", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardDataSourceDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Create, CloudGuardDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.additional_entities_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.interval_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.key_entities_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.operator", "GREATERTHANEQUALTO"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query", "search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg01, data.resourceId as cg02"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source_details.0.query_start_time.0.query_start_time"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.0.start_policy_type", "ABSOLUTE_TIME_START_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.regions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudGuardDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudGuardDataSourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.additional_entities_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.interval_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.key_entities_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.operator", "GREATERTHANEQUALTO"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query", "search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg01, data.resourceId as cg02"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source_details.0.query_start_time.0.query_start_time"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.0.start_policy_type", "ABSOLUTE_TIME_START_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.regions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + CloudGuardDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Update, CloudGuardDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.additional_entities_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.interval_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.key_entities_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.operator", "GREATERTHANEQUALTO"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query", "search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg02, data.resourceId as cg01"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source_details.0.query_start_time.0.query_start_time"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.query_start_time.0.start_policy_type", "ABSOLUTE_TIME_START_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.regions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_sources", "test_data_sources", acctest.Optional, acctest.Update, CloudGuardCloudGuardDataSourceDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Optional, acctest.Update, CloudGuardDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "data_source_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_source_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_source", "test_data_source", acctest.Required, acctest.Create, CloudGuardCloudGuardDataSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardDataSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_source_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.additional_entities_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.interval_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.logging_query_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.logging_query_details.0.key_entities_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.logging_query_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.logging_query_type", "INSIGHT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.operator", "GREATERTHANEQUALTO"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.query", "search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 5 | select data.eventName as cgkey01, data.message as cg02, data.resourceId as cg01"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.query_start_time.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_source_details.0.query_start_time.0.query_start_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.query_start_time.0.start_policy_type", "ABSOLUTE_TIME_START_POLICY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.regions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_details.0.threshold", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_detector_mapping_info.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_feed_provider", "LOGGINGQUERY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region_status_detail.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardDataSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardDataSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_data_source" {
			noResourceFound = false
			request := oci_cloud_guard.GetDataSourceRequest{}

			tmp := rs.Primary.ID
			request.DataSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetDataSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudGuardDataSource") {
		resource.AddTestSweepers("CloudGuardDataSource", &resource.Sweeper{
			Name:         "CloudGuardDataSource",
			Dependencies: acctest.DependencyGraph["dataSource"],
			F:            sweepCloudGuardDataSourceResource,
		})
	}
}

func sweepCloudGuardDataSourceResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	dataSourceIds, err := getCloudGuardDataSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, dataSourceId := range dataSourceIds {
		if ok := acctest.SweeperDefaultResourceId[dataSourceId]; !ok {
			deleteDataSourceRequest := oci_cloud_guard.DeleteDataSourceRequest{}

			deleteDataSourceRequest.DataSourceId = &dataSourceId

			deleteDataSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteDataSource(context.Background(), deleteDataSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting DataSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataSourceId, CloudGuardDataSourceSweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardDataSourceSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardDataSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listDataSourcesRequest := oci_cloud_guard.ListDataSourcesRequest{}
	listDataSourcesRequest.CompartmentId = &compartmentId
	listDataSourcesRequest.LifecycleState = oci_cloud_guard.ListDataSourcesLifecycleStateActive
	listDataSourcesResponse, err := cloudGuardClient.ListDataSources(context.Background(), listDataSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataSource := range listDataSourcesResponse.Items {
		id := *dataSource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataSourceId", id)
	}
	return resourceIds, nil
}

func CloudGuardDataSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataSourceResponse, ok := response.Response.(oci_cloud_guard.GetDataSourceResponse); ok {
		return dataSourceResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardDataSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetDataSource(context.Background(), oci_cloud_guard.GetDataSourceRequest{
		DataSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
