// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbNodeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbNode,
		Schema: map[string]*schema.Schema{
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// todo: // @codegen omits this
			"software_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readDbNode(d *schema.ResourceData, m interface{}) error {
	sync := &DbNodeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbNodeDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_database.DatabaseClient
	Res    *oci_database.DbNode
}

func (s *DbNodeDataSourceCrud) Get() error {
	request := oci_database.GetDbNodeRequest{}

	dbNodeId := s.D.Get("db_node_id")
	tmp := dbNodeId.(string)
	request.DbNodeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetDbNode(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbNode
	return nil
}

func (s *DbNodeDataSourceCrud) SetData() {
	s.D.SetId(*s.Res.Id)

	if s.Res.BackupVnicId != nil {
		s.D.Set("backup_vnic_id", *s.Res.BackupVnicId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	// todo: @codegen includes this but misses schema entry
	if s.Res.SoftwareStorageSizeInGB != nil {
		s.D.Set("software_storage_size_in_gb", *s.Res.SoftwareStorageSizeInGB)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

}
