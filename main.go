package main

import (
	"log"
	"os"
	"runtime"

	cli "github.com/urfave/cli/v2"
)

func main() {
	concurrencyN := runtime.NumCPU()
	app := &cli.App{
		Name:  "downloader",
		Usage: "File concurrency downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "`URL` to download",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output `filename`",
			},
			&cli.IntFlag{
				Name:    "concurrency",
				Aliases: []string{"n"},
				Value:   concurrencyN,
				Usage:   "Concurrency `number`",
			},
		},
		Action: func(c *cli.Context) error {
			strURL := c.String("url")
			filename := c.String("output")
			concurrency := c.Int("concurrency")
			return NewDownloader(concurrency).Download(strURL, filename)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// func test() error {
// 	resp, err := http.Head("https://studygolang.com/dl/golang/go1.16.5.src.tar.gz")
// 	if err != nil {
// 		fmt.Printf("not support - %v", err)
// 		return err
// 	}
// 	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
// 		fmt.Println("support")
// 		return nil
// 	}
// 	fmt.Println(resp.Header.Get("Accept-Ranges"))
// 	return nil
// }
