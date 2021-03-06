package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/joeabbey/diver/pkg/ucp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	// UCP ROLES flags
	ucpAuthRolesGet.Flags().StringVar(&name, "rolename", "", "Name of the role to retrieve")
	ucpAuthRolesGet.Flags().StringVar(&id, "id", "", "ID of the role to retrieve")

	ucpAuthRolesCreate.Flags().StringVar(&name, "rolename", "", "Name of the role to create")
	ucpAuthRolesCreate.Flags().StringVar(&ruleset, "ruleset", "", "Path to a ruleset (JSON) to be used")
	ucpAuthRolesCreate.Flags().BoolVar(&admin, "service", false, "New role is a system role")

	ucpAuthRolesDelete.Flags().StringVar(&name, "id", "", "ID of the role to delete")

	// UCP ROLES
	ucpAuth.AddCommand(ucpAuthRoles)
	ucpAuthRoles.AddCommand(ucpAuthRolesList)
	ucpAuthRoles.AddCommand(ucpAuthRolesTotal)
	if !DiverRO {
		ucpAuthRoles.AddCommand(ucpAuthRolesGet)
		ucpAuthRoles.AddCommand(ucpAuthRolesCreate)
		ucpAuthRoles.AddCommand(ucpAuthRolesDelete)

	}
}

var ucpAuthRoles = &cobra.Command{
	Use:   "roles",
	Short: "Manage Docker EE Roles",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		cmd.Help()
	},
}

var ucpAuthRolesList = &cobra.Command{
	Use:   "list",
	Short: "List Docker EE Roles",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		client, err := ucp.ReadToken()
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = client.GetRoles()
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

var ucpAuthRolesGet = &cobra.Command{
	Use:   "get",
	Short: "List all rules for a particular role",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		if name == "" && id == "" {
			cmd.Help()
			log.Fatalln("No role specified to download")
		}
		client, err := ucp.ReadToken()
		if err != nil {
			log.Fatalf("%v", err)
		}
		rules, err := client.GetRoleRuleset(name, id)
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s", rules)
	},
}

var ucpAuthRolesTotal = &cobra.Command{
	Use:   "totalrole",
	Short: "returns the TOTAL ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))

		client, err := ucp.ReadToken()
		if err != nil {
			log.Fatalf("%v", err)
		}
		rules, err := client.TotalRole()
		if err != nil {
			log.Fatalf("%v", err)
		}
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, rules, "", "\t")
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s", prettyJSON.Bytes())
	},
}

var ucpAuthRolesCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a new role based upon a ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		if name == "" {
			cmd.Help()
			log.Fatalln("No role specified to download")
		}

		rulefile, err := ioutil.ReadFile(ruleset)
		if err != nil {
			log.Fatalf("%v", err)
		}

		client, err := ucp.ReadToken()
		if err != nil {
			log.Fatalf("%v", err)
		}

		err = client.CreateRole(name, name, string(rulefile), admin)
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Infof("Role [%s] created succesfully", name)
	},
}

var ucpAuthRolesDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Docker EE Organisation",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		if name == "" {
			cmd.Help()
			log.Fatalln("No Role ID specified")
		}
		client, err := ucp.ReadToken()
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = client.DeleteRole(name)
		if err != nil {
			// Fatal error if can't read the token
			log.Fatalf("%v", err)
		}
	},
}
