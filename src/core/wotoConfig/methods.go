package wotoConfig

func (v *ValueSection) GetSectionName() string {
	return v.sectionName
}

func (v *ValueSection) SetSectionName(name string) {
	v.sectionName = name
}

//---------------------------------------------------------

func (d *DatabaseBackupType) IsInvalidType() bool {
	return false
}
