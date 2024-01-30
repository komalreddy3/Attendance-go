package userAdapter

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
)

func Userrec(students []userModels.User) []userServiceBean.CustomUserInfo {
	var customUserInfos []userServiceBean.CustomUserInfo

	for _, student := range students {
		customUserInfo := userServiceBean.CustomUserInfo{
			Username: student.Username,
		}

		// Append the custom user info to the new slice
		customUserInfos = append(customUserInfos, customUserInfo)
	}
	return customUserInfos
}
