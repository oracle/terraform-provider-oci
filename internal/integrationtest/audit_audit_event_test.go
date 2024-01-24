// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AuditAuditauditEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"end_time":       acctest.Representation{RepType: acctest.Required, Create: `${timestamp()}`},
		"start_time":     acctest.Representation{RepType: acctest.Required, Create: `${timeadd(timestamp(), "-1m")}`},
	}

	AuditAuditEventResourceConfig = ""
)

// issue-routing-tag: audit/default
func TestAuditAuditEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAuditAuditEventResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_audit_events.test_audit_events"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_audit_events", "test_audit_events", acctest.Required, acctest.Create, AuditAuditauditEventDataSourceRepresentation) +
				compartmentIdVariableStr + AuditAuditEventResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
