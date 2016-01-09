package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "beek"
	app.Usage = `Peek at binary files.

EXAMPLES:
    # read first 16 bytes of test.bin in hex format
    beek test.bin --format x -l 16
	# read last 8 bytes of test.bin in binary format
	beek --format b test.bin`

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "length, l",
			Value: 8,
			Usage: "Number of bytes to read from file. Default: 8.",
		},
		cli.StringFlag{
			Name:  "format, s",
			Value: "d",
			Usage: `Format of output:
            'x' : hex
            'X' : hex with capitol letters
            'd' : base10
            'a' : ascii`,
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: "",
			Usage: "delimiter. Default is none.",
		},
		cli.IntFlag{
			Name:  "offset, o",
			Value: 0,
			Usage: "offset of reading in file. 0 for head, 10 for starting with 10th byte. offset of -10 reads last 10 bytes starting at -10,-9,-8,...",
		},
	}

	app.Action = func(c *cli.Context) {
		arr := readFile(c.Args().First(), int64(c.Int("offset")), c.Int("length"))
		var s string
		for i, b := range arr {
			s += fmt.Sprintf(fmt.Sprintf("%"+c.String("format"), b))
			if i < len(arr)-1 {
				s += c.String("delimiter")
			}
		}
		fmt.Println(s)
	}

	app.Run(os.Args)
}

func readFile(name string, offset int64, size int) []byte {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if size < 0 {
		log.Fatal("Size must not be out of Range")
	}
	buf := make([]byte, size)
	if offset < 0 {
		info, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		offset = info.Size() + offset
	}
	_, err = f.ReadAt(buf, offset)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return buf
}
