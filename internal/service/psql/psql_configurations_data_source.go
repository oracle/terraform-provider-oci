// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPsqlConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(PsqlConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readPsqlConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.ListConfigurationsResponse
}

func (s *PsqlConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlConfigurationsDataSourceCrud) Get() error {
	request := oci_psql.ListConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_psql.ConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.ListConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsqlConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlConfigurationsDataSource-", PsqlConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	configuration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConfigurationSummaryToMap(item))
	}
	configuration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsqlConfigurationsDataSource().Schema["configuration_collection"].Elem.(*schema.Resource).Schema)
		configuration["items"] = items
	}

	resources = append(resources, configuration)
	if err := s.D.Set("configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
