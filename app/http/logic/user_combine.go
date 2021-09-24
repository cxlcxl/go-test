package logic

import "goskeleton/app/model"

// CombineUserParent 组合用户列表的父级账号等...
func CombineUserParent(users []*model.UsersModel) {
	if len(users) > 0 {
		parentIds := make([]int64, 0)
		for i, user := range users {
			if user.ParentId != 0 {
				parentIds = append(parentIds, user.ParentId)
			}
			// 用户类型
			users[i].UserTypeName = model.UserTypes[user.UserType]
		}
		if len(parentIds) > 0 {
			parents := model.CreateUserFactory("").GetsByIds(parentIds)
			parentUserMap := make(map[int64]string, len(parents))
			for _, parent := range parents {
				parentUserMap[parent.Id] = parent.UserName
			}
			for i, user := range users {
				if user.ParentId != 0 {
					if name, ok := parentUserMap[user.ParentId]; ok {
						users[i].ParentName = name
					}
				} else {
					users[i].ParentName = "无"
				}
			}
		}
	}
}
