package mappers

import (
	"github.com/serdarkalayci/gitdoc/adapters/comm/rest/dto"
	"github.com/serdarkalayci/gitdoc/domain"
)

func MapdocumentRequestDTO2document(doc dto.DocumentRequestDTO) domain.Document {
	return domain.Document{
		Name: doc.Name,
	}
}

func Mapdocument2documentResponseDTO(doc domain.Document) dto.DocumentResponseDTO {
	return dto.DocumentResponseDTO{
		ID:   doc.ID,
		Name: doc.Name,
	}
}
