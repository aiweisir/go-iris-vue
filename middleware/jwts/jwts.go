package jwts

import (
	"casbin-demo/supports"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"time"

	"casbin-demo/conf/parse"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// iris provides some basic middleware, most for your learning courve.
// You can use any net/http compatible middleware with iris.FromStd wrapper.
//
// JWT net/http video tutorial for golang newcomers: https://www.youtube.com/watch?v=dgJFeqeXVKw
//
// Unlike the other middleware, this middleware was cloned from external source: https://github.com/auth0/go-jwt-middleware
// (because it used "context" to define the user but we don't need that so a simple iris.FromStd wouldn't work as expected.)
// jwt_test.go also didn't created by me:
// 28 Jul 2016
// @heralight heralight add jwts unit test.
//
// So if this doesn't works for you just try other net/http compatible middleware and bind it via `iris.FromStd(myHandlerWithNext)`,
// It's here for your learning curve.

// A function called whenever an error is encountered
type errorHandler func(context.Context, string)

// TokenExtractor is a function that takes a context as input and returns
// either a token or an error.  An error should only be returned if an attempt
// to specify a token was found, but the information was somehow incorrectly
// formed.  In the case where a token is simply not present, this should not
// be treated as an error.  An empty string should be returned in that case.
type TokenExtractor func(context.Context) (string, error)

// Middleware the middleware for JSON Web tokens authentication method
type Jwts struct {
	Config Config
}

// Serve the middleware's action
func (m *Jwts) Serve(ctx context.Context) {
	path := ctx.Path()
	fmt.Println("to jwts path:" + path)

	if err := m.CheckJWT(ctx); err != nil {
		//supports.Unauthorized(ctx, supports.Token_failur, nil)
		//ctx.StopExecution()
		return
	}
	// If everything ok then call next.
	ctx.Next()
}


// below 3 method is get token from url
// FromAuthHeader is a "TokenExtractor" that takes a give context and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// below 3 method is get token from url
// FromParameter returns a function that extracts the token from the specified
// query string parameter
func FromParameter(param string) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

// below 3 method is get token from url
// FromFirst returns a function that runs multiple token extractors and takes the
// first token it finds
func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

func (m *Jwts) logf(format string, args ...interface{}) {
	if m.Config.Debug {
		log.Printf(format, args...)
	}
}

// Get returns the user (&token) information for this client/request
func (m *Jwts) Get(ctx context.Context) *jwt.Token {
	return ctx.Values().Get(m.Config.ContextKey).(*jwt.Token)
}


// CheckJWT the main functionality, checks for token
func (m *Jwts) CheckJWT(ctx context.Context) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// Use the specified token extractor to extract a token from the request
	token, err := m.Config.Extractor(ctx)

	// If debugging is turned on, log the outcome
	if err != nil {
		m.logf("Error extracting JWT: %v", err)
	} else {
		//m.logf("Token extracted: %s", token)
	}

	// If an error occurs, call the error handler and return an error
	if err != nil {
		m.Config.ErrorHandler(ctx, err.Error())
		return fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			m.logf("  No credentials found (CredentialsOptional=true)")
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		// If we get here, the required token is missing
		errorMsg := "Required authorization token not found"
		m.Config.ErrorHandler(ctx, errorMsg)
		m.logf("  Error: No credentials found (CredentialsOptional=false)")
		return fmt.Errorf(errorMsg)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		m.logf("Error parsing token: %v", err)
		m.Config.ErrorHandler(ctx, err.Error())
		return fmt.Errorf("Error parsing token: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		m.logf("Error validating token algorithm: %s", message)
		m.Config.ErrorHandler(ctx, errors.New(message).Error())
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		m.logf("Token is invalid")
		m.Config.ErrorHandler(ctx, "The token isn't valid")
		return fmt.Errorf("Token is invalid")
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf("Token is expired")
			}
		}
	}

	m.logf("JWT: %v", parsedToken)

	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}

// ------------------------------------------------------------------------
// ------------------------------------------------------------------------
const SECRET = "My Secret"
// OnError default error handler
//func OnError(ctx context.Context, err string) {
//	supports.Error(ctx, iris.StatusUnauthorized, supports.Token_Failur, nil)
//}
// jwt中间件配置
func ConfigJWT() *Jwts {
	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(SECRET), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, err string) {
			supports.Error(ctx, iris.StatusUnauthorized, supports.Token_failur, nil)
		},
		// 指定func用于提取请求中的token
		Extractor: FromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration: true,
		Debug: true,
		EnableAuthOnOptions: false,
	}
	return &Jwts{Config: c}
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
// 在登录成功的时候生成token
func GenerateToken(username, password string) (string, error) {
	//expireTime := time.Now().Add(60 * time.Second)
	expireTime := time.Now().Add(time.Duration(parse.O.JWTTimeout) * time.Second)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//tokenClaims := jwts.NewWithClaims(jwts.SigningMethodHS256, jwts.MapClaims{
	//	"nick_name": "iris",
	//	"email":     "go-iris@qq.com",
	//	"id":        "1",
	//	"iss":       "Iris",
	//	"iat":       time.Now().Unix(),
	//	"jti":       "9527",
	//	"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
	//})

	token, err := tokenClaims.SignedString([]byte(SECRET))
	return token, err
}
