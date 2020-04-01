// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterDatasource("oci_database_backup_destinations", DatabaseBackupDestinationsDataSource())
}

func DatabaseBackupDestinationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseBackupDestinations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_destinations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseBackupDestinationResource()),
			},
		},
	}
}

func readDatabaseBackupDestinations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseBackupDestinationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListBackupDestinationResponse
}

func (s *DatabaseBackupDestinationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseBackupDestinationsDataSourceCrud) Get() error {
	request := oci_database.ListBackupDestinationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListBackupDestination(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackupDestination(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseBackupDestinationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backupDestination := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		associatedDatabases := []interface{}{}
		for _, item := range r.AssociatedDatabases {
			associatedDatabases = append(associatedDatabases, AssociatedDatabaseDetailsToMap(item))
		}
		backupDestination["associated_databases"] = associatedDatabases

		if r.ConnectionString != nil {
			backupDestination["connection_string"] = *r.ConnectionString
		}

		if r.DefinedTags != nil {
			backupDestination["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			backupDestination["display_name"] = *r.DisplayName
		}

		backupDestination["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			backupDestination["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			backupDestination["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.LocalMountPointPath != nil {
			backupDestination["local_mount_point_path"] = *r.LocalMountPointPath
		}

		backupDestination["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			backupDestination["time_created"] = r.TimeCreated.String()
		}

		backupDestination["type"] = r.Type

		backupDestination["vpc_users"] = r.VpcUsers

		resources = append(resources, backupDestination)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseBackupDestinationsDataSource().Schema["backup_destinations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backup_destinations", resources); err != nil {
		return err
	}

	return nil
}
