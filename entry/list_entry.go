package entry

import (
	"crypto/md5"
	"encoding/hex"
	"icesos/full_path"
	"icesos/set"
	"os"
	"time"
)

type ListEntry struct {
	FullPath full_path.FullPath // file full full_path
	Set      set.Set            // own set
	Mtime    string             // time of last modification
	Ctime    string             // time of creation
	Mode     os.FileMode        // file mode
	Mime     string             // MIME type
	Md5      string             // MD5
	FileSize uint64             // file size
	Fid      string             // fid
}

func (ent *Entry) ToListEntry() *ListEntry {
	md5Str := ""
	if len(ent.Md5) == md5.Size {
		md5Str = hex.EncodeToString(ent.Md5)
	}

	return &ListEntry{
		FullPath: ent.FullPath,
		Set:      ent.Set,
		Mtime:    ent.Mtime.Format(time.RFC3339),
		Ctime:    ent.Ctime.Format(time.RFC3339),
		Mode:     ent.Mode,
		Mime:     ent.Mime,
		Md5:      md5Str,
		FileSize: ent.FileSize,
		Fid:      ent.Fid,
	}
}

func ToListEntris(ents []Entry) []ListEntry {
	ret := make([]ListEntry, len(ents))

	for i, u := range ents {
		ret[i] = *u.ToListEntry()
	}

	return ret
}
