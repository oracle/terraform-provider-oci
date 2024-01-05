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
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreKnowledgeBaseDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	AdmKnowledgeBaseRequiredOnlyResource = AdmKnowledgeBaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, knowledgeBaseRepresentation)

	AdmKnowledgeBaseResourceConfig = AdmKnowledgeBaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, knowledgeBaseRepresentation)

	AdmknowledgeBaseSingularDataSourceRepresentation = map[string]interface{}{
		"knowledge_base_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_knowledge_base.test_knowledge_base.id}`},
	}

	AdmknowledgeBaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_adm_knowledge_base.test_knowledge_base.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: knowledgeBaseDataSourceFilterRepresentation}}
	knowledgeBaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_adm_knowledge_base.test_knowledge_base.id}`}},
	}

	knowledgeBaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreKnowledgeBaseDefinedTagsChangesRepresentation},
	}

	AdmKnowledgeBaseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: adm/default
func TestAdmKnowledgeBaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAdmKnowledgeBaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_adm_knowledge_base.test_knowledge_base"
	datasourceName := "data.oci_adm_knowledge_bases.test_knowledge_bases"
	singularDatasourceName := "data.oci_adm_knowledge_base.test_knowledge_base"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AdmKnowledgeBaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create, knowledgeBaseRepresentation), "adm", "knowledgeBase", t)

	acctest.ResourceTest(t, testAccCheckAdmKnowledgeBaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AdmKnowledgeBaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, knowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AdmKnowledgeBaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AdmKnowledgeBaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create, knowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AdmKnowledgeBaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(knowledgeBaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + AdmKnowledgeBaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, knowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_knowledge_bases", "test_knowledge_bases", acctest.Optional, acctest.Update, AdmknowledgeBaseDataSourceRepresentation) +
				compartmentIdVariableStr + AdmKnowledgeBaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, knowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "knowledge_base_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "knowledge_base_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, AdmknowledgeBaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AdmKnowledgeBaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "knowledge_base_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AdmKnowledgeBaseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAdmKnowledgeBaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApplicationDependencyManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_adm_knowledge_base" {
			noResourceFound = false
			request := oci_adm.GetKnowledgeBaseRequest{}

			tmp := rs.Primary.ID
			request.KnowledgeBaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")

			response, err := client.GetKnowledgeBase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_adm.KnowledgeBaseLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AdmKnowledgeBase") {
		resource.AddTestSweepers("AdmKnowledgeBase", &resource.Sweeper{
			Name:         "AdmKnowledgeBase",
			Dependencies: acctest.DependencyGraph["knowledgeBase"],
			F:            sweepAdmKnowledgeBaseResource,
		})
	}
}

func sweepAdmKnowledgeBaseResource(compartment string) error {
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()
	knowledgeBaseIds, err := getAdmKnowledgeBaseIds(compartment)
	if err != nil {
		return err
	}
	for _, knowledgeBaseId := range knowledgeBaseIds {
		if ok := acctest.SweeperDefaultResourceId[knowledgeBaseId]; !ok {
			deleteKnowledgeBaseRequest := oci_adm.DeleteKnowledgeBaseRequest{}

			deleteKnowledgeBaseRequest.KnowledgeBaseId = &knowledgeBaseId

			deleteKnowledgeBaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")
			_, error := applicationDependencyManagementClient.DeleteKnowledgeBase(context.Background(), deleteKnowledgeBaseRequest)
			if error != nil {
				fmt.Printf("Error deleting KnowledgeBase %s %s, It is possible that the resource is already deleted. Please verify manually \n", knowledgeBaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &knowledgeBaseId, AdmknowledgeBasesSweepWaitCondition, time.Duration(3*time.Minute),
				AdmknowledgeBasesSweepResponseFetchOperation, "adm", true)
		}
	}
	return nil
}

func getAdmKnowledgeBaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KnowledgeBaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()

	listKnowledgeBasesRequest := oci_adm.ListKnowledgeBasesRequest{}
	listKnowledgeBasesRequest.CompartmentId = &compartmentId
	listKnowledgeBasesRequest.LifecycleState = oci_adm.KnowledgeBaseLifecycleStateActive
	listKnowledgeBasesResponse, err := applicationDependencyManagementClient.ListKnowledgeBases(context.Background(), listKnowledgeBasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KnowledgeBase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, knowledgeBase := range listKnowledgeBasesResponse.Items {
		id := *knowledgeBase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KnowledgeBaseId", id)
	}
	return resourceIds, nil
}

func AdmknowledgeBasesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if knowledgeBaseResponse, ok := response.Response.(oci_adm.GetKnowledgeBaseResponse); ok {
		return knowledgeBaseResponse.LifecycleState != oci_adm.KnowledgeBaseLifecycleStateDeleted
	}
	return false
}

func AdmknowledgeBasesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ApplicationDependencyManagementClient().GetKnowledgeBase(context.Background(), oci_adm.GetKnowledgeBaseRequest{
		KnowledgeBaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
