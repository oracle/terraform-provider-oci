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

func MysqlMysqlDbSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlDbSystems,
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
			"database_management": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_heat_wave_cluster_attached": {
				Type:     schema.TypeBool,
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
				Elem:     tfresource.GetDataSourceItemSchema(MysqlMysqlDbSystemResource()),
			},
		},
	}
}

func readMysqlMysqlDbSystems(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
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

	if databaseManagement, ok := s.D.GetOkExists("database_management"); ok {
		interfaces := databaseManagement.([]interface{})
		tmp := make([]oci_mysql.ListDbSystemsDatabaseManagementEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_mysql.ListDbSystemsDatabaseManagementEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("database_management") {
			request.DatabaseManagement = tmp
		}
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isHeatWaveClusterAttached, ok := s.D.GetOkExists("is_heat_wave_cluster_attached"); ok {
		tmp := isHeatWaveClusterAttached.(bool)
		request.IsHeatWaveClusterAttached = &tmp
	}

	if isUpToDate, ok := s.D.GetOkExists("is_up_to_date"); ok {
		tmp := isUpToDate.(bool)
		request.IsUpToDate = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.DbSystemLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlMysqlDbSystemsDataSource-", MysqlMysqlDbSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlDbSystem := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			mysqlDbSystem["availability_domain"] = *r.AvailabilityDomain
		}

		if r.BackupPolicy != nil {
			mysqlDbSystem["backup_policy"] = []interface{}{BackupPolicyToMap(r.BackupPolicy)}
		} else {
			mysqlDbSystem["backup_policy"] = nil
		}

		mysqlDbSystem["crash_recovery"] = r.CrashRecovery

		if r.CurrentPlacement != nil {
			mysqlDbSystem["current_placement"] = []interface{}{DbSystemPlacementToMap(r.CurrentPlacement)}
		} else {
			mysqlDbSystem["current_placement"] = nil
		}

		mysqlDbSystem["database_management"] = r.DatabaseManagement

		if r.DefinedTags != nil {
			mysqlDbSystem["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DeletionPolicy != nil {
			mysqlDbSystem["deletion_policy"] = []interface{}{DeletionPolicyDetailsToMap(r.DeletionPolicy)}
		} else {
			mysqlDbSystem["deletion_policy"] = nil
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

		if r.HeatWaveCluster != nil {
			mysqlDbSystem["heat_wave_cluster"] = []interface{}{HeatWaveClusterSummaryToMap(r.HeatWaveCluster)}
		} else {
			mysqlDbSystem["heat_wave_cluster"] = nil
		}

		if r.Id != nil {
			mysqlDbSystem["id"] = *r.Id
		}

		if r.IsHeatWaveClusterAttached != nil {
			mysqlDbSystem["is_heat_wave_cluster_attached"] = *r.IsHeatWaveClusterAttached
		}

		if r.IsHighlyAvailable != nil {
			mysqlDbSystem["is_highly_available"] = *r.IsHighlyAvailable
		}

		if r.MysqlVersion != nil {
			mysqlDbSystem["mysql_version"] = *r.MysqlVersion
		}

		if r.ShapeName != nil {
			mysqlDbSystem["shape_name"] = *r.ShapeName
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlMysqlDbSystemsDataSource().Schema["db_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_systems", resources); err != nil {
		return err
	}

	return nil
}
