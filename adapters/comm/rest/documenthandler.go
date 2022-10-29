package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/gitdoc/adapters/comm/rest/dto"
	"github.com/serdarkalayci/gitdoc/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/gitdoc/adapters/comm/rest/middleware"
	"github.com/serdarkalayci/gitdoc/application"
)

type validateddocument struct{}

// swagger:route GET /people document GetDocuments
// Return all the documents
// responses:
//	200: OK
//	500: errorResponse

// GetDocuments gets all the documents of the Titanic
func (ctx *APIContext) GetDocuments(rw http.ResponseWriter, r *http.Request) {
	span := createSpan("Titanic.ListAll", r)
	defer span.Finish()

	DocumentService := application.NewDocumentService(ctx.documentRepo)
	documents, err := DocumentService.List()
	if err != nil {
		respondWithError(rw, r, 500, "Cannot get documents from database")
	} else {
		documentDTOs := make([]dto.DocumentResponseDTO, 0)
		for _, p := range documents {
			pDTO := mappers.Mapdocument2documentResponseDTO(p)
			documentDTOs = append(documentDTOs, pDTO)
		}
		respondWithJSON(rw, r, 200, documentDTOs)
	}

}

// swagger:route POST /people document Adddocument
// Adds a new document
// responses:
//	201: Created
//	500: errorResponse

// Adddocument adds a new documents to the Titanic
func (ctx *APIContext) Adddocument(rw http.ResponseWriter, r *http.Request) {
	span := createSpan("Titanic.Add", r)
	defer span.Finish()
	// Get document data from payload
	documentDTO := r.Context().Value(validateddocument{}).(dto.DocumentRequestDTO)
	document := mappers.MapdocumentRequestDTO2document(documentDTO)
	DocumentService := application.NewDocumentService(ctx.documentRepo)
	document, err := DocumentService.Add(document)
	if err != nil {
		respondWithError(rw, r, 500, err.Error())
	} else {
		pDTO := mappers.Mapdocument2documentResponseDTO(document)
		respondWithJSON(rw, r, 201, pDTO)
	}
}

// swagger:route GET /people/{id} document GetDocument
// Return the document with the given id
// responses:
//	200: OK
//  400: Bad Request
//	500: errorResponse

// GetDocument gets the documents of the Titanic with the given id
func (ctx *APIContext) GetDocument(rw http.ResponseWriter, r *http.Request) {
	span := createSpan("Titanic.GetOne", r)
	defer span.Finish()

	// parse the document id from the url
	vars := mux.Vars(r)
	id := vars["id"]
	DocumentService := application.NewDocumentService(ctx.documentRepo)
	document, err := DocumentService.Get(id)
	if err != nil {
		switch err.(type) {
		case *application.ErrorIDFormat:
			respondWithError(rw, r, 400, "Cannot process with the given id")
		case *application.ErrorCannotFinddocument:
			respondWithError(rw, r, 404, "Cannot get document from database")
		default:
			respondWithError(rw, r, 500, "Internal server error")
		}
	} else {
		pDTO := mappers.Mapdocument2documentResponseDTO(document)
		respondWithJSON(rw, r, 200, pDTO)
	}
}

// swagger:route PUT /people{id} document UpdateDocument
// Updates an existing document
// responses:
//	201: Created
//  400: Bad Request
//	500: errorResponse

// UpdateDocument updates an existing documents on the Titanic
func (ctx *APIContext) UpdateDocument(rw http.ResponseWriter, r *http.Request) {
	span := createSpan("Titanic.Update", r)
	defer span.Finish()

	// parse the document id from the url
	vars := mux.Vars(r)
	id := vars["id"]
	// Get document data from payload
	documentDTO := r.Context().Value(validateddocument{}).(dto.DocumentRequestDTO)
	document := mappers.MapdocumentRequestDTO2document(documentDTO)
	DocumentService := application.NewDocumentService(ctx.documentRepo)
	err := DocumentService.Update(id, document)
	if err != nil {
		switch err.(type) {
		case *application.ErrorIDFormat:
			respondWithError(rw, r, 400, "Cannot process with the given id")
		case *application.ErrorCannotFinddocument:
			respondWithError(rw, r, 404, "Cannot get document from database")
		default:
			respondWithError(rw, r, 500, "Internal server error")
		}
	} else {
		respondEmpty(rw, r, 201)
	}
}

// swagger:route DELETE /people/{id} document DeleteDocument
// Deletes the document with the given id
// responses:
//	200: OK
//  400: Bad Request
//	500: errorResponse

// DeleteDocument deletes the documents of the Titanic with the given id
func (ctx *APIContext) DeleteDocument(rw http.ResponseWriter, r *http.Request) {
	span := createSpan("Titanic.Delete", r)
	defer span.Finish()

	// parse the document id from the url
	vars := mux.Vars(r)
	id := vars["id"]
	DocumentService := application.NewDocumentService(ctx.documentRepo)
	err := DocumentService.Delete(id)
	if err != nil {
		switch err.(type) {
		case *application.ErrorIDFormat:
			respondWithError(rw, r, 400, "Cannot process with the given id")
		case *application.ErrorCannotFinddocument:
			respondWithError(rw, r, 404, "Cannot get document from database")
		default:
			respondWithError(rw, r, 500, "Internal server error")
		}
	} else {
		respondEmpty(rw, r, 200)
	}
}

// MiddlewareValidateNewDocument Checks the integrity of new document in the request and calls next if ok
func (ctx *APIContext) MiddlewareValidateNewDocument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := middleware.ExtractAdddocumentPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the user
		errs := ctx.validation.Validate(user)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the document")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validateddocument{}, *user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
