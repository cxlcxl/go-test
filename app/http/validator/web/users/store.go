package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	validator_web "goskeleton/app/http/validator/web"
	"goskeleton/app/utils/response"
	"regexp"
)

type Store struct {
	BaseField
	// 表单参数验证结构体支持匿名结构体嵌套、以及匿名结构体与普通字段组合
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	RealName string `form:"real_name" json:"real_name" binding:"max=30"`
	Remark   string `form:"remark" json:"remark"`
	Avatar   string `json:"avatar"`
}

// 特别注意: 表单参数验证器结构体的函数，绝对不能绑定在指针上
// 我们这部分代码项目启动后会加载到容器，如果绑定在指针，一次请求之后，会造成容器中的代码段被污染

func (s Store) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&s); err != nil {
		response.ErrorParam(context, validator_web.Translate(err))
		return
	}

	if errMsg := s.customCheck(); errMsg != "" {
		response.ErrorParam(context, errMsg)
		return
	}

	//2.继续验证具有中国特色的参数，例如 身份证号码等，基本语法校验了长度18位，然后可以自行编写正则表达式等更进一步验证每一部分组成
	// r.CardNo  获取值继续校验，这里省略.....

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(s, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Store(extraAddBindDataContext)
	}
}

func (s Store) customCheck() string {
	if len(s.Email) == 0 && len(s.Phone) == 0 {
		return "邮箱和手机号至少填写一项"
	} else {
		if len(s.Phone) > 0 {
			if ok, _ := regexp.MatchString(consts.PhoneRule, s.Phone); !ok {
				return "手机号填写格式有误"
			}
		}
		if len(s.Email) > 0 {
			if ok, _ := regexp.MatchString(consts.EmailRule, s.Email); !ok {
				return "邮箱填写格式有误"
			}
		}
		return ""
	}
}
