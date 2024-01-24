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

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsDynamicResourceGroupRequiredOnlyResource = IdentityDomainsDynamicResourceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Required, acctest.Create, IdentityDomainsDynamicResourceGroupRepresentation)

	IdentityDomainsDynamicResourceGroupResourceConfig = IdentityDomainsDynamicResourceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Update, IdentityDomainsDynamicResourceGroupRepresentation)

	IdentityDomainsIdentityDomainsDynamicResourceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"dynamic_resource_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_dynamic_resource_group.test_dynamic_resource_group.id}`},
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":            acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsDynamicResourceGroupDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"dynamic_resource_group_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"dynamic_resource_group_filter": acctest.Representation{RepType: acctest.Optional, Create: `displayName eq \"displayName2\"`},
		"attribute_sets":                acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                   acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsDynamicResourceGroupRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"matching_rule":  acctest.Representation{RepType: acctest.Required, Create: `Any {Any {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}`, Update: `Any {All {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}`},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:DynamicResourceGroup`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsDynamicResourceGroupTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextension_oci_tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsDynamicResourceGroup},
	}

	ignoreChangeForIdentityDomainsDynamicResourceGroup = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags`,
			`schemas`,
		}},
	}
	IdentityDomainsDynamicResourceGroupTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsRepresentation = map[string]interface{}{
		"defined_tags":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation},
		"freeform_tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation},
	}
	IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsDefinedTagsRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag.tag1.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.tag-namespace1.name}`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsDynamicResourceGroupUrnietfparamsscimschemasoracleidcsextensionOCITagsFreeformTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `freeformKey`, Update: `freeformKey2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `freeformValue`, Update: `freeformValue2`},
	}

	IdentityDomainsDynamicResourceGroupResourceDependencies = DefinedTagsDependencies + TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsDynamicResourceGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsDynamicResourceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_dynamic_resource_group.test_dynamic_resource_group"
	datasourceName := "data.oci_identity_domains_dynamic_resource_groups.test_dynamic_resource_groups"
	singularDatasourceName := "data.oci_identity_domains_dynamic_resource_group.test_dynamic_resource_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsDynamicResourceGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Create, IdentityDomainsDynamicResourceGroupRepresentation), "identitydomains", "dynamicResourceGroup", t)

	print(config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Create, IdentityDomainsDynamicResourceGroupRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsDynamicResourceGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Required, acctest.Create, IdentityDomainsDynamicResourceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", "Any {Any {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Create, IdentityDomainsDynamicResourceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", "Any {Any {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "dynamicResourceGroups", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Update, IdentityDomainsDynamicResourceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", "Any {All {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_dynamic_resource_groups", "test_dynamic_resource_groups", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsDynamicResourceGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Optional, acctest.Update, IdentityDomainsDynamicResourceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "dynamic_resource_group_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "dynamic_resource_groups.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "dynamic_resource_groups.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_dynamic_resource_group", "test_dynamic_resource_group", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsDynamicResourceGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dynamic_resource_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.key", "freeformKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextension_oci_tags.0.freeform_tags.0.value", "freeformValue2"),
			),
		},

		// reset to required only resource for import
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsDynamicResourceGroupRequiredOnlyResource,
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsDynamicResourceGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_dynamic_resource_group", "dynamicResourceGroups"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"matching_rule",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsDynamicResourceGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_dynamic_resource_group" {
			noResourceFound = false
			request := oci_identity_domains.GetDynamicResourceGroupRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.DynamicResourceGroupId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetDynamicResourceGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsDynamicResourceGroup") {
		resource.AddTestSweepers("IdentityDomainsDynamicResourceGroup", &resource.Sweeper{
			Name:         "IdentityDomainsDynamicResourceGroup",
			Dependencies: acctest.DependencyGraph["dynamicResourceGroup"],
			F:            sweepIdentityDomainsDynamicResourceGroupResource,
		})
	}
}

func sweepIdentityDomainsDynamicResourceGroupResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	dynamicResourceGroupIds, err := getIdentityDomainsDynamicResourceGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, dynamicResourceGroupId := range dynamicResourceGroupIds {
		if ok := acctest.SweeperDefaultResourceId[dynamicResourceGroupId]; !ok {
			deleteDynamicResourceGroupRequest := oci_identity_domains.DeleteDynamicResourceGroupRequest{}

			deleteDynamicResourceGroupRequest.DynamicResourceGroupId = &dynamicResourceGroupId

			deleteDynamicResourceGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteDynamicResourceGroup(context.Background(), deleteDynamicResourceGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting DynamicResourceGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", dynamicResourceGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsDynamicResourceGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DynamicResourceGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listDynamicResourceGroupsRequest := oci_identity_domains.ListDynamicResourceGroupsRequest{}
	listDynamicResourceGroupsResponse, err := identityDomainsClient.ListDynamicResourceGroups(context.Background(), listDynamicResourceGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DynamicResourceGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dynamicResourceGroup := range listDynamicResourceGroupsResponse.Resources {
		id := *dynamicResourceGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DynamicResourceGroupId", id)
	}
	return resourceIds, nil
}
