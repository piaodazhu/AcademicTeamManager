<template>
    <a-layout has-sider>
        <a-layout-sider class="layout-sider" width="150">
            <div class="logo">
                <div><img src="../assets/logo.svg"
                        style="width: 32px;height: 32px;filter: drop-shadow(2px 2px 3px #1283FF);" /></div>
                <div v-if="collapsed == false" class="title"><b style="color: #1283FF;">ATM</b></div>
            </div>
            <a-menu style="border-right: none;width: 149px;" v-model:selectedKeys="selectedKeys" mode="inline">
                <a-menu-item :key="item.key" v-for="item in menuItem">
                    <router-link :to="item.to" @click="store.selectedKeys = item.key" replace>
                        <component :is="item.icon" />
                        <span>{{ item.name }}</span>
                    </router-link>
                </a-menu-item>
            </a-menu>
        </a-layout-sider>
        <a-layout :style="{ marginLeft: '150px', background: '#FAFAFA' }">
            <a-layout-header class="header">
                <div style="display: flex;align-items: center;justify-items: center;margin-right: 10px;">
                    <QuestionCircleFilled @click="toDocs"
                        style="color: #909399;font-size: 18px;margin: 2px 15px 0 0;cursor: pointer;" />
                    <a-popover placement="bottomRight" :overlayStyle="{ width: '320px' }" trigger="click">
                        <template #content>
                            <div style="max-height: 250px;overflow-y: scroll;">
                                <a-list item-layout="horizontal" :data-source="data.noticeList" size="small">
                                    <template #renderItem="{ item }">
                                        <a-list-item style="cursor: pointer;" @click="onReadNotice(item.id)">
                                            <div style="display: inline-flex;align-items: center;">
                                                <a-avatar shape="square" :size="20" v-if="item.status == 1">
                                                    <template #icon>
                                                        <BellFilled style="font-size: 18px;" />
                                                    </template>
                                                </a-avatar>
                                                <a-avatar shape="square"
                                                    style="color: #476FFF; background-color: #ccd6fa" :size="20" v-else>
                                                    <template #icon>
                                                        <BellFilled style="font-size: 18px;" />
                                                    </template>
                                                </a-avatar>
                                                <div v-if="item.status == 1" style="color: #717171;">
                                                    &nbsp;&nbsp;&nbsp;{{ item.content }}</div>
                                                <div v-else>&nbsp;&nbsp;&nbsp;{{ item.content }}</div>
                                            </div>
                                            <template #actions>
                                                <span v-if="item.status == 2" style="color: black;">{{
                                                    formatDate(item.created)
                                                }}</span>
                                                <span v-else>{{ formatDate(item.created) }}</span>
                                            </template>
                                        </a-list-item>
                                    </template>
                                </a-list>
                            </div>
                            <div style="margin-top: 10px;display: flex;align-items: center;justify-content: center;">
                                <a-button v-if="data.noticeList.length != 0" @click="onDeleteNotice" type="primary"
                                    style="width: 92%;" shape="round">???????????? {{ data.noticeList.length }} ?????????</a-button>
                            </div>
                        </template>
                        <a-badge :count="data.noticeCount">
                            <BellFilled style="color: #909399; font-size: 20px;cursor: pointer;" @click="onNotice" />
                        </a-badge>
                    </a-popover>
                    <a-dropdown :trigger="['click']">
                        <a-avatar @click="onUserAvatar" class="avatar" :size="28">U</a-avatar>
                        <template #overlay>
                            <a-menu style="width: 120px;">
                                <a-menu-item @click="visible = true">
                                    <ExclamationCircleOutlined /> ????????????
                                </a-menu-item>
                                <a-menu-item @click="onLogout">
                                    <LogoutOutlined /> ????????????
                                </a-menu-item>
                            </a-menu>
                        </template>
                    </a-dropdown>
                </div>
                <!-- ????????????????????? -->
                <a-modal v-model:visible="visible" title="????????????" @ok="onConfirm" @cancel="onCancel" cancelText="??????"
                    okText="??????" width="400px" :centered="true">
                    <a-alert message="?????????????????????????????????????????????????????????" type="warning" show-icon /><br />
                    <a-form :model="user" layout="vertical" @finish="onSubmit" :rules="rules">
                        <a-form-item name="email">
                            <a-input v-model:value="user.email" placeholder="??????" disabled />
                        </a-form-item>
                        <a-form-item name="code">
                            <a-input v-model:value="user.code" style="width: 55%;" placeholder="?????????" />
                            <a-button @click="onGetCode" style="width: 40%;float: right;" :loading="loading"
                                :disabled="disabled">
                                {{ buttonText }}</a-button>
                        </a-form-item>
                    </a-form>
                </a-modal>
            </a-layout-header>
            <a-layout-content :style="{ margin: '10px', background: '#fff', overflow: 'initial', borderRadius: '5px' }">
                <transition name="fade">
                    <router-view v-slot="{ Component }">
                        <component :is="Component" />
                    </router-view>
                </transition>
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>

<script setup>
import { reactive, ref, onBeforeMount } from 'vue';
import { useRouter } from 'vue-router'
import { useStore } from '../store/index';
import { message } from 'ant-design-vue';
import { getUserInfo, getVerifyCode, userDelete } from '../api/user';
import { updateNotice, getNoticeCount, getNoticeList, deleteNotice } from '../api/notice';
import { DashboardOutlined, IdcardOutlined, RocketOutlined, FileDoneOutlined, ProfileOutlined, CrownOutlined } from '@ant-design/icons-vue';
import { QuestionCircleFilled, BellFilled, ExclamationCircleOutlined, LogoutOutlined } from '@ant-design/icons-vue';
import moment from 'moment'

// ????????????
const menuItem = reactive([{
    key: "dashboard",
    to: "/dashboard",
    icon: DashboardOutlined,
    name: "-?????????-"
}, {
    key: "Student",
    to: "/Student",
    icon: IdcardOutlined,
    name: "????????????"
}, {
    key: "project",
    to: "/project",
    icon: RocketOutlined,
    name: "????????????"
}, {
    key: "output",
    to: "/output",
    icon: FileDoneOutlined,
    name: "????????????"
}, {
    key: "config",
    to: "/config",
    icon: ProfileOutlined,
    name: "????????????"
}, 
// {
//     key: "subscribe",
//     to: "/subscribe",
//     icon: CrownOutlined,
//     name: "????????????"
// }
])

// ????????????
const rules = {
    email: [{
        required: true,
        message: '???????????????!',
        trigger: 'blur',
    }, {
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z.]{2,8})$/,
        message: '?????????????????????',
        trigger: 'blur',
    }],
    code: [{ required: true, message: '??????????????????!' }],
};

const store = useStore();

const selectedKeys = ref([store.selectedKeys])
const collapsed = ref(false)

const data = reactive({
    noticeCount: 0,
    noticeList: []
})

store.$subscribe((mutation, state) => {
    selectedKeys.value = [state.selectedKeys]
})

const router = useRouter()

const user = reactive({
    name: undefined,
    email: undefined,
    verison: undefined,
    code: undefined,
    versionText: undefined
})

const visible = ref(false)
const visibleLogo = ref(false)
const loading = ref(false)
const disabled = ref(false)
const buttonText = ref('???????????????')

// ???????????????
onBeforeMount(() => {
    store.selectedKeys = 'dashboard'
    router.push('dashboard')
    // noticeCount()
})

// ??????????????????
const onUserAvatar = () => {
    getUserInfo().then((res) => {
        if (res.data.code == 0) {
            user.name = res.data.data.name
            user.email = res.data.data.email
            user.version = res.data.data.version
        }
    })
}

// ?????????????????????
const toDocs = () => {
    // window.open("https://docs.zocrm.cloud")
    window.open("http://127.0.0.1:5173")
}

// ?????????????????????
const onGetCode = () => {
    if (user.email == '') {
        message.warn('??????????????????')
        return
    }
    loading.value = true
    let param = {
        email: user.email
    }
    getVerifyCode(param).then((res) => {
        if (res.data.code == 0) {
            loading.value = false
            disabled.value = true
            buttonText.value = '??????????????????'
        }
        if (res.data.code == 10004) {
            loading.value = false
            message.error('?????????????????????')
        }
    })
}

// ????????????????????????
const onConfirm = () => {
    let param = {
        email: user.email,
        code: user.code
    }
    userDelete(param).then((res) => {
        if (res.data.code == 0) {
            router.push('/')
            message.success('???????????????')
        }
    })
}

// ???????????????
const formatDate = (timeStamp) => {
    let now = (Date.parse(new Date())) / 1000
    if (now - timeStamp < 60) {
        return "??????"
    }
    if ((new Date().getDate()) == (new Date(timeStamp * 1000).getDate())) {
        return "?????? " + moment(timeStamp * 1000).format('HH:mm')
    }
    return moment(timeStamp * 1000).format('YYYY-MM-DD')
}

// ??????????????????
const onReadNotice = (id) => {
    updateNotice({ id: id }).then((res) => {
        if (res.data.code == 0) {
            onNotice()
            noticeCount()
        }
    })
}

// ??????????????????
const noticeCount = () => {
    getNoticeCount().then((res) => {
        if (res.data.code == 0) {
            data.noticeCount = res.data.data.count
        }
    })
}

// ??????????????????
const onNotice = () => {
    getNoticeList().then((res) => {
        if (res.data.code == 0) {
            data.noticeList = res.data.data
        }
    })
}

// ????????????
const onDeleteNotice = () => {
    let ids = []
    for (let index = 0; index < data.noticeList.length; index++) {
        ids[index] = data.noticeList[index].id
    }
    deleteNotice({ ids: ids }).then((res) => {
        if (res.data.code == 0) {
            data.noticeList = res.data.data
            onNotice()
        }
    })
}

// ??????????????????
const onLogout = () => {
    localStorage.removeItem("uid")
    localStorage.removeItem("token")
    router.push('/')
}

// ??????????????????
const onCancel = () => {
    disabled.value = false
    modalFormRef.value.resetFields()
    visible.value = false
}
</script>

<style scoped>
.layout-sider {
    left: 0;
    top: 0;
    bottom: 0;
    overflow: auto;
    height: 100vh;
    position: fixed;
    background: #fff;
    border-right: 0.5px solid #F0F2F5;
}

.header {
    padding: 0 12px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    background: #fff;
}

.trigger {
    font-size: 18px;
    padding: 0 8px;
    cursor: pointer;
    transition: color 0.3s;
}

.trigger:hover {
    color: #476FFF;
}

.logo {
    height: 32px;
    margin: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.avatar {
    color: #f56a00;
    background-color: #fde3cf;
    cursor: pointer;
    margin-left: 20px;
}

.popover {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 10px 0;
}

.title {
    font-size: 20px;
    color: rgba(31, 31, 31, 0.85);
    font-weight: 620;
    margin-left: 10px;
    overflow: hidden;
}
</style>