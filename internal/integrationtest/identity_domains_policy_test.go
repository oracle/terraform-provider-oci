// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	IdentityDomainsPolicyRequiredOnlyResource = IdentityDomainsPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Required, acctest.Create, IdentityDomainsPolicyRepresentation)

	IdentityDomainsPolicyResourceConfig = IdentityDomainsPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Optional, acctest.Update, IdentityDomainsPolicyRepresentation)

	IdentityDomainsPolicySingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_policy.test_policy.id}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsPolicyDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"policy_count":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"policy_filter":  acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"sort_by":        acctest.Representation{RepType: acctest.Optional, Create: `id`},
	}

	IdentityDomainsPolicyRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"policy_type":    acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsPolicyPolicyTypeRepresentation},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Policy`}},
		"active":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_id":    acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"policy_groovy":  acctest.Representation{RepType: acctest.Optional, Create: `policyGroovy`, Update: `policyGroovy2`},
		"rules":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsPolicyRulesRepresentation},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsPolicyTagsRepresentation},
	}
	IdentityDomainsPolicyPolicyTypeRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `IdentityProvider`},
	}
	IdentityDomainsPolicyRulesRepresentation = map[string]interface{}{
		"sequence": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"value":    acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_rule.test_rule.id}`},
	}
	IdentityDomainsPolicyTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsTestRulePolicyTypeRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `IdentityProvider`},
	}

	IdentityDomainsTestRuleReturnRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `LocalIDPs`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `[\"UserNamePassword\"]`},
	}

	IdentityDomainsPolicyResourceDependencies = TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(IdentityDomainsRuleRepresentation, map[string]interface{}{
				"policy_type": acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsTestRulePolicyTypeRepresentation},
				"return":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsTestRuleReturnRepresentation},
			}),
		)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_policy.test_policy"
	datasourceName := "data.oci_identity_domains_policies.test_policies"
	singularDatasourceName := "data.oci_identity_domains_policy.test_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Optional, acctest.Create, IdentityDomainsPolicyRepresentation), "identitydomains", "policy", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Required, acctest.Create, IdentityDomainsPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "IdentityProvider"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Optional, acctest.Create, IdentityDomainsPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_groovy", "policyGroovy"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "IdentityProvider"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.sequence", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "rules.0.value"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "policies", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Optional, acctest.Update, IdentityDomainsPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "policy_groovy", "policyGroovy2"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "IdentityProvider"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.sequence", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "rules.0.value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_policies", "test_policies", acctest.Optional, acctest.Update, IdentityDomainsPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Optional, acctest.Update, IdentityDomainsPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "policy_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "policies.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_policy", "test_policy", acctest.Required, acctest.Create, IdentityDomainsPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_groovy", "policyGroovy2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_type.0.value", "IdentityProvider"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsPolicyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_policy", "policies"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"rules",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_policy" {
			noResourceFound = false
			request := oci_identity_domains.GetPolicyRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetPolicy(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsPolicy") {
		resource.AddTestSweepers("IdentityDomainsPolicy", &resource.Sweeper{
			Name:         "IdentityDomainsPolicy",
			Dependencies: acctest.DependencyGraph["policy"],
			F:            sweepIdentityDomainsPolicyResource,
		})
	}
}

func sweepIdentityDomainsPolicyResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	policyIds, err := getIdentityDomainsPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, policyId := range policyIds {
		if ok := acctest.SweeperDefaultResourceId[policyId]; !ok {
			deletePolicyRequest := oci_identity_domains.DeletePolicyRequest{}

			deletePolicyRequest.PolicyId = &policyId

			deletePolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeletePolicy(context.Background(), deletePolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting Policy %s %s, It is possible that the resource is already deleted. Please verify manually \n", policyId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listPoliciesRequest := oci_identity_domains.ListPoliciesRequest{}
	listPoliciesResponse, err := identityDomainsClient.ListPolicies(context.Background(), listPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Policy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, policy := range listPoliciesResponse.Resources {
		id := *policy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PolicyId", id)
	}
	return resourceIds, nil
}
