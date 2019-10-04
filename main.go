package main
import (
    "fmt"
    "os"
		"github.com/urfave/cli"
		"net/http"
		"io/ioutil"
)

func main () {

	app := cli.NewApp()
	// meta info
	app.Name = "Blog Refresh"
	app.Usage = " 'blog refresh' only"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name: "refresh",
			Usage: "blog refresh",
			Action: func(c *cli.Context) error {
				url := "https://us-central1-cloudstragetest-2f062.cloudfunctions.net/helloWorld"
				resp, _ := http.Get(url)
				defer resp.Body.Close()
				byteArray, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(byteArray))
				return nil
			},
		},
	}
	app.Run(os.Args)
}