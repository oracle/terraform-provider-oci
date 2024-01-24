// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseSqlTuningSetDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningSet,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"statement_counts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningSet(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningSetsResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningSetsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		tmp := owner.(string)
		request.Owner = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningSetDataSource-", DatabaseManagementManagedDatabaseSqlTuningSetDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningSetSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
