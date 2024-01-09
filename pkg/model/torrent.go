package model

import (
	"fmt"
)

// 种子摘要信息
type ReTorrentInfo struct {
	Id              uint64
	SiteKey         string              // 站点key
	SiteTorrentId   string              // 站内种子id
	Name            string              // 种子名称
	PieceLength     uint32              // piece大小,Byte
	PiecesHash      *string             // 同NexusPhp的算法相同
	Hash            string              // 种子Hash
	TotalFileLength uint32              // 内部文件大小总和,Byte
	Files           []ReTorrentFileInfo // 内部文件
}

// 种子内部文件摘要信息
type ReTorrentFileInfo struct {
	TorrentId      uint64
	Index          uint32  // 内部顺序，0开始
	Length         uint32  // 文件大小，Byte
	Ext            string  // 文件扩展名，小写
	Path           string  // 文件路径
	PieceQuickHash *string // piece的快速hash值
	QuickHash      *string // 文件的快速hash
	Md5            *string // 文件的md5
}

type PieceQuickHash struct {
	PieceLength uint32 // block大小
	OffsetHead  uint64 // 头部偏移
	OffsetTail  uint64 // 尾部偏移
	Md5         string
	hash        string
}

func (ph *PieceQuickHash) ToHash() string {
	if ph.hash == "" {
		ph.hash = fmt.Sprintf("%d|%d|%d|%s", ph.PieceLength, ph.OffsetHead, ph.OffsetTail, ph.Md5)
	}
	return ph.hash
}
