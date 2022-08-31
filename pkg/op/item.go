package op

import (
	"time"
)

func GetItem(name string, vault string) (Item, error) {
	args := []string{"item", "get", name, "--format", "json", "--iso-timestamps"}
	if vault != "" {
		args = append(args, "--vault", vault)
	}
	var item Item
	err := run(args, &item)
	return item, err
}

func ListItem(vault string) ([]Item, error) {
	args := []string{"item", "list", "--format", "json", "--iso-timestamps", "--categories", "API Credential"}
	if vault != "" {
		args = append(args, "--vault", vault)
	}
	items := []Item{}
	err := run(args, &items)
	return items, err
}

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
		if field.ID == key || field.Label == key {
			res = field
		}
	}
	return res
}

type Category string

const (
	CategoryAPICredential        = "API_CREDENTIAL"
	CategoryBankAccount          = "BANK_ACCOUNT"
	CategoryCreditCard           = "CREDIT_CARD"
	CategoryDocument             = "DOCUMENT"
	CategoryDriverLicense        = "DRIVER_LICENSE"
	CategoryIdentity             = "IDENTITY"
	CategoryLogin                = "LOGIN"
	CategoryMembership           = "MEMBERSHIP"
	CategoryPassport             = "PASSPORT"
	CategoryPassword             = "PASSWORD"
	CategorySecureNote           = "SECURE_NOTE"
	CategoryServer               = "SERVER"
	CategorySocialSecurityNumber = "SOCIAL_SECURITY_NUMBER"
	CategorySoftwareLicense      = "SOFTWARE_LICENSE"
	CategorySSHKey               = "SSH_KEY"
	CategoryWirelessRouter       = "WIRELESS_ROUTER"
)

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

type Vault struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
