//go:generate mockery -all -dir ../pkg/group -output ../pkg/group/mocks
//go:generate mockery -all -dir ../pkg/health -output ../pkg/health/mocks
//go:generate mockery -all -dir ../pkg/policy -output ../pkg/policy/mocks
//go:generate mockery -all -dir ../pkg/service -output ../pkg/service/mocks
//go:generate mockery -all -dir ../pkg/user -output ../pkg/user/mocks
//go:generate mockery -all -dir ../pkg/project -output ../pkg/project/mocks
//go:generate mockery -all -dir ../pkg/admin -output ../pkg/admin/mocks
//go:generate mockery -all -dir ../pkg/billing -output ../pkg/admin/mocks

package main

import (

	"fmt"
	"net/http"
	"time"
	"log"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/EZChain-core/price-service/config"

)

var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "go-realworld-clean",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

type Say struct{}


func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}


func main() {
	//repo := &Repository{}

	rootCmd.AddCommand(versionCmd)
	//cobra.OnInitialize(config.CobraInitialization)
	//config.LoggerConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}


}

func run() {

	config := config.InitConfig()

	ginServer := NewServer(
		config.ServerPort,
		DebugMode,
		config,
	)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", ginServer.Port),
		Handler:        ginServer.Router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}