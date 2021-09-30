// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v48/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_identity "github.com/oracle/oci-go-sdk/v48/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TagNamespaceRequiredOnlyResource = TagNamespaceResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Required, Create, tagNamespaceRepresentation)

	tagNamespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"include_subcompartments": Representation{RepType: Optional, Create: `false`},
		"state":                   Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                  RepresentationGroup{Required, tagNamespaceDataSourceFilterRepresentation}}
	tagNamespaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_tag_namespace.test_tag_namespace.id}`}},
	}

	tagNamespaceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"description":    Representation{RepType: Required, Create: `This namespace contains tags that will be used in billing.`, Update: `description2`},
		"name":           Representation{RepType: Required, Create: `BillingTags`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	TagNamespaceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityTagNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_identity_tag_namespace.test_tag_namespace"
	datasourceName := "data.oci_identity_tag_namespaces.test_tag_namespaces"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+TagNamespaceResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Create, tagNamespaceRepresentation), "identity", "tagNamespace", t)

	ResourceTest(t, testAccCheckIdentityTagNamespaceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Required, Create, tagNamespaceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
				resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Create, tagNamespaceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TagNamespaceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Create,
					RepresentationCopyWithNewProperties(tagNamespaceRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Update, tagNamespaceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_identity_tag_namespaces", "test_tag_namespaces", Optional, Update, tagNamespaceDataSourceRepresentation) +
				compartmentIdVariableStr + TagNamespaceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Update, tagNamespaceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "include_subcompartments", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.is_retired", "false"),
				resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.name", "BillingTags"),
				resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.time_created"),
			),
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

func testAccCheckIdentityTagNamespaceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_tag_namespace" {
			noResourceFound = false
			request := oci_identity.GetTagNamespaceRequest{}

			tmp := rs.Primary.ID
			request.TagNamespaceId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetTagNamespace(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.TagNamespaceLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("IdentityTagNamespace") {
		resource.AddTestSweepers("IdentityTagNamespace", &resource.Sweeper{
			Name:         "IdentityTagNamespace",
			Dependencies: DependencyGraph["tagNamespace"],
			F:            sweepIdentityTagNamespaceResource,
		})
	}
}

func sweepIdentityTagNamespaceResource(compartment string) error {
	// prevent tag deletion when testing, as its a time consuming and sequential operation permitted one per tenancy.
	importIfExists, _ := strconv.ParseBool(getEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		return nil
	}

	identityClient := GetTestClients(&schema.ResourceData{}).identityClient()
	tagNamespaceIds, err := getTagNamespaceIds(compartment)
	if err != nil {
		return err
	}

	// clean all tags in namespaces
	err = sweepIdentityTagResource(compartment)
	if err != nil {
		return err
	}

	for _, tagNamespaceId := range tagNamespaceIds {
		if ok := SweeperDefaultResourceId[tagNamespaceId]; !ok {
			deleteTagNamespaceRequest := oci_identity.DeleteTagNamespaceRequest{}

			deleteTagNamespaceRequest.TagNamespaceId = &tagNamespaceId

			deleteTagNamespaceRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")
			_, error := identityClient.DeleteTagNamespace(context.Background(), deleteTagNamespaceRequest)
			if error != nil {
				fmt.Printf("Error deleting TagNamespace %s %s, It is possible that the resource is already deleted. Please verify manually \n", tagNamespaceId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &tagNamespaceId, tagNamespaceSweepWaitCondition, time.Duration(3*time.Minute),
				tagNamespaceSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getTagNamespaceIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "TagNamespaceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient()

	listTagNamespacesRequest := oci_identity.ListTagNamespacesRequest{}
	listTagNamespacesRequest.CompartmentId = &compartmentId
	listTagNamespacesRequest.LifecycleState = oci_identity.TagNamespaceLifecycleStateActive
	listTagNamespacesResponse, err := identityClient.ListTagNamespaces(context.Background(), listTagNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TagNamespace list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tagNamespace := range listTagNamespacesResponse.Items {
		id := *tagNamespace.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "TagNamespaceId", id)
	}
	return resourceIds, nil
}

func tagNamespaceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tagNamespaceResponse, ok := response.Response.(oci_identity.GetTagNamespaceResponse); ok {
		return tagNamespaceResponse.LifecycleState != oci_identity.TagNamespaceLifecycleStateDeleted
	}
	return false
}

func tagNamespaceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient().GetTagNamespace(context.Background(), oci_identity.GetTagNamespaceRequest{
		TagNamespaceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
