package provider

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRepoCreate,
		ReadContext:   resourceRepoRead,
		UpdateContext: resourceRepoUpdate,
		DeleteContext: resourceRepoDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRepoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	repository := d.Get("repository").(string)
	repositorySplit := strings.Split(repository, "/")
	active := d.Get("active").(bool)

	var repo *drone.Repo
	var err error

	if active {
		repo, err = client.RepoEnable(repositorySplit[0], repositorySplit[1])
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err = client.RepoDisable(repositorySplit[0], repositorySplit[1])
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository)

	resourceRepoSet(d, repo)

	return nil
}

func resourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRepoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRepoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRepoSet(d *schema.ResourceData, r *drone.Repo) {
	d.Set("active", r.Active)
}
