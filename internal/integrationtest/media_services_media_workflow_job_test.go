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
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MediaServicesMediaWorkflowJobRequiredOnlyResource = MediaServicesMediaWorkflowJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Required, acctest.Create, MediaServicesMediaWorkflowJobRepresentation)

	MediaServicesMediaWorkflowJobResourceConfig = MediaServicesMediaWorkflowJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowJobRepresentation)

	MediaServicesMediaServicesMediaWorkflowJobSingularDataSourceRepresentation = map[string]interface{}{
		"media_workflow_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
	}

	MediaServicesMediaServicesMediaWorkflowJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACCEPTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesMediaWorkflowJobDataSourceFilterRepresentation}}

	MediaServicesMediaWorkflowJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_media_workflow_job.test_media_workflow_job.id}`}},
	}

	MediaServicesMediaWorkflowJobRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"workflow_identifier_type":         acctest.Representation{RepType: acctest.Required, Create: `ID`},
		"is_lock_override":                 acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaWorkflowJobLocksRepresentation},
		"media_workflow_configuration_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`}},
		"media_workflow_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow.test_media_workflow.id}`},
		"media_workflow_name":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow.test_media_workflow.display_name}`},
		"parameters":                       acctest.Representation{RepType: acctest.Required, Create: `{\"videos\":{\"inputObject\":\"horseInSuit.mp4\",\"outputObject\":\"iHaveBeenOBOed.mp4\"}}`, Update: `{\"videos\":{\"inputObject\":\"horseInSuit.mp4\",\"outputObject\":\"iHaveBeenOBOed.mp4\"}}`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	MediaServicesMediaWorkflowJobLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesMediaWorkflowJobResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaWorkflowRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Required, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation)
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaWorkflowJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaWorkflowJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_media_services_media_workflow_job.test_media_workflow_job"
	datasourceName := "data.oci_media_services_media_workflow_jobs.test_media_workflow_jobs"
	singularDatasourceName := "data.oci_media_services_media_workflow_job.test_media_workflow_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesMediaWorkflowJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowJobRepresentation), "mediaservices", "mediaWorkflowJob", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Required, acctest.Create, MediaServicesMediaWorkflowJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "workflow_identifier_type", "ID"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_id"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_name"),
				resource.TestCheckResourceAttrSet(resourceName, "parameters"),
				resource.TestCheckResourceAttr(resourceName, "workflow_identifier_type", "ID"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment_for_workflow_job", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MediaServicesMediaWorkflowJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MediaServicesMediaWorkflowJobRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_id"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_name"),
				resource.TestCheckResourceAttrSet(resourceName, "parameters"),
				resource.TestCheckResourceAttr(resourceName, "workflow_identifier_type", "ID"),

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
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_id"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_name"),
				resource.TestCheckResourceAttrSet(resourceName, "parameters"),
				resource.TestCheckResourceAttr(resourceName, "workflow_identifier_type", "ID"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_jobs", "test_media_workflow_jobs", acctest.Optional, acctest.Update, MediaServicesMediaServicesMediaWorkflowJobDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACCEPTED"),

				resource.TestCheckResourceAttr(datasourceName, "media_workflow_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "media_workflow_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "media_workflow_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "outputs.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesMediaWorkflowJobRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"media_workflow_name",
				"workflow_identifier_type",
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MediaServicesMediaWorkflowJob") {
		resource.AddTestSweepers("MediaServicesMediaWorkflowJob", &resource.Sweeper{
			Name:         "MediaServicesMediaWorkflowJob",
			Dependencies: acctest.DependencyGraph["mediaWorkflowJob"],
			F:            sweepMediaServicesMediaWorkflowJobResource,
		})
	}
}

func sweepMediaServicesMediaWorkflowJobResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	mediaWorkflowJobIds, err := getMediaServicesMediaWorkflowJobIds(compartment)
	if err != nil {
		return err
	}
	for _, mediaWorkflowJobId := range mediaWorkflowJobIds {
		if ok := acctest.SweeperDefaultResourceId[mediaWorkflowJobId]; !ok {
			deleteMediaWorkflowJobRequest := oci_media_services.DeleteMediaWorkflowJobRequest{}

			deleteMediaWorkflowJobRequest.MediaWorkflowJobId = &mediaWorkflowJobId

			deleteMediaWorkflowJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteMediaWorkflowJob(context.Background(), deleteMediaWorkflowJobRequest)
			if error != nil {
				fmt.Printf("Error deleting MediaWorkflowJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", mediaWorkflowJobId, error)
				continue
			}
		}
	}
	return nil
}

func getMediaServicesMediaWorkflowJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MediaWorkflowJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listMediaWorkflowJobsRequest := oci_media_services.ListMediaWorkflowJobsRequest{}
	listMediaWorkflowJobsRequest.CompartmentId = &compartmentId
	listMediaWorkflowJobsResponse, err := mediaServicesClient.ListMediaWorkflowJobs(context.Background(), listMediaWorkflowJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MediaWorkflowJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mediaWorkflowJob := range listMediaWorkflowJobsResponse.Items {
		id := *mediaWorkflowJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MediaWorkflowJobId", id)
	}
	return resourceIds, nil
}
