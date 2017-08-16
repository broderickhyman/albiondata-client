package client

type multiUploader struct {
	uploaders []iuploader
}

func newMultiUploader(uploaders []iuploader) iuploader {
	return &multiUploader{
		uploaders: uploaders,
	}
}

func (u *multiUploader) sendToIngest(body []byte, queue string) {
	for _, mu := range u.uploaders {
		mu.sendToIngest(body, queue)
	}
}
