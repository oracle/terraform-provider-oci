// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementOccCustomerGroupOccCustomerRequiredOnlyResource = CapacityManagementOccCustomerGroupOccCustomerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group_occ_customer", "test_occ_customer_group_occ_customer", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupOccCustomerRepresentation)

	occCustomerDisplayName        = "TERSI customer"
	occCustomerUpdatedDisplayName = "Updated TERSI customer"
	occCustomerDescription        = "This is a test customer created for TERSI testing"
	occCustomerUpdatedDescription = "This is the updated test customer description for TERSI testing"

	CapacityManagementOccCustomerGroupOccCustomerRepresentation = map[string]interface{}{
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: occCustomerDisplayName, Update: occCustomerUpdatedDisplayName},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_customer_group.test_occ_customer_group.id}`},
		"tenancy_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.occ_customer_id}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: occCustomerDescription, Update: occCustomerUpdatedDescription},
		"status":                acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}

	CapacityManagementOccCustomerGroupOccCustomerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupRepresentation)
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccCustomerGroupOccCustomerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccCustomerGroupOccCustomerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	occCustomerId := utils.GetEnvSettingWithBlankDefault("customer_tenancy_id")
	occCustomerIdStr := fmt.Sprintf("variable \"occ_customer_id\" { default = \"%s\" }\n", occCustomerId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_id")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	resourceName := "oci_capacity_management_occ_customer_group_occ_customer.test_occ_customer_group_occ_customer"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+occCustomerIdStr+compartmentIdVariableStr+occCustomerGroupIdVariableStr+CapacityManagementOccCustomerGroupOccCustomerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group_occ_customer", "test_occ_customer_group_occ_customer", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupOccCustomerRepresentation), "capacitymanagement", "occCustomerGroupOccCustomer", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccCustomerGroupOccCustomerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + occCustomerIdStr + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupOccCustomerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group_occ_customer", "test_occ_customer_group_occ_customer", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupOccCustomerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", occCustomerDisplayName),
				resource.TestCheckResourceAttr(resourceName, "tenancy_id", occCustomerId),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + occCustomerIdStr + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupOccCustomerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + occCustomerIdStr + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupOccCustomerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group_occ_customer", "test_occ_customer_group_occ_customer", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupOccCustomerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", occCustomerDescription),
				resource.TestCheckResourceAttr(resourceName, "display_name", occCustomerDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

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
			Config: config + occCustomerIdStr + compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupOccCustomerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group_occ_customer", "test_occ_customer_group_occ_customer", acctest.Optional, acctest.Update, CapacityManagementOccCustomerGroupOccCustomerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", occCustomerUpdatedDescription),
				resource.TestCheckResourceAttr(resourceName, "display_name", occCustomerUpdatedDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete the customer group as well
		{
			Config: config + occCustomerIdStr + compartmentIdVariableStr + occCustomerGroupIdVariableStr,
		},
	})
}

func testAccCheckCapacityManagementOccCustomerGroupOccCustomerDestroy(s *terraform.State) error {
	noResourceFound := true

	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
