package cognito

import (
	"context"
	"errors"
	"restful-rds-golang-products/internal/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// AuthClient es una interfaz que define las operaciones de autenticación con Cognito.
type AuthClient interface {
	SignUp(ctx context.Context, email, password string) error
	ConfirmSignUp(ctx context.Context, email, code string) error
	SignIn(ctx context.Context, email, password string) (*types.AuthenticationResultType, error)
}

// CognitoAuth implementa la interfaz AuthClient para interactuar con AWS Cognito.
type CognitoAuth struct {
	Client     *cognitoidentityprovider.Client
	UserPoolID string
	ClientID   string
}

// NewCognitoAuth crea una nueva instancia de CognitoAuth.
func NewCognitoAuth(client *cognitoidentityprovider.Client, userPoolID, clientID string) *CognitoAuth {
	return &CognitoAuth{
		Client:     client,
		UserPoolID: userPoolID,
		ClientID:   clientID,
	}
}

// SignUp registra un nuevo usuario en Cognito.
func (ca *CognitoAuth) SignUp(ctx context.Context, email, password string) error {
	logger.DebugLog.Printf("Attempting to sign up user: %s", email)
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(ca.ClientID),
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	_, err := ca.Client.SignUp(ctx, input)
	if err != nil {
		logger.ErrorLog.Printf("Error signing up user %s: %v", email, err)
		return formatCognitoError(err)
	}
	logger.InfoLog.Printf("User %s signed up successfully. Confirmation required.", email)
	return nil
}

// ConfirmSignUp confirma el registro de un usuario en Cognito.
func (ca *CognitoAuth) ConfirmSignUp(ctx context.Context, email, code string) error {
	logger.DebugLog.Printf("Attempting to confirm sign up for user: %s with code: %s", email, code)
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(ca.ClientID),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	}

	_, err := ca.Client.ConfirmSignUp(ctx, input)
	if err != nil {
		logger.ErrorLog.Printf("Error confirming sign up for user %s: %v", email, err)
		return formatCognitoError(err)
	}
	logger.InfoLog.Printf("User %s confirmed successfully.", email)
	return nil
}

// SignIn autentica a un usuario y devuelve los tokens de sesión.
func (ca *CognitoAuth) SignIn(ctx context.Context, email, password string) (*types.AuthenticationResultType, error) {
	logger.DebugLog.Printf("Attempting to sign in user: %s", email)
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(ca.ClientID),
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	}

	result, err := ca.Client.InitiateAuth(ctx, input)
	if err != nil {
		logger.ErrorLog.Printf("Error initiating auth for user %s: %v", email, err)
		return nil, formatCognitoError(err)
	}
	logger.InfoLog.Printf("User %s signed in successfully.", email)
	return result.AuthenticationResult, nil
}

// formatCognitoError convierte errores específicos de Cognito en errores más amigables.
func formatCognitoError(err error) error {
	var (
		usernameExistsErr *types.UsernameExistsException
		userNotFoundErr   *types.UserNotFoundException
		codeMismatchErr   *types.CodeMismatchException
		expiredCodeErr    *types.ExpiredCodeException
		notAuthorizedErr  *types.NotAuthorizedException
		limitExceededErr  *types.LimitExceededException
		invalidParamErr   *types.InvalidParameterException
	)

	switch {
	case errors.As(err, &usernameExistsErr):
		return errors.New("el usuario con este correo electrónico ya existe")
	case errors.As(err, &userNotFoundErr):
		return errors.New("usuario no encontrado o credenciales inválidas")
	case errors.As(err, &codeMismatchErr):
		return errors.New("código de confirmación incorrecto")
	case errors.As(err, &expiredCodeErr):
		return errors.New("el código de confirmación ha expirado")
	case errors.As(err, &notAuthorizedErr):
		return errors.New("credenciales inválidas o usuario no confirmado")
	case errors.As(err, &limitExceededErr):
		return errors.New("demasiados intentos, por favor inténtalo de nuevo más tarde")
	case errors.As(err, &invalidParamErr):
		return errors.New("parámetros de solicitud inválidos")
	default:
		return errors.New("error de Cognito desconocido: " + err.Error())
	}
}
