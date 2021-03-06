package main

import (
	"github.com/spf13/cobra"
)

func main() {

	// var cmdPrint = &cobra.Command{
	// 	Use:   "print [string to print]",
	// 	Short: "Print anything to the screen",
	// 	Long: `print is for printing anything back to the screen.
	// For many years people have printed back to the screen.`,
	// 	Args: cobra.MinimumNArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("Print: " + strings.Join(args, " "))
	// 	},
	// }

	// var cmdEcho = &cobra.Command{
	// 	Use:   "echo [string to echo]",
	// 	Short: "Echo anything to the screen",
	// 	Long: `echo is for echoing anything back.
	// Echo works a lot like print, except it has a child command.`,
	// 	Args: cobra.MinimumNArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("Echo: " + strings.Join(args, " "))
	// 	},
	// }

	var cmdGenerateGinServer = &cobra.Command{
		Use:   "generate",
		Short: "Generates a go server by the defined generator",
		Long:  "Generates a new go server in the selected output location using the selected generator",
		Run: func(cmd *cobra.Command, args []string) {
			//swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile("swagger.json")
		},
	}

	var rootCmd = &cobra.Command{Use: "gogen"}
	rootCmd.AddCommand(cmdGenerateGinServer)
	rootCmd.Execute()
}
