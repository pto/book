package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	filename := os.Args[1]
	extension := filepath.Ext(filename)
	if !(extension == ".m3u") && !(extension == ".pls") {
		printUsage()
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else if extension == ".m3u" {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	} else {
		songs := readPlsPlaylist(string(rawBytes))
		writeM3uPlaylist(songs)
	}
}

func printUsage() {
	fmt.Printf("usage: %s <file.m3u|.pls>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "[playlist]") {
			continue
		}
		if strings.HasPrefix(line, "File") {
			rawName := strings.SplitN(line, "=", 2)[1]
			song.Filename = strings.Map(mapPlatformDirSeparator, rawName)
		} else if strings.HasPrefix(line, "Title") {
			song.Title = strings.SplitN(line, "=", 2)[1]
		} else if strings.HasPrefix(line, "Length") {
			secondsStr := strings.SplitN(line, "=", 2)[1]
			var err error
			if song.Seconds, err = strconv.Atoi(secondsStr); err != nil {
				log.Printf("invalid length: %s\n", secondsStr)
				song.Seconds = -1
			}
		} else if strings.HasPrefix(line, "NumberOfEntries") {
			entriesStr := strings.SplitN(line, "=", 2)[1]
			if entries, err := strconv.Atoi(entriesStr); err != nil {
				log.Printf("invalid NumberOfEntries: %s\n", entriesStr)
			} else if entries != len(songs) {
				log.Printf("NumberOfEntries %d doesn't match "+
					"number of entries %d\n", entries, len(songs))
			}
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n",
					title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writeM3uPlaylist(songs []Song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Println(song.Filename)
	}
}
