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
	IdentityUserGroupMembershipRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", acctest.Required, acctest.Create, IdentityUserGroupMembershipRepresentation)

	IdentityIdentityUserGroupMembershipDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"group_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.id}`},
		"user_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityUserGroupMembershipDataSourceFilterRepresentation}}
	IdentityUserGroupMembershipDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_user_group_membership.test_user_group_membership.id}`}},
	}

	IdentityUserGroupMembershipRepresentation = map[string]interface{}{
		"group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.id}`},
		"user_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	IdentityUserGroupMembershipResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, IdentityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUserGroupMembershipResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"
	datasourceName := "data.oci_identity_user_group_memberships.test_user_group_memberships"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityUserGroupMembershipResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", acctest.Required, acctest.Create, IdentityUserGroupMembershipRepresentation), "identity", "userGroupMembership", t)

	acctest.ResourceTest(t, testAccCheckIdentityUserGroupMembershipDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityUserGroupMembershipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", acctest.Required, acctest.Create, IdentityUserGroupMembershipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_user_group_memberships", "test_user_group_memberships", acctest.Optional, acctest.Update, IdentityIdentityUserGroupMembershipDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityUserGroupMembershipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", acctest.Optional, acctest.Update, IdentityUserGroupMembershipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "memberships.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.user_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + IdentityUserGroupMembershipRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityUserGroupMembershipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user_group_membership" {
			noResourceFound = false
			request := oci_identity.GetUserGroupMembershipRequest{}

			tmp := rs.Primary.ID
			request.UserGroupMembershipId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetUserGroupMembership(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.UserGroupMembershipLifecycleStateDeleted): true,
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
