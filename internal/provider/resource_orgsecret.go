package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrgSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrgSecretCreate,
		ReadContext:   resourceOrgSecretRead,
		UpdateContext: resourceOrgSecretUpdate,
		DeleteContext: resourceOrgSecretDelete,

		Schema: map[string]*schema.Schema{
			"data": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOrgSecretCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	data := d.Get("data").(string)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	orgSecret, err := client.OrgSecretCreate(namespace, &drone.Secret{
		Data: data,
		Name: name,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", namespace, name))

	resourceOrgSecretSet(d, orgSecret)

	return nil
}

func resourceOrgSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	idSplit := strings.Split(d.Id(), "/")

	orgSecret, err := client.OrgSecret(idSplit[0], idSplit[1])
	if err != nil {
		return diag.FromErr(err)
	}

	resourceOrgSecretSet(d, orgSecret)

	return nil
}

func resourceOrgSecretUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	idSplit := strings.Split(d.Id(), "/")

	data := d.Get("data").(string)
	name := d.Get("name").(string)
	// namespace := d.Get("namespace").(string)

	orgSecretUpdate := &drone.Secret{
		Name: name,
	}

	if d.HasChange("data") {
		orgSecretUpdate.Data = data
	}

	orgSecret, err := client.OrgSecretUpdate(idSplit[0], orgSecretUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceOrgSecretSet(d, orgSecret)

	return nil
}

func resourceOrgSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	namespace := d.Get("namespace").(string)

	if err := client.OrgSecretDelete(namespace, d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceOrgSecretSet(d *schema.ResourceData, s *drone.Secret) {
	d.Set("data", s.Data)
	d.Set("name", s.Name)
}
