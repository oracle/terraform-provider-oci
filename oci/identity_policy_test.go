// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

var (
	PolicyRequiredOnlyResource = PolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Required, Create, policyRepresentation)

	policyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":         RepresentationGroup{Required, policyDataSourceFilterRepresentation}}
	policyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_policy.test_policy.id}`}},
	}

	policyRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `Policy for users who need to launch instances, attach volumes, manage images`, update: `description2`},
		"name":           Representation{repType: Required, create: `LaunchInstances`},
		"statements":     Representation{repType: Required, create: []string{`Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}`}},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"version_date":   Representation{repType: Optional, create: ``, update: `2018-01-01`},
	}

	PolicyResourceDependencies = DefinedTagsDependencies + `
resource "oci_identity_compartment" "t" {
	name = "Network"
	description = "For network components"
	compartment_id = "${var.tenancy_ocid}"
}

resource "oci_identity_group" "t" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "group for policy test"
	name = "GroupName"
}
`
)

func TestIdentityPolicyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_policy.test_policy"
	datasourceName := "data.oci_identity_policies.test_policies"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityPolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Required, Create, policyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
					resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PolicyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Create, policyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckNoResourceAttr(resourceName, "version_date"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Update, policyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "version_date", "2018-01-01"),

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
					generateDataSourceFromRepresentationMap("oci_identity_policies", "test_policies", Optional, Update, policyDataSourceRepresentation) +
					compartmentIdVariableStr + PolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Update, policyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.name", "LaunchInstances"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.statements.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "policies.0.version_date", "2018-01-01"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// ETag, lastUpdateETag, and policyHash are non-API fields that
					// get computed during resource Create/Update but omitted from Get calls.
					// These are internally used for diff suppression and not needed for imports.
					// Omit them in the import verification.
					"ETag",
					"lastUpdateETag",
					"policyHash",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckIdentityPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_policy" {
			noResourceFound = false
			request := oci_identity.GetPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			response, err := client.GetPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.PolicyLifecycleStateDeleted): true,
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
	resource.AddTestSweepers("IdentityPolicy", &resource.Sweeper{
		Name:         "IdentityPolicy",
		Dependencies: DependencyGraph["policy"],
		F:            sweepIdentityPolicyResource,
	})
}

func sweepIdentityPolicyResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	policyIds, err := getPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, policyId := range policyIds {
		if ok := SweeperDefaultResourceId[policyId]; !ok {
			deletePolicyRequest := oci_identity.DeletePolicyRequest{}

			deletePolicyRequest.PolicyId = &policyId

			deletePolicyRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeletePolicy(context.Background(), deletePolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting Policy %s %s, It is possible that the resource is already deleted. Please verify manually \n", policyId, error)
				continue
			}
			waitTillCondition(testAccProvider, &policyId, policySweepWaitCondition, time.Duration(3*time.Minute),
				policySweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getPolicyIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listPoliciesRequest := oci_identity.ListPoliciesRequest{}
	listPoliciesRequest.CompartmentId = &compartmentId
	listPoliciesResponse, err := identityClient.ListPolicies(context.Background(), listPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Policy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, policy := range listPoliciesResponse.Items {
		id := *policy.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "PolicyId", id)
	}
	return resourceIds, nil
}

func policySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if policyResponse, ok := response.Response.(oci_identity.GetPolicyResponse); ok {
		return policyResponse.LifecycleState == oci_identity.PolicyLifecycleStateDeleted
	}
	return false
}

func policySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetPolicy(context.Background(), oci_identity.GetPolicyRequest{
		PolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
