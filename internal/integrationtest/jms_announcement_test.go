// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsAnnouncementDataSourceRepresentation = map[string]interface{}{
		"summary_contains": acctest.Representation{RepType: acctest.Optional, Create: `random nonexisting text lorem ipsum dolor`},
		"time_end":         acctest.Representation{RepType: acctest.Optional, Create: `2023-05-20T01:00:00Z`},
		"time_start":       acctest.Representation{RepType: acctest.Optional, Create: `2023-05-01T01:00:00Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsAnnouncementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsAnnouncementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_announcements.test_announcements"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_announcements",
					"test_announcements",
					acctest.Optional,
					acctest.Create,
					JmsAnnouncementDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "summary_contains", "random nonexisting text lorem ipsum dolor"),
				resource.TestCheckResourceAttr(datasourceName, "time_end", "2023-05-20T01:00:00Z"),
				resource.TestCheckResourceAttr(datasourceName, "time_start", "2023-05-01T01:00:00Z"),

				// we expect it returns zero results because of summary_contains parameter value
				resource.TestCheckResourceAttr(datasourceName, "announcement_collection.0.items.#", "0"),
			),
		},
	})
}
