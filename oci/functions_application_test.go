// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_functions "github.com/oracle/oci-go-sdk/v25/functions"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApplicationRequiredOnlyResource = ApplicationResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation)

	ApplicationResourceConfig = ApplicationResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Optional, Update, applicationRepresentation)

	applicationSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": Representation{repType: Required, create: `${oci_functions_application.test_application.id}`},
	}

	applicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`},
		"id":             Representation{repType: Optional, create: `${oci_functions_application.test_application.id}`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, applicationDataSourceFilterRepresentation}}
	applicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_functions_application.test_application.id}`}},
	}

	applicationRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`},
		"subnet_ids":     Representation{repType: Required, create: []string{`${oci_core_subnet.test_subnet.id}`}},
		"config":         Representation{repType: Optional, create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ApplicationResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestFunctionsApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsApplicationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_functions_application.test_application"
	datasourceName := "data.oci_functions_applications.test_applications"
	singularDatasourceName := "data.oci_functions_application.test_application"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFunctionsApplicationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApplicationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_application", "test_application", Optional, Create, applicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_application", "test_application", Optional, Create,
						representationCopyWithNewProperties(applicationRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_application", "test_application", Optional, Update, applicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_functions_applications", "test_applications", Optional, Update, applicationDataSourceRepresentation) +
					compartmentIdVariableStr + ApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_application", "test_application", Optional, Update, applicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
					//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "applications.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "applications.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "applications.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_updated"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApplicationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ApplicationResourceConfig,
			},
			// verify resource import
			{
				Config:                  config + generateResourceImportConfig("oci_functions_application", "test_application"),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckFunctionsApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).functionsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_functions_application" {
			noResourceFound = false
			request := oci_functions.GetApplicationRequest{}

			tmp := rs.Primary.ID
			request.ApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")

			response, err := client.GetApplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_functions.ApplicationLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("FunctionsApplication") {
		resource.AddTestSweepers("FunctionsApplication", &resource.Sweeper{
			Name:         "FunctionsApplication",
			Dependencies: DependencyGraph["application"],
			F:            sweepFunctionsApplicationResource,
		})
	}
}

func sweepFunctionsApplicationResource(compartment string) error {
	functionsManagementClient := GetTestClients(&schema.ResourceData{}).functionsManagementClient()
	applicationIds, err := getApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, applicationId := range applicationIds {
		if ok := SweeperDefaultResourceId[applicationId]; !ok {
			deleteApplicationRequest := oci_functions.DeleteApplicationRequest{}

			deleteApplicationRequest.ApplicationId = &applicationId

			deleteApplicationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")
			_, error := functionsManagementClient.DeleteApplication(context.Background(), deleteApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Application %s %s, It is possible that the resource is already deleted. Please verify manually \n", applicationId, error)
				continue
			}
			waitTillCondition(testAccProvider, &applicationId, applicationSweepWaitCondition, time.Duration(3*time.Minute),
				applicationSweepResponseFetchOperation, "functions", true)
		}
	}
	return nil
}

func getApplicationIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	functionsManagementClient := GetTestClients(&schema.ResourceData{}).functionsManagementClient()

	listApplicationsRequest := oci_functions.ListApplicationsRequest{}
	listApplicationsRequest.CompartmentId = &compartmentId
	listApplicationsRequest.LifecycleState = oci_functions.ApplicationLifecycleStateActive
	listApplicationsResponse, err := functionsManagementClient.ListApplications(context.Background(), listApplicationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Application list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, application := range listApplicationsResponse.Items {
		id := *application.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ApplicationId", id)
	}
	return resourceIds, nil
}

func applicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if applicationResponse, ok := response.Response.(oci_functions.GetApplicationResponse); ok {
		return applicationResponse.LifecycleState != oci_functions.ApplicationLifecycleStateDeleted
	}
	return false
}

func applicationSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.functionsManagementClient().GetApplication(context.Background(), oci_functions.GetApplicationRequest{
		ApplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
