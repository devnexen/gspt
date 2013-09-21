
package gspt

import (
  "testing"
  "bytes"
  "strings"
  "time"

  "encoding/binary"
  "encoding/hex"

  "crypto/md5"
  "os/exec"
)


func randomMD5() string {
  str    := md5.New()
  random := new(bytes.Buffer)

  binary.Write(random, binary.LittleEndian, time.Now().UTC().UnixNano())

  str.Write(random.Bytes())

  return hex.EncodeToString(str.Sum(nil))
}


func TestSetProcTitle(t *testing.T) {
  title := randomMD5()

  SetProcTitle(title)

  out, err := exec.Command("/bin/ps", "a").Output()
  if err != nil {
    t.SkipNow()
  } else if !strings.Contains(string(out), title) {
    t.FailNow()
  }
}

