package ocvp

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpRetrievePasswordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpRetrievePassword,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sddc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sddc_password": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func readOcvpRetrievePassword(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpRetrievePasswordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpRetrievePasswordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.RetrievePasswordResponse
}

func (s *OcvpRetrievePasswordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpRetrievePasswordDataSourceCrud) Get() error {
	request := oci_ocvp.RetrievePasswordRequest{}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	if passwordType, ok := s.D.GetOkExists("type"); ok {
		if request.Type, ok = oci_ocvp.GetMappingRetrievePasswordTypeEnum(passwordType.(string)); !ok {
			return fmt.Errorf("unsupported password type: %s", passwordType.(string))
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.RetrievePassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpRetrievePasswordDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpRetrievePasswordDataSource-", OcvpRetrievePasswordDataSource(), s.D))

	s.D.Set("sddc_password", SddcPasswordToMap(s.Res.SddcPassword))

	return nil
}

func SddcPasswordToMap(obj oci_ocvp.SddcPassword) map[string]interface{} {
	result := map[string]interface{}{}

	result["passwordType"] = obj.PasswordType
	result["value"] = *obj.Value

	return result
}
