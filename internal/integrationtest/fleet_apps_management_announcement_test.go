// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementAnnouncementDataSourceRepresentation = map[string]interface{}{
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"summary_contains": acctest.Representation{RepType: acctest.Optional, Create: `summaryContains`},
	}

	FleetAppsManagementAnnouncementResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementAnnouncementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementAnnouncementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fleet_apps_management_announcements.test_announcements"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_announcements", "test_announcements", acctest.Required, acctest.Create, FleetAppsManagementAnnouncementDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementAnnouncementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "announcement_collection.#"),
				resource.TestMatchResourceAttr(datasourceName, "announcement_collection.0.items.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
