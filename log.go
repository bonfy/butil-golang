package butil

import (
	"io"
	"log"
	"os"
)

// InitLogFileWithPrefix : setting log path
func InitLogFileWithPrefix(f *os.File, prefix string) {
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetPrefix(prefix)

	// UTC Time
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	// Local Time
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// InitLogFile : setting log path
func InitLogFile(f *os.File) {
	prefix := "[GO-LOG]"
	InitLogFileWithPrefix(f, prefix)
}
