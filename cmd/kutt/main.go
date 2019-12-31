package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/raahii/kutt-go"
	"github.com/spf13/cobra"
)

var cfgPath string

var RootCmd *cobra.Command

func init() {
	cobra.OnInitialize()

	RootCmd = rootCmd()
	RootCmd.AddCommand(
		apikeyCmd(),
		submitCmd(),
		listCmd(),
		deleteCmd(),
	)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func loadApiKey() (string, error) {
	if !exists(cfgPath) {
		return "", errors.New("api key is not registerd. run 'kutt apikey <key>'")
	}

	// read the file
	fp, err := os.Open(cfgPath)
	if err != nil {
		return "", fmt.Errorf("reading api key from file: %w", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	apiKey := scanner.Text()

	if apiKey == "" {
		return "", fmt.Errorf("api key is empty. remove %s and run 'kutt apikey' again.", cfgPath)
	}

	return apiKey, nil
}

// root command
func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kutt",
		Short: "CLI tool for Kutt.it (URL Shortener)",
	}

	cmd.PersistentFlags().StringP("apikey", "k", "", "api key for Kutt.it")

	return cmd
}

// apikey command
func apikeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apikey <key>",
		Short: "Register your api key to cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("apikey is required.")
			}
			key := args[0]

			if exists(cfgPath) {
				err := os.Remove(cfgPath)
				if err != nil {
					return fmt.Errorf("removing %s: %w", cfgPath, err)
				}
			}

			fp, err := os.Create(cfgPath)
			if err != nil {
				return fmt.Errorf("creating %s to store api key: %w", cfgPath, err)
			}
			defer fp.Close()

			fp.Write(([]byte)(key))

			return nil
		},
	}

	cmd.Flags().StringP("customurl", "c", "", "Custom ID for custom URL")
	cmd.Flags().StringP("password", "p", "", "Password for the URL")
	cmd.Flags().BoolP("reuse", "r", false, "Return last object of target if exists")
	cmd.Flags().BoolP("verbose", "v", false, "Show detailed output for the created url")

	return cmd
}

// submit command
func submitCmd() *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "submit <URL>",
		Short: "Submit a new short URL",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			apiKey, err = loadApiKey()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("target URL is required.")
			}
			target := args[0]
			flags := cmd.Flags()

			opts := []kutt.SubmitOption{}

			// customURL
			if flags.Lookup("customURL") != nil {
				customURL, err := flags.GetString("customURL")
				if err != nil {
					return err
				}
				opts = append(opts, kutt.WithCustomURL(customURL))
			}

			// password
			if flags.Lookup("password") != nil {
				password, err := flags.GetString("password")
				if err != nil {
					return err
				}
				opts = append(opts, kutt.WithPassword(password))
			}

			// reuse
			if flags.Lookup("reuse") != nil {
				reuse, err := flags.GetBool("reuse")
				if err != nil {
					return err
				}
				opts = append(opts, kutt.WithReuse(reuse))
			}

			cli := kutt.NewClient(apiKey)
			u, err := cli.Submit(target, opts...)
			if err != nil {
				return err
			}

			verbose, err := flags.GetBool("verbose")
			if err != nil {
				return err
			}

			if verbose {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(URLHeader())
				table.Append(tabulate(u))
				table.Render()
			} else {
				fmt.Println(u.ShortURL)
			}

			return nil
		},
	}

	cmd.Flags().StringP("customurl", "c", "", "Custom ID for custom URL")
	cmd.Flags().StringP("password", "p", "", "Password for the URL")
	cmd.Flags().BoolP("reuse", "r", false, "Return last object of target if exists")
	cmd.Flags().BoolP("verbose", "v", false, "Show detailed output for the created url")

	return cmd
}

// list command
func listCmd() *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List of last 5 URL objects.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			apiKey, err = loadApiKey()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := kutt.NewClient(apiKey)
			us, err := cli.List()
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(URLHeader())
			for _, u := range us {
				table.Append(tabulate(u))
			}
			table.Render()

			return nil
		},
	}

	return cmd
}

// delete command
func deleteCmd() *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "delete <ID>",
		Short: "Delete a shorted link (Give me url id or url shorted)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			apiKey, err = loadApiKey()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("target URL ID is required.")
			}
			ID := args[0]
			flags := cmd.Flags()

			opts := []kutt.DeleteOption{}

			// domain
			if flags.Lookup("domain") != nil {
				domain, err := flags.GetString("domain")
				if err != nil {
					return err
				}
				opts = append(opts, kutt.WithDomain(domain))
			}

			cli := kutt.NewClient(apiKey)
			err := cli.Delete(ID, opts...)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("domain", "d", "", "Domain for the URL")

	return cmd
}

func URLHeader() []string {
	return []string{"id", "target", "short url", "password", "visit", "created at"}
}

func tabulate(u *kutt.URL) []string {
	layout := "2006-01-02 15:04:05 MST"

	password := "false"
	if u.Password {
		password = "true"
	}

	return []string{
		u.ID,
		u.Target,
		u.ShortURL,
		password,
		strconv.Itoa(u.VisitCount),
		u.CreatedAt.In(time.Now().Location()).Format(layout),
	}
}

func main() {
	// config file path
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	cfgPath = filepath.Join(home, ".kutt")

	// execute
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}
