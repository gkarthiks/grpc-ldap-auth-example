// Package cmd
/*
Copyright © 2022 Karthikeyan Govindaraj
*/
package cmd

import (
	"github.com/spf13/cobra"
	"grpc-ldap-auth-example/greet/handlers"
	"grpc-ldap-auth-example/greet/prompts"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "To geet the greetings",
	Long: `To get the greetings for your name, if name is 
not provided you will be greeted as a stranger`,
	Run: func(cmd *cobra.Command, args []string) {
		yourName := prompts.PromptUser("What is your name?", nil)
		confirmationAndAuth := prompts.ConfirmSubmitAndGenerateAuth()

		handlers.SubmitRequestToGreet(yourName, confirmationAndAuth)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
