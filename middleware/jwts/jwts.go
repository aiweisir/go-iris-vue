package jwts

import (
	"fmt"
	"go-iris/web/supports"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"

	"time"

	"go-iris/inits/parse"

	"go-iris/web/models"

	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
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

type (
	// A function called whenever an error is encountered
	errorHandler func(context.Context, string)

	// TokenExtractor is a function that takes a context as input and returns
	// either a token or an error.  An error should only be returned if an attempt
	// to specify a token was found, but the information was somehow incorrectly
	// formed.  In the case where a token is simply not present, this should not
	// be treated as an error.  An empty string should be returned in that case.
	TokenExtractor func(context.Context) (string, error)

	// Middleware the middleware for JSON Web tokens authentication method
	Jwts struct {
		Config Config
	}
)

var (
	jwts *Jwts
	lock sync.Mutex
)

// Serve the middleware's action
func Serve(ctx context.Context) bool {
	ConfigJWT()
	if err := jwts.CheckJWT(ctx); err != nil {
		//supports.Unauthorized(ctx, supports.Token_failur, nil)
		//ctx.StopExecution()
		golog.Errorf("Check jwt error, %s", err)
		return false
	}
	return true
	// If everything ok then call next.
	//ctx.Next()
}

// 解析token的信息为当前用户
func ParseToken(ctx context.Context) (*models.User, bool) {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)

	id, ok1 := mapClaims["id"].(float64)
	username, ok2 := mapClaims["username"].(string)

	//golog.Infof("*** MapClaims=%v, [id=%f, ok1=%t]; [username=%s, ok2=%t]", mapClaims, id, ok1, username, ok2)
	if !ok1 || !ok2 {
		supports.Error(ctx, iris.StatusInternalServerError, supports.TokenParseFailur, nil)
		return nil, false
	}

	user := models.User{
		Id:       int64(id),
		Username: username,
	}
	return &user, true
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
	// If an error occurs, call the error handler and return an error
	if err != nil {
		m.logf("Error extracting JWT: %v", err)
		m.Config.ErrorHandler(ctx, supports.TokenExactFailur)
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

		m.logf("  Error: No credentials found (CredentialsOptional=false)")
		// If we get here, the required token is missing
		m.Config.ErrorHandler(ctx, supports.TokenParseFailurAndEmpty)
		return fmt.Errorf(supports.TokenParseFailurAndEmpty)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		m.logf("Error parsing token1: %v", err)
		m.Config.ErrorHandler(ctx, supports.TokenExpire)
		return fmt.Errorf("Error parsing token2: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		m.logf("Error validating token algorithm: %s", message)
		m.Config.ErrorHandler(ctx, supports.TokenParseFailur) // 算法错误
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		m.logf(supports.TokenParseFailurAndInvalid)
		m.Config.ErrorHandler(ctx, supports.TokenParseFailurAndInvalid)
		return fmt.Errorf(supports.TokenParseFailurAndInvalid)
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf(supports.TokenExpire)
			}
		}
	}

	//m.logf("JWT: %v", parsedToken)

	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}

// ------------------------------------------------------------------------
// ------------------------------------------------------------------------

// OnError default error handler
//func OnError(ctx context.Context, err string) {
//	supports.Error(ctx, iris.StatusUnauthorized, supports.Token_Failur, nil)
//}
// jwt中间件配置
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if jwts != nil {
		return
	}

	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(parse.O.Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, errMsg string) {
			supports.Error(ctx, iris.StatusUnauthorized, errMsg, nil)
		},
		// 指定func用于提取请求中的token
		Extractor: FromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	jwts = &Jwts{Config: c}
	//return &Jwts{Config: c}
}

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	//Password string `json:"password"`
	//User models.User `json:"user"`
	jwt.StandardClaims
}

// 在登录成功的时候生成token
func GenerateToken(user *models.User) (string, error) {
	//expireTime := time.Now().Add(60 * time.Second)
	expireTime := time.Now().Add(time.Duration(parse.O.JWTTimeout) * time.Second)

	claims := Claims{
		user.Id,
		user.Username,
		//user.Password,
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

	token, err := tokenClaims.SignedString([]byte(parse.O.Secret))
	return token, err
}
