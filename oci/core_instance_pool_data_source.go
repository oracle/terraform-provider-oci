// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstancePoolDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularInstancePool,
		Schema: map[string]*schema.Schema{
			"instance_pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"instance_configuration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"placement_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"availability_domain": {
							Type:     schema.TypeString,
							Required: true,
						},
						"primary_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"secondary_vnic_subnets": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"size": {
				Type:     schema.TypeInt,
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

func readSingularInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

type InstancePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.GetInstancePoolResponse
}

func (s *InstancePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstancePoolDataSourceCrud) Get() error {
	request := oci_core.GetInstancePoolRequest{}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *InstancePoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, InstancePoolPlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
