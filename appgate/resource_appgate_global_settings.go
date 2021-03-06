package appgate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Global settings is as the name suggest global settings, and should be treated as a singleton object.
// we can never delete global_settings, but we can UPDATE and DELETE (reset) it.
// https://discuss.hashicorp.com/t/singleton-resource/9869
func resourceGlobalSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalSettingsCreate,
		Read:   resourceGlobalSettingsRead,
		Update: resourceGlobalSettingsUpdate,
		Delete: resourceGlobalSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"claims_token_expiration": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the Claims Token is valid both for administrators and clients.",
				Optional:    true,
				Computed:    true,
			},
			"entitlement_token_expiration": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the Entitlement Token is valid for clients.",
				Optional:    true,
				Computed:    true,
			},
			"administration_token_expiration": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the administration Token is valid for administrators.",
				Optional:    true,
				Computed:    true,
			},
			"vpn_certificate_expiration": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the VPN certificates is valid for clients.",
				Optional:    true,
				Computed:    true,
			},
			"login_banner_message": {
				Type:        schema.TypeString,
				Description: "The configured message will be displayed on the login UI.",
				Optional:    true,
				Computed:    true,
			},
			"message_of_the_day": {
				Type:        schema.TypeString,
				Description: "The configured message will be displayed after a successful loging.",
				Optional:    true,
				Computed:    true,
			},
			"backup_api_enabled": {
				Type:        schema.TypeBool,
				Description: "Whether the backup API is enabled or not.",
				Optional:    true,
				Computed:    true,
			},
			"has_backup_passphrase": {
				Type:        schema.TypeBool,
				Description: "Whether there is a backup passphrase set or not. Deprecated as of 5.0. Use backupApiEnabled instead.",
				Optional:    true,
				Computed:    true,
			},
			"backup_passphrase": {
				Type:        schema.TypeString,
				Description: "The passphrase to encrypt Appliance Backups when backup API is used.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"fips": {
				Type:        schema.TypeBool,
				Description: "FIPS 140-2 Compliant Tunneling.",
				Optional:    true,
				Computed:    true,
			},
			"geo_ip_updates": {
				Type:        schema.TypeBool,
				Description: "Whether the automatic GeoIp updates are enabled or not.",
				Optional:    true,
				Computed:    true,
			},
			"audit_log_persistence_mode": {
				Type:        schema.TypeString,
				Description: "Audit Log persistence mode.",
				Optional:    true,
				Computed:    true,
			},
			"app_discovery_domains": {
				Type:        schema.TypeSet,
				Description: "Domains to monitor for for App Discovery feature.",
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"collective_id": {
				Type:        schema.TypeString,
				Description: "A randomly generated ID during first installation to identify the Collective.",
				Computed:    true,
			},
		},
	}
}

func resourceGlobalSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceGlobalSettingsUpdate(d, meta)
}

func resourceGlobalSettingsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Global settings id: %+v", d.Id())
	token := meta.(*Client).Token
	api := meta.(*Client).API.GlobalSettingsApi
	ctx := context.TODO()
	request := api.GlobalSettingsGet(ctx)
	settings, _, err := request.Authorization(token).Execute()
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Failed to read Global settings, %+v", err)
	}
	d.SetId(settings.GetCollectiveId())
	d.Set("claims_token_expiration", settings.GetClaimsTokenExpiration())
	d.Set("entitlement_token_expiration", settings.GetEntitlementTokenExpiration())
	d.Set("administration_token_expiration", settings.GetAdministrationTokenExpiration())
	d.Set("vpn_certificate_expiration", settings.GetVpnCertificateExpiration())
	d.Set("login_banner_message", settings.GetLoginBannerMessage())
	d.Set("message_of_the_day", settings.GetMessageOfTheDay())
	d.Set("backup_api_enabled", settings.GetBackupApiEnabled())
	d.Set("has_backup_passphrase", settings.GetHasBackupPassphrase())
	if val, ok := d.GetOk("backup_passphrase"); ok {
		d.Set("backup_passphrase", val)
	} else {
		d.Set("backup_passphrase", settings.GetBackupPassphrase())
	}
	d.Set("fips", settings.GetFips())
	d.Set("geo_ip_updates", settings.GetGeoIpUpdates())
	d.Set("audit_log_persistence_mode", settings.GetAuditLogPersistenceMode())
	d.Set("app_discovery_domains", settings.GetAppDiscoveryDomains())
	d.Set("collective_id", settings.GetCollectiveId())

	return nil
}

func resourceGlobalSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating Global settings")
	token := meta.(*Client).Token
	api := meta.(*Client).API.GlobalSettingsApi
	ctx := context.TODO()
	request := api.GlobalSettingsGet(ctx)
	originalsettings, _, err := request.Authorization(token).Execute()
	if err != nil {
		return fmt.Errorf("Failed to read Global settings while updating, %+v", err)
	}

	if d.HasChange("claims_token_expiration") {
		originalsettings.SetClaimsTokenExpiration(float32(d.Get("claims_token_expiration").(int)))
	}
	if d.HasChange("entitlement_token_expiration") {
		originalsettings.SetEntitlementTokenExpiration(float32(d.Get("entitlement_token_expiration").(int)))
	}
	if d.HasChange("administration_token_expiration") {
		originalsettings.SetAdministrationTokenExpiration(float32(d.Get("administration_token_expiration").(int)))
	}
	if d.HasChange("vpn_certificate_expiration") {
		originalsettings.SetVpnCertificateExpiration(float32(d.Get("vpn_certificate_expiration").(int)))
	}
	if d.HasChange("login_banner_message") {
		originalsettings.SetLoginBannerMessage(d.Get("login_banner_message").(string))
	}
	if d.HasChange("message_of_the_day") {
		originalsettings.SetMessageOfTheDay(d.Get("message_of_the_day").(string))
	}
	if d.HasChange("backup_api_enabled") {
		originalsettings.SetBackupApiEnabled(d.Get("backup_api_enabled").(bool))
	}
	if d.HasChange("has_backup_passphrase") {
		originalsettings.SetHasBackupPassphrase(d.Get("has_backup_passphrase").(bool))
	}
	if d.HasChange("backup_passphrase") {
		originalsettings.SetBackupPassphrase(d.Get("backup_passphrase").(string))
	}
	if d.HasChange("fips") {
		originalsettings.SetFips(d.Get("fips").(bool))
	}
	if d.HasChange("geo_ip_updates") {
		originalsettings.SetGeoIpUpdates(d.Get("geo_ip_updates").(bool))
	}
	if d.HasChange("audit_log_persistence_mode") {
		originalsettings.SetAuditLogPersistenceMode(d.Get("audit_log_persistence_mode").(string))
	}
	if d.HasChange("app_discovery_domains") {
		_, n := d.GetChange("app_discovery_domains")
		domains, err := readArrayOfStringsFromConfig(n.(*schema.Set).List())
		if err != nil {
			return err
		}
		originalsettings.SetAppDiscoveryDomains(domains)
	}
	log.Printf("[DEBUG] Updating Global settings %+v", originalsettings)
	req := api.GlobalSettingsPut(ctx)
	_, err = req.GlobalSettings(originalsettings).Authorization(token).Execute()
	if err != nil {
		return fmt.Errorf("Could not update Global settings %+v", prettyPrintAPIError(err))
	}

	return resourceGlobalSettingsRead(d, meta)
}

func resourceGlobalSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Delete Global settings")
	token := meta.(*Client).Token
	api := meta.(*Client).API.GlobalSettingsApi

	request := api.GlobalSettingsDelete(context.TODO())

	_, err := request.Authorization(token).Execute()
	if err != nil {
		return fmt.Errorf("Could reset Global settings %+v", prettyPrintAPIError(err))
	}
	d.SetId("")
	return nil
}
