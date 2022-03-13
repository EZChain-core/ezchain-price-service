package middleware

import "time"

type BalancePayload struct {
	Security float64 `json:"security",omitempty,omitnil`
	CDN float64 `json:"cdn",omitempty,omitnil`
	SytemAdmin float64 `json:"sysadmin",omitempty,omitnil`
	VPN float64 `json:"vpn",omitempty,omitnil`
	ContainerRegistry float64 `json:"container_registry",omitempty,omitnil`
	CloudStorage float64 `json:"cloud_storage",omitempty,omitnil`
	LoadBalancer float64 `json:"load_balancer",omitempty,omitnil`
	Primary float64 `json:"primary",omitempty,omitnil`
	MailBox float64 `json:"mail_inbox",omitempty,omitnil`
	Pentest float64 `json:"pentest",omitempty,omitnil`
	CallCenter float64 `json:"call_center",omitempty,omitnil`
}

type ServicePayload struct {
	CanonicalName string `json:"canonical_name",omitempty,omitnil`
	Name string `json:"name",omitempty,omitnil`
	ServiceUrl string `json:"service_url",omitempty,omitnil`
}

type IAMPayload struct {
	Token string  `json:"token",omitempty,omitnil`
	Expire string `json:"expire",omitempty,omitnil`
	TenantID string `json:"tenant_id",omitempty,omitnil`
	TenantName string `json:"tenant_name",omitempty,omitnil`
	VerifiedEmail bool `json:"verified_email",omitempty,omitnil`
	VerifiedPhone bool `json:"verified_phone",omitempty,omitnil`
	VerifiedPayment bool `json:"verified_payment",omitempty,omitnil`
	Role *string `json:"role",omitempty,omitnil`
}

type DomainPayload struct {
	Domain string `json:"domain",omitempty,omitnil`
	Active bool `json:"active",omitempty,omitnil`
	Verify bool `json:"verify",omitempty,omitnil`
}

type UserPayload struct {
	Service *string `json:"service",omitempty,omitnil`
	UrlType *string `json:"url_type",omitempty,omitnil`
	Origin *string `json:"origin",omitempty,omitnil`
	ClientType string  `json:"client_type",omitempty,omitnil`
	IP string `json:"ip",omitempty,omitnil`
	UserAgent string `json:"user-agent",omitempty,omitnil`
	BillingBalance int64  `json:"billing_balance",omitempty,omitnil`
	Balances BalancePayload `json:"balances",omitempty,omitnil`
	PaymentMethod string `json:"payment_method",omitempty,omitnil`

	BillingAccID string `json:"billing_acc_id",omitempty,omitnil`
	Debit bool `json:"debit",omitempty,omitnil`
	Email string `json:"email",omitempty,omitnil`
	Phone string `json:"phone",omitempty,omitnil`
	FullName string `json:"fullname",omitempty,omitnil`
	VerifiedEmail bool `json:"verified_email",omitempty,omitnil`
	VerifiedPhone bool `json:"verified_phone",omitempty,omitnil`
	LoginAlert bool `json:"login_alert",omitempty,omitnil`
	VerifiedPayment bool `json:"verified_payment",omitempty,omitnil`
	LastRegion string `json:"last_region",omitempty,omitnil`
	Type string `json:"type",omitempty,omitnil`
	OTP bool `json:"otp",omitempty,omitnil`
	Services []ServicePayload `json:"services",omitempty,omitnil`
	Whitelist []string `json:"whitelist",omitempty,omitnil`
	Token string `json:"token",omitempty,omitnil`
	Expires *time.Time `json:"expires",omitempty,omitnil`

	TenantID string `json:"tenant_id",omitempty,omitnil`
	TenantName string `json:"tenant_name",omitempty,omitnil`
	KsUserID string `json:"ks_user_id",omitempty,omitnil`
	IAM IAMPayload `json:"iam",omitempty,omitnil`
	Domains []DomainPayload `json:"domains",omitempty,omitnil`

}

