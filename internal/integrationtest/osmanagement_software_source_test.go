// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SoftwareSourceRequiredOnlyResource = SoftwareSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Required, acctest.Create, softwareSourceRepresentation)

	SoftwareSourceResourceConfig = SoftwareSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Update, softwareSourceRepresentation)

	softwareSourceSingularDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_software_source.id}`},
	}

	softwareSourceDisplayName       = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName")
	softwareSourceUpdateDisplayName = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName2")

	softwareSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: softwareSourceDisplayName, Update: softwareSourceUpdateDisplayName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: softwareSourceDataSourceFilterRepresentation}}
	softwareSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_osmanagement_software_source.test_software_source.id}`}},
	}

	softwareSourceRepresentation = map[string]interface{}{
		"arch_type":        acctest.Representation{RepType: acctest.Required, Create: `IA_32`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: softwareSourceDisplayName, Update: softwareSourceUpdateDisplayName},
		"checksum_type":    acctest.Representation{RepType: acctest.Optional, Create: `SHA1`, Update: `SHA256`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"maintainer_email": acctest.Representation{RepType: acctest.Optional, Create: `maintainerEmail`, Update: `maintainerEmail2`},
		"maintainer_name":  acctest.Representation{RepType: acctest.Optional, Create: `maintainerName`, Update: `maintainerName2`},
		"maintainer_phone": acctest.Representation{RepType: acctest.Optional, Create: `maintainerPhone`, Update: `maintainerPhone2`},
	}

	SoftwareSourceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementSoftwareSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementSoftwareSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_osmanagement_software_source.test_software_source"
	datasourceName := "data.oci_osmanagement_software_sources.test_software_sources"
	singularDatasourceName := "data.oci_osmanagement_software_source.test_software_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SoftwareSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Create, softwareSourceRepresentation), "osmanagement", "softwareSource", t)

	acctest.ResourceTest(t, testAccCheckOsmanagementSoftwareSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SoftwareSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Required, acctest.Create, softwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "IA_32"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", softwareSourceDisplayName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SoftwareSourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SoftwareSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Create, softwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "IA_32"),
				resource.TestCheckResourceAttr(resourceName, "checksum_type", "SHA1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", softwareSourceDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_email", "maintainerEmail"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_name", "maintainerName"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_phone", "maintainerPhone"),
				resource.TestCheckResourceAttrSet(resourceName, "repo_type"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SoftwareSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(softwareSourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "IA_32"),
				resource.TestCheckResourceAttr(resourceName, "checksum_type", "SHA1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", softwareSourceDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_email", "maintainerEmail"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_name", "maintainerName"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_phone", "maintainerPhone"),
				resource.TestCheckResourceAttrSet(resourceName, "repo_type"),

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
			Config: config + compartmentIdVariableStr + SoftwareSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Update, softwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "IA_32"),
				resource.TestCheckResourceAttr(resourceName, "checksum_type", "SHA256"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", softwareSourceUpdateDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_email", "maintainerEmail2"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_name", "maintainerName2"),
				resource.TestCheckResourceAttr(resourceName, "maintainer_phone", "maintainerPhone2"),
				resource.TestCheckResourceAttrSet(resourceName, "repo_type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_software_sources", "test_software_sources", acctest.Optional, acctest.Update, softwareSourceDataSourceRepresentation) +
				compartmentIdVariableStr + SoftwareSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Optional, acctest.Update, softwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", softwareSourceUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "software_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "software_sources.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "software_sources.0.display_name", softwareSourceUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "software_sources.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_sources.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_sources.0.packages"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_sources.0.repo_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_sources.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_sources.0.status"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_software_source", "test_software_source", acctest.Required, acctest.Create, softwareSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SoftwareSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "IA_32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "checksum_type", "SHA256"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", softwareSourceUpdateDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintainer_email", "maintainerEmail2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintainer_name", "maintainerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintainer_phone", "maintainerPhone2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "packages"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repo_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SoftwareSourceResourceConfig,
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

func testAccCheckOsmanagementSoftwareSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_osmanagement_software_source" {
			noResourceFound = false
			request := oci_osmanagement.GetSoftwareSourceRequest{}

			tmp := rs.Primary.ID
			request.SoftwareSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "osmanagement")

			response, err := client.GetSoftwareSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_osmanagement.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsmanagementSoftwareSource") {
		resource.AddTestSweepers("OsmanagementSoftwareSource", &resource.Sweeper{
			Name:         "OsmanagementSoftwareSource",
			Dependencies: acctest.DependencyGraph["softwareSource"],
			F:            sweepOsmanagementSoftwareSourceResource,
		})
	}
}

func sweepOsmanagementSoftwareSourceResource(compartment string) error {
	osManagementClient := acctest.GetTestClients(&schema.ResourceData{}).OsManagementClient()
	softwareSourceIds, err := getSoftwareSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, softwareSourceId := range softwareSourceIds {
		if ok := acctest.SweeperDefaultResourceId[softwareSourceId]; !ok {
			deleteSoftwareSourceRequest := oci_osmanagement.DeleteSoftwareSourceRequest{}

			deleteSoftwareSourceRequest.SoftwareSourceId = &softwareSourceId

			deleteSoftwareSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "osmanagement")
			_, error := osManagementClient.DeleteSoftwareSource(context.Background(), deleteSoftwareSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting SoftwareSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", softwareSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &softwareSourceId, softwareSourceSweepWaitCondition, time.Duration(3*time.Minute),
				softwareSourceSweepResponseFetchOperation, "osmanagement", true)
		}
	}
	return nil
}

func getSoftwareSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SoftwareSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	osManagementClient := acctest.GetTestClients(&schema.ResourceData{}).OsManagementClient()

	listSoftwareSourcesRequest := oci_osmanagement.ListSoftwareSourcesRequest{}
	listSoftwareSourcesRequest.CompartmentId = &compartmentId
	listSoftwareSourcesRequest.LifecycleState = oci_osmanagement.ListSoftwareSourcesLifecycleStateActive
	listSoftwareSourcesResponse, err := osManagementClient.ListSoftwareSources(context.Background(), listSoftwareSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SoftwareSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, softwareSource := range listSoftwareSourcesResponse.Items {
		id := *softwareSource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SoftwareSourceId", id)
	}
	return resourceIds, nil
}

func softwareSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if softwareSourceResponse, ok := response.Response.(oci_osmanagement.GetSoftwareSourceResponse); ok {
		return softwareSourceResponse.LifecycleState != oci_osmanagement.LifecycleStatesDeleted
	}
	return false
}

func softwareSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OsManagementClient().GetSoftwareSource(context.Background(), oci_osmanagement.GetSoftwareSourceRequest{
		SoftwareSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
