// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const exadbXsCloneDatabasePassword = "BEstrO0ng_#11"

type exadbXsCloneSuite struct {
	expectedShapeAttribute     string
	sourceClusterVar           string
	remoteTargetClusterVar     string
	crossVaultTargetClusterVar string
	dbHomeVersion              string
	dbNamePrefix               string
}

type exadbXsCloneCase struct {
	name             string
	source           string
	isThinClone      bool
	targetClusterVar string
	crossVault       bool
}

func exadbXsBlockStorageCloneSuite() exadbXsCloneSuite {
	return exadbXsCloneSuite{
		expectedShapeAttribute:     "BLOCK_STORAGE",
		sourceClusterVar:           "exadb_vm_cluster_block_storage_id",
		remoteTargetClusterVar:     "exadb_vm_cluster_block_remote_target_id",
		crossVaultTargetClusterVar: "exadb_vm_cluster_block_cross_vault_target_id",
		dbHomeVersion:              "19.0.0.0",
		dbNamePrefix:               "B",
	}
}

func exadbXsSmartStorageCloneSuite() exadbXsCloneSuite {
	return exadbXsCloneSuite{
		expectedShapeAttribute:     "SMART_STORAGE",
		sourceClusterVar:           "exadb_vm_cluster_smart_storage_id",
		remoteTargetClusterVar:     "exadb_vm_cluster_smart_remote_target_id",
		crossVaultTargetClusterVar: "exadb_vm_cluster_smart_cross_vault_target_id",
		dbHomeVersion:              "26.0.0.0.0",
		dbNamePrefix:               "S",
	}
}

// Pre-requisite to run block storage clone tests:
//
//	set TF_VAR_exadb_vm_cluster_block_storage_id to an ExaDB-XS VM cluster with BLOCK_STORAGE shapeAttribute
//	set TF_VAR_exadb_vm_cluster_block_remote_target_id to a different BLOCK_STORAGE ExaDB-XS VM cluster sharing the source vault
//	set TF_VAR_exadb_vm_cluster_block_cross_vault_target_id to a different BLOCK_STORAGE ExaDB-XS VM cluster using a different vault
//
// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_local_thin_clone_from_snapshot(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thin_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      true,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_local_thick_clone_from_snapshot(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thick_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      false,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_remote_thin_clone_from_snapshot(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thin_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      true,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_remote_thick_clone_from_snapshot(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      false,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_local_thin_clone_from_another_db(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thin_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      true,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_local_thick_clone_from_another_db(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thick_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_remote_thin_clone_from_another_db(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thin_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      true,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_remote_thick_clone_from_another_db(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_block_storage_remote_thick_clone_from_another_db_cross_vault(t *testing.T) {
	suite := exadbXsBlockStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_another_db_cross_vault",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.crossVaultTargetClusterVar,
		crossVault:       true,
	})
}

// Pre-requisite to run smart storage clone tests:
//
//	set TF_VAR_exadb_vm_cluster_smart_storage_id to an ExaDB-XS VM cluster with SMART_STORAGE shapeAttribute
//	set TF_VAR_exadb_vm_cluster_smart_remote_target_id to a different SMART_STORAGE ExaDB-XS VM cluster sharing the source vault
//	set TF_VAR_exadb_vm_cluster_smart_cross_vault_target_id to a different SMART_STORAGE ExaDB-XS VM cluster using a different vault
//
// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_local_thin_clone_from_snapshot(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thin_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      true,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_local_thick_clone_from_snapshot(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thick_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      false,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_remote_thin_clone_from_snapshot(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thin_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      true,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_remote_thick_clone_from_snapshot(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_snapshot",
		source:           "DB_SNAPSHOT",
		isThinClone:      false,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_local_thin_clone_from_another_db(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thin_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      true,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_local_thick_clone_from_another_db(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "local_thick_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.sourceClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_remote_thin_clone_from_another_db(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thin_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      true,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_remote_thick_clone_from_another_db(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_another_db",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.remoteTargetClusterVar,
	})
}

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_exadbxs_smart_storage_remote_thick_clone_from_another_db_cross_vault(t *testing.T) {
	suite := exadbXsSmartStorageCloneSuite()
	testDatabaseDatabaseResourceExadbXsClone(t, suite, exadbXsCloneCase{
		name:             "remote_thick_clone_from_another_db_cross_vault",
		source:           "DB_CLONE",
		isThinClone:      false,
		targetClusterVar: suite.crossVaultTargetClusterVar,
		crossVault:       true,
	})
}

func testDatabaseDatabaseResourceExadbXsClone(t *testing.T, suite exadbXsCloneSuite, tc exadbXsCloneCase) {
	scenarioName := currentTestName()
	httpreplay.SetScenario(scenarioName)
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceDbName := suite.dbNamePrefix + "SRCDB"
	targetDbName := exadbXsCloneTargetDbName(suite, tc)
	fullConfig := config + compartmentIdVariableStr + exadbXsCloneConfig(suite, tc, sourceDbName, targetDbName)

	acctest.SaveConfigContent(fullConfig, "database", "database", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{
		{
			Config: fullConfig,
			Check:  exadbXsCloneChecks(suite, tc, sourceDbName, targetDbName),
		},
	})
}

func exadbXsCloneConfig(suite exadbXsCloneSuite, tc exadbXsCloneCase, sourceDbName string, targetDbName string) string {
	targetClusterVar := tc.targetClusterVar
	if targetClusterVar == "" {
		targetClusterVar = suite.sourceClusterVar
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf(`
variable "%s" {}
`, suite.sourceClusterVar))
	if targetClusterVar != suite.sourceClusterVar {
		b.WriteString(fmt.Sprintf(`
variable "%s" {}
`, targetClusterVar))
	}
	b.WriteString(fmt.Sprintf(`
data "oci_database_exadb_vm_cluster" "source" {
  exadb_vm_cluster_id = var.%s
}

data "oci_database_exadb_vm_cluster" "target" {
  exadb_vm_cluster_id = var.%s
}

resource "oci_database_db_home" "source_db_home" {
  db_system_id = var.%s
  db_version   = "%s"
  source       = "NONE"
  display_name = "tf-src-%s"
}

resource "oci_database_database" "source_database" {
  db_home_id = oci_database_db_home.source_db_home.id
  source     = "NONE"

  database {
    admin_password = "%s"
    db_name        = "%s"
%s
  }
}
`, suite.sourceClusterVar, targetClusterVar, suite.sourceClusterVar, suite.dbHomeVersion, strings.ToLower(suite.expectedShapeAttribute), exadbXsCloneDatabasePassword, sourceDbName, exadbXsCloneStorageSizeDetails(suite)))

	if tc.source == "DB_SNAPSHOT" {
		b.WriteString(`
resource "oci_database_database_snapshot" "source_snapshot" {
  database_id = oci_database_database.source_database.id
  name        = "tf-clone-source-snapshot"
}
`)
	}

	b.WriteString(fmt.Sprintf(`
resource "oci_database_db_home" "target_db_home" {
  db_system_id = var.%s
  db_version   = "%s"
  source       = "NONE"
  display_name = "tf-target-%s"
}

resource "oci_database_database" "cloned_database" {
  db_home_id = oci_database_db_home.target_db_home.id
  source     = "%s"

  database {
    admin_password       = "%s"
    db_name              = "%s"
    is_thin_clone        = %s
    source_admin_password = "%s"
    source_tde_password  = "%s"
    tde_wallet_password  = "%s"
%s%s
  }
}
`, targetClusterVar, suite.dbHomeVersion, strings.ToLower(suite.expectedShapeAttribute), tc.source, exadbXsCloneDatabasePassword, targetDbName, strconv.FormatBool(tc.isThinClone), exadbXsCloneDatabasePassword, exadbXsCloneDatabasePassword, exadbXsCloneDatabasePassword, exadbXsCloneSourceReference(tc), exadbXsCloneStorageSizeDetails(suite)))

	return b.String()
}

func exadbXsCloneStorageSizeDetails(suite exadbXsCloneSuite) string {
	if suite.expectedShapeAttribute != "BLOCK_STORAGE" {
		return ""
	}
	return `
    storage_size_details {
      data_storage_size_in_gb  = 60
      reco_storage_size_in_gbs = 44
    }
`
}

func exadbXsCloneSourceReference(tc exadbXsCloneCase) string {
	switch tc.source {
	case "DB_SNAPSHOT":
		return `    source_snapshot_id   = oci_database_database_snapshot.source_snapshot.id
`
	case "DB_CLONE":
		return `    source_database_id   = oci_database_database.source_database.id
`
	default:
		return ""
	}
}

func exadbXsCloneChecks(suite exadbXsCloneSuite, tc exadbXsCloneCase, sourceDbName string, targetDbName string) resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("data.oci_database_exadb_vm_cluster.source", "shape_attribute", suite.expectedShapeAttribute),
		resource.TestCheckResourceAttr("data.oci_database_exadb_vm_cluster.target", "shape_attribute", suite.expectedShapeAttribute),
		resource.TestCheckResourceAttr("oci_database_database.source_database", "source", "NONE"),
		resource.TestCheckResourceAttr("oci_database_database.source_database", "database.0.db_name", sourceDbName),
		resource.TestCheckResourceAttr("oci_database_database.cloned_database", "source", tc.source),
		resource.TestCheckResourceAttr("oci_database_database.cloned_database", "database.0.db_name", targetDbName),
		resource.TestCheckResourceAttr("oci_database_database.cloned_database", "database.0.is_thin_clone", strconv.FormatBool(tc.isThinClone)),
		resource.TestCheckResourceAttrSet("oci_database_database.cloned_database", "db_home_id"),
		resource.TestCheckResourceAttrSet("oci_database_database.cloned_database", "id"),
		resource.TestCheckResourceAttrSet("oci_database_database.cloned_database", "state"),
	}

	if suite.expectedShapeAttribute == "BLOCK_STORAGE" {
		checks = append(checks,
			resource.TestCheckResourceAttr("oci_database_database.cloned_database", "database.0.storage_size_details.#", "1"),
			resource.TestCheckResourceAttr("oci_database_database.cloned_database", "database.0.storage_size_details.0.data_storage_size_in_gb", "60"),
			resource.TestCheckResourceAttr("oci_database_database.cloned_database", "database.0.storage_size_details.0.reco_storage_size_in_gbs", "44"),
		)
	}

	switch tc.source {
	case "DB_SNAPSHOT":
		checks = append(checks,
			resource.TestCheckResourceAttrSet("oci_database_database.cloned_database", "database.0.source_snapshot_id"),
		)
	case "DB_CLONE":
		checks = append(checks,
			resource.TestCheckResourceAttrSet("oci_database_database.cloned_database", "database.0.source_database_id"),
		)
	}

	if tc.targetClusterVar != suite.sourceClusterVar {
		if tc.crossVault {
			checks = append(checks, testCheckResourceAttributesNotEqual(
				"data.oci_database_exadb_vm_cluster.source", "exascale_db_storage_vault_id",
				"data.oci_database_exadb_vm_cluster.target", "exascale_db_storage_vault_id"))
		} else {
			checks = append(checks, acctest.TestCheckResourceAttributesEqual(
				"data.oci_database_exadb_vm_cluster.source", "exascale_db_storage_vault_id",
				"data.oci_database_exadb_vm_cluster.target", "exascale_db_storage_vault_id"))
		}
	}

	return acctest.ComposeAggregateTestCheckFuncArrayWrapper(checks)
}

func testCheckResourceAttributesNotEqual(name1, key1, name2, key2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		val1, err := acctest.FromInstanceState(s, name1, key1)
		if err != nil {
			return err
		}
		val2, err := acctest.FromInstanceState(s, name2, key2)
		if err != nil {
			return err
		}
		if val1 == val2 {
			return fmt.Errorf("%s: attribute %q value %#v should not equal %s attribute %q value %#v", name1, key1, val1, name2, key2, val2)
		}
		return nil
	}
}

func exadbXsCloneTargetDbName(suite exadbXsCloneSuite, tc exadbXsCloneCase) string {
	// Examples: BLSNV = block/local/snapshot/thin/same-vault, BRDKX = block/remote/db-clone/thick/cross-vault,
	// SLSNV = smart/local/snapshot/thin/same-vault, SRDKX = smart/remote/db-clone/thick/cross-vault.
	location := "L"
	if tc.targetClusterVar != suite.sourceClusterVar {
		location = "R"
	}
	source := "D"
	if tc.source == "DB_SNAPSHOT" {
		source = "S"
	}
	thickness := "K"
	if tc.isThinClone {
		thickness = "N"
	}
	vault := "V"
	if tc.crossVault {
		vault = "X"
	}
	return suite.dbNamePrefix + location + source + thickness + vault
}

func currentTestName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "TestDatabaseDatabaseResource_exadbxs_clone"
	}
	fullName := runtime.FuncForPC(pc).Name()
	return fullName[strings.LastIndex(fullName, ".")+1:]
}
