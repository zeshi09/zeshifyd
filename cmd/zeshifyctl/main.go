package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/spf13/cobra"
	"github.com/zeshi09/zeshifyd/internal/model"
)

var (
	path        = os.Getenv("HOME") + "/.cache/zeshifyd/zeshifyd.json"
	validFields = []string{
		"app_name",
		"summary",
		"body",
		"icon",
		"time",
		"timeout",
	}
)

func baseRead() *model.Notification {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var n model.Notification
	if err := json.Unmarshal(data, &n); err != nil {
		log.Fatalf("failed to decode json: %v", err)
	}

	return &n
}

func getFullJson() {
	n := baseRead()

	output, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		log.Fatalf("failed to marshal notification: %v", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", output)
}

func getField(f string) {
	n := baseRead()

	if !slices.Contains(validFields, f) {
		fmt.Fprintf(os.Stderr, "❌ unknown field %s\n", f)
		return
	}

	switch f {
	case "app_name":
		fmt.Println(n.Appname)
	case "summary":
		fmt.Println(n.Summary)
	case "body":
		fmt.Println(n.Body)
	case "icon":
		fmt.Println(n.Icon)
	case "time":
		fmt.Println(n.Time.Format("2006-01-02 15:04:05"))
	case "timeout":
		fmt.Println(n.Timeout)
	default:
		fmt.Println("❌ shouldn't happen")
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "zeshifyctl",
		Short: "CLI for zeshifyd notificaton daemon",
	}

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show latest notification",
		Run: func(cmd *cobra.Command, args []string) {
			jsonFlag, _ := cmd.Flags().GetBool("json")
			fieldFlag, _ := cmd.Flags().GetString("field")
			listFlag, _ := cmd.Flags().GetBool("list-fields")
			if jsonFlag {
				getFullJson()
			} else if fieldFlag != "" {
				getField(fieldFlag)
			} else if listFlag {
				fmt.Fprintln(os.Stdout, "List of available fields for show --field command below:")
				for _, v := range validFields {
					fmt.Fprintf(os.Stdout, "  • %s\n", v)
				}
			} else {
				n := baseRead()
				fmt.Fprintf(os.Stdout, "🔔 %s: %s\n", n.Summary, n.Body)
			}
		},
	}

	showCmd.Flags().Bool("json", false, "A show for full json")
	showCmd.Flags().String("field", "", "A show certain field")
	showCmd.Flags().Bool("list-fields", false, "Show available fields for show --field")

	rootCmd.AddCommand(showCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
