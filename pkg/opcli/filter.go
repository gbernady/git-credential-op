package opcli

import (
	"strings"
)

// Filter represents a filter flag passed to 1Password CLI some commands.
type Filter func() []string

// IncludeArchive expands item list to include items in the Archive.
func IncludeArchive() Filter {
	return func() []string {
		return []string{"--include-archive"}
	}
}

// WithCategories limits results to only include items matching given categories.
func WithCategories(categories ...Category) Filter {
	return func() []string {
		if len(categories) == 0 {
			return nil
		}
		var s []string
		for _, c := range categories {
			s = append(s, string(c))
		}
		return []string{"--categories", strings.Join(s, ",")}
	}
}

// WithGroup limits results to only include resources from the given group.
func WithGroup(name string) Filter {
	return func() []string {
		if name == "" {
			return nil
		}
		return []string{"--group", name}
	}
}

// WithFavorite limits results to only include favorite items.
func WithFavorite() Filter {
	return func() []string {
		return []string{"--favorite"}
	}
}

// WithTags limits results to only include items matching given tags.
func WithTags(tags ...string) Filter {
	return func() []string {
		if len(tags) == 0 {
			return nil
		}
		return []string{"--tags", strings.Join(tags, ",")}
	}
}

// WithVault limits results to only include resources from the given vault.
func WithVault(name string) Filter {
	return func() []string {
		if name == "" {
			return nil
		}
		return []string{"--vault", name}
	}
}

func applyFilters(cmd []string, filters []Filter) []string {
	for _, f := range filters {
		cmd = append(cmd, f()...)
	}
	return cmd
}
