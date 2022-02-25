package provider

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"admin": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"machine": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	user, err := client.UserCreate(&drone.User{
		Active:  d.Get("active").(bool),
		Admin:   d.Get("admin").(bool),
		Email:   d.Get("email").(string),
		Login:   d.Get("login").(string),
		Machine: d.Get("machine").(bool),
		Token:   d.Get("token").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Login)

	resourceUserSet(d, user)

	return nil
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	user, err := client.User(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	resourceUserSet(d, user)

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	userPatch := &drone.UserPatch{}

	if d.HasChange("active") {
		active := d.Get("active").(bool)
		userPatch.Active = &active
	}

	if d.HasChange("admin") {
		admin := d.Get("admin").(bool)
		userPatch.Admin = &admin
	}

	if d.HasChange("machine") {
		machine := d.Get("machine").(bool)
		userPatch.Machine = &machine
	}

	if d.HasChange("token") {
		token := d.Get("token").(string)
		userPatch.Token = &token
	}

	user, err := client.UserUpdate(d.Id(), userPatch)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceUserSet(d, user)

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(drone.Client)

	if err := client.UserDelete(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceUserSet(d *schema.ResourceData, u *drone.User) {
	d.Set("active", u.Active)
	d.Set("admin", u.Admin)
	d.Set("email", u.Email)
	d.Set("login", u.Login)
	d.Set("machine", u.Machine)
	d.Set("token", u.Token)
}
