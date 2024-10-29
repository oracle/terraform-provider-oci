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
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentDataSourceRequiredOnlyResource = GenerativeAiAgentDataSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Required, acctest.Create, GenerativeAiAgentDataSourceRepresentation)

	GenerativeAiAgentDataSourceResourceConfig = GenerativeAiAgentDataSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Optional, acctest.Update, GenerativeAiAgentDataSourceRepresentation)

	GenerativeAiAgentDataSourceSingularDataSourceRepresentation = map[string]interface{}{
		"data_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_data_source.test_data_source.id}`},
	}

	GenerativeAiAgentDataSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		//"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayNameDSTest`, Update: `displayNameDSTest2`},
		"knowledge_base_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.knowledgeBaseId_env}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentDataSourceDataSourceFilterRepresentation}}
	GenerativeAiAgentDataSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_data_source.test_data_source.id}`}},
	}

	GenerativeAiAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentDataSourceDataSourceConfigRepresentation},
		"knowledge_base_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.knowledgeBaseId_env}`},
		// "defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `descriptionDSTest`, Update: `descriptionDSTest2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayNameDSTest`, Update: `displayNameDSTest2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentDataSourceDataSourceConfigRepresentation = map[string]interface{}{
		"data_source_config_type": acctest.Representation{RepType: acctest.Required, Create: `OCI_OBJECT_STORAGE`},
		"object_storage_prefixes": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentDataSourceDataSourceConfigObjectStoragePrefixesRepresentation},
	}
	GenerativeAiAgentDataSourceDataSourceConfigObjectStoragePrefixesRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_env}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_namespace_env}`},
		"prefix":    acctest.Representation{RepType: acctest.Optional, Create: `${var.bucket_prefix_env}`},
	}

	GenerativeAiAgentDataSourceResourceDependencies = ``
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentDataSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentDataSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	bucketName := utils.GetEnvSettingWithBlankDefault("bucket_name")
	bucketNameVariableStr := fmt.Sprintf("variable \"bucket_env\" { default = \"%s\" }\n", bucketName)

	bucketNamespace := utils.GetEnvSettingWithBlankDefault("bucket_namespace")
	bucketNamespaceVariableStr := fmt.Sprintf("variable \"bucket_namespace_env\" { default = \"%s\" }\n", bucketNamespace)

	bucketPrefix := utils.GetEnvSettingWithBlankDefault("bucket_prefix")
	bucketPrefixVariableStr := fmt.Sprintf("variable \"bucket_prefix_env\" { default = \"%s\" }\n", bucketPrefix)

	knowledgeBaseId := utils.GetEnvSettingWithBlankDefault("knowledgeBaseId_for_update")
	knowledgeBaseIdUVariableStr := fmt.Sprintf("variable \"knowledgeBaseId_env\" { default = \"%s\" }\n", knowledgeBaseId)

	resourceName := "oci_generative_ai_agent_data_source.test_data_source"
	datasourceName := "data.oci_generative_ai_agent_data_sources.test_data_sources"
	singularDatasourceName := "data.oci_generative_ai_agent_data_source.test_data_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+knowledgeBaseIdUVariableStr+bucketPrefixVariableStr+bucketNamespaceVariableStr+bucketNameVariableStr+compartmentIdVariableStr+GenerativeAiAgentDataSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Optional, acctest.Create, GenerativeAiAgentDataSourceRepresentation), "generativeaiagent", "dataSource", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentDataSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Required, acctest.Create, GenerativeAiAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.data_source_config_type", "OCI_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.0.bucket", bucketName),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.0.namespace", bucketNamespace),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Optional, acctest.Create, GenerativeAiAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.data_source_config_type", "OCI_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.0.bucket", bucketName),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.0.namespace", bucketNamespace),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.0.prefix", bucketPrefix),
				resource.TestCheckResourceAttr(resourceName, "description", "descriptionDSTest"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameDSTest"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify updates to updatable parameters
		{
			Config: config + knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Optional, acctest.Update, GenerativeAiAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.data_source_config_type", "OCI_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_config.0.object_storage_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "descriptionDSTest2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameDSTest2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_data_sources", "test_data_sources", acctest.Optional, acctest.Update, GenerativeAiAgentDataSourceDataSourceRepresentation) +
				knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Optional, acctest.Update, GenerativeAiAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "data_source_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_source_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_data_source", "test_data_source", acctest.Required, acctest.Create, GenerativeAiAgentDataSourceSingularDataSourceRepresentation) +
				knowledgeBaseIdUVariableStr + bucketPrefixVariableStr + bucketNamespaceVariableStr + bucketNameVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_source_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_config.0.data_source_config_type", "OCI_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_config.0.object_storage_prefixes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "descriptionDSTest2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameDSTest2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentDataSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentDataSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_data_source" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetDataSourceRequest{}

			tmp := rs.Primary.ID
			request.DataSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetDataSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.DataSourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentDataSource") {
		resource.AddTestSweepers("GenerativeAiAgentDataSource", &resource.Sweeper{
			Name:         "GenerativeAiAgentDataSource",
			Dependencies: acctest.DependencyGraph["dataSource"],
			F:            sweepGenerativeAiAgentDataSourceResource,
		})
	}
}

func sweepGenerativeAiAgentDataSourceResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	dataSourceIds, err := getGenerativeAiAgentDataSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, dataSourceId := range dataSourceIds {
		if ok := acctest.SweeperDefaultResourceId[dataSourceId]; !ok {
			deleteDataSourceRequest := oci_generative_ai_agent.DeleteDataSourceRequest{}

			deleteDataSourceRequest.DataSourceId = &dataSourceId

			deleteDataSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteDataSource(context.Background(), deleteDataSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting DataSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataSourceId, GenerativeAiAgentDataSourceSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentDataSourceSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentDataSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listDataSourcesRequest := oci_generative_ai_agent.ListDataSourcesRequest{}
	listDataSourcesRequest.CompartmentId = &compartmentId
	listDataSourcesRequest.LifecycleState = oci_generative_ai_agent.DataSourceLifecycleStateActive
	listDataSourcesResponse, err := generativeAiAgentClient.ListDataSources(context.Background(), listDataSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataSource := range listDataSourcesResponse.Items {
		id := *dataSource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataSourceId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentDataSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataSourceResponse, ok := response.Response.(oci_generative_ai_agent.GetDataSourceResponse); ok {
		return dataSourceResponse.LifecycleState != oci_generative_ai_agent.DataSourceLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentDataSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetDataSource(context.Background(), oci_generative_ai_agent.GetDataSourceRequest{
		DataSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
