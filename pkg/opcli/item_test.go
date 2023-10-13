package opcli

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItemField(t *testing.T) {
	tests := []struct {
		name   string
		item   Item
		key    string
		result *Field
	}{
		{
			name: "MatchID",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "passwd",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
				},
			},
			key: "username",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "uname",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
		{
			name: "MatchLabel",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "passwd",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
				},
			},
			key: "uname",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "uname",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
		{
			name: "MatchFirst",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID: "06CDE696F7B54212BE47E7F99CF674F0",
						Section: Section{
							ID: "Section_FFD16B98A713452695E49DA0EB32BFD0",
						},
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "wat",
						Reference: "op://Personal/Foo/Section_FFD16B98A713452695E49DA0EB32BFD0/username",
					},
				},
			},
			key: "username",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "uname",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, test.item.Field(test.key))
		})
	}
}

func TestItemFindFields(t *testing.T) {
	tests := []struct {
		name   string
		item   Item
		matchF func(f Field) bool
		result []Field
	}{
		{
			name: "MatchSingle",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Label == "password" },
			result: []Field{
				{
					ID:        "password",
					Type:      FieldTypeConcealed,
					Label:     "password",
					Value:     "bar",
					Reference: "op://Personal/Foo/password",
				},
			},
		},
		{
			name: "MatchMultiple",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Type == FieldTypeString },
			result: []Field{
				{
					ID:        "username",
					Type:      FieldTypeString,
					Label:     "username",
					Value:     "foo",
					Reference: "op://Personal/Foo/username",
				},
				{
					ID:        "email",
					Type:      FieldTypeString,
					Label:     "email",
					Value:     "foo@example.com",
					Reference: "op://Personal/Foo/email",
				},
			},
		},
		{
			name: "MatchNone",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Type == FieldTypeDate },
			result: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, test.item.FindFields(test.matchF))
		})
	}
}

func TestCategoryUnmarshalText(t *testing.T) {
	tests := []struct {
		text   string
		result Category
	}{
		{
			text:   "API_CREDENTIAL",
			result: CategoryAPICredential,
		},
		{
			text:   "BANK_ACCOUNT",
			result: CategoryBankAccount,
		},
		{
			text:   "CREDIT_CARD",
			result: CategoryCreditCard},
		{
			text:   "DATABASE",
			result: CategoryDatabase},
		{
			text:   "DOCUMENT",
			result: CategoryDocument,
		},
		{
			text:   "DRIVER_LICENSE",
			result: CategoryDriverLicense,
		},
		{
			text:   "EMAIL_ACCOUNT",
			result: CategoryEmailAccount,
		},
		{
			text:   "IDENTITY",
			result: CategoryIdentity,
		},
		{
			text:   "LOGIN",
			result: CategoryLogin,
		},
		{
			text:   "MEDICAL_RECORD",
			result: CategoryMedicalRecord,
		},
		{
			text:   "MEMBERSHIP",
			result: CategoryMembership,
		},
		{
			text:   "OUTDOOR_LICENSE",
			result: CategoryOutdoorLicense,
		},
		{
			text:   "PASSPORT",
			result: CategoryPassport,
		},
		{
			text:   "PASSWORD",
			result: CategoryPassword,
		},
		{
			text:   "REWARD_PROGRAM",
			result: CategoryRewardProgram,
		},
		{
			text:   "SECURE_NOTE",
			result: CategorySecureNote,
		},
		{
			text:   "SERVER",
			result: CategoryServer,
		},
		{
			text:   "SOCIAL_SECURITY_NUMBER",
			result: CategorySocialSecurityNumber,
		},
		{
			text:   "SOFTWARE_LICENSE",
			result: CategorySoftwareLicense,
		},
		{
			text:   "SSH_KEY",
			result: CategorySSHKey,
		},
		{
			text:   "WIRELESS_ROUTER",
			result: CategoryWirelessRouter,
		},
		{
			text:   "invalid",
			result: Category(""),
		},
	}
	for _, test := range tests {
		var c Category
		err := c.UnmarshalText([]byte(test.text))
		assert.NoError(t, err)
		assert.Equal(t, test.result, c)
	}
}

func TestFieldTypeUnmarshalText(t *testing.T) {
	tests := []struct {
		text   string
		result FieldType
	}{
		{
			text:   "ADDRESS",
			result: FieldTypeAddress,
		},
		{
			text:   "CONCEALED",
			result: FieldTypeConcealed,
		},
		{
			text:   "CREDIT_CARD_NUMBER",
			result: FieldTypeCreditCardNumber,
		},
		{
			text:   "CREDIT_CARD_TYPE",
			result: FieldTypeCreditCardType,
		},
		{
			text:   "DATE",
			result: FieldTypeDate,
		},
		{
			text:   "EMAIL",
			result: FieldTypeEmail,
		},
		{
			text:   "FILE",
			result: FieldTypeFile,
		},
		{
			text:   "GENDER",
			result: FieldTypeGender,
		},
		{
			text:   "MENU",
			result: FieldTypeMenu,
		},
		{
			text:   "MONTH_YEAR",
			result: FieldTypeMonthYear,
		},
		{
			text:   "OTP",
			result: FieldTypeOTP,
		},
		{
			text:   "PHONE",
			result: FieldTypePhone,
		},
		{
			text:   "REFERENCE",
			result: FieldTypeReference,
		},
		{
			text:   "SSHKEY",
			result: FieldTypeSSHKey,
		},
		{
			text:   "STRING",
			result: FieldTypeString,
		},
		{
			text:   "UNKNOWN",
			result: FieldTypeUnknown,
		},
		{
			text:   "URL",
			result: FieldTypeURL,
		},
		{
			text:   "invalid",
			result: FieldTypeUnknown,
		},
	}
	for _, test := range tests {
		var ft FieldType
		err := ft.UnmarshalText([]byte(test.text))
		assert.NoError(t, err)
		assert.Equal(t, test.result, ft)
	}
}

func TestListItems(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp []Item
		err  string
	}{
		{
			name: "All",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems()
			},
			resp: []Item{
				{
					ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
					Title:    "Foo",
					Favorite: true,
					Version:  1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:              CategoryLogin,
					LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					AdditionalInformation: "foo@example.com",
					URLs: []URL{
						{
							Label:   "website",
							Primary: true,
							HRef:    "https://example.com",
						},
					},
				},
				{
					ID:      "utfq63h5szb3jeembehuoioc4f",
					Title:   "Evil Corp.",
					Version: 1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7g",
						Name: "Bar",
					},
					Category:     CategoryMembership,
					LastEditedBy: "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Archived",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems(IncludeArchive())
			},
			resp: []Item{
				{
					ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
					Title:    "Foo",
					Favorite: true,
					Version:  1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:              CategoryLogin,
					LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					AdditionalInformation: "foo@example.com",
					URLs: []URL{
						{
							Label:   "website",
							Primary: true,
							HRef:    "https://example.com",
						},
					},
				},
				{
					ID:      "utfq63h5szb3jeembehuoioc4f",
					Title:   "Evil Corp.",
					Version: 1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7g",
						Name: "Bar",
					},
					Category:     CategoryMembership,
					LastEditedBy: "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
				{
					ID:      "ypaehfyfzrc5xmvosxywp5rwr4",
					Title:   "Some note",
					Version: 1,
					State:   ItemStateArchived,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:     CategorySecureNote,
					LastEditedBy: "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Categories",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems(WithCategories(CategoryAPICredential, CategoryLogin))
			},
			resp: []Item{
				{
					ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
					Title:    "Foo",
					Favorite: true,
					Version:  1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:              CategoryLogin,
					LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					AdditionalInformation: "foo@example.com",
					URLs: []URL{
						{
							Label:   "website",
							Primary: true,
							HRef:    "https://example.com",
						},
					},
				},
			},
		},
		{
			name: "Favorite",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems(WithFavorite())
			},
			resp: []Item{
				{
					ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
					Title:    "Foo",
					Favorite: true,
					Version:  1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:              CategoryLogin,
					LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					AdditionalInformation: "foo@example.com",
					URLs: []URL{
						{
							Label:   "website",
							Primary: true,
							HRef:    "https://example.com",
						},
					},
				},
			},
		},
		{
			name: "Tags",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems(WithTags("foo", "bar baz"))
			},
			resp: []Item{
				{
					ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
					Title:    "Foo",
					Favorite: true,
					Tags:     []string{"bar baz"},
					Version:  1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7f",
						Name: "Personal",
					},
					Category:              CategoryLogin,
					LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					AdditionalInformation: "foo@example.com",
					URLs: []URL{
						{
							Label:   "website",
							Primary: true,
							HRef:    "https://example.com",
						},
					},
				},
			},
		},
		{
			name: "Vault",
			call: func(cli *CLI) (any, error) {
				return cli.ListItems(WithVault("Bar"))
			},
			resp: []Item{
				{
					ID:      "utfq63h5szb3jeembehuoioc4f",
					Title:   "Evil Corp.",
					Version: 1,
					Vault: Vault{
						ID:   "ynghx4vwntpezvhqyeglcp7v7g",
						Name: "Bar",
					},
					Category:     CategoryMembership,
					LastEditedBy: "F7GSLUVENFGZVF2HVACL3IAS7F",
					CreatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestCreateItem(t *testing.T) {
	// FIXME: implement
}

func TestGetItem(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *Item
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) { return cli.GetItem("Foo") },
			resp: &Item{
				ID:       "ijfuujah5bfehb4rnx6rkxzpv5",
				Title:    "Foo",
				Favorite: true,
				Tags:     []string{"bar baz"},
				Version:  1,
				Vault: Vault{
					ID:   "ynghx4vwntpezvhqyeglcp7v7f",
					Name: "Personal",
				},
				Category:              CategoryLogin,
				LastEditedBy:          "F7GSLUVENFGZVF2HVACL3IAS7F",
				CreatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				UpdatedAt:             time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				AdditionalInformation: "foo@example.com",
				URLs: []URL{
					{
						Label:   "website",
						Primary: true,
						HRef:    "https://example.com",
					},
				},
				Sections: []Section{
					{
						ID: "add more",
					},
				},
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "passwd",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID: "06CDE696F7B54212BE47E7F99CF674F0",
						Section: Section{
							ID: "add more",
						},
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "wat",
						Reference: "op://Personal/Foo/add more/username",
					},
				},
				Files: []File{
					{
						ID:          "toairadal5cpbfbqs72qzrokxb",
						Name:        "Test file",
						Size:        13,
						ContentPath: "/v1/vaults/ynghx4vwntpezvhqyeglcp7v7f/items/ijfuujah5bfehb4rnx6rkxzpv5/files/toairadal5cpbfbqs72qzrokxb/content",
						Section: Section{
							ID: "add more",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestDeleteItem(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) error
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) error { return cli.DeleteItem("Foo") },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			err := test.call(cli)
			if test.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestArchiveItem(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) error
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) error { return cli.ArchiveItem("Foo") },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			err := test.call(cli)
			if test.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}
