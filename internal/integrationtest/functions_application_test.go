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
	oci_functions "github.com/oracle/oci-go-sdk/v56/functions"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApplicationRequiredOnlyResource = ApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation)

	ApplicationResourceConfig = ApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Update, applicationRepresentation)

	applicationSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
	}

	applicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: applicationDisplayName},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_application.test_application.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: applicationDataSourceFilterRepresentation}}
	applicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_functions_application.test_application.id}`}},
	}

	applicationDisplayName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	applicationRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: applicationDisplayName},
		"subnet_ids":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_subnet.id}`}},
		"config":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_policy_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: applicationImagePolicyConfigRepresentation},
		"network_security_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}},
		"syslog_url":                 acctest.Representation{RepType: acctest.Optional, Create: `tcp://syslog.test:80`, Update: `tcp://syslog2.test:80`},
		"trace_config":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: applicationTraceConfigRepresentation},
	}
	applicationImagePolicyConfigRepresentation = map[string]interface{}{
		"is_policy_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"key_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: applicationImagePolicyConfigKeyDetailsRepresentation},
	}
	applicationTraceConfigRepresentation = map[string]interface{}{
		"domain_id":  acctest.Representation{RepType: acctest.Optional, Create: "${oci_apm_apm_domain.test_apm_domain.id}"},
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	applicationImagePolicyConfigKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}

	ApplicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group1", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

// issue-routing-tag: functions/default
func TestFunctionsApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_functions_application.test_application"
	datasourceName := "data.oci_functions_applications.test_applications"
	singularDatasourceName := "data.oci_functions_application.test_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Create, applicationRepresentation), "functions", "application", t)

	acctest.ResourceTest(t, testAccCheckFunctionsApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Create, applicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "syslog_url", "tcp://syslog.test:80"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "trace_config.0.domain_id"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(applicationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "syslog_url", "tcp://syslog.test:80"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "trace_config.0.domain_id"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + ApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Update, applicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "syslog_url", "tcp://syslog2.test:80"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "trace_config.0.domain_id"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_applications", "test_applications", acctest.Optional, acctest.Update, applicationDataSourceRepresentation) +
				compartmentIdVariableStr + ApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Optional, acctest.Update, applicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", applicationDisplayName),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "applications.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.subnet_ids.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.trace_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.trace_config.0.domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.trace_config.0.is_enabled", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", applicationDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "syslog_url", "tcp://syslog2.test:80"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.0.is_enabled", "true"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ApplicationResourceConfig,
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

func testAccCheckFunctionsApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FunctionsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_functions_application" {
			noResourceFound = false
			request := oci_functions.GetApplicationRequest{}

			tmp := rs.Primary.ID
			request.ApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "functions")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("FunctionsApplication") {
		resource.AddTestSweepers("FunctionsApplication", &resource.Sweeper{
			Name:         "FunctionsApplication",
			Dependencies: acctest.DependencyGraph["application"],
			F:            sweepFunctionsApplicationResource,
		})
	}
}

func sweepFunctionsApplicationResource(compartment string) error {
	functionsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FunctionsManagementClient()
	applicationIds, err := getApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, applicationId := range applicationIds {
		if ok := acctest.SweeperDefaultResourceId[applicationId]; !ok {
			deleteApplicationRequest := oci_functions.DeleteApplicationRequest{}

			deleteApplicationRequest.ApplicationId = &applicationId

			deleteApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "functions")
			_, error := functionsManagementClient.DeleteApplication(context.Background(), deleteApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Application %s %s, It is possible that the resource is already deleted. Please verify manually \n", applicationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &applicationId, applicationSweepWaitCondition, time.Duration(3*time.Minute),
				applicationSweepResponseFetchOperation, "functions", true)
		}
	}
	return nil
}

func getApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	functionsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FunctionsManagementClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApplicationId", id)
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

func applicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FunctionsManagementClient().GetApplication(context.Background(), oci_functions.GetApplicationRequest{
		ApplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
