// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import "testing"

func TestDbmulticloudResourceGraphProvidesRequiredDataSourceArguments(t *testing.T) {
	rootAssociations := dbmulticloudResourceGraph["oci_identity_compartment"]
	resourceTypes := map[string]bool{}
	for _, association := range rootAssociations {
		if association.ResourceClass == "oci_dbmulticloud_multi_cloud_resource_discovery" {
			resourceTypes[association.DatasourceQueryParams["resource_type"]] = true
		}
	}

	for _, resourceType := range []string{"'VAULTS'", "'STORAGE'", "'GCP_KEY_RINGS'", "'AWS_KEYS'"} {
		if !resourceTypes[resourceType] {
			t.Errorf("resource discovery export is missing resource_type %s", resourceType)
		}
	}

	vaultAssociations := dbmulticloudResourceGraph["oci_dbmulticloud_oracle_db_azure_vault"]
	for _, association := range vaultAssociations {
		if association.ResourceClass == "oci_dbmulticloud_oracle_db_azure_vault_association" && association.DatasourceQueryParams["oracle_db_azure_vault_id"] == "id" {
			return
		}
	}

	t.Error("vault association export is missing the parent vault ID query parameter")
}
