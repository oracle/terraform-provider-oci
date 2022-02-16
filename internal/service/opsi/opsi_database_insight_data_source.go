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

func OpsiDatabaseInsightDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_insight_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiDatabaseInsightResource(), fieldMap, readSingularOpsiDatabaseInsight)
}

func readSingularOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiDatabaseInsightDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.DatabaseInsight
}

func (s *OpsiDatabaseInsightDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiDatabaseInsightDataSourceCrud) Get() error {
	request := oci_opsi.GetDatabaseInsightRequest{}

	if databaseInsightId, ok := s.D.GetOkExists("database_insight_id"); ok {
		tmp := databaseInsightId.(string)
		request.DatabaseInsightId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseInsight
	return nil
}

func (s *OpsiDatabaseInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiDatabaseInsightsSingularDataSource-", OpsiDatabaseInsightsDataSource(), s.D))

	switch v := (*s.Res).(type) {
	case oci_opsi.EmManagedExternalDatabaseInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_DATABASE")

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

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
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
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", *s.Res)
	}

	return nil
}
