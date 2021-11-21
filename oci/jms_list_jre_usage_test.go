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
	listJreUsageSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"host_id":          Representation{RepType: Optional, Create: `my_host_id_1`},
		"application_id":   Representation{RepType: Optional, Create: `my_application_id_1`},
		"application_name": Representation{RepType: Optional, Create: `my_appplication_name_1`},
		"time_end":         Representation{RepType: Optional, Create: `2021-11-20T01:00:00Z`},
		"time_start":       Representation{RepType: Optional, Create: `2021-11-01T01:00:00Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsListJreUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsListJreUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_jms_list_jre_usage.test_list_jre_usage"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_jms_list_jre_usage", "test_list_jre_usage", Optional, Create, listJreUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_end", "2021-11-20T01:00:00Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_start", "2021-11-01T01:00:00Z"),

				// JRE usage can be zero
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
