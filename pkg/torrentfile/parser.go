package torrentfile

import (
	"github.com/anacrolix/torrent/metainfo"
	"github.com/dustin/go-humanize"
)

const (
	// HashBlockNum 计算block的hash所需的要的数量
	HashBlockNum = 3
	// HashBlockMinFileSize 满足大小的文件才会计算
	HashBlockMinFileSize = 10 * humanize.MiByte
)

// ParseFile 解析种子文件 fp是全路径
func ParseFile(fp string) error {
	meta, err := metainfo.LoadFromFile(fp)
	if err != nil {
		return err
	}
	info, _ := meta.UnmarshalInfo()
}
