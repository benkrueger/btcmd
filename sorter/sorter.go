package sorter


import(
	"fmt"
	"path/filepath"
	"io/ioutil"
	"strings"
	"os"
	"../torrent"
)
const (
	Tracker int = iota
	Alpha
	Content
)
var inpath string
var outpath string

var sortmode int
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
func SetMode(m int){
	sortmode = m
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
func OpenTfile(path string) ([]byte,bool){
	f, err := os.Open(path)
	if err != nil {
		return nil,false
	}
	defer f.Close()
	b,err := ioutil.ReadAll(f)
	if err != nil {
		return nil,false
	}
	return b,true
}

func SortByTracker() bool {
	return true
}
//Sorting by tracker is all I can think of how to do....
func Sort(mode *string, inpath string, outpath string, dryrun bool) bool{
	func Sort(mode *string, inpath string, outpath string) bool{
		switch *mode {
		case "tracker":
		
		}
	return true
}
	return true
}

func PrintDryRun() {
	FindTorrentFiles()

	for _,f := range unsorted_files {
		buf,ok := OpenTfile(f)
		if(!ok){
			fmt.Println("Could not open torrent file :",f)
		}
		var t torrent.Torrent
		t.SetFilepath(f)
		t.UnmarshalTfBytes(buf)
		//fmt.Println(t.GetTracker())
		sorted_torrents[t.GetTracker()] = append(sorted_torrents[t.GetTracker()],t.GetFilePath())
	}
	for k,v := range sorted_torrents {
		fmt.Println(k,"===========> ",len(v)," torrents")
	}
}

func PruneDeadTorrents() {
	fmt.Println("Pruning dead Torrents")
}

