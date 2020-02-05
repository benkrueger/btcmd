package torrent
import bencode "github.com/IncSW/go-bencode"
import(
	"fmt"
	"regexp"
	"crypto/sha1"
	"strings"
	"net/http"
	"os"
	"io/ioutil"
)
var ANNOUCEREG = regexp.MustCompile("(udp|https?)://([^/^:]+)")

type Torrent struct {
	filepath string
	announce string
	scrape string
	tracker string
	name string
	length int64
	contents []string
	infohash string
	private bool
}
func (t *Torrent) SetFilepath(s string){
	t.filepath = s
}
func (t *Torrent) UnmarshalTfBytes(b []byte) {
	data,err := bencode.Unmarshal(b)
	if err == nil {
		
		buf_m_s_i := data.(map[string]interface{})
		announce_url_str := string(buf_m_s_i["announce"].([]byte))
		info_dict := buf_m_s_i["info"].(map[string]interface{})
		info_bytes,_ := bencode.Marshal(info_dict)
		
		t.infohash = fmt.Sprintf("%x", sha1.Sum([]byte(info_bytes)))
		t.name = string(info_dict["name"].([]byte))

		if(info_dict["private"] != nil){
			if(info_dict["private"].(int64) == 1){
				t.private = true
			}
		}else{
			t.private = false
		}
		files_list :=info_dict["files"]
		t.announce = announce_url_str
		t.SetTracker()

		if files_list != nil {
			
			contained_files_list := files_list.([]interface{})
			for _,f := range contained_files_list {
				
				f_dict := f.(map[string]interface{})
				f_paths := f_dict["path"].([]interface{})
				f_size := f_dict["length"].(int64)
				t.length += f_size
				for _,p := range f_paths {
					path_string := string(p.([]byte))
					t.contents = append(t.contents,path_string)
				}
			}

		}else{
			t.length = info_dict["length"].(int64)
		}
	}
}
func (t *Torrent)GetTracker() string{
	return t.tracker
}
func (t *Torrent)GetFilePath() string{
	return t.filepath
}
func (t *Torrent)GetLength()  int64{
	return t.length
}
func (t *Torrent)GetName()  string{
	return t.name
}
func (t *Torrent)GetPrivate()  bool{
	return t.private
}
func (t *Torrent)GetInfohash()  string{
	return t.infohash
}
func (t *Torrent)GetAnnouce()  string{
	return t.announce
}

func (t *Torrent)SetTracker(){
	if t.announce != "" {
		t.tracker = ANNOUCEREG.FindStringSubmatch(t.announce)[2]
		t.scrape = strings.Replace(t.announce,"announce","scrape",-1)
	}	
}
//TODO: figure out how to talk to trackers
//This is actually really hard, I might never do it.
func (t *Torrent)ScrapeTorrent() {
	reqstring := fmt.Sprintf("%s?info_hash=%s?peer_id=%s",t.scrape,t.infohash,"-LT2060-")
	resp,err := http.Get(reqstring)
	if err != nil {
		fmt.Println("Could not scrape, request failed.")
		
	}else{
		typestring := "none"
		fmt.Println("Header :",resp.Header)
		if resp.Header["Content-Type"] != nil {
			typestring = resp.Header["Content-Type"][0]
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Could not decipher body")
		}
		switch typestring {
		case "text/plain":
			fmt.Println("Text response:",string(body))
			break
		case "text/html":
			fmt.Println("HTML response:",string(body))
			break
		case "none":
			fmt.Println("No response possible")
		default:
			fmt.Println("IDK response:",string(body))
		}

		defer resp.Body.Close()
	}
	
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
func (t *Torrent)PrintTorrentInfo() {
	fmt.Println("File: ",t.filepath)
	fmt.Println("Name :",t.name)
	fmt.Println("Size :",t.length)
	fmt.Println("Tracker :",t.tracker)
	fmt.Println("Infohash :",string(t.infohash))
	fmt.Println("Announce :",t.announce)
	fmt.Println("Scrape :",t.scrape)
	fmt.Println("Private :",t.private)
}