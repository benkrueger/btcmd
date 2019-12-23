package main

import(
	"fmt"
	"os"
	"flag"
	"./sorter"
)
/*
	tsort [-i path] [-o path] [-d] [-a] [-s tracker|size|name]
	-i - input
	-o = output
	-d = dry run
	-dd = deduplicate
	-a = archive (gzip)
	-s = sort criteria. Options are [tracker|alpha|content]
	-g = torrent grep.
	-p = find and prune dead Torrents
*/
func main() {

	input_dir_ptr := flag.String("i","./","input path")
	output_dir_ptr := flag.String("o","./","output path")
	dry_run := flag.Bool("d",false,"Dry Run option")
	deduplicate_mode := flag.Bool("dd",false,"Deduplicate mode")
	archive_mode := flag.Bool("a",false,"Archive mode")
	sort_criteria := flag.String("s","","Sort criteria. Options are [tracker|alpha|content]")
	prune_mode := flag.Bool("p",false,"find and prune dead Torrents")
	flag.Parse()

	dir,_ := exists(*input_dir_ptr)
	if !dir {
		fmt.Println("Specify valid input dir.")
		os.Exit(1)
	}
	dir,_ = exists(*output_dir_ptr)
	if !dir {
		fmt.Println("Specify valid output dir.")
		os.Exit(1)
	}

	fmt.Println("Input path: ",*input_dir_ptr)
	fmt.Println("Output path: ",*output_dir_ptr)
	fmt.Println("Dry run: ",*dry_run)
	fmt.Println("Deduplicate: ",*deduplicate_mode)
	fmt.Println("Archive: ",*archive_mode)
	fmt.Println("Sort criteria: ",*sort_criteria)
	fmt.Println("Prune mode: ",*prune_mode)
	sorter.SetInpath(*input_dir_ptr)
	sorter.SetOutpath(*output_dir_ptr)


	if *dry_run {
		sorter.PrintDryRun()
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