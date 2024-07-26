// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuCollectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fsu_collection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetSoftwareUpdateFsuCollectionResource(), fieldMap, readSingularFleetSoftwareUpdateFsuCollection)
}

func readSingularFleetSoftwareUpdateFsuCollection(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

type FleetSoftwareUpdateFsuCollectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.GetFsuCollectionResponse
}

func (s *FleetSoftwareUpdateFsuCollectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuCollectionDataSourceCrud) Get() error {
	request := oci_fleet_software_update.GetFsuCollectionRequest{}

	if fsuCollectionId, ok := s.D.GetOkExists("fsu_collection_id"); ok {
		tmp := fsuCollectionId.(string)
		request.FsuCollectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.GetFsuCollection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetSoftwareUpdateFsuCollectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.FsuCollection).(type) {
	case oci_fleet_software_update.DbCollection:
		s.D.Set("type", "DB")

		if v.FleetDiscovery != nil {
			fleetDiscoveryArray := []interface{}{}
			if fleetDiscoveryMap := DbFleetDiscoveryDetailsToMap(&v.FleetDiscovery); fleetDiscoveryMap != nil {
				fleetDiscoveryArray = append(fleetDiscoveryArray, fleetDiscoveryMap)
			}
			s.D.Set("fleet_discovery", fleetDiscoveryArray)
		} else {
			s.D.Set("fleet_discovery", nil)
		}

		s.D.Set("source_major_version", v.SourceMajorVersion)

		if v.ActiveFsuCycle != nil {
			s.D.Set("active_fsu_cycle", []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)})
		} else {
			s.D.Set("active_fsu_cycle", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("service_type", v.ServiceType)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetCount != nil {
			s.D.Set("target_count", *v.TargetCount)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_fleet_software_update.GiCollection:
		s.D.Set("type", "GI")

		if v.FleetDiscovery != nil {
			fleetDiscoveryArray := []interface{}{}
			if fleetDiscoveryMap := GiFleetDiscoveryDetailsToMap(&v.FleetDiscovery); fleetDiscoveryMap != nil {
				fleetDiscoveryArray = append(fleetDiscoveryArray, fleetDiscoveryMap)
			}
			s.D.Set("fleet_discovery", fleetDiscoveryArray)
		} else {
			s.D.Set("fleet_discovery", nil)
		}

		s.D.Set("source_major_version", v.SourceMajorVersion)

		if v.ActiveFsuCycle != nil {
			s.D.Set("active_fsu_cycle", []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)})
		} else {
			s.D.Set("active_fsu_cycle", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("service_type", v.ServiceType)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetCount != nil {
			s.D.Set("target_count", *v.TargetCount)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.FsuCollection)
		return nil
	}

	return nil
}
