package utils

const (
    AREA_INFO_KEY = "area_info"
)

func GetNameKey(sessionId string) string {
    return sessionId + "name"
}

func GetUserIdKey(sessionId string) string {
    return sessionId + "user_id"
}

func GetMobileKey(sessionId string) string {
    return sessionId + "mobile"
}
