package memory

import (
	"carrega/daemon/models"
)

type DnNode struct {
	Download *models.DownloadProcess
	Next     *DnNode
}

var (
	Ongoing  DnNode = DnNode{}
	Finished DnNode = DnNode{}
)

func (n *DnNode) add(download *models.DownloadProcess) {
	if n.Next != nil {
		n.Next.add(download)
		return
	}

	if n.Download == nil {
		n.Download = download
		return
	}

	n.Next = &DnNode{Download: download}
}

func (n *DnNode) remove(download *models.DownloadProcess) {
	/* first on the list exceptions, ran once */
	if n.Download == download {
		if n.Next == nil {
			n.Download = nil
			return
		}

		*n = *n.Next
	}

	/* normal conditions */
	if n.Next == nil {
		return
	}

	if n.Next.Download == download {
		n.Next = n.Next.Next
	}

	n.Next = &DnNode{Download: download}
}

func (n *DnNode) ToBytes() []byte {
	if n.Next == nil {
		if n.Download == nil {
			return nil
		}

		return []byte(n.Download.FileName)
	}

	return append(
		append([]byte(n.Download.FileName), []byte(";;")...),
		n.Next.ToBytes()...,
	)
}
