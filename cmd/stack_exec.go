package cmd

import (
	//	"fmt"
	"io/ioutil"

	"github.com/mobingi/mobingi-cli/pkg/cli"
	//	"github.com/mobingilabs/mobingi-sdk-go/client"
	"github.com/mobingilabs/mobingi-sdk-go/mobingi/sesha3"
	"github.com/mobingilabs/mobingi-sdk-go/pkg/cmdline"
	d "github.com/mobingilabs/mobingi-sdk-go/pkg/debug"
	//	"github.com/mobingilabs/mobingi-sdk-go/pkg/pretty"
	//	"github.com/pkg/errors"
	//	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

func StackExecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "execute a script to your instance",
		Long: `execute a script to your instance.

Examples:

  $ ` + cmdline.Args0() + ` stack exec --stackid xxxx --target ip1,ip2,ip3,ipn --script /path/to/script`,
		Run: stackExec,
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().String("target", "", "ip1,ip2,ip3,ipn")
	cmd.Flags().String("stackid", "", "your stackid")
	cmd.Flags().String("script", "", "your script path")
	cmd.Flags().String("flag", "", "configuration flag")
	cmd.Flags().String("user", "ec2-user", "ssh username")
	return cmd
}

func stackExec(cmd *cobra.Command, args []string) {
	sess, err := clisession()
	cli.ErrorExit(err, 1)

	svc := sesha3.New(sess)
	sfile := cli.GetCliStringFlag(cmd, "script")
	scriptdata, err := ioutil.ReadFile(sfile)
	cli.ErrorExit(err, 1)
	in := &sesha3.GetExecResponseInput{
		StackId:  cli.GetCliStringFlag(cmd, "stackid"),
		Target:   cli.GetCliStringFlag(cmd, "target"),
		Script:   string(scriptdata),
		InstUser: cli.GetCliStringFlag(cmd, "user"),
		Flag:     cli.GetCliStringFlag(cmd, "flag"),
	}
	_, _, u, err := svc.GetExecResponse(in)
	cli.ErrorExit(err, 1)
	d.Info(u)

	//	cli.ErrorExit(err, 1)
	//	exitOn401(resp)
	//
	//	out := cli.GetCliStringFlag(cmd, "out")
	//	pfmt := cli.GetCliStringFlag(cmd, "fmt")
	//	switch pfmt {
	//	case "raw":
	//		fmt.Println(string(body))
	//		if out != "" {
	//			err = ioutil.WriteFile(out, body, 0644)
	//			cli.ErrorExit(err, 1)
	//		}
	//	case "json":
	//		indent := cli.GetCliIntFlag(cmd, "indent")
	//		js := pretty.JSON(string(body), indent)
	//		fmt.Println(js)
	//
	//		// write to file option
	//		if out != "" {
	//			err = ioutil.WriteFile(out, []byte(js), 0644)
	//			cli.ErrorExit(err, 1)
	//		}
	//	default:
	//		if browser {
	//			d.Info("open link with a browser (if not opened automatically):", u)
	//			_ = open.Run(u)
	//			return
	//		}
	//
	//		sshcli, err := sesha3.NewClient(&sesha3.SeshaClientInput{URL: u})
	//		cli.ErrorExit(err, 1)
	//
	//		err = sshcli.Run()
	//		if err != nil {
	//			d.Error("session return:", err)
	//		}
	//	}
}
