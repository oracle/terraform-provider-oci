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

func PsqlDefaultConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPsqlDefaultConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"default_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"configuration_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"allowed_values": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"config_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"default_config_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_overridable": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"is_restart_required": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"db_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_memory_size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_ocpu_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"is_flexible": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"shape": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readPsqlDefaultConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDefaultConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlDefaultConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.ListDefaultConfigurationsResponse
}

func (s *PsqlDefaultConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDefaultConfigurationsDataSourceCrud) Get() error {
	request := oci_psql.ListDefaultConfigurationsRequest{}

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

	response, err := s.Client.ListDefaultConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDefaultConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsqlDefaultConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlDefaultConfigurationsDataSource-", PsqlDefaultConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	defaultConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DefaultConfigurationSummaryToMap(item))
	}
	defaultConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsqlDefaultConfigurationsDataSource().Schema["default_configuration_collection"].Elem.(*schema.Resource).Schema)
		defaultConfiguration["items"] = items
	}

	resources = append(resources, defaultConfiguration)
	if err := s.D.Set("default_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}

func DefaultConfigParamsToMap(obj oci_psql.DefaultConfigParams) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedValues != nil {
		result["allowed_values"] = string(*obj.AllowedValues)
	}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.DefaultConfigValue != nil {
		result["default_config_value"] = string(*obj.DefaultConfigValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsOverridable != nil {
		result["is_overridable"] = bool(*obj.IsOverridable)
	}

	if obj.IsRestartRequired != nil {
		result["is_restart_required"] = bool(*obj.IsRestartRequired)
	}

	return result
}

func DefaultConfigurationDetailsToMap(obj *oci_psql.DefaultConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DefaultConfigParamsToMap(item))
	}
	result["items"] = items

	return result
}

func DefaultConfigurationSummaryToMap(obj oci_psql.DefaultConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceMemorySizeInGBs != nil {
		result["instance_memory_size_in_gbs"] = int(*obj.InstanceMemorySizeInGBs)
	}

	if obj.InstanceOcpuCount != nil {
		result["instance_ocpu_count"] = int(*obj.InstanceOcpuCount)
	}

	if obj.IsFlexible != nil {
		result["is_flexible"] = bool(*obj.IsFlexible)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
