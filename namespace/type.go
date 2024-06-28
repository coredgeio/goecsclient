package namespace

type CreateNamespaceReq struct {
	Namespace                string   `json:"namespace,omitempty"`
	DefaultObjectProject     string   `json:"default_object_project,omitempty"`
	DefaultDataServicesVpool string   `json:"default_data_services_vpool,omitempty"`
	AllowedVpoolsList        []string `json:"allowed_vpools_list,omitempty"`
	DisallowedVpoolsList     []string `json:"disallowed_vpools_list,omitempty"`
	NamespaceAdmins          string   `json:"namespace_admins,omitempty"`
	UserMapping              []struct {
		Domain     string `json:"domain,omitempty"`
		Attributes []struct {
			Attribute []struct {
				Key   string   `json:"key,omitempty"`
				Value []string `json:"value,omitempty"`
			} `json:"attribute,omitempty"`
		} `json:"attributes,omitempty"`
		Groups []string `json:"groups,omitempty"`
	} `json:"user_mapping,omitempty"`
	IsEncryptionEnabled          bool   `json:"is_encryption_enabled,omitempty"`
	DefaultBucketBlockSize       int64  `json:"default_bucket_block_size,omitempty"`
	ExternalGroupAdmins          string `json:"external_group_admins,omitempty"`
	IsStaleAllowed               bool   `json:"is_stale_allowed,omitempty"`
	ComplianceEnabled            bool   `json:"compliance_enabled,omitempty"`
	DefaultAuditDeleteExpiration int64  `json:"default_audit_delete_expiration,omitempty"`
	RootUserPassword             string `json:"root_user_password,omitempty"`
}

type CreateNamespaceResp struct {
	DefaultDataServicesVpool string   `json:"default_data_services_vpool,omitempty"`
	AllowedVpoolsList        []string `json:"allowed_vpools_list,omitempty"`
	DisallowedVpoolsList     []string `json:"disallowed_vpools_list,omitempty"`
	NamespaceAdmins          string   `json:"namespace_admins,omitempty"`
	IsEncryptionEnabled      string   `json:"is_encryption_enabled,omitempty"`
	UserMapping              []struct {
		Domain     string `json:"domain,omitempty"`
		Attributes []struct {
			Attribute []struct {
				Key   string   `json:"key,omitempty"`
				Value []string `json:"value,omitempty"`
			} `json:"attribute,omitempty"`
		} `json:"attributes,omitempty"`
		Groups []string `json:"groups,omitempty"`
	} `json:"user_mapping,omitempty"`
	DefaultBucketBlockSize int64  `json:"default_bucket_block_size,omitempty"`
	ExternalGroupAdmins    string `json:"external_group_admins,omitempty"`
	IsStaleAllowed         bool   `json:"is_stale_allowed,omitempty"`
	IsComplianceEnabled    bool   `json:"is_compliance_enabled,omitempty"`
	NotificationSize       int64  `json:"notification_size,omitempty"`
	BlockSize              int64  `json:"block_size,omitempty"`
	RetentionClass     []struct {
		Name   string `json:"name,omitempty"`
		Period int64  `json:"period,omitempty"`
	} `json:"retention_class,omitempty"`
	RootUserName     string `json:"root_user_name,omitempty"`
	RootUserPassword string `json:"root_user_password,omitempty"`
	Name             string `json:"name,omitempty"`
	ID               string `json:"id,omitempty"`
	Link             struct {
		Rel  string `json:"rel,omitempty"`
		Href string `json:"href,omitempty"`
	} `json:"link,omitempty"`
	CreationTime int64 `json:"creation_time,omitempty"`
	Inactive     bool      `json:"inactive,omitempty"`
	Global       bool      `json:"global,omitempty"`
	Remote       bool      `json:"remote,omitempty"`
	Vdc          struct {
		ID   string `json:"id,omitempty"`
		Link struct {
			Rel  string `json:"rel,omitempty"`
			Href string `json:"href,omitempty"`
		} `json:"link,omitempty"`
	} `json:"vdc,omitempty"`
	Internal bool `json:"internal,omitempty"`
}

type UpdateNamespaceReq struct {
	DefaultDataServicesVpool              string   `json:"default_data_services_vpool,omitempty"`
	VpoolsAddedToAllowedVpoolsList        []string `json:"vpools_added_to_allowed_vpools_list,omitempty"`
	VpoolsAddedToDisallowedVpoolsList     []string `json:"vpools_added_to_disallowed_vpools_list,omitempty"`
	VpoolsRemovedFromAllowedVpoolsList    []string `json:"vpools_removed_from_allowed_vpools_list,omitempty"`
	VpoolsRemovedFromDisallowedVpoolsList []string `json:"vpools_removed_from_disallowed_vpools_list,omitempty"`
	NamespaceAdmins                       string   `json:"namespace_admins,omitempty"`
	UserMapping                           []struct {
		Domain     string `json:"domain,omitempty"`
		Attributes []struct {
			Attribute []struct {
				Key   string   `json:"key,omitempty"`
				Value []string `json:"value,omitempty"`
			} `json:"attribute,omitempty"`
		} `json:"attributes,omitempty"`
		Groups []string `json:"groups,omitempty"`
	} `json:"user_mapping,omitempty"`
	DefaultBucketBlockSize       int64  `json:"default_bucket_block_size,omitempty"`
	ExternalGroupAdmins          string `json:"external_group_admins,omitempty"`
	IsEncryptionEnabled          bool   `json:"is_encryption_enabled,omitempty"`
	IsStaleAllowed               bool   `json:"is_stale_allowed,omitempty"`
	DefaultAuditDeleteExpiration int64  `json:"default_audit_delete_expiration,omitempty"`
	CurrentRootUserPassword      string `json:"current_root_user_password,omitempty"`
	NewRootUserPassword          string `json:"new_root_user_password,omitempty"`
}

type SetNamespaceQuotaReq struct {
	BlockSize int64 `json:"blockSize,omitempty"`
	NotificationSize int64 `json:"notificationSize,omitempty"`
}