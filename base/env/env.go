package env

/**
  log package 独立
  业务模块应该在 env init 后进行加载

  env 只以来init
*/

/**


  O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
   O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
   O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
   O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
   O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
   O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
   O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
   O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
*/
import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var env = make(map[string]interface{})

const LOGPATH = "config.yaml"

func init() {
	log.Print("init env package")
	readEnv()
}

func GetEnv() map[string]interface{} {
	log.Print("get info")
	return env
}

func GetEnvByKey(key string) interface{} {
	log.Print("get info by key")
	return env[key]
}
func readEnv() {
	loadYaml()
	storeYaml(env)
}

func storeYaml(config map[string]interface{}) {
	bytes, err := yaml.Marshal(config)

	if err != nil {

		log.Print("marshal config  err")
	}
	err = ioutil.WriteFile(LOGPATH, bytes, 0744)

	if err != nil {
		log.Print("store yaml file err")
	}
}

func loadYaml() {
	// read bytes
	if !exist(LOGPATH) {
		_, err := os.Create(LOGPATH)
		if err != nil {
			log.Println("create config.yaml error")
		}
	}

	bytes, err := ioutil.ReadFile(LOGPATH)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, &env)
	if err != nil {
		log.Println("Unmarshal yaml file error")
	}
}
func exist(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}
