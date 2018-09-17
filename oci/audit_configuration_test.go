// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ConfigurationResourceConfig = ConfigurationResourceDependencies + `
resource "oci_audit_configuration" "test_configuration" {
	compartment_id = "${var.tenancy_ocid}"
	retention_period_days = "${var.configuration_retention_period_days_value}"
}
`
	ConfigurationPropertyVariables = `
variable "configuration_retention_period_days_value" { default = "100" }

`
	ConfigurationResourceDependencies = ""
)

func TestAuditConfigurationResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	resourceName := "oci_audit_configuration.test_configuration"

	singularDatasourceName := "data.oci_audit_configuration.test_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ConfigurationPropertyVariables + ConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "retention_period_days", "100"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "configuration_retention_period_days_value" { default = "91" }

                ` + ConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "retention_period_days", "91"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "configuration_retention_period_days_value" { default = "91" }

data "oci_audit_configuration" "test_configuration" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
                ` + ConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "retention_period_days", "91"),
				),
			},
		},
	})
}
