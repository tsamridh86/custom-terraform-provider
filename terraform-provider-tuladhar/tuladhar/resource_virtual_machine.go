package tuladhar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineCreate,
		Read:   resourceVirtualMachineRead,
		Update: resourceVirtualMachineUpdate,
		Delete: resourceVirtualMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVirtualMachineCreate(d *schema.ResourceData, m interface{}) error {
	// Implement create logic
	print("Creating a new custom virtual machine")
	d.SetId("some-id")
	return resourceVirtualMachineRead(d, m)
}

func resourceVirtualMachineRead(d *schema.ResourceData, m interface{}) error {
	// Implement read logic
	print("Reading from a virtual machine")
	return nil
}

func resourceVirtualMachineUpdate(d *schema.ResourceData, m interface{}) error {
	// Implement update logic
	print("Updating a custom virtual machine")
	return resourceVirtualMachineRead(d, m)
}

func resourceVirtualMachineDelete(d *schema.ResourceData, m interface{}) error {
	// Implement delete logic
	print("Deleting a custom virtual machine")
	d.SetId("")
	return nil
}
