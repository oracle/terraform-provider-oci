package visual_builder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_visual_builder "github.com/oracle/oci-go-sdk/v56/visualbuilder"
)

func VisualBuilderVbInstanceApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Read: readVisualBuilderVbInstanceApplication,
		Schema: map[string]*schema.Schema{
			"vb_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_open_id": {
				Type:      schema.TypeString,
				Optional:  true,
				StateFunc: utils.GetMd5Hash,
				Sensitive: true,
			},
			"application_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							// Elem:     GetDataSourceItemSchema(VisualBuilderVbInstanceApplicationsDataSource()),
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type:     schema.TypeString,
										Required: true,
									},
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version": {
										Type:     schema.TypeString,
										Required: true,
									},
									"project_id": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readVisualBuilderVbInstanceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.ReadResource(sync)
}

type VisualBuilderVbInstanceApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_visual_builder.VbInstanceClient
	Res    *oci_visual_builder.RequestSummarizedApplicationsResponse
}

func (s *VisualBuilderVbInstanceApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VisualBuilderVbInstanceApplicationsDataSourceCrud) Get() error {
	request := oci_visual_builder.RequestSummarizedApplicationsRequest{}

	if vbInstanceId, ok := s.D.GetOkExists("vb_instance_id"); ok {
		tmp := vbInstanceId.(string)
		request.VbInstanceId = &tmp
	}

	if idcsOpenId, ok := s.D.GetOkExists("idcs_open_id"); ok {
		tmp := idcsOpenId.(string)
		request.IdcsOpenId = &tmp
	} else {
		tmp := ""
		request.IdcsOpenId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "visual_builder")

	response, err := s.Client.RequestSummarizedApplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VisualBuilderVbInstanceApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("VisualBuilderVbInstanceApplicationsDataSource-", VisualBuilderVbInstanceApplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	vbInstanceApplication := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, VbApplicationSummaryToMap(item))
	}
	vbInstanceApplication["items"] = items

	resources = append(resources, vbInstanceApplication)
	if err := s.D.Set("application_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func VbApplicationSummaryToMap(obj oci_visual_builder.ApplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["state"] = string(obj.State)
	result["id"] = *obj.Id
	result["version"] = *obj.Version
	result["project_id"] = *obj.ProjectId

	return result
}
