// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v58/resourcemanager"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerStackTfStateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularResourcemanagerStackTfState,
		Schema: map[string]*schema.Schema{
			"stack_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"local_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularResourcemanagerStackTfState(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackTfStateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

type ResourcemanagerStackTfStateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resourcemanager.ResourceManagerClient
	Res    *oci_resourcemanager.GetStackTfStateResponse
}

func (s *ResourcemanagerStackTfStateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourcemanagerStackTfStateDataSourceCrud) Get() error {
	request := oci_resourcemanager.GetStackTfStateRequest{}

	if stackId, ok := s.D.GetOkExists("stack_id"); ok {
		tmp := stackId.(string)
		request.StackId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.GetStackTfState(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourcemanagerStackTfStateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourcemanagerStackTfStateDataSource-", ResourcemanagerStackTfStateDataSource(), s.D))

	path, _ := s.D.GetOkExists("local_path")

	byteArr, err := ioutil.ReadAll(s.Res.Content)
	if err != nil {
		log.Printf("Unable to read Stack Tf State from response. Error: %q", err)
		return err
	}

	err = ioutil.WriteFile(path.(string), byteArr, 0644)
	if err != nil {
		log.Printf("Unable to write Stack Tf State to file. Error: %q", err)
		return err
	}

	return nil
}
