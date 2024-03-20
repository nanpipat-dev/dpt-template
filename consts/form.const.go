package consts

const (
	FormStatusWaitConsider = "wait_consider" // รอพิจารณา
	FormStatusWaitApprove  = "wait_approve"  // รออนุมัติ

	FormStatusWaitConsiderEdit   = "wait_consider_edit"   // รอพิจารณาแก้ไข
	FormStatusWaitApproveEdit    = "wait_approve_edit"    // รออนุมัติแก้ไข
	FormStatusObsolete           = "obsolete"             // ไม่ใช้แล้ว
	FormStatusWaitConsiderCancel = "wait_consider_cancel" // รอพิจารณายกเลิก
	FormStatusWaitApproveCancel  = "wait_approve_cancel"  // รออนุมัติยกเลิก
	FormStatusCanceled           = "canceled"             // ยกเลิก
	FormStatusEditRejected       = "edit_rejected"        // ถูกปฏิเสธการแก้ไข
	FormStatusCancelRejected     = "cancel_rejected"      // ถูกปฏิเสธการยกเลิก

	FormStatusWaitReport = "wait_report" // รอแจ้งผล
	FormStatusReported   = "reported"    // แจ้งผลแล้ว

	FormAttachmentCodeCancelChangeLetter = "cancel_change_letter" // แจ้งเปลี่ยนหนังสือ
	FormAttachmentCodeCancelAdditional   = "cancel_additional"    // แจ้งเพิ่มเติม

	FormStatusApproved = "approved" // อนุมัติ
	FormStatusRejected = "rejected" // ถูกปฏิเสธ

	RequestFormStatusApprove = "approve"
	RequestFormStatusReject  = "reject"
)
