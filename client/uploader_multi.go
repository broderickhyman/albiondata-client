package client

type multiUploader struct {
	uploaders []iuploader
}

func newMultiUploader(uploaders []iuploader) iuploader {
	return &multiUploader{
		uploaders: uploaders,
	}
}

func (u *multiUploader) private() bool {
	return true
}

func (u *multiUploader) sendToPublicIngest(body []byte, queue string) {
	for _, mu := range u.uploaders {
		if mu.private() {
			mu.sendToPublicIngest(body, queue)
		}
	}
}

func (u *multiUploader) sendToPrivateIngest(body []byte, queue string) {
	for _, mu := range u.uploaders {
		mu.sendToPrivateIngest(body, queue)
	}
}
