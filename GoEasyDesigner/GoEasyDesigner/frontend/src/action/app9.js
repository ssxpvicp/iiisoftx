import {useCounterStore} from '@/stores/counter'
import {onMounted, ref} from "vue";
import {
    E保存,
    E保存件对话框,
    E停止命令,
    E打开文件对话框,
    E检查更新,
    E读入文件,
    E运行命令
} from "../../wailsjs/go/main/App";
import {取父目录, 生成辅助代码} from "@/public";
import {ElMessage} from "element-plus";
import {BrowserOpenURL, EventsOn} from "../../wailsjs/runtime";

export const appAction = {};
let store = {}; // 我想在这里作类型标注 useCounterStore()怎么处理怎么写
appAction.init = function () {
    console.log("abc")
    appAction.store = useCounterStore()
    store = useCounterStore()
}

appAction.新建 = function () {
    function 创建窗口() {
        return {
            "id": "1",
            "名称": "窗口",
            "组件名称": "窗口",
            "标题": "窗口",
            "top": "0",
            "left": "0",
            "width": "500",
            "height": "400",
            "background": "rgba(0, 0, 0, 0.05)",
            "禁止放置": false,
            "禁止拖动": true,
            "禁止": false,
            "可视": true,
            "层级": 0,
            "子组件": []
        }
    }

    appAction.store.list = [创建窗口()]
    appAction.store.indexMap = {}
}
appAction.打开 = function () {
    if (appAction.store.客户端模式 == false) {
        const input = document.createElement('input')
        input.type = 'file'
        input.accept = '.json'
        input.onchange = e => {
            const file = e.target.files[0]
            const reader = new FileReader()
            reader.readAsText(file)
            reader.onload = () => {
                const data = reader.result
                console.log(data)
                // 初始化界面(data)
                store.list = JSON.parse(data)
            }
        }
        input.click()
        return
    }
    console.log("打开")
    E打开文件对话框().then((文件路径) => {
        if (文件路径 == "") {
            return
        }
        appAction._打开文件加载界面(文件路径)
    })
}

appAction._打开文件加载界面 = function (filepath) {
    store.项目信息.设计文件路径 = filepath
    store.项目信息.窗口事件文件路径 = 取父目录(filepath) + "/窗口事件.js"
    store.项目信息.辅助代码文件路径 = 取父目录(filepath) + "/辅助代码.js"
    store.项目信息.项目管理目录 = 取父目录(filepath)
    store.项目信息.项目根目录 = 取父目录(取父目录(取父目录(取父目录(filepath))))


    console.log("设计文件路径", store.项目信息.设计文件路径)
    console.log("窗口事件文件路径", store.项目信息.窗口事件文件路径)
    E读入文件(store.项目信息.设计文件路径).then((文件内容) => {
        console.log(文件内容)
        // 初始化界面(文件内容)
        store.list = JSON.parse(文件内容)
    })
    // E读入文件(store.项目信息.窗口事件文件路径).then((res) => {
    //   console.log(res)
    //   code.value = res
    // })
    store.项目管理刷新()
}

// 使用 $refs 来引用滚动容器
const scrollContainer = ref(null);

try {
    EventsOn("运行命令", function (data) {
        console.log("运行命令", data)
        store.调试信息 = store.调试信息 + "<br / >" + data
        scrollContainer.value.scrollTop = scrollContainer.value.scrollHeight;
        if (data == "命令已完成") {
            store.运行按钮文本 = "运行"
            store.编译按钮文本 = "编译"
        }
    })
} catch (e) {
    console.log('非客户端模式')
}

function 键盘按下(event, index) {
    console.log("键盘按下", event.key, index)
    if (event.key == 'Delete') {

    }
}


appAction.保存设计文件 = function () {
    let njson = JSON.stringify(store.list, null, 2)
    console.log("???", store)
    let 辅助代码 = 生成辅助代码(store.list[0].子组件)


    if (store.客户端模式 == false) {
        //浏览器打开就发起保存
        const blob = new Blob([njson], {type: 'application/json'})
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = '设计文件.json'
        link.click()

        const blob2 = new Blob([辅助代码], {type: 'application/json'})
        const link2 = document.createElement('a')
        link2.href = URL.createObjectURL(blob2)
        link2.download = '辅助代码.js'
        link2.click()
        return;
    }

    // 客户端直接保存
    function _保存(p, d) {
        d = String(d)
        p = String(p)
        E保存(p, d)
    }

    if (store.项目信息.设计文件路径 == "") {
        E保存件对话框().then((res) => {
            if (res == "") {
                ElMessage({
                    message: "未选择文件",
                    type: 'success',
                    duration: 3000, // 设置显示时间为5秒，单位为毫秒
                });
                return
            }

            store.项目信息.设计文件路径 = res
            store.项目信息.窗口事件文件路径 = 取父目录(res) + "/窗口事件.js"
            store.项目信息.辅助代码文件路径 = 取父目录(res) + "/辅助代码.js"
            store.项目信息.项目管理目录 = 取父目录(res)
            store.项目信息.项目根目录 = 取父目录(取父目录(取父目录(取父目录(res))))

            store.项目管理刷新()

            console.log("窗口事件文件路径", store.项目信息.窗口事件文件路径)
            if (store.代码编辑器内容 !== "") {
                _保存(store.项目信息.窗口事件文件路径, store.代码编辑器内容)
                _保存(store.项目信息.辅助代码文件路径, 辅助代码)
            }
            _保存(store.项目信息.设计文件路径, njson)
        })
        return
    }
    if (store.代码编辑器内容 !== "") {
        _保存(store.项目信息.窗口事件文件路径, store.代码编辑器内容)
        _保存(store.项目信息.辅助代码文件路径, 辅助代码)
    }
    _保存(store.项目信息.设计文件路径, njson)

}

appAction.运行 = function () {
    if (store.客户端模式 == false) {
        //弹出提示
        ElMessage({
            message: "当前为浏览器模式 不能运行 请自行在项目根目录运行 wails dev",
            type: 'success',
            duration: 3000, // 设置显示时间为5秒，单位为毫秒
        });
        return
    }
    if (store.项目信息.项目根目录 == "") {
        ElMessage({
            message: "请先保存",
            type: 'success',
            duration: 3000, // 设置显示时间为5秒，单位为毫秒
        });
        return
    }

    if (store.运行按钮文本 == '运行') {
        store.运行按钮文本 = '停止'
        store.调试信息 = "运行中 ..."
        store.选择夹_底部现行选中项 = "1"
        E运行命令(store.项目信息.项目根目录, "wails dev")
    } else {
        store.调试信息 = "已停止 ..."
        store.运行按钮文本 = '运行'
        E停止命令()
    }

}
appAction.编译 = function () {
    if (store.客户端模式 == false) {
        //弹出提示
        ElMessage({
            message: "当前为浏览器模式 不能编译 请自行在项目根目录运行 wails build",
            type: 'success',
            duration: 3000, // 设置显示时间为5秒，单位为毫秒
        });
        return
    }
    if (store.项目信息.项目根目录 == "") {
        ElMessage({
            message: "请先保存",
            type: 'success',
            duration: 3000, // 设置显示时间为5秒，单位为毫秒
        });
        return
    }
    if (store.编译按钮文本 == '编译') {
        store.编译按钮文本 = '停止'
        store.调试信息 = "编译中 ..."
        store.选择夹_底部现行选中项 = "1"
        E运行命令(store.项目信息.项目根目录, "wails build")
    } else {
        store.调试信息 = "已停止 ..."
        store.编译按钮文本 = '编译'
        E停止命令()
    }

}

appAction.帮助 = function () {
    if (store.客户端模式) {
        BrowserOpenURL("https://github.com/duolabmeng6/GoEasyDesigner")
    } else {
        //浏览器打开新页面
        window.open("https://github.com/duolabmeng6/GoEasyDesigner")
    }
}

appAction.下载客户端 = function () {
    window.open("https://github.com/duolabmeng6/GoEasyDesigner/releases")
}


appAction.运行环境检测 = function () {
    if (store.客户端模式 == false) {
        //弹出提示
        ElMessage({
            message: "当前为浏览器模式 不能运行 请自行在项目根目录运行 wails doctor",
            type: 'success',
            duration: 3000, // 设置显示时间为5秒，单位为毫秒
        });
        return
    }
    store.选择夹_底部现行选中项 = "1"
    store.调试信息 = "运行环境检测 ..."
    let 结果;
    结果 = E运行命令(store.项目信息.项目根目录, "wails doctor")
    console.log("结果")
    // 检查node是否安装

}

appAction.检查更新 = function () {
    E检查更新()
}

export default {
    appAction,
}