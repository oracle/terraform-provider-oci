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

func LogAnalyticsNamespaceLookupsAppendDataManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createLogAnalyticsNamespaceLookupsAppendDataManagement,
		Read:   readLogAnalyticsNamespaceLookupsAppendDataManagement,
		Delete: deleteLogAnalyticsNamespaceLookupsAppendDataManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"append_lookup_file": {
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

func createLogAnalyticsNamespaceLookupsAppendDataManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupsAppendDataManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceLookupsAppendDataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsNamespaceLookupsAppendDataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsNamespaceLookupsAppendDataManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.AppendLookupDataResponse
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceLookupsAppendDataManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceLookupsAppendDataManagementResource-", LogAnalyticsNamespaceLookupsAppendDataManagementResource(), s.D)
}

func (s *LogAnalyticsNamespaceLookupsAppendDataManagementResourceCrud) Create() error {
	request := oci_log_analytics.AppendLookupDataRequest{}
	var namespaceName string

	if appendLookupFile, ok := s.D.GetOkExists("append_lookup_file"); ok {
		tmp := appendLookupFile.(string)
		contents, err := ioutil.ReadFile(tmp)
		if err != nil {
			return fmt.Errorf("Error while reading the specified file: %q", err)
		}
		request.AppendLookupFileBody = ioutil.NopCloser(bytes.NewReader(contents))
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

	response, err := s.Client.AppendLookupData(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, workRequestErr := namespaceLookupWaitForWorkRequest(&namespaceName, workId, "log_analytics",
		oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
	return workRequestErr
}

func (s *LogAnalyticsNamespaceLookupsAppendDataManagementResourceCrud) SetData() error {
	return nil
}
