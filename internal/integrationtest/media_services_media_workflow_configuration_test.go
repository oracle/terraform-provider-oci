// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MediaServicesMediaWorkflowConfigurationRequiredOnlyResource = MediaServicesMediaWorkflowConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Required, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation)

	MediaServicesMediaWorkflowConfigurationResourceConfig = MediaServicesMediaWorkflowConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowConfigurationRepresentation)

	MediaServicesMediaServicesMediaWorkflowConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"media_workflow_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`},
	}

	MediaServicesMediaServicesMediaWorkflowConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesMediaWorkflowConfigurationDataSourceFilterRepresentation}}
	MediaServicesMediaWorkflowConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`}},
	}

	MediaServicesMediaWorkflowConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"parameters":       acctest.Representation{RepType: acctest.Required, Create: `{\"storage\":{\"inputbucket\":\"myvideobucket\",\"outputbucket\":\"myoutputBucket\"}}`, Update: `{\"storage\":{\"inputbucket\":\"myvideobucket1\",\"outputbucket\":\"myoutputBucket1\"}}`},
		"is_lock_override": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaWorkflowConfigurationLocksRepresentation},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	ignoreDefinedTagsAndLocks = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`locks`, `defined_tags`, `is_lock_override`}},
	}

	MediaServicesMediaWorkflowConfigurationLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesMediaWorkflowConfigurationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaWorkflowConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaWorkflowConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_media_services_media_workflow_configuration.test_media_workflow_configuration"
	datasourceName := "data.oci_media_services_media_workflow_configurations.test_media_workflow_configurations"
	singularDatasourceName := "data.oci_media_services_media_workflow_configuration.test_media_workflow_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesMediaWorkflowConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation), "mediaservices", "mediaWorkflowConfiguration", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesMediaWorkflowConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Required, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket\",\"outputbucket\":\"myoutputBucket\"}}"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket\",\"outputbucket\":\"myoutputBucket\"}}"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MediaServicesMediaWorkflowConfigurationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket\",\"outputbucket\":\"myoutputBucket\"}}"),

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
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket1\",\"outputbucket\":\"myoutputBucket1\"}}"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_configurations", "test_media_workflow_configurations", acctest.Optional, acctest.Update, MediaServicesMediaServicesMediaWorkflowConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket1\",\"outputbucket\":\"myoutputBucket1\"}}"),
				resource.TestCheckResourceAttr(datasourceName, "media_workflow_configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "media_workflow_configuration_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "media_workflow_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters", "{\"storage\":{\"inputbucket\":\"myvideobucket1\",\"outputbucket\":\"myoutputBucket1\"}}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesMediaWorkflowConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesMediaWorkflowConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_media_workflow_configuration" {
			noResourceFound = false
			request := oci_media_services.GetMediaWorkflowConfigurationRequest{}

			tmp := rs.Primary.ID
			request.MediaWorkflowConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetMediaWorkflowConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.MediaWorkflowConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MediaServicesMediaWorkflowConfiguration") {
		resource.AddTestSweepers("MediaServicesMediaWorkflowConfiguration", &resource.Sweeper{
			Name:         "MediaServicesMediaWorkflowConfiguration",
			Dependencies: acctest.DependencyGraph["mediaWorkflowConfiguration"],
			F:            sweepMediaServicesMediaWorkflowConfigurationResource,
		})
	}
}

func sweepMediaServicesMediaWorkflowConfigurationResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	mediaWorkflowConfigurationIds, err := getMediaServicesMediaWorkflowConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, mediaWorkflowConfigurationId := range mediaWorkflowConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[mediaWorkflowConfigurationId]; !ok {
			deleteMediaWorkflowConfigurationRequest := oci_media_services.DeleteMediaWorkflowConfigurationRequest{}

			deleteMediaWorkflowConfigurationRequest.MediaWorkflowConfigurationId = &mediaWorkflowConfigurationId

			deleteMediaWorkflowConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteMediaWorkflowConfiguration(context.Background(), deleteMediaWorkflowConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting MediaWorkflowConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", mediaWorkflowConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mediaWorkflowConfigurationId, MediaServicesMediaWorkflowConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesMediaWorkflowConfigurationSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesMediaWorkflowConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MediaWorkflowConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listMediaWorkflowConfigurationsRequest := oci_media_services.ListMediaWorkflowConfigurationsRequest{}
	listMediaWorkflowConfigurationsRequest.CompartmentId = &compartmentId
	listMediaWorkflowConfigurationsRequest.LifecycleState = oci_media_services.MediaWorkflowLifecycleStateActive
	listMediaWorkflowConfigurationsResponse, err := mediaServicesClient.ListMediaWorkflowConfigurations(context.Background(), listMediaWorkflowConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MediaWorkflowConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mediaWorkflowConfiguration := range listMediaWorkflowConfigurationsResponse.Items {
		id := *mediaWorkflowConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MediaWorkflowConfigurationId", id)
	}
	return resourceIds, nil
}

func MediaServicesMediaWorkflowConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mediaWorkflowConfigurationResponse, ok := response.Response.(oci_media_services.GetMediaWorkflowConfigurationResponse); ok {
		return mediaWorkflowConfigurationResponse.LifecycleState != oci_media_services.MediaWorkflowConfigurationLifecycleStateDeleted
	}
	return false
}

func MediaServicesMediaWorkflowConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetMediaWorkflowConfiguration(context.Background(), oci_media_services.GetMediaWorkflowConfigurationRequest{
		MediaWorkflowConfigurationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
