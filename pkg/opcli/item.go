package opcli

import (
	"time"
)

type Item struct {
	ID                    string    `json:"id"`
	Title                 string    `json:"title"`
	Favorite              bool      `json:"favorite"`
	Tags                  []string  `json:"tags"`
	Version               int       `json:"version"`
	State                 ItemState `json:"state"`
	Vault                 Vault     `json:"vault"`
	Category              Category  `json:"category"`
	LastEditedBy          string    `json:"last_edited_by"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	AdditionalInformation string    `json:"additional_information"`
	URLs                  []URL     `json:"urls"`

	Sections []Section `json:"sections"`
	Fields   []Field   `json:"fields"`
	Files    []File    `json:"files"`
}

func (i Item) Field(name string) *Field {
	var f *Field
	for _, field := range i.Fields {
		if field.ID == name || field.Label == name {
			f = &field
			break
		}
	}
	return f
}

func (i Item) FindFields(matching func(f Field) bool) []Field {
	var f []Field
	for _, field := range i.Fields {
		if matching(field) {
			f = append(f, field)
		}
	}
	return f
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
	}
	return nil
}

type Field struct {
	ID              string          `json:"id"`
	Section         Section         `json:"section"`
	Type            FieldType       `json:"type"`
	Purpose         FieldPurpose    `json:"purpose"`
	Label           string          `json:"label"`
	Value           string          `json:"value"`
	TOTP            string          `json:"totp"`
	Entropy         int64           `json:"entropy"`
	PasswordDetails PasswordDetails `json:"password_details"`
	Reference       string          `json:"reference"`
}

type FieldAssignment struct {
	Label   string              `json:"label"`
	Type    FieldAssignmentType `json:"type"`
	Value   string              `json:"value"`
	Purpose FieldPurpose        `json:"purpose"`
}

type FieldAssignmentType string

const (
	FieldAssignmentTypeConcealed = "concealed"
	FieldAssignmentTypeText      = "text"
	FieldAssignmentTypeEmail     = "email"
	FieldAssignmentTypeURL       = "url"
	FieldAssignmentTypeDate      = "date"
	FieldAssignmentTypeMonthYear = "monthYear"
	FieldAssignmentTypePhone     = "phone"
)

type FieldType string

const (
	FieldTypeAddress          = "address"
	FieldTypeConcealed        = "concealed"
	FieldTypeCreditCardNumber = "ccnum"
	FieldTypeCreditCardType   = "cctype"
	FieldTypeDate             = "date"
	FieldTypeEmail            = "email"
	FieldTypeFile             = "file"
	FieldTypeGender           = "gender"
	FieldTypeMenu             = "menu"
	FieldTypeMonthYear        = "monthYear"
	FieldTypeOTP              = "OTP"
	FieldTypePhone            = "phone"
	FieldTypeReference        = "reference"
	FieldTypeSSHKey           = "sshkey"
	FieldTypeString           = "string"
	FieldTypeUnknown          = ""
	FieldTypeURL              = "URL"
)

func (f *FieldType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "ADDRESS":
		*f = FieldTypeAddress
	case "CONCEALED":
		*f = FieldTypeConcealed
	case "CREDIT_CARD_NUMBER":
		*f = FieldTypeCreditCardNumber
	case "CREDIT_CARD_TYPE":
		*f = FieldTypeCreditCardType
	case "DATE":
		*f = FieldTypeDate
	case "EMAIL":
		*f = FieldTypeEmail
	case "FILE":
		*f = FieldTypeFile
	case "GENDER":
		*f = FieldTypeGender
	case "MENU":
		*f = FieldTypeMenu
	case "MONTH_YEAR":
		*f = FieldTypeMonthYear
	case "OTP":
		*f = FieldTypeOTP
	case "PHONE":
		*f = FieldTypePhone
	case "REFERENCE":
		*f = FieldTypeReference
	case "SSHKEY":
		*f = FieldTypeSSHKey
	case "STRING":
		*f = FieldTypeString
	case "URL":
		*f = FieldTypeURL
	case "UNKNOWN":
		fallthrough
	default:
		*f = FieldTypeUnknown
	}
	return nil
}

type FieldPurpose string

const (
	FieldPurposeNotes    = "NOTES"
	FieldPurposePassword = "PASSWORD"
	FieldPurposeUsername = "USERNAME"
)

type File struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Size        int64   `json:"size"`
	ContentPath string  `json:"content_path"`
	Section     Section `json:"section"`
}

type ItemState string

const (
	ItemStateArchived = "ARCHIVED"
)

type PasswordDetails struct {
	Entropy   int64            `json:"entropy"`
	Generated bool             `json:"generated"`
	Strength  PasswordStrength `json:"strength"`
}

type PasswordStrength string

const (
	PasswordStrengthTerrible  = "TERRIBLE"
	PasswordStrengthWeak      = "WEAK"
	PasswordStrengthFair      = "FAIR"
	PasswordStrengthGood      = "GOOD"
	PasswordStrengthVeryGood  = "VERY_GOOD"
	PasswordStrengthExcellent = "EXCELLENT"
	PasswordStrengthFantastic = "FANTASTIC"
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

// ListItems returns a list of all items the account has read access to.
// Excludes items in the Archive by default.
//
// Supported filters:
//
//   - WithIncludeArchive()   Include items in the Archive.
//   - WithCategories()       Only list items in these categories (comma-separated).
//   - WithFavorite()         Only list favorite items
//   - WithTags()             Only list items with these tags (comma-separated).
//   - WithVault()            Only list items in this vault.
func (c *CLI) ListItems(filters ...Filter) ([]Item, error) {
	var val []Item
	err := c.execJSON(applyFilters([]string{"item", "list"}, filters), nil, &val)
	return val, err
}

// CreateItem creates a new item and returns it with all the fileds like ID filled.
//
//	--dry-run                      Perform a dry run of the command and output a preview of the resulting item.
//	--generate-password[=recipe]   Give the item a randomly generated password.
func (c *CLI) CreateItem(item *Item) (*Item, error) {
	return nil, nil
}

// GetItem returns the details of an item specified by its name, ID, or sharing link.
//
// Supported filters:
//
//   - WithIncludeArchive()   Include items in the Archive.
//   - WithVault()            Only list items in this vault.
func (c *CLI) GetItem(name string, filters ...Filter) (*Item, error) {
	var val *Item
	err := c.execJSON(applyFilters([]string{"item", "get", name}, filters), nil, &val)
	return val, err
}

// DeleteItem permanently deletes an item specified by its name, ID, or sharing link.
func (c *CLI) DeleteItem(name string) error {
	_, err := c.execRaw([]string{"item", "delete", name}, nil)
	return err
}

// ArchiveItem archives the item specified by its name, ID, or sharing link.
func (c *CLI) ArchiveItem(name string) error {
	_, err := c.execRaw([]string{"item", "delete", name, "--archive"}, nil)
	return err
}
