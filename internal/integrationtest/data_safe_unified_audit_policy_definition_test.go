// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeUnifiedAuditPolicyDefinitionRequiredOnlyResource = DataSafeUnifiedAuditPolicyDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Required, acctest.Create, DataSafeUnifiedAuditPolicyDefinitionRepresentation)

	DataSafeUnifiedAuditPolicyDefinitionResourceConfig = DataSafeUnifiedAuditPolicyDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyDefinitionRepresentation)

	DataSafeUnifiedAuditPolicyDefinitionSingularDataSourceRepresentation = map[string]interface{}{
		"unified_audit_policy_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${var.unified_audit_policy_definition_id}`},
	}

	DataSafeUnifiedAuditPolicyDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                       acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_seeded":                          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"unified_audit_policy_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${var.unified_audit_policy_definition_id}`},
		"filter":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeUnifiedAuditPolicyDefinitionDataSourceFilterRepresentation}}
	DataSafeUnifiedAuditPolicyDefinitionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.unified_audit_policy_definition_id}`}},
	}

	DataSafeUnifiedAuditPolicyDefinitionRepresentation = map[string]interface{}{
		"unified_audit_policy_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${var.unified_audit_policy_definition_id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreUnifiedAuditPolicyDefinitionTagsChangesRep},
	}

	ignoreUnifiedAuditPolicyDefinitionTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeUnifiedAuditPolicyDefinitionResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnifiedAuditPolicyDefinitionResource_basic(t *testing.T) {
	t.Skip("Skipping this test as there is no direct create for unified audit policy definition")
	httpreplay.SetScenario("TestDataSafeUnifiedAuditPolicyDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	unifiedAuditPolicyDefinitionId := utils.GetEnvSettingWithBlankDefault("unified_audit_policy_definition_id")
	unifiedAuditPolicyDefinitionIdVariableStr := fmt.Sprintf("variable \"unified_audit_policy_definition_id\" { default = \"%s\" }\n", unifiedAuditPolicyDefinitionId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition"
	datasourceName := "data.oci_data_safe_unified_audit_policy_definitions.test_unified_audit_policy_definitions"
	singularDatasourceName := "data.oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+unifiedAuditPolicyDefinitionIdVariableStr+DataSafeUnifiedAuditPolicyDefinitionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataSafeUnifiedAuditPolicyDefinitionRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "unifiedAuditPolicyDefinition", t)

	acctest.ResourceTest(t, testAccCheckDataSafeUnifiedAuditPolicyDefinitionDestroy, []resource.TestStep{
		// verify Change compartment
		{
			Config: config + compartmentIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + compartmentIdUVariableStr + DataSafeUnifiedAuditPolicyDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeUnifiedAuditPolicyDefinitionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeUnifiedAuditPolicyDefinitionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definitions", "test_unified_audit_policy_definitions", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyDefinitionDataSourceRepresentation) +
				compartmentIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_seeded", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "unified_audit_policy_definition_id"),

				resource.TestCheckResourceAttr(datasourceName, "unified_audit_policy_definition_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "unified_audit_policy_definition_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_unified_audit_policy_definition", "test_unified_audit_policy_definition", acctest.Required, acctest.Create, DataSafeUnifiedAuditPolicyDefinitionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unified_audit_policy_definition_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_seeded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_definition_statement"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeUnifiedAuditPolicyDefinitionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`unified_audit_policy_definition_id`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeUnifiedAuditPolicyDefinitionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_unified_audit_policy_definition" {
			noResourceFound = false
			request := oci_data_safe.GetUnifiedAuditPolicyDefinitionRequest{}

			tmp := rs.Primary.ID
			request.UnifiedAuditPolicyDefinitionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetUnifiedAuditPolicyDefinition(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeUnifiedAuditPolicyDefinition") {
		resource.AddTestSweepers("DataSafeUnifiedAuditPolicyDefinition", &resource.Sweeper{
			Name:         "DataSafeUnifiedAuditPolicyDefinition",
			Dependencies: acctest.DependencyGraph["unifiedAuditPolicyDefinition"],
			F:            sweepDataSafeUnifiedAuditPolicyDefinitionResource,
		})
	}
}

func sweepDataSafeUnifiedAuditPolicyDefinitionResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	unifiedAuditPolicyDefinitionIds, err := getDataSafeUnifiedAuditPolicyDefinitionIds(compartment)
	if err != nil {
		return err
	}
	for _, unifiedAuditPolicyDefinitionId := range unifiedAuditPolicyDefinitionIds {
		if ok := acctest.SweeperDefaultResourceId[unifiedAuditPolicyDefinitionId]; !ok {
			deleteUnifiedAuditPolicyDefinitionRequest := oci_data_safe.DeleteUnifiedAuditPolicyDefinitionRequest{}

			deleteUnifiedAuditPolicyDefinitionRequest.UnifiedAuditPolicyDefinitionId = &unifiedAuditPolicyDefinitionId

			deleteUnifiedAuditPolicyDefinitionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteUnifiedAuditPolicyDefinition(context.Background(), deleteUnifiedAuditPolicyDefinitionRequest)
			if error != nil {
				fmt.Printf("Error deleting UnifiedAuditPolicyDefinition %s %s, It is possible that the resource is already deleted. Please verify manually \n", unifiedAuditPolicyDefinitionId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeUnifiedAuditPolicyDefinitionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UnifiedAuditPolicyDefinitionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listUnifiedAuditPolicyDefinitionsRequest := oci_data_safe.ListUnifiedAuditPolicyDefinitionsRequest{}
	listUnifiedAuditPolicyDefinitionsRequest.CompartmentId = &compartmentId
	listUnifiedAuditPolicyDefinitionsResponse, err := dataSafeClient.ListUnifiedAuditPolicyDefinitions(context.Background(), listUnifiedAuditPolicyDefinitionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UnifiedAuditPolicyDefinition list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, unifiedAuditPolicyDefinition := range listUnifiedAuditPolicyDefinitionsResponse.Items {
		id := *unifiedAuditPolicyDefinition.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UnifiedAuditPolicyDefinitionId", id)
	}
	return resourceIds, nil
}
