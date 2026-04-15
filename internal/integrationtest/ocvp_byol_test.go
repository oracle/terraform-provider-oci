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
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// If you want tests to create tags, uncomment below and comment other defined tag variables
	// If tags are created already, comment below and uncomment other variables
	//ocvpDefinedTag             = `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`
	//ocvpDefinedTagUpdate       = `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`
	//ocvpDefinedTagsDependencies = DefinedTagsDifDependencies

	ocvpDefinedTag              = `${map("example-tag-namespace-all.example-tag", "value")}`
	ocvpDefinedTagUpdate        = `${map("example-tag-namespace-all.example-tag", "updatedValue")}`
	ocvpDefinedTagsDependencies = ``

	OcvpByolRequiredOnlyResource = OcvpByolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Required, acctest.Create, OcvpByolRepresentation)

	OcvpByolResourceConfig = OcvpByolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Update, OcvpByolRepresentation)

	OcvpByolSingularDataSourceRepresentation = map[string]interface{}{
		"byol_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_byol.test_byol.id}`},
	}

	OcvpByolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"available_units_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"byol_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_byol.test_byol.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"software_type": acctest.Representation{RepType: acctest.Optional, Create: `VCF`, Update: `VSAN`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpByolDataSourceFilterRepresentation}}
	OcvpByolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_byol.test_byol.id}`}},
	}

	OcvpByolRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"entitlement_key": acctest.Representation{RepType: acctest.Required, Create: `AAAAA-BBBBB-CCCCC-DDDDD-EEEEE`, Update: `AAAAA-BBBBB-CCCCC-DDDDD-HHHHH`},
		"software_type":   acctest.Representation{RepType: acctest.Required, Create: `VCF`, Update: `VSAN`},
		"time_term_end":   acctest.Representation{RepType: acctest.Required, Create: `2028-03-23T01:23:45.678Z`, Update: `2028-03-24T01:23:45.678Z`},
		"time_term_start": acctest.Representation{RepType: acctest.Required, Create: `2026-03-23T01:23:45.678Z`, Update: `2026-03-24T01:23:45.678Z`},
		"total_units":     acctest.Representation{RepType: acctest.Required, Create: `1000`, Update: `1100`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: ocvpDefinedTag, Update: ocvpDefinedTagUpdate},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}

	OcvpByolVdefendRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(OcvpByolRepresentation, []string{"software_type"}),
		map[string]interface{}{
			"software_type": acctest.Representation{RepType: acctest.Required, Create: `VDEFEND`},
		})

	OcvpByolLbRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(OcvpByolRepresentation, []string{"software_type"}),
		map[string]interface{}{
			"software_type": acctest.Representation{RepType: acctest.Required, Create: `AVI_LOAD_BALANCER`},
		})

	OcvpByolResourceDependencies = ocvpDefinedTagsDependencies
)

// issue-routing-tag: ocvp/default
func TestOcvpByolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpByolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_byol.test_byol"
	datasourceName := "data.oci_ocvp_byols.test_byols"
	singularDatasourceName := "data.oci_ocvp_byol.test_byol"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpByolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Create, OcvpByolRepresentation), "ocvp", "byol", t)

	acctest.ResourceTest(t, testAccCheckOcvpByolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpByolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Required, acctest.Create, OcvpByolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entitlement_key", "AAAAA-BBBBB-CCCCC-DDDDD-EEEEE"),
				resource.TestCheckResourceAttr(resourceName, "software_type", "VCF"),
				resource.TestCheckResourceAttr(resourceName, "time_term_end", "2028-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttr(resourceName, "time_term_start", "2026-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttr(resourceName, "total_units", "1000"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpByolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpByolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Create, OcvpByolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entitlement_key", "AAAAA-BBBBB-CCCCC-DDDDD-EEEEE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_type", "VCF"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_term_end", "2028-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttr(resourceName, "time_term_start", "2026-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "total_units", "1000"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OcvpByolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OcvpByolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entitlement_key", "AAAAA-BBBBB-CCCCC-DDDDD-EEEEE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_type", "VCF"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_term_end", "2028-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttr(resourceName, "time_term_start", "2026-03-23T01:23:45.678Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "total_units", "1000"),

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
			Config: config + compartmentIdVariableStr + OcvpByolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Update, OcvpByolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "entitlement_key", "AAAAA-BBBBB-CCCCC-DDDDD-HHHHH"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_type", "VSAN"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_term_end", "2028-03-24T01:23:45.678Z"),
				resource.TestCheckResourceAttr(resourceName, "time_term_start", "2026-03-24T01:23:45.678Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "total_units", "1100"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_byols", "test_byols", acctest.Optional, acctest.Update, OcvpByolDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpByolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Optional, acctest.Update, OcvpByolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "available_units_greater_than_or_equal_to", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "byol_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "software_type", "VSAN"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "byol_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "byol_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Required, acctest.Create, OcvpByolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpByolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "byol_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_units"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlement_key", "AAAAA-BBBBB-CCCCC-DDDDD-HHHHH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_type", "VSAN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_term_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_term_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "total_units", "1100"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpByolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpByolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ByolClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_byol" {
			noResourceFound = false
			request := oci_ocvp.GetByolRequest{}

			tmp := rs.Primary.ID
			request.ByolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetByol(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.ByolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpByol") {
		resource.AddTestSweepers("OcvpByol", &resource.Sweeper{
			Name:         "OcvpByol",
			Dependencies: acctest.DependencyGraph["byol"],
			F:            sweepOcvpByolResource,
		})
	}
}

func sweepOcvpByolResource(compartment string) error {
	byolClient := acctest.GetTestClients(&schema.ResourceData{}).ByolClient()
	byolIds, err := getOcvpByolIds(compartment)
	if err != nil {
		return err
	}
	for _, byolId := range byolIds {
		if ok := acctest.SweeperDefaultResourceId[byolId]; !ok {
			deleteByolRequest := oci_ocvp.DeleteByolRequest{}

			deleteByolRequest.ByolId = &byolId

			deleteByolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := byolClient.DeleteByol(context.Background(), deleteByolRequest)
			if error != nil {
				fmt.Printf("Error deleting Byol %s %s, It is possible that the resource is already deleted. Please verify manually \n", byolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &byolId, OcvpByolSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpByolSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpByolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ByolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	byolClient := acctest.GetTestClients(&schema.ResourceData{}).ByolClient()

	listByolsRequest := oci_ocvp.ListByolsRequest{}
	listByolsRequest.CompartmentId = &compartmentId
	listByolsRequest.LifecycleState = oci_ocvp.ByolLifecycleStateActive
	listByolsResponse, err := byolClient.ListByols(context.Background(), listByolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Byol list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, byol := range listByolsResponse.Items {
		id := *byol.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ByolId", id)
	}
	return resourceIds, nil
}

func OcvpByolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if byolResponse, ok := response.Response.(oci_ocvp.GetByolResponse); ok {
		return byolResponse.LifecycleState != oci_ocvp.ByolLifecycleStateDeleted
	}
	return false
}

func OcvpByolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ByolClient().GetByol(context.Background(), oci_ocvp.GetByolRequest{
		ByolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
