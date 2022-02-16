// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInstanceResource()),
			},
		},
	}
}

func readCoreInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstancesDataSourceCrud struct {
	tfresource.BaseCrud
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstancesResponse
}

func (s *CoreInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancesDataSourceCrud) Get() error {
	request := oci_core.ListInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstancesDataSource-", CoreInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AgentConfig != nil {
			instance["agent_config"] = []interface{}{InstanceAgentConfigToMap(r.AgentConfig)}
		} else {
			instance["agent_config"] = nil
		}

		if r.AvailabilityConfig != nil {
			instance["availability_config"] = []interface{}{InstanceAvailabilityConfigToMap(r.AvailabilityConfig)}
		} else {
			instance["availability_config"] = nil
		}

		if r.AvailabilityDomain != nil {
			instance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CapacityReservationId != nil {
			instance["capacity_reservation_id"] = *r.CapacityReservationId
		}

		if r.DedicatedVmHostId != nil {
			instance["dedicated_vm_host_id"] = *r.DedicatedVmHostId
		}

		if r.DefinedTags != nil {
			instance["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			instance["display_name"] = *r.DisplayName
		}

		if r.ExtendedMetadata != nil {
			instance["extended_metadata"] = convertNestedMapToFlatMap(r.ExtendedMetadata)
		}

		if r.FaultDomain != nil {
			instance["fault_domain"] = *r.FaultDomain
		}

		instance["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instance["id"] = *r.Id
		}

		if r.ImageId != nil {
			instance["image"] = *r.ImageId
		}

		if r.InstanceOptions != nil {
			instance["instance_options"] = []interface{}{InstanceOptionsToMap(r.InstanceOptions)}
		} else {
			instance["instance_options"] = nil
		}

		if r.IpxeScript != nil {
			instance["ipxe_script"] = *r.IpxeScript
		}

		instance["launch_mode"] = r.LaunchMode

		if r.LaunchOptions != nil {
			instance["launch_options"] = []interface{}{LaunchOptionsToMap(r.LaunchOptions)}
		} else {
			instance["launch_options"] = nil
		}

		if r.Metadata != nil {
			instance["metadata"] = r.Metadata
		}

		if r.PlatformConfig != nil {
			platformConfigArray := []interface{}{}
			if platformConfigMap := PlatformConfigToMap(&r.PlatformConfig); platformConfigMap != nil {
				platformConfigArray = append(platformConfigArray, platformConfigMap)
			}
			instance["platform_config"] = platformConfigArray
		} else {
			instance["platform_config"] = nil
		}

		if r.PreemptibleInstanceConfig != nil {
			instance["preemptible_instance_config"] = []interface{}{PreemptibleInstanceConfigDetailsToMap(r.PreemptibleInstanceConfig)}
		} else {
			instance["preemptible_instance_config"] = nil
		}

		if r.Region != nil {
			instance["region"] = *r.Region
		}

		if r.Shape != nil {
			instance["shape"] = *r.Shape
		}

		if r.ShapeConfig != nil {
			instance["shape_config"] = []interface{}{InstanceShapeConfigToMap(r.ShapeConfig)}
		} else {
			instance["shape_config"] = nil
		}

		if r.SourceDetails != nil {
			sourceDetailsArray := []interface{}{}
			if sourceDetailsMap := InstanceSourceDetailsToMap(&r.SourceDetails, nil, nil); sourceDetailsMap != nil {
				sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			}
			instance["source_details"] = sourceDetailsArray
		} else {
			instance["source_details"] = nil
		}

		instance["state"] = r.LifecycleState

		if r.SystemTags != nil {
			instance["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			instance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeMaintenanceRebootDue != nil {
			instance["time_maintenance_reboot_due"] = r.TimeMaintenanceRebootDue.String()
		}

		resources = append(resources, instance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		return err
	}

	return nil
}

func convertNestedMapToFlatMap(m map[string]interface{}) map[string]string {
	flatMap := make(map[string]string)
	var ok bool
	for key, val := range m {
		if flatMap[key], ok = val.(string); !ok {
			mapValStr, err := json.Marshal(val)
			if err != nil {
				mapValStr = []byte{}
			}
			flatMap[key] = string(mapValStr)
		}
	}
	return flatMap
}
