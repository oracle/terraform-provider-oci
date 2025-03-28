package tfresource

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func RegisterResource(name string, resourceSchema *schema.Resource) {
	if globalvar.OciResources == nil {
		globalvar.OciResources = make(map[string]*schema.Resource)
	}
	globalvar.OciResources[name] = resourceSchema
}

func RegisterDatasource(name string, datasourceSchema *schema.Resource) {
	if globalvar.OciDatasources == nil {
		globalvar.OciDatasources = make(map[string]*schema.Resource)
	}
	globalvar.OciDatasources[name] = datasourceSchema
}

func RegisterFrameworkDatasource(ds func() datasource.DataSource) {
	if globalvar.OciFrameworkDataSources == nil {
		globalvar.OciFrameworkDataSources = make([]func() datasource.DataSource, 0)
	}
	globalvar.OciFrameworkDataSources = append(globalvar.OciFrameworkDataSources, ds)
}

func RegisterFrameworkResource(ds func() resource.Resource) {
	if globalvar.OciFrameworkResources == nil {
		globalvar.OciFrameworkResources = make([]func() resource.Resource, 0)
	}
	globalvar.OciFrameworkResources = append(globalvar.OciFrameworkResources, ds)
}
