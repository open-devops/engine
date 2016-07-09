package dockercompose

//go:generate counterfeiter -o ./fakeOpRunExitCodeReader.go --fake-name fakeOpRunExitCodeReader ./ opRunExitCodeReader

import (
  "os/exec"
  "strings"
  dockerEngine "github.com/docker/engine-api/client"
  "golang.org/x/net/context"
)

type opRunExitCodeReader interface {
  read(
  opBundlePath string,
  opName string,
  opNamespace string,
  ) (opExitCode int, err error)
}

func newOpRunExitCodeReader(
dockerEngine *dockerEngine.Client,
) opRunExitCodeReader {

  return &_opRunExitCodeReader{
    dockerEngine:dockerEngine,
  }

}

type _opRunExitCodeReader struct {
  dockerEngine *dockerEngine.Client
}

func (this _opRunExitCodeReader) read(
opBundlePath string,
opName string,
opRunId string,
) (
opExitCode int,
err error,
) {

  // get container id
  dockerComposePsCmd :=
  exec.Command(
    "docker-compose",
    "-p",
    opRunId,
    "ps",
    "-q",
    opName,
  )

  dockerComposePsCmd.Dir = opBundlePath

  var dockerComposePsCmdRawOutput []byte
  dockerComposePsCmdRawOutput, err = dockerComposePsCmd.Output()
  if (nil != err) {
    return
  }

  container, err := this.dockerEngine.ContainerInspect(
    context.Background(),
    strings.TrimSpace(string(dockerComposePsCmdRawOutput)),
  )
  if (nil != err) {
    return
  }
  opExitCode = container.State.ExitCode

  return

}
