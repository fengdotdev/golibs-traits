package trait

type DataTransferObject[DTO any] interface {
	ToDTO() (DTO, error)
	FromDTO(dto DTO) error
}
