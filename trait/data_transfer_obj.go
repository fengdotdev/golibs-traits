package trait

// NOT READY FOR PRODUCTION
type DataTransferObject[DTO any] interface {
	ToDTO() (DTO, error)
	FromDTO(dto DTO) error
}
