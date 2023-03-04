<template>
    <div :style="{ padding: '20px 20px 12px 20px' }">
        <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
            <a-space>
                <a-input v-model:value="query.name" placeholder="成果名称" style="width: 280px; margin-right: 10px;">
                    <template #suffix>
                        <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="outputList()" />
                    </template>
                </a-input>
                <a-button :type="buttonType.bt1" @click="onOutputs">全部成果</a-button>
                <a-button :type="buttonType.bt2" @click="onFilter">
                    <template #icon>
                        <FilterOutlined />
                    </template>高级筛选</a-button>
                <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除成果</a-button>
            </a-space>
            <div>
                <a-space size="middle">
                    <a-button type="primary" @click="onCreate">新建成果</a-button>
                    <a-button type="primary" @click="onExport" ghost>
                        <template #icon>
                            <ExportOutlined />
                        </template>导出</a-button>
                </a-space>
            </div>
        </div>

        <a-modal v-model:visible="visibleFilter" title="高级筛选" @ok="confirmFilter" @cancel="cancelFilter" cancelText="取消"
            okText="确定" width="800px" style="top: 150px;">
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-select v-model:value="query.type" placeholder="成果类型" style="width: 100%;"
                        :allowClear="true">
                        <a-select-option :value="0">未设置</a-select-option>
                        <a-select-option :value="1">期刊论文</a-select-option>
                        <a-select-option :value="2">会议论文</a-select-option>
                        <a-select-option :value="3">学术专著</a-select-option>
                        <a-select-option :value="4">发明专利</a-select-option>
                    </a-select>
                </a-col>
                <a-col :span="6">
                    <a-select v-model:value="query.status" placeholder="完成状态" style="width: 100%;"
                        :allowClear="true">
                        <a-select-option :value="0">未设置</a-select-option>
                        <a-select-option :value="1">已完成</a-select-option>
                        <a-select-option :value="2">推进中</a-select-option>
                    </a-select>
                </a-col>
            </a-row>
        </a-modal>

        <a-table rowKey="id" :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :columns="columns" :data-source="data.outputList"
            :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
            :scroll="{ y: '59vh' }" class="ant-table-striped"
            :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'name'">
                    <a @click="onEdit(record)">{{ text }}</a>
                </template>
                <template v-if="column.dataIndex === 'status'">
                    <Spot v-if="text == 1" type="success" title="已完成" />
                    <Spot v-if="text == 2" type="danger" title="推进中" />
                </template>
                <template v-if="column.dataIndex === 'type'">
                    <span v-if="text == 1">期刊论文</span>
                    <span v-if="text == 2">会议论文</span>
                    <span v-if="text == 3">学术专著</span>
                    <span v-if="text == 4">发明专利</span>
                </template>
                <template v-if="column.dataIndex === 'weight'">
                    <span style="color: #ff991f">{{ text }}</span>
                </template>
            </template>
        </a-table>
        <!-- 新建、编辑产品 -->
        <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
            width="800px" :centered="true">
            <div style="height: 35vh;overflow-y: scroll;padding: 0 15px;">
                <a-form ref="outputFormRef" :model="output" layout="vertical" name="output" :rules="rules">
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="成果名称" name="name">
                                <a-input v-model:value="output.name" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="成果类型" name="type">
                                <a-select v-model:value="output.type" placeholder="请选择">
                                    <a-select-option :value="1">期刊论文</a-select-option>
                                    <a-select-option :value="2">会议论文</a-select-option>
                                    <a-select-option :value="3">学术专著</a-select-option>
                                    <a-select-option :value="4">发明专利</a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="12">
                            <a-form-item label="成果权重" name="weight">
                                <a-input-number v-model:value="output.weight" style="width: 100%" placeholder="请填写" />
                            </a-form-item>
                        </a-col>
                        <a-col :span="12">
                            <a-form-item label="完成状态" name="status">
                                <a-select v-model:value="output.status" placeholder="请选择">
                                    <a-select-option :value="1">
                                        <Spot type="success" title="已完成" />
                                    </a-select-option>
                                    <a-select-option :value="2">
                                        <Spot type="danger" title="推进中" />
                                    </a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                    </a-row>
                    <a-row :gutter="16">
                        <a-col :span="24">
                            <a-form-item label="成果描述" name="description">
                                <a-textarea v-model:value="output.description" :rows="8" />
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
import { SearchOutlined, ExclamationCircleOutlined, ExportOutlined } from '@ant-design/icons-vue';
import moment from 'moment'
import { createOutput, updateOutput, queryOutputList, deleteOutput, queryOutputInfo, outputExport } from '../api/output';
import Spot from '../components/Spot.vue';
import { message, Modal } from 'ant-design-vue';


const query = reactive({
    name: undefined,
    type: undefined,
    status: undefined,
})

// 表格字段
const columns = [{
    title: '成果名称',
    dataIndex: 'name',
    width: 100,
    fixed: 'left',
    ellipsis: true,
}, {
    title: '完成状态',
    dataIndex: 'status',
    width: 60,
}, {
    title: '成果类型',
    dataIndex: 'type',
    width: 50,
}, {
    title: '成果权重',
    dataIndex: 'weight',
    width: 50,
}, {
    title: '成果描述',
    dataIndex: 'description',
    width: 240,
    ellipsis: true,
}];


const rules = {
    name: [{ required: true, message: '请输入成果名称', trigger: 'blur' }],
    type: [{ required: true, message: '请选择成果类型' }],
    code: [{ pattern: /^\d+$/, message: '成果编号格式不正确', trigger: 'blur' }],
    weight: [{ required: true, message: '请输入产品权重' }],
    status: [{ required: true, message: '请选择完成状态' }]
};


let output = reactive({
    id: undefined,
    name: undefined,
    type: undefined,
    weight: undefined,
    status: undefined,
    description: undefined,
});


const data = reactive({
    outputList: [],
    selectedIds: []
})


let pagination = reactive({
    current: 1,
    pageSize: 10,
    total: undefined,
})

const title = ref('');
const visible = ref(false);
const disabled = ref(true)
const operation = ref(0);
const outputFormRef = ref();

// 按钮状态
const buttonType = reactive({
    bt1: 'primary',
    bt2: 'default',
})

const onOutputs = () => {
    buttonType.bt1 = 'primary'
    buttonType.bt2 = 'default'
    for (const key in query) {
        query[key] = undefined
    }
    outputList()
}

const visibleFilter = ref(false)

const onFilter = () => {
    buttonType.bt1 = 'default'
    buttonType.bt2 = 'primary'
    visibleFilter.value = true
}

const confirmFilter = () => {
    outputList()
    visibleFilter.value = false
}

// 初始化数据
onMounted(() => { outputList() })

// 表格记录多选
const onSelectChange = selectedRowKeys => {
    data.selectedIds = selectedRowKeys
    if (data.selectedIds.length !== 0) {
        disabled.value = false
    } else {
        disabled.value = true
    }
};

const onCreate = () => {
    title.value = '添加成果'
    operation.value = 1
    visible.value = true
}

const onEdit = (row) => {
    title.value = '编辑成果'
    operation.value = 2
    let param = { id: row.id }
    queryOutputInfo(param).then((res) => {
        if (res.data.code == 0) {
            let p = res.data.data
            output.id = p.id
            output.name = p.name
            output.type = p.type
            output.weight = p.weight
            output.status = p.status
            output.description = p.description
        }
    })
    visible.value = true
}

const onSave = () => {
    outputFormRef.value.validateFields().then(() => {
        if (operation.value == 1) {
            createOutput(output).then((res) => {
                if (res.data.code == 0) {
                    message.success('保存成功')
                    outputFormRef.value.resetFields()
                    visible.value = false;
                    outputList()
                }
                if (res.data.code == 40001) {
                    message.error('成果名称已经存在')
                }
            })
        }
        if (operation.value == 2) {
            updateOutput(output).then((res) => {
                if (res.data.code == 0) {
                    outputFormRef.value.resetFields()
                    visible.value = false;
                    message.success('保存成功')
                    outputList()
                }
            })
        }
    });
};

const onDelete = () => {
    Modal.confirm({
        title: '确定删除选中的' + data.selectedIds.length + '项吗?',
        icon: createVNode(ExclamationCircleOutlined),
        centered: true,
        cancelText: '取消',
        okText: '确定',
        onOk() {
            deleteOutput({ ids: data.selectedIds }).then((res) => {
                if (res.data.code == 0) {
                    outputList()
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


const onPagination = (page) => {
    pagination.current = page
    outputList()
}


const outputList = () => {
    // setButtonType(query.type)
    let param = {
        name: query.name,
        type: query.type,
        status: query.status,
        pageNum: pagination.current,
        pageSize: pagination.pageSize,
    }
    queryOutputList(param).then((res) => {
        if (res.data.code == 0) {
            pagination.total = res.data.data.total
            data.outputList = res.data.data.list
        }
    })
}


const onExport = () => {
    outputExport().then((res) => {
        if (res.data.type == 'application/json') {
            message.error('导出错误！')
        } else {
            let blob = new Blob([res.data], {
                type: "application/vnd.ms-excel"
            })
            let a = document.createElement('a')
            a.setAttribute("download", "成果信息.xlsx");
            a.href = window.URL.createObjectURL(blob)
            a.click()
            window.URL.revokeObjectURL(a.href)
        }
    })
}

// 点击取消按钮
const onCancel = () => {
    outputFormRef.value.resetFields()
    visible.value = false
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>