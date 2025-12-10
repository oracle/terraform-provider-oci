// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsLogAnalyticsEntityAssociationsRemoveResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsEntityAssociationsRemove,
		Read:     readLogAnalyticsLogAnalyticsEntityAssociationsRemove,
		Delete:   deleteLogAnalyticsLogAnalyticsEntityAssociationsRemove,
		Schema: map[string]*schema.Schema{
			// Required
			"association_entities": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"log_analytics_entity_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createLogAnalyticsLogAnalyticsEntityAssociationsRemove(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsEntityAssociationsRemove(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsLogAnalyticsEntityAssociationsRemove(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceCrud) ID() string {
	return s.D.Get("log_analytics_entity_id").(string)
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceCrud) Create() error {
	request := oci_log_analytics.RemoveEntityAssociationsRequest{}

	if associationEntities, ok := s.D.GetOkExists("association_entities"); ok {
		interfaces := associationEntities.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("association_entities") {
			request.AssociationEntities = tmp
		}
	}

	if logAnalyticsEntityId, ok := s.D.GetOkExists("log_analytics_entity_id"); ok {
		tmp := logAnalyticsEntityId.(string)
		request.LogAnalyticsEntityId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.RemoveEntityAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceCrud) SetData() error {
	return nil
}
