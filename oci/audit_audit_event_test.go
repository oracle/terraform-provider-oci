// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	auditEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"end_time":       Representation{repType: Required, create: `${timestamp()}`},
		"start_time":     Representation{repType: Required, create: `${timeadd(timestamp(), "-1m")}`},
		"limit":          Representation{repType: Required, create: `1`},
	}

	AuditEventResourceConfig = ""
)

func TestAuditAuditEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditAuditEventResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_audit_events.test_audit_events"

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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "end_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time"),

					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.event_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.event_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.event_source"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.event_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.event_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.principal_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.request_action"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.request_agent"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.request_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.request_origin"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.request_resource"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.response_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.response_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "audit_events.0.tenant_id"),
				),
			},
		},
	})
}
