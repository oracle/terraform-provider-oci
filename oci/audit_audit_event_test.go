// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	auditEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"end_time":       Representation{RepType: Required, Create: `${timestamp()}`},
		"start_time":     Representation{RepType: Required, Create: `${timeadd(timestamp(), "-1m")}`},
	}

	AuditEventResourceConfig = ""
)

// issue-routing-tag: audit/default
func TestAuditAuditEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditAuditEventResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_audit_events.test_audit_events"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_audit_events", "test_audit_events", Required, Create, auditEventDataSourceRepresentation) +
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
	})
}
