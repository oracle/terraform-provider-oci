// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDbVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_database_software_image_supported": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_upgrade_supported": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"storage_management": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_latest_for_major_version": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_preview_db_version": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_upgrade_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"supports_pdb": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbVersionsResponse
}

func (s *DatabaseDbVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbVersionsDataSourceCrud) Get() error {
	request := oci_database.ListDbVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if dbSystemShape, ok := s.D.GetOkExists("db_system_shape"); ok {
		tmp := dbSystemShape.(string)
		request.DbSystemShape = &tmp
	}

	if isDatabaseSoftwareImageSupported, ok := s.D.GetOkExists("is_database_software_image_supported"); ok {
		tmp := isDatabaseSoftwareImageSupported.(bool)
		request.IsDatabaseSoftwareImageSupported = &tmp
	}

	if isUpgradeSupported, ok := s.D.GetOkExists("is_upgrade_supported"); ok {
		tmp := isUpgradeSupported.(bool)
		request.IsUpgradeSupported = &tmp
	}

	if storageManagement, ok := s.D.GetOkExists("storage_management"); ok {
		request.StorageManagement = oci_database.DbSystemOptionsStorageManagementEnum(storageManagement.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbVersionsDataSource-", DatabaseDbVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbVersion := map[string]interface{}{}

		if r.IsLatestForMajorVersion != nil {
			dbVersion["is_latest_for_major_version"] = *r.IsLatestForMajorVersion
		}

		if r.IsPreviewDbVersion != nil {
			dbVersion["is_preview_db_version"] = *r.IsPreviewDbVersion
		}

		if r.IsUpgradeSupported != nil {
			dbVersion["is_upgrade_supported"] = *r.IsUpgradeSupported
		}

		if r.SupportsPdb != nil {
			dbVersion["supports_pdb"] = *r.SupportsPdb
		}

		if r.Version != nil {
			dbVersion["version"] = *r.Version
		}

		resources = append(resources, dbVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbVersionsDataSource().Schema["db_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_versions", resources); err != nil {
		return err
	}

	return nil
}
