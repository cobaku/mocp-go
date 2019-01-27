package mocp

import (
	"os/exec"
	"strconv"
	"strings"
	"github.com/mitchellh/mapstructure"
)

type MocInfo struct {
	File        string `json:"file,omitempty"`
	Title       string `json:"title,omitempty"`
	Artist      string `json:"artist,omitempty"`
	SongTitle   string `json:"song,omitempty"`
	Album       string `json:"album,omitempty"`
	TotalTime   string `json:"total,omitempty"`
	TimeLeft    string `json:"left,omitempty"`
	TotalSec    int    `json:"totalSec,omitempty"`
	CurrentTime string `json:"current,omitempty"`
	CurrentSec  int    `json:"currentSec,omitempty"`
	Bitrate     string `json:"bitrate,omitempty"`
	AvgBitrate  string `json:"avgBitrate,omitempty"`
	Rate        string `json:"rate,omitempty"`
}

func Info() (MocInfo, error) {
	info, err := Exec("info")

	if err != nil {
		return MocInfo{}, err
	}

	entries := strings.Split(strings.TrimSpace(info), "\n")

	m := make(map[string]string)
	for _, e := range entries {
		parts := strings.Split(e, ": ")
		m[parts[0]] = parts[1]
	}

	result := MocInfo{}
	mapstructure.Decode(m, &result)
	return result, nil
}

func Exec(command string, arg ...string) (string, error) {
	args := make([]string, 0, len(arg)+1)
	args = append(args, "--"+command)
	args = append(args, arg...)
	out, err := exec.Command("mocp", args...).Output()
	if err == nil {
		return string(out), nil
	} else {
		return "", err
	}
}

func Run() {
	Exec("server")
}

func Exit() {
	Exec("exit")
}

func Pause() {
	Exec("pause")
}

func Next() {
	Exec("next")
}

func Prev() {
	Exec("previous")
}

func Play() {
	Exec("play")
}

func Stop() {
	Exec("stop")
}

func StartInDir(dir string) {
	Exec("music-dir", dir)
}

func Enqueue(files ...string) {
	Exec("enqueue", files...)
}

func ClearPlaylist() {
	Exec("clear")
}

func Append(files ...string) {
	Exec("append", files...)
}

func Seek(time int) {
	Exec("seek", strconv.Itoa(time))
}

func Config(file string) {
	Exec("config", file)
}

func Sync() {
	Exec("sync")
}

func NoSync() {
	Exec("nosync")
}

func Volume(vol int) {
	Exec("volume", strconv.Itoa(vol))
}

func Jump(val string) {
	Exec("jump", val)
}

func IsRunning() bool {
	check, _ := exec.Command("pidof", "mocp").Output()
	return len(check) > 0
}
