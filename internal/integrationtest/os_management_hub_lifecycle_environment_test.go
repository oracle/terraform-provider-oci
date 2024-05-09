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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubLifecycleEnvironmentRequiredOnlyResource = OsManagementHubLifecycleEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation)

	OsManagementHubLifecycleEnvironmentResourceConfig = OsManagementHubLifecycleEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Update, OsManagementHubLifecycleEnvironmentRepresentation)

	OsManagementHubLifecycleEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"lifecycle_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id}`},
	}

	OsManagementHubLifecycleEnvironmentDataSourceRepresentation = map[string]interface{}{
		"arch_type":                acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}, Update: []string{`displayName2`}},
		"display_name_contains":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"lifecycle_environment_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id}`},
		"location":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"location_not_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: []string{`ON_PREMISE`}},
		"os_family":                acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubLifecycleEnvironmentDataSourceFilterRepresentation}}
	OsManagementHubLifecycleEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id}`}},
	}

	OsManagementHubLifecycleEnvironmentRepresentation = map[string]interface{}{
		"arch_type":      acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"os_family":      acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"stages":         []acctest.RepresentationGroup{{RepType: acctest.Required, Group: OsManagementHubLifecycleEnvironmentStagesRepresentation}, {RepType: acctest.Required, Group: OsManagementHubLifecycleEnvironmentStagesProdRepresentation}},
		"vendor_name":    acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: ignoreDefinedTagsChangesForOsmhLERep},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"location":       acctest.Representation{RepType: acctest.Optional, Create: `OCI_COMPUTE`},
	}
	OsManagementHubLifecycleEnvironmentStagesRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"rank":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		//"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"defined_tags": acctest.Representation{RepType: acctest.Optional, Create: ignoreDefinedTagsChangesForOsmhLERep},
		//"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	OsManagementHubLifecycleEnvironmentStagesProdRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `prod`, Update: `prod2`},
		"rank":         acctest.Representation{RepType: acctest.Required, Create: `2`},
		//"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"defined_tags": acctest.Representation{RepType: acctest.Optional, Create: ignoreDefinedTagsChangesForOsmhLERep},
		//"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ignoreDefinedTagsChangesForOsmhLERep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	OsManagementHubLifecycleEnvironmentResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_os_management_hub_lifecycle_environment.test_lifecycle_environment"
	datasourceName := "data.oci_os_management_hub_lifecycle_environments.test_lifecycle_environments"
	singularDatasourceName := "data.oci_os_management_hub_lifecycle_environment.test_lifecycle_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubLifecycleEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation), "osmanagementhub", "lifecycleEnvironment", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubLifecycleEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "stages.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.rank", "1"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.display_name", "prod"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.rank", "2"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "location", "OCI_COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "stages.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "stages.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.display_name", "displayName"),
				//resource.TestCheckResourceAttr(resourceName, "stages.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.rank", "1"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.display_name", "prod"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.rank", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OsManagementHubLifecycleEnvironmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "location", "OCI_COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "stages.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.freeform_tags.%", "0"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.rank", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Update, OsManagementHubLifecycleEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "location", "OCI_COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "stages.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "stages.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.display_name", "displayName2"),
				//resource.TestCheckResourceAttr(resourceName, "stages.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "stages.0.rank", "1"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.display_name", "prod2"),
				resource.TestCheckResourceAttr(resourceName, "stages.1.rank", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_environments", "test_lifecycle_environments", acctest.Optional, acctest.Update, OsManagementHubLifecycleEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Update, OsManagementHubLifecycleEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "location.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "location_not_equal_to.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_environment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_environment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "location", "OCI_COMPUTE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_ids.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.arch_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.0.display_name", "displayName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "stages.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.lifecycle_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.location"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.0.managed_instance_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.os_family"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.0.rank", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.time_modified"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.0.vendor_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.arch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.1.display_name", "prod2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "stages.1.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.lifecycle_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.os_family"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.1.rank", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.time_modified"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.1.vendor_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_modified"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "ORACLE"),
			),
		},
		// verify resource import
		{
			Config:                  config + OsManagementHubLifecycleEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOsManagementHubLifecycleEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LifecycleEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_lifecycle_environment" {
			noResourceFound = false
			request := oci_os_management_hub.GetLifecycleEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.LifecycleEnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetLifecycleEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsManagementHubLifecycleEnvironment") {
		resource.AddTestSweepers("OsManagementHubLifecycleEnvironment", &resource.Sweeper{
			Name:         "OsManagementHubLifecycleEnvironment",
			Dependencies: acctest.DependencyGraph["lifecycleEnvironment"],
			F:            sweepOsManagementHubLifecycleEnvironmentResource,
		})
	}
}

func sweepOsManagementHubLifecycleEnvironmentResource(compartment string) error {
	lifecycleEnvironmentClient := acctest.GetTestClients(&schema.ResourceData{}).LifecycleEnvironmentClient()
	lifecycleEnvironmentIds, err := getOsManagementHubLifecycleEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, lifecycleEnvironmentId := range lifecycleEnvironmentIds {
		if ok := acctest.SweeperDefaultResourceId[lifecycleEnvironmentId]; !ok {
			deleteLifecycleEnvironmentRequest := oci_os_management_hub.DeleteLifecycleEnvironmentRequest{}

			deleteLifecycleEnvironmentRequest.LifecycleEnvironmentId = &lifecycleEnvironmentId

			deleteLifecycleEnvironmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := lifecycleEnvironmentClient.DeleteLifecycleEnvironment(context.Background(), deleteLifecycleEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting LifecycleEnvironment %s %s, It is possible that the resource is already deleted. Please verify manually \n", lifecycleEnvironmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &lifecycleEnvironmentId, OsManagementHubLifecycleEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubLifecycleEnvironmentSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubLifecycleEnvironmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LifecycleEnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	lifecycleEnvironmentClient := acctest.GetTestClients(&schema.ResourceData{}).LifecycleEnvironmentClient()

	listLifecycleEnvironmentsRequest := oci_os_management_hub.ListLifecycleEnvironmentsRequest{}
	listLifecycleEnvironmentsRequest.CompartmentId = &compartmentId
	listLifecycleEnvironmentsRequest.LifecycleState = oci_os_management_hub.LifecycleEnvironmentLifecycleStateActive
	listLifecycleEnvironmentsResponse, err := lifecycleEnvironmentClient.ListLifecycleEnvironments(context.Background(), listLifecycleEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LifecycleEnvironment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, lifecycleEnvironment := range listLifecycleEnvironmentsResponse.Items {
		id := *lifecycleEnvironment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LifecycleEnvironmentId", id)
	}
	return resourceIds, nil
}

func OsManagementHubLifecycleEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if lifecycleEnvironmentResponse, ok := response.Response.(oci_os_management_hub.GetLifecycleEnvironmentResponse); ok {
		return lifecycleEnvironmentResponse.LifecycleState != oci_os_management_hub.LifecycleEnvironmentLifecycleStateDeleted
	}
	return false
}

func OsManagementHubLifecycleEnvironmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LifecycleEnvironmentClient().GetLifecycleEnvironment(context.Background(), oci_os_management_hub.GetLifecycleEnvironmentRequest{
		LifecycleEnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
