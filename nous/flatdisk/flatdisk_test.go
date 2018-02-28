package flatdisk_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/leeola/nous"
	"github.com/leeola/nous/nous/flatdisk"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	fooContent = "foo"
	fooHash    = "2DrjgbGi2vdPGAwVDt42guEdjVYK3HbJjvcsNLVA873KHiW6pC"
)

func TestStore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "flatdisk")
	if err != nil {
		t.Fatalf("failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	Convey("Given valid input", t, func() {
		c := flatdisk.Config{
			Path: tmpDir,
		}
		n, err := flatdisk.New(c)
		So(err, ShouldBeNil)

		i := nous.Information{
			Content: fooContent,
		}

		Convey("When stored", func() {
			hash, err := n.Store(i)

			Convey("It should not error", func() {
				So(err, ShouldBeNil)
			})

			Convey("It return the hash", func() {
				So(hash, ShouldEqual, fooHash)
			})
		})
	})
}
