package tool

const (
	UserCheckPass    = 1
	UserCheckNotPass = -1
	UserCheckIng     = 0
	OpenStatus       = 1
	CloseStatus      = 1
	SuperAdmin       = 0
	OperatorUser     = 1
	AdsUser          = 2
	DeveloperUser    = 3
	GDTUser          = 4
	SPMUser          = 5
	IDTUser          = 6

	DeleteFlag    = 1
	NotDeleteFlag = 0

	AccountTaskMaxExpireTime = 2147483647
	OneDayForSeconds         = 86400

	AndroidPlatform = 1
	IosPlatform     = 2
	H5Platform      = 3

	ScreenCross    = 1
	ScreenVertical = 2

	DefaultCpcVideoCtr  = 0.05
	DefaultCpcPicCtr    = 0.1
	DefaultCpcCustomCtr = 0.02
	DefaultCpcSashCtr   = 0.08
	DefaultCpcEmbedCtr  = 0.02

	RateBase           = 10000
	DspDefaultBidPrice = 20.00

	IntegrationAdType = 1 // 广告类型
	ChannelAdType     = 2
	DspAdType         = 3

	VideoAdSubType                = 1  // 视频广告
	PicAdSubType                  = 2  // 插图广告
	CustomAdSubType               = 3  // 自定义广告
	SplashAdSubType               = 4  // 开屏广告
	EmbedAdSubType                = 5  // 原生1.0广告
	InteractiveAdSubType          = 6  // 互动广告
	BannerAdSubType               = 7  // banner广告
	FeedAdSubType                 = 8  // 信息流广告
	InteractiveIncentiveAdSubType = 9  // 互动激励广告
	Embed2AdSubType               = 10 // 原生自渲染广告
	EmbedTemplateAdSubType        = 11 // 原生模板广告
	FocusCustomAdSubType          = 31 // 交叉推广 精品橱窗－焦点图
	WallCustomAdSubType           = 32 // 交叉推广 精品橱窗－应用墙
	NativeCustomAdSubType         = 33 // 交叉推广 原生Banner
	SingleEmbedAdSubType          = 51 // 原生广告 横版单图
	CombinationEmbedAdSubType     = 52 // 原生广告 组图
	SingleVerticalEmbedAdSubType  = 53 // 原生广告 竖版单图
	SingleBannerAdSubType         = 71 // banner广告 单图

	ChargeTypeCpm = 1 // 计费类型
	ChargeTypeCpc = 2 // 计费类型
	ChargeTypeCpa = 3 // 计费类型

	EventTypeView   = 5  // 计费事件类型
	EventTypeClick  = 6  // 计费事件类型
	EventTypeActive = 45 // 计费事件类型

	StatMOBGI   = 1 // 统计版本
	StatHOUSEAD = 2 // 统计版本
	StatADX     = 3 // 统计版本

	VideoAds                = 1  // 视频广告
	PicAds                  = 2  // 插图广告
	CustomAds               = 3  // 自定义广告
	SplashAds               = 4  // 开屏广告
	EmbedAds                = 5  // 原生1.0广告
	BannerAds               = 7  // banner广告
	FeedAds                 = 8  // 信息流广告
	InteractiveIncentiveAds = 9  // 互动激励广告
	Embed2Ads               = 10 // 原生自渲染广告
	EmbedTemplateAds        = 11 // 原生模板广告
	CustomFocus             = 31 // 交叉推广 精品橱窗－焦点图
	CustomWall              = 32 // 交叉推广 精品橱窗－应用墙
	CustomNative            = 33 // 交叉推广 原生Banner

	FieldMiss    = 50001 // 接口异常编码
	FieldInvalid = 50002

	StatusActive    = 1 // 配置状态码
	StatusNotActive = 0

	EtoronDspId   = "Etoron_DSP"   // dsp的编号id
	DomobDspId    = "Domob_DSP"    // dsp的编号id
	HouseAdDspId  = "Housead_DSP"  // dsp的编号id
	UniplayDspId  = "Uniplay_DSP"  // dsp的编号id
	SmaatoDspId   = "Smaato_DSP"   // dsp的编号id
	ToutiaoDspId  = "Toutiao_DSP"  // dsp的编号id
	InmobiDspId   = "Inmobi_DSP"   // dsp的编号id
	OperaDspId    = "Opera_DSP"    // dsp的编号id
	ZhiziyunDspId = "Zhiziyun_DSP" // dsp的编号id
	AdinDspId     = "Adin_DSP"     // dsp的编号id
	YomobDspId    = "Yomob_DSP"    // dsp的编号id
	BulemobiDspId = "Bulemobi_DSP" // dsp的编号id
	IflytekDspId  = "Iflytek_DSP"  // dsp的编号id
	AdviewDspId   = "Adview_DSP"   // dsp的编号id
	AiclkDspId    = "Aiclk_DSP"    // dsp的编号id
	WangmaiDspId  = "Wangmai_DSP"  // dsp的编号id
	AdroiDspId    = "Adroi_DSP"    // dsp的编号id
	MobvistaDspId = "Mobvista_DSP" // dsp的编号id
	SogouDspId    = "Sougou_DSP"   // dsp的编号id
	AdtalosDspId  = "Adtalos_DSP"  // dsp的编号id
)

var (
	ConfTypes = map[int]string{1: "全局配置", 2: "定向配置"}
	AppTypes  = map[int]string{3: "联盟流量", 1: "自有流量", 4: "小游戏流量", 5: "休闲游戏变现"}
	UserType  = map[int]string{SuperAdmin: "超级管理员", OperatorUser: "运营用户", AdsUser: "广告主用户", DeveloperUser: "开发者用户", GDTUser: "广点通用户", SPMUser: "投放用户", IDTUser: "IDT用户"}
	// UserCheckType 用户审核类型
	UserCheckType = map[int]string{UserCheckPass: "已审核", UserCheckIng: "审核中", UserCheckNotPass: "审核不通过"}
	MPlatformDesc = map[int]string{IosPlatform: "IOS", AndroidPlatform: "Andriod"}
	MAdType       = map[int]string{IntegrationAdType: "聚合广告", ChannelAdType: "渠道广告", DspAdType: "DSP广告"}
	// MAdSubType 广告子类型
	MAdSubType = map[int]string{VideoAdSubType: "视频广告", PicAdSubType: "插页广告", SplashAdSubType: "开屏广告", EmbedAdSubType: "原生1.0广告", BannerAdSubType: "banner广告", FeedAdSubType: "信息流广告", Embed2AdSubType: "原生自渲染广告", EmbedTemplateAdSubType: "原生模板广告"}
	// MAdPosType 广告位对应的广告类型
	MAdPosType = map[int]string{VideoAdSubType: "VIDEO_INTERGRATION", PicAdSubType: "PIC_INTERGRATION", SplashAdSubType: "SPLASH_INTERGRATION", EmbedAdSubType: "ENBED_INTERGRATION", InteractiveAdSubType: "INTERATIVE_AD", BannerAdSubType: "BANNER_AD", FeedAdSubType: "FEED_AD", Embed2AdSubType: "ENBED2_INTERGRATION", EmbedTemplateAdSubType: "ENBED_TEMPLATE_INTERGRATION"}
	// MAdSubTypeDesc 广告子类型描述
	MAdSubTypeDesc = map[int]string{VideoAdSubType: "video", PicAdSubType: "pic", SplashAdSubType: "splash", EmbedAdSubType: "enbed", BannerAdSubType: "banner", FeedAdSubType: "feed", Embed2AdSubType: "enbed2", EmbedTemplateAdSubType: "enbed_template"}
	// MCustomSubType 交叉推广子类型
	MCustomSubType = map[int]string{FocusCustomAdSubType: "精品橱窗－焦点图", WallCustomAdSubType: "精品橱窗－应用墙", NativeCustomAdSubType: "原生Banner"}
	// MEmbedSubType 原生1.0广告子类型
	MEmbedSubType = map[int]string{SingleEmbedAdSubType: "单图", CombinationEmbedAdSubType: "组图"}
	// MNewEmbedSubType 原生2.0广告子类型
	MNewEmbedSubType = map[int]string{SingleEmbedAdSubType: "横版单图", CombinationEmbedAdSubType: "横版三小图", SingleVerticalEmbedAdSubType: "竖版单图"}
	// MTemplateSubType 原生模板广告子类型
	MTemplateSubType = map[int]string{90: "上图下文(图片尺寸1280×720)", 91: "上文下图(图片尺寸1280×720)", 92: "左图右文(图片尺寸1200×800)", 93: "左文右图(图片尺寸1200×800)", 94: "双图双文(大图尺寸1280×720)", 95: "纯图片(图片尺寸1280×720)", 96: "纯图片(图片尺寸800×1200)"}
	// MEmbedSize 原生1.0广告尺寸列表
	MEmbedSize = map[string]string{"16:9": "1280*720", "2:3": "640*960", "3:2": "960*640", "32:5": "640*100", "1:1": "1200*1200", "1200:627": "1200*627"}
	// MBannerSize Banner广告尺寸列表
	MBannerSize = []string{"320*50", "640*100"}
	// MAdPosTypeName 广告位类型对应的名称
	MAdPosTypeName = map[string]string{"VIDEO_INTERGRATION": "视频广告", "PIC_INTERGRATION": "插屏广告", "SPLASH_INTERGRATION": "开屏广告", "ENBED_INTERGRATION": "原生1.0广告", "INTERATIVE_AD": "互动广告", "BANNER_AD": "banner广告", "FEED_AD": "信息流广告", "ENBED2_INTERGRATION": "原生自渲染广告", "ENBED_TEMPLATE_INTERGRATION": "原生模板广告"}
	// MNewEmbedSize 原生新广告尺寸列表
	MNewEmbedSize = []map[string]map[string]string{
		{"51": {"16:9": "1280*720"}},
		{"52": {"3:2": "480*320"}},
		{"53": {"9:16": "1080*1920"}},
	}
)
