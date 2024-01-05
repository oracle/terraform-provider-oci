// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityPolicyRequiredOnlyResource = IdentityPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation)

	IdentityIdentityPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `LaunchInstances`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityPolicyDataSourceFilterRepresentation}}
	IdentityPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_policy.test_policy.id}`}},
	}

	IdentityPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Policy for users who need to launch instances, attach volumes, manage images`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `LaunchInstances`},
		"statements":     acctest.Representation{RepType: acctest.Required, Create: []string{`Allow Group Administrators to read instances in tenancy`}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"version_date":   acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `2018-01-01`},
	}

	IdentityPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_policy.test_policy"
	datasourceName := "data.oci_identity_policies.test_policies"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Optional, acctest.Create, IdentityPolicyRepresentation), "identity", "policy", t)

	acctest.ResourceTest(t, testAccCheckIdentityPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Optional, acctest.Create, IdentityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckNoResourceAttr(resourceName, "version_date"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Optional, acctest.Update, IdentityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "version_date", "2018-01-01"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_policies", "test_policies", acctest.Optional, acctest.Update, IdentityIdentityPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Optional, acctest.Update, IdentityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.compartment_id", tenancyId),
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
			Config:            config + IdentityPolicyRequiredOnlyResource,
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
	})
}

func testAccCheckIdentityPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_policy" {
			noResourceFound = false
			request := oci_identity.GetPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

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
