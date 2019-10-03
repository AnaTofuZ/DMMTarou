package cmd

import (
	"fmt"
	"github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
	"github.com/spf13/cobra"
	"github.com/srvc/fail"
	"os"
)

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get contents",

		RunE: func(_ *cobra.Command, args []string) error {
			affiliateID := os.Getenv("dmm_affiID")
			apiID := os.Getenv("dmm_apiID")
			client := dmm.New(affiliateID, apiID)
			dmmapi := client.Product
			dmmapi.SetSite(api.SiteGeneral)
			dmmapi.SetService("mono")
			dmmapi.SetFloor("videoa")
			dmmapi.SetLength(10)
			result, err := dmmapi.Execute()

			if err != nil {
				return fail.Wrap(err)
			}

			fmt.Println(result)
			return nil
		},
	}
	return cmd
}
