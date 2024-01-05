// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OsmanagementOsmanagementManagedInstanceEventReportSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"managed_instance_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"latest_timestamp_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `latestTimestampGreaterThanOrEqualTo`},
		"latest_timestamp_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `latestTimestampLessThan`},
	}

	OsmanagementManagedInstanceEventReportResourceConfig = ManagedInstanceManagementResourceDependencies
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceEventReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceEventReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_osmanagement_managed_instance_event_report.test_managed_instance_event_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceEventReportResourceConfig,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] OS Management Resource should be created after 2 minutes as OS Agent takes time to activate")
				time.Sleep(5 * time.Minute)
				return nil
			},
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_event_report", "test_managed_instance_event_report", acctest.Required, acctest.Create, OsmanagementOsmanagementManagedInstanceEventReportSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "counts"),
			),
		},
	})
}
