// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityNetworkSourceRequiredOnlyResource = IdentityNetworkSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Required, acctest.Create, IdentityNetworkSourceRepresentation)

	IdentityNetworkSourceResourceConfig = IdentityNetworkSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Optional, acctest.Update, IdentityNetworkSourceRepresentation)

	IdentityIdentityNetworkSourceSingularDataSourceRepresentation = map[string]interface{}{
		"network_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_network_source.test_network_source.id}`},
	}

	IdentityIdentityNetworkSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `corpnet`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityNetworkSourceDataSourceFilterRepresentation}}
	IdentityNetworkSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_network_source.test_network_source.id}`}},
	}

	IdentityNetworkSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":         acctest.Representation{RepType: acctest.Required, Create: `corporate ip ranges to be used for ip based authorization`, Update: `description2`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `corpnet`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"public_source_list":  acctest.Representation{RepType: acctest.Optional, Create: []string{`128.2.13.5`}, Update: []string{`128.2.13.5`, `128.2.13.6`}},
		"services":            acctest.Representation{RepType: acctest.Optional, Create: []string{`none`}, Update: []string{`all`}},
		"virtual_source_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityNetworkVirtualSourceListRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityNetworkSourceIgnoreChangesRepresentation},
	}

	IdentityNetworkVirtualSourceListRepresentation = map[string]interface{}{
		"vcn_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"ip_ranges": acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.0/16`}},
	}

	IdentityNetworkSourceIgnoreChangesRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	IdentityNetworkSourceResourceDependencies = DefinedTagsDependencies + acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityNetworkSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityNetworkSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_network_source.test_network_source"
	datasourceName := "data.oci_identity_network_sources.test_network_sources"
	singularDatasourceName := "data.oci_identity_network_source.test_network_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityNetworkSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Optional, acctest.Create, IdentityNetworkSourceRepresentation), "identity", "networkSource", t)

	acctest.ResourceTest(t, testAccCheckIdentityNetworkSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityNetworkSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Required, acctest.Create, IdentityNetworkSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "corporate ip ranges to be used for ip based authorization"),
				resource.TestCheckResourceAttr(resourceName, "name", "corpnet"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityNetworkSourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityNetworkSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Optional, acctest.Create, IdentityNetworkSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "corporate ip ranges to be used for ip based authorization"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "corpnet"),
				resource.TestCheckResourceAttr(resourceName, "public_source_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "virtual_source_list.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityNetworkSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Optional, acctest.Update, IdentityNetworkSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "corpnet"),
				resource.TestCheckResourceAttr(resourceName, "public_source_list.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "virtual_source_list.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_network_sources", "test_network_sources", acctest.Optional, acctest.Update, IdentityIdentityNetworkSourceDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityNetworkSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Optional, acctest.Update, IdentityNetworkSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "corpnet"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "network_sources.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_sources.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.name", "corpnet"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.public_source_list.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.services.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_sources.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_sources.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "network_sources.0.virtual_source_list.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Required, acctest.Create, IdentityIdentityNetworkSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityNetworkSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_source_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "corpnet"),
				resource.TestCheckResourceAttr(singularDatasourceName, "public_source_list.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "services.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "virtual_source_list.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + IdentityNetworkSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityNetworkSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_network_source" {
			noResourceFound = false
			request := oci_identity.GetNetworkSourceRequest{}

			tmp := rs.Primary.ID
			request.NetworkSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetNetworkSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.NetworkSourcesLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IdentityNetworkSource") {
		resource.AddTestSweepers("IdentityNetworkSource", &resource.Sweeper{
			Name:         "IdentityNetworkSource",
			Dependencies: acctest.DependencyGraph["networkSource"],
			F:            sweepIdentityNetworkSourceResource,
		})
	}
}

func sweepIdentityNetworkSourceResource(compartment string) error {
	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()
	networkSourceIds, err := getIdentityNetworkSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, networkSourceId := range networkSourceIds {
		if ok := acctest.SweeperDefaultResourceId[networkSourceId]; !ok {
			deleteNetworkSourceRequest := oci_identity.DeleteNetworkSourceRequest{}

			deleteNetworkSourceRequest.NetworkSourceId = &networkSourceId

			deleteNetworkSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			_, error := identityClient.DeleteNetworkSource(context.Background(), deleteNetworkSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkSourceId, IdentityNetworkSourceSweepWaitCondition, time.Duration(3*time.Minute),
				IdentityNetworkSourceSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getIdentityNetworkSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()

	listNetworkSourcesRequest := oci_identity.ListNetworkSourcesRequest{}
	listNetworkSourcesRequest.CompartmentId = &compartmentId
	listNetworkSourcesRequest.LifecycleState = oci_identity.NetworkSourcesLifecycleStateActive
	listNetworkSourcesResponse, err := identityClient.ListNetworkSources(context.Background(), listNetworkSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkSource := range listNetworkSourcesResponse.Items {
		id := *networkSource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkSourceId", id)
	}
	return resourceIds, nil
}

func IdentityNetworkSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkSourceResponse, ok := response.Response.(oci_identity.GetNetworkSourceResponse); ok {
		return networkSourceResponse.LifecycleState != oci_identity.NetworkSourcesLifecycleStateDeleted
	}
	return false
}

func IdentityNetworkSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IdentityClient().GetNetworkSource(context.Background(), oci_identity.GetNetworkSourceRequest{
		NetworkSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
