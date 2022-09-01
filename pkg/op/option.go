package op

import (
	"fmt"
	"strings"
)

type Flag func() []string

// Global flags

func WithAccount(account string) Flag {
	return func() []string {
		if account == "" {
			return nil
		}
		return []string{"--account", account}
	}
}

// Item flags

func WithCategories(categories ...Category) Flag {
	return func() []string {
		if len(categories) == 0 {
			return nil
		}
		var s []string
		for _, c := range categories {
			s = append(s, string(c))
		}
		return []string{"--categories", fmt.Sprintf(`"%s"`, strings.Join(s, ","))}
	}
}

func WithFavorite() Flag {
	return func() []string {
		return []string{"--favorite"}
	}
}

func WithIncludeArchive() Flag {
	return func() []string {
		return []string{"--include-archive"}
	}
}

func WithTags(tags ...string) Flag {
	return func() []string {
		if len(tags) == 0 {
			return nil
		}
		return []string{"--tags", fmt.Sprintf(`"%s"`, strings.Join(tags, ","))}
	}
}

func WithVault(vault string) Flag {
	return func() []string {
		if vault == "" {
			return nil
		}
		return []string{"--vault", vault}
	}
}
