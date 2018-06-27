// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DatabaseDataSource(),
			},
		},
	}
}

func readDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDatabasesResponse
}

func (s *DatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabasesDataSourceCrud) Get() error {
	request := oci_database.ListDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	// @CODEGEN "page" was never wired up, omit

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabasesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		database := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"db_home_id":     *r.DbHomeId,
		}

		if r.CharacterSet != nil {
			database["character_set"] = *r.CharacterSet
		}

		if r.DbBackupConfig != nil {
			database["db_backup_config"] = []interface{}{dbBackupConfigToMap(r.DbBackupConfig)}
		}

		if r.DbName != nil {
			database["db_name"] = *r.DbName
		}

		if r.DbUniqueName != nil {
			database["db_unique_name"] = *r.DbUniqueName
		}

		if r.DbWorkload != nil {
			database["db_workload"] = *r.DbWorkload
		}

		if r.Id != nil {
			database["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			database["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NcharacterSet != nil {
			database["ncharacter_set"] = *r.NcharacterSet
		}

		if r.PdbName != nil {
			database["pdb_name"] = *r.PdbName
		}

		database["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			database["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, database)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabasesDataSource().Schema["databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("databases", resources); err != nil {
		panic(err)
	}

	return
}
