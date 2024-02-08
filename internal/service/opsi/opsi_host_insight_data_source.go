// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiHostInsightDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["host_insight_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiHostInsightResource(), fieldMap, readSingularOpsiHostInsight)
}

func readSingularOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiHostInsightDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetHostInsightResponse
}

func (s *OpsiHostInsightDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiHostInsightDataSourceCrud) Get() error {
	request := oci_opsi.GetHostInsightRequest{}

	if hostInsightId, ok := s.D.GetOkExists("host_insight_id"); ok {
		tmp := hostInsightId.(string)
		request.HostInsightId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiHostInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiHostInsightsSingularDataSource-", OpsiHostInsightsDataSource(), s.D))
	switch v := (s.Res.HostInsight).(type) {
	case oci_opsi.EmManagedExternalHostInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_HOST")

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			s.D.Set("enterprise_manager_entity_display_name", *v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			s.D.Set("enterprise_manager_entity_identifier", *v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			s.D.Set("enterprise_manager_entity_name", *v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			s.D.Set("enterprise_manager_entity_type", *v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			s.D.Set("enterprise_manager_identifier", *v.EnterpriseManagerIdentifier)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
		}

		if v.PlatformName != nil {
			s.D.Set("platform_name", *v.PlatformName)
		}

		s.D.Set("platform_type", v.PlatformType)

		if v.PlatformVersion != nil {
			s.D.Set("platform_version", *v.PlatformVersion)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

	case oci_opsi.MacsManagedCloudHostInsight:
		s.D.Set("entity_source", "MACS_MANAGED_CLOUD_HOST")

		if v.GetCompartmentId() != nil {
			s.D.Set("compartment_id", *v.GetCompartmentId())
		}

		if v.GetHostName() != nil {
			s.D.Set("host_name", *v.GetHostName())
		}

		s.D.Set("freeform_tags", v.GetFreeformTags())

		if v.GetDefinedTags() != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.GetDefinedTags()))
		}

		if v.GetTimeCreated() != nil {
			s.D.Set("time_created", v.GetTimeCreated().String())
		}

		if v.GetHostDisplayName() != nil {
			s.D.Set("host_display_name", *v.GetHostDisplayName())
		}

		if v.GetHostType() != nil {
			s.D.Set("host_type", *v.GetHostType())
		}

		if v.GetProcessorCount() != nil {
			s.D.Set("processor_count", *v.GetProcessorCount())
		}

		if v.GetSystemTags() != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.GetSystemTags()))
		}

		if v.GetTimeUpdated() != nil {
			s.D.Set("time_updated", v.GetTimeUpdated().String())
		}

		if v.GetLifecycleDetails() != nil {
			s.D.Set("lifecycle_details", *v.GetLifecycleDetails())
		}

		s.D.Set("status", v.GetStatus())

		s.D.Set("state", v.GetLifecycleState())

	case oci_opsi.MacsManagedExternalHostInsight:
		s.D.Set("entity_source", "MACS_MANAGED_EXTERNAL_HOST")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.HostDisplayName != nil {
			s.D.Set("host_display_name", *v.HostDisplayName)
		}

		if v.HostName != nil {
			s.D.Set("host_name", *v.HostName)
		}

		if v.HostType != nil {
			s.D.Set("host_type", *v.HostType)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", s.Res.HostInsight)
		return nil
	}

	return nil
}
