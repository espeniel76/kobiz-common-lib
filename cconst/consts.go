package cconst

const DATE_FORMAT_KST = "2006-02-01 15:04:05"
const MAIN_LIST_SIZE = 5

type ErrorCode struct {
	Code   int
	Detail string
}

func GetErrorCode(code int) ErrorCode {
	if err, ok := errorCodes[code]; ok {
		return err
	}
	return ErrorCode{Code: code, Detail: "알 수 없는 에러"}
}

func GetErrorMessage(code int) string {
	err := GetErrorCode(code)
	return err.Detail
}

const (
	OK                             = 1000 // 성공
	SUCCESS                        = 1001 // 성공
	NONE                           = 1002 // 없음
	SYSTEM_UNKNOWN_ERROR           = 0    // 알 수 없는 에러가 발생하였습니다.
	BINDING_ERROR                  = 1    // Binding 에러
	EMPTY_TOKEN                    = 10   // 토큰(로그인정보)이 없습니다.
	INVALID_TOKEN                  = 11   // 토큰(로그인정보)이 유효하지 않습니다.
	EMPTY_APP_VERSION              = 21   // Application 버전 정보가 없습니다.
	UPGRADE_APP_VERSION            = 22   // Application 버전 업그레이드가 필요합니다.
	INVALID_USER_AGENT             = 23   // User-Agent 정보가 올바르지 않습니다.
	FILE_NOT_FOUND                 = 91   // 파일이 존재하지 않습니다.
	DATA_NOT_FOUND                 = 93   // 데이터가 존재하지 않습니다.
	USERID_DUPLICATE               = 105  // 중복된 아이디가 있습니다.
	EMAIL_DUPLICATE                = 106  // 중복된 이메일이 있습니다.
	FAIL_CONVERT_BASE64            = 107  // 파일 변환에 실패하였습니다.
	FAIL_SEND_MAIL                 = 108  // 이메일 발송에 실패하였습니다.
	FAIL_SEND_SMS                  = 109  // SMS 발송에 실패하였습니다.
	LOGIN_FAIL                     = 112  // 로그인이 실패하였습니다.
	EMPTY_USER                     = 113  // 입력하신 정보와 일치하는 회원정보가 없습니다. 다시 확인하여 주십시오.
	TOO_MUCH_REQ                   = 114  // 잠시 후에 다시 시도해주세요.
	WRONG_OTP                      = 115  // 인증번호가 올바르지 않습니다.
	NO_VALID_DATA                  = 116  // 유효하지 않은 데이터입니다.
	NO_VALID_PARAM                 = 117  // 유효하지 않은 파라미터입니다.
	INVALID_PWD                    = 118  // 유효한 패스워드가 아닙니다.
	EMPTY_ROLE                     = 119  // 권한이 충분하지 않습니다.
	INVALID_PATH_PARAM             = 120  // 유효하지 않은 경로 파라미터입니다.
	INSUFFICIENT_PARAM             = 501  // 필수 파라미터가 부족합니다.
	NOT_FOUND_USER                 = 502  // 사용자를 찾을 수 없습니다.
	WITHDRAWN_MEMBER               = 503  // 탈퇴한 회원입니다.
	WRONG_PASSWD                   = 504  // 잘못된 비밀번호입니다.
	AUTH_TIMEOUT                   = 505  // 인증 시간이 초과되었습니다.
	AUTH_FAIL                      = 506  // 인증에 실패하였습니다.
	ALREADY_EXIST_MEMBER           = 507  // 이미 존재하는 회원입니다.
	DUPLICATE_LOGIN                = 508  // 중복 로그인입니다.
	NO_LOGIN_INFORMATION           = 509  // 로그인 정보가 없습니다.
	WRONG_LOGIN_TYPE               = 510  // 잘못된 로그인 유형입니다.
	FAIL_ADD_INFO                  = 511  // 정보 추가에 실패하였습니다.
	FAIL_LIMIT_PASSWD              = 512  // 비밀번호 제한에 실패하였습니다.
	DISABLE_FIND_PASSWD_SNS        = 513  // SNS 비밀번호 찾기가 비활성화되었습니다.
	ACCOUNT_EMAIL                  = 514  // 이메일 계정입니다.
	ACCOUNT_SNS                    = 515  // SNS 계정입니다.
	DISABLE_LOGIN_ACCOUNT          = 516  // 로그인이 비활성화된 계정입니다.
	LOGIN_ERROR                    = 527  // 로그인 오류입니다.
	FAIL_SEND_TMP_PASSWD           = 518  // 임시 비밀번호 전송에 실패하였습니다.
	FAIL_JOIN_MEMBER               = 519  // 회원 가입에 실패하였습니다.
	FAIL_SEND_AUTHCODE             = 520  // 인증 코드 전송에 실패하였습니다.
	FAIL_UPDATE_AUTH               = 521  // 인증 정보 업데이트에 실패하였습니다.
	FAIL_UPDATE_PRIVATE            = 523  // 개인 정보 업데이트에 실패하였습니다.
	FAIL_UPDATE_PASSWD             = 524  // 비밀번호 업데이트에 실패하였습니다.
	FAIL_SET_SMARTALM              = 527  // 스마트 알람 설정에 실패하였습니다.
	FAIL_SET_UNLOCK_WAIT           = 529  // 잠금 해제 대기 설정에 실패하였습니다.
	FAIL_WITHDRAWN                 = 531  // 회원 탈퇴에 실패하였습니다.
	FAIL_SRCH_CONTENT              = 544  // 콘텐츠 검색에 실패하였습니다.
	FAIL_SAVE_FAVOR                = 548  // 즐겨찾기 저장에 실패하였습니다.
	FAIL_DEL_FAVOR                 = 549  // 즐겨찾기 삭제에 실패하였습니다.
	FAIL_SET_LIFEINFO              = 559  // 생활 정보 설정에 실패하였습니다.
	LOGOUT_PROC_ERROR              = 560  // 로그아웃 처리 중 오류가 발생하였습니다.
	FAIL_UPLOAD_USEINFO            = 563  // 사용 정보 업로드에 실패하였습니다.
	FAIL_CHK_PASSWD                = 565  // 비밀번호 확인에 실패하였습니다.
	FAIL_SRCH_SETINFO              = 566  // 설정 정보 검색에 실패하였습니다.
	FAIL_READ_S3                   = 567  // S3 읽기에 실패하였습니다.
	DISABLE_CHANGE_SAME_PASSWD     = 568  // 동일한 비밀번호로 변경할 수 없습니다.
	DISABLE_SAME_PASSWD_EMAIL      = 573  // 이메일과 동일한 비밀번호는 사용할 수 없습니다.
	DISABLE_SAME_PASSWD_PHONE      = 574  // 전화번호와 동일한 비밀번호는 사용할 수 없습니다.
	DUPLICATE_DATA_ERROR           = 575  // 중복된 데이터 오류입니다.
	SQL_GRAMMER_ERROR              = 576  // SQL 문법 오류입니다.
	DATA_INTEGRITY_VIOLATION_ERROR = 577  // 데이터 무결성 위반 오류입니다.
	DB_ERROR                       = 578  // 데이터베이스 오류입니다.
	NO_EXIST_SET_CONTENT           = 579  // 설정된 콘텐츠가 존재하지 않습니다.
	ALREADY_EXIST_REPLY            = 580  // 이미 답변이 존재합니다.
	NO_EXIST_REDIS_DATA            = 582  // Redis 데이터가 존재하지 않습니다.
	NO_VALID_DATE                  = 600  // 유효하지 않은 날짜입니다.
	FILE_UPLOAD_FAILED             = 601  // 파일 업로드에 실패하였습니다.
	DISABLE_CHANGE_SAME_PHONE      = 602  // 동일한 전화번호로 변경할 수 없습니다.
	NOW_LOGINED_ACCOUNT            = 702  // 현재 로그인된 계정입니다.
	ACCOUNT_ALREADY_EXIST          = 703  // 계정이 이미 존재합니다.
	NO_EXIST_TOKEN                 = 800  // 인증 토큰이 없습니다.
	INCORRECT_HTTP_METHOD          = 801  // 요청 타입이 맞지 않습니다.
	EXPIRED_TOKEN                  = 802  // 만료된 토큰입니다.
	CONTENT_TYPE_MISMATCH          = 803  // 컨텐츠 타입이 맞지 않습니다.
	JSON_SYNTAX_ERROR              = 804  // JSON 구문 오류
	NO_EXIST_CONTENT_TYPE          = 805  // Header Content-Type 이 존재하지 않습니다.
	NO_EXIST_USER_AGENT_SIMULATION = 808  // Header User-Agent가 존재하지 않습니다.
	NO_EXIST_API                   = 809  // 정의되지 않은 API
	INCORRECT_PASSWORD             = 810  // 비밀번호가 틀렸습니다.
	SIGN_ERROR_TOKEN               = 811  // 토큰 서명 에러
	ALREADY_EXIST_NICKNAME         = 819  // 이미 존재하는 닉네임입니다.
	FAIL_UPLOAD_S3                 = 820  // S3 파일 업로드 전송 실패
	NO_EXIST_ACCOUNT               = 821  // 계정이 존재하지 않습니다.
)

var errorCodes = map[int]ErrorCode{
	SYSTEM_UNKNOWN_ERROR:           {SYSTEM_UNKNOWN_ERROR, "알수없는 에러가 발생하였습니다."},
	BINDING_ERROR:                  {BINDING_ERROR, "Binding 에러 "},
	EMPTY_TOKEN:                    {EMPTY_TOKEN, "토큰(로그인정보)이 없습니다."},
	INVALID_TOKEN:                  {INVALID_TOKEN, "토큰(로그인정보)이 유효하지 않습니다."},
	EMPTY_APP_VERSION:              {EMPTY_APP_VERSION, "Application 버전 정보가 없습니다."},
	UPGRADE_APP_VERSION:            {UPGRADE_APP_VERSION, "Application 버전 업그레이드가 필요합니다."},
	INVALID_USER_AGENT:             {INVALID_USER_AGENT, "User-Agent 정보가 올바르지 않습니다."},
	FILE_NOT_FOUND:                 {FILE_NOT_FOUND, "파일이 존재하지 않습니다."},
	DATA_NOT_FOUND:                 {DATA_NOT_FOUND, "데이터가 존재하지 않습니다."},
	USERID_DUPLICATE:               {USERID_DUPLICATE, "중복된 아이디가 있습니다."},
	EMAIL_DUPLICATE:                {EMAIL_DUPLICATE, "중복된 이메일이 있습니다."},
	FAIL_CONVERT_BASE64:            {FAIL_CONVERT_BASE64, "파일 변환에 실패하였습니다."},
	FAIL_SEND_MAIL:                 {FAIL_SEND_MAIL, "이메일 발송에 실패하였습니다."},
	FAIL_SEND_SMS:                  {FAIL_SEND_SMS, "SMS 발송에 실패하였습니다."},
	LOGIN_FAIL:                     {LOGIN_FAIL, "로그인이 실패하였습니다."},
	EMPTY_USER:                     {EMPTY_USER, "입력하신 정보와 일치하는 회원정보가 없습니다.다시 확인하여 주십시오."},
	TOO_MUCH_REQ:                   {TOO_MUCH_REQ, "잠시후에 다시 시도해주세요."},
	WRONG_OTP:                      {WRONG_OTP, "인증번호가 올바르지 않습니다."},
	NO_VALID_DATA:                  {NO_VALID_DATA, "유효하지 않은 데이터입니다."},
	NO_VALID_PARAM:                 {NO_VALID_PARAM, "유효하지 않은 파라미터입니다."},
	INVALID_PWD:                    {INVALID_PWD, "유효한 패스워드가 아닙니다."},
	EMPTY_ROLE:                     {EMPTY_ROLE, "권한이 충분하지 않습니다."},
	INVALID_PATH_PARAM:             {INVALID_PATH_PARAM, "유효하지 않은 경로 파라미터입니다."},
	INSUFFICIENT_PARAM:             {INSUFFICIENT_PARAM, "필수 파라미터가 부족합니다."},
	NOT_FOUND_USER:                 {NOT_FOUND_USER, "사용자를 찾을 수 없습니다."},
	WITHDRAWN_MEMBER:               {WITHDRAWN_MEMBER, "탈퇴한 회원입니다."},
	WRONG_PASSWD:                   {WRONG_PASSWD, "잘못된 비밀번호입니다."},
	AUTH_TIMEOUT:                   {AUTH_TIMEOUT, "인증 시간이 초과되었습니다."},
	AUTH_FAIL:                      {AUTH_FAIL, "인증에 실패하였습니다."},
	ALREADY_EXIST_MEMBER:           {ALREADY_EXIST_MEMBER, "이미 존재하는 회원입니다."},
	DUPLICATE_LOGIN:                {DUPLICATE_LOGIN, "중복 로그인입니다."},
	NO_LOGIN_INFORMATION:           {NO_LOGIN_INFORMATION, "로그인 정보가 없습니다."},
	WRONG_LOGIN_TYPE:               {WRONG_LOGIN_TYPE, "잘못된 로그인 유형입니다."},
	FAIL_ADD_INFO:                  {FAIL_ADD_INFO, "정보 추가에 실패하였습니다."},
	FAIL_LIMIT_PASSWD:              {FAIL_LIMIT_PASSWD, "비밀번호 제한에 실패하였습니다."},
	DISABLE_FIND_PASSWD_SNS:        {DISABLE_FIND_PASSWD_SNS, "SNS 비밀번호 찾기가 비활성화되었습니다."},
	ACCOUNT_EMAIL:                  {ACCOUNT_EMAIL, "이메일 계정입니다."},
	ACCOUNT_SNS:                    {ACCOUNT_SNS, "SNS 계정입니다."},
	DISABLE_LOGIN_ACCOUNT:          {DISABLE_LOGIN_ACCOUNT, "로그인이 비활성화된 계정입니다."},
	LOGIN_ERROR:                    {LOGIN_ERROR, "로그인 오류입니다."},
	FAIL_SEND_TMP_PASSWD:           {FAIL_SEND_TMP_PASSWD, "임시 비밀번호 전송에 실패하였습니다."},
	FAIL_JOIN_MEMBER:               {FAIL_JOIN_MEMBER, "회원 가입에 실패하였습니다."},
	FAIL_SEND_AUTHCODE:             {FAIL_SEND_AUTHCODE, "인증 코드 전송에 실패하였습니다."},
	FAIL_UPDATE_AUTH:               {FAIL_UPDATE_AUTH, "인증 정보 업데이트에 실패하였습니다."},
	FAIL_UPDATE_PRIVATE:            {FAIL_UPDATE_PRIVATE, "개인 정보 업데이트에 실패하였습니다."},
	FAIL_UPDATE_PASSWD:             {FAIL_UPDATE_PASSWD, "비밀번호 업데이트에 실패하였습니다."},
	FAIL_SET_UNLOCK_WAIT:           {FAIL_SET_UNLOCK_WAIT, "잠금 해제 대기 설정에 실패하였습니다."},
	FAIL_WITHDRAWN:                 {FAIL_WITHDRAWN, "회원 탈퇴에 실패하였습니다."},
	FAIL_SRCH_CONTENT:              {FAIL_SRCH_CONTENT, "콘텐츠 검색에 실패하였습니다."},
	FAIL_SAVE_FAVOR:                {FAIL_SAVE_FAVOR, "즐겨찾기 저장에 실패하였습니다."},
	FAIL_DEL_FAVOR:                 {FAIL_DEL_FAVOR, "즐겨찾기 삭제에 실패하였습니다."},
	FAIL_SET_LIFEINFO:              {FAIL_SET_LIFEINFO, "생활 정보 설정에 실패하였습니다."},
	LOGOUT_PROC_ERROR:              {LOGOUT_PROC_ERROR, "로그아웃 처리 중 오류가 발생하였습니다."},
	FAIL_UPLOAD_USEINFO:            {FAIL_UPLOAD_USEINFO, "사용 정보 업로드에 실패하였습니다."},
	FAIL_CHK_PASSWD:                {FAIL_CHK_PASSWD, "비밀번호 확인에 실패하였습니다."},
	FAIL_SRCH_SETINFO:              {FAIL_SRCH_SETINFO, "설정 정보 검색에 실패하였습니다."},
	FAIL_READ_S3:                   {FAIL_READ_S3, "S3 읽기에 실패하였습니다."},
	DISABLE_CHANGE_SAME_PASSWD:     {DISABLE_CHANGE_SAME_PASSWD, "동일한 비밀번호로 변경할 수 없습니다."},
	DISABLE_SAME_PASSWD_EMAIL:      {DISABLE_SAME_PASSWD_EMAIL, "이메일과 동일한 비밀번호는 사용할 수 없습니다."},
	DISABLE_SAME_PASSWD_PHONE:      {DISABLE_SAME_PASSWD_PHONE, "전화번호와 동일한 비밀번호는 사용할 수 없습니다."},
	DUPLICATE_DATA_ERROR:           {DUPLICATE_DATA_ERROR, "중복된 데이터 오류입니다."},
	SQL_GRAMMER_ERROR:              {SQL_GRAMMER_ERROR, "SQL 문법 오류입니다."},
	DATA_INTEGRITY_VIOLATION_ERROR: {DATA_INTEGRITY_VIOLATION_ERROR, "데이터 무결성 위반 오류입니다."},
	DB_ERROR:                       {DB_ERROR, "데이터베이스 오류입니다."},
	NO_EXIST_SET_CONTENT:           {NO_EXIST_SET_CONTENT, "설정된 콘텐츠가 존재하지 않습니다."},
	ALREADY_EXIST_REPLY:            {ALREADY_EXIST_REPLY, "이미 답변이 존재합니다."},
	NO_EXIST_REDIS_DATA:            {NO_EXIST_REDIS_DATA, "Redis 데이터가 존재하지 않습니다."},
	NO_VALID_DATE:                  {NO_VALID_DATE, "유효하지 않은 날짜입니다."},
	FILE_UPLOAD_FAILED:             {FILE_UPLOAD_FAILED, "파일 업로드에 실패하였습니다."},
	DISABLE_CHANGE_SAME_PHONE:      {DISABLE_CHANGE_SAME_PHONE, "동일한 전화번호로 변경할 수 없습니다."},
	NOW_LOGINED_ACCOUNT:            {NOW_LOGINED_ACCOUNT, "현재 로그인된 계정입니다."},
	ACCOUNT_ALREADY_EXIST:          {ACCOUNT_ALREADY_EXIST, "계정이 이미 존재합니다."},
	NO_EXIST_TOKEN:                 {NO_EXIST_TOKEN, "인증 토큰이 없습니다."},
	INCORRECT_HTTP_METHOD:          {INCORRECT_HTTP_METHOD, "요청타입이 맞지 않습니다."},
	EXPIRED_TOKEN:                  {EXPIRED_TOKEN, "만료된 토큰 입니다."},
	CONTENT_TYPE_MISMATCH:          {CONTENT_TYPE_MISMATCH, "컨텐츠 타입이 맞지 않습니다."},
	JSON_SYNTAX_ERROR:              {JSON_SYNTAX_ERROR, "JSON 구문 오류."},
	NO_EXIST_CONTENT_TYPE:          {NO_EXIST_CONTENT_TYPE, "Header Content-Type 이 존재하지 않습니다."},
	NO_EXIST_USER_AGENT_SIMULATION: {NO_EXIST_USER_AGENT_SIMULATION, "Header User-Agent가 존재하지 않습니다."},
	NO_EXIST_API:                   {NO_EXIST_API, "정의되지 않은 API"},
	INCORRECT_PASSWORD:             {INCORRECT_PASSWORD, "비밀번호가 틀렸습니다."},
	SIGN_ERROR_TOKEN:               {SIGN_ERROR_TOKEN, "토큰 서명 에러"},
	ALREADY_EXIST_NICKNAME:         {ALREADY_EXIST_NICKNAME, "이미 존재하는 닉네임 입니다."},
	FAIL_UPLOAD_S3:                 {FAIL_UPLOAD_S3, "S3파일 업로드 전송 실패"},
	NO_EXIST_ACCOUNT:               {NO_EXIST_ACCOUNT, "계정이 존재하지 않습니다"},
}
