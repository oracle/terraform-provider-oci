// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlMysqlBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"soft_delete": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_delete": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_preparation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"backup_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"retention_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mysql_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"immediate_source_backup_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"original_source_backup_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_copy_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_system_snapshot_summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"encrypt_data": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key_generation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readMysqlMysqlBackups(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbBackupsClient
	Res    *oci_mysql.ListBackupsResponse
}

func (s *MysqlMysqlBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlBackupsDataSourceCrud) Get() error {
	request := oci_mysql.ListBackupsRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	if backupPreparationStatus, ok := s.D.GetOkExists("backup_preparation_status"); ok {
		request.BackupPreparationStatus = oci_mysql.BackupValidationDetailsBackupPreparationStatusEnum(backupPreparationStatus.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if creationType, ok := s.D.GetOkExists("creation_type"); ok {
		request.CreationType = oci_mysql.BackupCreationTypeEnum(creationType.(string))
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if softDelete, ok := s.D.GetOkExists("soft_delete"); ok {
		request.SoftDelete = oci_mysql.ListBackupsSoftDeleteEnum(softDelete.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.BackupLifecycleStateEnum(state.(string))
	}

	if validationStatus, ok := s.D.GetOkExists("validation_status"); ok {
		request.ValidationStatus = oci_mysql.BackupValidationDetailsValidationStatusEnum(validationStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlMysqlBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlMysqlBackupsDataSource-", MysqlMysqlBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlBackup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		mysqlBackup["backup_preparation_status"] = oci_mysql.BackupValidationDetailsBackupPreparationStatusEnum(r.BackupPreparationStatus)

		if r.BackupSizeInGBs != nil {
			mysqlBackup["backup_size_in_gbs"] = *r.BackupSizeInGBs
		}

		mysqlBackup["backup_type"] = r.BackupType

		mysqlBackup["creation_type"] = r.CreationType

		if r.DataStorageSizeInGBs != nil {
			mysqlBackup["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DbSystemId != nil {
			mysqlBackup["db_system_id"] = *r.DbSystemId
		}

		if r.DbSystemSnapshotSummary != nil {
			mysqlBackup["db_system_snapshot_summary"] = []interface{}{DbSystemSnapshotSummaryToMap(r.DbSystemSnapshotSummary)}
		} else {
			mysqlBackup["db_system_snapshot_summary"] = nil
		}

		if r.DefinedTags != nil {
			mysqlBackup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			mysqlBackup["description"] = *r.Description
		}

		if r.DisplayName != nil {
			mysqlBackup["display_name"] = *r.DisplayName
		}

		if r.EncryptData != nil {
			mysqlBackup["encrypt_data"] = []interface{}{EncryptDataDetailsToMap(r.EncryptData)}
		} else {
			mysqlBackup["encrypt_data"] = nil
		}

		mysqlBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlBackup["id"] = *r.Id
		}

		if r.ImmediateSourceBackupId != nil {
			mysqlBackup["immediate_source_backup_id"] = *r.ImmediateSourceBackupId
		}

		if r.LifecycleDetails != nil {
			mysqlBackup["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MysqlVersion != nil {
			mysqlBackup["mysql_version"] = *r.MysqlVersion
		}

		if r.OriginalSourceBackupId != nil {
			mysqlBackup["original_source_backup_id"] = *r.OriginalSourceBackupId
		}

		if r.RetentionInDays != nil {
			mysqlBackup["retention_in_days"] = *r.RetentionInDays
		}

		if r.ShapeName != nil {
			mysqlBackup["shape_name"] = *r.ShapeName
		}

		mysqlBackup["soft_delete"] = r.SoftDelete

		mysqlBackup["state"] = r.LifecycleState

		if r.SystemTags != nil {
			mysqlBackup["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCopyCreated != nil {
			mysqlBackup["time_copy_created"] = r.TimeCopyCreated.String()
		}

		if r.TimeCreated != nil {
			mysqlBackup["time_created"] = r.TimeCreated.String()
		}

		mysqlBackup["validation_status"] = oci_mysql.BackupValidationDetailsBackupPreparationStatusEnum(r.ValidationStatus)

		resources = append(resources, mysqlBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlMysqlBackupsDataSource().Schema["backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backups", resources); err != nil {
		return err
	}

	return nil
}
