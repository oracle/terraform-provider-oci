// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceLookupsUpdateDataManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createLogAnalyticsNamespaceLookupsUpdateDataManagement,
		Read:   readLogAnalyticsNamespaceLookupsUpdateDataManagement,
		Delete: deleteLogAnalyticsNamespaceLookupsUpdateDataManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"update_lookup_file": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lookup_name": {
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
			"char_encoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"expect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_force": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createLogAnalyticsNamespaceLookupsUpdateDataManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupsUpdateDataManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceLookupsUpdateDataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsNamespaceLookupsUpdateDataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsNamespaceLookupsUpdateDataManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.UpdateLookupDataResponse
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceLookupsUpdateDataManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceLookupsUpdateDataManagementResource-", LogAnalyticsNamespaceLookupsUpdateDataManagementResource(), s.D)
}

func (s *LogAnalyticsNamespaceLookupsUpdateDataManagementResourceCrud) Create() error {
	request := oci_log_analytics.UpdateLookupDataRequest{}
	var namespaceName string

	if updateLookupFile, ok := s.D.GetOkExists("update_lookup_file"); ok {
		tmp := updateLookupFile.(string)
		contents, err := ioutil.ReadFile(tmp)
		if err != nil {
			return fmt.Errorf("Error while reading the specified file: %q", err)
		}
		request.UpdateLookupFileBody = ioutil.NopCloser(bytes.NewReader(contents))
	}

	if charEncoding, ok := s.D.GetOkExists("char_encoding"); ok {
		tmp := charEncoding.(string)
		request.CharEncoding = &tmp
	}

	if expect, ok := s.D.GetOkExists("expect"); ok {
		tmp := expect.(string)
		request.Expect = &tmp
	}

	if isForce, ok := s.D.GetOkExists("is_force"); ok {
		tmp := isForce.(bool)
		request.IsForce = &tmp
	}

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		request.LookupName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateLookupData(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, workRequestErr := namespaceLookupWaitForWorkRequest(&namespaceName, workId, "log_analytics",
		oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
	return workRequestErr
}

func (s *LogAnalyticsNamespaceLookupsUpdateDataManagementResourceCrud) SetData() error {
	return nil
}
