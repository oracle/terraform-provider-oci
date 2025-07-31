// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupArtifactResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelGroupArtifact,
		Read:     readDatascienceModelGroupArtifact,
		Delete:   deleteDatascienceModelGroupArtifact,
		Schema: map[string]*schema.Schema{
			// Required
			"model_group_artifact": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content_length": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"model_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatascienceModelGroupArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelGroupArtifact(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatascienceModelGroupArtifact(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatascienceModelGroupArtifactResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	DisableNotFoundRetries bool
}

func (s *DatascienceModelGroupArtifactResourceCrud) ID() string {
	return "nil"
}

func (s *DatascienceModelGroupArtifactResourceCrud) Create() error {
	request := oci_datascience.CreateModelGroupArtifactRequest{}

	if modelGroupArtifact, ok := s.D.GetOkExists("model_group_artifact"); ok {
		tmp := []byte(modelGroupArtifact.(string))
		request.ModelGroupArtifact = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if modelGroupId, ok := s.D.GetOkExists("model_group_id"); ok {
		tmp := modelGroupId.(string)
		request.ModelGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateModelGroupArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatascienceModelGroupArtifactResourceCrud) SetData() error {
	return nil
}
