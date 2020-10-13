// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v27/mysql"
)

func init() {
	RegisterDatasource("oci_mysql_mysql_db_systems", MysqlMysqlDbSystemsDataSource())
}

func MysqlMysqlDbSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlDbSystems,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_up_to_date": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(MysqlMysqlDbSystemResource()),
			},
		},
	}
}

func readMysqlMysqlDbSystems(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	return ReadResource(sync)
}

type MysqlMysqlDbSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbSystemClient
	Res    *oci_mysql.ListDbSystemsResponse
}

func (s *MysqlMysqlDbSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlDbSystemsDataSourceCrud) Get() error {
	request := oci_mysql.ListDbSystemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isUpToDate, ok := s.D.GetOkExists("is_up_to_date"); ok {
		tmp := isUpToDate.(bool)
		request.IsUpToDate = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.DbSystemLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "mysql")

	response, err := s.Client.ListDbSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlMysqlDbSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlDbSystem := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			mysqlDbSystem["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			mysqlDbSystem["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			mysqlDbSystem["description"] = *r.Description
		}

		if r.DisplayName != nil {
			mysqlDbSystem["display_name"] = *r.DisplayName
		}

		endpoints := []interface{}{}
		for _, item := range r.Endpoints {
			endpoints = append(endpoints, DbSystemEndpointToMap(item))
		}
		mysqlDbSystem["endpoints"] = endpoints

		if r.FaultDomain != nil {
			mysqlDbSystem["fault_domain"] = *r.FaultDomain
		}

		mysqlDbSystem["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlDbSystem["id"] = *r.Id
		}

		if r.MysqlVersion != nil {
			mysqlDbSystem["mysql_version"] = *r.MysqlVersion
		}

		mysqlDbSystem["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			mysqlDbSystem["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			mysqlDbSystem["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, mysqlDbSystem)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MysqlMysqlDbSystemsDataSource().Schema["db_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_systems", resources); err != nil {
		return err
	}

	return nil
}
