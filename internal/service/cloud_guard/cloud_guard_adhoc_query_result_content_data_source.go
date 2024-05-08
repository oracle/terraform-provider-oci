// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardAdhocQueryResultContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCloudGuardAdhocQueryResultContent,
		Schema: map[string]*schema.Schema{
			"adhoc_query_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularCloudGuardAdhocQueryResultContent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryResultContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardAdhocQueryResultContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetAdhocQueryResultContentResponse
}

func (s *CloudGuardAdhocQueryResultContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardAdhocQueryResultContentDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetAdhocQueryResultContentRequest{}

	if adhocQueryId, ok := s.D.GetOkExists("adhoc_query_id"); ok {
		tmp := adhocQueryId.(string)
		request.AdhocQueryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetAdhocQueryResultContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardAdhocQueryResultContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardAdhocQueryResultContentDataSource-", CloudGuardAdhocQueryResultContentDataSource(), s.D))

	return nil
}
*/
