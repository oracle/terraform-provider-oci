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
	JmsJavaDownloadsJavaLicenseSingularDataSourceRepresentation = map[string]interface{}{
		"license_type": acctest.Representation{RepType: acctest.Required, Create: `OTN`},
	}

	JmsJavaDownloadsJavaLicenseDataSourceRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `Oracle Technology Network`},
		"license_type": acctest.Representation{RepType: acctest.Optional, Create: `OTN`},
	}
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaLicenseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaLicenseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_java_downloads_java_licenses.test_java_licenses"
	singularDatasourceName := "data.oci_jms_java_downloads_java_license.test_java_license"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_licenses",
					"test_java_licenses",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaLicenseDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Oracle Technology Network"),
				resource.TestCheckResourceAttr(datasourceName, "license_type", "OTN"),

				resource.TestCheckResourceAttrSet(datasourceName, "java_license_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license",
					"test_java_license",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaLicenseSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_url"),
			),
		},
	})
}
