package manager

import "github.com/louisun/allure-go-v2/allure"

// WithAttachments adds attachment to report in case of current execution context
func (a *allureManager) WithAttachments(attachments ...*allure.Attachment) {
	a.ExecutionContext().AddAttachments(attachments...)
}

// WithNewAttachment creates and adds attachment to report in case of current execution context
func (a *allureManager) WithNewAttachment(name string, mimeType allure.MimeType, content []byte) {
	a.ExecutionContext().AddAttachments(allure.NewAttachment(name, mimeType, content))
}
