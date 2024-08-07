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

func PsqlDefaultConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularPsqlDefaultConfiguration,
		Schema: map[string]*schema.Schema{
			"default_configuration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularPsqlDefaultConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDefaultConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlDefaultConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetDefaultConfigurationResponse
}

func (s *PsqlDefaultConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDefaultConfigurationDataSourceCrud) Get() error {
	request := oci_psql.GetDefaultConfigurationRequest{}

	if defaultConfigurationId, ok := s.D.GetOkExists("default_configuration_id"); ok {
		tmp := defaultConfigurationId.(string)
		request.DefaultConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetDefaultConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlDefaultConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", []interface{}{DefaultConfigurationDetailsToMap(s.Res.ConfigurationDetails)})
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.InstanceMemorySizeInGBs != nil {
		s.D.Set("instance_memory_size_in_gbs", *s.Res.InstanceMemorySizeInGBs)
	}

	if s.Res.InstanceOcpuCount != nil {
		s.D.Set("instance_ocpu_count", *s.Res.InstanceOcpuCount)
	}

	if s.Res.IsFlexible != nil {
		s.D.Set("is_flexible", *s.Res.IsFlexible)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
