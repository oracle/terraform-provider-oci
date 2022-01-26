// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	TagRequiredOnlyResource = TagResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Required, acctest.Create, tagRepresentation)

	TagResourceConfigWithoutValidator = TagResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithRemovedProperties(tagRepresentation, []string{"validator"}))

	tagSingularDataSourceRepresentation = map[string]interface{}{
		"tag_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag.test_tag.name}`},
		"tag_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.id}`},
	}

	tagDataSourceRepresentation = map[string]interface{}{
		"tag_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: tagDataSourceFilterRepresentation}}
	tagDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_tag.test_tag.name}`}},
	}

	tagRepresentation = map[string]interface{}{
		"description":      acctest.Representation{RepType: acctest.Required, Create: `This tag will show the cost center that will be used for billing of associated resources.`, Update: `description2`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `TFTestTag`},
		"tag_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_cost_tracking": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"validator":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: tagValidatorRepresentation},
	}
	tagValidatorRepresentation = map[string]interface{}{
		"validator_type": acctest.Representation{RepType: acctest.Required, Create: `ENUM`},
		"values":         acctest.Representation{RepType: acctest.Required, Create: []string{`value1`, `value2`}},
	}

	TagResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag.test_tag"
	datasourceName := "data.oci_identity_tags.test_tags"
	singularDatasourceName := "data.oci_identity_tag.test_tag"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+TagResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Optional, acctest.Create, tagRepresentation), "identity", "tag", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TagResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Required, acctest.Create, tagRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestTag"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TagResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TagResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Optional, acctest.Create, tagRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_cost_tracking", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestTag"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "validator.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.validator_type", "ENUM"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.values.#", "2"),

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
			Config: config + compartmentIdVariableStr + TagResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithRemovedProperties(tagRepresentation, []string{"validator"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_cost_tracking", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestTag"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tags", "test_tags", acctest.Optional, acctest.Update, tagDataSourceRepresentation) +
				compartmentIdVariableStr + TagResourceConfigWithoutValidator,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "tag_namespace_id"),

				resource.TestCheckResourceAttr(datasourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tags.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "tags.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "tags.0.is_cost_tracking", "true"),
				resource.TestCheckResourceAttr(datasourceName, "tags.0.is_retired", "false"),
				resource.TestCheckResourceAttr(datasourceName, "tags.0.name", "TFTestTag"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag", "test_tag", acctest.Required, acctest.Create, tagSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TagResourceConfigWithoutValidator,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_namespace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_cost_tracking", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TFTestTag"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TagResourceConfigWithoutValidator,
		},
		// verify resource import
		{
			Config:                  config,
			ImportStateIdFunc:       getTagCompositeId(resourceName),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getTagCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("tagNamespaces/%s/tags/%s", rs.Primary.Attributes["tag_namespace_id"], rs.Primary.Attributes["name"]), nil
	}
}

func testAccCheckIdentityTagDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_tag" {
			noResourceFound = false
			request := oci_identity.GetTagRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.TagName = &value
			}

			if value, ok := rs.Primary.Attributes["tag_namespace_id"]; ok {
				request.TagNamespaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetTag(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.TagLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IdentityTag") {
		resource.AddTestSweepers("IdentityTag", &resource.Sweeper{
			Name:         "IdentityTag",
			Dependencies: acctest.DependencyGraph["tag"],
			F:            sweepIdentityTagResource,
		})
	}
}

func sweepIdentityTagResource(compartment string) error {
	// prevent tag deletion when testing, as its a time consuming and sequential operation permitted one per tenancy.
	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		return nil
	}

	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()
	tagIds, err := getTagIds(compartment)
	if err != nil {
		return err
	}
	for _, tagId := range tagIds {
		if ok := acctest.SweeperDefaultResourceId[tagId]; !ok {
			deleteTagRequest := oci_identity.DeleteTagRequest{}

			deleteTagRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			_, error := identityClient.DeleteTag(context.Background(), deleteTagRequest)
			if error != nil {
				fmt.Printf("Error deleting Tag %s %s, It is possible that the resource is already deleted. Please verify manually \n", tagId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &tagId, tagSweepWaitCondition, time.Duration(3*time.Minute),
				tagSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getTagIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TagId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()

	listTagsRequest := oci_identity.ListTagsRequest{}
	tagNamespaceIds, error := getTagNamespaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting tagNamespaceId required for Tag resource requests \n")
	}
	for _, tagNamespaceId := range tagNamespaceIds {
		listTagsRequest.TagNamespaceId = &tagNamespaceId

		listTagsRequest.LifecycleState = oci_identity.TagLifecycleStateActive
		listTagsResponse, err := identityClient.ListTags(context.Background(), listTagsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Tag list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, tag := range listTagsResponse.Items {
			id := *tag.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TagId", id)
		}

	}
	return resourceIds, nil
}

func tagSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tagResponse, ok := response.Response.(oci_identity.GetTagResponse); ok {
		return tagResponse.LifecycleState != oci_identity.TagLifecycleStateDeleted
	}
	return false
}

func tagSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IdentityClient().GetTag(context.Background(), oci_identity.GetTagRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
