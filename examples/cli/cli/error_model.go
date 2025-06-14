// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/go-swagger/go-swagger/examples/cli/models"
)

// Schema cli for Error

// register flags to command
func registerModelErrorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorPropCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorPropMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorPropCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagCodeDescription := ``

	var flagCodeName string
	if cmdPrefix == "" {
		flagCodeName = "code"
	} else {
		flagCodeName = fmt.Sprintf("%v.code", cmdPrefix)
	}

	var flagCodeDefault int64

	_ = cmd.PersistentFlags().Int64(flagCodeName, flagCodeDefault, flagCodeDescription)

	return nil
}

func registerErrorPropMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagMessageDescription := `Required. `

	var flagMessageName string
	if cmdPrefix == "" {
		flagMessageName = "message"
	} else {
		flagMessageName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var flagMessageDefault string

	_ = cmd.PersistentFlags().String(flagMessageName, flagMessageDefault, flagMessageDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CodeAdded := retrieveErrorPropCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CodeAdded

	err, MessageAdded := retrieveErrorPropMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MessageAdded

	return nil, retAdded
}

func retrieveErrorPropCodeFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCodeName := fmt.Sprintf("%v.code", cmdPrefix)
	if cmd.Flags().Changed(flagCodeName) {

		var flagCodeName string
		if cmdPrefix == "" {
			flagCodeName = "code"
		} else {
			flagCodeName = fmt.Sprintf("%v.code", cmdPrefix)
		}

		flagCodeValue, err := cmd.Flags().GetInt64(flagCodeName)
		if err != nil {
			return err, false
		}
		m.Code = flagCodeValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorPropMessageFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagMessageName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(flagMessageName) {

		var flagMessageName string
		if cmdPrefix == "" {
			flagMessageName = "message"
		} else {
			flagMessageName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		flagMessageValue, err := cmd.Flags().GetString(flagMessageName)
		if err != nil {
			return err, false
		}
		m.Message = &flagMessageValue

		retAdded = true
	}

	return nil, retAdded
}
