package globalvar

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var OciResources map[string]*schema.Resource
var OciDatasources map[string]*schema.Resource

var OciFrameworkDataSources []func() datasource.DataSource
var OciFrameworkResources []func() resource.Resource
