package example

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExample() *schema.Resource {
	return &schema.Resource{
		Create: resourceExampleCreate,
		Read:   resourceExampleRead,
		Update: resourceExampleUpdate,
		Delete: resourceExampleDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceExampleCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("example-id")
	return resourceExampleRead(d, m)
}

func resourceExampleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceExampleUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceExampleRead(d, m)
}

func resourceExampleDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
