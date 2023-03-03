<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
        <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
            <a-space>
                <a-input v-model:value="query.name" placeholder="学生名称" style="width: 250px; margin-right: 10px;">
                    <template #suffix>
                        <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="studentList()" />
                    </template>
                </a-input>
                <a-button :type="buttonType.bt1" @click="onStudents">全体学生</a-button>
                <a-button :type="buttonType.bt2" @click="onFilter">
                    <template #icon>
                        <FilterOutlined />
                    </template>高级筛选</a-button>
                <a-button type="primary" @click="onDelete" :disabled="disabled" danger>移除成员</a-button>
            </a-space>
            <div>
                <a-space size="middle">
                    <a-button type="primary" @click="onCreate">新增成员</a-button>
                    <a-button type="primary" @click="onExport" ghost>
                        <template #icon>
                            <ExportOutlined />
                        </template>导出</a-button>
                </a-space>
            </div>
            <a-modal v-model:visible="visibleFilter" title="高级筛选" @ok="confirmFilter" @cancel="cancelFilter"
                cancelText="取消" okText="确定" width="800px" style="top: 150px;">
                <a-row :gutter="20">
                    <a-col :span="6">
                        <a-select v-model:value="query.type" :options="options.type" placeholder="学位类型"
                            style="width: 100%;" :allowClear="true"/>
                    </a-col>
                    <a-col :span="6">
                        <a-select v-model:value="query.status" :options="options.status" placeholder="研讨状态"
                            style="width: 100%;" :allowClear="true" />
                    </a-col>
                </a-row>
            </a-modal>
        </div>
        <a-table rowKey="id" :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :columns="columns" :data-source="data.studentList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'name'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'operation'">
                    <a-button type="link"><template #icon>
                            <PhoneTwoTone two-tone-color="#31C27C" @click="callUp(record.phone)" />
                        </template></a-button>
                    <a-button type="link"><template #icon>
                            <MailTwoTone two-tone-color="#476FFF" @click="onMail(record.name, record.email)" />
                        </template></a-button>
                </template>
                <template v-if="column.dataIndex === 'status'">
                    <Spot v-if="text == 0" type="warning" title="未设置" />
                    <Spot v-if="text == 1" type="success" title="已研讨" />
                    <Spot v-if="text == 2" type="danger" title="待研讨" />
                </template>
            </template>
        </a-table>

        <!-- 新建、编辑学生 -->
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" :centered="true">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="studentFormRef" :model="student" layout="vertical" name="student" :rules="rules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="姓名" name="name">
                                <a-input v-model:value="student.name" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="学位" name="type">
                                <a-select v-model:value="student.type" :options="options.type" placeholder="请选择" allowClear:false />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="手机号" name="phone">
                                <a-input v-model:value="student.phone" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="邮箱" name="email">
                                <a-input v-model:value="student.email" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="上次研讨" name="lastdis">
                                <a-date-picker v-model:value="student.lastdis" placeholder="选择日期" style="width: 100%"
                                    format="YYYY-MM-DD" valueFormat="YYYY-MM-DD" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="下次研讨" name="nextdis">
                                <a-date-picker v-model:value="student.nextdis" placeholder="选择日期" style="width: 100%"
                                    format="YYYY-MM-DD" valueFormat="YYYY-MM-DD" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="备注" name="remark">
                                <a-textarea v-model:value="student.remark" :auto-size="{ minRows: 3, maxRows: 3 }" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
        <!-- 发送邮件对话框 -->
        <a-modal v-model:visible="visibleMail" title="发送邮件" @ok="onSend" @cancel="onCancel" cancelText="取消" okText="发送"
            width="600px" :centered="true">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="mailSendFormRef" :model="mail" layout="vertical" :rules="mailRules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="学生姓名" name="studentName">
                                <a-input v-model:value="mail.studentName" disabled />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="收件邮箱" name="receiver">
                                <a-input v-model:value="mail.receiver" disabled />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="邮件主题" name="subject">
                                <a-input v-model:value="mail.subject" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="邮件内容" name="content">
                                <a-textarea v-model:value="mail.content" placeholder="邮件正文支持HTML语法"
                                    :auto-size="{ minRows: 6, maxRows: 100 }" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="24">
                            <a-form-item label="上传附件" name="attachment">
                                <a-upload-dragger v-model:fileList="fileList" name="file" :multiple="true"
                                    :headers="{ 'X-Requested-With': null }" :action="action" @change="upload"
                                    @drop="upload" @remove="remove">
                                    <p class="ant-upload-drag-icon">
                                        <inbox-outlined></inbox-outlined>
                                    </p>
                                    <p style="font-size: 14px;color: #00000073;">单击或拖动文件到此区域</p>
                                </a-upload-dragger>
                            </a-form-item>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, createVNode } from 'vue';
import { SearchOutlined, ExclamationCircleOutlined, FilterOutlined, ExportOutlined, PhoneTwoTone, MailTwoTone, InboxOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createStudent, updateStudent, sendMailToStudent, queryStudentList, queryStudentInfo, deleteStudent, studentExport } from '../api/student';
import { message, Modal } from 'ant-design-vue';
import Spot from '../components/Spot.vue';
import { fileRemove } from '../api/common';
import regionData from '../assets/region';

// 条件筛选
const query = reactive({
    name: undefined,
    type: undefined,
    status: undefined,
})


const buttonType = reactive({
    bt1: 'primary',
    bt2: 'default',
})


const onStudents = () => {
    buttonType.bt1 = 'primary'
    buttonType.bt2 = 'default'
    for (const key in query) {
        query[key] = undefined
    }
    studentList()
}

const visibleFilter = ref(false)


const onFilter = () => {
    buttonType.bt1 = 'default'
    buttonType.bt2 = 'primary'
    visibleFilter.value = true
}


const confirmFilter = () => {
    studentList()
    visibleFilter.value = false
}


const cancelFilter = () => {
    buttonType.bt1 = 'primary'
    buttonType.bt2 = 'default'
    for (const key in query) {
        query[key] = undefined
    }
}


const options = reactive({
    type: [{
        value: 1,
        label: '博士',
    }, {
        value: 2,
        label: '硕士',
    }, {
        value: 3,
        label: '本科',
    }],
    status: [{
        value: 0,
        label: '未设置',
    },{
        value: 1,
        label: '已讨论',
    }, {
        value: 2,
        label: '待讨论',
    }],
    regionData
})


const columns = [{
    title: '姓名',
    dataIndex: 'name',
    width: 150,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '学位',
    dataIndex: 'type',
    width: 80,
    customRender: type => {
        if (type.value == 1) {
            return "博士"
        } else if (type.value == 2) {
            return "硕士"
        } else if (type.value == 3) {
            return "本科"
        } else {
            return "未知"
        }
    }
}, {
    title: '手机号',
    dataIndex: 'phone',
    width: 150,
}, {
    title: '邮箱',
    dataIndex: 'email',
    width: 200,
}, {
    title: '上次讨论',
    dataIndex: 'lastdis',
    width: 120,
    // customRender: text => {
    //     return text == 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD')
    // }
}, {
    title: '下次讨论',
    dataIndex: 'nextdis',
    width: 120,
    // customRender: text => {
    //     return text == 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD')
    // }
}, {
    title: '状态',
    dataIndex: 'status',
    width: 100,
}, {
    title: '备注',
    dataIndex: 'remark',
    width: 150,
    ellipsis: true,
}, {
    title: '操作',
    dataIndex: 'operation',
    width: 100,
    fixed: 'right',
    ellipsis: true,
}];

// 表单校验
const rules = {
    name: [{ required: true, message: '请输入学生名称', trigger: 'blur' }],
    phone: [{
        pattern: /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/,
        message: '手机格式不正确',
        trigger: 'blur',
    }],
    email: [{
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z.]{2,8})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }],
};

const data = reactive({
    studentList: [],
    selectedIds: []
})


const onSelectChange = selectedRowKeys => {
    data.selectedIds = selectedRowKeys
    if (data.selectedIds.length !== 0) {
        disabled.value = false
    } else {
        disabled.value = true
    }
};

// 学生属性
let student = reactive({
    id: undefined,
    name: undefined,
    type: undefined,
    phone: undefined,
    email: undefined,
    remark: undefined,
    lastdis: '',
    nextdis: '',
    status: undefined,
});

// 表格分页
let pagination = reactive({
    current: 1,
    pageSize: 10,
    total: undefined,
})

const title = ref('');
const visible = ref(false);
const disabled = ref(true)
const operation = ref(0);
const studentFormRef = ref();
const visibleMail = ref(false)

// 点击新建学生
const onCreate = () => {
    title.value = '新建学生'
    operation.value = 1
    visible.value = true
}

// 点击学生名称
const onEdit = (row) => {
    title.value = '编辑学生'
    operation.value = 2
    let param = { id: row.id }
    queryStudentInfo(param).then((res) => {
        if (res.data.code == 0) {
            let p = res.data.data
            student.id = p.id
            student.name = p.name
            student.phone = p.phone
            student.email = p.email
            student.type = p.type
            student.remark = p.remark
            student.status = p.status
            student.nextdis = p.nextdis
            student.lastdis = p.lastdis
        }
    })
    visible.value = true
}

// 点击保存学生
const onSave = () => {
    studentFormRef.value.validateFields().then(() => {
        // student.lastdis = new Date(student.lastdis).getTime() / 1000;
        // student.nextdis = new Date(student.nextdis).getTime() / 1000;
        var now = new Date()
        var next = new Date(student.nextdis)
        if (next.getTime() < now.getTime()) {
            student.status = 1 
        } else {
            student.status = 2 
        }
        if (operation.value == 1) {
            createStudent(student).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    studentFormRef.value.resetFields()
                    visible.value = false;
                    studentList()
                }
                if (res.data.code == 20001) {
                    message.error('学生姓名已经存在')
                }
            })
        }
        if (operation.value == 2) {
            updateStudent(student).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    studentFormRef.value.resetFields()
                    visible.value = false;
                    studentList()
                }
            })
        }
    });
};

// 点击删除学生
const onDelete = () => {
    let param = {
        ids: data.selectedIds
    }
    Modal.confirm({
        title: '确定删除选中的' + data.selectedIds.length + '项吗?',
        icon: createVNode(ExclamationCircleOutlined),
        centered: true,
        cancelText: '取消',
        okText: '确定',
        onOk() {
            deleteStudent(param).then((res) => {
                if (res.data.code == 0) {
                    studentList()
                    disabled.value = true
                    message.success('删除成功')
                }
            })
        },
        onCancel() {
            console.log('Cancel');
        },
    });
}

// 分页查询学生列表
const onPagination = (page) => {
    pagination.current = page
    studentList()
}

// 初始化数据
onMounted(() => { studentList() })

const studentList = () => {
    let param = {
        name: query.name,
        type: query.type,
        status: query.status,
        pageNum: pagination.current,
        pageSize: pagination.pageSize,
    }
    queryStudentList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.studentList = res.data.data.list
        }
    })
}

// 导出表格
const onExport = () => {
    studentExport().then((res) => {
        if (res.data.type == 'application/json') {
            message.error('导出错误！')
        } else {
            let blob = new Blob([res.data], {
                type: "application/vnd.ms-excel"
            })
            let a = document.createElement('a')
            a.setAttribute("download", "学生信息.xlsx");
            a.href = window.URL.createObjectURL(blob)
            a.click()
            window.URL.revokeObjectURL(a.href)
        }
    })
}

// 打电话
const callUp = (phone) => {
    window.location.href = 'tel://' + phone
}

const mail = reactive({
    studentName: '',
    receiver: '',
    subject: '',
    content: '',
    attachment: undefined
})

const mailSendFormRef = ref()

// 邮件发送表单校验
const mailRules = {
    receiver: [{
        required: true,
        pattern: /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z.]{2,8})$/,
        message: '邮箱格式不正确',
        trigger: 'blur',
    }],
    subject: [{ required: true, message: '请输入邮件主题', trigger: 'blur' }],
    content: [{ required: true, message: '请输入邮件内容', trigger: 'blur' }],
};

// 点击邮件
const onMail = (cname, email) => {
    mail.studentName = cname
    mail.receiver = email
    visibleMail.value = true
}

// 上传附件
const action = ref(import.meta.env.VITE_FILE_UPLOAD_URL)
const upload = (file) => {
    if (file.file.status == 'done') {
        if (file.file.response.code == 0) {
            mail.attachment = file.file.response.data.url
            fileName.value = file.file.response.data.name
        } else {
            message.error('附件上传失败')
        }
    }
}

// 移除附件
const fileName = ref(undefined)
const remove = (file) => {
    if (file.status == 'done') {
        fileRemove({ name: fileName.value })
    }
}

// 点击发送邮件
const onSend = () => {
    mailSendFormRef.value.validateFields().then(() => {
        let param = {
            receiver: mail.receiver,
            subject: mail.subject,
            content: mail.content,
            attachment: mail.attachment
        }
        sendMailToStudent(param).then((res) => {
            if (res.data.code == 0) {
                message.success("邮件已发送")
                visibleMail.value = false
            }
            if (res.data.code == 50002) {
                message.error("邮件发送失败")
            }
            if (res.data.code == 50003) {
                message.warn("邮件服务未开启")
            }
        })
    })
}

// 点击取消按钮
const onCancel = () => {
    studentFormRef.value.resetFields()
    visible.value = false
}

// const selectedOptions = (value) => {
//     student.region = value
// }
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>