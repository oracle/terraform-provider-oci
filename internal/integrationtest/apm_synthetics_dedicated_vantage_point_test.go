// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v55/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v55/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DedicatedVantagePointRequiredOnlyResource = DedicatedVantagePointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Required, acctest.Create, dedicatedVantagePointRepresentation)

	DedicatedVantagePointResourceConfig = DedicatedVantagePointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Optional, acctest.Update, dedicatedVantagePointRepresentation)

	dedicatedVantagePointSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"dedicated_vantage_point_id": Representation{RepType: Required, Create: `${oci_apm_synthetics_dedicated_vantage_point.test_dedicated_vantage_point.id}`},
	}

	dedicatedVantagePointDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: dedicatedVantagePointDataSourceFilterRepresentation}}
	dedicatedVantagePointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_dedicated_vantage_point.test_dedicated_vantage_point.id}`}},
	}

	dedicatedVantagePointRepresentation = map[string]interface{}{
		"apm_domain_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"dvp_stack_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: dedicatedVantagePointDvpStackDetailsRepresentation},
		"region":            acctest.Representation{RepType: acctest.Required, Create: `region`, Update: `region2`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":            acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}
	dedicatedVantagePointDvpStackDetailsRepresentation = map[string]interface{}{
		"dvp_stack_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_resourcemanager_stack.test_stack.id}`},
		"dvp_stack_type": acctest.Representation{RepType: acctest.Required, Create: `ORACLE_RM_STACK`},
		"dvp_stream_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_streaming_stream.test_stream.id}`},
		"dvp_version":    acctest.Representation{RepType: acctest.Required, Create: `dvpVersion`, Update: `dvpVersion2`},
	}

	DedicatedVantagePointResourceDependencies = GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
		DefinedTagsDependencies +
		GenerateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", Required, Create, stackDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsDedicatedVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsDedicatedVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_dedicated_vantage_point.test_dedicated_vantage_point"
	datasourceName := "data.oci_apm_synthetics_dedicated_vantage_points.test_dedicated_vantage_points"
	singularDatasourceName := "data.oci_apm_synthetics_dedicated_vantage_point.test_dedicated_vantage_point"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DedicatedVantagePointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Optional, acctest.Create, dedicatedVantagePointRepresentation), "apmsynthetics", "dedicatedVantagePoint", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsDedicatedVantagePointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Required, acctest.Create, dedicatedVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stack_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_stack_type", "ORACLE_RM_STACK"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stream_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_version", "dvpVersion"),
				resource.TestCheckResourceAttr(resourceName, "region", "region"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVantagePointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DedicatedVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Optional, acctest.Create, dedicatedVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stack_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_stack_type", "ORACLE_RM_STACK"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stream_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_version", "dvpVersion"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "region", "region"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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
			Config: config + compartmentIdVariableStr + DedicatedVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Optional, acctest.Update, dedicatedVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stack_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_stack_type", "ORACLE_RM_STACK"),
				resource.TestCheckResourceAttrSet(resourceName, "dvp_stack_details.0.dvp_stream_id"),
				resource.TestCheckResourceAttr(resourceName, "dvp_stack_details.0.dvp_version", "dvpVersion2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "region", "region2"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_points", "test_dedicated_vantage_points", acctest.Optional, acctest.Update, dedicatedVantagePointDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Optional, acctest.Update, dedicatedVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "dedicated_vantage_point_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vantage_point_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_dedicated_vantage_point", "test_dedicated_vantage_point", acctest.Required, acctest.Create, dedicatedVantagePointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVantagePointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vantage_point_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dvp_stack_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dvp_stack_details.0.dvp_stack_type", "ORACLE_RM_STACK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dvp_stack_details.0.dvp_version", "dvpVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_status_count_map.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region", "region2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DedicatedVantagePointResourceConfig,
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

func testAccCheckApmSyntheticsDedicatedVantagePointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_dedicated_vantage_point" {
			noResourceFound = false
			request := oci_apm_synthetics.GetDedicatedVantagePointRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.DedicatedVantagePointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")

			_, err := client.GetDedicatedVantagePoint(context.Background(), request)

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("ApmSyntheticsDedicatedVantagePoint") {
		resource.AddTestSweepers("ApmSyntheticsDedicatedVantagePoint", &resource.Sweeper{
			Name:         "ApmSyntheticsDedicatedVantagePoint",
			Dependencies: DependencyGraph["dedicatedVantagePoint"],
			F:            sweepApmSyntheticsDedicatedVantagePointResource,
		})
	}
}

func sweepApmSyntheticsDedicatedVantagePointResource(compartment string) error {
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()
	dedicatedVantagePointIds, err := getDedicatedVantagePointIds(compartment)
	if err != nil {
		return err
	}
	for _, dedicatedVantagePointId := range dedicatedVantagePointIds {
		if ok := SweeperDefaultResourceId[dedicatedVantagePointId]; !ok {
			deleteDedicatedVantagePointRequest := oci_apm_synthetics.DeleteDedicatedVantagePointRequest{}

			deleteDedicatedVantagePointRequest.DedicatedVantagePointId = &dedicatedVantagePointId

			deleteDedicatedVantagePointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteDedicatedVantagePoint(context.Background(), deleteDedicatedVantagePointRequest)
			if error != nil {
				fmt.Printf("Error deleting DedicatedVantagePoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", dedicatedVantagePointId, error)
				continue
			}
		}
	}
	return nil
}

func getDedicatedVantagePointIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DedicatedVantagePointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()

	listDedicatedVantagePointsRequest := oci_apm_synthetics.ListDedicatedVantagePointsRequest{}
	listDedicatedVantagePointsRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for DedicatedVantagePoint resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listDedicatedVantagePointsRequest.ApmDomainId = &apmDomainId

		listDedicatedVantagePointsResponse, err := apmSyntheticClient.ListDedicatedVantagePoints(context.Background(), listDedicatedVantagePointsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DedicatedVantagePoint list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dedicatedVantagePoint := range listDedicatedVantagePointsResponse.Items {
			id := *dedicatedVantagePoint.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "DedicatedVantagePointId", id)
		}

	}
	return resourceIds, nil
}
