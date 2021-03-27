package main

import(
	"fmt"
	"os"
	"flag"
)
/*
	btdmp:
	Dump bittorrent files to command line.
	options:
	-i <input path>
	-j json output
	-t text output
	-p scrape tracker annouce to get list of peers
*/

func main() {
	input_file := flag.String("i","","input file")
	json := flag.Bool("j",false,"Dump torrent file as json")
	text := flag.Bool("t",false,"Dump torrent file as plaintext")
	peers := flag.Bool("p",false,"Check if torrent is alive from tr")
	flag.Parse()
	if(*input_file != "") {
		f_exists,f_err := exists(*input_file);
		if(!f_exists) {
			fmt.Println("Input file does not exist")
			os.Exit(1)
		}
		if(f_err != nil) {
			fmt.Println(f_err)
			os.Exit(1)
		}
	} 
	if(*text) {

	}
	if(*json) {
		fmt.Println("Outputing torrent contents as JSON dict")
	}
	if(*peers) {
		fmt.Println("Outputting torrent peer list")
	}
	os.Exit(0)
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
