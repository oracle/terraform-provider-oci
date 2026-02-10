package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ignoreRemediationRunDefinedTagsChangesRepresentationStandby1 = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`,
			`db_home[0].database[0].protection_mode`, `db_home[0].database[0].transport_type`}},
	}

	ignoreRemediationRunDefinedTagsChangesRepresentationStandby2 = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`,
			`db_home[0].database[0].protection_mode`, `db_home[0].database[0].transport_type`}},
	}

	ignoreRemediationRunDefinedTagsChangesRepresentationPrimary = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DBSystemRepresentationPrimary = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myoracledbprimary`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemAmdPrimary`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemOptionsPrimary},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDbHomeGroupPrimary},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRemediationRunDefinedTagsChangesRepresentationPrimary},
	}

	DbSystemOptionsPrimary = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `LVM`},
	}

	DbSystemAmdDbHomeGroupPrimary = map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `tfDbHomePrimary`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDatabaseGroupPrimary},
	}

	DbSystemAmdDatabaseGroupPrimary = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"character_set":  acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set": acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":    acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"pdb_name":       acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
	}

	DBSystemRepresentationStandby1 = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myoracledb1`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"source":                  acctest.Representation{RepType: acctest.Optional, Create: `DATAGUARD`},
		"primary_db_system_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system_primary.id}`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemAmdStandby1`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemOptionsStandby1},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDbHomeGroupStandby1},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRemediationRunDefinedTagsChangesRepresentationStandby1},
	}

	DbSystemAmdDbHomeGroupStandby1 = map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `tfDbHomeStandby1`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDatabaseGroupStandby1},
	}

	DbSystemOptionsStandby1 = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `LVM`},
	}

	DbSystemAmdDatabaseGroupStandby1 = map[string]interface{}{
		"admin_password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"tde_wallet_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":             acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"character_set":       acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set":      acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":         acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"pdb_name":            acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
		"protection_mode":     acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`},
		"transport_type":      acctest.Representation{RepType: acctest.Required, Create: `ASYNC`},
	}

	DBSystemRepresentationStandby2 = map[string]interface{}{
		"depends_on":              acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_data_guard_standby_1"}},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myOracleDBStandby2`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"source":                  acctest.Representation{RepType: acctest.Required, Create: `DATAGUARD`},
		"primary_db_system_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system_primary.id}`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemAmdStandby2`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemOptionsStandby2},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDbHomeGroupStandby2},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRemediationRunDefinedTagsChangesRepresentationStandby2},
	}

	DbSystemOptionsStandby2 = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `LVM`},
	}

	DbSystemAmdDbHomeGroupStandby2 = map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `tfDbHomeStandby2`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDatabaseGroupStandby2},
	}

	DbSystemAmdDatabaseGroupStandby2 = map[string]interface{}{
		"admin_password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"tde_wallet_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":             acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"character_set":       acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set":      acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":         acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"pdb_name":            acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
		"protection_mode":     acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`},
		"transport_type":      acctest.Representation{RepType: acctest.Required, Create: `ASYNC`},
	}

	DGConfigPrimary = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system_primary", acctest.Optional, acctest.Create, DBSystemRepresentationPrimary)
	DGConfig1       = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_data_guard_standby_1", acctest.Optional, acctest.Create, DBSystemRepresentationStandby1)
	DGConfig2       = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_data_guard_standby_2", acctest.Optional, acctest.Create, DBSystemRepresentationStandby2)
)

func TestResourceDbSystemDataGuardGroup(t *testing.T) {
	httpreplay.SetScenario("TestResourceDbSystemDataGuardGroup")
	defer httpreplay.SaveScenario()

	resourceNamePrimary := "oci_database_db_system.test_db_system_primary"
	resourceNameStandby1 := "oci_database_db_system.test_data_guard_standby_1"
	resourceNameStandby2 := "oci_database_db_system.test_data_guard_standby_2"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: ResourceDatabaseBaseConfig + DGConfigPrimary,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "id"),
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "time_created"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(resourceNamePrimary, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(resourceNamePrimary, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
			),
		},
		{
			Config: ResourceDatabaseBaseConfig + DGConfigPrimary + DGConfig1,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameStandby1, "id"),
				resource.TestCheckResourceAttrSet(resourceNameStandby1, "time_created"),
				resource.TestCheckResourceAttr(resourceNameStandby1, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceNameStandby1, "cpu_core_count", "2"),
			),
		},
		{
			Config: ResourceDatabaseBaseConfig + DGConfigPrimary + DGConfig1 + DGConfig2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameStandby2, "id"),
				resource.TestCheckResourceAttrSet(resourceNameStandby2, "time_created"),
				resource.TestCheckResourceAttr(resourceNameStandby2, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceNameStandby2, "cpu_core_count", "2"),
			),
		},
	})
}
