// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeHostDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreComputeHost,
		Schema: map[string]*schema.Schema{
			"compute_host_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"additional_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_host_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"check_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"configuration_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"firmware_bundle_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recycle_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_last_apply": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"configuration_state": {
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
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"firmware_bundle_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hpc_island_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"impacted_component_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"local_block_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpu_memory_fabric_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_block_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recycle_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_host_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recycle_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_configuration_check": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreComputeHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeHostDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeHostResponse
}

func (s *CoreComputeHostDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeHostDataSourceCrud) Get() error {
	request := oci_core.GetComputeHostRequest{}

	if computeHostId, ok := s.D.GetOkExists("compute_host_id"); ok {
		tmp := computeHostId.(string)
		request.ComputeHostId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeHostDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_data", readData(s.Res.AdditionalData))

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeHostGroupId != nil {
		s.D.Set("compute_host_group_id", *s.Res.ComputeHostGroupId)
	}

	if s.Res.ConfigurationData != nil {
		s.D.Set("configuration_data", []interface{}{ComputeHostConfigurationDataToMap(s.Res.ConfigurationData)})
	} else {
		s.D.Set("configuration_data", nil)
	}

	s.D.Set("configuration_state", s.Res.ConfigurationState)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.FirmwareBundleId != nil {
		s.D.Set("firmware_bundle_id", *s.Res.FirmwareBundleId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("health", s.Res.Health)

	if s.Res.HpcIslandId != nil {
		s.D.Set("hpc_island_id", *s.Res.HpcIslandId)
	}

	s.D.Set("impacted_component_details", readData(s.Res.ImpactedComponentDetails))

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.LocalBlockId != nil {
		s.D.Set("local_block_id", *s.Res.LocalBlockId)
	}

	if s.Res.GpuMemoryFabricId != nil {
		s.D.Set("gpu_memory_fabric_id", *s.Res.GpuMemoryFabricId)
	}

	if s.Res.NetworkBlockId != nil {
		s.D.Set("network_block_id", *s.Res.NetworkBlockId)
	}

	if s.Res.Platform != nil {
		s.D.Set("platform", *s.Res.Platform)
	}

	if s.Res.RecycleDetails != nil {
		s.D.Set("recycle_details", []interface{}{RecycleDetailsToMap(s.Res.RecycleDetails)})
	} else {
		s.D.Set("recycle_details", nil)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeConfigurationCheck != nil {
		s.D.Set("time_configuration_check", s.Res.TimeConfigurationCheck.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func readData(Data interface{}) string {
	buf, err := json.Marshal(Data)
	if err != nil {
		log.Printf("error Marshalling Data: %v", err)
	}
	return string(buf)
}
