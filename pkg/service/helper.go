package service

import (
	"fmt"

	ovirtsdk "github.com/ovirt/go-ovirt"
)

type AttachmentNotFoundError struct {
	vmId   string
	diskId string
}

func (e *AttachmentNotFoundError) Error() string {
	return fmt.Sprintf("failed to find attachment by disk %s for VM %s", e.diskId, e.vmId)
}

func diskAttachmentByVmAndDisk(connection *ovirtsdk.Connection, vmId string, diskId string) (*ovirtsdk.DiskAttachment, error) {
	vmService := connection.SystemService().VmsService().VmService(vmId)
	attachments, err := vmService.DiskAttachmentsService().List().Send()
	if err != nil {
		return nil, err
	}

	for _, attachment := range attachments.MustAttachments().Slice() {
		if diskId == attachment.MustDisk().MustId() {
			return attachment, nil
		}
	}
	return nil, &AttachmentNotFoundError{
		vmId:   vmId,
		diskId: diskId,
	}
}
