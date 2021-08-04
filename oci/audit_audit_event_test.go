// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	auditEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"end_time":       Representation{repType: Required, create: `${timestamp()}`},
		"start_time":     Representation{repType: Required, create: `${timeadd(timestamp(), "-1m")}`},
	}

	AuditEventResourceConfig = ""
)

// issue-routing-tag: audit/default
func TestAuditAuditEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditAuditEventResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_audit_events.test_audit_events"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_audit_events", "test_audit_events", Required, Create, auditEventDataSourceRepresentation) +
					compartmentIdVariableStr + AuditEventResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "end_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time"),

					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.#"),
				),
				// Non empty plan expected because the data source input relies on interpolation syntax
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
