// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbNodes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
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
			"db_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DbNodeDataSource(),
			},
		},
	}
}

func readDbNodes(d *schema.ResourceData, m interface{}) error {
	sync := &DbNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbNodesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbNodesResponse
}

func (s *DbNodesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbNodesDataSourceCrud) Get() error {
	request := oci_database.ListDbNodesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbNodes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbNodes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbNodesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNode := map[string]interface{}{
			"db_system_id": *r.DbSystemId,
		}

		if r.BackupVnicId != nil {
			dbNode["backup_vnic_id"] = *r.BackupVnicId
		}

		if r.Hostname != nil {
			dbNode["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbNode["id"] = *r.Id
			dbNode["db_node_id"] = *r.Id // maintain legacy vanity id
		}

		// @CODEGEN not present in schema
		if r.SoftwareStorageSizeInGB != nil {
			dbNode["software_storage_size_in_gb"] = *r.SoftwareStorageSizeInGB
		}

		dbNode["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbNode["time_created"] = r.TimeCreated.String()
		}

		if r.VnicId != nil {
			dbNode["vnic_id"] = *r.VnicId
		}

		resources = append(resources, dbNode)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DbNodesDataSource().Schema["db_nodes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_nodes", resources); err != nil {
		panic(err)
	}

	return
}
