package main

import(
	"fmt"
	"os"
	"flag"
	"./sorter"
	"regexp"
)
/*
	tsort [-i path] [-o path] [-d] [-a] [-f]
	-i - input
	-o = output. 
	-dr = dry run
	-dd = deduplicate torrent content
	-s = sort torrents to output dir
	-a = archive (gzip)
	-f = field of torrent you wish to sort/search by. ex: -f  tracker sorts all files one dir deep by tracker  
	-g = torrent grep. -g <pattern> 
	-p = find and prune dead Torrents
	-d = run as a service. Listen to torrents coming into input folder and sort them to output.
*/
func main() {

	input_dir_ptr := flag.String("i","./","input path")
	output_dir_ptr := flag.String("o","./","output path")
	field_ptr := flag.String("f","","Field of torrent you wish to sort/search with.")
	dry_run := flag.Bool("d",false,"Dry Run option")
	deduplicate_mode := flag.Bool("dd",false,"Deduplicate mode")
	archive_mode := flag.Bool("a",false,"Archive mode")
	search_pattern := flag.String("g","","Pattern to search in torrent file")
	prune_mode := flag.Bool("p",false,"find and prune dead Torrents")
	sort_mode := flag.Bool("s",false,"Sort to output")
	flag.Parse()
	ok,err := CheckInput()
	if(!ok || err != nil){
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Input path: ",*input_dir_ptr)
	fmt.Println("Output path: ",*output_dir_ptr)
	fmt.Println("Dry run: ",*dry_run)
	fmt.Println("Deduplicate: ",*deduplicate_mode)
	fmt.Println("Archive: ",*archive_mode)
	fmt.Println("Sort criteria: ",*field_ptr)
	fmt.Println("Prune mode: ",*prune_mode)
	fmt.Println("Run as daemon:",*service_mode)
	
	if(){

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

func CheckInput(inpath string,outpath string,field string,search_pattern string)(bool,error,*Regexp){
	dir,err := exists(inpath)
	if !dir {
		fmt.Println("Specify valid input dir.")
		return false,err,nil
	}
	dir,_ = exists(outpath)
	if !dir {
		fmt.Println("Specify valid output dir.")
		return false,err,nil
	}
	search_re := regexp.MustCompilePOSIX(search_pattern)



}