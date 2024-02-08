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

func OpsiImportableComputeEntitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiImportableComputeEntity,
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
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compute_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compute_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
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

func readOpsiImportableComputeEntities(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiImportableComputeEntitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiImportableComputeEntitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListImportableComputeEntitiesResponse
}

func (s *OpsiImportableComputeEntitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiImportableComputeEntitiesDataSourceCrud) Get() error {
	request := oci_opsi.ListImportableComputeEntitiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListImportableComputeEntities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListImportableComputeEntities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiImportableComputeEntitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiImportableComputeEntitiesDataSource-", OpsiImportableComputeEntitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	importableComputeEntity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ImportableComputeEntitySummaryToMap(item))
	}
	importableComputeEntity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiImportableComputeEntitiesDataSource().Schema["importable_compute_entity_summary_collection"].Elem.(*schema.Resource).Schema)
		importableComputeEntity["items"] = items
	}

	resources = append(resources, importableComputeEntity)
	if err := s.D.Set("importable_compute_entity_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ImportableComputeEntitySummaryToMap(obj oci_opsi.ImportableComputeEntitySummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.CloudImportableComputeEntitySummary:
		result["entity_source"] = "MACS_MANAGED_CLOUD_HOST"

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		result["platform_type"] = string(v.PlatformType)

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.ComputeDisplayName != nil {
			result["compute_display_name"] = string(*v.ComputeDisplayName)
		}

		if v.ComputeId != nil {
			result["compute_id"] = string(*v.ComputeId)
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", obj)
		return nil
	}

	return result
}
