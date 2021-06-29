// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_identity "github.com/oracle/oci-go-sdk/v43/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PolicyRequiredOnlyResource = PolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Required, Create, policyRepresentation)

	policyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"name":           Representation{repType: Optional, create: `LaunchInstances`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, policyDataSourceFilterRepresentation}}
	policyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_policy.test_policy.id}`}},
	}

	policyRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `Policy for users who need to launch instances, attach volumes, manage images`, update: `description2`},
		"name":           Representation{repType: Required, create: `LaunchInstances`},
		"statements":     Representation{repType: Required, create: []string{`Allow group Administrators to read instances in tenancy`}},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"version_date":   Representation{repType: Optional, create: ``, update: `2018-01-01`},
	}

	PolicyResourceDependencies = DefinedTagsDependencies
)

func TestIdentityPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_policy.test_policy"
	datasourceName := "data.oci_identity_policies.test_policies"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PolicyResourceDependencies+
		generateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Create, policyRepresentation), "identity", "policy", t)

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
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
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
					resource.TestCheckResourceAttr(datasourceName, "name", "LaunchInstances"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

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
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_policy" {
			noResourceFound = false
			request := oci_identity.GetPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

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
