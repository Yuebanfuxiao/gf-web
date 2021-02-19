package global

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
	"io/ioutil"
	"strings"
	"time"
)

const (
	JwtErrorMissingSecretKey uint32 = 1 << iota // Token is malformed
	JwtErrorExpiredToken
	JwtErrorEmptyAuthHeader
	JwtErrorMissingExpField
	JwtErrorWrongFormatOfExp
	JwtErrorInvalidPublicKey
	JwtErrorInvalidPrivateKey
	JwtErrorNoPublicKeyFile
	JwtErrorNoPrivateKeyFile
	JwtErrorInvalidSigningAlgorithm
	JwtErrorEmptyToken
	JwtErrorInvalidAuthHeader
	JwtErrorEmptyQueryToken
	JwtErrorEmptyCookieToken
	JwtErrorEmptyParamToken
	JwtErrorFailedTokenCreation
	JwtErrorFailedTokenDestroy
	JwtErrorInvalidToken
	JwtErrorAuthorizeElsewhere
)

var (
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = NewJwtError("secret key is required", JwtErrorMissingSecretKey)

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = NewJwtError("failed to create JWT Token", JwtErrorFailedTokenCreation)

	// ErrFailedTokenDestroy indicates JWT Token failed to destroy, reason unknown
	ErrFailedTokenDestroy = NewJwtError("failed to destroy JWT Token", JwtErrorFailedTokenDestroy)

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = NewJwtError("token is expired", JwtErrorExpiredToken)

	// ErrAuthElsewhere authorize elsewhere.
	ErrAuthorizeElsewhere = NewJwtError("sign in elsewhere", JwtErrorAuthorizeElsewhere)

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = NewJwtError("auth header is empty", JwtErrorEmptyAuthHeader)

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = NewJwtError("missing exp field", JwtErrorMissingExpField)

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = NewJwtError("exp must be float64 format", JwtErrorWrongFormatOfExp)

	// ErrInvalidToken indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidToken = NewJwtError("token is invalid", JwtErrorInvalidToken)

	// ErrEmptyToken can be thrown if authing token is invalid, the token variable is empty
	ErrEmptyToken = NewJwtError("token is empty", JwtErrorEmptyToken)

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = NewJwtError("auth header is invalid", JwtErrorInvalidAuthHeader)

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = NewJwtError("query token is empty", JwtErrorEmptyQueryToken)

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cookie is empty
	ErrEmptyCookieToken = NewJwtError("cookie token is empty", JwtErrorEmptyCookieToken)

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = NewJwtError("parameter token is empty", JwtErrorEmptyParamToken)

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = NewJwtError("invalid signing algorithm", JwtErrorInvalidSigningAlgorithm)

	// ErrNoPrivateKeyFile indicates that the given private key is unreadable
	ErrNoPrivateKeyFile = NewJwtError("private key file unreadable", JwtErrorNoPrivateKeyFile)

	// ErrNoPublicKeyFile indicates that the given public key is unreadable
	ErrNoPublicKeyFile = NewJwtError("public key file unreadable", JwtErrorNoPublicKeyFile)

	// ErrInvalidPrivateKey indicates that the given private key is invalid
	ErrInvalidPrivateKey = NewJwtError("private key invalid", JwtErrorInvalidPrivateKey)

	// ErrInvalidPublicKey indicates the the given public key is invalid
	ErrInvalidPublicKey = NewJwtError("public key invalid", JwtErrorInvalidPublicKey)
)

func NewJwtError(errorText string, errorFlags uint32) *JwtError {
	return &JwtError{
		text:   errorText,
		Errors: errorFlags,
	}
}

type JwtError struct {
	Inner  error
	Errors uint32
	text   string
}

// Validation error is an error type
func (o *JwtError) Error() string {
	if o.Inner != nil {
		return o.Inner.Error()
	}

	return o.text
}

// No errors
func (o *JwtError) valid() bool {
	return o.Errors == 0
}

type Jwt struct {
	Realm          string          // 适用范围
	Algorithm      string          // 算法
	Secret         string          // 密钥
	Timeout        time.Duration   // TOKEN超时时间
	Refresh        time.Duration   // TOKEN刷新时间
	Unique         bool            // 唯一限制
	IdentityKey    string          // 身份识别KEY
	TokenLookup    string          // TOKEN查找位置
	TokenHeadName  string          // TOKEN头名
	PublicKeyFile  string          // 公钥文件
	PrivateKeyFile string          // 私钥文件
	publicKey      *rsa.PublicKey  // 公钥
	privateKey     *rsa.PrivateKey // 私钥
	key            []byte          // 密钥
}

type Token struct {
	Token  string
	Expire string
	Type   string
}

// 创建JWT
func NewJwt(o *Jwt) (*Jwt, error) {
	if err := o.init(); err != nil {
		return nil, err
	}

	return o, nil
}

// 初始化
func (o *Jwt) init() error {
	if o.TokenLookup == "" {
		o.TokenLookup = "header:Authorization"
	}

	if o.Algorithm == "" {
		o.Algorithm = "HS256"
	}

	if o.Timeout == 0 {
		o.Timeout = time.Hour
	}

	o.TokenHeadName = strings.TrimSpace(o.TokenHeadName)
	if len(o.TokenHeadName) == 0 {
		o.TokenHeadName = "Bearer"
	}

	if o.usingRsaKeyAlgo() {
		return o.readRsaKeys()
	}

	o.key = []byte(o.Secret)
	if o.key == nil {
		return ErrMissingSecretKey
	}

	return nil
}

// 生成TOKEN
func (o *Jwt) GenerateToken(data interface{}) (*Token, error) {
	token := jwt.New(jwt.GetSigningMethod(o.Algorithm))

	claims := token.Claims.(jwt.MapClaims)

	for key, value := range o.handlePayload(data) {
		claims[key] = value
	}

	expire := time.Now().Add(o.Timeout)

	claims["exp"] = expire.Unix()

	claims["orig_iat"] = time.Now().Unix()

	claims["uic"] = guid.S()

	tokenString, err := o.signedString(token)

	if err != nil {
		return nil, ErrFailedTokenCreation
	}

	if o.Unique {
		o.SetUniqueIdentificationCode(claims)
	}

	return &Token{
		Token:  tokenString,
		Expire: expire.Format(time.RFC3339),
		Type:   o.TokenHeadName,
	}, nil
}

// 刷新TOKEN
func (o *Jwt) RefreshToken(request *ghttp.Request) (*Token, error) {
	claims, err := o.CheckIfTokenExpire(request)

	if err != nil {
		return nil, err
	}

	if err := o.checkUniqueIdentificationCode(claims); err != nil {
		return nil, err
	}

	newToken := jwt.New(jwt.GetSigningMethod(o.Algorithm))

	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire := time.Now().Add(o.Timeout)

	newClaims["exp"] = expire.Unix()

	newClaims["orig_iat"] = time.Now().Unix()

	newClaims["uic"] = guid.S()

	tokenString, err := o.signedString(newToken)

	if err != nil {
		return nil, ErrFailedTokenCreation
	}

	if o.Unique {
		o.SetUniqueIdentificationCode(newClaims)
	}

	return &Token{
		Token:  tokenString,
		Expire: expire.Format(time.RFC3339),
		Type:   o.TokenHeadName,
	}, nil
}

// 销毁TOKEN
func (o *Jwt) DestroyToken(request *ghttp.Request) error {
	if o.Unique {
		if claims, err := o.GetClaims(request); err != nil {
			return err
		} else {
			if claims["uic"] == o.GetUniqueIdentificationCode(claims[o.IdentityKey]) {
				return o.DelUniqueIdentificationCode(claims[o.IdentityKey])
			}
		}
	}

	return nil
}

// 检测TOKEN是否过期
func (o *Jwt) CheckIfTokenExpire(r *ghttp.Request) (jwt.MapClaims, error) {
	token, err := o.parseToken(r)

	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)

		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, ErrInvalidToken
		}
	}

	if token == nil {
		return nil, ErrEmptyToken
	}

	claims := token.Claims.(jwt.MapClaims)

	origIat := int64(claims["orig_iat"].(float64))

	if origIat < time.Now().Add(-o.Refresh).Unix() {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// 解析TOKEN
func (o *Jwt) parseToken(request *ghttp.Request) (*jwt.Token, error) {
	var tokenString string
	var err error

	for _, method := range strings.Split(o.TokenLookup, ",") {
		if len(tokenString) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			tokenString, err = o.getTokenFromHeader(request, v)
		case "query":
			tokenString, err = o.getTokenFromQuery(request, v)
		case "cookie":
			tokenString, err = o.getTokenFromCookie(request, v)
		case "param":
			tokenString, err = o.getTokenFromParam(request, v)
		}
	}

	if err != nil {
		return nil, err
	}

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(o.Algorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}

		if o.usingRsaKeyAlgo() {
			return o.publicKey, nil
		}

		// save token string if valid
		request.SetParam("JWT_TOKEN", tokenString)

		return o.key, nil
	})
}

// 中间件
func (o *Jwt) Middleware(request *ghttp.Request) error {
	claims, err := o.GetClaimsFromToken(request)

	if err != nil {
		return err
	}

	if claims["exp"] == nil {
		return ErrMissingExpField
	}

	if _, ok := claims["exp"].(float64); !ok {
		return ErrWrongFormatOfExp
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return ErrExpiredToken
	}

	if err := o.checkUniqueIdentificationCode(claims); err != nil {
		return err
	}

	o.SetClaims(request, claims)

	return nil
}

// 获取约定参数
func (o *Jwt) GetClaimsFromToken(request *ghttp.Request) (jwt.MapClaims, error) {
	token, err := o.parseToken(request)

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims := jwt.MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims, nil
}

// 获取身份识别
func (o *Jwt) GetIdentity(request *ghttp.Request) interface{} {
	if claims, err := o.GetClaims(request); err != nil {
		return nil
	} else {
		return claims[o.IdentityKey]
	}
}

// 获取约定参数
func (o *Jwt) GetClaims(request *ghttp.Request) (jwt.MapClaims, error) {
	claims := request.GetParam("JWT_PAYLOAD")

	if claims != nil {
		return claims.(jwt.MapClaims), nil
	}

	return o.GetClaimsFromToken(request)
}

// 设置约定参数
func (o *Jwt) SetClaims(request *ghttp.Request, claims jwt.MapClaims) {
	request.SetParam("JWT_PAYLOAD", claims)
}

// 读取公钥私钥
func (o *Jwt) readRsaKeys() error {
	if err := o.readPrivateKey(); err != nil {
		return err
	}

	if err := o.readPublicKey(); err != nil {
		return err
	}

	return nil
}

// 读取私钥
func (o *Jwt) readPrivateKey() error {
	keyData, err := ioutil.ReadFile(o.PrivateKeyFile)

	if err != nil {
		return ErrNoPrivateKeyFile
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	if err != nil {
		return ErrInvalidPrivateKey
	}

	o.privateKey = key

	return nil
}

// 读取公钥
func (o *Jwt) readPublicKey() error {
	keyData, err := ioutil.ReadFile(o.PublicKeyFile)

	if err != nil {
		return ErrNoPublicKeyFile
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)

	if err != nil {
		return ErrInvalidPublicKey
	}

	o.publicKey = key

	return nil
}

// 签名字符串
func (o *Jwt) signedString(token *jwt.Token) (string, error) {
	var tokenString string

	var err error

	if o.usingRsaKeyAlgo() {
		tokenString, err = token.SignedString(o.privateKey)
	} else {
		tokenString, err = token.SignedString(o.key)
	}

	if err != nil {
		return "", ErrFailedTokenCreation
	}

	return tokenString, nil
}

// 检测是否使用公私钥算法
func (o *Jwt) usingRsaKeyAlgo() bool {
	switch o.Algorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

// 处理载荷数据
func (o *Jwt) handlePayload(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}

	params := data.(map[string]interface{})

	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}

	return claims
}

// 从HEADER获取TOKEN
func (o *Jwt) getTokenFromHeader(request *ghttp.Request, key string) (string, error) {
	authHeader := request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == o.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

// 从QUERY中获取TOKEN
func (o *Jwt) getTokenFromQuery(request *ghttp.Request, key string) (string, error) {
	token := request.GetString(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

// 从COOKIE中获取TOKEN
func (o *Jwt) getTokenFromCookie(request *ghttp.Request, key string) (string, error) {
	cookie := request.Cookie.Get(key)

	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

// 从PARAM中获取TOKEN
func (o *Jwt) getTokenFromParam(request *ghttp.Request, key string) (string, error) {
	token := request.GetString(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}

// 检测唯一识别码
func (o *Jwt) checkUniqueIdentificationCode(claims jwt.MapClaims) error {
	if o.Unique {
		if uic := o.GetUniqueIdentificationCode(claims[o.IdentityKey]); uic != "" {
			if uic != claims["uic"] {
				return ErrAuthorizeElsewhere
			}
		} else {
			return ErrInvalidToken
		}
	}

	return nil
}

// 保存唯一识别码
func (o *Jwt) SetUniqueIdentificationCode(claims jwt.MapClaims) {
	gcache.Set("jwt:"+o.Realm+":"+gconv.String(claims[o.IdentityKey]), claims["uic"], o.Timeout)
}

// 获取唯一识别码
func (o *Jwt) GetUniqueIdentificationCode(identity interface{}) string {
	if uic, err := gcache.Get("jwt:" + o.Realm + ":" + gconv.String(identity)); err != nil {
		return ""
	} else {
		return gconv.String(uic)
	}
}

// 删除唯一识别码
func (o *Jwt) DelUniqueIdentificationCode(identity interface{}) error {
	if _, err := gcache.Remove("jwt:" + o.Realm + ":" + gconv.String(identity)); err != nil {
		return ErrFailedTokenDestroy
	} else {
		return nil
	}
}
