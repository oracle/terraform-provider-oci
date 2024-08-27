// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsGroupRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation)

	IdentityDomainsGroupResourceConfig = IdentityDomainsGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Update, IdentityDomainsGroupRepresentation)

	IdentityDomainsIdentityDomainsGroupSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsGroupDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"group_count":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"group_filter":   acctest.Representation{RepType: acctest.Optional, Create: `displayName sw \"displayName2\"`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsGroupRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"idcs_endpoint":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":                 acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:core:2.0:Group`}},
		"attribute_sets":          acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":             acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"force_delete":            acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"members":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupMembersRepresentation},
		"non_unique_display_name": acctest.Representation{RepType: acctest.Optional, Create: `nonUniqueDisplayName`, Update: `nonUniqueDisplayName2`},
		"tags":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextension_oci_tags":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensiongroup_group":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensiongroupGroupRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionposix_group":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionposixGroupRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionrequestable_group": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionrequestableGroupRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsGroup},
	}

	ignoreChangeForIdentityDomainsGroup = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags`,
			`schemas`,
		}},
	}
	IdentityDomainsGroupMembersRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}
	IdentityDomainsGroupTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation = map[string]interface{}{
		"defined_tags":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation},
		"freeform_tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensiongroupGroupRepresentation = map[string]interface{}{
		"creation_mechanism": acctest.Representation{RepType: acctest.Optional, Create: `api`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"owners":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensiongroupGroupOwnersRepresentation},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionposixGroupRepresentation = map[string]interface{}{
		"gid_number": acctest.Representation{RepType: acctest.Optional, Create: `500`, Update: `501`},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionrequestableGroupRepresentation = map[string]interface{}{
		"requestable": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag.tag1.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.name}`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `freeformKey`, Update: `freeformKey2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `freeformValue`, Update: `freeformValue2`},
	}
	IdentityDomainsGroupUrnietfparamsscimschemasoracleidcsextensiongroupGroupOwnersRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}

	IdentityDomainsGroupResourceDependencies = DefinedTagsDependencies + TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_group.test_group"
	datasourceName := "data.oci_identity_domains_groups.test_groups"
	singularDatasourceName := "data.oci_identity_domains_group.test_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Create, IdentityDomainsGroupRepresentation), "identitydomains", "group", t)

	print(config + compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Create, IdentityDomainsGroupRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Create, IdentityDomainsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "members.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.value"),
				resource.TestCheckResourceAttr(resourceName, "non_unique_display_name", "nonUniqueDisplayName"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.creation_mechanism", "api"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.0.value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_group.0.gid_number", "500"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_group.0.requestable", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "groups", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Update, IdentityDomainsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "members.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "members.0.value"),
				resource.TestCheckResourceAttr(resourceName, "non_unique_display_name", "nonUniqueDisplayName2"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.creation_mechanism", "api"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.owners.0.value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionposix_group.0.gid_number", "501"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionrequestable_group.0.requestable", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_groups", "test_groups", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Optional, acctest.Update, IdentityDomainsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "group_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "groups.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "groups.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "non_unique_display_name", "nonUniqueDisplayName2"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensiongroup_group.0.description", "description2"),
			),
		},

		// reset to required only resource for import
		{
			Config: config + compartmentIdVariableStr + TestDomainDependencies,
		},
		{
			Config: config + compartmentIdVariableStr + TestDomainDependencies + IdentityDomainsGroupRequiredOnlyResource,
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_group", "groups"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"resource_type_schema_version",
				"attribute_sets",
				"attributes",
				"force_delete",
				"external_id",
				"non_unique_display_name",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_group" {
			noResourceFound = false
			request := oci_identity_domains.GetGroupRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.GroupId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsGroup") {
		resource.AddTestSweepers("IdentityDomainsGroup", &resource.Sweeper{
			Name:         "IdentityDomainsGroup",
			Dependencies: acctest.DependencyGraph["group"],
			F:            sweepIdentityDomainsGroupResource,
		})
	}
}

func sweepIdentityDomainsGroupResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	groupIds, err := getIdentityDomainsGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, groupId := range groupIds {
		if ok := acctest.SweeperDefaultResourceId[groupId]; !ok {
			deleteGroupRequest := oci_identity_domains.DeleteGroupRequest{}

			deleteGroupRequest.GroupId = &groupId

			deleteGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteGroup(context.Background(), deleteGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting Group %s %s, It is possible that the resource is already deleted. Please verify manually \n", groupId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "GroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listGroupsRequest := oci_identity_domains.ListGroupsRequest{}
	listGroupsResponse, err := identityDomainsClient.ListGroups(context.Background(), listGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Group list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, group := range listGroupsResponse.Resources {
		id := *group.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "GroupId", id)
	}
	return resourceIds, nil
}
