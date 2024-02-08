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

func OpsiExadataInsightDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["exadata_insight_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiExadataInsightResource(), fieldMap, readSingularOpsiExadataInsight)
}

func readSingularOpsiExadataInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiExadataInsightDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiExadataInsightDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetExadataInsightResponse
}

func (s *OpsiExadataInsightDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiExadataInsightDataSourceCrud) Get() error {
	request := oci_opsi.GetExadataInsightRequest{}

	if exadataInsightId, ok := s.D.GetOkExists("exadata_insight_id"); ok {
		tmp := exadataInsightId.(string)
		request.ExadataInsightId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetExadataInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiExadataInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiExadataInsightsSingularDataSource-", OpsiExadataInsightDataSource(), s.D))
	switch v := (s.Res.ExadataInsight).(type) {
	case oci_opsi.EmManagedExternalExadataInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_EXADATA")

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

		if v.IsAutoSyncEnabled != nil {
			s.D.Set("is_auto_sync_enabled", *v.IsAutoSyncEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.ExadataDisplayName != nil {
			s.D.Set("exadata_display_name", *v.ExadataDisplayName)
		}

		if v.ExadataName != nil {
			s.D.Set("exadata_name", *v.ExadataName)
		}

		s.D.Set("exadata_rack_type", v.ExadataRackType)

		s.D.Set("exadata_type", v.ExadataType)

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsVirtualizedExadata != nil {
			s.D.Set("is_virtualized_exadata", *v.IsVirtualizedExadata)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
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
	case oci_opsi.PeComanagedExadataInsight:
		s.D.Set("entity_source", "PE_COMANAGED_EXADATA")

		if v.ExadataInfraId != nil {
			s.D.Set("exadata_infra_id", *v.ExadataInfraId)
		}

		s.D.Set("exadata_infra_resource_type", v.ExadataInfraResourceType)

		if v.ExadataShape != nil {
			s.D.Set("exadata_shape", *v.ExadataShape)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.ExadataDisplayName != nil {
			s.D.Set("exadata_display_name", *v.ExadataDisplayName)
		}

		if v.ExadataName != nil {
			s.D.Set("exadata_name", *v.ExadataName)
		}

		s.D.Set("exadata_rack_type", v.ExadataRackType)

		s.D.Set("exadata_type", v.ExadataType)

		s.D.Set("freeform_tags", v.FreeformTags)
		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsVirtualizedExadata != nil {
			s.D.Set("is_virtualized_exadata", *v.IsVirtualizedExadata)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
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
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", s.Res.ExadataInsight)
		return nil
	}

	return nil
}
