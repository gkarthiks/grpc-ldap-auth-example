package prompts

import (
	"encoding/base64"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// PromptUser prompts the user with the provided questions
func PromptUser(label interface{}, validate func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// PromptPassword prompts the user with the provided questions and masks the input
func PromptPassword(label interface{}, validate func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
		Mask:      '⎈',
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// PromptYesNo prompts Yes/No question
func PromptYesNo(label interface{}) (promptResult string) {
	items := []string{"Yes", "No"}
	index := -1
	var err error
	for index < 0 {
		prompt := promptui.Select{
			Label: label,
			Items: items,
		}
		index, promptResult, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	}
	return
}

// ConfirmSubmitAndGenerateAuth prompts the user for confirmation
func ConfirmSubmitAndGenerateAuth() (encodedAuthBytes string) {
	confirmSubmit := PromptYesNo("Confirm submit the request?")
	if strings.Compare(confirmSubmit, "Yes") == 0 {
		userName := PromptUser("Enter your username:", UserNameValidate)
		passwd := PromptPassword("Enter your password:", PasswdLenValidate)
		encodedAuthBytes = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userName, passwd)))
	} else {
		logrus.Info("Please start over!!! Thanks!!!")
		os.Exit(0)
	}
	return
}
