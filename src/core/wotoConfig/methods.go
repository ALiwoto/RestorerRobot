package wotoConfig

func (v *ValueSection) GetSectionName() string {
	return v.sectionName
}

func (v *ValueSection) SetSectionName(name string) {
	v.sectionName = name
}
