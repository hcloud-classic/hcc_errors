package hcc_errors

const (
	PiccoloInternalInitFail            = piccolo + internal + initFail
	PiccoloInternalConnectionFail      = piccolo + internal + connectionFail
	PiccoloInternalUUIDGenerationError = piccolo + internal + UUIDGenerationError

	PiccoloGrpcRequestError = piccolo + grpc + requestError

	PiccoloGraphQLTimestampConversionError = piccolo + graphql + timestampConversionError
	PiccoloGraphQLArgumentError            = piccolo + graphql + argumentError
	PiccoloGraphQLLoginFailed              = piccolo + graphql + loginFailed
	PiccoloGraphQLTokenGenerationError     = piccolo + graphql + tokenGenerationError
	PiccoloGraphQLTokenExpired             = piccolo + graphql + tokenExpired
	PiccoloGraphQLInvalidToken             = piccolo + graphql + invalidToken
	PiccoloGraphQLUserExist                = piccolo + graphql + userExist

	PiccoloMySQLPrepareError = piccolo + sql + prepareError
	PiccoloMySQLExecuteError = piccolo + sql + executeError
)
