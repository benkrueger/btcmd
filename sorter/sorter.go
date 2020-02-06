package sorter


import(
	"fmt"
	"path/filepath"
	"strings"
	/*"../torrent"*/
)

var inpath string
var outpath string

var unsorted_files []string
var sorted_torrents map[string][]string = make(map[string][]string)
func init()(){

}
func SetInpath(ipath string){
	inpath = ipath
}
func SetOutpath(opath string){
	outpath = opath
}

func ListInpath() bool{
	FindTorrentFiles()
	for _,f := range unsorted_files{
		fmt.Println(f)
	}
	return true
}
func FindTorrentFiles() bool {
	tstring := strings.Join([]string{inpath,"*.torrent"},"")
	files,err := filepath.Glob(tstring)
	if err != nil {
		return false
	}
	unsorted_files = files
	return true
}

func Sort(mode *string, inpath string, outpath string,field string) bool{
	return false
}

