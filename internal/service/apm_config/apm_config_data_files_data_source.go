// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmConfigDataFilesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readApmConfigDataFilesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"apm_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_modified_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_modified_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Manually defining this as the generated code returns DataFileSummaryCollection, but we actually want
			// The list of DataFileSummary defined within DataFileSummaryCollection response.
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"apm_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"md5": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Elem:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readApmConfigDataFilesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ApmConfigDataFilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_config.ConfigClient
	Res    *oci_apm_config.ListDataFilesResponse
}

func (s *ApmConfigDataFilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmConfigDataFilesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_apm_config.ListDataFilesRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if apmType, ok := s.D.GetOkExists("apm_type"); ok {
		tmp := apmType.(string)
		request.ApmType = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp, err := metaDataObjectMapToMetaDataJsonString(metadata.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("could not convert map to string %q", err)
		}
		request.Metadata = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if timeLastModifiedAfter, ok := s.D.GetOkExists("time_last_modified_after"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastModifiedAfter.(string))
		if err != nil {
			return err
		}
		request.TimeLastModifiedAfter = &oci_common.SDKTime{Time: tmp}
	}

	if timeLastModifiedBefore, ok := s.D.GetOkExists("time_last_modified_before"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastModifiedBefore.(string))
		if err != nil {
			return err
		}
		request.TimeLastModifiedBefore = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_config")

	response, err := s.Client.ListDataFiles(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataFiles(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmConfigDataFilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmConfigDataFilesDataSource-", ApmConfigDataFilesDataSource(), s.D))

	items := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataFileSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFilters(f.(*schema.Set), items, ApmConfigDataFilesDataSource().Schema["items"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("items", items); err != nil {
		return err
	}

	return nil
}
