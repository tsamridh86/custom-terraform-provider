package tuladhar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tuladhar_virtual_machine": resourceVirtualMachine(),
		},
	}
}
