// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiImportableAgentEntitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiImportableAgentEntity,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"entity_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_agent_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_agent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOpsiImportableAgentEntities(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiImportableAgentEntitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiImportableAgentEntitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListImportableAgentEntitiesResponse
}

func (s *OpsiImportableAgentEntitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiImportableAgentEntitiesDataSourceCrud) Get() error {
	request := oci_opsi.ListImportableAgentEntitiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListImportableAgentEntities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListImportableAgentEntities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiImportableAgentEntitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiImportableAgentEntitiesDataSource-", OpsiImportableAgentEntitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	importableAgentEntity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ImportableAgentEntitySummaryToMap(item))
	}
	importableAgentEntity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiImportableAgentEntitiesDataSource().Schema["importable_agent_entity_summary_collection"].Elem.(*schema.Resource).Schema)
		importableAgentEntity["items"] = items
	}

	resources = append(resources, importableAgentEntity)
	if err := s.D.Set("importable_agent_entity_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ImportableAgentEntitySummaryToMap(obj oci_opsi.ImportableAgentEntitySummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.HostImportableAgentEntitySummary:
		result["entity_source"] = "MACS_MANAGED_EXTERNAL_HOST"

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		result["platform_type"] = string(v.PlatformType)

		if v.ManagementAgentDisplayName != nil {
			result["management_agent_display_name"] = string(*v.ManagementAgentDisplayName)
		}

		if v.ManagementAgentId != nil {
			result["management_agent_id"] = string(*v.ManagementAgentId)
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
		return nil
	}

	return result
}
