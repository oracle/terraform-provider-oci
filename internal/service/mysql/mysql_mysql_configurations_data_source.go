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

func MysqlMysqlConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MysqlMysqlConfigurationResource()),
			},
		},
	}
}

func readMysqlMysqlConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.ListConfigurationsResponse
}

func (s *MysqlMysqlConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlConfigurationsDataSourceCrud) Get() error {
	request := oci_mysql.ListConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.ConfigurationLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_mysql.ListConfigurationsTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_mysql.ListConfigurationsTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

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

func (s *MysqlMysqlConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlMysqlConfigurationsDataSource-", MysqlMysqlConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlConfiguration := map[string]interface{}{}

		if r.CompartmentId != nil {
			mysqlConfiguration["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			mysqlConfiguration["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			mysqlConfiguration["description"] = *r.Description
		}

		if r.DisplayName != nil {
			mysqlConfiguration["display_name"] = *r.DisplayName
		}

		mysqlConfiguration["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlConfiguration["id"] = *r.Id
		}

		if r.ShapeName != nil {
			mysqlConfiguration["shape_name"] = *r.ShapeName
		}

		mysqlConfiguration["state"] = r.LifecycleState

		if r.SystemTags != nil {
			mysqlConfiguration["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			mysqlConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			mysqlConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		mysqlConfiguration["type"] = r.Type

		resources = append(resources, mysqlConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlMysqlConfigurationsDataSource().Schema["configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("configurations", resources); err != nil {
		return err
	}

	return nil
}
