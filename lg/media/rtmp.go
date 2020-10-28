package media

import (
	"context"
	. "lg/eventbus"
	"log"
	"os"
	"os/exec"
)

func init() {

	var ctx context.Context
	var cancel context.CancelFunc

	var cmd *exec.Cmd

	/*(

	cmd := exec.Command("ffmpeg",
	"-f", "avfoundation",
	"-i", "1",
	"-vcodec", "libx264",
	"-preset", "ultrafast",
	"-acodec", "libfaac",
	"-f", "flv", "rtmp://video.nissanchina.cn/mec/456")
	*/

	GlobalBus.SubscribeAsync("rdstart", func(id string) {

		log.Println(cancel)
		ctx, cancel = context.WithCancel(context.Background())

		//app+id
		cmd = exec.CommandContext(ctx, "ffmpeg",
			"-f", "avfoundation",
			"-i", "1",
			"-vcodec", "libx264",
			"-s", "800x480",
			"-preset", "ultrafast",
			"-acodec", "libfaac",
			"-f", "flv", "rtmp://video.nissanchina.cn/mec/"+id)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		log.Println("asyncrdstart", id)
	}, false)

	GlobalBus.SubscribeAsync("rdstop", func(id string) {
		log.Println("asyncrdstop", id)

		log.Println("退出程序中...", cmd.Process.Pid)
		log.Println(cancel)

		cancel()

	}, false)

}
