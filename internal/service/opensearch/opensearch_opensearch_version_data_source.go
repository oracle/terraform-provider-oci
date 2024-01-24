// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpensearchOpensearchVersion,
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
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_opensearch_opensearch_version", "oci_opensearch_opensearch_versions"),
	}
}

func readSingularOpensearchOpensearchVersion(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

type OpensearchOpensearchVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opensearch.OpensearchClusterClient
	Res    *oci_opensearch.ListOpensearchVersionsResponse
}

func (s *OpensearchOpensearchVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpensearchOpensearchVersionDataSourceCrud) Get() error {
	request := oci_opensearch.ListOpensearchVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opensearch")

	response, err := s.Client.ListOpensearchVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpensearchOpensearchVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpensearchOpensearchVersionDataSource-", OpensearchOpensearchVersionDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OpensearchVersionsSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
