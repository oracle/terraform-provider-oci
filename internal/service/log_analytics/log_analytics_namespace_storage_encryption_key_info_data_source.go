// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespaceStorageEncryptionKeyInfo,
		Schema: map[string]*schema.Schema{
			"namespace": {
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
						"key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceStorageEncryptionKeyInfo(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListEncryptionKeyInfoResponse
}

func (s *LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSourceCrud) Get() error {
	request := oci_log_analytics.ListEncryptionKeyInfoRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListEncryptionKeyInfo(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSource-", LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EncryptionKeyInfoSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func EncryptionKeyInfoSummaryToMap(obj oci_log_analytics.EncryptionKeyInfoSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	result["key_source"] = string(obj.KeySource)

	result["key_type"] = string(obj.KeyType)

	return result
}
