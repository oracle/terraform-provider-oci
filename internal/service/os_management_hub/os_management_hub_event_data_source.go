// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubEventDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["event_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubEventResource(), fieldMap, readSingularOsManagementHubEvent)
}

func readSingularOsManagementHubEvent(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubEventDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.EventClient
	Res    *oci_os_management_hub.GetEventResponse
}

func (s *OsManagementHubEventDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubEventDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetEventRequest{}

	if eventId, ok := s.D.GetOkExists("event_id"); ok {
		tmp := eventId.(string)
		request.EventId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubEventDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Event).(type) {
	case oci_os_management_hub.AgentEvent:
		s.D.Set("type", "AGENT")

		if v.Data != nil {
			s.D.Set("data", []interface{}{AgentEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.ExploitAttemptEvent:
		s.D.Set("type", "EXPLOIT_ATTEMPT")

		if v.Data != nil {
			s.D.Set("data", []interface{}{ExploitAttemptEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KernelCrashEvent:
		s.D.Set("type", "KERNEL_CRASH")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KernelEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KernelOopsEvent:
		s.D.Set("type", "KERNEL_OOPS")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KernelEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KspliceUpdateEvent:
		s.D.Set("type", "KSPLICE_UPDATE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KspliceUpdateEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.ManagementStationEvent:
		s.D.Set("type", "MANAGEMENT_STATION")

		if v.Data != nil {
			s.D.Set("data", []interface{}{ManagementStationEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.SoftwareSourceEvent:
		s.D.Set("type", "SOFTWARE_SOURCE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{SoftwareSourceEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.SoftwareUpdateEvent:
		s.D.Set("type", "SOFTWARE_UPDATE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{SoftwareUpdateEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.Event)
		return nil
	}

	return nil
}
