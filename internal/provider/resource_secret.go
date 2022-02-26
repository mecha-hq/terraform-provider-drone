package provider

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretCreate,
		ReadContext:   resourceSecretRead,
		UpdateContext: resourceSecretUpdate,
		DeleteContext: resourceSecretDelete,

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
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSecretCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	repository := d.Get("repository").(string)
	repositorySplit := strings.Split(repository, "/")

	secret, err := client.SecretCreate(repositorySplit[0], repositorySplit[1], &drone.Secret{
		Data: d.Get("data").(string),
		Name: d.Get("name").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(secret.Name)

	resourceSecretSet(d, secret)

	return nil
}

func resourceSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	repository := d.Get("repository").(string)
	repositorySplit := strings.Split(repository, "/")

	secret, err := client.Secret(repositorySplit[0], repositorySplit[1], d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	resourceSecretSet(d, secret)

	return nil
}

func resourceSecretUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	repository := d.Get("repository").(string)
	repositorySplit := strings.Split(repository, "/")

	secretUpdate := &drone.Secret{
		Name: d.Get("name").(string),
	}

	if d.HasChange("data") {
		secretUpdate.Data = d.Get("data").(string)
	}

	secret, err := client.SecretUpdate(repositorySplit[0], repositorySplit[1], secretUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceSecretSet(d, secret)

	return nil
}

func resourceSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	repository := d.Get("repository").(string)
	repositorySplit := strings.Split(repository, "/")

	if err := client.SecretDelete(repositorySplit[0], repositorySplit[1], d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceSecretSet(d *schema.ResourceData, s *drone.Secret) {
	d.Set("data", s.Data)
	d.Set("name", s.Name)
}
