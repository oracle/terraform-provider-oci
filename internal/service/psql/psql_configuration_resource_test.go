// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// TestPsqlConfigurationImportPreservesOverrides reproduces
// https://github.com/oracle/terraform-provider-oci/issues/2668 : importing an
// existing oci_psql_configuration whose configuration_details report an
// overridden value must populate db_configuration_overrides, otherwise
// Terraform sees it as missing and plans a replacement (db_configuration_overrides
// is ForceNew).
func TestPsqlConfigurationImportPreservesOverrides(t *testing.T) {
	data, err := schema.InternalMap(PsqlConfigurationResource().Schema).Data(&terraform.InstanceState{
		ID: "ocid1.postgresqlconfiguration.oc1..example",
		Attributes: map[string]string{
			"compartment_id":               "ocid1.compartment.oc1..example",
			"db_version":                   "16",
			"display_name":                 "import-example",
			"db_configuration_overrides.#": "0",
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	crud := &PsqlConfigurationResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}

	configId := "ocid1.postgresqlconfiguration.oc1..example"
	configKey := "oci.admin_enabled_extensions"
	overridenValue := "vector,pglogical"
	crud.Res = &oci_psql.Configuration{
		Id: &configId,
		ConfigurationDetails: &oci_psql.ConfigurationDetails{
			Items: []oci_psql.ConfigParams{
				{
					ConfigKey:            &configKey,
					OverridenConfigValue: &overridenValue,
				},
			},
		},
	}

	if err := crud.SetData(); err != nil {
		t.Fatal(err)
	}

	overrides := data.Get("db_configuration_overrides").([]interface{})
	if len(overrides) == 0 {
		t.Fatal("import did not populate db_configuration_overrides from the API response, so Terraform will plan a replacement")
	}
}
