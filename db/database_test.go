package db

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type mockDbConfig struct{}

func (m *mockDbConfig) Dialect() string { return "mockDb" }

func TestDatabase(t *testing.T) {
	Convey("Subject: Opening/closing database", t, func() {
		Convey("Given an SQLite3 configuration", func() {
			conf := &Sqlite3Config{path: filepath.Join(os.TempDir(), "dbtest.db")}

			Convey("It opens the SQLite3 database successfully", func() {
				db, err := OpenDb(conf)
				So(err, ShouldBeNil)
				So(db, ShouldNotBeNil)

				Convey("It can close the database successfully", func() {
					So(db.Close(), ShouldBeNil)
				})
			})
		})

		Convey("Given an unknown configuration", func() {
			conf := &mockDbConfig{}

			Convey("It returns an error", func() {
				db, err := OpenDb(conf)
				So(err, ShouldNotBeNil)
				So(db, ShouldBeNil)
			})
		})

	})

}
