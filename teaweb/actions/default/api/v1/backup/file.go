package backup

import (
	"github.com/TeaWeb/code/teaweb/actions/default/settings/backup/backuputils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
)

type FileAction actions.Action

// 下载文件
func (this *FileAction) RunGet(params struct {
	Filename string
}) {
	backuputils.DownloadFile(params.Filename, this.ResponseWriter, func() {
		this.ResponseWriter.WriteHeader(http.StatusNotFound)
		this.WriteString("file not found")
	})
}
