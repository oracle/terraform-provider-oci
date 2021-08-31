// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedInstanceEventReportSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                            Representation{repType: Required, create: `${var.compartment_id}`},
		"managed_instance_id":                       Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"latest_timestamp_greater_than_or_equal_to": Representation{repType: Optional, create: `latestTimestampGreaterThanOrEqualTo`},
		"latest_timestamp_less_than":                Representation{repType: Optional, create: `latestTimestampLessThan`},
	}

	ManagedInstanceEventReportResourceConfig = ManagedInstanceManagementResourceDependencies
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceEventReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceEventReportResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_osmanagement_managed_instance_event_report.test_managed_instance_event_report"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceEventReportResourceConfig,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] OS Management Resource should be created after 2 minutes as OS Agent takes time to activate")
					time.Sleep(5 * time.Minute)
					return nil
				},
			},
			// verify singular datasource
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_event_report", "test_managed_instance_event_report", Required, Create, managedInstanceEventReportSingularDataSourceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "counts"),
				),
			},
		},
	})
}
