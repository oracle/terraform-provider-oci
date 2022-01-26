// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigration,
		Read:     readDatabaseMigration,
		Delete:   deleteDatabaseMigration,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"additional_migrations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cloud_exadata_infrastructure_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_vm_cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigration(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseMigration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseMigrationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExadataDbSystemMigration
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseMigrationResourceCrud) ID() string {
	return *s.Res.DbSystemId
}

func (s *DatabaseMigrationResourceCrud) Create() error {
	request := oci_database.MigrateExadataDbSystemResourceModelRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.MigrateExadataDbSystemResourceModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataDbSystemMigration
	return nil
}

func (s *DatabaseMigrationResourceCrud) SetData() error {
	additionalMigrations := []interface{}{}
	for _, item := range s.Res.AdditionalMigrations {
		additionalMigrations = append(additionalMigrations, ExadataDbSystemMigrationSummaryToMap(item))
	}
	s.D.Set("additional_migrations", additionalMigrations)

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.CloudVmClusterId != nil {
		s.D.Set("cloud_vm_cluster_id", *s.Res.CloudVmClusterId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	return nil
}

func ExadataDbSystemMigrationSummaryToMap(obj oci_database.ExadataDbSystemMigrationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudExadataInfrastructureId != nil {
		result["cloud_exadata_infrastructure_id"] = string(*obj.CloudExadataInfrastructureId)
	}

	if obj.CloudVmClusterId != nil {
		result["cloud_vm_cluster_id"] = string(*obj.CloudVmClusterId)
	}

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	return result
}
