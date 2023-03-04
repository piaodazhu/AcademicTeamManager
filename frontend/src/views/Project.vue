<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
        <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
            <a-space>
                <a-input v-model:value="query.id" placeholder="项目编号" style="width: 250px; margin-right: 10px;">
                    <template #suffix>
                        <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch" />
                    </template>
                </a-input>
                <a-button :type="buttonType.bt1" @click="onSortProject(1)">全部项目</a-button>
                <a-button :type="buttonType.bt2" @click="onSortProject(2)">按完成度降序</a-button>
                <a-button :type="buttonType.bt3" @click="onSortProject(3)">按完成度升序</a-button>
                <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
            </a-space>
            <div>
                <a-space size="middle">
                    <a-button type="primary" @click="onCreate">新建项目</a-button>
                    <a-button type="primary" @click="onExport" ghost>
                        <template #icon>
                            <ExportOutlined />
                        </template>导出</a-button>
                </a-space>
            </div>
        </div>
        <a-table rowKey="id"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectedProjectIds, getCheckboxProps: defaultSelected }"
            :columns="columns" :data-source="data.projectList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'id'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'finishrate'">
                    <span style="color: #ff991f">{{ text }}</span>
                </template>
                <!-- <template v-if="column.dataIndex === 'finishrate'">
                    <Spot v-if="text == 1" type="success" title="已完成" />
                    <Spot v-if="text == 2" type="danger" title="未完成" />
                </template> -->
            </template>
        </a-table>
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" :centered="true">
            <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="projectFormRef" :model="project" layout="vertical" name="project" :rules="rules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="项目编号" name="id">
                                <a-input v-model:value="project.id" :disabled="true" placeholder="根据编号规则自动生成" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="项目名称" name="name">
                                <a-input v-model:value="project.name" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <!-- <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="学生名称" name="cid">
                                <a-select v-model:value="project.cid" style="width: 100%" placeholder="请选择"
                                    :fieldNames="{ label: 'name', value: 'id' }" :options="data.studentOption"
                                    @focus="getStudentOption" @change="changeStudentOption"></a-select>
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="合同金额" name="amount">
                                <a-input v-model:value="project.amount" style="width: 100%" :disabled="true" />
                            </a-form-item>
                        </a-col>
                    </a-row> -->
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="开始时间" name="beginTime">
                                <a-date-picker v-model:value="project.beginTime" placeholder="选择日期" style="width: 100%"
                                    format="YYYY-MM-DD" valueFormat="YYYY-MM-DD" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="结项时间" name="overTime">
                                <a-date-picker v-model:value="project.overTime" placeholder="选择日期" style="width: 100%"
                                    format="YYYY-MM-DD" valueFormat="YYYY-MM-DD" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="24">
                            <a-form-item label="备注" name="remark">
                                <a-input v-model:value="project.remark" />
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="24">
                            <a-form-item label="项目成果" name="output">
                                <a-button type="primary" @click="onAddOutput" style="float: right;margin-bottom: 10px;">
                                    添加成果</a-button>
                                <a-table rowKey="id" :columns="outputColumns" :data-source="data.outputList"
                                    :scroll="{ y: '59vh' }" class="ant-table-striped"
                                    :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)"
                                    :pagination="false" bordered>
                                    <template #bodyCell="{ column, text, record }">
                                        <template v-if="column.dataIndex === 'type'">
                                            <span v-if="text == 1">期刊论文</span>
                                            <span v-if="text == 2">会议论文</span>
                                            <span v-if="text == 3">学术专著</span>
                                            <span v-if="text == 4">发明专利</span>
                                        </template>
                                        <template v-if="column.dataIndex === 'weight'">
                                            <span style="color: #ff991f">{{ text }}</span>
                                        </template>
                                        <template v-if="column.dataIndex === 'status'">
                                            <Spot v-if="text == 0" type="danger" title="未设置" />
                                            <Spot v-if="text == 1" type="success" title="已完成" />
                                            <Spot v-if="text == 2" type="warning" title="推进中" />
                                        </template>
                                        <template v-if="column.dataIndex === 'operation'">
                                            <a @click="delOutput(record)">删除</a>
                                        </template>
                                    </template>
                                </a-table>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row>
                        <a-col :span="24">
                            <div style="float: right;margin: 0 20px;">
                                <span>已选择成果：{{ data.outputList.length }}种&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;总权重：{{
                                    project.weight.toFixed(2) }}</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;完成度：<span
                                    style="width: 80px; color: #ff991f">{{ project.finishrate.toFixed(2) }}</span>
                            </div>
                        </a-col>
                    </a-row>
                </a-form>
            </div>
        </a-modal>
        <!-- 添加产品 -->
        <a-modal v-model:visible="outputListVisible" title="添加成果" @cancel="onCancelOutputList" @ok="onConfirm"
            cancelText="取消" okText="确定" width="800px" style="top: 80px" :getCheckboxProps="defaultSelected">
            <div style="height: 55vh;padding: 0 15px;">
                <a-table rowKey="id"
                    :row-selection="{ selectedRowKeys: data.defaultSelectedIds, onChange: onSelectedOutputIds }"
                    :columns="outputListColumns" :data-source="data.outputList"
                    :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPaginationOutput }"
                    :scroll="{ y: '40vh' }" class="ant-table-striped"
                    :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
                    <template #bodyCell="{ column, text, record }">
                        <template v-if="column.dataIndex === 'name'">
                            <a>{{ text }}</a>
                        </template>
                        <template v-if="column.dataIndex === 'status'">
                            <Spot v-if="text == 0" type="danger" title="未设置" />
                            <Spot v-if="text == 1" type="success" title="已完成" />
                            <Spot v-if="text == 2" type="warning" title="推进中" />
                        </template>
                        <template v-if="column.dataIndex === 'type'">
                            <span v-if="text == 1">默认</span>
                        </template>
                        <template v-if="column.dataIndex === 'weight'">
                            <span style="color: #ff991f">{{ text }}</span>
                        </template>
                    </template>
                </a-table>
            </div>
        </a-modal>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, createVNode } from 'vue';
import { SearchOutlined, ExclamationCircleOutlined, ExportOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createProject, updateProject, queryProjectList, queryProjectInfo, deleteProject, queryProjectOlist, projectExport } from '../api/project';
import { queryOutputList } from "../api/output";
import { queryStudentOption } from "../api/student";
import { message, Modal } from 'ant-design-vue';
import Spot from '../components/Spot.vue';

const query = reactive({
    id: undefined,
    finishrate: undefined
})

const sorttype = reactive(
    {
        type: 1
    }
)

const onSearch = () => {
    projectList()
}

const onProjectList = () => {
    query.id = undefined
    query.status = undefined
    pagination.current = 1
    setButtonType()
    projectList()
}

// const onProjectStatus = (isFinish) => {
//     if (isFinish) {}
//     query.status = status
//     pagination.current = 1
//     setButtonType(status)
//     projectList()
// }

const onSortProject = (type) => {
    sorttype.type = type
    projectList()
}

// 按钮默认类型
const buttonType = reactive({
    bt1: 'primary',
    bt2: 'default',
    bt3: 'default',
})

// 设置按钮类型
const setButtonType = (status) => {
    switch (status) {
        case 1:
            buttonType.bt1 = 'default'
            buttonType.bt2 = 'primary'
            buttonType.bt3 = 'default'
            break;
        case 2:
            buttonType.bt1 = 'default'
            buttonType.bt2 = 'default'
            buttonType.bt3 = 'primary'
            break;
        default:
            buttonType.bt1 = 'primary'
            buttonType.bt2 = 'default'
            buttonType.bt3 = 'default'
            break;
    }
}

const columns = [{
    title: '项目编号',
    dataIndex: 'id',
    width: 80,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '项目名称',
    dataIndex: 'name',
    width: 150,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '开始时间',
    dataIndex: 'beginTime',
    width: 120,
}, {
    title: '结项时间',
    dataIndex: 'overTime',
    width: 120,
}, {
    title: '完成度',
    dataIndex: 'finishrate',
    width: 80,
}, {
    title: '备注',
    dataIndex: 'remark',
    width: 240,
    ellipsis: true,
}];

const outputColumns = [{
    title: '成果名称',
    dataIndex: 'name',
    width: 100,
}, {
    title: '成果类别',
    dataIndex: 'type',
    width: 90,
}, {
    title: '成果权重',
    dataIndex: 'weight',
    width: 60,
}, {
    title: '完成状态',
    dataIndex: 'status',
    width: 100,
}, {
    title: '操作',
    dataIndex: 'operation',
    width: 100,
}]

const outputListColumns = [{
    title: '成果名称',
    dataIndex: 'name',
    width: 100,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '成果类型',
    dataIndex: 'type',
    width: 100,
}, {
    title: '成果权重',
    dataIndex: 'weight',
    width: 60,
}, {
    title: '完成状态',
    dataIndex: 'status',
    width: 120,
}, {
    title: '成果描述',
    dataIndex: 'description',
    width: 240,
    ellipsis: true,
}];

// 表单校验
const rules = {
    name: [{ required: true, message: '请输入项目名称', trigger: 'blur' }],
    // cid: [{ required: true, message: '请选择客户', trigger: 'blur' }],
    // status: [{ required: true, message: '请选择项目状态' }]
};

// 合同属性
let project = reactive({
    id: undefined,
    name: undefined,
    // amount: undefined,
    beginTime: '',
    overTime: '',
    // cid: undefined,
    remark: undefined,
    // status: undefined,
    outputlist: [],
    weight: 0,
    finishrate: 0
});

const data = reactive({
    projectId: 0,
    projectList: [],
    projectIds: [],
    outputList: [],
    outputIds: [],
    // addedOutputList: [],
    studentOption: [],
    defaultSelectedIds: []
})

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
const projectFormRef = ref();
const outputListVisible = ref(false);

const onCreate = () => {
    title.value = '新建项目'
    operation.value = 1
    visible.value = true
    project.id = undefined
    project.name = undefined
    project.weight = 0.0
    project.outputlist = []
    // data.addedOutputList = []
    data.outputList = []
    data.weight = 0
}

const onEdit = (row) => {
    title.value = '编辑项目'
    operation.value = 2
    // getStudentOption()
    let param = { id: row.id }
    queryProjectInfo(param).then((res) => {
        if (res.data.code == 0) {
            let p = res.data.data
            project.id = p.id
            project.name = p.name
            // project.cid = p.cid
            // project.amount = p.amount
            project.beginTime = p.beginTime
            project.overTime = p.overTime
            project.remark = p.remark
            project.finishrate = p.finishrate
            // project.status = p.status
            if (p.outputlist == null) {
                project.outputlist = []
                // data.addedOutputList = []
                data.outputList = []
            } else {
                project.outputlist = p.outputlist
                // data.addedOutputList = p.outputlist
                data.outputList = p.outputlist
            }

            calculatedWeight()
        }
    })
    data.projectId = row.id
    visible.value = true
}

const onSave = () => {
    projectFormRef.value.validateFields().then(() => {
        if (operation.value == 1) {
            let param = {
                name: project.name,
                // cid: project.cid,
                // amount: project.amount,
                beginTime: project.beginTime,
                overTime: project.overTime,
                remark: project.remark,
                // status: project.status,
                // outputlist: data.addedOutputList,
                outputlist: data.outputList,
                finishrate: project.finishrate
            }
            createProject(param).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    data.defaultSelectedIds = []
                    projectList()
                }
            })
        }
        if (operation.value == 2) {
            let param = {
                id: project.id,
                name: project.name,
                // cid: project.cid,
                // amount: project.amount,
                beginTime: project.beginTime,
                overTime: project.overTime,
                remark: project.remark,
                // status: project.status,
                // outputlist: data.addedOutputList,
                outputlist: data.outputList,
                finishrate: project.finishrate
            }
            updateProject(param).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    data.defaultSelectedIds = []
                    projectList()
                }
            })
        }
        projectFormRef.value.resetFields()
        visible.value = false;
    });
};


const onDelete = () => {
    let param = {
        ids: data.projectIds
    }
    Modal.confirm({
        title: '确定删除选中的' + data.projectIds.length + '项吗?',
        icon: createVNode(ExclamationCircleOutlined),
        centered: true,
        cancelText: '取消',
        okText: '确定',
        onOk() {
            deleteProject(param).then((res) => {
                if (res.data.code == 0) {
                    projectList()
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

onMounted(() => { projectList() })

const onPagination = (page) => {
    pagination.current = page
    projectList()
}
const projectList = () => {
    let param = {
        // id: parseInt(query.id == undefined || query.id == '' ? '0' : query.id),
        // status: query.status,
        Name: project.name,
        pageNum: pagination.current,
        pageSize: pagination.pageSize,
        sorttype: sorttype.type
    }
    if (query.id != undefined && query.id != '') {
        param.id = parseInt(query.id)
    }
    queryProjectList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.projectList = res.data.data.list
        }
    })
}

const onAddOutput = () => {
    let param = {
        pageNum: pagination.current,
        pageSize: pagination.pageSize
    }
    queryOutputList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.outputList = res.data.data.list
        }
    })
    data.defaultSelectedIds = []
    if (data.outputList.length > 0) {
        for (let i = 0; i < data.outputList.length; i++) {
            data.defaultSelectedIds[i] = data.outputList[i].id
        }
    }
    outputListVisible.value = true
}

const onPaginationOutput = (page) => {
    pagination.current = page
    let param = {
        pageNum: pagination.current,
        pageSize: pagination.pageSize
    }
    queryOutputList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.outputList = res.data.data.list
        }
    })
}

const onSelectedProjectIds = selectedRowKeys => {
    data.projectIds = selectedRowKeys
    if (data.projectIds.length !== 0) {
        disabled.value = false
    } else {
        disabled.value = true
    }
};


const onSelectedOutputIds = selectedRowKeys => {
    data.outputIds = selectedRowKeys
    data.defaultSelectedIds = selectedRowKeys
    console.log(data.outputIds)
};


const delOutput = (row) => {
    for (let i = 0; i < data.outputList.length; i++) {
        if (data.outputList[i].id == row.id) {
            data.outputList.splice(i, 1);
        }
    }
    // console.log(data.outputList)
    calculatedWeight()
}

const onConfirm = () => {
    // console.log("ids", data.outputIds)
    // console.log(data.defaultSelectedIds)
    let param = {
        id: data.projectId,
        oids: data.outputIds
    }
    queryProjectOlist(param).then((res) => {
        if (res.data.code == 0) {
            data.outputList = res.data.data
            calculatedWeight()
        }
    })
    outputListVisible.value = false
}

// const getStudentOption = () => {
//     queryStudentOption().then((res) => {
//         if (res.data.code == 0) {
//             data.studentOption = res.data.data
//             console.log("opt", data.studentOption)
//         }
//     })
// }

// const changeStudentOption = (value) => {
//     project.cid.value = value
// }

const calculatedWeight = () => {
    project.weight = 0
    let totalWeight = 0.0
    let finishWeight = 0.0
    for (let i = 0; i < data.outputList.length; i++) {
        totalWeight = totalWeight + data.outputList[i].weight
        if (data.outputList[i].status == 1) {
            finishWeight += data.outputList[i].weight
        }
    }
    project.weight = totalWeight
    project.finishrate = finishWeight / totalWeight
}

// 点击合同取消按钮
const onCancel = () => {
    projectFormRef.value.resetFields()
    data.outputList = []
    data.projectId = undefined
    visible.value = false
};

// 导出表格
const onExport = () => {
    projectExport().then((res) => {
        if (res.data.type == 'application/json') {
            message.error('导出错误！')
        } else {
            let blob = new Blob([res.data], {
                type: "application/vnd.ms-excel"
            })
            let a = document.createElement('a')
            a.setAttribute("download", "项目信息.xlsx");
            a.href = window.URL.createObjectURL(blob)
            a.click()
            window.URL.revokeObjectURL(a.href)
        }
    })
}

const onCancelOutputList = () => {
    outputListVisible.value = false
    data.projectId = undefined
    pagination.current = 1,
        pagination.total = undefined
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>