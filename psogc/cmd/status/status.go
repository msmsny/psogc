package status

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/go-playground/validator"
	"github.com/spf13/cobra"

	"github.com/msmsny/psogc/psogc"
)

// NewStatusCommand creates cobra.Command of psgoc status.
// opts argument is to test standard output.
func NewStatusCommand(opts ...StatusCommandOption) *cobra.Command {
	opt := &StatusOption{}
	config := &psogc.CharacterConfig{}
	output := outputFromOpts(opts)
	characterEnum := psogc.NewCharacterClassEnum()
	v := validator.New()
	v.RegisterValidation("isCharacterClass", characterEnum.ValuesValidator())
	// TODO remove isMod5
	v.RegisterValidation("isMod5", func(fl validator.FieldLevel) bool {
		return fl.Field().Int()%5 == 0
	})

	cmds := &cobra.Command{
		Use:   "status",
		Short: "View character status",
		Long:  "View character status",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validate(v, characterEnum, opt); err != nil {
				return err
			}

			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("os.Getwd: %s", err)
			}
			if config, err = psogc.LoadConfig(wd); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, character := range config.Characters {
				if character.Name != opt.Name {
					continue
				}
				for _, status := range character.Statuses {
					if opt.Level == 0 {
						output.viewStatusOneLine(character.Name, status)
					} else if status.Level == opt.Level {
						output.viewStatus(character.Name, status)
					}
				}
			}

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	flags := cmds.Flags()
	flags.SortFlags = false
	flags.StringVar(&opt.Name, "name", opt.Name, fmt.Sprintf("character name: %s", strings.Join(characterEnum.OrderedValues(), ", ")))
	flags.IntVar(&opt.Level, "level", opt.Level, "character level within 1-200")

	return cmds
}

type StatusOption struct {
	Name  string `validate:"required,isCharacterClass"`
	Level int    `validate:"omitempty,gte=1,lte=200,isMod5"` // TODO remove isMod5
}

func validate(v *validator.Validate, enum *psogc.CharacterClassEnum, opt *StatusOption) error {
	if err := v.Struct(opt); err != nil {
		messages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := ""
			switch e.Field() {
			case "Name":
				message = fmt.Sprintf("invalid name: %s, the name must be in %s", e.Value(), strings.Join(enum.OrderedValues(), ", "))
			case "Level":
				// TODO remove "and mod 5 is zero"
				message = fmt.Sprintf("invalid level: %d, the level must be 1-200 and mod 5 is zero", e.Value())
			}
			if message != "" {
				messages = append(messages, message)
			}
		}
		return fmt.Errorf(strings.Join(messages, "\n"))
	}

	return nil
}

type output struct {
	writer io.Writer
}

func NewOutput() *output {
	return &output{writer: os.Stdout}
}

func (o *output) viewStatus(name string, status *psogc.Status) {
	fmt.Fprint(o.writer, heredoc.Doc(fmt.Sprintf(`
		name:         %s
		level:        %4d
		HP:           %4d
		TP:           %4d
		Attack:       %4d
		Defense:      %4d
		MindStrength: %4d
		Accuracy:     %6.1f
		Evasion:      %4d
	`,
		name,
		status.Level,
		status.HP,
		status.TP,
		status.Attack,
		status.Defense,
		status.MindStrength,
		status.Accuracy,
		status.Evasion,
	)))
}

func (o *output) viewStatusOneLine(name string, status *psogc.Status) {
	fmt.Fprintf(
		o.writer,
		"name: %s, level: %3d, HP: %4d, TP: %4d, Attack: %4d, Defense: %3d, MindStrength: %4d, Accuracy: %5.1f, Evasion: %3d\n",
		name,
		status.Level,
		status.HP,
		status.TP,
		status.Attack,
		status.Defense,
		status.MindStrength,
		status.Accuracy,
		status.Evasion,
	)
}

func outputFromOpts(options []StatusCommandOption) *output {
	opts := &statusCommandOptions{output: NewOutput()}
	for _, option := range options {
		option.apply(opts)
	}

	return opts.output
}

type StatusCommandOption interface{ apply(*statusCommandOptions) }

type statusCommandOptions struct {
	output *output
}

type outputOption struct {
	output *output
}

func (o outputOption) apply(opts *statusCommandOptions) { opts.output = o.output }

// withOutput is unexported for using only test.
func withOutput(output *output) StatusCommandOption {
	return outputOption{output: output}
}
