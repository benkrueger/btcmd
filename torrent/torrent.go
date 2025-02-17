package torrent

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	bencode "github.com/IncSW/go-bencode"
)

var ANNOUCEREG = regexp.MustCompile("(udp|https?)://([^/^:]+)")

type Torrent struct {
	filepath    string
	announce    string
	scrape      string
	tracker     string
	name        string
	length      int64
	contents    []string
	infohashStr string
	private     bool
}

type beincodeInfo struct {
	Pieces string
}

func (t *Torrent) UnmarshalTfBytes(b []byte) error {
	data, err := bencode.Unmarshal(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bencode data: %v", err)
	}

	buf_m_s_i, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected type for bencode data")
	}

	announce, ok := buf_m_s_i["announce"].(string)
	if !ok {
		return fmt.Errorf("announce field missing or invalid")
	}
	t.announce = announce

	// Extracting `info` dictionary and calculate info hash
	infoMap, ok := buf_m_s_i["info"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("info field missing or invalid")
	}

	infoBytes, err := bencode.Marshal(infoMap)
	if err != nil {
		return fmt.Errorf("failed to marshal info dict: %v", err)
	}

	t.infohashStr = fmt.Sprintf("%x", sha1.Sum(infoBytes))

	// Set attributes
	t.name = infoMap["name"].(string)
	t.length = infoMap["length"].(int64)
	t.private = infoMap["private"].(int64) == 1 // assuming it's included
	t.SetTracker()

	// Handle files if they exist
	if files, ok := infoMap["files"]; ok {
		for _, file := range files.([]interface{}) {
			fileMap := file.(map[string]interface{})
			t.length += fileMap["length"].(int64)
			for _, pathElement := range fileMap["path"].([]interface{}) {
				t.contents = append(t.contents, pathElement.(string))
			}
		}
	} else {
		t.length = infoMap["length"].(int64)
	}

	return nil
}

func (t *Torrent) GetTracker() string {
	return t.tracker
}

func (t *Torrent) GetFilePath() string {
	return t.filepath
}

func (t *Torrent) GetLength() int64 {
	return t.length
}

func (t *Torrent) GetName() string {
	return t.name
}

func (t *Torrent) GetPrivate() bool {
	return t.private
}

func (t *Torrent) GetInfohash() string {
	return t.infohashStr
}

func (t *Torrent) GetAnnouce() string {
	return t.announce
}

func (t *Torrent) SetTracker() {
	if t.announce != "" {
		t.tracker = ANNOUCEREG.FindStringSubmatch(t.announce)[2]
		t.scrape = strings.Replace(t.announce, "announce", "scrape", -1)
	}
}
func OpenTfile(path string) ([]byte, bool) {
	f, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, false
	}
	return b, true
}

func (t *Torrent) PrintTorrentInfo() {
	fmt.Println("File: ", t.filepath)
	fmt.Println("Name :", t.name)
	fmt.Println("Size :", t.length)
	fmt.Println("Tracker :", t.tracker)
	fmt.Println("Infohash :", t.infohashStr)
	fmt.Println("Announce :", t.announce)
	fmt.Println("Scrape :", t.scrape)
	fmt.Println("Private :", t.private)
}

func LoadTorrent(filePath string) (*Torrent, error) {
	torrentData, isSuccess := OpenTfile(filePath)
	if !isSuccess {
		return nil, fmt.Errorf("failed to open file: %s", filePath)
	}

	torrent := &Torrent{}
	torrent.UnmarshalTfBytes(torrentData)
	return torrent, nil
}

func (t *Torrent) ToJSON() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (t *Torrent) OutputPeers() {
	fmt.Println("Simulated peer list for", t.name)
}
