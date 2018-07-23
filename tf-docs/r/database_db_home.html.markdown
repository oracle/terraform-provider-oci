---
layout: "oci"
page_title: "OCI: oci_database_db_home"
sidebar_current: "docs-oci-resource-database-db_home"
description: |-
Creates and manages an OCI DbHome
---

# oci_database_db_home
The `oci_database_db_home` resource creates and manages an OCI DbHome

Creates a new DB Home in the specified DB System based on the request parameters you provide.


## Example Usage

```hcl
resource "oci_database_db_home" "test_db_home" {
	#Required
	database {
		#Required
		admin_password = "${var.db_home_database_admin_password}"
		backup_id = "${oci_database_backup.test_backup.id}"
		backup_tde_password = "${var.db_home_database_backup_tde_password}"
		db_name = "${var.db_home_database_db_name}"

		#Optional
		character_set = "${var.db_home_database_character_set}"
		db_backup_config {

			#Optional
			auto_backup_enabled = "${var.db_home_database_db_backup_config_auto_backup_enabled}"
		}
		db_workload = "${var.db_home_database_db_workload}"
		defined_tags = "${var.db_home_database_defined_tags}"
		freeform_tags = "${var.db_home_database_freeform_tags}"
		ncharacter_set = "${var.db_home_database_ncharacter_set}"
		pdb_name = "${var.db_home_database_pdb_name}"
	}
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	db_version {
	}

	#Optional
	display_name = "${var.db_home_display_name}"
	source = "${var.db_home_source}"
}
```

## Argument Reference

The following arguments are supported:

* `database` - (Required) 
	* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	* `backup_id` - (Required) The backup OCID.
	* `backup_tde_password` - (Required) The password to open the TDE wallet.
	* `character_set` - (Optional) The character set for the database.  The default is AL32UTF8. Allowed values are:  AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
	* `db_backup_config` - (Optional) 
		* `auto_backup_enabled` - (Optional) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `db_name` - (Required) The database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	* `db_workload` - (Optional) Database workload type.
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `ncharacter_set` - (Optional) National character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
	* `pdb_name` - (Optional) Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `db_system_id` - (Required) The OCID of the DB System.
* `db_version` - (Required) (Updatable) A valid Oracle database version. To get a list of supported versions, use the [ListDbVersions](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbVersion/ListDbVersions) operation.
* `display_name` - (Optional) The user-provided name of the database home.
* `source` - (Optional) Source of database:   NONE for creating a new database   DB_BACKUP for creating a new database by restoring a backup 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_system_id` - The OCID of the DB System.
* `db_version` - The Oracle database version.
* `display_name` - The user-provided name for the database home. It does not need to be unique.
* `id` - The OCID of the database home.
* `last_patch_history_entry_id` - The OCID of the last patch history. This is updated as soon as a patch operation is started.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.

## Import

DbHomes can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_home.test_db_home "id"
```
