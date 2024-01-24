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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	IdentityDomainsGrantRequiredOnlyResource = IdentityDomainsGrantResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Required, acctest.Create, IdentityDomainsGrantRepresentation)

	IdentityDomainsGrantResourceConfig = IdentityDomainsGrantResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Optional, acctest.Update, IdentityDomainsGrantRepresentation)

	IdentityDomainsIdentityDomainsGrantSingularDataSourceRepresentation = map[string]interface{}{
		"grant_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_grant.test_grant.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsGrantDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"grant_count":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"grant_filter":   acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsGrantRepresentation = map[string]interface{}{
		"grant_mechanism": acctest.Representation{RepType: acctest.Required, Create: `IMPORT_APPROLE_MEMBERS`},
		"grantee":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsGrantGranteeRepresentation},
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":         acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Grant`}},
		"app":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsGrantAppRepresentation},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"entitlement":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGrantEntitlementRepresentation},
		"tags":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsGrantTagsRepresentation},
	}
	IdentityDomainsGrantGranteeRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}
	IdentityDomainsGrantAppRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_apps.test_grant_apps.apps.0.id}`},
	}
	IdentityDomainsGrantEntitlementRepresentation = map[string]interface{}{
		"attribute_name":  acctest.Representation{RepType: acctest.Required, Create: `appRoles`},
		"attribute_value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app_role.test_app_role.id}`},
	}
	IdentityDomainsGrantTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsGrantRepresentation2 = acctest.RepresentationCopyWithNewProperties(IdentityDomainsGrantRepresentation, map[string]interface{}{
		"grantee": acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsGrantGranteeRepresentation2},
	})
	IdentityDomainsGrantGranteeRepresentation2 = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user2.id}`},
	}

	IdentityDomainsGrantResourceDependencies = TestDomainDependencies + GrantTestAppDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user2", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(IdentityDomainsUserRepresentation, map[string]interface{}{
				"user_name": acctest.Representation{RepType: acctest.Required, Create: `userName2`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app_role", "test_app_role", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(IdentityDomainsAppRoleRepresentation, map[string]interface{}{
				"app": acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAppRoleAppRepresentationForGrants},
			}))

	IdentityDomainsAppRoleAppRepresentationForGrants = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_apps.test_grant_apps.apps.0.id}`},
	}
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsGrantResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsGrantResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_grant.test_grant"
	datasourceName := "data.oci_identity_domains_grants.test_grants"
	singularDatasourceName := "data.oci_identity_domains_grant.test_grant"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsGrantResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Optional, acctest.Create, IdentityDomainsGrantRepresentation), "identitydomains", "grant", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsGrantDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGrantResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Required, acctest.Create, IdentityDomainsGrantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "grant_mechanism", "IMPORT_APPROLE_MEMBERS"),
				resource.TestCheckResourceAttr(resourceName, "grantee.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "grantee.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGrantResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsGrantResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Optional, acctest.Create, IdentityDomainsGrantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "app.0.value"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_key"),
				resource.TestCheckResourceAttr(resourceName, "entitlement.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlement.0.attribute_name", "appRoles"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlement.0.attribute_value"),
				resource.TestCheckResourceAttr(resourceName, "grant_mechanism", "IMPORT_APPROLE_MEMBERS"),
				resource.TestCheckResourceAttr(resourceName, "grantee.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "grantee.0.type", "User"),
				resource.TestCheckResourceAttr(resourceName, "grantor.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "grantor.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "is_fulfilled"),
				resource.TestCheckResourceAttrSet(resourceName, "ocid"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "grants", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_grants", "test_grants", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsGrantDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant1", acctest.Optional, acctest.Update, IdentityDomainsGrantRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_grant", "test_grant2", acctest.Optional, acctest.Update, IdentityDomainsGrantRepresentation2) +
				compartmentIdVariableStr + IdentityDomainsGrantResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "grant_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "grants.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "grants.0.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.grantee.0.value"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.app.0.value"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.entitlement.0.attribute_value"),
				resource.TestCheckResourceAttr(datasourceName, "grants.1.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.1.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.1.grantee.0.value"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.1.app.0.value"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.1.entitlement.0.attribute_value"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_grant", "test_grant", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsGrantSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsGrantResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "app.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlement.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlement.0.attribute_name", "appRoles"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entitlement.0.attribute_value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grant_mechanism", "IMPORT_APPROLE_MEMBERS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grantee.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grantee.0.type", "User"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantee.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grantor.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grantor.0.type", "User"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantor.0.value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_fulfilled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsGrantRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_grant", "grants"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"composite_key",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsGrantDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_grant" {
			noResourceFound = false
			request := oci_identity_domains.GetGrantRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.GrantId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetGrant(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsGrant") {
		resource.AddTestSweepers("IdentityDomainsGrant", &resource.Sweeper{
			Name:         "IdentityDomainsGrant",
			Dependencies: acctest.DependencyGraph["grant"],
			F:            sweepIdentityDomainsGrantResource,
		})
	}
}

func sweepIdentityDomainsGrantResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	grantIds, err := getIdentityDomainsGrantIds(compartment)
	if err != nil {
		return err
	}
	for _, grantId := range grantIds {
		if ok := acctest.SweeperDefaultResourceId[grantId]; !ok {
			deleteGrantRequest := oci_identity_domains.DeleteGrantRequest{}

			deleteGrantRequest.GrantId = &grantId

			deleteGrantRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteGrant(context.Background(), deleteGrantRequest)
			if error != nil {
				fmt.Printf("Error deleting Grant %s %s, It is possible that the resource is already deleted. Please verify manually \n", grantId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsGrantIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "GrantId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listGrantsRequest := oci_identity_domains.ListGrantsRequest{}
	listGrantsResponse, err := identityDomainsClient.ListGrants(context.Background(), listGrantsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Grant list for compartment id : %s , %s \n", compartmentId, err)
	}

	for _, grant := range listGrantsResponse.Resources {
		id := *grant.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "GrantId", id)
	}
	return resourceIds, nil
}
