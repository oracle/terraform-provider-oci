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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	IdentityDomainsAppRoleRequiredOnlyResource = IdentityDomainsAppRoleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Required, acctest.Create, IdentityDomainsAppRoleRepresentation)

	IdentityDomainsAppRoleResourceConfig = IdentityDomainsAppRoleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Optional, acctest.Update, IdentityDomainsAppRoleRepresentation)

	IdentityDomainsIdentityDomainsAppRoleSingularDataSourceRepresentation = map[string]interface{}{
		"app_role_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app_role.test_app_role.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAppRoleDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"app_role_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"app_role_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":     acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsAppRoleRepresentation = map[string]interface{}{
		"app":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAppRoleAppRepresentation},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":              acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:AppRole`}},
		"admin_role":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"available_to_clients": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"available_to_groups":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"available_to_users":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"legacy_group_name":    acctest.Representation{RepType: acctest.Optional, Create: `legacyGroupName`},
		"public":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"tags":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAppRoleTagsRepresentation},
	}
	IdentityDomainsAppRoleAppRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app.test_app.id}`},
	}
	IdentityDomainsAppRoleTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsAppRoleResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Required, acctest.Create, IdentityDomainsAppRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAppRoleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAppRoleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_app_role.test_app_role"
	datasourceName := "data.oci_identity_domains_app_roles.test_app_roles"
	singularDatasourceName := "data.oci_identity_domains_app_role.test_app_role"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsAppRoleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Optional, acctest.Create, IdentityDomainsAppRoleRepresentation), "identitydomains", "appRole", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsAppRoleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppRoleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Required, acctest.Create, IdentityDomainsAppRoleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "app.0.value"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppRoleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAppRoleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Optional, acctest.Create, IdentityDomainsAppRoleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_role", "false"),
				resource.TestCheckResourceAttr(resourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "app.0.value"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "available_to_clients", "false"),
				resource.TestCheckResourceAttr(resourceName, "available_to_groups", "false"),
				resource.TestCheckResourceAttr(resourceName, "available_to_users", "false"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "legacy_group_name"),
				resource.TestCheckResourceAttr(resourceName, "public", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "appRoles", resId)
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_app_roles", "test_app_roles", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsAppRoleDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAppRoleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Optional, acctest.Update, IdentityDomainsAppRoleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "app_role_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "app_roles.#", regexp.MustCompile("[1-9]+")),
				resource.TestMatchResourceAttr(datasourceName, "app_roles.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAppRoleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAppRoleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "app_role_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "admin_role", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "app.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "available_to_clients", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "available_to_groups", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "available_to_users", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "public", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsAppRoleRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_app_role", "appRoles"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsAppRoleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_app_role" {
			noResourceFound = false
			request := oci_identity_domains.GetAppRoleRequest{}

			tmp := rs.Primary.ID
			request.AppRoleId = &tmp

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetAppRole(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsAppRole") {
		resource.AddTestSweepers("IdentityDomainsAppRole", &resource.Sweeper{
			Name:         "IdentityDomainsAppRole",
			Dependencies: acctest.DependencyGraph["appRole"],
			F:            sweepIdentityDomainsAppRoleResource,
		})
	}
}

func sweepIdentityDomainsAppRoleResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	appRoleIds, err := getIdentityDomainsAppRoleIds(compartment)
	if err != nil {
		return err
	}
	for _, appRoleId := range appRoleIds {
		if ok := acctest.SweeperDefaultResourceId[appRoleId]; !ok {
			deleteAppRoleRequest := oci_identity_domains.DeleteAppRoleRequest{}

			deleteAppRoleRequest.AppRoleId = &appRoleId

			deleteAppRoleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteAppRole(context.Background(), deleteAppRoleRequest)
			if error != nil {
				fmt.Printf("Error deleting AppRole %s %s, It is possible that the resource is already deleted. Please verify manually \n", appRoleId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsAppRoleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AppRoleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listAppRolesRequest := oci_identity_domains.ListAppRolesRequest{}
	listAppRolesResponse, err := identityDomainsClient.ListAppRoles(context.Background(), listAppRolesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AppRole list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, appRole := range listAppRolesResponse.Resources {
		id := *appRole.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AppRoleId", id)
	}
	return resourceIds, nil
}
