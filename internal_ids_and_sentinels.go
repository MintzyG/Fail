package fail

// IDs
var (
	UnregisteredError           = internalID(0, 0, false, "FailUnregisteredError")
	TranslateUntrustedError     = internalID(0, 1, false, "FailTranslateUntrustedError")
	TranslateNotFound           = internalID(0, 2, false, "FailTranslatorNotFound")
	TranslateUnsupportedError   = internalID(0, 3, false, "FailTranslateUnsupportedError")
	TranslatePanic              = internalID(0, 4, false, "FailTranslatorPanic")
	TranslateWrongType          = internalID(0, 5, false, "FailTranslateWrongType")
	MultipleErrors              = internalID(0, 6, false, "FailMultipleErrors")
	UnknownError                = internalID(0, 7, false, "FailUnknownError")
	NotMatchedInAnyMapper       = internalID(0, 8, false, "FailNotMatchedInAnyMapper")
	NoMapperRegistered          = internalID(0, 9, false, "FailNoMapperRegistered")
	TranslatorAlreadyRegistered = internalID(0, 10, false, "FailTranslatorAlreadyRegistered")
	RuntimeIDInvalid            = internalID(9, 11, false, "FailRuntimeIDInvalid")

	TranslatorNil       = internalID(0, 0, true, "FailTranslatorNil")
	TranslatorNameEmpty = internalID(0, 1, true, "FailTranslatorNameEmpty")
)

// Sentinels
var (
	errUnregisteredError           = Form(UnregisteredError, "error with ID(%s) is not registered in the registry", true, nil, "ID NOT SET")
	errTranslateWrongType          = Form(TranslateWrongType, "translator returned unexpected type", true, nil)
	errTranslateUntrustedError     = Form(TranslateUntrustedError, "tried translating unregistered error", true, nil)
	errTranslateNotFound           = Form(TranslateNotFound, "couldn't find translator", true, nil)
	errTranslateUnsupportedError   = Form(TranslateUnsupportedError, "error not supported by %s translator", true, nil, "UNNAMED")
	errTranslatePanic              = Form(TranslatePanic, "translator panicked during translation", true, nil)
	errTranslatorAlreadyRegistered = Form(TranslatorAlreadyRegistered, "translator already registered", true, nil)
	errTranslatorNil               = Form(TranslatorNil, "cannot register nil translator", true, nil)
	errTranslatorNameEmpty         = Form(TranslatorNameEmpty, "translator must have a non-empty name", true, nil)
	errNotMatchedInAnyMapper       = Form(NotMatchedInAnyMapper, "error wasn't matched/mapped by any mapper", true, nil)
	errNoMapperRegistered          = Form(NoMapperRegistered, "no mapper is registered in the registry", true, nil)
	errMultipleErrors              = Form(MultipleErrors, "multiple errors occurred", false, nil)
	errUnknownError                = Form(UnknownError, "unknown error", true, nil)
	errRuntimeInvalidID            = Form(RuntimeIDInvalid, "all error IDs must be defined at package initialization time and not runtime", true, nil)
)
