package main

import (
	"log"

	"github.com/kubaj/kubeauth/providers"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	gcloud := providers.GCloudProvider{}
	accs, err := gcloud.ReadAccounts()
	if err != nil {
		log.Fatalln(err)
	}

	if len(accs) == 0 {
		log.Fatalln("no accounts found. Run `gcloud auth login`")
	}

	if len(accs) > 1 {
		account := ""
		survey.AskOne(
			&survey.Select{
				Message: "Choose an account:",
				Options: accs,
			},
			&account,
			nil,
		)
		gcloud.SelectAccount(account)
	}
}
