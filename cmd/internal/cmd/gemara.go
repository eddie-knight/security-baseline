// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gemaraproj/go-gemara/gemaraconv"
	"github.com/spf13/cobra"

	"github.com/ossf/security-baseline/pkg/baseline"
)

type gemaraOptions struct {
	outPath      string
	baselinePath string
	markdown     bool
	noTOC        bool
}

func (o *gemaraOptions) Validate() error {
	if o.baselinePath == "" {
		return errors.New("baseline path not specified")
	}
	return nil
}

func (o *gemaraOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", defaultBaselinePath, "path to directory containing the baseline YAML data",
	)
	cmd.PersistentFlags().StringVarP(
		&o.outPath, "out", "o", "", "path to output file (defaults to STDOUT)",
	)
	cmd.PersistentFlags().BoolVar(
		&o.markdown, "markdown", false,
		"write Markdown via go-gemara (gemaraconv) instead of Gemara YAML",
	)
	cmd.PersistentFlags().BoolVar(
		&o.noTOC, "no-toc", false,
		"omit table of contents (only applies with --markdown)",
	)
}

func addGemara(parentCmd *cobra.Command) {
	opts := gemaraOptions{}
	gemaraCmd := &cobra.Command{
		Use:   "gemara -o baseline.gemara.yaml",
		Short: "Export the assembled baseline as Gemara YAML",
		Long: `Assembles the baseline YAML sources into a single Gemara ControlCatalog
and writes it to a file or STDOUT. By default the output is Gemara YAML, which
can be validated externally with cue vet. With --markdown, the same catalog is
rendered as Markdown using github.com/gemaraproj/go-gemara/gemaraconv.`,
		SilenceUsage:  false,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}
			cmd.SilenceUsage = true

			loader := baseline.NewLoader()
			loader.DataPath = opts.baselinePath

			bline, err := loader.Load()
			if err != nil {
				return err
			}

			w := os.Stdout
			if opts.outPath != "" {
				f, err := os.Create(opts.outPath)
				if err != nil {
					return fmt.Errorf("creating output file: %w", err)
				}
				defer f.Close() //nolint:errcheck
				w = f
			}

			gen := baseline.NewGenerator()
			if opts.markdown {
				mdOpts := []gemaraconv.MarkdownOption{}
				if opts.noTOC {
					mdOpts = append(mdOpts, gemaraconv.WithTOC(false))
				}
				if err := gen.ExportGemaraMarkdown(bline, w, mdOpts...); err != nil {
					return err
				}
				if opts.outPath != "" {
					fmt.Fprintf(os.Stderr, "Gemara Markdown written to %s\n", opts.outPath)
				}
				return nil
			}

			if err := gen.ExportGemara(bline, w); err != nil {
				return err
			}

			if opts.outPath != "" {
				fmt.Fprintf(os.Stderr, "Gemara YAML written to %s\n", opts.outPath)
			}
			return nil
		},
	}
	opts.AddFlags(gemaraCmd)
	parentCmd.AddCommand(gemaraCmd)
}
