package op

import (
	"fmt"
	"time"
)

type Item struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Version   int       `json:"version"`
	Vault     Vault     `json:"vault"`
	Category  Category  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Sections []Section `json:"sections"`
	Fields   []Field   `json:"fields"`
	URLs     []URL     `json:"urls"`
}

func (i Item) Field(key string) Field {
	var res Field
	for _, field := range i.Fields {
		if key == field.ID || key == field.Label {
			res = field
		}
	}
	return res
}

type Category string

const (
	CategoryAPICredential        = "API Credential"
	CategoryBankAccount          = "Bank Account"
	CategoryCreditCard           = "Credit Card"
	CategoryDatabase             = "Database"
	CategoryDocument             = "Document"
	CategoryDriverLicense        = "Driver License"
	CategoryEmailAccount         = "Email Account"
	CategoryIdentity             = "Identity"
	CategoryLogin                = "Login"
	CategoryMedicalRecord        = "Medical Record"
	CategoryMembership           = "Membership"
	CategoryOutdoorLicense       = "Outdoor License"
	CategoryPassport             = "Passport"
	CategoryPassword             = "Password"
	CategoryRewardProgram        = "Reward Program"
	CategorySecureNote           = "Secure Note"
	CategoryServer               = "Server"
	CategorySocialSecurityNumber = "Social Security Number"
	CategorySoftwareLicense      = "Software License"
	CategorySSHKey               = "SSH Key"
	CategoryWirelessRouter       = "Wireless Router"
)

func (c *Category) UnmarshalText(text []byte) error {
	switch string(text) {
	case "API_CREDENTIAL":
		*c = CategoryAPICredential
	case "BANK_ACCOUNT":
		*c = CategoryBankAccount
	case "CREDIT_CARD":
		*c = CategoryCreditCard
	case "DATABASE":
		*c = CategoryDatabase
	case "DOCUMENT":
		*c = CategoryDocument
	case "DRIVER_LICENSE":
		*c = CategoryDriverLicense
	case "EMAIL_ACCOUNT":
		*c = CategoryEmailAccount
	case "IDENTITY":
		*c = CategoryIdentity
	case "LOGIN":
		*c = CategoryLogin
	case "MEDICAL_RECORD":
		*c = CategoryMedicalRecord
	case "MEMBERSHIP":
		*c = CategoryMembership
	case "OUTDOOR_LICENSE":
		*c = CategoryOutdoorLicense
	case "PASSPORT":
		*c = CategoryPassport
	case "PASSWORD":
		*c = CategoryPassword
	case "REWARD_PROGRAM":
		*c = CategoryRewardProgram
	case "SECURE_NOTE":
		*c = CategorySecureNote
	case "SERVER":
		*c = CategoryServer
	case "SOCIAL_SECURITY_NUMBER":
		*c = CategorySocialSecurityNumber
	case "SOFTWARE_LICENSE":
		*c = CategorySoftwareLicense
	case "SSH_KEY":
		*c = CategorySSHKey
	case "WIRELESS_ROUTER":
		*c = CategoryWirelessRouter
	default:
		return fmt.Errorf("unrecognized category %q", string(text))
	}
	return nil
}

type Field struct {
	ID        string       `json:"id"`
	Section   Section      `json:"section"`
	Type      FieldType    `json:"type"`
	Purpose   FieldPurpose `json:"purpose"`
	Label     string       `json:"label"`
	Value     string       `json:"value"`
	TOTP      string       `json:"totp"`
	Reference string       `json:"reference"`
}

type FieldType string

const (
	FieldTypeConcealed = "CONCEALED"
	FieldTypeEmail     = "EMAIL"
	FieldTypeOTP       = "OTP"
	FieldTypeString    = "STRING"
)

type FieldPurpose string

const (
	FieldPurposeNotes    = "NOTES"
	FieldPurposePassword = "PASSWORD"
	FieldPurposeUsername = "USERNAME"
)

type Section struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type URL struct {
	Label   string `json:"label"`
	Primary bool   `json:"primary"`
	HRef    string `json:"href"`
}

func GetItem(name string, flags ...Flag) (Item, error) {
	cmd := []string{"item", "get", name}
	var item Item
	err := run(cmd, flags, &item)
	return item, err
}

func ListItem(flags ...Flag) ([]Item, error) {
	cmd := []string{"item", "list"}
	items := []Item{}
	err := run(cmd, flags, &items)
	return items, err
}
