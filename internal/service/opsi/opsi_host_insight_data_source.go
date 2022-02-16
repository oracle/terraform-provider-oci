// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
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
	Res    *oci_opsi.HostInsight
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

	s.Res = &response.HostInsight
	return nil
}

func (s *OpsiHostInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiHostInsightsSingularDataSource-", OpsiHostInsightsDataSource(), s.D))

	switch v := (*s.Res).(type) {
	case oci_opsi.MacsManagedExternalHostInsight:
		s.D.Set("entity_source", "MACS_MANAGED_EXTERNAL_HOST")

		s.D.SetId(*v.GetId())

		if v.GetCompartmentId() != nil {
			s.D.Set("compartment_id", *v.GetCompartmentId())
		}

		if v.GetDefinedTags() != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.GetDefinedTags()))
		}

		s.D.Set("freeform_tags", v.GetFreeformTags())

		if v.GetHostDisplayName() != nil {
			s.D.Set("host_display_name", *v.GetHostDisplayName())
		}

		if v.GetHostName() != nil {
			s.D.Set("host_name", *v.GetHostName())
		}

		if v.GetHostType() != nil {
			s.D.Set("host_type", *v.GetHostType())
		}

		if v.GetLifecycleDetails() != nil {
			s.D.Set("lifecycle_details", *v.GetLifecycleDetails())
		}

		if v.GetProcessorCount() != nil {
			s.D.Set("processor_count", *v.GetProcessorCount())
		}

		s.D.Set("state", v.GetLifecycleState())

		s.D.Set("status", v.GetStatus())

		if v.GetSystemTags() != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.GetSystemTags()))
		}

		if v.GetTimeCreated() != nil {
			s.D.Set("time_created", v.GetTimeCreated().String())
		}

		if v.GetTimeUpdated() != nil {
			s.D.Set("time_updated", v.GetTimeUpdated().String())
		}

	case oci_opsi.EmManagedExternalHostInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_HOST")

		s.D.SetId(*v.GetId())

		if v.GetCompartmentId() != nil {
			s.D.Set("compartment_id", *v.GetCompartmentId())
		}

		if v.GetDefinedTags() != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.GetDefinedTags()))
		}

		s.D.Set("freeform_tags", v.GetFreeformTags())

		if v.GetHostDisplayName() != nil {
			s.D.Set("host_display_name", *v.GetHostDisplayName())
		}

		if v.GetHostName() != nil {
			s.D.Set("host_name", *v.GetHostName())
		}

		if v.GetHostType() != nil {
			s.D.Set("host_type", *v.GetHostType())
		}

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

		if v.GetLifecycleDetails() != nil {
			s.D.Set("lifecycle_details", *v.GetLifecycleDetails())
		}

		if v.PlatformName != nil {
			s.D.Set("platform_name", *v.PlatformName)
		}

		s.D.Set("platform_type", v.PlatformType)

		if v.PlatformVersion != nil {
			s.D.Set("platform_name", *v.PlatformVersion)
		}

		if v.GetProcessorCount() != nil {
			s.D.Set("processor_count", *v.GetProcessorCount())
		}

		s.D.Set("state", v.GetLifecycleState())

		s.D.Set("status", v.GetStatus())

		if v.GetSystemTags() != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.GetSystemTags()))
		}

		if v.GetTimeCreated() != nil {
			s.D.Set("time_created", v.GetTimeCreated().String())
		}

		if v.GetTimeUpdated() != nil {
			s.D.Set("time_updated", v.GetTimeUpdated().String())
		}

	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}
