// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabaseVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_component": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_container_database_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_apps": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"end_of_support": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_certified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"release_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_app_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousContainerDatabaseVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousContainerDatabaseVersionsResponse
}

func (s *DatabaseAutonomousContainerDatabaseVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseVersionsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousContainerDatabaseVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if serviceComponent, ok := s.D.GetOkExists("service_component"); ok {
		request.ServiceComponent = oci_database.ListAutonomousContainerDatabaseVersionsServiceComponentEnum(serviceComponent.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousContainerDatabaseVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousContainerDatabaseVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousContainerDatabaseVersionsDataSource-", DatabaseAutonomousContainerDatabaseVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousContainerDatabaseVersion := map[string]interface{}{}

		if r.Details != nil {
			autonomousContainerDatabaseVersion["details"] = *r.Details
		}

		supportedApps := []interface{}{}
		for _, item := range r.SupportedApps {
			supportedApps = append(supportedApps, AppVersionSummaryToMap(item))
		}
		autonomousContainerDatabaseVersion["supported_apps"] = supportedApps

		if r.Version != nil {
			autonomousContainerDatabaseVersion["version"] = *r.Version
		}

		resources = append(resources, autonomousContainerDatabaseVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerDatabaseVersionsDataSource().Schema["autonomous_container_database_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_database_versions", resources); err != nil {
		return err
	}

	return nil
}

func AppVersionSummaryToMap(obj oci_database.AppVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndOfSupport != nil {
		result["end_of_support"] = string(*obj.EndOfSupport)
	}

	if obj.IsCertified != nil {
		result["is_certified"] = bool(*obj.IsCertified)
	}

	if obj.ReleaseDate != nil {
		result["release_date"] = string(*obj.ReleaseDate)
	}

	if obj.SupportedAppName != nil {
		result["supported_app_name"] = string(*obj.SupportedAppName)
	}

	return result
}
