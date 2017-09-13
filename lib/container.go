package lib

// ContainerUpload contains a list of items
type ContainerUpload struct {
	PrivateUpload
	Items           []*ItemContainer `json:"Items"`
	CurrentLocation int              `json:"CurrentLocation"`
	ContainerType   string           `json:"ContainerType"`
	ContainerGUID   CharacterID      `json:"ContainerGUID"`
}
