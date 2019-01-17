package utils

const (
    AREA_INFO_KEY = "area_info"
    HOME_INDEX_DATA_KEY = "home_index_data"
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

