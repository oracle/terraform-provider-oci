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
	IdentityDomainsSocialIdentityProviderRequiredOnlyResource = IdentityDomainsSocialIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Required, acctest.Create, IdentityDomainsSocialIdentityProviderRepresentation)

	IdentityDomainsSocialIdentityProviderResourceConfig = IdentityDomainsSocialIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsSocialIdentityProviderRepresentation)

	IdentityDomainsSocialIdentityProviderSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"social_identity_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_social_identity_provider.test_social_identity_provider.id}`},
	}

	IdentityDomainsSocialIdentityProviderDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"social_identity_provider_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"social_identity_provider_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                     acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsSocialIdentityProviderRepresentation = map[string]interface{}{
		"account_linking_enabled":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"consumer_key":                       acctest.Representation{RepType: acctest.Required, Create: `consumerKey`, Update: `consumerKey2`},
		"consumer_secret":                    acctest.Representation{RepType: acctest.Required, Create: `consumerSecret`, Update: `consumerSecret2`},
		"enabled":                            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"idcs_endpoint":                      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":                               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"registration_enabled":               acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"schemas":                            acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:SocialIdentityProvider`}},
		"service_provider_name":              acctest.Representation{RepType: acctest.Required, Create: `Google`},
		"show_on_login":                      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"access_token_url":                   acctest.Representation{RepType: acctest.Optional, Create: `https://something1.com/token`, Update: `https://something2.com/token`},
		"admin_scope":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`adminScope`}, Update: []string{`adminScope2`}},
		"authz_url":                          acctest.Representation{RepType: acctest.Optional, Create: `https://something1.com`, Update: `https://something2.com`},
		"auto_redirect_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"client_credential_in_payload":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"clock_skew_in_seconds":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"description":                        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"jit_prov_assigned_groups":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSocialIdentityProviderJitProvAssignedGroupsRepresentation},
		"jit_prov_group_static_list_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"profile_url":                        acctest.Representation{RepType: acctest.Optional, Create: `https://something.com/profileUrl1.png`, Update: `https://something.com/profileUrl2.png`},
		"redirect_url":                       acctest.Representation{RepType: acctest.Optional, Create: `https://redirectUrl1.com`, Update: `https://redirectUrl2.com`},
		"relay_idp_param_mappings":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSocialIdentityProviderRelayIdpParamMappingsRepresentation},
		"scope":                              acctest.Representation{RepType: acctest.Optional, Create: []string{`scope`}, Update: []string{`scope2`}},
		"social_jit_provisioning_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"status":                             acctest.Representation{RepType: acctest.Optional, Create: `created`, Update: `deleted`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsSocialIdentityProvider},
	}
	ignoreChangeForIdentityDomainsSocialIdentityProvider = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`tags`,
			`schemas`,
		}},
	}
	IdentityDomainsSocialIdentityProviderJitProvAssignedGroupsRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`, Update: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsSocialIdentityProviderRelayIdpParamMappingsRepresentation = map[string]interface{}{
		"relay_param_key":   acctest.Representation{RepType: acctest.Required, Create: `relayParamKey`, Update: `relayParamKey2`},
		"relay_param_value": acctest.Representation{RepType: acctest.Optional, Create: `relayParamValue`, Update: `relayParamValue2`},
	}
	IdentityDomainsSocialIdentityProviderTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsSocialIdentityProviderResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsSocialIdentityProviderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsSocialIdentityProviderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_social_identity_provider.test_social_identity_provider"
	datasourceName := "data.oci_identity_domains_social_identity_providers.test_social_identity_providers"
	singularDatasourceName := "data.oci_identity_domains_social_identity_provider.test_social_identity_provider"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsSocialIdentityProviderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Optional, acctest.Create, IdentityDomainsSocialIdentityProviderRepresentation), "identitydomains", "socialIdentityProvider", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsSocialIdentityProviderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IdentityDomainsSocialIdentityProviderRepresentation, map[string]interface{}{
						"enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
					},
					),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "account_linking_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "consumer_key", "consumerKey"),
				resource.TestCheckResourceAttr(resourceName, "consumer_secret", "consumerSecret"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "registration_enabled", "false"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_name"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Optional, acctest.Create, IdentityDomainsSocialIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_url", "https://something1.com/token"),
				resource.TestCheckResourceAttr(resourceName, "account_linking_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "admin_scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authz_url", "https://something1.com"),
				resource.TestCheckResourceAttr(resourceName, "auto_redirect_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_credential_in_payload", "false"),
				resource.TestCheckResourceAttr(resourceName, "clock_skew_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "consumer_key", "consumerKey"),
				resource.TestCheckResourceAttr(resourceName, "consumer_secret", "consumerSecret"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "jit_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_prov_assigned_groups.0.ref"),
				resource.TestCheckResourceAttr(resourceName, "jit_prov_group_static_list_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "profile_url", "https://something.com/profileUrl1.png"),
				resource.TestCheckResourceAttr(resourceName, "redirect_url", "https://redirectUrl1.com"),
				resource.TestCheckResourceAttr(resourceName, "registration_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.0.relay_param_key", "relayParamKey"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.0.relay_param_value", "relayParamValue"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "scope.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_name"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login", "false"),
				resource.TestCheckResourceAttr(resourceName, "social_jit_provisioning_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "status", "created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "socialIdentityProviders", resId) // replace "groups" with correct resource name
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
			Config: config + compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsSocialIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_url", "https://something2.com/token"),
				resource.TestCheckResourceAttr(resourceName, "account_linking_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "admin_scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authz_url", "https://something2.com"),
				resource.TestCheckResourceAttr(resourceName, "auto_redirect_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_credential_in_payload", "true"),
				resource.TestCheckResourceAttr(resourceName, "clock_skew_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "consumer_key", "consumerKey2"),
				resource.TestCheckResourceAttr(resourceName, "consumer_secret", "consumerSecret2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "jit_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_prov_assigned_groups.0.ref"),
				resource.TestCheckResourceAttr(resourceName, "jit_prov_group_static_list_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "profile_url", "https://something.com/profileUrl2.png"),
				resource.TestCheckResourceAttr(resourceName, "redirect_url", "https://redirectUrl2.com"),
				resource.TestCheckResourceAttr(resourceName, "registration_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.0.relay_param_key", "relayParamKey2"),
				resource.TestCheckResourceAttr(resourceName, "relay_idp_param_mappings.0.relay_param_value", "relayParamValue2"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "scope.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_name"),
				resource.TestCheckResourceAttr(resourceName, "show_on_login", "true"),
				resource.TestCheckResourceAttr(resourceName, "social_jit_provisioning_enabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_social_identity_providers", "test_social_identity_providers", acctest.Optional, acctest.Update, IdentityDomainsSocialIdentityProviderDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsSocialIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "social_identity_provider_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "social_identity_providers.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "social_identity_providers.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_social_identity_provider", "test_social_identity_provider", acctest.Required, acctest.Create, IdentityDomainsSocialIdentityProviderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSocialIdentityProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "social_identity_provider_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_token_url", "https://something2.com/token"),
				resource.TestCheckResourceAttr(singularDatasourceName, "account_linking_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_scope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authz_url", "https://something2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auto_redirect_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_credential_in_payload", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "clock_skew_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumer_key", "consumerKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumer_secret", "consumerSecret2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jit_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jit_prov_assigned_groups.0.ref"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_prov_group_static_list_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile_url", "https://something.com/profileUrl2.png"),
				resource.TestCheckResourceAttr(singularDatasourceName, "redirect_url", "https://redirectUrl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "registration_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "relay_idp_param_mappings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "relay_idp_param_mappings.0.relay_param_key", "relayParamKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "relay_idp_param_mappings.0.relay_param_value", "relayParamValue2"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "show_on_login", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "social_jit_provisioning_enabled", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsSocialIdentityProviderRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_social_identity_provider", "socialIdentityProviders"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsSocialIdentityProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_social_identity_provider" {
			noResourceFound = false
			request := oci_identity_domains.GetSocialIdentityProviderRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.SocialIdentityProviderId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetSocialIdentityProvider(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsSocialIdentityProvider") {
		resource.AddTestSweepers("IdentityDomainsSocialIdentityProvider", &resource.Sweeper{
			Name:         "IdentityDomainsSocialIdentityProvider",
			Dependencies: acctest.DependencyGraph["socialIdentityProvider"],
			F:            sweepIdentityDomainsSocialIdentityProviderResource,
		})
	}
}

func sweepIdentityDomainsSocialIdentityProviderResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	socialIdentityProviderIds, err := getIdentityDomainsSocialIdentityProviderIds(compartment)
	if err != nil {
		return err
	}
	for _, socialIdentityProviderId := range socialIdentityProviderIds {
		if ok := acctest.SweeperDefaultResourceId[socialIdentityProviderId]; !ok {
			deleteSocialIdentityProviderRequest := oci_identity_domains.DeleteSocialIdentityProviderRequest{}

			deleteSocialIdentityProviderRequest.SocialIdentityProviderId = &socialIdentityProviderId

			deleteSocialIdentityProviderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteSocialIdentityProvider(context.Background(), deleteSocialIdentityProviderRequest)
			if error != nil {
				fmt.Printf("Error deleting SocialIdentityProvider %s %s, It is possible that the resource is already deleted. Please verify manually \n", socialIdentityProviderId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsSocialIdentityProviderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SocialIdentityProviderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listSocialIdentityProvidersRequest := oci_identity_domains.ListSocialIdentityProvidersRequest{}
	listSocialIdentityProvidersResponse, err := identityDomainsClient.ListSocialIdentityProviders(context.Background(), listSocialIdentityProvidersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SocialIdentityProvider list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, socialIdentityProvider := range listSocialIdentityProvidersResponse.Resources {
		id := *socialIdentityProvider.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SocialIdentityProviderId", id)
	}
	return resourceIds, nil
}
