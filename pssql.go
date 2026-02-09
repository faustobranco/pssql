package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"pssql/utils"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

var version = "1.1.1"

func main() {
	home, _ := os.UserHomeDir()
	defaultConfig := filepath.Join(home, ".pssql", "pssql.json")

	connectFlag := flag.String("connect", "", "Name of the server to connect directly")
	configFlag := flag.String("config", defaultConfig, "Path to the JSON configuration file")
	listFlag := flag.Bool("list", false, "List servers in the configuration file")
	helpFlag := flag.Bool("help", false, "Show help for pssql")
	versionFlag := flag.Bool("version", false, "Show version")

	flag.Usage = func() {
		fmt.Printf("pssql - PostgreSQL Connection Manager\n\nUsage: pssql [flags]\n\nFlags:\n")
		fmt.Printf("  --connect string  Direct connection by server name\n")
		fmt.Printf("  --config string   Path to JSON file (default: %s)\n", defaultConfig)
		fmt.Printf("  --list            List all configured servers\n")
		fmt.Printf("  --help            Show this help\n")
		fmt.Printf("  --version         Show current version\n")
	}

	utils.ValidateStrictFlags()
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}
	if *versionFlag {
		pterm.Info.Println("pssql version", version)
		return
	}

	jsonFile, err := os.Open(*configFlag)
	if err != nil {
		pterm.Error.Printf("Error: %v\n", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var obj_Hosts utils.Struct_Hosts
	json.Unmarshal(byteValue, &obj_Hosts)

	if os.Getenv("PSSQL_COMPLETION_MODE") == "true" {
		for _, s := range obj_Hosts.Servers {
			fmt.Println(s.Name)
		}
		return
	}

	if *listFlag {
		for _, s := range obj_Hosts.Servers {
			fmt.Printf("%s | %s\n", s.Col(s.Name, 20), s.Host)
		}
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Col .CLI 5 | cyan }} | {{ .Col .Name 35 | cyan }} | {{ .Col .Host 100 | cyan }} | {{ .Col .Database 15 | cyan }} | {{ .Col .User 15 | cyan }}",
		Inactive: "   {{ .Col .CLI 5 | white }} | {{ .Col .Name 35 | white }} | {{ .Col .Host 100 | white }} | {{ .Col .Database 15 | white }} | {{ .Col .User 15 | white }}",
	}

	prompt := promptui.Select{
		Label: "Select Server", Items: obj_Hosts.Servers, Templates: templates, Size: 10,
		Searcher: func(input string, index int) bool {
			h := obj_Hosts.Servers[index]
			return strings.Contains(strings.ToLower(h.Name+" "+h.Host), strings.ToLower(input))
		},
	}

	var obj_selected utils.Struct_Server
	directConnect := *connectFlag != ""
	aws_token := ""

	for {
		if directConnect {
			found := false
			for _, s := range obj_Hosts.Servers {
				if strings.EqualFold(s.Name, *connectFlag) {
					obj_selected = s
					found = true
					break
				}
			}
			if !found {
				pterm.Error.Printf("Server '%s' not found.\n", *connectFlag)
				return
			}
			directConnect = false
		} else {
			fmt.Print("\033[H\033[2J")
			i, _, err := prompt.Run()
			if err != nil {
				return
			}
			obj_selected = obj_Hosts.Servers[i]
		}

		if obj_selected.Name == "" || obj_selected.Host == "" {
			pterm.Error.Println("Invalid server: Name or Host missing.")
			if *connectFlag != "" {
				return
			}
			pterm.Info.Println("Press Enter to go back...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}

		if obj_selected.User == "" {
			promptUser := promptui.Prompt{
				Label: "Inform the user: ",
			}

			result, err := promptUser.Run()
			if err != nil {
				pterm.Error.Printf("%v\n", err)
				return
			}
			obj_selected.User = result
		}

		if obj_selected.Database == "" {
			promptDatabase := promptui.Prompt{
				Label: "Inform the database: ",
			}

			result, err := promptDatabase.Run()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			obj_selected.Database = result
		}

		if obj_selected.Auth == "aws-iam" {
			if obj_selected.AWSIAM.Region == "" {
				pterm.Error.Printf("The authentication is 'aws-iam', but the 'aws-iam' (Region) setting is empty in the JSON for the server. '%s'.\n", obj_selected.Name)
				pterm.Info.Printf("Press [ENTER] to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				continue
			}
			obj_ctx := context.TODO()
			cfg, err := config.LoadDefaultConfig(obj_ctx, config.WithRegion(obj_selected.AWSIAM.Region), config.WithSharedConfigProfile(obj_selected.AWSIAM.Profile))
			if err != nil {
				pterm.Error.Printf("Error loading AWS config: %v\n", err)
				return
			}

			endpoint := fmt.Sprintf("%s:%v", obj_selected.Host, obj_selected.Port)

			token, err := auth.BuildAuthToken(
				obj_ctx,
				endpoint,
				obj_selected.AWSIAM.Region,
				obj_selected.User,
				cfg.Credentials,
			)

			pterm.Info.Printf("generating AWS IAM token...")

			if err != nil {
				pterm.Error.Printf("Error generating AWS IAM token: %v\n", err)
				return
			}
			aws_token = token
		}
		cmd := exec.Command(obj_selected.CLI, "--host", obj_selected.Host, "--username", obj_selected.User, "--dbname", obj_selected.Database)
		if aws_token != "" {
			cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", aws_token))
		}
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			pterm.Error.Printf("Run error: %v. Press Enter...", err)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}
		break
	}
}
