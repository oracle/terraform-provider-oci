// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"
)

func DataSafeTargetDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataSafeTargetDatabaseResource()),
			},
		},
	}
}

func readDataSafeTargetDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTargetDatabasesResponse
}

func (s *DataSafeTargetDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasesDataSourceCrud) Get() error {
	request := oci_data_safe.ListTargetDatabasesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListTargetDatabasesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if databaseType, ok := s.D.GetOkExists("database_type"); ok {
		request.DatabaseType = oci_data_safe.ListTargetDatabasesDatabaseTypeEnum(databaseType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureType, ok := s.D.GetOkExists("infrastructure_type"); ok {
		request.InfrastructureType = oci_data_safe.ListTargetDatabasesInfrastructureTypeEnum(infrastructureType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListTargetDatabasesLifecycleStateEnum(state.(string))
	}

	if targetDatabaseId, ok := s.D.GetOkExists("id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListTargetDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargetDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasesDataSource-", DataSafeTargetDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		targetDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			targetDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			targetDatabase["description"] = *r.Description
		}

		if r.DisplayName != nil {
			targetDatabase["display_name"] = *r.DisplayName
		}

		targetDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			targetDatabase["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			targetDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		targetDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			targetDatabase["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, targetDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeTargetDatabasesDataSource().Schema["target_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("target_databases", resources); err != nil {
		return err
	}

	return nil
}
