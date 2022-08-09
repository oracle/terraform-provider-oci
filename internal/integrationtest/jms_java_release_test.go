// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsjavaReleaseSingularDataSourceRepresentation = map[string]interface{}{
		"release_version": acctest.Representation{RepType: acctest.Required, Create: `17.0.2`},
	}

	JmsjavaReleaseDataSourceRepresentation = map[string]interface{}{
		"family_version":      acctest.Representation{RepType: acctest.Optional, Create: `11`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `UPDATE_REQUIRED`},
		"license_type":        acctest.Representation{RepType: acctest.Optional, Create: `OTN`},
		"release_type":        acctest.Representation{RepType: acctest.Optional, Create: `CPU`},
		"release_version":     acctest.Representation{RepType: acctest.Optional, Create: `11.0.11`},
	}
)

// issue-routing-tag: jms/default
func TestJmsJavaReleaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaReleaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_jms_java_releases.test_java_releases"
	singularDatasourceName := "data.oci_jms_java_release.test_java_release"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_java_releases", "test_java_releases", acctest.Optional, acctest.Create, JmsjavaReleaseDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "family_version", "11"),
				resource.TestCheckResourceAttr(datasourceName, "jre_security_status", "UPDATE_REQUIRED"),
				resource.TestCheckResourceAttr(datasourceName, "license_type", "OTN"),
				resource.TestCheckResourceAttr(datasourceName, "release_type", "CPU"),
				resource.TestCheckResourceAttr(datasourceName, "release_version", "11.0.11"),

				resource.TestCheckResourceAttrSet(datasourceName, "java_release_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_java_release", "test_java_release", acctest.Required, acctest.Create, JmsjavaReleaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_version"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifacts.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_version", "17"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_notes_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "release_version", "17.0.2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_status"),
			),
		},
	})
}
