package constant

// gva 常量文件，用于定义一些常量 , 如果修改此处，请自行修改前端中的/src/const/index.js 中对应的
const TOKEN_NAME = "x-token"                 // token ，此处是存储在localStorage中的token名称
const USER_INFO = "x-user"                   // 用户信息
const REQUEST_HEAER_TOKEN_NAME = TOKEN_NAME  // 请求头 token 名称 , 此处是后端可以从请求头中获取token的名称
const REQUEST_USER_ID = "x-user-id"          // 请求头 用户ID 名称
const REQUSET_COOKIE_TOKEN_NAME = TOKEN_NAME // cookie token 名称, 修改此处，请同时修改后端 cookie 的 name
const NEW_TOKEN_NAME = "new-token"           // 请求头中新token的名称
const OS_TYPE_NAME = "os-type"               // 请求头中操作系统类型的名称, 同时存储在localStorage中
const OS_TYPE_WIN = "WIN"                    // windows 操作系统
const OS_TYPE_MAC = "MAC"                    // mac 操作系统
const TOP_ACTIVE = "topActive"               // 顶部菜单激活项
const LOCAL_HISTORY = "historys"             // 存储的菜单
const TAB_ACTIVE_NAME = "activeValue"        // 当前活跃的菜单
const JWT_CLAIMS = "claims"                  // jwt 中的 claims
const NEW_TOKEN_EXPIRES_AT = "new-expires-at"
