// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	CpeRequiredOnlyResource = CpeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, cpeRepresentation)

	cpeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: cpeDataSourceFilterRepresentation}}
	cpeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_cpe.test_cpe.id}`}},
	}

	cpeRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ip_address":          acctest.Representation{RepType: acctest.Required, Create: `203.0.113.6`},
		"cpe_device_shape_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes.0.cpe_device_shape_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `MyCpe`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CpeResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreCpeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCpeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_cpe.test_cpe"
	datasourceName := "data.oci_core_cpes.test_cpes"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CpeResourceDependencies+
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Optional, acctest.Create, cpeRepresentation), "core", "cpe", t)

	acctest.ResourceTest(t, testAccCheckCoreCpeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CpeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, cpeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "203.0.113.6"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CpeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CpeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Optional, acctest.Create, cpeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_device_shape_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCpe"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "203.0.113.6"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CpeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(cpeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_device_shape_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCpe"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "203.0.113.6"),

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
			Config: config + compartmentIdVariableStr + CpeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Optional, acctest.Update, cpeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_device_shape_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "203.0.113.6"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpes", "test_cpes", acctest.Optional, acctest.Update, cpeDataSourceRepresentation) +
				compartmentIdVariableStr + CpeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Optional, acctest.Update, cpeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "cpes.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cpes.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.cpe_device_shape_id"),
				resource.TestCheckResourceAttr(datasourceName, "cpes.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cpes.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "cpes.0.ip_address", "203.0.113.6"),
				resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreCpeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cpe" {
			noResourceFound = false
			request := oci_core.GetCpeRequest{}

			tmp := rs.Primary.ID
			request.CpeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetCpe(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("CoreCpe") {
		resource.AddTestSweepers("CoreCpe", &resource.Sweeper{
			Name:         "CoreCpe",
			Dependencies: acctest.DependencyGraph["cpe"],
			F:            sweepCoreCpeResource,
		})
	}
}

func sweepCoreCpeResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	cpeIds, err := getCpeIds(compartment)
	if err != nil {
		return err
	}
	for _, cpeId := range cpeIds {
		if ok := acctest.SweeperDefaultResourceId[cpeId]; !ok {
			deleteCpeRequest := oci_core.DeleteCpeRequest{}

			deleteCpeRequest.CpeId = &cpeId

			deleteCpeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCpe(context.Background(), deleteCpeRequest)
			if error != nil {
				fmt.Printf("Error deleting Cpe %s %s, It is possible that the resource is already deleted. Please verify manually \n", cpeId, error)
				continue
			}
		}
	}
	return nil
}

func getCpeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CpeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listCpesRequest := oci_core.ListCpesRequest{}
	listCpesRequest.CompartmentId = &compartmentId
	listCpesResponse, err := virtualNetworkClient.ListCpes(context.Background(), listCpesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cpe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cpe := range listCpesResponse.Items {
		id := *cpe.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CpeId", id)
	}
	return resourceIds, nil
}
