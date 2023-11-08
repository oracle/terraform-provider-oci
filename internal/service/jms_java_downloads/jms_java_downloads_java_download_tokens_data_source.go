// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadTokensDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaDownloadsJavaDownloadTokens,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"family_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_by_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_download_token_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(JmsJavaDownloadsJavaDownloadTokenResource()),
						},
					},
				},
			},
		},
	}
}

func readJmsJavaDownloadsJavaDownloadTokens(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokensDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadTokensDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.ListJavaDownloadTokensResponse
}

func (s *JmsJavaDownloadsJavaDownloadTokensDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadTokensDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.ListJavaDownloadTokensRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if searchByUser, ok := s.D.GetOkExists("search_by_user"); ok {
		tmp := searchByUser.(string)
		request.SearchByUser = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_jms_java_downloads.ListJavaDownloadTokensLifecycleStateEnum(state.(string))
	}

	if value, ok := s.D.GetOkExists("value"); ok {
		tmp := value.(string)
		request.Value = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.ListJavaDownloadTokens(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaDownloadTokens(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaDownloadsJavaDownloadTokensDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaDownloadTokensDataSource-", JmsJavaDownloadsJavaDownloadTokensDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaDownloadToken := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaDownloadTokenSummaryToMap(item))
	}
	javaDownloadToken["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaDownloadsJavaDownloadTokensDataSource().Schema["java_download_token_collection"].Elem.(*schema.Resource).Schema)
		javaDownloadToken["items"] = items
	}

	resources = append(resources, javaDownloadToken)
	if err := s.D.Set("java_download_token_collection", resources); err != nil {
		return err
	}

	return nil
}
