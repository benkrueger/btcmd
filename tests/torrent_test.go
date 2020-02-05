package tests

import (
	"testing"
	"../torrent"
)

func TestSetfilepath(t *testing.T)  {
	var tor torrent.Torrent
	var testfilepath = "./testdata/tails-amd64-3.0.1.torrent"
	tor.SetFilepath(testfilepath)
	if tor.GetFilePath() != testfilepath {
		t.Errorf("Test failed")
	}
}

func TestUnmarshalTfBytes(t*testing.T) {
	var tor torrent.Torrent
	var testfilepath = "./testdata/tails-amd64-3.0.1.torrent"
	var ressize = 1209117505
	var restrack = "tracker.torrent.eu.org"
	var resname = "tails-amd64-3.0.1"
	var resannouce = "udp://tracker.torrent.eu.org:451"
	var resprivate = false
	var reshash = "3a85384aadf8a248457ace90d1dc29eb8dbd1df0"
	tor.SetFilepath(testfilepath)
	buf,err := torrent.OpenTfile(testfilepath)
	if(!err){
		t.Errorf("Opening torrent failed")
	}
	tor.UnmarshalTfBytes(buf)
	if tor.GetLength() != ressize {
		t.Errorf("Size field was found incorrect")
	}
	if tor.GetTracker() != restrack {
		t.Errorf("Tracker/Annouce url incorrect")
	}
	if tor.GetName() != "tails-amd64-3.0.1" {
		t.Errorf("Name field incorrect")
	}
	if tor.GetPrivate() {
		t.Errorf("Private field incorrect")
	}
	if tor.GetInfohash() != "3a85384aadf8a248457ace90d1dc29eb8dbd1df0" {
		t.Errorf("Infohash incorrect")
	}
}
func TestGetTracker(t *testing.T)  {
	t.Fail()
}
func TestSetTracker(t *testing.T)  {
	
}
