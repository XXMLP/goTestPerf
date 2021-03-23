package client

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/feiyuw/dubbo-go/common/logger"
	"github.com/feiyuw/dubbo-go/config"
	"github.com/feiyuw/dubbo-go/protocol/dubbo"
)

const (
	clientYamlContent = `# dubbo client yaml configure file

# client
request_timeout : "100ms"
# connect timeout
connect_timeout : "100ms"
check: true

protocol_conf:
  dubbo:
    reconnect_interval: 0
    connection_number: 2
    heartbeat_period: "5s"
    session_timeout: "20s"
    fail_fast_timeout: "5s"
    pool_size: 64
    pool_ttl: 600
    getty_session_param:
      compress_encoding: false
      tcp_no_delay: true
      tcp_keep_alive: true
      keep_alive_period: "120s"
      tcp_r_buf_size: 262144
      tcp_w_buf_size: 65536
      pkg_rq_size: 1024
      pkg_wq_size: 512
      tcp_read_timeout: "1s"
      tcp_write_timeout: "5s"
      wait_timeout: "1s"
      max_msg_len: 10240
      session_name: "client"`

	logYamlContent = `level: "warn"
development: false
disableCaller: true
disableStacktrace: true
sampling:
encoding: "console"

# encoder
encoderConfig:
  messageKey: "message"
  levelKey: "level"
  timeKey: "time"
  nameKey: "logger"
  callerKey: "caller"
  stacktraceKey: "stacktrace"
  lineEnding: ""
  levelEncoder: "capitalColor"
  timeEncoder: "iso8601"
  durationEncoder: "seconds"
  callerEncoder: "short"
  nameEncoder: ""

outputPaths:
  - "stderr"
errorOutputPaths:
  - "stderr"
initialFields:`
)

func init() {
	generateTempConfFiles()
	if err := logger.InitLog(); err != nil {
		log.Fatal(err)
	}
	config.InitConsumer()
	dubbo.InitClient()
}

// NewDubboClient return a new *dubbo.Client from pool
func NewDubboClient() *dubbo.Client {
	return dubbo.NewClient()
}

func generateTempConfFiles() {
	tmpClientYml, err := ioutil.TempFile("", "client")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := tmpClientYml.WriteString(clientYamlContent); err != nil {
		log.Fatal(err)
	}
	if err := tmpClientYml.Close(); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("CONF_CONSUMER_FILE_PATH", tmpClientYml.Name()); err != nil {
		log.Fatal(err)
	}

	tmpLogYml, err := ioutil.TempFile("", "log")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := tmpLogYml.WriteString(logYamlContent); err != nil {
		log.Fatal(err)
	}
	if err := tmpLogYml.Close(); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("APP_LOG_CONF_FILE", tmpLogYml.Name()); err != nil {
		log.Fatal(err)
	}
}
