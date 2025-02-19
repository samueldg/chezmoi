package cmd

import (
	"fmt"

	"github.com/twpayne/chezmoi/v2/pkg/chezmoi"
)

type pathStyle string

const (
	pathStyleAbsolute pathStyle = "absolute"
	pathStyleRelative pathStyle = "relative"
)

var (
	pathStyleStrings = []string{
		pathStyleAbsolute.String(),
		pathStyleRelative.String(),
	}

	pathStyleFlagCompletionFunc = chezmoi.FlagCompletionFunc(pathStyleStrings)
)

// Set implements github.com/spf13/pflag.Value.Set.
func (p *pathStyle) Set(s string) error {
	uniqueAbbreviations := chezmoi.UniqueAbbreviations(pathStyleStrings)
	pathStyleStr, ok := uniqueAbbreviations[s]
	if !ok {
		return fmt.Errorf("%s: unknown path style", s)
	}
	*p = pathStyle(pathStyleStr)
	return nil
}

func (p pathStyle) String() string {
	return string(p)
}

// Type implements github.com/spf13/pflag.Value.Type.
func (p pathStyle) Type() string {
	return "absolute|relative"
}
