package aquasec

import (
	"fmt"
	"github.com/aquasecurity/terraform-provider-aquasec/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func resourceKubernetesAssurancePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesAssurancePolicyCreate,
		Read:   resourceKubernetesAssurancePolicyRead,
		Update: resourceKubernetesAssurancePolicyUpdate,
		Delete: resourceKubernetesAssurancePolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			/*
				"assurance_type": {
					Type:        schema.TypeString,
					Description: "What type of assurance policy is described.",
					Optional:    true,
					Computed:    true,
				},
			*/
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"author": {
				Type:        schema.TypeString,
				Description: "Name of user account that created the policy.",
				Computed:    true,
			},
			"registry": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cvss_severity_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if the cvss severity is scanned.",
				Optional:    true,
			},
			"cvss_severity": {
				Type:        schema.TypeString,
				Description: "Identifier of the cvss severity.",
				Optional:    true,
			},
			"cvss_severity_exclude_no_fix": {
				Type:        schema.TypeBool,
				Description: "Indicates that policy should ignore cvss cases that do not have a known fix.",
				Optional:    true,
			},
			"custom_severity_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"maximum_score_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if exceeding the maximum score is scanned.",
				Optional:    true,
			},
			"maximum_score": {
				Type:        schema.TypeFloat,
				Description: "Value of allowed maximum score.",
				Optional:    true,
			},
			"control_exclude_no_fix": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"custom_checks_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if scanning should include custom checks.",
				Optional:    true,
			},
			"scap_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if scanning should include scap.",
				Optional:    true,
			},
			"cves_black_list_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if cves blacklist is relevant.",
				Optional:    true,
			},
			"packages_black_list_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if packages blacklist is relevant.",
				Optional:    true,
			},
			"packages_white_list_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if packages whitelist is relevant.",
				Optional:    true,
			},
			"only_none_root_users": {
				Type:        schema.TypeBool,
				Description: "Indicates if raise a warning for images that should only be run as root.",
				Optional:    true,
			},
			"trusted_base_images_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if list of trusted base images is relevant.",
				Optional:    true,
			},
			"scan_sensitive_data": {
				Type:        schema.TypeBool,
				Description: "Indicates if scan should include sensitive data in the image.",
				Optional:    true,
			},
			"audit_on_failure": {
				Type:        schema.TypeBool,
				Description: "Indicates if auditing for failures.",
				Optional:    true,
				Default:     true,
			},
			"block_failed": {
				Type:        schema.TypeBool,
				Description: "Indicates if failed images are blocked.",
				Optional:    true,
				Default:     true,
			},
			"disallow_malware": {
				Type:        schema.TypeBool,
				Description: "Indicates if malware should block the image.",
				Optional:    true,
			},
			"monitored_malware_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exceptional_monitored_malware_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"blacklisted_licenses_enabled": {
				Type:        schema.TypeBool,
				Description: "Lndicates if license blacklist is relevant.",
				Optional:    true,
			},
			"blacklisted_licenses": {
				Type:        schema.TypeList,
				Description: "List of blacklisted licenses.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"whitelisted_licenses_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if license blacklist is relevant.",
				Optional:    true,
			},
			"whitelisted_licenses": {
				Type:        schema.TypeList,
				Description: "List of whitelisted licenses.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_checks": {
				Type:        schema.TypeList,
				Description: "List of Custom user scripts for checks.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"script_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_modified": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"engine": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"snippet": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_only": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"author": {
							Type:        schema.TypeString,
							Description: "Name of user account that created the policy.",
							Optional:    true,
						},
					},
				},
			},
			"scap_files": {
				Type:        schema.TypeList,
				Description: "List of SCAP user scripts for checks.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"variables": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"registries": {
				Type:        schema.TypeList,
				Description: "List of registries.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:        schema.TypeList,
				Description: "List of labels.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:        schema.TypeList,
				Description: "List of images.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cves_black_list": {
				Type:        schema.TypeList,
				Description: "List of cves blacklisted items.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"packages_black_list": {
				Type:        schema.TypeSet,
				Description: "List of backlisted images.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"epoch": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version_range": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"release": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arch": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"license": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"packages_white_list": {
				Type:        schema.TypeSet,
				Description: "List of whitelisted images.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"epoch": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version_range": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"release": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arch": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"license": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"allowed_images": {
				Type:        schema.TypeList,
				Description: "List of explicitly allowed images.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"trusted_base_images": {
				Type:        schema.TypeSet,
				Description: "List of trusted images.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"registry": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"imagename": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"read_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"force_microenforcer": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"docker_cis_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"kube_cis_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enforce_excessive_permissions": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"function_integrity_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dta_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cves_white_list_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if cves whitelist is relevant.",
				Optional:    true,
			},
			"cves_white_list": {
				Type:        schema.TypeList,
				Description: "List of cves whitelisted licenses",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"kubernetes_controls_names": {
				Type:        schema.TypeList,
				Description: "List of kubernetes control names and available kubernetes controls are: 'Access to host IPC namespace', 'Access to host PID', 'Access to host network', 'Access to host ports', 'All container images must start with a GCR domain', 'All container images must start with an ECR domain', 'All container images must start with the *.azurecr.io domain', 'CPU not limited', 'CPU requests not specified', 'Can elevate its own privileges', 'ConfigMap with secrets', 'ConfigMap with sensitive content', 'Container images from public registries used', 'Default capabilitiessome containers do not drop all', 'Default capabilitiessome containers do not drop any', 'Delete pod logs', 'Exec into Pods', 'Image tag :latest used', 'Manage EKS IAM Auth ConfigMap', 'Manage Kubernetes RBAC resources', 'Manage Kubernetes networking', 'Manage Kubernetes workloads and pods', 'Manage all resources', 'Manage all resources at the namespace', 'Manage configmaps', 'Manage namespace secrets', 'Manage secrets', 'Manage webhookconfigurations', 'Manages /etc/hosts', 'Memory not limited', 'Memory requests not specified', 'Non-core volume types used.', 'Non-default /proc masks set', 'Privileged', 'Root file system is not read-only', 'Runs as root user', 'Runs with GID <= 10000', 'Runs with UID <= 10000', 'Runs with a root primary or supplementary GID', 'Runtime/Default AppArmor profile not set', 'Runtime/Default Seccomp profile not set', 'SELinux custom options set', 'SYS_ADMIN capability added', 'Seccomp policies disabled', 'Service with External IP', 'Specific capabilities added', 'Unsafe sysctl options set', 'User with admin access', 'Workloads in the default namespace', 'hostPath volume mounted with docker.sock', 'hostPath volumes mounted'",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"blacklist_permissions_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if blacklist permissions is relevant.",
				Optional:    true,
			},
			"blacklist_permissions": {
				Type:        schema.TypeList,
				Description: "List of function's forbidden permissions.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enforce": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enforce_after_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ignore_recently_published_vln": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_recently_published_vln_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ignore_risk_resources_enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if risk resources are ignored.",
				Optional:    true,
			},
			"ignored_risk_resources": {
				Type:        schema.TypeList,
				Description: "List of ignored risk resources.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_scopes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"auto_scan_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_scan_configured": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_scan_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iteration_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iteration": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"week_days": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"required_labels_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"required_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"forbidden_labels_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"forbidden_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Name of the container image.",
				Optional:    true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dta_severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_nfs_mounts": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"malware_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"partial_results_image_fail": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"maximum_score_exclude_no_fix": {
				Type:        schema.TypeBool,
				Description: "Indicates that policy should ignore cases that do not have a known fix.",
				Optional:    true,
			},
		},
	}
}

func resourceKubernetesAssurancePolicyCreate(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)
	name := d.Get("name").(string)
	assurance_type := "kubernetes"

	iap := expandAssurancePolicy(d, assurance_type)
	err := ac.CreateAssurancePolicy(iap, assurance_type)

	if err != nil {
		return err
	}
	d.SetId(name)
	return resourceKubernetesAssurancePolicyRead(d, m)

}

func resourceKubernetesAssurancePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)
	name := d.Get("name").(string)
	assurance_type := "kubernetes"

	if d.HasChanges("description", "registry", "cvss_severity_enabled", "cvss_severity", "cvss_severity_exclude_no_fix", "custom_severity_enabled", "maximum_score_enabled", "maximum_score", "control_exclude_no_fix", "custom_checks_enabled",
		"scap_enabled", "cves_black_list_enabled", "packages_black_list_enabled", "packages_white_list_enabled", "only_none_root_users", "trusted_base_images_enabled", "scan_sensitive_data", "audit_on_failure", "block_failed",
		"disallow_malware", "monitored_malware_paths", "exceptional_monitored_malware_paths", "blacklisted_licenses_enabled", "blacklisted_licenses", "whitelisted_licenses_enabled", "whitelisted_licenses", "custom_checks", "scap_files", "scope",
		"registries", "labels", "images", "cves_black_list", "packages_black_list", "packages_white_list", "allowed_images", "trusted_base_images", "read_only", "force_microenforcer", "docker_cis_enabled", "kube_cis_enabled", "enforce_excessive_permissions",
		"function_integrity_enabled", "dta_enabled", "cves_white_list", "kubernetes_controls_names", "cves_white_list_enabled", "blacklist_permissions_enabled", "blacklist_permissions", "enabled", "enforce", "enforce_after_days", "ignore_recently_published_vln", "ignore_recently_published_vln_period",
		"ignore_risk_resources_enabled", "ignored_risk_resources", "application_scopes", "auto_scan_enabled", "auto_scan_configured", "auto_scan_time", "required_labels_enabled", "required_labels", "forbidden_labels_enabled", "forbidden_labels", "domain_name",
		"domain", "description", "dta_severity", "scan_nfs_mounts", "malware_action", "partial_results_image_fail", "maximum_score_exclude_no_fix") {
		iap := expandAssurancePolicy(d, assurance_type)
		err := ac.UpdateAssurancePolicy(iap, assurance_type)
		if err == nil {
			err1 := resourceKubernetesAssurancePolicyRead(d, m)
			if err1 == nil {
				d.SetId(name)
			} else {
				return err1
			}
		} else {
			return err
		}
	}
	return nil
}

func resourceKubernetesAssurancePolicyRead(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)
	assurance_type := "kubernetes"

	iap, err := ac.GetAssurancePolicy(d.Id(), assurance_type)

	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", err), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	//d.Set("assurance_type", iap.AssuranceType)
	d.Set("name", iap.Name)
	d.Set("description", iap.Description)
	d.Set("author", iap.Author)
	d.Set("application_scopes", iap.ApplicationScopes)
	d.Set("registry", iap.Registry)
	d.Set("cvss_severity_enabled", iap.CvssSeverityEnabled)
	d.Set("cvss_severity", iap.CvssSeverity)
	d.Set("cvss_severity_exclude_no_fix", iap.CvssSeverityExcludeNoFix)
	d.Set("custom_severity_enabled", iap.CustomSeverityEnabled)
	d.Set("maximum_score_enabled", iap.MaximumScoreEnabled)
	d.Set("maximum_score", iap.MaximumScore)
	d.Set("control_exclude_no_fix", iap.ControlExcludeNoFix)
	d.Set("custom_checks_enabled", iap.CustomChecksEnabled)
	d.Set("scap_enabled", iap.ScapEnabled)
	d.Set("cves_black_list_enabled", iap.CvesBlackListEnabled)
	d.Set("packages_black_list_enabled", iap.PackagesBlackListEnabled)
	d.Set("packages_white_list_enabled", iap.PackagesWhiteListEnabled)
	d.Set("only_none_root_users", iap.OnlyNoneRootUsers)
	d.Set("trusted_base_images_enabled", iap.TrustedBaseImagesEnabled)
	d.Set("scan_sensitive_data", iap.ScanSensitiveData)
	d.Set("audit_on_failure", iap.AuditOnFailure)
	d.Set("block_failed", iap.BlockFailed)
	d.Set("disallow_malware", iap.DisallowMalware)
	d.Set("monitored_malware_paths", iap.MonitoredMalwarePaths)
	d.Set("exceptional_monitored_malware_paths", iap.ExceptionalMonitoredMalwarePaths)
	d.Set("blacklisted_licenses_enabled", iap.BlacklistedLicensesEnabled)
	d.Set("blacklisted_licenses", iap.BlacklistedLicenses)
	d.Set("whitelisted_licenses_enabled", iap.WhitelistedLicensesEnabled)
	d.Set("whitelisted_licenses", iap.WhitelistedLicenses)
	d.Set("custom_checks", flattenCustomChecks(iap.CustomChecks))
	d.Set("scap_files", iap.ScapFiles)
	d.Set("scope", flatteniapscope(iap.Scope))
	d.Set("registries", iap.Registries)
	d.Set("labels", iap.Labels)
	d.Set("images", iap.Images)
	d.Set("cves_black_list", iap.CvesBlackList)
	d.Set("packages_black_list", flattenPackages(iap.PackagesBlackList))
	d.Set("packages_white_list", flattenPackages(iap.PackagesWhiteList))
	d.Set("allowed_images", iap.AllowedImages)
	d.Set("trusted_base_images", flattenTrustedBaseImages(iap.TrustedBaseImages))
	d.Set("read_only", iap.ReadOnly)
	d.Set("force_microenforcer", iap.ForceMicroenforcer)
	d.Set("docker_cis_enabled", iap.DockerCisEnabled)
	d.Set("kube_cis_enabled", iap.KubeCisEnabled)
	d.Set("enforce_excessive_permissions", iap.EnforceExcessivePermissions)
	d.Set("function_integrity_enabled", iap.FunctionIntegrityEnabled)
	d.Set("dta_enabled", iap.DtaEnabled)
	d.Set("cves_white_list_enabled", iap.CvesWhiteListEnabled)
	d.Set("cves_white_list", iap.CvesWhiteList)
	d.Set("kubernetes_controls_names", iap.KubenetesControlsNames)
	d.Set("blacklist_permissions_enabled", iap.BlacklistPermissionsEnabled)
	d.Set("blacklist_permissions", iap.BlacklistPermissions)
	d.Set("enabled", iap.Enabled)
	d.Set("enforce", iap.Enforce)
	d.Set("enforce_after_days", iap.EnforceAfterDays)
	d.Set("ignore_recently_published_vln", iap.IgnoreRecentlyPublishedVln)
	d.Set("ignore_recently_published_vln_period", iap.IgnoreRecentlyPublishedVlnPeriod)
	d.Set("ignore_risk_resources_enabled", iap.IgnoreRiskResourcesEnabled)
	d.Set("ignored_risk_resources", iap.IgnoredRiskResources)
	d.Set("application_scopes", iap.ApplicationScopes)
	d.Set("auto_scan_enabled", iap.AutoScanEnabled)
	d.Set("auto_scan_configured", iap.AutoScanConfigured)
	d.Set("auto_scan_time", flattenAutoScanTime(iap.AutoScanTime))
	d.Set("required_labels_enabled", iap.RequiredLabelsEnabled)
	d.Set("required_labels", flattenLabels(iap.RequiredLabels))
	d.Set("forbidden_labels_enabled", iap.ForbiddenLabelsEnabled)
	d.Set("forbidden_labels", flattenLabels(iap.ForbiddenLabels))
	d.Set("domain_name", iap.DomainName)
	d.Set("domain", iap.Domain)
	d.Set("dta_severity", iap.DtaSeverity)
	d.Set("scan_nfs_mounts", iap.ScanNfsMounts)
	d.Set("malware_action", iap.MalwareAction)
	d.Set("partial_results_image_fail", iap.PartialResultsImageFail)
	d.Set("maximum_score_exclude_no_fix", iap.MaximumScoreExcludeNoFix)

	return nil
}

func resourceKubernetesAssurancePolicyDelete(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)
	name := d.Get("name").(string)
	assurance_type := "kubernetes"
	err := ac.DeleteAssurancePolicy(name, assurance_type)

	if err == nil {
		d.SetId("")
	} else {
		return err
	}
	return nil
}
