// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var TenantmanagercontrolplaneChildTenancyRepresentation = map[string]interface{}{
	"admin_email":       acctest.Representation{RepType: acctest.Required, Create: `${var.child_tenancy_admin_email}`},
	"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	"home_region":       acctest.Representation{RepType: acctest.Required, Create: `${var.child_tenancy_home_region}`},
	"tenancy_name":      acctest.Representation{RepType: acctest.Required, Create: `${var.child_tenancy_name}`},
	"governance_status": acctest.Representation{RepType: acctest.Optional, Create: `OPTED_IN`},
}

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneChildTenancyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneChildTenancyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	inputs := map[string]string{}
	for _, name := range []string{"compartment_ocid", "child_tenancy_admin_email", "child_tenancy_home_region", "child_tenancy_name", "child_tenancy_governed_name"} {
		inputs[name] = utils.GetEnvSettingWithBlankDefault(name)
		if os.Getenv("TF_ACC") != "" && inputs[name] == "" {
			t.Fatalf("TF_VAR_%s must be set; this test creates two persistent child tenancies", name)
		}
		variableName := name
		if name == "compartment_ocid" {
			variableName = "compartment_id"
		}
		config += fmt.Sprintf("variable %q { default = %q }\n", variableName, inputs[name])
	}
	if os.Getenv("TF_ACC") != "" {
		if inputs["child_tenancy_name"] == inputs["child_tenancy_governed_name"] {
			t.Fatal("child_tenancy_name and child_tenancy_governed_name must be distinct unused names")
		}
		// Delete is a no-op: the service retains these tenancies. Preserve each
		// applied configuration and state before the harness removes local state.
		t.Setenv("TF_ACC_PERSIST_WORKING_DIR", "1")
	}

	resourceName := "oci_tenantmanagercontrolplane_child_tenancy.test_child_tenancy"
	governedResourceName := "oci_tenantmanagercontrolplane_child_tenancy.test_child_tenancy_governed"
	requiredConfig := config + acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_child_tenancy", "test_child_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneChildTenancyRepresentation)
	governedRepresentation := acctest.GetUpdatedRepresentationCopy("tenancy_name", acctest.Representation{RepType: acctest.Required, Create: `${var.child_tenancy_governed_name}`}, TenantmanagercontrolplaneChildTenancyRepresentation)
	// Keep the first resource in the second step; a remove/recreate sequence
	// cannot reuse a tenancy name because this API has no delete operation.
	bothConfig := requiredConfig + acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_child_tenancy", "test_child_tenancy_governed", acctest.Optional, acctest.Create, governedRepresentation)
	acctest.SaveConfigContent(bothConfig, "tenantmanagercontrolplane", "childTenancy", t)

	checkInputs := func(address, name string) resource.TestCheckFunc {
		return resource.ComposeAggregateTestCheckFunc(
			resource.TestMatchResourceAttr(address, "id", regexp.MustCompile(`^ocid1\.tenancy\.`)),
			resource.TestCheckResourceAttr(address, "admin_email", inputs["child_tenancy_admin_email"]),
			resource.TestCheckResourceAttr(address, "compartment_id", inputs["compartment_ocid"]),
			resource.TestCheckResourceAttr(address, "home_region", inputs["child_tenancy_home_region"]),
			resource.TestCheckResourceAttr(address, "tenancy_name", name),
		)
	}
	var firstID string
	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: requiredConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				checkInputs(resourceName, inputs["child_tenancy_name"]),
				func(s *terraform.State) (err error) {
					firstID, err = acctest.FromInstanceState(s, resourceName, "id")
					t.Logf("Child created with governance omitted: name=%s id=%s", inputs["child_tenancy_name"], firstID)
					return err
				},
			),
		},
		{
			Config: bothConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				checkInputs(resourceName, inputs["child_tenancy_name"]),
				checkInputs(governedResourceName, inputs["child_tenancy_governed_name"]),
				resource.TestCheckResourceAttr(governedResourceName, "governance_status", "OPTED_IN"),
				func(s *terraform.State) error {
					if err := resource.TestCheckResourceAttr(resourceName, "id", firstID)(s); err != nil {
						return err
					}
					governedID, err := acctest.FromInstanceState(s, governedResourceName, "id")
					if err != nil {
						return err
					}
					if governedID == firstID {
						return fmt.Errorf("distinct child tenancies unexpectedly share ID %s", firstID)
					}
					t.Logf("Child created with governance OPTED_IN: name=%s id=%s", inputs["child_tenancy_governed_name"], governedID)
					return nil
				},
			),
		},
	})
}
