// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/common"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const (
	DefinedTagsDependencies = `
variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}
`
)

var (
	TagRequiredOnlyResource = TagResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Required, Create, tagRepresentation)

	tagDataSourceRepresentation = map[string]interface{}{
		"tag_namespace_id": Representation{repType: Required, create: `${oci_identity_tag_namespace.tag-namespace1.id}`},
		"state":            Representation{repType: Optional, create: `AVAILABLE`},
		"filter":           RepresentationGroup{Required, tagDataSourceFilterRepresentation}}
	tagDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_tag.test_tag.id}`}},
	}

	tagRepresentation = map[string]interface{}{
		"description":      Representation{repType: Required, create: `This tag will show the cost center that will be used for billing of associated resources.`, update: `description2`},
		"name":             Representation{repType: Required, create: `TFTestTag`},
		"tag_namespace_id": Representation{repType: Required, create: `${oci_identity_tag_namespace.tag-namespace1.id}`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_cost_tracking": Representation{repType: Optional, create: `false`, update: `true`},
		"validator":        RepresentationGroup{Optional, tagValidatorRepresentation},
	}
	tagValidatorRepresentation = map[string]interface{}{
		"validator_type": Representation{repType: Required, create: `ENUM`},
		"values":         Representation{repType: Required, create: []string{`value1`, `value2`}},
	}

	TagResourceDependencies = DefinedTagsDependencies
)

func TestIdentityTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag.test_tag"
	datasourceName := "data.oci_identity_tags.test_tags"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		//CheckDestroy: testAccCheckIdentityTagDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Required, Create, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(resourceName, "name", "TFTestTag"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Create, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartment(&resId, &compartmentId); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Update, representationCopyWithRemovedProperties(tagRepresentation, []string{"validator"})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_cost_tracking", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", "TFTestTag"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_identity_tags", "test_tags", Optional, Update, tagDataSourceRepresentation) +
					compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Create, representationCopyWithRemovedProperties(tagRepresentation, []string{"validator"})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespace_id"),

					resource.TestCheckResourceAttr(datasourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.is_cost_tracking", "false"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.is_retired", "false"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.name", "TFTestTag"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
				),
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
	client := testAccProvider.Meta().(*OracleClients).identityClient
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("IdentityTag") {
		resource.AddTestSweepers("IdentityTag", &resource.Sweeper{
			Name:         "IdentityTag",
			Dependencies: DependencyGraph["tag"],
			F:            sweepIdentityTagResource,
		})
	}
}

func sweepIdentityTagResource(compartment string) error {
	// prevent tag deletion when testing, as its a time consuming and sequential operation permitted one per tenancy.
	importIfExists, _ := strconv.ParseBool(getEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		return nil
	}

	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	tagIds, err := getTagIds(compartment)
	if err != nil {
		return err
	}
	for _, tagId := range tagIds {
		if ok := SweeperDefaultResourceId[tagId]; !ok {
			deleteTagRequest := oci_identity.DeleteTagRequest{}

			deleteTagRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeleteTag(context.Background(), deleteTagRequest)
			if error != nil {
				fmt.Printf("Error deleting Tag %s %s, It is possible that the resource is already deleted. Please verify manually \n", tagId, error)
				continue
			}
			waitTillCondition(testAccProvider, &tagId, tagSweepWaitCondition, time.Duration(3*time.Minute),
				tagSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getTagIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "TagId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "TagId", id)
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

func tagSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetTag(context.Background(), oci_identity.GetTagRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
