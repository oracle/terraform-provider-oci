// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApmSyntheticsOnPremiseVantagePointRequiredOnlyResource = ApmSyntheticsOnPremiseVantagePointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointRepresentation)

	ApmSyntheticsOnPremiseVantagePointResourceConfig = ApmSyntheticsOnPremiseVantagePointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointRepresentation)

	ApmSyntheticsOnPremiseVantagePointSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"on_premise_vantage_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id}`},
	}

	ApmSyntheticsOnPremiseVantagePointDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `OPVP-name`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `OPVP-name`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsOnPremiseVantagePointDataSourceFilterRepresentation}}
	ApmSyntheticsOnPremiseVantagePointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.name}`}},
	}

	ApmSyntheticsOnPremiseVantagePointRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `OPVP-name`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"type":          acctest.Representation{RepType: acctest.Optional, Create: `ON_PREMISE_DOCKER_VANTAGE_POINT`},
	}

	ApmSyntheticsOnPremiseVantagePointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsOnPremiseVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsOnPremiseVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point"
	datasourceName := "data.oci_apm_synthetics_on_premise_vantage_points.test_on_premise_vantage_points"
	singularDatasourceName := "data.oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsOnPremiseVantagePointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Optional, acctest.Create, ApmSyntheticsOnPremiseVantagePointRepresentation), "apmsynthetics", "onPremiseVantagePoint", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsOnPremiseVantagePointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "OPVP-name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Optional, acctest.Create, ApmSyntheticsOnPremiseVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "OPVP-name"),
				resource.TestCheckResourceAttr(resourceName, "type", "ON_PREMISE_DOCKER_VANTAGE_POINT"),
				resource.TestCheckResourceAttr(resourceName, "workers_summary.#", "1"),

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
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "OPVP-name"),
				resource.TestCheckResourceAttr(resourceName, "type", "ON_PREMISE_DOCKER_VANTAGE_POINT"),
				resource.TestCheckResourceAttr(resourceName, "workers_summary.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_points", "test_on_premise_vantage_points", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "OPVP-name"),
				resource.TestCheckResourceAttr(datasourceName, "name", "OPVP-name"),

				resource.TestCheckResourceAttr(datasourceName, "on_premise_vantage_point_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "on_premise_vantage_point_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "on_premise_vantage_point_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "OPVP-name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ON_PREMISE_DOCKER_VANTAGE_POINT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "workers_summary.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsOnPremiseVantagePointRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmSyntheticsOnPremiseVantagePointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_on_premise_vantage_point" {
			noResourceFound = false
			request := oci_apm_synthetics.GetOnPremiseVantagePointRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.OnPremiseVantagePointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")

			_, err := client.GetOnPremiseVantagePoint(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("ApmSyntheticsOnPremiseVantagePoint") {
		resource.AddTestSweepers("ApmSyntheticsOnPremiseVantagePoint", &resource.Sweeper{
			Name:         "ApmSyntheticsOnPremiseVantagePoint",
			Dependencies: acctest.DependencyGraph["onPremiseVantagePoint"],
			F:            sweepApmSyntheticsOnPremiseVantagePointResource,
		})
	}
}

func sweepApmSyntheticsOnPremiseVantagePointResource(compartment string) error {
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()
	onPremiseVantagePointIds, err := getApmSyntheticsOnPremiseVantagePointIds(compartment)
	if err != nil {
		return err
	}
	for _, onPremiseVantagePointId := range onPremiseVantagePointIds {
		if ok := acctest.SweeperDefaultResourceId[onPremiseVantagePointId]; !ok {
			deleteOnPremiseVantagePointRequest := oci_apm_synthetics.DeleteOnPremiseVantagePointRequest{}

			deleteOnPremiseVantagePointRequest.OnPremiseVantagePointId = &onPremiseVantagePointId

			deleteOnPremiseVantagePointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteOnPremiseVantagePoint(context.Background(), deleteOnPremiseVantagePointRequest)
			if error != nil {
				fmt.Printf("Error deleting OnPremiseVantagePoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", onPremiseVantagePointId, error)
				continue
			}
		}
	}
	return nil
}

func getApmSyntheticsOnPremiseVantagePointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OnPremiseVantagePointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()

	listOnPremiseVantagePointsRequest := oci_apm_synthetics.ListOnPremiseVantagePointsRequest{}
	//listOnPremiseVantagePointsRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for OnPremiseVantagePoint resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listOnPremiseVantagePointsRequest.ApmDomainId = &apmDomainId

		listOnPremiseVantagePointsResponse, err := apmSyntheticClient.ListOnPremiseVantagePoints(context.Background(), listOnPremiseVantagePointsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting OnPremiseVantagePoint list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, onPremiseVantagePoint := range listOnPremiseVantagePointsResponse.Items {
			id := *onPremiseVantagePoint.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OnPremiseVantagePointId", id)
		}

	}
	return resourceIds, nil
}
