package inventory

import (
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
)

var (
	FullInventoryReply = BuildFullInventoryInfo()
)

func BuildInventoryInfo(u *User) []byte {
	buf := make([]byte, 5+u.Inventory.NumOfItem*19)
	offset := 0
	WriteUint16(&buf, u.Inventory.NumOfItem, &offset)
	for k, v := range u.Inventory.Items {
		WriteUint16(&buf, uint16(k), &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint32(&buf, v.Id, &offset)
		WriteUint16(&buf, v.Count, &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint8(&buf, 0, &offset)
		WriteUint64(&buf, 0, &offset)
	}
	return buf[:offset]
}

func BuildFullInventoryInfo() []byte {
	buf := make([]byte, 5+uint16(len(FullInventoryItem))*19)
	offset := 0
	WriteUint16(&buf, uint16(len(FullInventoryItem)), &offset)
	for k, v := range FullInventoryItem {
		WriteUint16(&buf, uint16(k), &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint32(&buf, v.Id, &offset)
		WriteUint16(&buf, v.Count, &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint8(&buf, 0, &offset)
		WriteUint64(&buf, 0, &offset)
	}
	return buf[:offset]
}

func BuildDefaultInventoryInfo() []byte {
	DeafaultInventoryItem := CreateDefaultInventoryItem()
	buf := make([]byte, 5+len(DeafaultInventoryItem)*19)
	offset := 0
	WriteUint16(&buf, 25, &offset)
	for k, v := range DeafaultInventoryItem {
		WriteUint16(&buf, uint16(k), &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint32(&buf, v.Id, &offset)
		WriteUint16(&buf, v.Count, &offset)
		WriteUint8(&buf, 1, &offset)
		WriteUint8(&buf, 0, &offset)
		WriteUint64(&buf, 0, &offset)

	}
	return buf[:offset]
}

func BuildUnlockReply() []byte {
	buf := make([]byte, 4096)
	offset := 0
	WriteUint8(&buf, 1, &offset)                            //type ?
	WriteUint16(&buf, uint16(len(UnlockFullList)), &offset) //num of weapons
	for _, v := range UnlockFullList {
		WriteUint32(&buf, v.Itemid, &offset)
		WriteUint32(&buf, v.Seq, &offset)
		WriteUint8(&buf, v.CostType, &offset)
		WriteUint32(&buf, v.Price, &offset)
	}
	WriteUint16(&buf, 1, &offset) //num of weapons

	WriteUint32(&buf, 2, &offset) //前置
	WriteUint32(&buf, 1, &offset) //当前
	WriteUint32(&buf, 2, &offset) //杀敌数
	WriteUint16(&buf, 1, &offset)
	WriteUint16(&buf, 1, &offset)
	WriteUint16(&buf, 1, &offset)

	WriteUint16(&buf, 1, &offset) //unk

	WriteUint32(&buf, 1, &offset)

	return buf[:offset]
}

func BuildDefaultUnlockReply() []byte {
	return []byte{0x01, 0x4B, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x0B, 0x00, 0x00, 0x00, 0x01, 0xE8, 0x03, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x0C, 0x00,
		0x00, 0x00, 0x01, 0xDC, 0x05, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00, 0x0D, 0x00, 0x00, 0x00, 0x01,
		0xE8, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x01, 0xDC, 0x05, 0x00,
		0x00, 0x0B, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x01, 0x08, 0x07, 0x00, 0x00, 0x3C, 0x00,
		0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0x80, 0xBB, 0x00, 0x00, 0x1F, 0x00, 0x00, 0x00, 0x11,
		0x00, 0x00, 0x00, 0x01, 0xC0, 0x5D, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00,
		0x01, 0x08, 0x07, 0x00, 0x00, 0x1C, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x01, 0x4C, 0x1D,
		0x00, 0x00, 0x3B, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x01, 0x60, 0x61, 0x02, 0x00, 0x35,
		0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x01, 0x30, 0x75, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
		0x16, 0x00, 0x00, 0x00, 0x01, 0xA0, 0x0F, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00,
		0x00, 0x01, 0x98, 0x3A, 0x00, 0x00, 0x3F, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x01, 0xE0,
		0x93, 0x04, 0x00, 0x14, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00, 0x01, 0xA0, 0x0F, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00, 0x01, 0x98, 0x3A, 0x00, 0x00, 0x3E, 0x00, 0x00,
		0x00, 0x1B, 0x00, 0x00, 0x00, 0x01, 0xE0, 0x93, 0x04, 0x00, 0x05, 0x00, 0x00, 0x00, 0x1C, 0x00,
		0x00, 0x00, 0x01, 0x08, 0x07, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x1D, 0x00, 0x00, 0x00, 0x01,
		0x30, 0x75, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x1E, 0x00, 0x00, 0x00, 0x01, 0x88, 0x13, 0x00,
		0x00, 0x0C, 0x00, 0x00, 0x00, 0x1F, 0x00, 0x00, 0x00, 0x01, 0x20, 0x4E, 0x00, 0x00, 0x16, 0x00,
		0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x20, 0x4E, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x43,
		0x00, 0x00, 0x00, 0x01, 0x30, 0x75, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x57, 0x00, 0x00, 0x00,
		0x01, 0x20, 0xA1, 0x07, 0x00, 0x47, 0x00, 0x00, 0x00, 0x58, 0x00, 0x00, 0x00, 0x01, 0x20, 0xA1,
		0x07, 0x00, 0x4D, 0x00, 0x00, 0x00, 0x59, 0x00, 0x00, 0x00, 0x00, 0x90, 0x01, 0x00, 0x00, 0x55,
		0x00, 0x00, 0x00, 0x81, 0x00, 0x00, 0x00, 0x00, 0x70, 0x03, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x90, 0x00, 0x00, 0x00, 0x01, 0x30, 0x75, 0x00, 0x00, 0x1D, 0x00, 0x00, 0x00, 0x91, 0x00, 0x00,
		0x00, 0x01, 0x60, 0xEA, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x92, 0x00, 0x00, 0x00, 0x01, 0x48,
		0xE8, 0x01, 0x00, 0x2F, 0x00, 0x00, 0x00, 0x93, 0x00, 0x00, 0x00, 0x01, 0x40, 0x0D, 0x03, 0x00,
		0x6A, 0xBF, 0x00, 0x00, 0xA8, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x70, 0xBF, 0x00,
		0x00, 0xA9, 0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x6F, 0xBF, 0x00, 0x00, 0xAA, 0x00,
		0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x6E, 0xBF, 0x00, 0x00, 0xAB, 0x00, 0x00, 0x00, 0x00,
		0x50, 0x00, 0x00, 0x00, 0x69, 0xBF, 0x00, 0x00, 0xAC, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00,
		0x00, 0x72, 0xBF, 0x00, 0x00, 0xAD, 0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x6B, 0xBF,
		0x00, 0x00, 0xAE, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x6D, 0xBF, 0x00, 0x00, 0xAF,
		0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x4A, 0x00, 0x00, 0x00, 0xD7, 0x00, 0x00, 0x00,
		0x01, 0x50, 0xC3, 0x00, 0x00, 0x4B, 0x00, 0x00, 0x00, 0xD8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x77,
		0x01, 0x00, 0x4E, 0x00, 0x00, 0x00, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x70, 0x11, 0x01, 0x00, 0x52,
		0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x01, 0xC0, 0xD4, 0x01, 0x00, 0x5B, 0x00, 0x00, 0x00,
		0x06, 0x01, 0x00, 0x00, 0x01, 0xF0, 0x49, 0x02, 0x00, 0x5F, 0x00, 0x00, 0x00, 0x19, 0x01, 0x00,
		0x00, 0x01, 0x60, 0xEA, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x1A, 0x01, 0x00, 0x00, 0x01, 0xC0,
		0xD4, 0x01, 0x00, 0x64, 0x00, 0x00, 0x00, 0x38, 0x01, 0x00, 0x00, 0x01, 0xF0, 0x49, 0x02, 0x00,
		0x68, 0x00, 0x00, 0x00, 0x5C, 0x01, 0x00, 0x00, 0x01, 0x20, 0xA1, 0x07, 0x00, 0x6D, 0x00, 0x00,
		0x00, 0x82, 0x01, 0x00, 0x00, 0x01, 0xA0, 0x86, 0x01, 0x00, 0x6C, 0x00, 0x00, 0x00, 0x83, 0x01,
		0x00, 0x00, 0x01, 0xA0, 0x86, 0x01, 0x00, 0x6E, 0x00, 0x00, 0x00, 0x84, 0x01, 0x00, 0x00, 0x01,
		0xA0, 0x86, 0x01, 0x00, 0x42, 0x00, 0x00, 0x00, 0xFA, 0x01, 0x00, 0x00, 0x01, 0x30, 0x75, 0x00,
		0x00, 0x43, 0x00, 0x00, 0x00, 0xFB, 0x01, 0x00, 0x00, 0x01, 0x50, 0xC3, 0x00, 0x00, 0x78, 0x00,
		0x00, 0x00, 0xFC, 0x01, 0x00, 0x00, 0x01, 0x40, 0x0D, 0x03, 0x00, 0x79, 0x00, 0x00, 0x00, 0x07,
		0x02, 0x00, 0x00, 0x00, 0xA0, 0x00, 0x00, 0x00, 0x7C, 0x00, 0x00, 0x00, 0x08, 0x02, 0x00, 0x00,
		0x00, 0x04, 0x01, 0x00, 0x00, 0x7A, 0x00, 0x00, 0x00, 0x09, 0x02, 0x00, 0x00, 0x00, 0xE0, 0x01,
		0x00, 0x00, 0x7B, 0x00, 0x00, 0x00, 0x0A, 0x02, 0x00, 0x00, 0x00, 0x44, 0x02, 0x00, 0x00, 0x7D,
		0x00, 0x00, 0x00, 0x58, 0x02, 0x00, 0x00, 0x00, 0x44, 0x02, 0x00, 0x00, 0x7E, 0x00, 0x00, 0x00,
		0x59, 0x02, 0x00, 0x00, 0x00, 0x0C, 0x03, 0x00, 0x00, 0x81, 0x00, 0x00, 0x00, 0x91, 0x02, 0x00,
		0x00, 0x01, 0xF0, 0x49, 0x02, 0x00, 0x82, 0x00, 0x00, 0x00, 0x92, 0x02, 0x00, 0x00, 0x01, 0x00,
		0x53, 0x07, 0x00, 0x83, 0x00, 0x00, 0x00, 0x93, 0x02, 0x00, 0x00, 0x01, 0x60, 0x5B, 0x03, 0x00,
		0x85, 0x00, 0x00, 0x00, 0x94, 0x02, 0x00, 0x00, 0x00, 0x40, 0x01, 0x00, 0x00, 0x84, 0x00, 0x00,
		0x00, 0x95, 0x02, 0x00, 0x00, 0x00, 0x08, 0x02, 0x00, 0x00, 0x87, 0x00, 0x00, 0x00, 0x1F, 0x03,
		0x00, 0x00, 0x00, 0x08, 0x02, 0x00, 0x00, 0x8A, 0x00, 0x00, 0x00, 0xA4, 0x03, 0x00, 0x00, 0x01,
		0xE0, 0x93, 0x04, 0x00, 0x8F, 0x00, 0x00, 0x00, 0x44, 0x04, 0x00, 0x00, 0x01, 0x80, 0xA9, 0x03,
		0x00, 0x90, 0x00, 0x00, 0x00, 0x45, 0x04, 0x00, 0x00, 0x01, 0x40, 0x7E, 0x05, 0x00, 0x91, 0x00,
		0x00, 0x00, 0x46, 0x04, 0x00, 0x00, 0x01, 0x00, 0x53, 0x07, 0x00, 0x9B, 0x00, 0x00, 0x00, 0xA9,
		0x04, 0x00, 0x00, 0x01, 0xF0, 0x49, 0x02, 0x00, 0x9C, 0x00, 0x00, 0x00, 0xAA, 0x04, 0x00, 0x00,
		0x01, 0x40, 0x0D, 0x03, 0x00, 0x97, 0x00, 0x00, 0x00, 0xFC, 0x04, 0x00, 0x00, 0x01, 0x42, 0x99,
		0x00, 0x00, 0x98, 0x00, 0x00, 0x00, 0xFD, 0x04, 0x00, 0x00, 0x01, 0x86, 0x29, 0x02, 0x00, 0x99,
		0x00, 0x00, 0x00, 0xFE, 0x04, 0x00, 0x00, 0x01, 0x8C, 0xED, 0x02, 0x00, 0x10, 0x00, 0x03, 0x00,
		0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x42, 0x00, 0x00, 0x00, 0x43, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x07, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x1C, 0x00,
		0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00,
		0x35, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00,
		0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x13, 0x00, 0x00, 0x00, 0x4D, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x3E, 0x00, 0x00, 0x00, 0x08, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00, 0x3F, 0x00,
		0x00, 0x00, 0x1A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00,
		0x19, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00,
		0x06, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00,
		0x0B, 0x00, 0x00, 0x00, 0x0D, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00,
		0x10, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00,
		0x14, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00,
		0x1A, 0x00, 0x00, 0x00, 0x1C, 0x00, 0x00, 0x00, 0x6C, 0xBF, 0x00, 0x00, 0x71, 0xBF, 0x00, 0x00,
		0x42, 0x00, 0x00, 0x00, 0x94, 0x01, 0x00, 0x00}
}
